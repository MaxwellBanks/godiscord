package godiscordutil

import "testing"

var evenTable = [][]string{
	{"xxxx", "yyyy", "zzzz"},
	{"aa", "bb", "cc"},
}

var oddTable = [][]string{
	{"xxx", "yyy", "zzz"},
	{"a", "b", "c"},
}

var mixedTable = [][]string{
	{"xxx", "yyyy", "z"},
	{"aa", "bbbbb", "ccc"},
}

var evenSolution = "|xxxx|yyyy|zzzz|\n" +
	"----------------\n" +
	"| aa | bb | cc |\n"

var oddSolution = "|xxx|yyy|zzz|\n" +
	"-------------\n" +
	"| a | b | c |\n"

var mixedSolution = "|xxx|yyyy | z |\n" +
	"---------------\n" +
	"|aa |bbbbb|ccc|\n"

func TestGenTable(t *testing.T) {
	if GenTable(evenTable) != evenSolution {
		t.Errorf(
			"Generating even spaced table failed. Expected\n"+
				"%s\n, got \n%s\n", evenSolution, GenTable(evenTable),
		)
	}
	if GenTable(oddTable) != oddSolution {
		t.Errorf(
			"Generating odd spaced table failed. Expected\n"+
				"%s\n, got \n%s\n", oddSolution, GenTable(oddTable),
		)
	}
	if GenTable(mixedTable) != mixedSolution {
		t.Errorf(
			"Generating mixed spaced table failed. Expected\n"+
				"%s\n, got \n%s\n", mixedSolution, GenTable(mixedTable),
		)
	}
}
