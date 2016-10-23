package main

import (
	"fmt"

	"github.com/thylong/mongo_mock_go_example/model"
)

func main() {
	session := model.NewMockSession()
	documents, _ := session.DB("test").C("other_test").GetMyDocuments()

	fmt.Println(documents)
}
