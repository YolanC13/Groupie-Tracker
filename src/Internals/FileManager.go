package groupieTracker

import (
	"bufio"
	"fmt"
	"os"
)

func LoadFile(fileToLoad string) []string {
	file, err := os.Open(fileToLoad)
	if err != nil {
		fmt.Println("Error opening file:", err)
		fmt.Println("Vérifiez que le fichier " + fileToLoad + " est bien présent dans le répertoire du programme")
		return nil
	}
	defer file.Close()

	r := bufio.NewReader(file)
	words := []string{}
	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			words = append(words, string(line))
		}
		if err != nil {
			return words
		}
	}
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
