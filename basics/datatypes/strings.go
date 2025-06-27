package datatypes

import (
	"fmt"
	"go_lang/utils"
	"reflect"
	"strconv"
	"strings"
)

func String() {
	var firstname string = "Alice"
	var lastname string = "Wonderland"
	var fullname = firstname + " " + lastname
	utils.Heading("String", "")
	fmt.Println("string:", firstname)

	fmt.Println()

	utils.Heading("String Methods", "")
	utils.FormatOutput(
		"Concatenation (+)",
		fmt.Sprintf("Firstname: (%s) Lastname (%s)", firstname, lastname),
		fullname,
	)
	utils.FormatOutput("Length (len())", firstname, len(firstname))

	utils.FormatOutput(
		"Contains (strings.Contains())",
		fmt.Sprintf("String (%s) Substring (%s)", fullname, "Alice"),
		strings.Contains(fullname, "Alice"),
	)

	utils.FormatOutput(
		"ContainsAny (strings.ContainsAny())",
		fmt.Sprintf("String (%s) Substring (%s)", fullname, "lopicq"),
		strings.ContainsAny(fullname, "lopicq"),
	)
	utils.FormatOutput(
		"HasPrefix (strings.HasPrefix())",
		fmt.Sprintf("String (%s) Prefix (%s)", fullname, "Al"),
		strings.HasPrefix(fullname, "Al"),
	)
	utils.FormatOutput(
		"HasSuffix (strings.HasSuffix())",
		fmt.Sprintf("String (%s) Suffix (%s)", fullname, "land"),
		strings.HasSuffix(fullname, "land"),
	)
	utils.FormatOutput(
		"Index (strings.Index())",
		`String ("hello world world") Substring ("world")`,
		strings.Index("hello world world", "world"),
	)
	utils.FormatOutput(
		"LastIndex (strings.LastIndex())",
		`String ("hello world world") Substring ("world")`,
		strings.LastIndex("hello world world", "world"),
	)
	utils.FormatOutput(
		"IndexAny (strings.IndexAny())",
		`String ("hello world") Chars ("ord")`,
		strings.IndexAny("hello world", "ord"),
	)
	utils.FormatOutput(
		"Count (strings.Count())",
		`String ("aaaa") Substring ("a")`,
		strings.Count("aaaa", "a"),
	)
	utils.FormatOutput(
		"Split (strings.Split())",
		`String ("Orange,Mango") Sep (",")`,
		strings.Split("Orange,Mango", ","),
	)
	utils.FormatOutput(
		"SplitN (strings.SplitN())",
		`String ("a,b,c") Sep (",") N (4)`,
		strings.SplitN("a,b,c", ",", 4),
	)
	utils.FormatOutput(
		"Fields (strings.Fields())",
		`String ("Hello  World new people")`,
		strings.Fields("Hello  World new people"),
	)
	utils.FormatOutput(
		"Join (strings.Join())",
		`Slice (["pen", "pine"]) Sep (",")`,
		strings.Join([]string{"pen", "pine"}, ","),
	)
	utils.FormatOutput(
		"Replace (strings.Replace())",
		`String ("dark_strategy_user") Old ("_") New (" ") N (1)`,
		strings.Replace("dark_strategy_user", "_", " ", 1),
	)
	utils.FormatOutput(
		"ReplaceAll (strings.ReplaceAll())",
		`String ("dark_strategy_user") Old ("_") New (" ")`,
		strings.ReplaceAll("dark_strategy_user", "_", " "),
	)
	utils.FormatOutput(
		"Repeat (strings.Repeat())",
		`String ("s") Count (3)`,
		strings.Repeat("s", 3),
	)
	utils.FormatOutput(
		"Trim (strings.Trim())",
		`String ("_hello_") Cutset (",_")`,
		strings.Trim("_hello_", ",_"),
	)
	utils.FormatOutput(
		"TrimSpace (strings.TrimSpace())",
		`String (" hello\n")`,
		strings.TrimSpace(" hello\n"),
	)
	utils.FormatOutput(
		"TrimPrefix (strings.TrimPrefix())",
		`String ("__hello_") Prefix ("_")`,
		strings.TrimPrefix("__hello_", "_"),
	)
	utils.FormatOutput(
		"TrimSuffix (strings.TrimSuffix())",
		`String ("_hello__") Suffix ("_")`,
		strings.TrimSuffix("_hello__", "_"),
	)
	utils.FormatOutput(
		"TrimLeft (strings.TrimLeft())",
		`String ("__hello_") Cutset ("_")`,
		strings.TrimLeft("__hello_", "_"),
	)
	utils.FormatOutput(
		"TrimRight (strings.TrimRight())",
		`String ("_hello__") Cutset ("_")`,
		strings.TrimRight("_hello__", "_"),
	)
	utils.FormatOutput(
		"ToUpper (strings.ToUpper())",
		`String ("hello")`,
		strings.ToUpper("hello"),
	)
	utils.FormatOutput(
		"ToLower (strings.ToLower())",
		`String ("HELLO")`,
		strings.ToLower("HELLO"),
	)
	utils.FormatOutput(
		"EqualFold (strings.EqualFold())",
		`String1 ("HELLO") String2 ("hello")`,
		strings.EqualFold("HELLO", "hello"),
	)
	utils.FormatOutput(
		"strconv.Itoa()",
		`Int (123)`,
		reflect.TypeOf(strconv.Itoa(123)),
	)
	utils.FormatOutput(
		"strconv.Atoi()",
		`String ("123")`,
		func() any {
			i, err := strconv.Atoi("123")
			if err != nil {
				return fmt.Sprintf("Error: %v", err)
			}
			return i
		}(),
	)
	utils.FormatOutput(
		"strconv.ParseFloat()",
		`String ("123.456")`,
		func() any {
			f, err := strconv.ParseFloat("123.456", 64)
			if err != nil {
				return fmt.Sprintf("Error: %v", err)
			}
			return f
		}(),
	)
	utils.FormatOutput(
		"Repeat (strings.Repeat())",
		`String ("s") Count (3)`,
		strings.Repeat("sssssss", 3),
	)
}
