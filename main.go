package main

import (
	"fmt"
)

type Myerr struct {
	ParentErr *Myerr
	ErrMsg    string
}

func New(msg string, parentErr interface{}) error {
	var p *Myerr
	p = nil
	if parentErr != nil {
		p, _ = parentErr.(*Myerr)

	}
	return &Myerr{
		ParentErr: p,
		ErrMsg:    msg,
	}
}

func (err *Myerr) Error() string {
	if err.ParentErr == nil {
		return err.ErrMsg
	}
	return err.ParentErr.Error() + "\n==========\n" + err.ErrMsg
}

func test1() error {
	return New("test1 error", nil)
}

func test2() error {
	err := test1()
	if err != nil {
		return New("test2 error", err)
	}
	return nil
}

func main() {
	fmt.Println("hello world")
	err := test2()
	fmt.Println(err.Error())
}
