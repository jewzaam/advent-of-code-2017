package main

import "strconv"
import "bufio"
import "fmt"
import "strings"
import "os"

var Registers = make(map[string]int)

type Condition struct {
    register string
    op string
    value int
}

func (c Condition) Test() bool {
    output := false

    reg_val := Registers[c.register]

    switch c.op {
    case ">":
        if reg_val > c.value {
            output = true
        }
    case "<":
        if reg_val < c.value {
            output = true
        }
    case ">=":
        if reg_val >= c.value {
            output = true
        }
    case "<=":
        if reg_val <= c.value {
            output = true
        }
    case "!=":
        if reg_val != c.value {
            output = true
        }
    case "==":
        if reg_val == c.value {
            output = true
        }
    default:
        fmt.Println("Unable to parse Condition operator: %s\n", c.op)
        os.Exit(1)
    }

    return output
}

type Instruction struct {
    register string
    op string
    value int
    condition *Condition
}

func (i *Instruction) SetValues(vals []string) {
    i.register = vals[0]
    i.op = vals[1]
    i.value, _ = strconv.Atoi(vals[2])

    // psn 3 is always "if"
    if vals[3] != "if" {
        fmt.Println("Expected 'if': %s\n", vals[3])
        os.Exit(1)
    }

    c_val, _ := strconv.Atoi(vals[6])

    i.condition = &Condition{vals[4], vals[5],c_val}
}

func (i Instruction) Execute() {
    if i.condition.Test() {
        switch i.op {
        case "inc":
            Registers[i.register] += i.value
        case "dec":
            Registers[i.register] -= i.value
        default:
            fmt.Println("Unable to parse Instruction operator: %s\n", i.op)
            os.Exit(1)
        }
    }
}

func main() {
    // get the data from cli
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter puzzle input: \n")

    instructions := make([]*Instruction, 0)

    text, _ := reader.ReadString('\n')
    for text != "\n" {
        //fmt.Printf("text: %s\n", text)
        text = strings.Replace(text, "\n", "", -1)
        text = strings.TrimSpace(text)

        split := strings.Split(text, " ")

        i := &Instruction{}
        i.SetValues(split)

        instructions = append(instructions, i)

        text, _ = reader.ReadString('\n')
    }

    // walk through instructions and execute
    for _, instruction := range instructions {
        instruction.Execute()
    }

    // find largest register value
    big := 0
    for register := range Registers {
        value := Registers[register]
        if value > big {
            big = value
        }
    }

    fmt.Printf("biggest: %d\n", big)
}
