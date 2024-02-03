/**
*	@author Elia Renzoni
*	@date 03/02/2024
*	@brief OOP in Go - interfaces
*
 */

package duolingo

type Language interface {
	addWord(str, wordTrad string)
	getRandomWord() string
	wordCorrection(str string) bool
	getVocabolary() []string
}
