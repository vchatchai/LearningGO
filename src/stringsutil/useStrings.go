package stringsutil

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func UseStrings() {
	f("Tesst")
	upper := s.ToUpper("Hello there!")
	f("s.ToUpper(\"Hello there!\"):%s\n", upper)
	f("s.ToLower(\"Hello THERE\"): %s\n", s.ToLower("Hello THERE"))
	f("Title(\"tHis wiLL be A title!\"): %s\ns", s.Title("tHis wiLL be A title!"))
	f("EqualFold(\"Mihalis\", \"MIHAlis\"): %v\n", s.EqualFold("Mihalis", "MIHAlis"))
	f("EqualFold(\"Mihalis\", \"MIHAli\"): %v\n", s.EqualFold("Mihalis", "MIHAli"))
	f("HasPrefix(\"Mihalis\", \"Mi\"): %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("HasPrefix(\"Mihalis\", \"mi\"): %v\n", s.HasPrefix("Hihalis", "mi"))
	f("HasSuffix(\"Mihalis\", \"is\"): %v\n", s.HasSuffix("Mihalis", "is"))
	f("HasSuffix(\"Mihalis\", \"IS\"): %v\n", s.HasSuffix("Mihalis", "IS"))

	f("Index(\"Mihalis\", \"ha\"): %v\n", s.Index("Mihalis", "ha"))
	f("Index(\"Mihalis\", \"Ha\"): %v\n", s.Index("Mihalis", "Ha"))
	f("Count(\"Mihalis\", \"i\"): %v\n", s.Count("Mihalis", "i"))
	f("Count(\"Mihalis\", \"I\"): %v\n", s.Count("Mihalis", "I"))
	f("Repeat(\"ab\", 5): %s\n", s.Repeat("ab", 5))
	// f("Repeat ab*5 : %s\n", "ab"*5)

	f("TrimSpace(\" \\tThis is a line. \\n\"):'%s'\n", s.TrimSpace(" \tThis is a line. \n"))
	f("TrimLeft(\" \\t This is a\\t line. \\n\", \"\\n\\t\") '%s'\n", s.TrimLeft(" \t This is a\t line. \n", "\n\t"))
	f("TrimRight(\" \\t This is a\\t line. \\n\", \"\\n\\t\") '%s'\n", s.TrimRight(" \t This is a\t line. \n", " \n\t"))
	f("Compare(\"Mihalis\", \"MIHALIS\"): %v\n", s.Compare("Mihalis", "MIHALIS"))
	f("Compare(\"Mihalis\", \"Mihalis\"): %v\n", s.Compare("Mihalis", "Mihalis"))
	f("Compare(\"MIHALIS\", \"MIHalis\"): %v\n", s.Compare("MIHALIS", "MIHalis"))
	f("Fields(\"This is a string!\"): %v\n", s.Fields("This is a string!"))
	f("Fields(\"Thisis\\na\\tstring!\") %v\n", s.Fields("Thisis\na\tstring!"))
	f("Split(\"abcd efg\", \"\"):%s\n", s.Split("abcd efg", ""))
	f("Replace(\"abcd efg\", \"\", \"-\", -1):%s\n", s.Replace("abcd efg", "", "-", -1))
	f("Replace(\"abcd efg\", \"\", \"_\", 4):%s\n", s.Replace("abcd efg", "", "_", 4))
	f("Replace(\"abcd efg\", \"\", \"_\", 2): %s\n", s.Replace("abcd efg", "", "_", 2))
	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join(lines, \"+++\"): %s\n", s.Join(lines, "+++"))

	f("SplitAfter(\"123+++432++\", \"++\"): %s\n", s.SplitAfter("123+++432++", "++"))

	trimFuction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc(\"123 abc ABC  .\", trimFuction): %s\n", s.TrimFunc("123 abc ABC \t .", trimFuction))

}
