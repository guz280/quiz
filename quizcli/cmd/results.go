package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// resultsCmd represents the results command
var resultsCmd = &cobra.Command{
	Use:   "results",
	Short: "Send all answers & compare score",
	Long:  `"Send all answers & compare score.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("results called")
		postResults(args)
	},
}

func init() {
	rootCmd.AddCommand(resultsCmd)
}

type Message struct {
	GoodAnswers int
	BadAnswers  int
	Compared    string
}

// initialise
var MessageReceived Message

func postResults(args []string) {
	if len(args) != 7 {
		fmt.Println("Please enter 7 answers only")
	} else {
		url := "http://localhost:1000/results"
		responseBytes := postResultsGetStatistics(url, args)
		fmt.Println(string(responseBytes))

		msg := MessageReceived
		if err := json.Unmarshal(responseBytes, &msg); err != nil {
			fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		}

		fmt.Println("Good Answers:", msg.GoodAnswers)
		fmt.Println("Bad Answers:", msg.BadAnswers)
		fmt.Println("Compared:", msg.Compared)
	}
}

type Answers struct {
	Answers []Answer `json:"answers"`
}
type Answer struct {
	Questionid int `json:"questionid"`
	Answerid   int `json:"answerid"`
}

// initialise
var Results Answers

func postResultsGetStatistics(baseAPI string, args []string) []byte {

	//{"answers":[{"questionId":1,"answerId":3},{"questionId":2,"answerId":3},{"questionId":3,"answerId":1},{"questionId":4,"answerId":1},{"questionId":5,"answerId":1},{"questionId":6,"answerId":1},{"questionId":7,"answerId":3}]}
	for i := 0; i < len(args); i++ {
		var answer Answer
		id, err := strconv.Atoi(args[i])
		if err != nil {
			fmt.Println("error convert to int", err)
		}
		answer.Answerid = id
		answer.Questionid = i + 1

		Results.Answers = append(Results.Answers, answer)
	}
	fmt.Println(Results)
	output, err := json.Marshal(Results)
	r := strings.NewReader(string(output))

	request, err := http.NewRequest(
		http.MethodPost, //method
		baseAPI,         //url
		r,               //body
	)

	if err != nil {
		log.Printf("Could not request to get question ids. %v", err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}
	return responseBytes
}
