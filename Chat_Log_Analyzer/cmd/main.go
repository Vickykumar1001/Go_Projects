package main

import (
	"fmt"
	"go_assessment/internal/services"
	"log"
)

const filePath = "./data/chat_logs.log"
const spamTimeLimit = 2

func main() {

	// Reading the file
	content, err := services.ReadFile(filePath)
	// if err print error and return
	if err != nil {
		log.Println(err)
		return
	}
	// processing the file and saving it in form of array of struct
	messageDetails, err := services.ProcessLogs(content)
	if err != nil {
		log.Println(err)
		return
	}
	// Get Top 10 Words
	topWords := services.GetTopWords(messageDetails)
	fmt.Println("1. Top 10 most used words are: ")
	for idx, word := range topWords {
		fmt.Printf("%d) %s    ", idx+1, word)
	}

	// Get Top 5 most active users
	fmt.Println("\n\n2. Top 5 most active user are: ")
	topUsers := services.GetMostActiveUser(messageDetails)
	for idx, user := range topUsers {
		fmt.Printf("%d) %s    ", idx+1, user)
	}

	// Average Message Length Per User
	fmt.Println("\n\n3. Average Message Length Per User: ")
	avgLength := services.AverageMessageLengthPerUser(messageDetails)
	for user, length := range avgLength {
		fmt.Printf("%s:  %.2f    ", user, length)
	}

	// Spam Detection
	fmt.Println("\n\n4. Spammers are: ")
	spammers := services.DetectSpammers(messageDetails, spamTimeLimit)
	for user, count := range spammers {
		fmt.Printf("%s: spam count(%d)    ", user, count+1)
	}

}
