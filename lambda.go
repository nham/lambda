package main

import (
    "fmt"
    "strings"
)

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

func (t *Tree) String() string {
    s := ""
    if t.value == `\` || t.value == "A" {
        s = "(" + t.value + " "
        if t.left != nil { s += t.left.String() }
        s += " "
        if t.right != nil { s += t.right.String() }
        s += ")"

    } else {
        s = t.value
    }

    return s
}


/* For parsing */
func killWhitey(s string) string {
    ws := "\t\n\r "

    for _, v := range ws {
        s = strings.Replace(s, string(v), "", -1)
    }

    return s
}

func makeTree(val string, l, r *Tree) *Tree {
    return &Tree{l, val, r}
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func parse(s string) (*Tree, error) {
    tree_stack := make([]*Tree, 0)
    clen := 0

    for _, v := range reverse(killWhitey(s)) {
        clen = len(tree_stack)
        if v == '\\' || v == 'A' {
            tree_stack = append(tree_stack[:clen-2],
                makeTree(string(v), tree_stack[clen-1], tree_stack[clen-2]))
        } else {
            tree_stack = append(tree_stack, makeTree(string(v), nil, nil))
        }

    }

    return tree_stack[0], nil
}


/* Main */

func main() {
    t := &Tree{
        &Tree{nil, "e", &Tree{nil, "y", nil}},
        "h",
        nil,
    }
    fmt.Println(t.DFS())

    inp := `A \a \ b A b c x`
    s, _ := parse(inp)
    fmt.Println(s)
    fmt.Println(s.DFS())
}
