package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		return
	} else {
		url := "http://localhost:1000/results"
		responseBytes := postResultsGetStatistics(url, args)

		msg := MessageReceived
		if err := json.Unmarshal(responseBytes, &msg); err != nil {
			fmt.Printf("Could not unmarshal reponseBytes. %v", err)
			return
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

	retError := make([]byte, 0)

	for i := 0; i < len(args); i++ {
		var answer Answer
		id, err := strconv.Atoi(args[i])
		if err != nil {
			fmt.Println("error convert to int", err)
			return retError
		}
		answer.Answerid = id
		answer.Questionid = i + 1

		Results.Answers = append(Results.Answers, answer)
	}
	output, err := json.Marshal(Results)
	r := strings.NewReader(string(output))

	request, err := http.NewRequest(
		http.MethodPost, //method
		baseAPI,         //url
		r,               //body
	)

	if err != nil {
		fmt.Printf("Could not request to get question ids. %v", err)
		return retError
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Could not make a request. %v", err)
		return retError
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Could not read response body. %v", err)
		return retError
	}
	return responseBytes
}
