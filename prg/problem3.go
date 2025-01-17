package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Problem3 struct {
	Input  string
	Answer string
	Seed   int64
}

func CreateProblem3() *Problem3 {
	return &Problem3{}
}
func (p *Problem3) GetInput() string {
	return p.Input
}
func (p *Problem3) GetAnswer() string {
	return p.Answer
}
func (p *Problem3) SetSeed(i int64) {
	p.Seed = i
}
func (p *Problem3) checkAnswer(answer string) bool {
	if answer == p.Answer {
		return true
	} else {
		return false
	}
}

func (p *Problem3) GenerateInputAndAnswer() {
	// Generate data
	var buffer bytes.Buffer

	s2 := rand.NewSource(p.Seed + 2)
	r2 := rand.New(s2)

	for i := 0; i < 130; i++ {
		buffer.WriteString(strconv.Itoa(r2.Intn(100)))
		if i < 129 {
			buffer.WriteString(" ")
		}
	}

	p.Input = buffer.String()
	sinput := strings.Split(p.Input, " ")
	ianswer := 0
	before := 0
	// Generate answer

	for _, v := range sinput {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Wrong conv")
		}
		if i >= before {
			ianswer += i
		}
		before = i
	}

	p.Answer = strconv.Itoa(ianswer)
}
func (p *Problem3) Print() string {
	return fmt.Sprintf("Problem 3: Mer for loopar")
}
