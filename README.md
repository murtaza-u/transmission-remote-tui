# Transmission-Remote-Tui (trt)
> A TUI for BitTorrent client transmission

## Installation
```bash
$ git clone https://github.com/Murtaza-Udaipurwala/transmission-remote-tui
$ cd transmission-remote-tui
$ make
$ sudo make install
```

## Usage
- Transmission daemon must be running
```bash
$ transmission-daemon &
```

- Navigation
| keybinding | Action                                      |
|------------|---------------------------------------------|
| h, j, k, l | move around                                 |
| g          | scroll the to top of the page               |
| G          | scroll the to bottom of the page            |
| q          | quit / go back                              |
| Q          | kill the transmission daemon                |
| l, enter   | show more details about a torrent           |
| h, l       | increase/decrease file priority             |
| K          | move torrent up the queue                   |
| J          | move torrent down the queue                 |
| U          | move torrent at the top of the queue        |
| D          | move torrent at the bottom of the queue     |
| p          | pause/start torrent                         |
| r          | remove torrent                              |
| R          | remove torrent and delete all the files     |
| v          | verify torrent                              |

## Uninstalling
```bash
$ sudo make clean uninstall
```
