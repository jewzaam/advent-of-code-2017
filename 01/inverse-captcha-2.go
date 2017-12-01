package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "math"

func main() {
    text := ""

    // try to get text from command line arg
    if len(os.Args) > 1 {
        text = os.Args[1]
    } else {
        // get the data from cli
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter digits: ")
        text, _ = reader.ReadString('\n')
    }

    text = strings.TrimSpace(text)

    // verify we have an even number of bytes
    if math.Mod(float64(len(text)), 2) != 0 {
        fmt.Println("ERROR: Input must be an even number of digits")
        os.Exit(1)
    }

    psn := 0
    offset := len(text) / 2
    output := 0
    for psn < offset {
        this := text[psn]
        next := text[psn + offset]
//        fmt.Printf("psn: %d, this: %d, next: %d, output: %d\n", psn, this, next, output)
        if this == next {
            output += int(next - byte('0'))
        }
        psn += 1
    }
    output *= 2
    fmt.Printf("captcha: %d\n", output)
}
