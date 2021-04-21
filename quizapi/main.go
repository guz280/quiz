package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	// "github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func handleRequests() {
	// myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/", homePage)
	// myRouter.HandleFunc("/questions", returnQuestionIds)
	// myRouter.HandleFunc("/results", results).Methods("POST")
	// myRouter.HandleFunc("/question/{id}", returnQuestionAndAnswers)
	// log.Fatal(http.ListenAndServe(":1000", myRouter))
	http.HandleFunc("/", homePage)
	// return list of question ids
	http.HandleFunc("/questions", returnQuestionIds)
	http.HandleFunc("/question", returnQuestionAndAnswers)
	http.HandleFunc("/results", results)
	log.Fatal(http.ListenAndServe(":1000", nil))
}

// START - GET Question Ids
type QuestionId struct {
	Id int
}

// initialise
var QuestionIds []QuestionId

// questions ids
func returnQuestionIds(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDpoint Hit: returnQuestionIds")
	QuestionIds = []QuestionId{
		QuestionId{Id: 1},
		QuestionId{Id: 2},
		QuestionId{Id: 3},
		QuestionId{Id: 4},
		QuestionId{Id: 5},
		QuestionId{Id: 6},
		QuestionId{Id: 7},
	}
	json.NewEncoder(w).Encode(QuestionIds)
}

// END - GET Question Ids

// START - GET Question & Answer By Id
type QandAs struct {
	QandAs []QandA `json:"QandA"`
}
type QandA struct {
	//Id       int    `json:"id"`
	Question string `json:"question"`
	Answers  string `json:"answers"`
}

func returnQuestionAndAnswers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDpoint Hit: returnQuestionAndAnswers")

	var id = r.URL.Query().Get("id")
	questionId, err := strconv.Atoi(id)
	// handle error
	if err != nil {
		fmt.Println(err)
	}

	if (questionId > 7) || (questionId < 1) {
		fmt.Println("Please select an existing question id from 1 to 7")
	} else {
		// Open our jsonFile
		jsonFile, err := os.Open("QuestionsAndAnswers.json")
		// handle error
		if err != nil {
			fmt.Println(err)
		}
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
		// read our opened jsonFile as a byte array.
		byteValue, _ := ioutil.ReadAll(jsonFile)
		// initialise QandA array
		var qandas QandAs
		json.Unmarshal(byteValue, &qandas)
		fmt.Println(qandas.QandAs[questionId-1])
	}
}

// END - GET Question & Answer By Id

// START - POST Question Number & Answer Id

var GoodAnswerScores []int

type Answers struct {
	Answers []Answer `json:"Answers"`
}
type Answer struct {
	Id     int `json:"id"`
	Answer int `json:"answer"`
}

// I know that this needs to be done a proper POST, but had issues with library "github.com/gorilla/mux"
func results(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Has to be a POST Method")
	} else {
		//fmt.Println(r.Method)
		// get the body of our POST request
		reqBody, _ := ioutil.ReadAll(r.Body)
		//fmt.Println(string(reqBody))

		var ints []int
		err := json.Unmarshal([]byte(reqBody), &ints)
		if err != nil {
			log.Fatal(err)
		}

		var goodAnswers, badAnswers int = 0, 0
		// loop through array to submit good/bad answers

		// load answers from json
		// Open our jsonFile
		jsonFile, err := os.Open("Answers.json")
		// handle error
		if err != nil {
			fmt.Println(err)
		}
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
		// read our opened jsonFile as a byte array.
		byteValue, _ := ioutil.ReadAll(jsonFile)
		// initialise QandA array
		var answers Answers
		json.Unmarshal(byteValue, &answers)

		for index, answer := range ints {
			// check if answer is good or not, increment count accordingly
			if answers.Answers[index].Answer == answer {
				goodAnswers += 1
			} else {
				badAnswers += 1
			}
		}
		fmt.Println("goodAnswers:", goodAnswers)
		fmt.Println("badAnswers:", badAnswers)
		GoodAnswerScores = append(GoodAnswerScores, goodAnswers)
		fmt.Println("GoodAnswerScores:", GoodAnswerScores)
		fmt.Println("Length:", len(GoodAnswerScores))
		// submit score to global variable

		// issue percentage
		// You scored higher than 60% of all quizzers;
	}
}

// END - POST Question Number & Answer Id

func main() {
	handleRequests()
}
