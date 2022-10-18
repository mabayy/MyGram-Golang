package main

import "final-projek/router"

func main() {
	r := router.StartApp()
	r.Run(":3000")
}
