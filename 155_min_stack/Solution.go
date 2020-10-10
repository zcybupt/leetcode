package main

import "fmt"

type MinStack struct {
    stackA, stackB []int
    lastA, lastB   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        stackA: make([]int, 0),
        stackB: make([]int, 0),
        lastA: -1,
        lastB: -1,
    }
}

func (this *MinStack) Push(x int) {
    this.stackA = append(this.stackA, x)
    if this.lastB == -1 || this.stackB[this.lastB] >= x {
        this.stackB = append(this.stackB, x)
        this.lastB++
    }
    this.lastA++
}

func (this *MinStack) Pop() {
    if this.stackA[this.lastA] == this.stackB[this.lastB] {
        this.stackB = this.stackB[:this.lastB]
        this.lastB--
    }
    this.stackA = this.stackA[:this.lastA]
    this.lastA--
}

func (this *MinStack) Top() int {
    return this.stackA[this.lastA]
}

func (this *MinStack) GetMin() int {
    return this.stackB[this.lastB]
}

func main() {
    stack := Constructor()
    stack.Push(-2)
    stack.Push(0)
    stack.Push(-3)
    fmt.Println(stack.GetMin())
    stack.Pop()
    fmt.Println(stack.Top())
    fmt.Println(stack.GetMin())
}

