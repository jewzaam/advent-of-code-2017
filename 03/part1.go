package main

import "bufio"
import "fmt"
import "strings"
import "os"
import "strconv"
import "math/big"

func factorial(n int64) *big.Int {
   if n < 0 {
      return big.NewInt(1)
   }
   if (n==0) {
      return big.NewInt(1)
   }
   bigN := big.NewInt(n)
   return bigN.Mul(bigN, factorial(n-1))
}

func GetLayerWidth(layer int) int {
    return (layer - 1) * 2 + 1
}

func GetLayerCount(layer int) int {
    c := (GetLayerWidth(layer) - 1) * 4
    if c < 1 {
        // layer 1 (center) will result in a count of 0, but needs to be 1
        c = 1
    }
    return c
}

func GetInput() int {
    text := ""

    // try to get text from command line arg
    if len(os.Args) > 1 {
        text = os.Args[1]
    } else {
        // get the data from cli
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter puzzle input: ")
        text, _ = reader.ReadString('\n')
    }

    text = strings.TrimSpace(text)

    i, err := strconv.Atoi(text)
    if err == nil {
        return i
    } else {
        fmt.Printf("Unable to parse input as a number: %d\n", text)
        os.Exit(1)
    }

    return i
}

func Abs(x int) int {
    if x < 0 {
        return x * -1
    }
    return x
}

func main() {
    input := GetInput()

    layer := 1
    count := 0
    // artificial limit, just so it doesn't continue forever
    for layer < 500 {
        width := GetLayerWidth(layer)
        c := GetLayerCount(layer)

        // find min by adding prior counts
        min := count + 1
        max := min + c - 1

        // add current layer to count for next iteration
        count += c

//        fmt.Printf("layer: %d, size: %dx%d, count: %d, min: %d, max: %d\n", layer, width, width, count, min, max)

        if min <= input && input <= max {
            fmt.Printf("Input found in layer %d!\n", layer)

            // it doesn't matter what side
            // what does matter is where relative to the middle of any side
            x := Abs((input - min) % (width - 1) - ((width -1) / 2 - 1))

            steps := layer - 1 + x

            fmt.Printf("Steps required: %d\n", steps)

            break
        }

        layer += 1
    }
}
