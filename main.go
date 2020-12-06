package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	//"golang.org/x/crypto/bcrypt" // todo -> generate a key from a passphrase
	//"gopkg.in/yaml.v2" // todo -> read yaml configuration file
	//"github.com/google/uuid" // todo -> need a way to specify remote storage sandbox
)

type bcryptParameters struct {
	Cost int
}

type configuration struct {
	Password string
	Bcrypt   bcryptParameters
}

func generateKeyFromPassphrase() {

}

func readDefaultConfigurations() (*configuration, error) {
	_, err := os.Stat("configuration.yml")

	// todo -> consider generating a default configuration on execution when none found
	// pros; user wouldn't have to configure anything, just run the command
	// cons: generating keys that the user doesn't know about...
	if err != nil {
		fmt.Printf("no default configuration file found")
		return nil, errors.New("no default configuration file")
	}

	// todo -> unmarshall the yaml file
	return nil, nil
}

func main() {
	configuration, err := readDefaultConfigurations()

	if configuration == nil {
		fmt.Printf("no default configuration file found and we do not support generating one yet")
		return
	}

	var pathToFile string = os.Args[1]

	var fileInfo os.FileInfo

	fileInfo, err = os.Stat(os.Args[1])

	if err != nil {
		fmt.Printf("file '%s' could not be opened\n", os.Args[1])
		return
	}

	fmt.Printf("file '%s' exists -> %s\n", os.Args[1], fileInfo.Name())

	var f io.Reader

	f, err = os.Open(pathToFile)

	if err != nil {
		fmt.Printf("file '%s' could not be opened\n", fileInfo.Name())
		return
	}

	var fileReader = bufio.NewReader(f)

	data, err := fileReader.ReadString('\n')


	fmt.Printf("data read from file...\n%s", data)
}
