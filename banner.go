package main

import "log"

const strBanner = `
        _             _             _                 _          _            _
       / /\         /\ \           / /\              /\ \       /\ \         /\ \
      / /  \       /  \ \         / /  \             \_\ \      \_\ \       /  \ \
     / / /\ \__   / /\ \ \       / / /\ \            /\__ \     /\__ \     / /\ \ \
    / / /\ \___\ / / /\ \ \     / / /\ \ \          / /_ \ \   / /_ \ \   / / /\ \_\
    \ \ \ \/___// / /  \ \_\   / / /  \ \ \        / / /\ \ \ / / /\ \ \ / / /_/ / /
     \ \ \     / / /    \/_/  / / /___/ /\ \      / / /  \/_// / /  \/_// / /__\/ /
 _    \ \ \   / / /          / / /_____/ /\ \    / / /      / / /      / / /_____/
/_/\__/ / /  / / /________  / /_________/\ \ \  / / /      / / /      / / /\ \ \
\ \/___/ /  / / /_________\/ / /_       __\ \_\/_/ /      /_/ /      / / /  \ \ \
 \_____\/   \/____________/\_\___\     /____/_/\_\/       \_\/       \/_/    \_\/

`

func banner() {
	log.Printf(strBanner)
}
