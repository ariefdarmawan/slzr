package slzr

import (
    "errors"
    "github.com/eaciit/toolkit"
    "strings"
)

type IDriver interface{
    Config() *toolkit.M
    SetModel(func()interface{})
    Model()interface{}
    Next(interface{}) error
    Run(toolkit.M) error
    Save(m interface{}) error
}

type BaseDriver struct{
    fnm func()interface{}
    cfg *toolkit.M
}

func (b *BaseDriver) Config() *toolkit.M{
    if b.cfg==nil {
        b.cfg = new(toolkit.M)
    }
    return b.cfg
}

func (b *BaseDriver) Next(m interface{})error{
    return errors.New("slzr.BaseDriver.Fetch: method is not yet implemented")
}

func (b *BaseDriver) SetModel(f func()interface{}){
    b.fnm=f
}

func (b *BaseDriver) Model()interface{}{
    return b.fnm()
}

func (b *BaseDriver) Save(m interface{})error{
    return errors.New("slzr.BaseDriver.Save: method is not yet implemented")
}

func (b *BaseDriver) Run(config toolkit.M) error{
    for{
        model := b.Model() 
        e := b.Next(model)
        if e!=nil {
            if !strings.Contains(strings.ToLower(e.Error()),"eof"){
                return toolkit.Errorf("slzr.Run: %s", e.Error())
            } else {
                return nil
            }
        }

        b.Save(model)
    }
}