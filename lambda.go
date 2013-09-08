package main

import "fmt"

type Tree struct {
    left *Tree
    value string
    right *Tree
}

func (t *Tree) DFS() map[string]bool {
    vars := make(map[string]bool)

    vars[t.value] = true

    if t.left != nil {
        for k, v := range t.left.DFS() {
            vars[k] = v
        }
    }

    if t.right != nil {
        for k, v := range t.right.DFS() {
            vars[k] = v
        }
    }

    return vars
}


func main() {
    t := &Tree{
        &Tree{nil, "e", &Tree{nil, "y", nil}},
        "h",
        nil,
    }

    fmt.Println(t.DFS())
}
