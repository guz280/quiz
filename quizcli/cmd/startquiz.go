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
		fmt.Println("startquiz called")
		// get all questionIds
		responseBytes := getQuestionIds()
		questionids := QuestionIds
		if err := json.Unmarshal([]byte(responseBytes), &questionids); err != nil {
			fmt.Printf("ssssCould not unmarshal reponseBytes. %v", err)
		}

		// declare answer array & get questions one by one
		answers := []string{}
		for i := 0; i < len(questionids); i++ {
			var args []string
			idString := strconv.Itoa(questionids[i].Id)
			args = append(args, idString)
			getQuestionAndAnswerById(args)
			answer := ""
			fmt.Printf("Enter answer for question %v: ", questionids[i].Id)
			fmt.Println()
			fmt.Scanln(&answer)

			answers = append(answers, answer)
		}
		postResults(answers)
	},
}

func init() {
	rootCmd.AddCommand(startquizCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startquizCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startquizCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
