package main

import "strconv"
import "bufio"
import "fmt"
import "strings"
import "os"


type Program struct {
    name string
    weight int
    child_names []string
    parent Program
}

func main() {
    // get the data from cli
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter puzzle input: \n")

    programs := make(map[string]Program)

    text, _ := reader.ReadString('\n')
    for text != "\n" {
        fmt.Printf("text: %s\n", text)
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

        // children
        children := []string{}
        if len(split) > 2 {
            x := 3 // offset for split, gets us past the "->"
            children = make([]string, len(split) - 3)
            for x < len(split) {
                children[x-3] = strings.Replace(split[x], ",", "", -1)
                x += 1
            }
        }

        programs[name] = Program{ name, weight, children }

        fmt.Println(programs[name])

        text, _ = reader.ReadString('\n')
    }
}
