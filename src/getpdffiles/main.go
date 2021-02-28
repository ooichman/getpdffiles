package main

import (
  "os"
  "fmt"
	"github.com/secsy/goftp"
	"bytes"
  "log"
  "io/ioutil"
  "time"
//  "path"
)

func getEnv(key, fallback string) string {
        var value string
        value, exists := os.LookupEnv(key)
        if !exists {
                value = fallback
        }
        return value
}


func main() {

  config := goftp.Config {
    User:               "anonymous",
    Password:           "root@local.me",
    ConnectionsPerHost: 21,
    Timeout:            30 * time.Second,
    Logger:             os.Stderr,
  }

   ftpServer := getEnv("FTP_SERVER", "localhost" )
	 client, dailErr := goftp.DialConfig(config, ftpServer)

	 if dailErr != nil {
       log.Fatal(dailErr)
   		 panic(dailErr)
	 }

  dir := getEnv("FTP_DIRECTORY", "/" )
  download := getEnv("DOWNLOAD_FILES", "no" )
  //   files , err := client.ReadDir(dir)
   files , err := client.ReadDir(dir)
   
   if err != nil {
      panic(err)
   }

   for  _ , file := range files {
 //   if file.IsDir() {
 //      path.Join(dir, file.Name())
 //   } else {
          if download == "yes" {
            ret_file := file.Name()
            fmt.Println("Retrieving file: ", ret_file)
            buf := new(bytes.Buffer)
            fullPathFile := dir + ret_file
            rferr := client.Retrieve(fullPathFile, buf)

            if rferr != nil {
                panic(rferr)
            }

            fmt.Println("writing data to file", ret_file)
            fmt.Println("Opening file", ret_file,"for writing")
            w , _ := ioutil.ReadAll(buf)
            ferr :=  ioutil.WriteFile(ret_file, w , 0644)

            if ferr != nil {
                log.Fatal(ferr)
                panic(ferr)
            } else {
                fmt.Println("Writing", ret_file ," completed")
            }
          } else {
            fmt.Println("the file is:", file.Name())
          }
 //   }
   } 
}