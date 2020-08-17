package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Board struct {
	MaxDepth int
	Stacks   []*Stack
	Moves    []Move
}

type Move struct {
	From  int
	To    int
	Color Color
}

func (m *Move) Print() {
	fmt.Printf("%2d -> %2d %s\n", m.From+1, m.To+1, ColorMapFull[m.Color])
}

func NewBoard(data string) *Board {
	b := &Board{}
	lines := strings.Split(data, "\n")
	lastLine := lines[len(lines)-1]
	if len(lastLine) == 0 {
		lines = lines[:len(lines)-1]
	}

	b.MaxDepth = len(lines)
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		columns := strings.Split(line, " ")
		if len(b.Stacks) == 0 {
			for idx := 0; idx < len(columns); idx++ {
				b.Stacks = append(b.Stacks, new(Stack))
			}
			// always two empty stacks
			b.Stacks = append(b.Stacks, new(Stack))
			b.Stacks = append(b.Stacks, new(Stack))
		}
		if len(columns)+2 != len(b.Stacks) {
			panic("rows of unequal lengths")
		}
		for idx, column := range columns {
			stack := b.Stacks[idx]
			color := ReverseColorMap[column]
			stack.Push(color)
		}
	}

	return b
}

func (b *Board) Solved() bool {
	for _, stack := range b.Stacks {
		depth := stack.Depth()
		if depth != 0 && depth != b.MaxDepth {
			return false
		}
	}

	for _, stack := range b.Stacks {
		if stack.Depth() != 0 {
			color := stack.Items[0]
			for _, item := range stack.Items {
				if item != color {
					return false
				}
			}
		}
	}

	return true
}

func (b *Board) Print() {
	for i := b.MaxDepth - 1; i >= 0; i-- {
		for _, stack := range b.Stacks {
			if stack.Depth() > i {
				color := stack.Items[i]
				column := ColorMap[color]
				fmt.Printf("%s", column)
			} else {
				fmt.Printf("..")
			}
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
	}

	for _, move := range b.Moves {
		move.Print()
	}
}

func (b *Board) Solve() {
	var idle int

	for !b.Solved() && !b.NoMoves() && len(b.Moves) < 5000 {
		idle++
		if idle == 110 {
			return
		}
		to := b.RandomToColumn()
		from := b.RandomFromColumn(to)
		if from == -1 {
			continue
		}
		if b.checkReversesPrevious(from, to) {
			continue
		}
		if b.monochromeReversed(from, to) {
			from, to = to, from
		}
		stack1 := b.Stacks[from]
		stack2 := b.Stacks[to]

		idle = 0
		b.Moves = append(b.Moves, Move{From: from, To: to, Color: stack1.Top()})
		stack2.Push(stack1.Pop())
	}
}

func monochrome(stack *Stack) bool {
	color := stack.Top()
	for _, c := range stack.Items {
		if c != color {
			return false
		}
	}
	return true
}

func (b *Board) monochromeReversed(from, to int) bool {
	stack1 := b.Stacks[from]
	stack2 := b.Stacks[to]
	if stack1.Depth() == 3 && stack2.Depth() == 1 {
		if monochrome(stack1) {
			return true
		}
	}
	return false
}

func (b *Board) RandomFromColumn(to int) int {
	var candidates []int

	toStack := b.Stacks[to]
	if toStack.Depth() == 0 {
		// any non-monochromatic stack
		for idx, stack := range b.Stacks {
			if stack.Depth() < 2 {
				continue
			}
			color := stack.Top()
			for _, c := range stack.Items {
				if c != color {
					candidates = append(candidates, idx)
					break
				}
			}
		}
	} else {
		// any color-matching stack (that isn't the destination)
		color := toStack.Top()
		for idx, stack := range b.Stacks {
			if idx == to {
				continue
			}
			if stack.Depth() == 0 {
				continue
			}
			if stack.Top() == color {
				candidates = append(candidates, idx)
			}
		}
	}
	if len(candidates) == 0 {
		return -1
	}
	return candidates[rand.Int()%len(candidates)]
}

func (b *Board) RandomToColumn() int {
	var candidates []int
	for idx, stack := range b.Stacks {
		if stack.Depth() != b.MaxDepth {
			candidates = append(candidates, idx)
		}
	}

	return candidates[rand.Int()%len(candidates)]
}

func (b *Board) NoMoves() bool {
	for i := 0; i < len(b.Stacks); i++ {
		stack1 := b.Stacks[i]
		if stack1.Depth() == 0 {
			return false
		}
		if stack1.Depth() == b.MaxDepth {
			continue
		}
		color := stack1.Top()
		for j := 0; j < len(b.Stacks); j++ {
			if j == i {
				continue
			}
			stack2 := b.Stacks[j]
			if stack2.Depth() == 0 {
				return false
			}
			if stack2.Top() == color {
				return false
			}
		}
	}
	return true
}

func (b *Board) checkReversesPrevious(from, to int) bool {
	for i := len(b.Moves) - 1; i >= 0; i-- {
		m := b.Moves[i]
		if m.From == to && m.To == from {
			return true
		}
		if m.From == from || m.From == to || m.To == from || m.To == to {
			return false
		}
	}
	return false
}
