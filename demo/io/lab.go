package io

import "strings"

func Do()  {
	reader:=strings.NewReader("abcd")
	println(reader.Len())
}
