package main

import "bufio"
import "fmt"
import "strings"
import "os"
import "strconv"

func Min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func Max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    // get the data from cli by reading until get empty row
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter spreadsheet followed by empty row:\n")

    text := ""
    checksum := 0

    text, _ = reader.ReadString('\n')
    for text != "\n" {
        // cleanup
        text = strings.Replace(text, "\t", " ", -1)
        text = strings.Replace(text, "\n", "", -1)

        // parse and find max and min values
        numbers := strings.Split(text, " ")
        x := 0
        for x < len(numbers) {
            a, err := strconv.Atoi(numbers[x])
            x += 1
            if err != nil {
                continue
            }
            y := 0
            for y < len(numbers) {
                b, err := strconv.Atoi(numbers[y])
                y += 1
                if err != nil {
                    continue
                }
                if a == Min(a, b) {
                    // picking arbitrary pair to skip.. so we don't double count
                    continue
                }
                min := Min(a, b)
                max := Max(a, b)
                m := max % min
                if m == 0 {
                    checksum += (max / min)
                }
            }
        }

        text, _ = reader.ReadString('\n')
    }

    fmt.Printf("checksum: %d\n", checksum)
}
