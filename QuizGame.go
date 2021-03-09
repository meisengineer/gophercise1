// Welcome to the QuizGame, by Jesse Sauer circa March, 2021 (The future!!!)
// The QuizGame reads questions from a CSV and presents them to the
// contestant. Each question is shown until answered. When the timer expires
// the game is over and the contestant is shown how they've faired.

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	pwd, err1 := os.Getwd()
	check(err1)
	data, err := ioutil.ReadFile(pwd + "/data/sample")
	check(err)

	reader := csv.NewReader(strings.NewReader(string(data)))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%T\n", record)

		fmt.Println(record)
	}

}
