package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	wpress "wpress-extractor/inc"
)

func usage(message string) {
  if message != "" {
    fmt.Println(message);
  }
  fmt.Println("Usage: wpress-extractor --input <path/to/wpress/file> --output <path/to/output/directory>");
}

func main() {
  if ( len(os.Args) < 2 ) {
    fmt.Println("Show help. Please provide the name/path of wpress file as an argument.");
    return;
  }

  pathToFile := "";
  outputPath := "/tmp/extractions";

  for i := 0; i < len(os.Args)-1; i++ {
    switch os.Args[i] {
    case "--input":
      pathToFile = os.Args[i+1];
    case "--output":
      outputPath = os.Args[i+1];
    }
  }

  if pathToFile == "" {
    usage("Please provide the name/path of wpress file as an argument.");
    return;
  }

  if outputPath == "" {
    usage("Please provide the output path as an argument.");
    return;
  }

  /* Check if file exists */
  _, file_exists_err := os.Stat(pathToFile);
  if os.IsNotExist(file_exists_err) {
    fmt.Printf("Provided archive %s does not exist.", pathToFile);
    return;
  }

  // Get just the base file name and trim the extension (after the last dot)
	fileName := filepath.Base(pathToFile);
	name := strings.TrimSuffix(fileName, filepath.Ext(fileName));
  outputPath = outputPath + "/" + name;

  fmt.Printf("Attempting to extract %s\n", pathToFile);
  archiver, _ := wpress.NewReader(pathToFile);

  _ , extract_err := archiver.Extract(outputPath);
  if ( extract_err != nil ) {
    fmt.Println("Error = ");
    fmt.Println(extract_err);
  } else {
    fmt.Printf("All done!\nThe files are extracted under: %s\n", outputPath);
  }
}
