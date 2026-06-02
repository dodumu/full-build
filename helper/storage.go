package helper

import (
	"encoding/json"
	"os"
)

func SaveUsers(filename string, users []User) error {
	val, err := json.Marshal(users)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, val, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadUsers(filename string) ([]User, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
