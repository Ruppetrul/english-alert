package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
)

const appTitle = "English ALERT"

func main() {
	if 1 != len(os.Args) {
		handleArguments()
	}

	dictPath := filepath.Join(filepath.Dir(os.Args[0]), "dictionary.json")

	dictFile, err := os.Open(dictPath)
	if err != nil {
		log.Printf("Error opening dictionary dictFile: %v", err)
		log.Fatal(err)
	}

	dictFileContent, err := io.ReadAll(dictFile)
	if err != nil {
		log.Printf("Error read dictionary dictFile: %v", err)
		log.Fatal(err)
	}

	var dictionary map[string]any
	err = json.Unmarshal(dictFileContent, &dictionary)
	if err != nil {
		log.Fatal(err)
	}

	wordsCount := len(dictionary)
	if 0 == wordsCount {
		log.Fatal("dictionary is empty")
	}

	randomPos := rand.Intn(wordsCount)

	pos := 0
	for eng, rus := range dictionary {
		if pos != randomPos {
			pos++
			continue
		}
		cmd := exec.Command("notify-send", appTitle, fmt.Sprintf("%s - %s", eng, rus))
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		break
	}
}

func handleArguments() {
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		os.Exit(1)
	}

	handleAction(os.Args[1])

	//Run program with any param need only for configure.
	os.Exit(0)
}

func handleAction(action string) {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	config, err := NewServiceConfig("", execPath, filepath.Dir(execPath))
	if err != nil {
		log.Fatal(err)
	}

	switch action {
	case "enable":
		if err = enableService(config); err != nil {
			log.Fatal(err)
		}
	case "disable":
		if err = disableService(config); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("Unknown command. Use `enable`")
	}
}
