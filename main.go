package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

// DONE: Load apps
// TODO: Update system
// TODO: Install system apps
// TODO: Install node apps
// TODO: Install python apps
// TODO: Config .zshrc
type jsonData map[string][]string

func main() {
	apps := getApps()

	// Update system
	update()

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

	switch system {
	case "system":
		installer := getInstaller()
		cmd = installer + unpack(list)
	case "node":
		cmd = "npm i -g " + unpack(list)
	case "python":
		cmd = "pip install " + unpack(list)
	}

	run(cmd)
}

// Update system
func update() {
	run("yaourt -Syu --noconfirm ")
}

// Load config
func loadConfig() {
	run(strings.TrimSpace(`
		git clone https://github.com/marco-souza/zshrc.git &&
    cd zshrc &&
    ./apply.sh &&
    cd -
	`))
}

// Get installer
func getInstaller() string {
	// if runtime.GOOS == "darwin" {
	// 	return "brew install cask "
	// }
	return "yaourt -Sy --noconfirm "
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
