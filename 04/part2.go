package main

import "bufio"
import "fmt"
import "strings"
import "os"
import "sort"

func main() {
    valid_count := 0

    // get the data from cli
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter puzzle input (blank line to exit): \n")

    text, _ := reader.ReadString('\n')
    for text != "\n" {
        //fmt.Printf("text: %s\n", text)
        text = strings.TrimSpace(text)
        words := strings.Split(text, " ")

        // sort chars in each word (policy: cannot have anagrams)
        x := 0
        for x < len(words) {
            w := words[x]

            w_split := strings.Split(w, "")
            sort.Strings(w_split)
            w_join := strings.Join(w_split, "")

            words[x] = w_join

            x += 1
        }

        //fmt.Printf("words with chars sorted: %s", words)

        // check if each word is unique
        a := 0
        valid := true
        for a < len(words) && valid {
            //fmt.Printf("a: %d\n", a)
            w := words[a]


            b := 0
            count := 0
            for b < len(words) && valid {
                //fmt.Printf("b: %d\n", b)
                // count how many times the word shows up
                if words[b] == w {
                    count += 1
                }

                // if count > 1, jump ship
                if count > 1 {
                    valid = false
                    break
                }

                b += 1
            }

            a += 1
        }

        if valid {
            valid_count += 1
        }

        text, _ = reader.ReadString('\n')
    }

    fmt.Printf("Valid count: %d\n", valid_count)
}
