package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// DONE: Load apps
// DONE: Update system
// DONE: Install system apps
// DONE: Install node apps
// DONE: Install python apps
// DONE: Config .zshrc
type jsonData map[string][]string

func main() {
	apps := getApps()

	// Update system
	update()

	// Install Dependencies
	dependencies()

	// Load config
	loadConfig()

	// Install each system
	for i, v := range apps {
		install(i, v)
	}

	start()
}

func install(system string, list []string) {

	if len(list) == 0 {
		return
	}

	var cmd string

	installer := getInstaller(system)
	cmd = installer + unpack(list)

	if system != "1" {
		cmd = onZsh(cmd)
	}

	run(cmd)
}

// Update system
func update() {
	run("yaourt -Syu --noconfirm")
}

// Install Dependencies
func dependencies() {
	// TODO: Install yaourt (the hard way)
	run("sudo pacman -Sy --noconfirm git zsh")
}

// Load config
func loadConfig() {
	dest := "/tmp/zshrc"
	run("git clone https://github.com/marco-souza/zshrc.git " + dest)
	run("cd " + dest)
	run("./apply.sh")
	run("cd -")
}

// Get installer - Each installer is represented by a number
func getInstaller(system string) (result string) {
	switch system {
	case "1":
		result = "yaourt -Sy --noconfirm "
	case "2":
		result = "npm i -g "
	case "3":
		result = "pip install "
	}
	return
}

// Wrap command on zsh, to allow local configs setted on zsh
// like GOPATH, npm global repo, and so on
func onZsh(cmd string) string {
	return fmt.Sprintf("zsh -c '%s'", cmd)
}

// Get apps data
func getApps() jsonData {
	raw, err := ioutil.ReadFile("data/apps.json")
	if err != nil {
		log.Fatal(err)
	}

	var apps jsonData
	err = json.Unmarshal(raw, &apps)
	if err != nil {
		log.Fatal(err)
	}
	return apps
}
