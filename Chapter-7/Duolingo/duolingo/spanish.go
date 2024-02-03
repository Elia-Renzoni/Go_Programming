package duolingo

import (
	"math/rand"
	"time"
)

type Spanish struct {
	word     string
	wordTrad string
	nextWord *Spanish
	headList *Spanish
}

func (s *Spanish) addWord(str, wordTrad string) {
	var newWord = &Spanish{}
	newWord.word = str
	newWord.wordTrad = wordTrad

	if s.headList == nil {
		s.headList = newWord
	} else {
		lastNode := s.getLastNode()
		lastNode.nextWord = newWord
	}
}

func (s *Spanish) getLastNode() *Spanish {
	var nodes *Spanish
	for node := s.headList; node != nil; node = node.nextWord {
		nodes = node
	}
	return nodes
}

func (s *Spanish) getRandomWord() string {
	rand.Seed(time.Now().Unix())
	var randomWord string
	i := 0
	for node := s.headList; node != nil && i <= rand.Intn(100); node = node.nextWord {
		i++
		randomWord = node.word
	}

	return randomWord
}

func (s *Spanish) wordCorrection(str string) bool {
	var node *Spanish
	for node = s.headList; node != nil && node.wordTrad == str; node = node.headList {
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
