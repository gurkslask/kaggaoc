package main

type Problems interface {
	checkAnswer(answer string) bool
	Print() string
	SetSeed(i int64)
	GenerateInputAndAnswer()
	GetInput() string
	GetAnswer() string
}

func GetProblem(p int) Problems {
	// Get problem, if new problem add to map
	mm := map[int]Problems{
		1: CreateProblem(),
	}
	return mm[p]
	//c := CreateProblem()

	//return c
}
