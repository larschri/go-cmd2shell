package cmd2shell

import (
	"os/exec"
	"strings"
	"testing"
)

func TestPretty(t *testing.T) {
	c := exec.Command("no cmd", "1 2\t3\u2603", "456")
	c.Dir = "/home/foo"
	c.Env = []string{"x=y", "1 2 3"}

	ol := Pretty(c)

	expect := []string{"env -i -C /home/foo \\",
		"  x=y \\",
		"  '1 2 3=' \\",
		"  'no cmd' \\",
		"  '1 2\t3☃' \\",
		"  456"}

	if ol != strings.Join(expect, "\n") {
		t.Error(ol)
	}
}

func TestOneLiner(t *testing.T) {
	c := exec.Command("no cmd", "1 2\t3\u2603", "456")
	c.Dir = "/home/foo"
	c.Env = []string{"x=y", "1 2 3"}

	ol := OneLiner(c)

	if ol != "env -i -C /home/foo x=y '1 2 3=' 'no cmd' '1 2\t3☃' 456" {
		t.Error(ol)
	}
}
