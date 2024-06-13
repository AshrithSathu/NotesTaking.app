package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct{
	Title string `json:"title"`;
	Content string `json:"content"`;
	CreatedAt time.Time `json:"created_at"`;
}

func (note Note) DisplayNote(){
	fmt.Println("Title: ",note.Title);
	fmt.Println("Content: ",note.Content); 
}

func (note Note) SaveNote() (error){
	fmt.Println("Saving note");
	filename := strings.ReplaceAll(note.Title," ","_");
	//convert filename to lowercase and add .json extension and time stamp to avoid same file name
	//add date and time also to avoid same file name
	filename = strings.ToLower(filename) + "_" + note.CreatedAt.Format("2006-01-02_15:04:05") + ".json";

	jsonText, err := json.Marshal(note);

	if err != nil{
		fmt.Println("Error marshalling note");
		return err;
	} 
	//save note to file in json format

	return os.WriteFile(filename,jsonText, 0644);
}

func NewNote(title string , content string ) (Note, error) {

	if title == "" || content == ""{
		fmt.Println("Invalid input");
		return Note{}, errors.New("invalid input");
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil;
}

