package main

import (
	"GO_Tokenizer/GO_Tokenizer"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Settings we pull from PGPT
	settings := GO_Tokenizer.Settings{
		MaxNewTokens:  1024,
		ContextWindow: 7800,
	}

	jsonData, err := ioutil.ReadFile("test1.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}
	systemPrompt := "Start the response with saying Summary: then Summarize the following Zendesk ticket with bullet points and have the format be in MARKDOWN format., using the least words possible to provide a broad understanding of what has happened and what needs to happen. Include headers for Summary, What Has Been Done, What Needs to Happen, and Issue Suggestions. Gage the tempature of the tech and the customer overall response. If unknow state unkown. List any negative comments. Lastly, if the text was translated to english state: Translated text to english ( from whatever lanaguage)."
	userPrompt := "Summarize the document. Please provide extreme details on what has happened during this ticket. so that a tech that has not worked on this ticket understands what has happened and what may need to happen."

	// Tokenize and check for warnings
	_, _, additionalInfo, err := GO_Tokenizer.TokenizeInput(string(jsonData), systemPrompt, userPrompt, settings)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Output the results
	fmt.Printf("Additional Info: %+v\n", additionalInfo)
}
