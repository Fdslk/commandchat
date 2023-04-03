/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	commandchat "zqf.com/commandchat/commandchatChannel"
)

var prompt string

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "plain text chat with AI, don't input any sensitive infornmation",
	Long: `When you use this command, you can ask any question to chatGpt. It will output you wanted answer, but keep
	it in you mind, don't upload any personal or sensitive data to it`,
	Run: func(cmd *cobra.Command, args []string) {
		newRequestBytes, err := commandchat.CreateCompletionsRequest(prompt)

		if err != nil {
			fmt.Println("error occurred:", err)
			return
		}

		rawResponse, err := commandchat.Chat(newRequestBytes)

		if err != nil {
			fmt.Println("error occurred:", err)
			return
		}

		response, err := commandchat.CreateCompletionsResponse(rawResponse)

		if err != nil {
			fmt.Println("error occurred:", err)
			return
		}

		fmt.Printf(response.Choices[0].Text)
	},
}

func init() {
	chatCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "the question you want to ask chatgpt")
	chatCmd.MarkFlagRequired("prompt")
	rootCmd.AddCommand(chatCmd)
}
