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
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// questionsCmd represents the questions command
var questionsCmd = &cobra.Command{
	Use:   "questions",
	Short: "Get all questions ids",
	Long:  `Get all question Ids`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("questions called")
		getQuestionIds()
	},
}

func init() {
	rootCmd.AddCommand(questionsCmd)
}

type QuestionId struct {
	Id int
}

// initialise
var QuestionIds []QuestionId

func getQuestionIds() {
	url := "http://localhost:1000/questions"
	responseBytes := getQuestionsIdData(url)
	questionids := QuestionIds

	if err := json.Unmarshal([]byte(responseBytes), &questionids); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	fmt.Println("Choose question id from the below")
	for index := range questionids {
		fmt.Println("Question Id: ", questionids[index].Id)
	}
}

func getQuestionsIdData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet, //method
		baseAPI,        //url
		nil,            //body
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
