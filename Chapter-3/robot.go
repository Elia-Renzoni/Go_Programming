package main

import "fmt"

const (
  command1 = 'R'
  command2 = 'L'
)

func main() {
  var (
      robotMovement = []int16 {-5, -3, -1, 0, 4, 6, 7, 8}
      movementPath string = "RRLR"
      sum int16
      secondLatency int16
  )
  checkRobotMovement(robotMovement, movementPath, &sum, secondLatency)
  if sum != 0 {
    fmt.Printf("Somma delle distanze : %d", sum)
  } else {
    fmt.Printf("Somma delle distanze inesistente !")
  }
}

func checkRobotMovement(movementArr []int16, movementRules string, sum* int16, afterSecond int16) {
  var (
    firstRobotIndex = 0
    secondRobotIndex = len(movementArr) - 1
    val1, val2 int16
  )
  for indexValue := range movementRules {
    if movementRules[indexValue] == command1 {
      firstRobotIndex++
      val1 = movementArr[firstRobotIndex]
    } else {
      firstRobotIndex--
      val1 = movementArr[firstRobotIndex]
    }
    if movementRules[indexValue + 1] == command1	{
      secondRobotIndex++
      val2 = movementArr[secondRobotIndex]
    } else {
      secondRobotIndex--
      val2 = movementArr[secondRobotIndex]
    }
    if val2 != val1 {
      *sum += 1
    }
  }
}
