package main

import (
	"dating-app/cmd"
)

func main() {
	cmd.Execute().WithGracefulShutdown()
}
