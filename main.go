package main

import (
	"os"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func printUsage() {
	fmt.Println(
`
Usage:

	yamlfukr <command> <filename> <args>

---------------------

	Commands:

	update

	yamlfukr update file.yaml key value

	`,
	)

	defer fmt.Println("*** YAML Ain't Markup Language ***", os.Args)

}

func main() {

	// usage yamlfukr update path/to/yaml key value

	if len(os.Args) == 1 {
		printUsage()
		return
	}

	var action string

	switch os.Args[1] {

	case "update":
		action = os.Args[1]

		if len(os.Args) < 5 {
			printUsage()
			return
		}

	case "delete":
		action = os.Args[1]

	default:

		printUsage()
		return

	}

	filepath, filename := parseRawFilename(os.Args[2])

	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")

	if filepath == "" {
		filepath = "."
	}

	viper.AddConfigPath(filepath)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(
				"YAML file not found:",
				strings.Join([]string{filepath, filename}, "/"),
			)
			return
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}

	//fmt.Println(action, env)

	switch action {
	case "update":

		key := os.Args[3]
		value := os.Args[4]

		viper.Set(key, value)
		viper.WriteConfig()
	}

}

func parseRawFilename(rawFilename string) (string, string) {

	var filename string
	for _, pattern := range []string{".yaml", ".yml"} {
		filename = strings.Replace(rawFilename, pattern, "", 1)
	}

	path := strings.Split(filename, "/")
	if len(path) > 1 {
		filename = path[len(path)-1]
	}

	return strings.Join(path[:len(path) - 1], "/"), filename
}
