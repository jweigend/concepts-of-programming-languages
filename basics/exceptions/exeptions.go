package main

import (
	"fmt"
)

// TryCatchBlock is a codeblock with a try and optional catch, finally clause.
type TryCatchBlock struct {
	T func()
	C func(Exception)
	F func()
}

// Exception is a Exception type.
type Exception interface{}

// Throw is a alias for panic.
func Throw(up Exception) {
	panic(up)
}

// Do does call the try function als installs a catch and finally handler.
func (tcf TryCatchBlock) Do() {
	if tcf.F != nil {
		defer tcf.F()
	}
	if tcf.C != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.C(r)
			}
		}()
	}
	tcf.T()
}

// TryCatchFinally helper for a better syntax.
func TryCatchFinally(t func(), c func(ex Exception), f func()) {
	TryCatchBlock{T: t, C: c, F: f}.Do()
}

// Test exception handling.
func main() {
	fmt.Println("Starting ...")
	TryCatchFinally(
		func() {
			fmt.Println("Trying ...")
			Throw("Some Exception") // throws an exception
		},
		func(e Exception) {
			fmt.Printf("Caught %v\n", e)
		},
		func() {
			fmt.Println("Finally...")
		})
	fmt.Println("Shutdown gracefully")
}
