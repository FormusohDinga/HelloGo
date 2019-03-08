package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	// cmd := exec.Command("docker", "save", "--help")
	cmd := exec.Command("youtube-dl", "--extract-audio",
		"--audio-format", "mp3", "https://hw-mp3.datpiff.com/mixtapes/5/m21621c4/04%20-%20Charli%20Brown%20%20Walk%20on%20by.mp3")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "hh")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}
