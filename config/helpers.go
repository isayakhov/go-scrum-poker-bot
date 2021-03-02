package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getStrEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func getIntEnv(key string, defaultValue string) int {
	value, err := strconv.Atoi(getStrEnv(key, defaultValue))
	if err != nil {
		panic(fmt.Sprintf("Incorrect env value for %s", key))
	}

	return value
}

func getListStrEnv(key string, defaultValue string) []string {
	value := []string{}
	for _, item := range strings.Split(getStrEnv(key, defaultValue), ",") {
		value = append(value, strings.TrimSpace(item))
	}
	return value
}
