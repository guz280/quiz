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

func handleRequests() {
	// using "github.com/gorilla/mux" -> but had issues with this library
	// // myRouter := mux.NewRouter().StrictSlash(true)
	// // myRouter.HandleFunc("/", homePage)
	// // myRouter.HandleFunc("/questions", returnQuestionIds)
	// // myRouter.HandleFunc("/results", results).Methods("POST")
	// // myRouter.HandleFunc("/question/{id}", returnQuestionAndAnswers)
	// // log.Fatal(http.ListenAndServe(":1000", myRouter))

	// using "net/http"
	http.HandleFunc("/questions", returnQuestionIds)
	http.HandleFunc("/question", returnQuestionAndAnswers)
	http.HandleFunc("/results", results)
	log.Fatal(http.ListenAndServe(":1000", nil))
}

//
//
//
//
// START - GET Question Ids
type QuestionId struct {
	Id int
}

// initialise
var QuestionIds []QuestionId

// questions ids
// modify to take the ids from the json file not hard coded
func returnQuestionIds(w http.ResponseWriter, r *http.Request) {
	fmt.Println("return question ids called")
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
//
//
//
//
// START - GET Question & Answer By Id
type QandAs struct {
	QandAs []QandA `json:"QandA"`
}
type QandA struct {
	Id       int          `json:"id"`
	Question string       `json:"question"`
	Answers  []AllAnswers `json:"answers"`
	Answerid int          `json:"answerid"`
}
type AllAnswers struct {
	Id     int    `json:"id"`
	Answer string `json:"answer"`
}

func returnQuestionAndAnswers(w http.ResponseWriter, r *http.Request) {

	fmt.Println("return question and answers called")

	var id = r.URL.Query().Get("id")
	questionPassedId, err := strconv.Atoi(id)
	// handle error
	if err != nil {
		fmt.Println(err)
	}

	if (questionPassedId > 7) || (questionPassedId < 1) {
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
		if err := json.Unmarshal(byteValue, &qandas); err != nil {
			fmt.Printf("Could not unmarshal byteValue. %v", err)
		}
		json.NewEncoder(w).Encode(qandas.QandAs[questionPassedId-1])
	}
}

// END - GET Question & Answer By Id
//
//
//
//
// START - POST Question Number & Answer Id

var GoodAnswerScores []int

type Answers struct {
	Answers []Answer `json:"answers"`
}
type Answer struct {
	Questionid int `json:"questionid"`
	Answerid   int `json:"answerid"`
}

type Message struct {
	GoodAnswers int
	BadAnswers  int
	Compared    string
}

// I know that this needs to be done a proper POST, but had issues with library "github.com/gorilla/mux"
func results(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post results called")
	if r.Method != "POST" {
		fmt.Println("Call has to be a POST Method")
	} else {
		//fmt.Println(r.Method)
		// get the body of our POST request
		reqBody, _ := ioutil.ReadAll(r.Body)

		// read json file sent by user
		// initialise answers array
		var userAnswers Answers
		if err := json.Unmarshal(reqBody, &userAnswers); err != nil {
			fmt.Printf("Could not unmarshal reqBody. %v", err)
		}

		var goodAnswersCount, badAnswersCount int = 0, 0
		// loop through array to submit good/bad answers

		// load answers from json
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
		var answers QandAs
		if err := json.Unmarshal(byteValue, &answers); err != nil {
			fmt.Printf("Could not unmarshal byteValue. %v", err)
		}

		for index, answer := range []Answer(userAnswers.Answers) {
			// check if answer is good or not, increment count accordingly
			if answers.QandAs[index].Answerid == answer.Answerid {
				goodAnswersCount += 1
			} else {
				badAnswersCount += 1
			}
		}

		// append to array to keep track of +ve scores
		GoodAnswerScores = append(GoodAnswerScores, goodAnswersCount)

		// calculate +ve percentage
		var count int = 0
		for i := 0; i < len(GoodAnswerScores); i++ {
			if GoodAnswerScores[i] < goodAnswersCount {
				count += 1
			}
		}
		var stat int = (count * 100) / len(GoodAnswerScores)
		var comparedMsg string
		comparedMsg = "You scored higher than " + strconv.Itoa(stat) + "% of all quizzers"

		var m Message
		m = Message{
			GoodAnswers: goodAnswersCount,
			BadAnswers:  badAnswersCount,
			Compared:    comparedMsg,
		}
		json.NewEncoder(w).Encode(m)
	}
}

// END - POST Question Number & Answer Id
//
//
//
//
func main() {
	handleRequests()
}
