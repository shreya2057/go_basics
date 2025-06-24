package datatypes

import (
	"fmt"
	"go_lang/utils"
	"strings"
)

func String() {
	var firstname string = "Alice"
	var lastname string = "Wonderland"
	var fullname = firstname + " " + lastname
	utils.Heading("String", "")
	fmt.Println("string:", firstname)
	fmt.Println("string concatenation:", firstname+" "+lastname)
	fmt.Println("string length:", len(firstname), len(lastname))
	fmt.Println("string Contains:", strings.Contains(fullname, "Alice"))
	fmt.Println("string ContainsAny:", strings.ContainsAny(fullname, "lopicq"))
	fmt.Println("string HasPrefix:", strings.HasPrefix(fullname, "Al"))
	fmt.Println("string HasSuffix:", strings.HasSuffix(fullname, "land"))
	fmt.Println("string Index:", strings.Index("hello world world", "world"))
	fmt.Println("string LastIndex:", strings.LastIndex("hello world world", "world"))
	fmt.Println("string IndexAny:", strings.IndexAny("hello world", "ord"))
	fmt.Println("string Count:", strings.Count("aaaa", "a"))
	fmt.Println("string Split:", strings.Split("Orange,Mango", ","))
	fmt.Println("string SplitN:", strings.SplitN("a,b,c", ",", 4))
	fmt.Println("string Fields:", strings.Fields("Hello  World new people"))
	fmt.Println("string Join:", strings.Join([]string{"pen", "pine"}, ","))
	fmt.Println("string Replace:", strings.Replace("dark_strategy_user", "_", " ", 1))
	fmt.Println("string ReplaceAll:", strings.ReplaceAll("dark_strategy_user", "_", " "))
	fmt.Println("string Trim:", strings.Trim("_hello_", ",_"))
	fmt.Println("string TrimSpace:", strings.TrimSpace(" hello\n"))
	fmt.Println("string TrimPrefix:", strings.TrimPrefix("__hello_", "_"))
	fmt.Println("string TrimSuffix:", strings.TrimSuffix("_hello__", "_"))
	fmt.Println("string TrimLeft:", strings.TrimLeft("__hello_", "_"))
	fmt.Println("string TrimRight:", strings.TrimRight("_hello__", "_"))
	fmt.Println("string ToUpper:", strings.ToUpper("hello"))
	fmt.Println("string ToLower:", strings.ToLower("HELLO"))
	fmt.Println("string EqualFold:", strings.EqualFold("HELLO", "hello"))
}
