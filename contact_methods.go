package main

import (
	"io/ioutil"
	"os"
	"strings"
)

type contact struct{ Dir, Tel, Name, SdName string }

func (c *contact) showContact() error {
	text, err := ioutil.ReadFile(c.Dir + c.Tel)
	if err != nil {
		return err
	}
	textAr := strings.Fields(string(text))
	c.Name = textAr[0]
	c.SdName = textAr[1]
	return nil
}

func (c *contact) createContact() error {
	text := []byte(c.Name + " " + c.SdName)
	fileName := c.Dir + c.Tel
	return ioutil.WriteFile(fileName, text, 0777)
}

func (c *contact) updateContact() error {
	err := c.showContact()
	if err != nil {
		return err
	}
	err = c.deleteContact()
	if err != nil {
		return err
	}
	return c.createContact()
}

func (c *contact) deleteContact() error {
	return os.Remove(c.Dir + c.Tel)

}

func indexContact(dir string) (contacts []contact, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		c := contact{Dir: dir, Tel: file.Name()}
		if err := c.showContact(); err != nil {
			return contacts, err
		}
		contacts = append(contacts, c)

	}
	return
}
