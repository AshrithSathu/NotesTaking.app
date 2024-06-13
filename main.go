package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/note/note"
)

func main(){
	title,content := getNoteData();

	userNote , err :=note.NewNote(title,content); 

	if err != nil{
		fmt.Println("Error creating note");
		return;
	}

	userNote.DisplayNote(); 
	err = userNote.SaveNote();
	if err != nil{
		fmt.Println("Error saving note");
		return;
	}
	fmt.Println("Note saved successfully");
}


func getNoteData()(string , string){
	title:= getUserInput("get title := ");

	content := getUserInput("get content := ");


	return title,content;
}

func getUserInput(prompt string) (string){
	fmt.Print(prompt);
	
	var text string;
	reader := bufio.NewReader(os.Stdin);

	text, err := reader.ReadString('\n');

	if err != nil{
		fmt.Println("Error reading input");
		text = "";
		return "";
	}

	return text;

}