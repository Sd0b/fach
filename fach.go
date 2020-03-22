package main

import "io/ioutil"
import "fmt"
import "net/http"
import "log"
import "os"
import "bufio"


  

func main(){
	file,err := os.Open("domains.txt")
if err != nil {
        fmt.Println("File reading error", err)
       }

    scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var links []string
 
	for scanner.Scan() {
		links = append(links, scanner.Text())
	}
 
	file.Close()


      
      /*"""
    Handle a URL.
    """*/

for _, link := range links  {


    var default_protocol = "http"
    full_link := string(default_protocol) + string("://") + string(link)

        req , err := http.Get(full_link)
        if err != nil {
           log.Print(err)
           os.Exit(1)
           }
         
   
   fmt.Println("link: --> ",full_link)
   fmt.Println(req)

   defer req.Body.Close()
   body, err := ioutil.ReadAll(req.Body)
   fmt.Println(body)

 

	}  


}








