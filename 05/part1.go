package main

import "fmt"
import "bufio"
import "strings"
import "strconv"
import "os"


func main() {
    offsets := []int{}

    // get the data from cli
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter puzzle input (blank line to exit): \n")

    text, _ := reader.ReadString('\n')
    for text != "\n" {
        //fmt.Printf("text: %s\n", text)
        text = strings.TrimSpace(text)
        i, err := strconv.Atoi(text)
        if err == nil {
            // not efficient.  don't care.
            offsets = append(offsets, i)
        } else {
            fmt.Printf("Unable to parse input: %s\n", text)
            os.Exit(1)
        }

        text, _ = reader.ReadString('\n')
    }


    steps := 0
    psn := 0
    for true {
        v := offsets[psn]
        offsets[psn] += 1
        psn = psn + v

        steps += 1

        // did we make it out?
        if psn >= len(offsets) {
            break
        }
    }

    fmt.Printf("steps: %d\n", steps)
}
