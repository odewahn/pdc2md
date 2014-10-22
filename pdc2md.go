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
            val = val + fmt.Sprintf("\n%s %s\n\n", strings.Repeat("#",c.Level + 1), c.Source[0])	
         case "markdown":
	        val = val + "\n"
	        for _,s := range c.Source {
		       val = val + s
	        }
	        val = val + "\n"
        case "code":
	        val = val + "\n\n``` python\n"
	        for _,s := range c.Input {
		       val = val + s
	        }
	        val = val + "\n```\n\n"
	  }

   }
   return val	
}




func main() {
	
   BASE_PATH := "/Users/odewahn/Desktop/python-data-cookbook/notebooks"
   OUTPUT_PATH := "/Users/odewahn/Desktop/pdc-generated"
   base_dirs, _ := ioutil.ReadDir(BASE_PATH)
 
   for _,d := range base_dirs {
	  // pull all directories that have a "-" in the name
	  dir_name := strings.Split(d.Name(),"-")
	  if (d.IsDir() && len(dir_name) > 1) {
         chapter_fn := d.Name() + ".md"  // Root filename we're going to glob together
         chapter_title := "# " + strings.Join(dir_name[1:len(dir_name)], " ") + "\n\n"
         chapter_contents := chapter_title        
        
         // Process ipynb files in the directory and convert them into sections in the chapter
         sections, _ := ioutil.ReadDir(BASE_PATH + "/" + d.Name())
         for _,f := range sections {
	        test_fn := strings.Split(f.Name(),".")
            if (test_fn[len(test_fn)-1] == "ipynb") {
	           section_fn := BASE_PATH + "/" + d.Name() + "/" + f.Name()
	
			   // Grab the contents of the section
			   ipynb, err := ioutil.ReadFile(section_fn)
			   if err != nil {
			      panic(err)
			   }
			
			   // unmarshall the json data into a notebook structure
			   var notebook Notebook
			   err = json.Unmarshal(ipynb, &notebook)
			
			   // append the contents to the chapter output
	           chapter_contents = chapter_contents + md(notebook) + "\n"      
            }
         }
         err := ioutil.WriteFile(OUTPUT_PATH + "/" + chapter_fn, []byte(chapter_contents), 0644)
	     if err != nil {
	        panic(err)
	     }
      }
   }
	




}