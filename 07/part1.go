package main

import "strconv"
import "bufio"
import "fmt"
import "strings"
import "os"

var Programs = make(map[string]*Program)

type Program struct {
    name string
    weight int
    parent *Program
    children []*Program
}

func (p *Program) SetName(name string) {
    p.name = name
}

func (p *Program) SetWeight(weight int) {
    p.weight = weight
}

func (p *Program) SetParent(parent *Program) {
    p.parent = parent
}

func (p *Program) AddChild(child *Program) {
    p.children = append(p.children, child)
}

func GetProgram(name string) *Program {
    if Programs[name] == nil || Programs[name].name != name {
        p := &Program{}
        p.SetName(name)
        p.children = []*Program{}
        Programs[name] = p
    }

    return Programs[name]
}

func main() {
    // get the data from cli
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter puzzle input: \n")

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

        parent := GetProgram(name)
        parent.SetWeight(weight)

        // children
        if len(split) > 2 {
            x := 3 // offset for split, gets us past the "->"
            for x < len(split) {
                // create child, references parent
                child_name := strings.Replace(split[x], ",", "", -1)
                child := GetProgram(child_name)
                child.SetParent(parent)
                parent.AddChild(child)

                x += 1
            }
        }

        text, _ = reader.ReadString('\n')
    }

    // who is the "bottom"?  meaning who has no parent?
    for name := range Programs {
        program := Programs[name]
        if program.parent == nil || program.parent.name == "" {
            fmt.Printf("bottom: %s\n", program.name)
            os.Exit(0)
        }
    }
}
