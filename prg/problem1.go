package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Problem struct {
	Input  string
	Answer string
	Seed   int64
}

func CreateProblem() *Problem {
	return &Problem{}
}
func (p *Problem) GetInput() string {
	return p.Input
}
func (p *Problem) GetAnswer() string {
	return p.Answer
}
func (p *Problem) SetSeed(i int64) {
	p.Seed = i
}
func (p *Problem) checkAnswer(answer string) bool {
	if answer == p.Answer {
		return true
	} else {
		return false
	}
}

func (p *Problem) GenerateInputAndAnswer() {
	var buffer bytes.Buffer

	s2 := rand.NewSource(p.Seed)
	r2 := rand.New(s2)

	for i := 0; i < 100; i++ {
		buffer.WriteString(strconv.Itoa(r2.Intn(100)))
		if i < 99 {
			buffer.WriteString(" ")
		}
	}

	p.Input = buffer.String()
	sinput := strings.Split(p.Input, " ")
	ianswer := 0

	for _, v := range sinput {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Wrong conv")
		}
		ianswer += i
	}

	p.Answer = strconv.Itoa(ianswer)
}
func (p *Problem) Print() string {
	return fmt.Sprintf("Seed: %v\nInput: %v\nAnswer: %v \n", p.Seed, p.Input, p.Answer)
}
