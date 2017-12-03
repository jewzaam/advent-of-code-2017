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
// create large array to hold values.
const max_layers = 50
const z = max_layers*2+1
var data [z][z]int
const psn_center_x = max_layers + 1
const psn_center_y = max_layers + 1

func UpdateDataAt(target int, x int, y int) {
    if data[x][y] > 0 {
        // do not update values
        //fmt.Printf("skip [%d][%d]\n", x - psn_center_x, y - psn_center_y)
        return
    }

    value := data[x + 1][y + 0]
    value += data[x + 0][y + 1]
    value += data[x - 1][y + 0]
    value += data[x + 0][y - 1]
    value += data[x + 1][y + 1]
    value += data[x + 1][y - 1]
    value += data[x - 1][y - 1]
    value += data[x - 1][y + 1]

    data[x][y] = value
    //fmt.Printf("data[%d][%d] = %d\n", x - psn_center_x, y - psn_center_y, value)

    if target < value {
        fmt.Printf("First value larger than input: %d\n", value)
        os.Exit(0)
    }
}

func main() {
    input := GetInput()


    layer := 2
    psn_x := psn_center_x
    psn_y := psn_center_y

    // starting position gets value of 1
    data[psn_x][psn_y] = 1

    // artificial limit, just so it doesn't continue forever
    for layer <= max_layers {
        width := GetLayerWidth(layer)
        //fmt.Printf("layer = %d, width = %d\n", layer, width)

        // first position in layer is [x+1][y]
        psn_x += 1

        // update data at starting position
        UpdateDataAt(input, psn_x, psn_y)

        // move up
        y := 2 // slight offset because we already have done the starting position
        //fmt.Println("Move Up")
        for y < width {
            psn_y -= 1
            UpdateDataAt(input, psn_x, psn_y)
            y += 1
        }

        // move left
        //fmt.Println("Move Left")
        x := 1
        for x < width {
            psn_x -= 1
            UpdateDataAt(input, psn_x, psn_y)
            x += 1
        }

        // move down
        //fmt.Println("Move Down`")
        y = 1
        for y < width {
            psn_y += 1
            UpdateDataAt(input, psn_x, psn_y)
            y += 1
        }

        // move right
        //fmt.Println("Move Right`")
        x = 1
        for x < width {
            psn_x += 1
            UpdateDataAt(input, psn_x, psn_y)
            x += 1
        }

        layer += 1
    }

    fmt.Println("Did't find a value larger.. weird")
}
