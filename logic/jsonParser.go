package logic

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/afauble/FollowTracker/models"
)

func ParseFollowerJson(folderPath string) map[string]models.StringListData {
	results := make(map[string]models.StringListData)

	// Read the file content into memory
	data, err := os.ReadFile(folderPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	// Unmarshal JSON into struct
	var followers []models.Follower
	if err := json.Unmarshal(data, &followers); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	for _, f := range followers {
		results[f.GetValue()] = f.StringListData[0]
	}
	return results
}

func ParseFollowingJson(folderPath string) map[string]models.StringListData {
	results := make(map[string]models.StringListData)

	// Read the file content into memory
	data, err := os.ReadFile(folderPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	// Unmarshal JSON into struct
	var following models.Following
	if err := json.Unmarshal(data, &following); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	for _, f := range following.RelationshipsFollowing {
		results[f.GetValue()] = f.StringListData[0]
	}
	return results
}
