package cmd

import (
	"context"
	"fmt"

	"github.com/campoy/justforfunc/31-grpc/todo"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of the tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		filterDone, err := cmd.Flags().GetBool("todo_only")
		if err != nil {
			return err
		}
		return list(context.Background(), filterDone)
	},
}

func init() {
	listCmd.Flags().BoolP("todo_only", "t", false, "Show only tasks that are not completed yet")

	rootCmd.AddCommand(listCmd)
}

func list(ctx context.Context, filterDone bool) error {
	l, err := client.List(ctx, &todo.Void{})
	if err != nil {
		return fmt.Errorf("could not fetch tasks: %v", err)
	}
	for _, t := range l.Tasks {
		if t.Done {
			if filterDone {
				continue
			}
			fmt.Printf("👍")
		} else {
			fmt.Printf("😱")
		}
		fmt.Printf(" %s\n", t.Text)
	}
	return nil
}
