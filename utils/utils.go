package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Ints(input string) []int {
	s := Strings(input)
	var numbers []int
	for _, depth := range s {
		i, _ := strconv.Atoi(depth)
		numbers = append(numbers, i)
	}
	return numbers
}

func Binary(input string) []int {
	s := Strings(input)
	var numbers []int
	for _, depth := range s {
		i, _ := strconv.ParseInt(depth, 2, 64)
		numbers = append(numbers, int(i))
	}
	return numbers
}

func Strings(input string) []string {
	return strings.Split(input, "\n")
}

// Input returns the input for the specified year and day as a string,
// downloading it if it does not already exist on disk.
func Input(day int) string {
	os.Mkdir("inputs", 0755)
	filename := fmt.Sprintf("inputs/%02d.txt", day)
	if _, err := os.Stat(filename); err != nil {
		est, err := time.LoadLocation("EST")
		if err != nil {
			panic(err)
		}
		if t := time.Date(2021, time.December, day, 0, 0, 0, 0, est); time.Until(t) > 0 {
			fmt.Printf("Puzzle not unlocked yet! Sleeping for %v...\n", time.Until(t))
			time.Sleep(time.Until(t) + 3*time.Second) // don't want to fire too early
		}
		fmt.Println("Downloading input...")
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2021/day/%v/input", day), nil)
		fmt.Println(os.Getenv("AOC_SESSION"))
		req.AddCookie(&http.Cookie{
			Name:  "session",
			Value: os.Getenv("AOC_SESSION"),
		})
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		if err := ioutil.WriteFile(filename, data, 0660); err != nil {
			panic(err)
		}
	}
	return ReadInput(filename)
}

// ReadInput returns the contents of filename as a string.
func ReadInput(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes.TrimSpace(data))
}
