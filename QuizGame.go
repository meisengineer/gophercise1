// Welcome to the QuizGame, by Jesse Sauer circa March, 2021 (The future!!!)
// The QuizGame reads questions from a CSV and presents them to the
// contestant. Each question is shown until answered. When the timer expires
// the game is over and the contestant is shown how they've faired.

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Define the command line flags the program will accept. These will all be pointers.
	var filename = flag.String("f", "sample.csv", "file in the data/ directory to load")
	var timeLimitFlag = flag.String("t", "30", "time limit for the quiz, in seconds")
	//var random = flag.Bool("r", false, "randomize the questions")
	flag.Parse()

	// Define input variables
	var (
		record  []string
		qakey   map[string]string
		correct = 0
	)

	timeLimit, _ := time.ParseDuration(*timeLimitFlag)
	timer := time.NewTimer(time.Duration(timeLimit.Seconds()))

	// Read in the file. All quiz sources should be csv and in the data directory
	pwd, err1 := os.Getwd()
	check(err1)
	data, err := ioutil.ReadFile(pwd + "/data/" + *filename)
	check(err)
	reader := csv.NewReader(strings.NewReader(string(data)))
	qakey = make(map[string]string)
	for {
		record, err = reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			check(err)
		}
		qakey[record[0]] = record[1]
	}

	// Prompt the user to start
	fmt.Println("Welcome to the Kwiz Game!")
	fmt.Println("Press Enter to begin...")
	fmt.Scanln()
	timerc := make(chan bool)
	go func() {
		<-timer.C
		time.Sleep(30)
		return timer.Stop()
	}()
	for question, answer := range qakey {
		var temp string
		fmt.Println(question)
		fmt.Scanln(&temp)
		if temp == answer {
			correct++
		}
	}
	fmt.Println("Answers Correct: ", correct, " / ", len(qakey))
}
