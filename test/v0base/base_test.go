package v0base

import (
    "testing"
    "errors"
)

func check(t *testing.T, e error, pre string, stop bool){
    if pre!=""{
        pre=pre + ": "
    }
    
    if e!=nil {
        if stop {
            t.Fatalf("%s%s",pre,e.Error())
        } else {
            t.Errorf("%s%s",pre,e.Error())
        }
    }
}

func TestBase(t *testing.T){
    check(t, errors.New("test fatal"),"Base",true)
}