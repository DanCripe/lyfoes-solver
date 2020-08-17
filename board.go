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
	From int
	To   int
}

func (m *Move) Print() {
	fmt.Printf("%d -> %d\n", m.From, m.To)
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
		// for count := 0; !b.Solved() && !b.NoMoves() && count < len(b.Stacks)*len(b.Stacks)*10; count++
		idle++
		if idle == 100 {
			return
		}
		to := b.RandomToColumn()
		from := b.RandomFromColumn(to)
		if from == -1 {
			continue
		}
		stack1 := b.Stacks[from]
		stack2 := b.Stacks[to]
		/*
			if stack1.Depth() == 0 {
				fmt.Printf("from empty column\n")
				// can't move from empty column
				continue
			}
			if stack2.Depth() == b.MaxDepth {
				fmt.Printf("to full column\n")
				// can't move to full column
				continue
			}
			if stack2.Depth() != 0 && stack2.Top() != stack1.Top() {
				fmt.Printf("different colors\n")
				// can't move to partially full column with a different color
				continue
			}
			if stack1.Depth() == 1 && stack2.Depth() == 0 {
				// don't both moving from column with 1 item to column with 0 items
				fmt.Printf("1 to empty\n")
				continue
			}
			if stack2.Depth() == 0 {
				monochrome := true
				color := stack1.Top()
				for _, c := range stack2.Items {
					if c != color {
						monochrome = false
						break
					}
				}
				if monochrome {
					fmt.Printf("monochrome\n")
					continue
				}
			}
		*/
		if len(b.Moves) != 0 {
			lastMove := b.Moves[len(b.Moves)-1]
			if lastMove.From == to && lastMove.To == from {
				// don't undo a move that was just done
				continue
			}
		}
		idle = 0
		b.Moves = append(b.Moves, Move{From: from, To: to})
		stack2.Push(stack1.Pop())
	}
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
