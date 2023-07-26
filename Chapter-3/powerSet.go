/*
* @author Elia Renzoni
* @date 26/07/2023
* @brief return the power set of an array (integer)
*/

package main

import "fmt"
import "math"

func main() {
  var (
    inputSet []int = []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    powerSetSubSet int
    powerSet []int = setPowerSet(&powerSetSubSet, inputSet)
  )
  if powerSetSubSet != 0 {
    for _, value := range powerSet {
      fmt.Printf("Value Elem : %d", value)
    }
  } else {
    fmt.Printf("No PowerSet")
  }
}

func setPowerSet(subSetNumbers* int, inputSet []int) []int {
  var powerSet []int
  var j int
  if *subSetNumbers = math.Pow(2, len(inputSet)); *subSetNumbers != 0 {
      subSetCounter := 1
      for i := 0; i < subSetCounter && subSetCounter < *subSetNumbers; i++ {
        if i < len(inputSet) {
          for j < len(inputSet) {
            powerSet[j] = inputSet[i]
            j += 1
          }
          i = 0
          subSetCounter += 1
        } else {
          return powerSet
        }
      }
  } else {
    return powerSet
  }
  return powerSet
}
