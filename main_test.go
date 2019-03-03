package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var (
	c  = contact{"./test_data/", "79153423434", "Vasya", "Tupichkin"}
	c1 = contact{"./test_data/", "79456754744", "Irina", "Nyashmyash"}
)

func CreateTestEnvironment() {
	if err := os.Mkdir(c.Dir, 0777); err != nil {
		fmt.Println("Creating test env with error", err)
		os.Exit(2)
	}
}

func DeleteTestEnvironment() {
	if err := os.Remove(c.Dir); err != nil {
		fmt.Println("Deleting test env with error", err)
		os.Exit(2)
	}
}
func TestMain(m *testing.M) {
	CreateTestEnvironment()
	defer DeleteTestEnvironment()
}

func TestCreateContact(t *testing.T) {
	c.createContact()
	expectedFileContent := c.Name + " " + c.SdName
	fileContent, err := ioutil.ReadFile(c.Dir + c.Tel)
	if err != nil {
		t.Error("Can't read created file with err: ", err)
	} else if string(fileContent) != expectedFileContent {
		t.Error(
			"For ", c,
			"Expected ", expectedFileContent,
			"Got ", fileContent,
		)
	}
	os.Remove(c.Dir + c.Tel)
}
func TestShowContact(t *testing.T) {
	c.createContact()
	c2 := contact{Dir: c.Dir, Tel: c.Tel}
	c2.showContact()
	if c != c2 {
		t.Error(
			"Expected", c,
			"Got", c2,
		)
	}
	os.Remove(c.Dir + c.Tel)
}
func TestUpdateContact(t *testing.T) {

}
