package main

import "strconv"
import "bufio"
import "fmt"
import "strings"
import "os"

func GetMemoryHash(memory []int) string {
    sa := make([]string, len(memory))

    x := 0
    for x < len(memory) {
        s := strconv.Itoa(memory[x])
        sa[x] = s
        x += 1
    }

    return strings.Join(sa, " ")
}

func main() {
    // get the data from cli
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter puzzle input: \n")

    text, _ := reader.ReadString('\n')

    text = strings.Replace(text, "\t", " ", -1)
    text = strings.TrimSpace(text)

    memory_s := strings.Split(text, " ")
    memory := make([]int, len(memory_s))

    x := 0
    for x < len(memory_s) {
        i, err := strconv.Atoi(memory_s[x])

        if err != nil {
            fmt.Printf("Unable to convert to integer: %s\n", memory_s[x])
            os.Exit(1)
        }

        memory[x] = i
        x += 1
    }

    // use a map for the lookup of keys, values are the step it was seen
    history := make(map[string]int)
    // offset all values by +1 because int default is needed and it's 0
    history[GetMemoryHash(memory)] = 1

    running := true
    count := 1

    //fmt.Println(memory)
    steps := 0

    for running {
        //fmt.Printf("count: %d, memory: %s\n", count, memory)
        // find biggest bank
        b_psn := -1
        b_size := -1

        x := 0
        for x < len(memory) {
            if (memory[x] > b_size || (memory[x] == b_size && x < b_psn)) {
                b_size = memory[x]
                b_psn = x
            }
            x += 1
        }

        // reallocate
        // 1. zero out max block
        // 2. allocate block to memory (distribute)
        // 3. check if we have seen it before
        memory[b_psn] = 0

        for b_size > 0 {
            //fmt.Printf("b_size: %d, b_psn: %d\n", b_size, b_psn)
            // shift position
            b_psn += 1

            // check if we hit end of memory
            if b_psn >= len(memory) {
                // wrap
                b_psn = 0
            }

            // distribute
            memory[b_psn] += 1
            b_size -=1
        }

        // check if we have a duplicate condition
        h := GetMemoryHash(memory)

        //fmt.Println(memory)

        f := history[h]

        if f > 0 {
            // seen it before.  exit loops
            //fmt.Printf("FOUND: %d\n", f)
            running = false
            // increment by 1 because it's inclusive
            steps = count - f + 1
            break
        }

        // haven't seen it, add to history and continue
        count += 1
        history[h] = count // value is offset by 1 on purpose
    }

    fmt.Printf("count: %d\n", count)
    fmt.Printf("steps: %d\n", steps)
}
