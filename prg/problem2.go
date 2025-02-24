package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Problem2 struct {
	Input  string
	Answer string
	Seed   int64
}

func CreateProblem2() *Problem2 {
	return &Problem2{}
}
func (p *Problem2) GetInput() string {
	return p.Input
}
func (p *Problem2) GetAnswer() string {
	return p.Answer
}
func (p *Problem2) SetSeed(i int64) {
	p.Seed = i
}
func (p *Problem2) checkAnswer(answer string) bool {
	if answer == p.Answer {
		return true
	} else {
		return false
	}
}

func (p *Problem2) GenerateInputAndAnswer() {
	var buffer bytes.Buffer

	s2 := rand.NewSource(p.Seed + 1)
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

	for _, v := range sinput {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Wrong conv")
		}
		if i%2 == 0 {
			ianswer += i
		}
	}

	p.Answer = strconv.Itoa(ianswer)
}
func (p *Problem2) Print() string {
	return fmt.Sprintf("Problem 2: IF sats")
}
