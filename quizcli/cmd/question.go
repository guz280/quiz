/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// questionCmd represents the question command
var questionCmd = &cobra.Command{
	Use:   "question",
	Short: "Get your question & answrers",
	Long:  `Get your question & answers by entering id`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("question called")
		getQuestionAndAnswerById(args)
	},
}

func init() {
	rootCmd.AddCommand(questionCmd)
}

type QandA struct {
	Id       int
	Question string
	Answers  []AllAnswers
	Answerid int
}
type AllAnswers struct {
	Id     int
	Answer string
}

// initialise
var QuestionAndAnswers QandA

func getQuestionAndAnswerById(args []string) {
	url := "http://localhost:1000/question?id=" + args[0]
	responseBytes := getQuestionsAndAnswersData(url)
	qandas := QuestionAndAnswers
	if err := json.Unmarshal(responseBytes, &qandas); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		return
	}

	fmt.Println("Q:", qandas.Question)
	for index := range qandas.Answers {
		fmt.Println("A:", qandas.Answers[index].Id, "-", qandas.Answers[index].Answer)
	}
}

func getQuestionsAndAnswersData(baseAPI string) []byte {
	retError := make([]byte, 0)
	request, err := http.NewRequest(
		http.MethodGet, //method
		baseAPI,        //url
		nil,            //body
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
