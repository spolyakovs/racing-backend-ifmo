package store

import (
	"fmt"
	"strings"
)

func makeArgsString(queryArgs ...string) string {
	return strings.Join(queryArgs, ", ")
}

func makeValsString(queryArgs ...string) string {
	var querySB strings.Builder

	for index, _ := range queryArgs {
		if index == 0 {
			querySB.WriteString(fmt.Sprintf("$%d", index+1))
		} else {
			querySB.WriteString(fmt.Sprintf(", $%d", index+1))
		}
	}

	return querySB.String()
}

func makeUpdateSetString(queryArgs ...string) string {
	var querySB strings.Builder

	for index, arg := range queryArgs {
		if index == 0 {
			querySB.WriteString(fmt.Sprintf("%s = $%d", arg, index))
		} else {
			querySB.WriteString(fmt.Sprintf(", %s = $%d", arg, index))
		}
	}

	return querySB.String()
}
