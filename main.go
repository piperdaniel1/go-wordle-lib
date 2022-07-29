package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)   

type Solver struct {
    // TODO
}

// Struct to validate guesses
type Game struct {
    guess_list []string
    answer_list []string
    strict_mode bool
}

type WordlePlayer struct {
    guesses []string
    colors []string
}

func contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}

func arr_compare(arr1 []string, arr2 []string) bool {
    if len(arr1) != len(arr2) {
        return false
    }

    for i, elem1 := range arr1 {
        if elem1 != arr2[i] {
            return false
        }
    }
    
    return true
}

func num_chars(char string, word []string) int {
    num := 0
    for _, e := range word {
        if e == char {
            num++
        }
    }

    return num
}

// count the number of green and yellow letters that
// are already filled in in the color_list
// these letters must correspond to the right char
func num_detracting_chars(char string, guess []string, color_list []string) int {
    detracting_chars := 0

    for i, guess_letter := range guess {
        if guess_letter == char && (color_list[i] == "g" || color_list[i] == "y")  {
            detracting_chars++
        }
    }

    return detracting_chars
}

func gen_colors(guess *string, correct_word *string) string {
    colors := []string{"", "", "", "", ""}
    correct_word_arr := strings.Split(*correct_word, "")
    guess_arr := strings.Split(*guess, "")
    
    // initial round to get greens and blacks
    for i, char := range *guess {
        if string(char) == correct_word_arr[i] {
            colors[i] = "g"
        }
        if !contains(correct_word_arr, string(char)) {
            colors[i] = "b"
        }
    }

    // second round to get yellows
    for i, char := range *guess {
        if string(char) == correct_word_arr[i] || !contains(correct_word_arr, string(char)) {
            continue
        }

        total_chars := num_chars(string(char), correct_word_arr)
        detracting_chars := num_detracting_chars(string(char), guess_arr, colors)

        if total_chars - detracting_chars > 0 {
            colors[i] = "y"
        } else {
            colors[i] = "b"
        }
    }

    return strings.Join(colors, "")
}

func get_words_remaining(wp *WordlePlayer, ga *Game) []string {
    valid_words := []string{}

    for _, word := range ga.answer_list {
        valid_flag := true
        for i, guess := range wp.guesses {
            color_str := gen_colors(&guess, &word)

            if color_str != wp.colors[i] {
                valid_flag = false
                break
            }
        }

        if valid_flag {
            valid_words = append(valid_words, word)
        }
    }
    
    return valid_words
}

// Adds a guess to the wordle player
// Returns status based on whether the guess was valid
// true = added successfully
// false = invalid guess
func (wp *WordlePlayer) add_guess(ga *Game, guess string) bool {
    if !ga.validate(guess) {
        return false
    } else {
        wp.guesses = append(wp.guesses, guess)
        return true
    }
}

// Adds a color to the wordle player
// Returns status based on whether the color was valid
// true = added successfully
// false = invalid guess
func (wp *WordlePlayer) add_color(ga *Game, color string) bool {
    if !verify_color(color) {
        return false
    } else {
        wp.colors = append(wp.colors, color)
        return true
    }
}


// Verify that a given string is a valid guess
func (ga Game) validate (guess string) bool {
    if len(guess) != 5 {
        return false
    }

    if ga.strict_mode {
        match_flag := false
        for _, word := range ga.guess_list {
            if word == guess {
                match_flag = true
            }
        }
        if !match_flag {
            return false
        }
    }

    return true
}

// Import an array of valid guesses from
// the all_guesses.txt file
func get_valid_guesses() []string {
    data, err := os.ReadFile("all_guesses.txt")
    check(err)

    guess_arr := strings.Split(string(data), "\n")

    sanitized_guesses := []string{}

    for _, guess := range guess_arr {
        if len(guess) == 5 {
            sanitized_guesses = append(sanitized_guesses, guess)
        }
    }
    
    return sanitized_guesses
}

// Import an array of valid colors from
// the all_guesses.txt file
func get_valid_answers() []string {
    data, err := os.ReadFile("answer_list.txt")
    check(err)

    guess_arr := strings.Split(string(data), "\n")

    sanitized_guesses := []string{}

    for _, guess := range guess_arr {
        if len(guess) == 5 {
            sanitized_guesses = append(sanitized_guesses, guess)
        }
    }
    
    return sanitized_guesses
}

// Verify that a given string is a valid color string
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

// Collect guesses from the user, adding them to the passed
// curr_guesses array
func collect_guesses(wp *WordlePlayer, ga *Game) {
    fmt.Println("Enter your guesses:")
    for {
        var guess string
        fmt.Print(" > ")
        fmt.Scanln(&guess)

        if guess == "" {
            break
        }

        if !wp.add_guess(ga, guess) {
            fmt.Println("Error, that's not a valid guess.")
            continue
        }
    }

    fmt.Println("You have entered the following guesses:")
    fmt.Println(wp.guesses)
}

// Collect colors from the user, adding them to the passed
// curr_colors array. Can only collect colors until the colors
// array is equal in length to the passed num_guesses int
func collect_colors(wp *WordlePlayer, ga *Game) {
    fmt.Println("Enter your colors:")
    for len(wp.colors) < len(wp.guesses) {
        var color string
        fmt.Print(" > ")
        fmt.Scanln(&color)

        if !wp.add_color(ga, color) {
            fmt.Println("Error, that's not a valid color string.")
            continue
        }
    }
}

// Panic on error
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// entry point
func main() {
    test()
    fmt.Println("Welcome to WORDLE Solve Go!")
    fmt.Println("Wordle Solve Go is a standalone rewrite from the original Python version in an effort to increase speed.")

    ga := Game{get_valid_guesses(), get_valid_answers(), true}
    wp := WordlePlayer{[]string{}, []string{}}

    // collect guesses and colors from the user
    collect_guesses(&wp, &ga)
    collect_colors(&wp, &ga)

    stime := time.Now()
    words_rem := get_words_remaining(&wp, &ga)
    etime := time.Now()
    fmt.Println("Time elapsed:", etime.Sub(stime))
    fmt.Println(wp)
    fmt.Println(len(words_rem), "remaining words: ")
    fmt.Println(words_rem)
}
