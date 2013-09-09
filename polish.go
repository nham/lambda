package main

import (
    "fmt"
    "errors"
    "container/list"
)

type OTree struct {
    value string
    children *list.List
}

type Operator struct {
    symbol rune
    arity int
}


func main() {
    // input is a sequence of runes
    vars := make(map[rune]bool)
    ops := make(map[rune]*Operator)
    ops['$'] = &Operator{'$', 3}
    vars['a'] = true
    vars['b'] = true
    vars['c'] = true

    t, err := parse("$abc", ops, vars)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(t)

}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func makeTree(val string, children []*OTree) *OTree {
    // children is a slice of OTrees. we need to build
    // a linked list in the order
    // actually, we should probably use a real stack datastructure
    // instead of a slice. then have a general function for
    // walking the stack, populating the linked list backwards?
    l := list.New()
    if len(children) > 0 {
        // recall the slice is in reverse order. yeah this shouldve been taken
        // care of prior calling this function, but I dont want to do it properly
        // now. just do it in reverse order.
        for _, v := range children {
            fmt.Println("pushing", v)
            l.PushFront(v)
        }
    }

    return &OTree{val, l}
}

func parse(s string, ops map[rune]*Operator, vars map[rune]bool) (*OTree, error) {
    stack := make([]*OTree, 0)
    curr := 0

    for _, sym := range reverse(s) {
        fmt.Println("sym = ", string(sym))
        curr = len(stack)

        if op, ok := ops[sym]; ok {
            // this is an operator, parse the whole subtree
            // this should actually be like:
            // pop op.arity values off the stack. saving them
            // backwards in a linked list, and pass that to makeTree
            fmt.Println("arity: ", op.arity)
            stack = append(stack[:curr - op.arity],
                    makeTree(string(sym), stack[curr - op.arity : curr]))
            fmt.Println(" stack is now: ", stack)
        } else if vars[sym] {
            stack = append(stack, makeTree(string(sym), nil))
            fmt.Println(" stack is now: ", stack)
        } else {
            return nil, errors.New("Unrecognized symbol")
        }
    }

    return stack[0], nil
}

func (t *OTree) String() string {
    s := ""
    if t.children != nil && t.children.Front() != nil {
        s = "(" + t.value
        for n := t.children.Front(); n != nil; n = n.Next() {
            s += " "
            s += n.Value.(*OTree).String()
        }
        s += ")"

    } else {
        s = t.value
    }

    return s
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

*/
