package main

import (
  "github.com/ugorji/go/codec"
  "fmt"
  "io"
)


var (
        v interface{} // value to decode/encode into
        r io.Reader
        w io.Writer
        b []byte
        mh codec.MsgpackHandle
    )

func main(){
    dec := codec.NewDecoder(r, &mh)
    dec = codec.NewDecoderBytes(b, &mh)
    s := dec.Decode(&v) 
    fmt.Printf("%v", s)
    //enc := codec.NewEncoder(w, &mh)
    //enc = codec.NewEncoderBytes(&b, &mh)
    //enc.Encode(v)
}
