package game

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/githbert/wordlet/fileop"
	"github.com/githbert/wordlet/misc"
	"github.com/githbert/wordlet/rnd"
)

type Result struct {
	nrWords int
	correct int
}

func getUniqueNumber(m int, s []int) (int, error) {

	for i := 0; i < m; i++ {
		r := rnd.GetNumber(0, m)
		if !misc.SliceContains(s, r) {
			return r, nil
		}
	}

	return 0, errors.New("unable to find a unique number")
}

func getRandomWords(w []string, n int) []string {
	var uniqueNr []int
	var wordList []string

	// create a list of uniquenumbers
	for i := 0; i < n; i++ {
		un, err := getUniqueNumber(len(w)-1, uniqueNr)
		if err != nil {
			panic(err)
		}
		uniqueNr = append(uniqueNr, un)
	}

	// popluate a wordlist of those numbers
	for _, i := range uniqueNr {
		wordList = append(wordList, w[i])
	}

	return wordList
}

func getAnswer() string {
	var answer string
	fmt.Scanln(&answer)

	if len(answer) == 0 {
		return ""
	}

	return answer
}

func showWords(wl []string, w int) {

	fmt.Printf("Kom ihåg dessa ord ")
	fmt.Println(strings.Join(wl[:], ", "))
	time.Sleep(time.Duration(w) * time.Second)
}

func typeWords(wl []string) []string {
	var typed []string
	fmt.Println("Skriv alla orden som visades (dom behöver inte vara i ordning):")
	for range wl {
		fmt.Printf("> ")
		typed = append(typed, getAnswer())
	}

	return typed
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func calcResult(wl []string, al []string, s *Result) {
	for _, word := range wl {
	inner:
		for _, answer := range al {
			if word == answer {
				s.correct++
				break inner
			}
		}
	}
}

func printResults(wl []string, al []string, s *Result) {
	fmt.Printf("Ord  : ")
	fmt.Println(wl)
	fmt.Printf("Svar : ")
	fmt.Println(al)
	fmt.Println("--------+---------")
	fmt.Printf("Words   : %d\n", s.nrWords)
	fmt.Printf("Correkt : %d\n", s.correct)
	fmt.Println("--------+---------")
}

func GameLoop(path string, nr int, w int) {
	stat := Result{correct: 0, nrWords: nr}
	words, err := fileop.ReadFile(path)
	if err != nil {
		log.Fatalf("ReadFile: %s", err)
	}

	selectedWords := getRandomWords(words, nr)
	showWords(selectedWords, w)
	clearScreen()
	typedWords := typeWords(selectedWords)
	calcResult(selectedWords, typedWords, &stat)
	printResults(selectedWords, typedWords, &stat)

}
