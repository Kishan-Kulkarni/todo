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
	if count!=0{
		for i := 0; i < len(strArr); i++ {
			curr := strArr[i]
			last := curr[len(curr)-1:]
			if last == "Y"{
				done[i]=true
			}
		}
	}
	if len(args) ==1{
		if count!=0{
			PrintList(strArr, done)
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
				PrintList(strArr, done)
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
				temp = temp + "N"
				_, err=file.WriteString(temp)
				if err!=nil{
					log.Panic("Couldn't insert the todo:", args[2])
				}
				strArr = append(strArr,temp)
				done=append(done, false)
				PrintList(strArr[1:2], done)
			}else{
				temp:="\n"
				strTemp :=""
				for i := 2; i < len(args); i++ {
					temp=temp+args[i]+" "
					strTemp=strTemp+args[i]+" "
				}
				temp = temp + "N"
				strTemp = strTemp + "N"
				_, err=file.WriteString(temp)
				if err!=nil{
					log.Panic("Couldn't insert the todo:", args[2])
				}
				strArr = append(strArr,strTemp)
				done=append(done, false)
				PrintList(strArr,done)
			}
		}else if args[1]== "done"{
			err =file.Truncate(0)
			if err!=nil{
				log.Panic("Something went wrong erasing test.txt")
			}
			_, err = file.Seek(0,0)	
			if err!=nil{
				log.Panic("Something went wrong seeking test.txt")
			}
			if len(args)==2{
				for i := 0; i < len(strArr); i++ {
					curr := strArr[i]
					curr=curr[:len(curr)-1]
					curr=curr+"Y"
					strArr[i] = curr
					done[i]=true					
				}
			}else{
				idx , err :=strconv.ParseInt(args[2], 10, 64)
				if err != nil || int(idx) > len(done) || int(idx) <= 0{
					log.Panic("Enter a valid index to delete")
				}
				curr:=strArr[idx-1]
				curr=curr[:len(curr)-1]
				curr=curr+"Y"
				strArr[idx-1] = curr
				done[idx-1]=true
			}
			for i := 0; i < len(strArr); i++ {
				temp := strArr[i] 
				if i!=len(strArr)-1 {
					temp = temp + "\n"
				}
				_, err=file.WriteString(temp)
				if err!=nil{
					log.Panic("Couldn't complete the todo")
				}
			}
			PrintList(strArr, done)
		}else if args[1]=="clear"{
			err =file.Truncate(0)
			if err!=nil{
				log.Panic("Something went wrong erasing test.txt")
			}
			_, err = file.Seek(0,0)	
			if err!=nil{
				log.Panic("Something went wrong seeking test.txt")
			}
			newStr:=make([]string,0)
			newDone:=make([]bool,0)
			for i := 0; i < len(strArr); i++ {
				if !done[i]{
					newStr=append(newStr, strArr[i])
					newDone=append(newDone, false)
				}
			}
			strArr=newStr
			done=newDone
			for i := 0; i < len(strArr); i++ {
				temp := strArr[i] 
				if i!=len(strArr)-1 {
					temp = temp + "\n"
				}
				_, err=file.WriteString(temp)
				if err!=nil{
					log.Panic("Couldn't clear the list")
				}
			}
			PrintList(strArr, done)
		}
	}	
}

func PrintList(strArr []string, done []bool) {
	if len(strArr) ==1 {
		valid:=strArr[0]
		valid = valid[0:len(valid)-1]
		temp := fmt.Sprintf("%d.	%s",1,  valid)
		if !done[0]{
			color.Red(temp)
		}else{
			color.Green(temp)
		}
		return
	}
	for i := 0; i < len(strArr); i++ {
		valid:=strArr[i]
		valid = valid[0:len(valid)-1]
		temp := fmt.Sprintf("%d.	%s",i+1,  valid)
		if !done[i]{
			color.Red(temp)
		}else{
			color.Green(temp)
		}
	}
}
