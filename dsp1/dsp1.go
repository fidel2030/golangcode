package main

import (
  "fmt"
  "math"
  "math/cmplx"

)

func main(){
  const N = 7
  //var x [7] int = [7]int{2, 1, -2, 3, -1, -1, 1}
  x := [7]int{2, 1, -2, 3, -1, -1, 1}
  var a [7] complex128
  var b complex128
  for k := 0; k < N; k++{

    for n := 0; n < N; n++{
      //a[k] = complex(float64(k)*1.0,0.0)
      b = 1i*2*math.Pi//*complex(float64(k),0)*complex(float64(n),0)/N
      a[k] = a[k] // + 1.0 / N * x[n] * cmplx.Exp(b)
      fmt.Printf("%d\t%d", k, n)
      fmt.Printf("\t%.1f\n", cmplx.Exp(math.Pi*1i))
    }
  }

}
