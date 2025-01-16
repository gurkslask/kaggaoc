package main

import (
	"errors"
	"fmt"
)

type Problems interface {
	checkAnswer(answer string) bool
	Print() string
	SetSeed(i int64)
	GenerateInputAndAnswer()
	GetInput() string
	GetAnswer() string
}

func GetProblem(p int) (Problems, error) {
	// Get problem, if new problem add to map
	mm := map[int]Problems{
		1: CreateProblem(),
		2: CreateProblem2(),
	}
	if P, ok := mm[p]; ok {
		return P, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Challenge %v does not exist", p))
	}
}
