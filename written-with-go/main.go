package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var towns = []string{"domates", "patates", "elma", "armut"}

func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Println(`Nothingness - a role-playing game on command line
use help command to show help page`)
	gameLoop()
}

func gameLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	var command string
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			command = scanner.Text()
			response := runCommand(command)
			fmt.Println(response)
		}
	}
}

func runCommand(command string) string {
	commands := strings.Split(command, " ")
	switch commands[0] {
	case "help":
		return helpPage()
	case "status":
		return statusPage()
	case "map":
		return mapPage()
	case "go":
		return goPage(commands)
	case "show":
		return showPage(commands[1])
	case "fight":
		return fightPage(commands[1])
	case "exit":
		os.Exit(0)
		return ""
	default:
		return invalidCommandPage()
	}
}

func helpPage() string {
	response := `Help
	help                    -> show help page
	status                  -> show status of your hero
	map                     -> show the map
	go <town-name>          -> go to given name of town
	show everyone | build   -> show everyone | build in current town
	fight <monster-name>    -> fight to given name of monster
	exit                    -> exit from the game
`
	return response
}

func statusPage() string {
	return "this is status page"
}

func mapPage() string {
	return "this is map page"
}

func goPage(commands []string) string {
	if len(commands) != 2 {
		return invalidCommandPage()
	} else if !inArray(strings.ToLower(commands[1]), towns) {
		return "There is not any town that given name\n" + mapPage()
	}
	return commands[1]
}

func showPage(thing string) string {
	return thing
}

func fightPage(monsterName string) string {
	return "fighting page"
}

func invalidCommandPage() string {
	return "The command is invalid!\n" + helpPage()
}
