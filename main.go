package main

import(

	"fmt"
)


func handleMain( w http.ResponseWriter, r *http.Request){

	_, err := w.Write([]byte("handleMain"))
	if err != nil{
		fmt.Println("error handlemain", err)
		
	}
}