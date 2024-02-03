package duolingo

import (
	"math/rand"
	"time"
)

type English struct {
	word     string
	wordTrad string
	nextWord *English
	headList *English
}

func (e *English) addWord(str, wordTrad string) {
	var newWord = &English{}
	newWord.word = str
	newWord.wordTrad = wordTrad

	if e.headList == nil {
		e.headList = newWord
	} else {

	}

}

func (e *English) getLastNode() *English {
	var nodes *English
	for node := e.headList; node != nil; node = node.nextWord {
		nodes = node
	}
	return nodes
}

func (e *English) getRandomWord() string {
	rand.Seed(time.Now().Unix())
	var randomWord string
	i := 0
	for node := e.headList; node != nil && i <= rand.Intn(100); node = node.nextWord {
		i++
		randomWord = node.word
	}

	return randomWord
}

func (e *English) wordCorrection(str string) bool {
	var node *English
	for node = e.headList; node != nil && node.wordTrad == str; node = node.headList {
	}

	switch {
	case node != nil:
		fallthrough
	case node.wordTrad == str:
		return true
	case node == nil, node.wordTrad != str:
		return false
	}

	return false
}
