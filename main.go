package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"

)


Func main() {
	discord, err := discordgo.new("")
	if err != nil {
		fmt.Println("Could not start session")

	}
	fmt.Println("Session:", discord)
}
