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


    x := 0
    psn := 0
    for x < len(offsets) {
        v := offsets[psn]
        offsets[psn] += 1
        psn = psn + v

        x += 1
    }

    fmt.Printf("steps: %d\n", x)
}
