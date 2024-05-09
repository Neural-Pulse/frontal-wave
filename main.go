package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	ID           int
	Preferences  map[string]int
}

func main() {
	// Checking the existence of data files
	preferenceDataAvailable := checkFile("sample-data.csv")
	salesDataAvailable := checkFile("sample-sales.csv")

	// Choosing the approach based on data availability
	if preferenceDataAvailable {
		fmt.Println("Preference score data available. Using the preference score approach.")
		usePreferenceScoreApproach()
	} else if salesDataAvailable {
		fmt.Println("Sales data available. Using the sales-based approach.")
		useSalesBasedApproach()
	} else {
		fmt.Println("No suitable approach found. Check the data files.")
	}
}

func checkFile(fileName string) bool {
	_, err := os.Stat(fileName)
	if err == nil {
		return true // File found
	} else if os.IsNotExist(err) {
		return false // File not found
	} else {
		fmt.Println("Error checking the file:", err)
		return false
	}
}

func usePreferenceScoreApproach() {
	// Open the preference data file
	file, err := os.Open("sample-data.csv")
	if err != nil {
		fmt.Println("Error opening preference data file:", err)
		return
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading preference data:", err)
		return
	}

	// Process preference data
	users := make(map[int]*User)
	for _, record := range records[1:] { // Skip header
		userID, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Println("Error parsing user ID:", err)
			continue
		}

		product := record[1]
		score, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Println("Error parsing preference score:", err)
			continue
		}

		if _, ok := users[userID]; !ok {
			users[userID] = &User{ID: userID, Preferences: make(map[string]int)}
		}

		users[userID].Preferences[product] = score
	}

	// Generate recommendations based on preference score
	for userID, user := range users {
		fmt.Printf("Recommendations for user %d:\n", userID)
		for product, score := range user.Preferences {
			fmt.Printf("- Product: %s, Preference Score: %d\n", product, score)
		}
		fmt.Println()
	}
}

func useSalesBasedApproach() {
	// Open the sales data file
	file, err := os.Open("sales.csv")
	if err != nil {
		fmt.Println("Error opening sales data file:", err)
		return
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading sales data:", err)
		return
	}

	// Process sales data
	users := make(map[int]*User)
	for _, record := range records[1:] { // Skip header
		userID, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Println("Error parsing user ID:", err)
			continue
		}

		product := record[1]

		if _, ok := users[userID]; !ok {
			users[userID] = &User{ID: userID, Preferences: make(map[string]int)}
		}

		users[userID].Preferences[product]++
	}

	// Generate recommendations based on sales data
	for userID, user := range users {
		fmt.Printf("Recommendations for user %d:\n", userID)
		for product, sales := range user.Preferences {
			fmt.Printf("- Product: %s, Sales: %d\n", product, sales)
		}
		fmt.Println()
	}
}
