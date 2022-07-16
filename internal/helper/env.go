package helper

import (
	"bufio"
	"fmt"
	"forum/internal/helper/constraints"
	"os"
	"strings"
)

// setting up environment
func SetEnv() {
	err := readEnv()
	if err != nil {
		return
	}

	value, found := os.LookupEnv("DB_PATH")
	if found && len(value) > 0 {
		constraints.DB_PATH = value
	}

	value, found = os.LookupEnv("DB_NAME")
	if found && len(value) > 0 {
		constraints.DB_NAME = value
	}

	value, found = os.LookupEnv("DB_USERNAME")
	if found && len(value) > 0 {
		constraints.DB_USERNAME = value
	}

	value, found = os.LookupEnv("DB_PASSWORD")
	if found && len(value) > 0 {
		constraints.DB_PASSWORD = value
	}
}

// read the env file
// for future: maybe add overwriting env
func readEnv() error {
	file, err := os.Open("./configs/.env")
	if err != nil {
		return fmt.Errorf("helper.readEnv: %w", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		envs := strings.Split(text, "=")
		if len(envs) != 2 {
			// logging
			continue
		}

		// check for already existing env
		if _, found := os.LookupEnv(envs[0]); found {
			continue
		}

		os.Setenv(envs[0], envs[1])
	}

	file.Close()
	return nil
}
