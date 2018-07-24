package main

import ("fmt"
        "github.com/user/stringutil"
        "os"
        "math"
        "math/cmplx"
)

func main(){
  fmt.Printf("Hello, world.\n")
  fmt.Printf(stringutil.Reverse("\nHello Caiman"))
  var s, sep string
  for i := 1; i < len(os.Args); i++{
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Println(s)
  fmt.Printf("%.1f\n", cmplx.Exp(1i*math.Pi/2)+1)
}
