package main

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func convertWordsToDigits(str string) string {
    tokens:=[]string{"zero","one","two","three","four","five","six","seven","eight","nine"}
    representations:=[]string{"0","1","2","3","4","5","6","7","8","9"}

    for start:=0; start<len(str); start+=1{
        for i,token := range tokens{
            if len(str) < start+len(token){
                continue
            }

            slice := str[start:start+len(token)]

            if slice==token{
                left:=[]rune(str[:start])
                right:=str[start+len(token):]
                str = string(append(left,[]rune(representations[i])...))
                str = str + right
                
                break //there is no word greater than five so the window is not greater than 5 and two digits can not fit in a single window
            }
        }
    }


    return str

}

func convertWordsToDigitsReverse(str string) string {
    tokens:=[]string{"zero","one","two","three","four","five","six","seven","eight","nine"}
    representations:=[]string{"0","1","2","3","4","5","6","7","8","9"}


    ogLength:=len(str)

    for start:=ogLength-1; start>=0; start-=1{
        for i,token := range tokens{
            if len(str) < start+len(token){
                continue
            }

            slice := str[start:start+len(token)]

            if slice==token{
                left:=[]rune(str[:start])
                right:=str[start+len(token):]
                str = string(append(left,[]rune(representations[i])...))
                str = str + right
                break //there is no word greater than five so the window is not greater than 5 and two digits can not fit in a single window
            }
        }
    }


    return str

}

func getLeftDigit(str string) int{
    digit:=-1
    for i:=0; i< len(str); i++{
        char:=str[i]

        if (unicode.IsDigit(rune(char))){
            digit,_= strconv.Atoi(string(char))
            break

        }

    }

    return digit
}

func getRightDigit(str string) int{
    digit:=-1
    for i:=len(str)-1; i> -1; i--{
        char:=str[i]

        if (unicode.IsDigit(rune(char))){
            digit,_= strconv.Atoi(string(char))
            break

        }

    }

    return digit
}


func getDigits(str string) int{
    start:=0
    end:=len(str) - 1
    first:=-1
    last:=-1


    for start<=end{
        char:=str[start]
        if (unicode.IsDigit(rune(char))){
            first,_= strconv.Atoi(string(char))
        }
        char2:=str[end]
        if (unicode.IsDigit(rune(char2))){
            last,_= strconv.Atoi(string(char2))
        }
        if first==-1{
            start+=1
        }
        if last ==-1{
            end-=1
        }
        if first != -1 && last != -1 {
            break
        }

    }

    value:=0
    if last ==-1 && first != -1 {
        value=first*10 + first
    }else if first == -1 && last != -1{
        value=last * 10 + last
    }else if first ==-1{
        value=last * 10 + last
    }else if last == -1{
        value=first*10 + first
    }else{

        value=first*10 + last

    }
    if value<0{
        return 0 
    }

    return value
}

func getDigitsP2(str string) int{


    first:=getLeftDigit(convertWordsToDigits(str))
    last:=getRightDigit(convertWordsToDigitsReverse(str))

    value:=0
    if last ==-1 && first != -1 {
        value=first*10 + first
    }else if first == -1 && last != -1{
        value=last * 10 + last
    }else if first ==-1{
        value=last * 10 + last
    }else if last == -1{
        value=first*10 + first
    }else{

        value=first*10 + last

    }
    if value<0{
        return 0 
    }

    return value
}

//DRIVER funcitons ---------------------------

func dayOne() int{
    dat, err := os.ReadFile("dayOne.txt")
    if err != nil{
        panic("Couldn't read file")
    }

    values:=strings.Split(string(dat),"\n")

    sum:=0
    for _,val:= range values{
        sum+=getDigits(val)
    }

    return sum


}

func dayOneP2() int{
    dat, err := os.ReadFile("dayOneP2.txt")
    if err != nil{
        panic("Couldn't read file")
    }

    values:=strings.Split(string(dat),"\n")

    sum:=0
    for _,val:= range values{
        sum+=getDigitsP2(val)
    }

    return sum


}
