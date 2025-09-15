package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yani-/wpress"
)

func main() {

  // Initialize outputPath variable with results as default value
  outputPath := "./results"


  if ( len(os.Args) < 2 ) {
    fmt.Println("Please provide the name/path of .wpress file as the first argument");
  }

  pathToFile := os.Args[1]
  /* Check if file exists */
  _, file_exists_err := os.Stat(pathToFile);
  if os.IsNotExist(file_exists_err) {
    fmt.Println("File does not exist");
    return;
  }

  // Get just the base file name
	fileName := filepath.Base(pathToFile)
	// Trim the extension (after the last dot)
	name := strings.TrimSuffix(fileName, filepath.Ext(fileName))
  outputPath = outputPath + "/" + name

  fmt.Printf("Attempting to extract %s\n", pathToFile);
  archiver, _ := wpress.NewReader(pathToFile)

  _ , extract_err := archiver.Extract(outputPath);
  if ( extract_err != nil ) {
    fmt.Println("Error = ");
    fmt.Println(extract_err);
  } else {
    fmt.Printf("All done! The files are extracted under: %s\n", outputPath);
  }

}
