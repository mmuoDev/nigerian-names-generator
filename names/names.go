package names

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

//Person represents names of persons
type Person struct {
	Names []string `json:"names"`
}

//go:embed files/*
var content embed.FS

func IsValidTribe(tribe string) bool {
	tribes := [3]string{"yoruba", "igbo", "hausa"}
	for _, t := range tribes {
		if tribe == t {
			return true
		}
	}
	return false
}

//fileToStruct converts a file to a struct
func fileToStruct(filepath string, s interface{}) error {
	//bb, err := ioutil.ReadFile(filepath)
	
	//bb, e := content.ReadFile(filepath)
	files, err := fs.ReadDir(content, "files")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Print("file name", file.Name())
		if file.Name() == filepath {
			bb, err := content.ReadFile(filepath)
			if err != nil {
				panic(err)
			}
			if err := json.Unmarshal(bb, s); err != nil {
				return errors.Wrap(err, "Unable to unmarshal struct")
			}
		}
	}
	
	// if err != nil {
	// 	return errors.Wrapf(err, "Unable to read file at path=%s", filepath)
	// }
	
	return nil
}

func retrieveNamesFromFiles(tribe, gender string, p Person) (Person, error) {
	switch tribe {
	case "igbo":

		if gender == "male" {
			err := fileToStruct(filepath.Join("files", "igbo_male.json"), &p)
			return p, err
		}
		err := fileToStruct(filepath.Join("files", "igbo_female.json"), &p)
		return p, err
	case "yoruba":
		if gender == "male" {
			err := fileToStruct(filepath.Join("files", "yoruba_male.json"), &p)
			return p, err
		}
		err := fileToStruct(filepath.Join("files", "yoruba_female.json"), &p)
		return p, err
	case "hausa":
		if gender == "male" {
			err := fileToStruct(filepath.Join("files", "hausa_male.json"), &p)
			return p, err
		}
		err := fileToStruct(filepath.Join("files", "hausa_female.json"), &p)
		return p, err
	}
	return p, nil
}

func GetNames(tribe, gender string) ([]string, error) {
	var person Person
	switch tribe {
	case "igbo":
		p, err := retrieveNamesFromFiles(tribe, gender, person)
		if err != nil {
			return person.Names, err
		}
		return p.Names, nil
	case "yoruba":
		p, err := retrieveNamesFromFiles(tribe, gender, person)
		if err != nil {
			return person.Names, err
		}
		return p.Names, nil
	case "hausa":
		p, err := retrieveNamesFromFiles(tribe, gender, person)
		if err != nil {
			return person.Names, err
		}
		return p.Names, nil
	}
	return person.Names, nil
}

func GenerateRandomNames(tribe, gender string, count int) ([]string, error) {
	var res []string
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	isValid := IsValidTribe(tribe)
	if !isValid {
		return res, fmt.Errorf("The tribe, %s is an invalid tribe", tribe)
	}
	names, err := GetNames(tribe, gender)
	if err != nil {
		return res, fmt.Errorf("Something went wrong, err=%v", err)
	}
	for i := 1; i <= count; i++ {
		n := names[rand.Intn(len(names))]
		res = append(res, n)
	}

	return res, nil
}

func TestFxn(s string) bool {
	return false
}
