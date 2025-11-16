# viddy-go

<p align="center">
<img src="images/logo.png" width="200" alt="viddy" title="viddy" />
</p>

An interactive `watch` command.

_Viddy well, gopher. Viddy well._

## Fork notice

This is a fork of the last Go commit of [sachaos/viddy](https://github.com/sachaos/viddy) before it was ported to Rust.
You should use the Rust version of Viddy by default.
The Rust version is more stable and better maintained.
Use viddy-go when you want to take advantage of the Go toolchain for cross-compilation or want a specific feature like `--pty`.

## Features

* Basic features of the original watch command:
    * Execute a command periodically and display its result.
    * Colorized output.
    * Diff highlighting.
* Time machine mode:
    * Rewind like a video.
    * Go to the past and back to the future.
* View output in a pager.
* Vim-like keymaps.
* Search for text.
* Suspend and restart execution.
* Force commands to run at precise intervals.
* Shell alias support.
* Customizable key mappings.
* Customizable colors.

## Requirements

- Go 1.24 or later

## Install

### Go

```shell
go install dbohdan.com/viddy-go@latest
```

### Other

Download binaries from the [release page](https://github.com/dbohdan/viddy-go/releases).

## Keymap

| Key       |                                                |
|-----------|------------------------------------------------|
| Space     | Toggle time machine mode                       |
| b         | Toggle terminal <ins>b</ins>ell ringing        |
| d         | Toggle <ins>d</ins>iff                         |
| f         | Toggle <ins>f</ins>old                         |
| s         | Toggle <ins>s</ins>uspend execution            |
| t         | Toggle header/<ins>t</ins>itle display         |
| ?         | Toggle help view                               |
| /         | Search text                                    |
| j         | Pager: Move to next line                       |
| k         | Pager: Move to previous line                   |
| Control-F | Pager: Page down                               |
| Control-B | Pager: Page up                                 |
| g         | Pager: Go to top                               |
| Shift-G   | Pager: Go to bottom                            |
| Shift-J   | (Time machine mode) Go to the past             |
| Shift-K   | (Time machine mode) Back to the future         |
| Shift-F   | (Time machine mode) Go further into the past   |
| Shift-B   | (Time machine mode) Go further into the future |
| Shift-O   | (Time machine mode) Go to oldest position      |
| Shift-N   | (Time machine mode) Go to current position     |

## Configuration

Create your config file at `$XDG_CONFIG_HOME/viddy-go/config.toml`.
On macOS, the path is `~/Library/Application Support/viddy-go/config.toml`.

```toml
[general]
no_shell = false
shell = "zsh"
shell_options = ""
skip_empty_diffs = false

[keymap]
timemachine_go_to_past = "Down"
timemachine_go_to_more_past = "Shift-Down"
timemachine_go_to_future = "Up"
timemachine_go_to_more_future = "Shift-Up"
timemachine_go_to_now = "Ctrl-Shift-Up"
timemachine_go_to_oldest = "Ctrl-Shift-Down"

[color]
background = "white"  # Default value uses terminal color.
```

## What is "viddy"?

"Viddy" is a nadsat word meaning "to see".
Nadsat is the fictional argot of teenage gangs in the dystopian book and movie [_A Clockwork Orange_](https://en.wikipedia.org/wiki/A_Clockwork_Orange_(novel)).

## Credits

The gopher logo for Viddy is licensed under the Creative Commons 3.0 Attribution license.

The original Go gopher was designed by [Renée French](https://en.wikipedia.org/wiki/Ren%C3%A9e_French).
