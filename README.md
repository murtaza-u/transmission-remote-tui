# Transmission-Remote-Tui (trt)
> A TUI for BitTorrent client transmission

## Installation
- Build dependencies: `go`

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
| K          | move torrent up the queue                   |
| J          | move torrent down the queue                 |
| U          | move torrent at the top of the queue        |
| D          | move torrent at the bottom of the queue     |
| p          | pause/start torrent                         |
| r          | remove torrent                              |
| R          | remove torrent and delete all the files     |
| v          | verify torrent                              |
| t          | ask trackers for more peers                 |

- Changing file's priority

| keybinding | Action                                      |
|------------|---------------------------------------------|
| i, d       | increase/decrease file priority             |
| o          | change priority of focused file to 'off'    |
| l          | change priority of focused file to 'low'    |
| n          | change priority of focused file to 'normal' |
| h          | change priority of focused file to 'high'   |
| O          | change priority of all files to 'off'       |
| L          | change priority of all files to 'low'       |
| N          | change priority of all files to 'normal'    |
| H          | change priority of all files to 'high'      |

## Uninstalling
```bash
$ sudo make clean uninstall
```
