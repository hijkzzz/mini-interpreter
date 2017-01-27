package environment

import (

	"fmt"
	"strconv"
	"time"
	"bufio"
	"os"
	"reflect"
)

type NativeEnv struct {
	*NestedEnv
}

func NewNativeEnv() *NativeEnv {
	env := NewNestedEnv()
	appendNatives(env)
	return &NativeEnv{env}
}

func appendNatives(env Environment) {
	append1(env, "read", read)
	append1(env, "print", print)
	append1(env, "size", size)
	append1(env, "atoi", atoi)
	append1(env, "itoa", itoa)
	append1(env, "timestamp", timestamp)
}

func append1(env Environment, name string, fun interface{}) {
	env.Set(name, NewNativeFunction(name, reflect.ValueOf(fun)))
}

// native methods

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
func read() string {
	scanner.Split(bufio.ScanWords)
	if scanner.Scan() {
		return scanner.Text()
	} else {
		return ""
	}
}

func print(obj interface{}) int {
	fmt.Print(obj)
	return 0
}

func size(s string) int {
	return len(s)
}

func atoi(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic("atoi error")
	}
	return x
}

func itoa(i int) string {
	x := strconv.Itoa(i)
	return x
}

func timestamp() int {
	return int(time.Now().Unix())
}
