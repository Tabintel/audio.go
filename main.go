package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	//"golang.org/x/text/message"
)

func playtheAudio() {
	f, err := os.Open("play.mp3")
	if err != nil {
		fmt.Println("Error opening audio file:", err)
		return
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Println("Error decoding audio file:", err)
		return
	}
	defer streamer.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		fmt.Println("Error initializing audio playback:", err)
		return
	}

	done := make(chan bool)

	speaker.Play(beep.Seq(streamer, beep.Callback(func ()  {
		done <- true
	})))

	<-done
}

func readandDisplay(name string, weekEarn, monthEarn int) {
	message := fmt.Sprintf("%s, you are a great technical writer and Go developer.\n", name)

	message += fmt.Sprintf("You earn $%d weekly and $%d monthly. You write and build.", weekEarn, monthEarn)

	cmd := exec.Command("gtts-cli", message, "-o", "play.mp3")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error generating audio:", err)
		return
	}

	fmt.Println("=== Ekemini's Earnings Summary ===")
	fmt.Println(message)
	fmt.Println("==================================")

	playtheAudio()
}

func main() {
	name := "Ekemini"
	weekEarn := 2000
	monthEarn := 10000

	for {
		readandDisplay(name, weekEarn, monthEarn)
		time.Sleep(24 * time.Hour)
	}
}