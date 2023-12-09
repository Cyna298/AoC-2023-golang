package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isValid(char byte) bool{
    return char!='.' && !unicode.IsDigit(rune(char))  
}

func hasAdjacentSymbol(matrix []string,row int, col int) bool{
    
    width := len(matrix)
    height := len(matrix)

    for i:=-1; i<2; i++{
        for j:=-1; j<2; j++{
            
            newRow:=int(math.Max(math.Min(float64(row+i),float64(width-2)),0))
            newCol:=int(math.Max(math.Min(float64(col+j),float64(height-2)),0))
            if isValid(matrix[newRow][newCol]){
                return true
            }


        }
    }
    return false 
}

func Day3(){
    
      
    data, err := os.ReadFile("day3Input.txt")
    if err != nil{
        panic("Couldn't read file")
    }

    values:=strings.Split(string(data),"\n")
    sum:=0
    for j,line := range values{
        left := -1
        right := -1
        i := 0 
        for i <= len(line){
            char:=byte('.')
            if i<len(line){
                char = line[i]

            }
            if unicode.IsDigit(rune(char)){
                if left == -1{
                    left = i
                }

            }else if left != -1{
                right = i
            }

            if left != -1 && right != -1{

                digits,err:= strconv.Atoi(line[left:right])
                if err != nil{
                    panic(err.Error())
                }

                

                //check directions
                for k:=left; k<right;k++{
                    if hasAdjacentSymbol(values,j,k){
                        sum+=digits
                        break


                    }
                }


                if err != nil{
                    panic(err.Error())
                }

                left = -1
                right = -1
            }

            i+=1
        }

    }
    fmt.Println("SUM",sum)

}

type Index struct   {
    row int
    col int
}
type Res struct {
    value int
    left int
    length int
}

func Day3P2(){
    data, err := os.ReadFile("day3Input.txt")
    if err != nil{
        panic("Couldn't read file")
    }

    values:=strings.Split(string(data),"\n")
    sum:=0
    numMap:=make(map[Index]Res)
    for j,line := range values{
        left := -1
        right := -1
        i := 0 
        for i <= len(line){
            char:=byte('.')
            if i<len(line){
                char = line[i]

            }
            if unicode.IsDigit(rune(char)){
                if left == -1{
                    left = i
                }

            }else if left != -1{
                right = i
            }

            if left != -1 && right != -1{

                digits,err:= strconv.Atoi(line[left:right])
                if err != nil{
                    panic(err.Error())
                }


                

                //check directions
                for k:=left; k<right;k++{
                    numMap[Index{j,k}] = Res{digits,left,right-left}
                    // if hasAdjacentSymbol(values,j,k){
                    //     sum+=digits
                    //     break
                    //
                    //
                    // }
                }


                if err != nil{
                    panic(err.Error())
                }

                left = -1
                right = -1
            }


            i+=1
        }

    }

    for i,line := range values{
        for j,char := range line{
            if char=='*'{

                width := len(values)
                height := len(values)
                numbers := []int{}

                for x:=-1; x<2; x++{
                    for y:=-1; y<2; y++{

                        newRow:=int(math.Max(math.Min(float64(i+x),float64(width-2)),0))
                        newCol:=int(math.Max(math.Min(float64(j+y),float64(height-2)),0))
                        val,ok := numMap[Index{newRow,newCol}]
                        if ok{
                            numbers = append(numbers, val.value)
                            for z:=val.left; z<val.left+val.length;z++{
                                delete(numMap,Index{newRow,z})
                            }
                        }


                    }
                }
                if len(numbers)==2{
                    sum+= numbers[0]*numbers[1]
                }
                
            }
        }
    }

    fmt.Println("SUM",sum)

}
