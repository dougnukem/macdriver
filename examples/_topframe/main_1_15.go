// +build !go1.16

package main

import (
	"log"
)

func main() {
	log.Printf("Embed not supported in go < 1.16")
}
