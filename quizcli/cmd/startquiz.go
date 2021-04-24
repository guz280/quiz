package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// startquizCmd represents the startquiz command
var startquizCmd = &cobra.Command{
	Use:   "startquiz",
	Short: "Start the whole quiz",
	Long:  `Start the whole quiz. Will guide you through all the steps`,

	Run: func(cmd *cobra.Command, args []string) {

		// get all questionIds
		responseBytes := getQuestionIds()
		questionids := QuestionIds
		if err := json.Unmarshal([]byte(responseBytes), &questionids); err != nil {
			fmt.Printf("ssssCould not unmarshal reponseBytes. %v", err)
		}

		// declare answer array & get questions one by one waiting for response from user
		answers := []string{}
		for i := 0; i < len(questionids); i++ {
			var args []string
			idString := strconv.Itoa(questionids[i].Id)
			args = append(args, idString)
			// get question and answers
			getQuestionAndAnswerById(args)

			// wait for answer from user
			answer := ""
			fmt.Printf("Enter answer for question %v: ", questionids[i].Id)
			fmt.Println()
			fmt.Scanln(&answer)

			answers = append(answers, answer)
		}
		// send results
		postResults(answers)
	},
}

func init() {
	rootCmd.AddCommand(startquizCmd)
}
