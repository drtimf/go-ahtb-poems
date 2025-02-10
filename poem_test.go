package ahtbpoems

import (
	"encoding/json"
	"testing"

	"gotest.tools/assert"
)

func TestPoemID(t *testing.T) {
	var id int64 = 12345678901234
	pid := PoemIDFromInt(id)
	assert.Equal(t, pid.ToIntID(), id)

	pid2 := NewPoemID()
	assert.Equal(t, pid2, PoemIDFromInt(pid2.ToIntID()))
	assert.Equal(t, pid2.ToIntID(), PoemIDFromInt(pid2.ToIntID()).ToIntID())
}

func TestPoemBuild(t *testing.T) {
	poem0 := PoemEnglishChapman(PoemIDFromInt(0))
	for i := 0; i < len(poem0.Lines); i++ {
		assert.Equal(t, poem0.Lines[i], queneau_english_chapman[0][i])
	}

	poemX := PoemFrench(PoemIDFromInt(1234567890123))
	assert.Equal(t, poemX.Id, int64(1234567890123))
	for i := 0; i < len(poemX.Lines); i++ {
		assert.Equal(t, poemX.Lines[i], queneau_french[i%10][i])
	}
}

func TestToString(t *testing.T) {
	poem := PoemEnglishChapman(PoemIDFromInt(1234567890123))
	assert.Equal(t, poem.ToString(), `Don Pedro from his shirt has washed the fleas
Since Elgin left his nostrils in the stone
Upon his old oak chest he cuts his cheese
Which neither time nor tide can long postpone
Old Galileo's Pisan offerings
Filching the lolly country thrift helped save
The fertile mmother changeling drops like kings
That every verbal shock aims to deprave
     The wolf devours both sheep and shepherdess
     Or grinning like a pale-faced golliwog
     And played their mountain croquet jungle chess
     Their sculptors did our best our hulks they clog
     On fish-slab whale nor seal has never swum
     And lessors' dates have all too short a sum
`)
}

func TestPoemJSON(t *testing.T) {
	poem := PoemEnglishChapman(PoemIDFromInt(1234567890123))
	poemJSON, err := json.Marshal(poem)
	assert.NilError(t, err)
	assert.Equal(t, string(poemJSON), `{"Id":1234567890123,"Lines":["Don Pedro from his shirt has washed the fleas","Since Elgin left his nostrils in the stone","Upon his old oak chest he cuts his cheese","Which neither time nor tide can long postpone","Old Galileo's Pisan offerings","Filching the lolly country thrift helped save","The fertile mmother changeling drops like kings","That every verbal shock aims to deprave","     The wolf devours both sheep and shepherdess","     Or grinning like a pale-faced golliwog","     And played their mountain croquet jungle chess","     Their sculptors did our best our hulks they clog","     On fish-slab whale nor seal has never swum","     And lessors' dates have all too short a sum"]}`)
}
