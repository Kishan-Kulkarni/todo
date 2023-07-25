package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	args := os.Args
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	path := os.Getenv("PATH")
	if _, err :=os.Stat(path); errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(path)
		if err != nil {
			log.Panic("Failed to create test.txt file")
		}
	}
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		log.Panic("Failed to open test.txt file")
	}
	defer file.Close()
	data := make([]byte, 100000)
	count , _ := file.Read(data)
	str := string(data[:count])
	strArr := strings.Split(str, "\n")
	done := make([]bool, len(strArr))
	if len(args) ==1{
		if count!=0{
			for i := 0; i < len(strArr); i++ {
				temp := fmt.Sprintf("%d.	%s",i+1,  strArr[i])
				if !done[i]{
					color.Red(temp)
				}else{
					color.Green(temp)
				}
			}	
		}
	}else {
		if args[1] == "delete" {
			err =file.Truncate(0)
				if err!=nil{
					log.Panic("Something went wrong erasing test.txt")
				}
				_, err = file.Seek(0,0)	
				if err!=nil{
					log.Panic("Something went wrong seeking test.txt")
				}
			if len(args) > 2 {
				idx , err :=strconv.ParseInt(args[2], 10, 64)
				if err != nil || int(idx) > len(strArr) || int(idx) <= 0{
					log.Panic("Enter a valid index to delete")
				}
				strArr = append(strArr[:idx-1], strArr[int(idx):]...)
				for i := 0; i < len(strArr); i++ {
					temp := strArr[i] 
					if i!=len(strArr)-1 {
						temp = temp + "\n"
					}
					_, err=file.WriteString(temp)
					if err!=nil{
						log.Panic("Couldn't delete the index:", idx)
					}
				}
			}	
		}else if args[1] == "insert"{
			if len(args) ==2{
				log.Panic("Enter a valid todo to insert")
			}
			if count==0{
				temp:=""
				for i := 2; i < len(args); i++ {
					temp=temp+args[i]+" "
				}
				_, err=file.WriteString(temp)
				if err!=nil{
					log.Panic("Couldn't insert the todo:", args[2])
				}
			}else{
				temp:="\n"
				for i := 2; i < len(args); i++ {
					temp=temp+args[i]+" "
				}
				_, err=file.WriteString(temp)
				if err!=nil{
					log.Panic("Couldn't insert the todo:", args[2])
				}
			}
		}
	}	
}
