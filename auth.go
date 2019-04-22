package main
import (
"fmt"
"net/http"
)

func Identify(r *http.Request) error {
	fmt.Println("plugin loading ............")
	if len(r.Header) > 0 {
      for k,v := range r.Header {
         fmt.Printf("%s=%s\n", k, v[0])
      }
   }
	return nil
}
