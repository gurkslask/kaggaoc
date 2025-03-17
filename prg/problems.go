package main

import (
	"bytes"
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
type ProblemStruct struct {
	pProblems map[int]Problems
}

func CreateProblemStruct() ProblemStruct {
	return ProblemStruct{pProblems: map[int]Problems{
		1: CreateProblem(),
		2: CreateProblem2(),
		3: CreateProblem3()},
	}

}

func (p ProblemStruct) GetProblem(pi int) (Problems, error) {
	if P, ok := p.pProblems[pi]; ok {
		return P, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Challenge %v does not exist", p))
	}
}
func (p ProblemStruct) GetProblems() string {
	var buffer bytes.Buffer
	for k, v := range p.pProblems {
		buffer.WriteString(fmt.Sprintf("%v: %v\n", k, v.Print()))
	}
	return buffer.String()

}
