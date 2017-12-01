package main

import "fmt"
import "os"
import "bufio"
import "strings"

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

    // Start off with the 'last' value from the end of the string
    // because we wrap for comparison.  This means you do nothing
    // sipecial to handle that condition.
    psn := 0
    last := text[len(text) - 1]
    output := 0
    for psn < len(text) {
        next := text[psn]
//        fmt.Printf("psn: %d, last: %d, next: %d, output: %d\n", psn, last, next, output)
        if last == next {
            output += int(next - byte('0'))
        }
        last = next
        psn += 1
    }
    fmt.Printf("captcha: %d\n", output)
}
