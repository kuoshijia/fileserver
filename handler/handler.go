package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	//if r.Method == "GET" {
	//
	//}
	switch r.Method {
		case "GET":
			data,err := ioutil.ReadFile("./static/view/index.html")
			if err!=nil {
				io.WriteString(w,"internal server error!")
				return
			}
			io.WriteString(w,string(data))
			return
		case "POST":
			file,head,err := r.FormFile("file")
			if err!=nil {
				fmt.Printf("no input file data found,err:%s",err.Error())
				return
			}
			defer file.Close()
			newFile,err := os.Create("tmp/" + head.Filename)
			if err!=nil{
				fmt.Printf("cannot create new file:%s,error: %s",head.Filename,err.Error())
				return
			}
			defer newFile.Close()
			_, err2:=io.Copy(newFile,file)
			if err2!=nil{
				fmt.Printf("error while saving file: %s",err2.Error())
				return
			}
			http.Redirect(w,r,"/file/upload/success",http.StatusFound)
	}
}

func UploadSuccessHandler(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w,"Upload finished!")
}
