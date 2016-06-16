package main

import (
    _ "github.com/ariefdarmawan/slzr"
    "github.com/eaciit/toolkit"
    "flag"
)

var (
    log *toolkit.LogEngine
    dsfrom = flag.String("from","","Data source origin")
    dsto = flag.String("to","","Data source destination")
    mapname = flag.String("map","","Name of data transfer mapping")
)

func main(){
    flag.Parse()

    log, _ = toolkit.NewLog(true,false,"","","")
    defer log.Close()

    if *dsfrom=="" {
        log.Error("From - Data source origin can't be empty")
        return
    }

    if *dsto=="" {
        log.Error("To - Data source destination can't be empty")
        return
    }
}

