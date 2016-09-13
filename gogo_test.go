package gogo

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func calcPart1a() error {
	fmt.Println("calcPart1a....")
	time.Sleep(1 * time.Second)
	return nil
}

func calcPart1b() error {
	fmt.Println("calcPart1b...")
	time.Sleep(1 * time.Second)
	return nil
}

func calcPart2() error {
	fmt.Println("calcPart2...")
	time.Sleep(1 * time.Second)
	return errors.New("some error just occured")
}

func calcPart3() error {
	fmt.Println("calcPart3...")
	return nil
}

func TestGOGO(t *testing.T) {
	err := Run(Fns{calcPart1a, calcPart1b}, Fns{calcPart2}, Fns{calcPart3})
	if err == nil {
		t.Fatal("should fail on calcpart2")
	}
}

func calcPanic() error {
	panic("ppp")
	return nil
}

func TestError(t *testing.T) {
	err := Run(Fns{calcPart1a, calcPanic, calcPart1b, calcPanic, calcPanic, calcPanic, calcPanic, calcPanic, calcPanic, calcPanic})
	if err == nil {
		t.Fatal("should have error happend!")
	}
	fmt.Println(err)
}

func funcFail() error {
	time.Sleep(100 * time.Millisecond)
	return errors.New("just fail")
}

func TestFail(t *testing.T) {
	err := Run(Fns{calcPart1a, funcFail, calcPart1b})
	if err == nil {
		t.Fatal("should have error happend!")
	}
	fmt.Println(err)
}
