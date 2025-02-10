package ahtbpoems

import "math/rand"

type Poem struct {
	Id    int64
	Lines [14]string
}

type PoemID struct {
	ids [14]int
}

func NewPoemID() (pid PoemID) {
	for i := 0; i < len(pid.ids); i++ {
		pid.ids[i] = rand.Intn(10)
	}

	return
}

func PoemIDFromInt(id int64) (pid PoemID) {
	for i := len(pid.ids) - 1; i >= 0; i-- {
		pid.ids[i] = int(id % 10)
		id = id / 10
	}
	return
}

func (pid PoemID) ToIntID() (id int64) {
	for _, i := range pid.ids {
		id = (id * 10) + int64(i)
	}

	return
}

func buildPoem(pid PoemID, texts [10][14]string) (poem Poem) {
	poem.Id = pid.ToIntID()
	for i := 0; i < len(poem.Lines); i++ {
		poem.Lines[i] = texts[pid.ids[i]][i]
	}

	return
}

func PoemFrench(pid PoemID) (poem Poem) {
	return buildPoem(pid, queneau_french)
}

func RandomPoemFrench() (poem Poem) {
	return PoemFrench(NewPoemID())
}

func PoemEnglishRowe(pid PoemID) (poem Poem) {
	return buildPoem(pid, queneau_english_rowe)
}

func RandomPoemEnglishRowe() (poem Poem) {
	return PoemEnglishRowe(NewPoemID())
}

func PoemEnglishChapman(pid PoemID) (poem Poem) {
	return buildPoem(pid, queneau_english_chapman)
}

func RandomPoemEnglishChapman() (poem Poem) {
	return PoemEnglishChapman(NewPoemID())
}

func (poem Poem) ToString() (s string) {
	for _, l := range poem.Lines {
		s += (l + "\n")
	}

	return
}
