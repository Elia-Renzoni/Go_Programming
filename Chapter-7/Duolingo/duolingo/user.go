/**
*	@author Elia Renzoni
*	@date 03/02/2024
*	@brief OOP in Go - interfaces
 */

package duolingo

type User interface {
	setUserName()
	setPassword()
	setLanguageToLearn()
}
