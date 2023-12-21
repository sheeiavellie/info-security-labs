package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MatrixParams struct {
	Ops             Operations
	NumberOfUsers   int
	NumberOfObjects int
}

func NewMatrixParams(
	opr *Operations,
	numOfUsers,
	numOfObjects int,
) *MatrixParams {
	return &MatrixParams{
		Ops:             *opr,
		NumberOfUsers:   numOfUsers,
		NumberOfObjects: numOfObjects,
	}
}

type Operations struct {
	Lenght int
	Ops    map[string]int
}

func NewOperations() *Operations {
	ops := map[string]int{
		"Assignment": 0,
		"Write":      0,
		"Read":       0,
	}

	return &Operations{
		Lenght: 3,
		Ops:    ops,
	}
}

type UniqueRand struct {
	generated       map[int]bool
	generationRange int
}

func NewUniqueRand(genRange int) *UniqueRand {
	return &UniqueRand{
		generated:       make(map[int]bool),
		generationRange: genRange,
	}
}

func (u *UniqueRand) Intn() int {
	for {
		i := rand.Intn(u.generationRange)
		if !u.generated[i] {
			u.generated[i] = true
			return i
		}
	}
}

func main() {
	operations := NewOperations()

	matrixParams := NewMatrixParams(operations, 4, 3)

	matrix := generateMatrix(*matrixParams)

	users := []string{"Ivan", "Sergey", "Boris", "Alexander"}

	var currentUser int
	var currentObj int
	var menu string

	quit := false
	fmt.Print("Enter command >")
	fmt.Scan(&menu)

	for !quit {

		switch menu {
		case "init":
			fmt.Printf("Welcome! Enter your id: ")
			fmt.Scan(&currentUser)

			if currentUser > len(users)-1 {
				fmt.Printf("No user with id %d\n", currentUser)
				break
			} else {
				fmt.Printf("User: %s\n", users[currentUser])
			}
			fmt.Println()
			fmt.Print("Enter command >")
			fmt.Scan(&menu)

		case "read":
			fmt.Println()
			fmt.Print("Enter object id: ")
			fmt.Scan(&currentObj)
			fmt.Println(execRead(currentUser, currentObj, matrix))

			fmt.Println()
			fmt.Print("Enter command >")
			fmt.Scan(&menu)

		case "write":
			fmt.Println()
			fmt.Print("Enter object id: ")
			fmt.Scan(&currentObj)
			fmt.Println(execWrite(currentUser, currentObj, matrix))

			fmt.Println()
			fmt.Print("Enter command >")
			fmt.Scan(&menu)

		case "assign":
			fmt.Println()
			fmt.Print("Enter object id: ")
			fmt.Scan(&currentObj)

			var userAssignId int
			fmt.Println()
			fmt.Print("Enter user id: ")
			fmt.Scan(&userAssignId)

			var rightId int
			fmt.Println()
			fmt.Print("Enter right id: ")
			fmt.Scan(&rightId)
			fmt.Println(execAssign(currentUser, currentObj, userAssignId, rightId, matrix))

			fmt.Println()
			fmt.Print("Enter command >")
			fmt.Scan(&menu)
		case "quit":
			fmt.Println("Goodbye!")
			quit = true
		default:
			fmt.Printf("Command %s does not exist\n", menu)
			fmt.Print("Enter command >")
			fmt.Scan(&menu)
		}
	}
}

func execRead(userId, objId int, m [][][3]int) string {
	if m[userId][objId][2] == 1 {
		return "Command Executed!"
	} else {
		return "No access!"
	}
}
func execWrite(userId, objId int, m [][][3]int) string {
	if m[userId][objId][1] == 1 {
		return "Command Executed!"
	} else {
		return "No access!"
	}
}
func execAssign(userId, objId, userAssignId, rightId int, m [][][3]int) string {
	if m[userId][objId][0] == 1 {
		m[userAssignId][objId][rightId] = 1
		return "Command Executed!"
	} else {
		return "No access!"
	}
}

func generateMatrix(mp MatrixParams) [][][3]int {
	matrix := make([][][3]int, mp.NumberOfUsers)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range matrix {
		matrix[i] = make([][3]int, mp.NumberOfObjects)
		for j := range matrix[i] {
			randOp := r.Intn(mp.Ops.Lenght)
			for k := 0; k < randOp; k++ {
				randOpIndex := r.Intn(mp.Ops.Lenght)
				elem := [3]int{}
				elem[k] = randOpIndex

				matrix[i][j] = elem
			}
		}
	}

	for k := range matrix[0] {
		elem := [3]int{1, 1, 1}
		matrix[0][k] = elem
	}

	return matrix
}
