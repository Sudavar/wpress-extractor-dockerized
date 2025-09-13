package main

import (
	"fmt"
	"os"

	"github.com/yani-/wpress"
)

func main() {

	if ( len(os.Args) == 2 ) {
		pathTofile := os.Args[1]
		fmt.Printf("Attempting to extract %s", pathTofile);

		archiver, _ := wpress.NewReader(pathTofile)

		_ , err := archiver.Extract();
		if ( err != nil ) {
			fmt.Println("Error = ");
			fmt.Println(err);
		} else {
			fmt.Println("All done!");
		}
	} else {
		fmt.Println("Please provide the name/path of .wpress file as the first argument");
	}
}
