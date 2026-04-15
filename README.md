#Updarch - An Arch linux update TUI

## Dependencies
- [paru](https://github.com/morganamilo/paru)

Built with Bubble Tea.
Description

updarch lists all pending package updates from your AUR helper paru, lets you choose which packages to update, and then runs paru -S <selected> to apply the updates. This provides a convenient, keyboard-driven way to avoid mass updates and only upgrade what you need.
Features

    - Fetches available updates using paru -Qu

    - Full keyboard‑driven TUI (Vim‑style or arrow keys)

    - Select/deselect individual packages with Space

    - Pagination for large update lists (Page Up/Down via ←/→ or h/l)

    - Executes paru on the chosen packages and quits

    - Automatically adapts to terminal size

## Installation
### Prerequisites

    - Arch Linux (or derivative like EndeavourOS, Manjaro)

    - paru installed and configured

    - Go 1.25+ (to build from source)

### Build from source

` git clone https://github.com/yourusername/updarch.git
cd updarch
go build -o updarch
`

Optionally move the binary to your PATH:


## Usage

### Keyboard controls
|Key|	Action|
|----|-----------|
|↑ / k|Move cursor up|
|↓ / j|Move cursor down|
|← / h|Previous page|
|→ / l|Next page|
|Space|Select / deselect package|
|Enter|Confirm selection and run -S|
|q / Ctrl+c|Quit without updating|

## [!CAUTION]

> Using updarch can lead your system into a partial upgrade state, which is dangerous.

> Partial upgrades occur when you update only a subset of packages without updating the entire system. This often causes:

    Broken dependencies

    Mismatched shared library versions

    Unpredictable application crashes

    System instability or unbootable states


### Why is this risky?

Package updates frequently assume that all other packages are at the latest version. If you skip some updates, critical libraries (e.g., glibc, gcc‑libs, openssl) may remain outdated while other packages expect newer symbols – leading to severe breakage.

### Recommended safe practices

   - **Prefer full system updates with paru -Syu or sudo pacman -Syu**

   -  **Use updarch only when you understand the implications and for packages that are completely independent (e.g., user‑space applications with no system library dependencies)**

**The author of updarch is not responsible for any system breakage resulting from partial updates. You have been warned.**

