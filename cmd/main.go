package main

import "dating-app/src/app"

func main() {
	app.Execute().WithGracefulShutdown()
}
