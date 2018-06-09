package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"runtime"
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
}

func install(system string, list []string) {

	if len(list) == 0 {
		return
	}

	var cmd string

	switch system {
	case "system":
		installer := "yaourt -Syu --noconfirm "
		if runtime.GOOS == "darwin" {
			installer = "brew install cask "
		}

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
	run("brew cask upgrade")
}

// Load config
func loadConfig() {
	run(`
		git clone https://github.com/marco-souza/zshrc.git &&
    cd zshrc &&
    ./apply.sh &&
    cd -
	`)
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
