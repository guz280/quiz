package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	//fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	// return list of question ids
	http.HandleFunc("/questions", returnQuestionIds)
	http.HandleFunc("/question", returnQuestionAndAnswers)
	log.Fatal(http.ListenAndServe(":1000", nil))
}

// start - GET Question Ids
type QuestionId struct {
	Id int
}

// initialise
var QuestionIds []QuestionId

// questions ids
func returnQuestionIds(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnQuestionIds")
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

// end - GET Question Ids

// start - GET Question & Answer By Id
type QandAs struct {
	QandAs []QandA `json:"QandA"`
}
type QandA struct {
	//Id       int    `json:"id"`
	Question string `json:"question"`
	Answers  string `json:"answers"`
}

func returnQuestionAndAnswers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnQuestionAndAnswers")

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
		// initialis QandA array
		var qandas QandAs
		json.Unmarshal(byteValue, &qandas)
		fmt.Println(qandas.QandAs[questionId-1])
		// jq := gojsonq.New().File("./QuestionAndAnswers.json")
		// res := jq.From("QandA").Where("id", "==", id).Get()
		// fmt.Println(res)

		//json.NewEncoder(w).Encode(QuestionIds)
	}
}

// end - GET Question & Answer By Id

func main() {
	handleRequests()
}
