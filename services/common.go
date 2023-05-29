package services

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func checkName(name *string) error {
	for _, r := range *name {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && r != ' ' {
			return fmt.Errorf("bad request, check name: it cannot contain special characters and nums")
		}
	}

	stringName := strings.Fields(*name)

	if len(stringName) != 1 {
		return fmt.Errorf("bad request, name, last name or middle name has too many words")
	}
	*name = strings.ReplaceAll(*name, " ", "")
	*name = cases.Title(language.Und).String(*name)

	return nil
}
