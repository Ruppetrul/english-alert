package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
)

func main() {
	file, err := os.Open("dictionary.json")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	var dictionary map[string]any
	err = json.Unmarshal(fileContent, &dictionary)

	if err != nil {
		log.Fatal(err)
	}

	dictionarySize := len(dictionary)
	if dictionarySize == 0 {
		log.Fatal("dictionary is empty")
	}

	randomPos := rand.Intn(dictionarySize)

	pos := 0
	for eng, rus := range dictionary {
		if pos == randomPos {
			cmd := exec.Command("notify-send", "English TIME!", fmt.Sprintf("%s - %s", eng, rus))
			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Notification successfully sent")
			}
			break
		}
		pos++
	}
}
