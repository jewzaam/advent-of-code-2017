package main

import "strconv"
import "bufio"
import "fmt"
import "strings"
import "os"

func main() {
    // get the data from cli
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter puzzle input: \n")

    weights := make(map[string]int)
    children := make(map[string][]string)
    parent := make(map[string]string)

    text, _ := reader.ReadString('\n')
    for text != "\n" {
        //fmt.Printf("text: %s\n", text)
        text = strings.Replace(text, "\n", "", -1)
        text = strings.TrimSpace(text)

        split := strings.Split(text, " ")
        // contents of splits
        // <name> (<weight>) -> (child1), (child2), ...

        name := split[0]
        weight, err := strconv.Atoi(strings.Replace(strings.Replace(split[1], "(", "", -1), ")", "", -1))
        if err != nil {
            fmt.Printf("Unable to parse weight: %s\n", split[1])
            os.Exit(1)
        }

        weights[name] = weight

        // children
        if len(split) > 2 {
            x := 3 // offset for split, gets us past the "->"
            c := make([]string, len(split) - 3)
            for x < len(split) {
                // create child, references parent
                c[x-3] = strings.Replace(split[x], ",", "", -1)
                parent[c[x-3]] = name

                x += 1
            }

            children[name] = c
        }

        text, _ = reader.ReadString('\n')
    }

    // who is the "bottom"?  meaning who has no parent?
    for k := range children {
        //fmt.Printf("k: %s, parent: %s\n", k, parent[k])
        if parent[k] == "" {
            fmt.Printf("bottom: %s\n", k)
            os.Exit(0)
        }
    }
}
