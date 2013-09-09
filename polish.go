package main

import (
    "fmt"
    "strings"
    "container/list"
)

type OTree struct {
    value string
    children *List
}

type Operator struct {
    symbol rune
    arity int
}


func main() {
    // input is a sequence of runes
    var vars map[rune]bool
    var ops map[rune]Operator

    t, err := parse(inp, ops, vars)

}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func makeTree(val string, children []*OTree) *Tree {
    // children is a slice of OTrees. we need to build
    // a linked list in the order
    // actually, we should probably use a real stack datastructure
    // instead of a slice. then have a general function for
    // walking the stack, populating the linked list backwards?
    return &OTree{val, nil}
}

func parse(s string, ops, vars) (*OTree, error) {
    stack := make([]*OTree, 0)
    curr := 0

    for _, sym := range reverse(s) {
        curr = len(tree_stack)

        if op, ok := ops[sym]; ok {
            // this is an operator, parse the whole subtree
            // this should actually be like:
            // pop op.arity values off the stack. saving them
            // backwards in a linked list, and pass that to makeTree
            stack = append(stack[:curr-2],
                makeTree(string(v), stack[curr-1], stack[curr-2]))
        } else if vars[sym] {
            stack = append(stack, makeTree(string(v), nil))
        } else {
            return nil, error.New("Unrecognized symbol")
        }
    }

    return tree_stack[0], nil
}

/*
func (t *OTree) DFS() map[string]bool {
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
*/
