package services

import (
	"fmt"
	"go_assessment/internal/model"
	"os"
	"sort"
	"strings"
	"time"
)

const (
	TimeStampStartIdx = 1
	TimeStampEndIdx   = 20
	UserNameStartIdx  = 22
	UserNameEndIdx    = 29
	MessageStartIdx   = 31
)

// reading file from the log
func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	return string(data), nil

}

// converting the log string to array of meessageDetail struct
func ProcessLogs(fileContent string) ([]model.MessageDetail, error) {
	//spliting the lines
	lines := strings.Split(fileContent, "\n")
	var messageDetails []model.MessageDetail

	const timestampFormat = "2006-01-02 15:04:05"

	for _, line := range lines {
		if line == "" { //if blank line continue
			continue
		}
		// slicing time part
		timestampStr := line[TimeStampStartIdx:TimeStampEndIdx]
		timestamp, err := time.Parse(timestampFormat, timestampStr) //converting to time data type
		if err != nil {
			return nil, fmt.Errorf("could not parse timestamp %s: %v", timestampStr, err)
		}
		// slicing username
		user := line[UserNameStartIdx:UserNameEndIdx]

		// slicing message
		message := line[MessageStartIdx:]
		// creating object of messageDetail
		messageDetail := model.NewMessageDetail(timestamp, user, message)
		messageDetails = append(messageDetails, *messageDetail) // adding it to array
	}
	return messageDetails, nil
}

func GetTopWords(messageDetails []model.MessageDetail) []string {
	wordCount := make(map[string]int)

	for _, messageDetail := range messageDetails {
		// spliting to get array of words
		words := strings.Split(messageDetail.Message, " ")
		for _, word := range words {
			// counting freq of each word
			wordCount[strings.ToLower(word)]++
		}
	}

	// converting map to slice so that we can sort it
	type wordFreq struct {
		Word  string
		Count int
	}

	var wordFreqs []wordFreq
	for word, count := range wordCount {
		wordFreqs = append(wordFreqs, wordFreq{Word: word, Count: count})
	}
	// sorting it
	sort.Slice(wordFreqs, func(i, j int) bool {
		return wordFreqs[i].Count > wordFreqs[j].Count
	})

	var topWords []string
	// storing only top 10 words
	for i := 0; i < 10 && i < len(wordFreqs); i++ {
		topWords = append(topWords, wordFreqs[i].Word)
	}

	return topWords
}

func GetMostActiveUser(messageDetails []model.MessageDetail) []string {
	userCount := make(map[string]int)

	for _, messageDetail := range messageDetails {
		// counting freq of message of each user
		userCount[messageDetail.User]++
	}

	// converting map to slice so that we can sort it
	type userFreq struct {
		user  string
		Count int
	}

	var userFreqs []userFreq
	for user, count := range userCount {
		userFreqs = append(userFreqs, userFreq{user: user, Count: count})
	}
	// sorting it
	sort.Slice(userFreqs, func(i, j int) bool {
		return userFreqs[i].Count > userFreqs[j].Count
	})

	var topusers []string
	// storing only top 5 users
	for i := 0; i < 5 && i < len(userFreqs); i++ {
		topusers = append(topusers, userFreqs[i].user)
	}

	return topusers

}

func AverageMessageLengthPerUser(messages []model.MessageDetail) map[string]float64 {
	userStats := make(map[string]struct {
		totalLength int
		count       int
	})

	for _, msg := range messages {
		userData := userStats[msg.User]
		userData.totalLength += len(msg.Message)
		userData.count++
		userStats[msg.User] = userData
	}

	averageLength := make(map[string]float64)
	for user, data := range userStats {
		averageLength[user] = float64(data.totalLength) / float64(data.count)
	}

	return averageLength
}

func DetectSpammers(messages []model.MessageDetail, thresholdSeconds int) map[string]int {
	userSpamCount := make(map[string]int)
	lastMessageTime := make(map[string]time.Time)

	for _, msg := range messages {
		lastTime, ok := lastMessageTime[msg.User]
		if ok {
			timeDiff := int(msg.TimeStamp.Sub(lastTime).Seconds())
			if timeDiff < thresholdSeconds {
				userSpamCount[msg.User]++
			}
		}
		lastMessageTime[msg.User] = msg.TimeStamp
	}

	return userSpamCount
}
