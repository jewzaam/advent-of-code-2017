package main

import "bufio"
import "fmt"
import "strings"
import "os"
import "strconv"
import "math"

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
        max := 0
        min := int(math.MaxInt32)

        // cleanup
        text = strings.Replace(text, "\t", " ", -1)
        text = strings.Replace(text, "\n", "", -1)

        // parse and find max and min values
        numbers := strings.Split(text, " ")
        x := 0
        for x < len(numbers) {
            s := numbers[x]
            i, err := strconv.Atoi(s)
            if err == nil {
//                fmt.Printf("x: %d, s: %s, i: %d\n", x, s, i)
                max = Max(max, i)
                min = Min(min, i)
            }
            x += 1
        }

        // get diff and add to checksum
        checksum += (max - min)

//        fmt.Printf("min: %d, max: %d, checksum: %d", min, max, checksum)

        text, _ = reader.ReadString('\n')
    }

    fmt.Printf("checksum: %d\n", checksum)
}
