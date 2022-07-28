package main

import (
    "fmt"
    "os"
    "strings"
)   

func get_valid_guesses() []string {
    data, err := os.ReadFile("all_guesses.txt")
    check(err)

    guess_arr := strings.Split(string(data), "\n")
    return guess_arr
}

func get_valid_answers() []string {
    data, err := os.ReadFile("answer_list.txt")
    check(err)

    guess_arr := strings.Split(string(data), "\n")
    return guess_arr
}

func verify_guess(guess string) bool {
    if len(guess) != 5 {
        return false
    }

    return true
}

func verify_color(color string) bool {
    if len(color) != 5 {
        return false
    }

    for _, char := range color {
        if char != 'b' && char != 'g' && char != 'y' {
            return false
        }
    }

    return true
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    fmt.Println("Welcome to WORDLE Solve Go!")
    fmt.Println("Wordle Solve Go is a standalone rewrite from the original Python version in an effort to increase speed.")

    //valid_guesses := get_valid_guesses()
    //valid_answers := get_valid_answers()

    var guesses []string
    var colors []string

    fmt.Println("Enter your guesses:")
    for {
        var guess string
        fmt.Scanln(&guess)

        if guess == "" {
            break
        }
        if !verify_guess(guess) {
            fmt.Println("Error, that's not a valid guess.")
        }

        guesses = append(guesses, guess)
    }

    fmt.Println("Enter your colors:")
    for i := 0; i < len(guesses); i++ {
        var color string
        fmt.Scanln(&color)

        if !verify_color(color) {
            fmt.Println("Error, that's not a valid color string.")
        }

        colors = append(colors, color)
    }

    fmt.Println(guesses)
    fmt.Println(colors)
}
