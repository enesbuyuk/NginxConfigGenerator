package main

import(
	"fmt"
	"os"
	"log"
	"bufio"
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	domainWWWPath   := "/var/www/ex.enesbuyuk.com"
	domain 		:= "ex.enesbuyuk.com, www.ex.enesbuyuk.com"
	domainFilename  := "ex.enesbuyuk.com"
	configContent	:= "server{listen 80;root "+domainWWWPath+";index index.html index.htm;server_name "+domain+";}"
    
	f, err := os.Create("/etc/nginx/sites-available/"+domainFilename)
 	if err != nil {
        	log.Fatal(err)
	}
    	defer f.Close()
    	
    	f.Sync()
    	w := bufio.NewWriter(f)
    	n4, err := w.WriteString(configContent)
    	check(err)
    	fmt.Println("----> "+f.Name()+" successfully created!")
    	fmt.Printf("----> wrote %d bytes\n", n4)

    	w.Flush()

    	if _, err := os.Lstat("/etc/nginx/sites-enabled/"+domainFilename); err == nil {
    	    os.Remove("/etc/nginx/sites-enabled/"+domainFilename)
   	}
    
    	err = os.Symlink("/etc/nginx/sites-available/"+domainFilename, "/etc/nginx/sites-enabled/"+domainFilename)
    	if err != nil {
        log.Fatal(err)
    	}else{
		fmt.Println("----> "+"Symlink successfully created!\n")
    	}
}
