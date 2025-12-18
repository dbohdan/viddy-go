package main

import (
	"context"
	"time"
)

type newSnapFunc func(int64, *Snapshot, chan<- struct{}) *Snapshot

func ClockSnapshot(ctx context.Context, begin int64, newSnap newSnapFunc, interval time.Duration) (<-chan *Snapshot, chan<- bool) {
	c := make(chan *Snapshot)
	isSuspendedQueue := make(chan bool)

	go func() {
		defer close(c)
		var s *Snapshot
		var isSuspended bool

		t := time.NewTicker(interval)
		defer t.Stop()

		for {
			select {
			case <-ctx.Done():
				return

			case suspend := <-isSuspendedQueue:
				isSuspended = suspend

			case now := <-t.C:
				if isSuspended {
					continue
				}

				finish := make(chan struct{})
				id := (now.UnixNano() - begin) / int64(time.Millisecond)

				s = newSnap(id, s, finish)
				select {
				case c <- s:

				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return c, isSuspendedQueue
}

func PreciseSnapshot(ctx context.Context, newSnap newSnapFunc, interval time.Duration) (<-chan *Snapshot, chan<- bool) {
	c := make(chan *Snapshot)
	isSuspendedQueue := make(chan bool)

	go func() {
		defer close(c)
		var s *Snapshot
		var isSuspended bool

		begin := time.Now().UnixNano()

		for {
			select {
			case <-ctx.Done():
				return

			case suspend := <-isSuspendedQueue:
				isSuspended = suspend

			default:
			}

			if isSuspended {
				select {
				case <-ctx.Done():
					return

				case <-time.After(interval):
					continue
				}
			}

			finish := make(chan struct{})
			start := time.Now()
			id := (start.UnixNano() - begin) / int64(time.Millisecond)
			ns := newSnap(id, s, finish)
			s = ns

			select {
			case c <- ns:

			case <-ctx.Done():
				return
			}

			select {
			case <-finish:

			case <-ctx.Done():
				return
			}

			pTime := time.Since(start)

			if pTime > interval {
				continue
			}

			select {
			case <-ctx.Done():
				return

			case <-time.After(interval - pTime):
			}
		}
	}()

	return c, isSuspendedQueue
}

func SequentialSnapshot(ctx context.Context, newSnap newSnapFunc, interval time.Duration) (<-chan *Snapshot, chan<- bool) {
	c := make(chan *Snapshot)
	isSuspendedQueue := make(chan bool)

	go func() {
		defer close(c)
		var s *Snapshot
		var isSuspended bool

		begin := time.Now().UnixNano()

		for {
			select {
			case <-ctx.Done():
				return

			case suspend := <-isSuspendedQueue:
				isSuspended = suspend

			default:
			}

			if isSuspended {
				select {
				case <-ctx.Done():
					return

				case <-time.After(interval):
					continue
				}
			}

			finish := make(chan struct{})
			id := (time.Now().UnixNano() - begin) / int64(time.Millisecond)

			s = newSnap(id, s, finish)

			select {
			case c <- s:

			case <-ctx.Done():
				return
			}

			select {
			case <-finish:

			case <-ctx.Done():
				return
			}

			select {
			case <-ctx.Done():
				return

			case <-time.After(interval):
			}
		}
	}()

	return c, isSuspendedQueue
}
