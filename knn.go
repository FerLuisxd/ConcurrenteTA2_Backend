// KNN project main.go
// Using data from https://archive.ics.uci.edu/ml/machine-learning-databases/iris/iris.data
package main

import (
	"fmt"
	"log"
	"math"
	// "math/rand"
	"sort"
	// "strconv"
)
var testSet []Adult
var trainSet []Adult
var k int 
func main3() {


	// var recordSet []Adult = adults

	k = 36
	fmt.Println("total ")
	fmt.Println(len(adults))
	for i := range adults {
		// if rand.Float64() < 0.9 {
		// if i < 10 {
		// 	testSet = append(testSet, adults[i])
		// } else {
			trainSet = append(trainSet, adults[i])
		// }
	}

	var predictions []string
	fmt.Println("test lenght")
	fmt.Println(len(testSet))

	for x := 0; x < len(testSet); x++ {
        
    }
	for x := range testSet {
		// neighbors := getNeighbors(trainSet, testSet[x], k)
		// result := getResponse(neighbors)
		result := testCase(trainSet,testSet[x],k)
		predictions = append(predictions, result[0].key)
		fmt.Printf("Predicted: %s, Actual: %s\n", result[0].key, testSet[x].Ke)
	}

	accuracy := getAccuracy(testSet, predictions)
	fmt.Printf("Accuracy: %f%s\n", accuracy, "%")
}

func testCase(trainSetA []Adult,testSetObject Adult, k int) sortedClassVotes {
	fmt.Println(testSetObject)
	neighbors := getNeighbors(trainSetA, testSetObject, k)
	result := getResponse(neighbors)
	return result
}


func getAccuracy(testSet []Adult, predictions []string) float64 {
	correct := 0

	for x := range testSet {
		if testSet[x].Ke == predictions[x] {
			correct += 1
		}
	}

	return (float64(correct) / float64(len(testSet))) * 100.00
}

type classVote struct {
	key   string
	value int
}

type sortedClassVotes []classVote

func (scv sortedClassVotes) Len() int           { return len(scv) }
func (scv sortedClassVotes) Less(i, j int) bool { return scv[i].value < scv[j].value }
func (scv sortedClassVotes) Swap(i, j int)      { scv[i], scv[j] = scv[j], scv[i] }

func getResponse(neighbors []Adult) sortedClassVotes {
	classVotes := make(map[string]int)

	for x := range neighbors {
		response := neighbors[x].Ke
		if contains(classVotes, response) {
			classVotes[response] += 1
		} else {
			classVotes[response] = 1
		}
	}

	scv := make(sortedClassVotes, len(classVotes))
	i := 0
	for k, v := range classVotes {
		scv[i] = classVote{k, v}
		i++
	}

	sort.Sort(sort.Reverse(scv))
	return scv
}

type distancePair struct {
	record   Adult
	distance float64
}

type distancePairs []distancePair

func (slice distancePairs) Len() int           { return len(slice) }
func (slice distancePairs) Less(i, j int) bool { return slice[i].distance < slice[j].distance }
func (slice distancePairs) Swap(i, j int)      { slice[i], slice[j] = slice[j], slice[i] }

func getNeighbors(trainingSet []Adult, testRecord Adult, k int) []Adult {
	var distances distancePairs
	for i := range trainingSet {
		// Se suma y calcula la distnacia concurrentemente
	// 	go func(instanceOne Adult, instanceTwo Adult){
	// 		var distance float64
	// distance += math.Pow(float64((instanceOne.Fnlwgt - instanceTwo.Fnlwgt)), 2)
	// distance += math.Pow(float64((instanceOne.EducationNum - instanceTwo.EducationNum)), 2)
	// distance += math.Pow(float64((instanceOne.CapitalGain - instanceTwo.CapitalGain)), 2)
	// distance += math.Pow(float64((instanceOne.CapitalLoss - instanceTwo.CapitalLoss)), 2)
	// distance += math.Pow(float64((instanceOne.Hours - instanceTwo.Hours)), 2)
	// distance += math.Pow(float64((instanceOne.Sex - instanceTwo.Hours)), 2)

	// distances = append(distances, distancePair{instanceTwo, math.Sqrt(distance)})

	// 	}(testRecord,trainingSet[i])

		dist := euclidianDistance(testRecord, trainingSet[i])
		distances = append(distances, distancePair{trainingSet[i], dist})
	}

	sort.Sort(distances)

	var neighbors []Adult

	for x := 0; x < k; x++ {
		neighbors = append(neighbors, distances[x].record)
	}

	return neighbors
}
func concurrent(){
		// 	go func(instanceOne Adult, instanceTwo Adult){
	// 		var distance float64

	// distance += math.Pow(float64((instanceOne.Fnlwgt - instanceTwo.Fnlwgt)), 2)
	// distance += math.Pow(float64((instanceOne.EducationNum - instanceTwo.EducationNum)), 2)
	// distance += math.Pow(float64((instanceOne.CapitalGain - instanceTwo.CapitalGain)), 2)
	// distance += math.Pow(float64((instanceOne.CapitalLoss - instanceTwo.CapitalLoss)), 2)
	// distance += math.Pow(float64((instanceOne.Hours - instanceTwo.Hours)), 2)
	// distance += math.Pow(float64((instanceOne.Sex - instanceTwo.Hours)), 2)
	
	// dist := math.Sqrt(distance)

	// distances = append(distances, distancePair{instanceTwo, dist})

	// 	}(testRecord,trainingSet[i])
}

func euclidianDistance(instanceOne Adult, instanceTwo Adult) float64 {
	var distance float64

	distance += math.Pow(float64((instanceOne.Fnlwgt - instanceTwo.Fnlwgt)), 2)
	distance += math.Pow(float64((instanceOne.EducationNum - instanceTwo.EducationNum)), 2)
	distance += math.Pow(float64((instanceOne.CapitalGain - instanceTwo.CapitalGain)), 2)
	distance += math.Pow(float64((instanceOne.CapitalLoss - instanceTwo.CapitalLoss)), 2)
	distance += math.Pow(float64((instanceOne.Hours - instanceTwo.Hours)), 2)
	distance += math.Pow(float64((instanceOne.Sex - instanceTwo.Hours)), 2)

	return math.Sqrt(distance)
}


func errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func contains(votesMap map[string]int, name string) bool {
	for s, _ := range votesMap {
		if s == name {
			return true
		}
	}

	return false
}