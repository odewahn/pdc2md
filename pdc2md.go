package main

import (
   "fmt"
   "io/ioutil"
   "encoding/json"
   "strings"
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


func md(n Notebook) string {
   val := ""
   for _,c := range n.Worksheets[0].Cells {
	  switch c.CellType {
         case "heading":
            val = val + fmt.Sprintf("%s %s\n\n", strings.Repeat("#",c.Level), c.Source[0])	
         case "markdown":
	        for _,s := range c.Source {
		       val = val + s
	        }
	        val = val + "\n"
        case "code":
	        val = val + "\n```\n"
	        for _,s := range c.Input {
		       val = val + s
	        }
	        val = val + "\n```\n\n"
	  }

   }
   return val	
}


func main() {
	
   // Read the raw byte steram from the file
   ipynb, err := ioutil.ReadFile("test.ipynb")
   if err != nil {
      panic(err)
   }

   var notebook Notebook

   err = json.Unmarshal(ipynb, &notebook)

   fmt.Println(md(notebook))

}