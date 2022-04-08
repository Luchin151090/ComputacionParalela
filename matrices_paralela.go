package main

import (
	"fmt"
	"math/rand"
	"time"
)
type Matriz struct{
  fila int
  columna int
}
func (m Matriz) run(){
  a:=make([][]int,m.fila)
  for i:=range a{
    a[i]=make([]int,m.columna)
  }

  b:=make([][]int,m.fila)
  for i:=range b{
    b[i]=make([]int,m.columna)
  }

  c:=make([][]int,m.fila)
  for i:=range c{
    c[i]=make([]int,m.columna)
  }

  for  i:=0;i<m.fila;i++{
    for j:=0;j<m.columna;j++{
    a[i][j]=rand.Intn(10)
    b[i][j]=rand.Intn(10)
  }
}

  if(m.fila%2==0)&&(m.columna%2==0){
    cua1:=m.fila/2
    cua2:=cua1+1
    cua3:=cua2
    for i:=0;i<m.fila;i++{
      for j:=0;j<m.columna;j++{
        c[i][j]=0
        for k:=0;k<m.columna;k++{
          if i<cua1 && j<cua1{
		  time.Sleep(551*time.Millisecond)
            c[i][j]=c[i][j]+(a[i][k]*b[k][j])
          }else if i<cua2 && j>=cua2{
		   time.Sleep(551*time.Millisecond)
            c[i][j]=c[i][j]+(a[i][k]*b[k][j])
          }else if i>=cua3 && j<cua3{
		   time.Sleep(551*time.Millisecond)
            c[i][j]=c[i][j]+(a[i][k]*b[k][j])
	  }else{
		time.Sleep(551*time.Millisecond)
            c[i][j]=c[i][j]+(a[i][k]*b[k][j])
	  }
        }
      }
    }
  }else{
    cua1:=m.fila/2
    cua2:=cua1
    cua3:=cua2
    for i:=0;i<m.fila;i++{
      for j:=0;j<m.columna;j++{
        c[i][j]=0
        for k:=0;k<m.columna;k++{
          if i<cua1 && j<cua1{
            c[i][j]=c[i][j]+(a[i][k]*b[k][j])
          }else if i<cua2 && j>=cua2{
            c[i][j]=c[i][j]+(a[i][k]*b[k][j])
          }else if i>=cua3 && j<cua3{
            c[i][j]=c[i][j]+(a[i][k]*b[k][j])
	  }else{
		time.Sleep(551*time.Millisecond)
            c[i][j]=c[i][j]+(a[i][k]*b[k][j])
	  }
        }
      }
    }    
  }

  //Mostrar
  /*
  fmt.Println("MATRIZ A")
  for i:=0;i<m.fila;i++{
    for j:=0;j<m.columna;j++{
      fmt.Print(a[i][j]," ")
    }
    fmt.Println()
  }

    fmt.Println("MATRIZ B")
  for i:=0;i<m.fila;i++{
    for j:=0;j<m.columna;j++{
      fmt.Print(b[i][j]," ")
    }
    fmt.Println()
  }

    fmt.Println("MATRIZ C")
  for i:=0;i<m.fila;i++{
    for j:=0;j<m.columna;j++{
      fmt.Print(c[i][j]," ")
    }
    fmt.Println()
  } 
*/
}

func main(){
  fmt.Println("Comenzando...")

  m1 := Matriz{1000,1000}
  m2 := Matriz{1000,1000}
  m3 := Matriz{1000,1000}
  //Temporizador General
  start:=time.Now()

  hilos:=[3]Matriz{m1,m2,m3}

  
  fmt.Println("Tiempo de las multiplicaciones...")
  start2:=time.Now()
  for i:=0;i<3;i++{
    go hilos[i].run()
  }
  
  fmt.Println(time.Since(start2))
  fmt.Println("Tiempo general: ",time.Since(start))
}
