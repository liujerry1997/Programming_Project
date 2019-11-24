package main

/*
func TestGenerateAgent(t *testing.T) {
	k := 10000
	row := make([]box, k)
	sensorArmLength := 7
	sensorDiagonalL := 5 * math.Sqrt(2)
	sensorAngle := math.Pi / float64(4)
	row2 := GenerateAgent(row, sensorArmLength, sensorDiagonalL, sensorAngle)
	c := 0
	for i := range row2 {
		if row2[i].IsAgent == true {
			c++
			if row2[i].agent.sensorDiagonalL != sensorDiagonalL {
				t.Error("The sensorDiagonalL hasn't been sent to sensor")
			}
		}
	}
	fmt.Println(c)

}
*/
/*
func TestInitializeBoard(t *testing.T) {
	board1 := InitializeBoard(1, 1)
	board2 := make(multiAgentMatrix, 1)
	board2[0] = make([]box, 1)
	board2[0][0].IsAgent = false
	board2[0][0].IsFood = false
	board2[0][0].foodChemo = 0.0
	board2[0][0].trailChemo = 0.0
	board2[0][0].agent = nil
	if board1[0][0].IsAgent != board2[0][0].IsAgent {
		t.Error("Expected IsAgent is false but received", board1[0][0].IsAgent)
	} else if board1[0][0].IsFood != board2[0][0].IsFood {
		t.Error("Expected IsFood is false but received", board1[0][0].IsFood)
	} else if board1[0][0].foodChemo != board2[0][0].foodChemo {
		t.Error("Expected foodChemo is 0.0 but received", board1[0][0].foodChemo)
	} else if board1[0][0].trailChemo != board2[0][0].trailChemo {
		t.Error("Expected IsAgent is 0.0 but received", board1[0][0].IsAgent)
	} else if board1[0][0].agent != board2[0][0].agent {
		t.Error("Expected IsAgent is nil but received", board1[0][0].agent)
	}
}
*/
/*
func TestAverageNeighbor(t *testing.T) {
	// filter for nutrient 5x5
	filterN := 5
	CN := 10.0
	// filter for trail 3x3
	filterT := 3

	board := InitializeBoard(7, 7)

	//Put food source center in 3,3
	for i := 2; i <= 4; i++ {
		for j := 2; j <= 4; j++ {
			board[i][j].IsFood = true
			board[i][j].foodChemo = CN //10
		}
	}
	for i := 2; i <= 4; i++ {
		for j := 2; j <= 4; j++ {
			board[i][j].trailChemo = CN //10
		}
	}
	//Check if updateFoodChemo is right
	currBoard1 := CopyBoard(board)
	currBoard1Expected := CopyBoard(board)
	currBoard1Expected[1][1].foodChemo = 2.5
	currBoard1Expected[2][1].foodChemo = 3.0

	//Check if updateTrial chemo is right
	currBoard2 := CopyBoard(board)
	currBoard2Expected := CopyBoard(board)
	a1 := float64(10) / float64(9)
	a2 := float64(20) / float64(9)
	a3 := float64(30) / float64(9)
	currBoard2Expected[1][1].trailChemo = a1

	currBoard2Expected[1][2].trailChemo = a2
	currBoard2Expected[1][3].trailChemo = a3

	board1Copy := CopyBoard(currBoard1)
	board2Copy := CopyBoard(currBoard2)
	//For the real currBoard after changin chemo
	currBoard1.AverageNeighbor(board1Copy, 1, 1, filterN, "food")
	currBoard1.AverageNeighbor(board1Copy, 2, 1, filterN, "food")
	if currBoard1Expected[1][1].foodChemo != currBoard1[1][1].foodChemo {
		t.Errorf("Expected foodchemo at board 1 1 is %f but received %f", currBoard1Expected[1][1].foodChemo, currBoard1[1][1].foodChemo)
	}
	if currBoard1[2][1].foodChemo != currBoard1Expected[2][1].foodChemo {
		t.Errorf("Expected foodchemo at board 2 1 is %f but received %f", currBoard1Expected[2][1].foodChemo, currBoard1[2][1].foodChemo)
	}

	currBoard2.AverageNeighbor(board2Copy, 1, 1, filterT, "trial")
	currBoard2.AverageNeighbor(board2Copy, 1, 2, filterT, "trial")
	currBoard2.AverageNeighbor(board2Copy, 1, 3, filterT, "trial")
	if currBoard2[1][1].trailChemo != currBoard2Expected[1][1].trailChemo {
		t.Errorf("Expected trailChemo at board 1 1 is %f but received %f", currBoard2Expected[1][1].trailChemo, currBoard2[1][1].trailChemo)
	}
	if currBoard2[1][2].trailChemo != currBoard2Expected[1][2].trailChemo {
		t.Errorf("Expected trailChemo at board 1 2 is %f but received %f", currBoard2Expected[1][2].trailChemo, currBoard2[1][2].trailChemo)
	}
	if currBoard2[1][3].trailChemo != currBoard2Expected[1][3].trailChemo {
		t.Errorf("Expected trailChemo at board 1 3 is %f but received %f", currBoard2Expected[1][3].trailChemo, currBoard2[1][3].trailChemo)
	}
}
*/

