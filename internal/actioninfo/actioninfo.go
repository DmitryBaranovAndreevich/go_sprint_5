package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, str := range dataset {
		err := dp.Parse(str)

		if err != nil {
			log.Printf("error parsing data: %v", err)
			continue
		}

		info, err := dp.ActionInfo()

		if err != nil {
			log.Printf("error getting action info: %v", err)
			continue
		}

		fmt.Println(info)
	}
}
