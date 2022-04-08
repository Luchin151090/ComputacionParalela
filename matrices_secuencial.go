package main

import (
	"fmt"
	"math/rand"
	"time"
)
func multiplicar(fila int,columna int){
  a:=make([][]int,fila)
  for i:=range a{
    a[i]=make([]int,columna)
  }

  b:=make([][]int,fila)
  for i:=range b{
    b[i]=make([]int,columna)
  }

  c:=make([][]int,fila)
  for i:=range c{
    c[i]=make([]int,columna)
  }

  for  i:=0;i<fila;i++{
    for j:=0;j<columna;j++{
    a[i][j]=rand.Intn(10)
    b[i][j]=rand.Intn(10)
  }
}

  if(fila%2==0)&&(columna%2==0){
    cua1:=fila/2
    cua2:=cua1+1
    cua3:=cua2
    for i:=0;i<fila;i++{
      for j:=0;j<columna;j++{
        c[i][j]=0
        for k:=0;k<columna;k++{
          if i<cua1 && j<cua1{
            c[i][j]=c[i][j]+(a[i][k]*b[k][j])
          }else if i<cua2 && j>=cua2{
            c[i][j]=c[i][j]+(a[i][k]*b[k][j])
          }else if i>=cua3 && j<cua3{
            c[i][j]=c[i][j]+(a[i][k]*b[k][j])
	  }else{
		    c[i][j]=c[i][j]+(a[i][k]*b[k][j])
	  }
        }
      }
    }
  }else{
    cua1:=fila/2
    cua2:=cua1
    cua3:=cua2
    for i:=0;i<fila;i++{
      for j:=0;j<columna;j++{
        c[i][j]=0
        for k:=0;k<columna;k++{
          if i<cua1 && j<cua1{
            c[i][j]=c[i][j]+(a[i][k]*b[k][j])
          }else if i<cua2 && j>=cua2{
            c[i][j]=c[i][j]+(a[i][k]*b[k][j])
          }else if i>=cua3 && j<cua3{
            c[i][j]=c[i][j]+(a[i][k]*b[k][j])
	  }else{
		 c[i][j]=c[i][j]+(a[i][k]*b[k][j])
	  }
        }
      }
    }    
  }

  //Mostrar
  fmt.Println("MATRIZ A")
  for i:=0;i<fila;i++{
    for j:=0;j<columna;j++{
      fmt.Print(a[i][j]," ")
    }
    fmt.Println()
  }

    fmt.Println("MATRIZ B")
  for i:=0;i<fila;i++{
    for j:=0;j<columna;j++{
      fmt.Print(b[i][j]," ")
    }
    fmt.Println()
  }

    fmt.Println("MATRIZ C")
  for i:=0;i<fila;i++{
    for j:=0;j<columna;j++{
      fmt.Print(c[i][j]," ")
    }
    fmt.Println()
  } 
}

func main(){
  fmt.Println("Comenzando...")
  start:=time.Now()

  fmt.Println("Tiempo de las multiplicaciones...")
  start2:=time.Now()
  multiplicar(3,3)
  fmt.Println(time.Since(start2))

  fmt.Println("Tiempo general: ",time.Since(start))
}