/*
func TestAverageFilter(t *testing.T) {
	// filter for nutrient 5x5
	filterN := 5
	CN := 10.0
	// filter for trail 3x3
	filterT := 3

	board := InitializeBoard(7, 7)

	//Put food source center in 3,3
	for i := 2; i <= 4; i++ {
		for j := 2; j <= 4; j++ {
			board[i][j].IsFood = true
			board[i][j].foodChemo = CN  //10
			board[i][j].trailChemo = CN //10
		}
	}
	//Check if updateFoodChemo is right
	currBoard1 := CopyBoard(board)
	currBoard1Expected := CopyBoard(board)
	currBoard1Expected[1][1].foodChemo = 2.5
	currBoard1Expected[5][1].foodChemo = 2.5
	currBoard1Expected[1][5].foodChemo = 2.5
	currBoard1Expected[5][5].foodChemo = 2.5
	for j := 2; j <= 4; j++ {
		currBoard1Expected[1][j].foodChemo = 3.0
		currBoard1Expected[5][j].foodChemo = 3.0
	}
	for i := 2; i <= 4; i++ {
		currBoard1Expected[i][1].foodChemo = 3.0
		currBoard1Expected[i][5].foodChemo = 3.0
	}
	//Check if updateTrial chemo is right
	currBoard2 := CopyBoard(board)
	currBoard2Expected := CopyBoard(board)
	a1 := float64(10) / float64(9)
	a2 := float64(20) / float64(9)
	a3 := float64(30) / float64(9)
	a4 := float64(40) / float64(9)
	a5 := float64(60) / float64(9)
	a6 := float64(90) / float64(9)
	currBoard2Expected[1][1].trailChemo = a1
	currBoard2Expected[5][1].trailChemo = a1
	currBoard2Expected[1][5].trailChemo = a1
	currBoard2Expected[5][5].trailChemo = a1
	for j := 2; j <= 4; j = j + 2 {
		currBoard2Expected[1][j].trailChemo = a2
		currBoard2Expected[2][j].trailChemo = a4
		currBoard2Expected[5][j].trailChemo = a2
		currBoard2Expected[4][j].trailChemo = a4
	}
	for i := 2; i <= 4; i = i + 2 {
		currBoard2Expected[i][1].trailChemo = a2
		currBoard2Expected[i][5].trailChemo = a2
	}
	currBoard2Expected[1][3].trailChemo = a3
	currBoard2Expected[3][1].trailChemo = a3
	currBoard2Expected[5][3].trailChemo = a3
	currBoard2Expected[3][5].trailChemo = a3

	currBoard2Expected[2][3].trailChemo = a5
	currBoard2Expected[3][2].trailChemo = a5
	currBoard2Expected[4][3].trailChemo = a5
	currBoard2Expected[3][4].trailChemo = a5
	currBoard2Expected[3][3].trailChemo = a6
	//For the real currBoard after changin chemo

	currBoard1.AverageFilter(filterN, "food")

	//Check foodChemo
	for i := range currBoard1 {
		for j := range currBoard1[i] {
			if currBoard1[i][j].foodChemo != currBoard1Expected[i][j].foodChemo {
				t.Errorf("Expected foodchemo at board %d %d is %f but received %f", i, j, currBoard1Expected[i][j].foodChemo, currBoard1[i][j].foodChemo)
			}
		}
	}
	currBoard2.AverageFilter(filterT, "trial")
	for i := range currBoard2 {
		for j := range currBoard2[i] {
			if currBoard2[i][j].trailChemo != currBoard2Expected[i][j].trailChemo {
				t.Errorf("Expected trailChemo at board %d %d is %f but received %f", i, j, currBoard2Expected[i][j].trailChemo, currBoard2[i][j].trailChemo)
			}
		}
	}
}
*/
/*
func TestDamp(t *testing.T) {
	// filter for nutrient 5x5
	dampN := 0.2

	CN := 10.0
	// filter for trail 3x3
	dampT := 0.1
	board := InitializeBoard(7, 7)

	//Put food source center in 3,3
	for i := 2; i <= 4; i++ {
		for j := 2; j <= 4; j++ {
			board[i][j].IsFood = true
			board[i][j].foodChemo = CN  //10
			board[i][j].trailChemo = CN //10
		}
	}

	//Put some trialchemo and foodchemo out of food resources
	for j := 1; j <= 5; j++ {
		board[1][j].foodChemo = 3.0
		board[1][j].trailChemo = 3.0

	}
	//Check if Damp food is right
	currBoard1 := CopyBoard(board)
	currBoard1Expected := CopyBoard(board)
	for j := 1; j <= 5; j++ {
		currBoard1Expected[1][j].foodChemo = 2.400000000000000 //10
	}

	currBoard2 := CopyBoard(board)
	currBoard2Expected := CopyBoard(board)
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 5; j++ {
			currBoard2Expected[i][j].trailChemo *= 0.9 //10
		}
	}

	currBoard1.Damp(dampN, "food")
	currBoard2.Damp(dampT, "trial")

	//check the expected and the calculated
	for i := range currBoard1 {
		for j := range currBoard1[0] {
			if currBoard1[i][j].foodChemo != currBoard1Expected[i][j].foodChemo {
				t.Errorf("Expected foodchemo at board %d %d is %f but received %f", i, j, currBoard1Expected[i][j].foodChemo, currBoard1[i][j].foodChemo)
			}
		}
	}

	for i := range currBoard2 {
		for j := range currBoard2[i] {
			if currBoard2[i][j].trailChemo != currBoard2Expected[i][j].trailChemo {
				t.Errorf("Expected trailChemo at board %d %d is %f but received %f", i, j, currBoard2Expected[i][j].trailChemo, currBoard2[i][j].trailChemo)
			}
		}
	}

}
*/
/*
func TestUpdateChemo(t *testing.T) {
	// filter for nutrient 5x5
	filterN := 5
	dampN := 0.2

	CN := 10.0
	// filter for trail 3x3
	filterT := 3
	dampT := 0.1
	board := InitializeBoard(7, 7)

	//Put food source center in 3,3
	for i := 2; i <= 4; i++ {
		for j := 2; j <= 4; j++ {
			board[i][j].IsFood = true
			board[i][j].foodChemo = CN  //10
			board[i][j].trailChemo = CN //10
		}
	}
	//Check if updateFoodChemo is right
	currBoard1 := CopyBoard(board)
	currBoard1Expected := CopyBoard(board)
	currBoard1Expected[1][1].foodChemo = 2.0
	currBoard1Expected[5][1].foodChemo = 2.0
	currBoard1Expected[1][5].foodChemo = 2.0
	currBoard1Expected[5][5].foodChemo = 2.0
	for j := 2; j <= 4; j++ {
		currBoard1Expected[1][j].foodChemo = 2.4
		currBoard1Expected[5][j].foodChemo = 2.4
	}
	for i := 2; i <= 4; i++ {
		currBoard1Expected[i][1].foodChemo = 2.4
		currBoard1Expected[i][5].foodChemo = 2.4
	}
	//Check if updateTrial chemo is right
	currBoard2 := CopyBoard(board)
	currBoard2Expected := CopyBoard(board)
	a1 := float64(10) * 0.9 / float64(9)
	a2 := float64(20) * 0.9 / float64(9)
	a3 := float64(30) * 0.9 / float64(9)
	a4 := float64(40) * 0.9 / float64(9)
	a5 := float64(60) * 0.9 / float64(9)
	a6 := float64(90) * 0.9 / float64(9)
	currBoard2Expected[1][1].trailChemo = a1
	currBoard2Expected[5][1].trailChemo = a1
	currBoard2Expected[1][5].trailChemo = a1
	currBoard2Expected[5][5].trailChemo = a1
	for j := 2; j <= 4; j = j + 2 {
		currBoard2Expected[1][j].trailChemo = a2
		currBoard2Expected[2][j].trailChemo = a4
		currBoard2Expected[5][j].trailChemo = a2
		currBoard2Expected[4][j].trailChemo = a4
	}
	for i := 2; i <= 4; i = i + 2 {
		currBoard2Expected[i][1].trailChemo = a2
		currBoard2Expected[i][5].trailChemo = a2
	}
	currBoard2Expected[1][3].trailChemo = a3
	currBoard2Expected[3][1].trailChemo = a3
	currBoard2Expected[5][3].trailChemo = a3
	currBoard2Expected[3][5].trailChemo = a3

	currBoard2Expected[2][3].trailChemo = a5
	currBoard2Expected[3][2].trailChemo = a5
	currBoard2Expected[4][3].trailChemo = a5
	currBoard2Expected[3][4].trailChemo = a5
	currBoard2Expected[3][3].trailChemo = a6
	//For the real currBoard after changin chemo

	currBoard1 = UpdateChemo(filterN, dampN, currBoard1, "food")

	//Check foodChemo
	for i := range currBoard1 {
		for j := range currBoard1[i] {
			if currBoard1[i][j].foodChemo != currBoard1Expected[i][j].foodChemo {
				t.Errorf("Expected foodchemo at board %d %d is %f but received %f", i, j, currBoard1Expected[i][j].foodChemo, currBoard1[i][j].foodChemo)
			}
		}
	}
	currBoard2 = UpdateChemo(filterT, dampT, currBoard2, "trial")
	for i := range currBoard2 {
		for j := range currBoard2[i] {
			if currBoard2[i][j].trailChemo != currBoard2Expected[i][j].trailChemo {
				t.Errorf("Expected trailChemo at board %d %d is %f but received %f", i, j, currBoard2Expected[i][j].trailChemo, currBoard2[i][j].trailChemo)
			}
		}
	}
}
*/
/*
func TestSynthesisComparator(t *testing.T) {

}
*/
