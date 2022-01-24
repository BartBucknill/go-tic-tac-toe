package game

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/kylelemons/godebug/diff"
)

type BoardTest struct {
	input  string
	expect string
	cmd    string
}

var boardTests = []BoardTest{
	{
		expect: `  1  2  3  

		A  .  .  .  
			
		B  .  .  .  
			
		C  .  .  .  
		`,
	},
	{
		cmd: `A1`,
		expect: `  1  2  3  

		A  O  .  .  
			
		B  .  .  .  
			
		C  .  .  .  
		`,
	},
}

func normaliseOutput(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), "\n", ""), "\t", "")
}

func TestBoard(t *testing.T) {
	for _, test := range boardTests {
		var out bytes.Buffer
		var in bytes.Buffer
		var err bytes.Buffer
		New(&out, &in, &err).Play()
		expectNorm := normaliseOutput(test.expect)
		actualNorm := normaliseOutput(out.String())
		if actualNorm != expectNorm {
			fmt.Println("Actual output does not match expected!")
			fmt.Println(diff.Diff(expectNorm, actualNorm))
		}
	}
}

type updateBoardTest struct {
	input       board
	expect      board
	cmd         string
	shouldError bool
}

var updateBoardTests = []updateBoardTest{
	{
		cmd: "A1",
		input: board{
			{' ', '1', '2', '3'},
			{'A', '.', '.', '.'},
			{'B', '.', '.', '.'},
			{'C', '.', '.', '.'},
		},
		expect: board{
			{' ', '1', '2', '3'},
			{'A', 'O', '.', '.'},
			{'B', '.', '.', '.'},
			{'C', '.', '.', '.'},
		},
	},
	{
		cmd: "A1",
		input: board{
			{' ', '1', '2', '3'},
			{'A', 'O', '.', '.'},
			{'B', '.', '.', '.'},
			{'C', '.', '.', '.'},
		},
		shouldError: true,
	},
}

func TestUpdateBoard(t *testing.T) {
	for _, test := range updateBoardTests {
		err := updateBoard('O', "A1", test.input)
		if test.shouldError {
			if err == nil {
				t.Fatalf("expected error but got none: %v", err)
			}
			return
		}
		if err != nil {
			t.Fatalf("updateBoard errored: %v", err)
		}
		if diff := cmp.Diff(test.expect, test.input); diff != "" {
			fmt.Println("Expected:")
			test.expect.render(os.Stdout)
			fmt.Println("Got:")
			test.input.render(os.Stdout)
			t.Errorf("Failed")
		}
	}

}
