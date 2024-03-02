package config

import (
	"strconv"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Port   int
	DBhost string
	DBuser string
	DBpass string
	DBname string
	DBport int
}

var config *Configuration

func new() *Configuration {
	return &Configuration{
		Port:   3000,
		DBhost: "localhost",
		DBuser: "root",
		DBpass: "",
		DBname: "test",
		DBport: 5432,
	}
}

func Get() *Configuration {
	if config == nil {
		config = load()
	}
	return config
}

func load() *Configuration {
	result := new()
	var vars map[string]string
	vars, err := godotenv.Read()
	if err != nil {
		panic(err)
	}
	if value := vars["APP_PORT"]; len(value) > 0 {
		if intVal, err := strconv.Atoi(value); err == nil {
			result.Port = intVal
		}
	}
	if value := vars["DB_PORT"]; len(value) > 0 {
		if intVal, err := strconv.Atoi(value); err == nil {
			result.DBport = intVal
		}
	}
	if value := vars["DB_USER"]; len(value) > 0 {
		result.DBuser = value
	}
	if value := vars["DB_PASS"]; len(value) > 0 {
		result.DBpass = value
	}
	if value := vars["DB_HOST"]; len(value) > 0 {
		result.DBhost = value
	}
	if value := vars["DB_NAME"]; len(value) > 0 {
		result.DBname = value
	}
	return result
}
