package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
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
	// fmt.Println()
	// apps := getApps()
	// listApps := unpack(apps["system"])
	// fmt.Println(listApps)

	// Update system
	// update()

	// Install each system
	// for i, v := range apps {
	// 	install(i, v)
	// }

	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)

}

func install(system string, list []string) {

	if len(list) == 0 {
		return
	}

	var cmd *exec.Cmd
	switch system {
	case "system":
		args := " -Sy --noconfirm " + unpack(list)
		cmd = exec.Command("yaourt", args)
		// fmt.Println("yaourt" + args)

	case "node":
		args := " i -g " + unpack(list)
		cmd = exec.Command("npm", args)

	case "python":
		args := " install " + unpack(list)
		cmd = exec.Command("pip", args)
	}
	// err := cmd.Start()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	fmt.Printf("%s \n", system)
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s \n", out)
		log.Fatal(err)
	}
	fmt.Printf("%s \n", out)
}

// Unpack list to string
func unpack(list []string) string {
	result := ""
	for _, v := range list {
		result += fmt.Sprintf(" %s", v)
	}
	return strings.TrimSpace(result)
}

// Update system
func update() {
	exec.Command("yaourt", "-Syu")
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
