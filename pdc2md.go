package main

import (
   "fmt"
   "io/ioutil"
   "encoding/json"
)


type Notebook struct {
	Worksheets []struct {
		Cells []struct {
		   CellType string   `json:"cell_type"`
     	   Level    int      `json:"level"`
		   Source   []string `json:"source"`	
		   Input    []string `json:"input"`
		} `json:"cells"`
	} `json:"worksheets"`
}



func main() {
	
   // Read the raw byte steram from the file
   ipynb, err := ioutil.ReadFile("test.ipynb")
   if err != nil {
      panic(err)
   }

   var notebook Notebook

   err = json.Unmarshal(ipynb, &notebook)

   for _,v := range notebook.Worksheets[0].Cells {
	  if len(v.Source) > 0 {
         fmt.Println(v.Source[0])	
         fmt.Println("\n***\n")
      }
   }
 

}