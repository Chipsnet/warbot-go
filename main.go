package main

import (
	"github.com/go-sharp/color"
)

// Version バージョン情報
const Version = "1.0.0"

func main() {
	color.Yellow(` _       _____    ____  ____        __             
| |     / /   |  / __ \/ __ )____  / /_ ____ _____ 
| | /| / / /| | / /_/ / __  / __ \/ __// __ / __ \
| |/ |/ / ___ |/ _, _/ /_/ / /_/ / /__/ /_/ / /_/ /
|__/|__/_/  |_/_/ |_/_____/\____/\__(_)__, /\____/ 
				    /____/`)

	color.Cyan("Version: " + Version)
}
