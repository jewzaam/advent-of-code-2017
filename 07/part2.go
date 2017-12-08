package main

import "strconv"
import "bufio"
import "fmt"
import "strings"
import "os"
import "sort"

var Programs = make(map[string]*Program)

type Program struct {
    name string
    weight int
    stack_weight int
    child_weight int
    parent *Program
    children []*Program
}

func (p Program) GetTargetWeight() int {
    // target is to balance my stack weight with my peers.
    // target is the median weight of peer's stacks.
    // then, my weight should be peer median peer stack weight - my children's weight

    if p.parent == nil {
        return p.weight
    }

    weights := make([]int, len(p.parent.children))
    for x, child := range p.parent.children {
        weights[x] = child.stack_weight
    }

    sort.Ints(weights)
    median := weights[1] // shortcut, this cannot work unless everybody w/ child has 3+ children

    return median - p.child_weight
}

func (p Program) IsBalanced() bool {
    if p.GetTargetWeight() == p.weight {
        return true
    }
    return false
}

func (p *Program) SetName(name string) {
    p.name = name
}

func (p *Program) SetWeight(weight int) {
    p.weight = weight
}

func (p *Program) SetStackWeight(sw int) {
    p.stack_weight = sw
}

func (p *Program) SetChildWeight(cw int) {
    p.child_weight = cw
}

func (p *Program) SetParent(parent *Program) {
    p.parent = parent
}

func (p *Program) AddChild(child *Program) {
    p.children = append(p.children, child)
}

func (p *Program) Print() {
    p.PrintIndented(0, true)
}

func (p *Program) PrintIndented(indent int, recurse bool) {
    prefix := ""
    x := indent
    for x > 0 {
        prefix += "\t"
        x -= 1
    }

    fmt.Printf("%s%s (%d=%d+%d|%d)\n", prefix, p.name, p.stack_weight, p.weight, p.child_weight, p.GetTargetWeight())

    if recurse {
        for _, child := range p.children {
            child.PrintIndented(indent + 1, recurse)
        }
    }
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

func CalculateStackWeight(program *Program) {
    sw := 0

    for _, child := range program.children {
        CalculateStackWeight(child)
        sw += child.stack_weight
    }

    program.SetChildWeight(sw)
    program.SetStackWeight(sw + program.weight)
}

func GetUnbalanced(program *Program) *Program {
    for _, child := range program.children {
        unbalanced := GetUnbalanced(child)
        if unbalanced != nil {
            return unbalanced
        }
    }

    if !program.IsBalanced() {
        return program
    }

    return nil
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
    bottom := &Program{}
    for name := range Programs {
        program := Programs[name]
        if program.parent == nil || program.parent.name == "" {
            bottom = program
        }
    }

    fmt.Printf("bottom: %s\n", bottom.name)
    // we have the bottom.. get all the stack weights now
    CalculateStackWeight(bottom)

//    bottom.Print()

    // which one is unbalanced?
    unbalanced := GetUnbalanced(bottom)

    fmt.Printf("unbalanced: %s\n", unbalanced.name)

    fmt.Printf("target: %d\n", unbalanced.GetTargetWeight())

}
