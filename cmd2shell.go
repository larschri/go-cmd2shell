package cmd2shell

import (
	"os"
	"os/exec"
	"strings"

	shellquote "github.com/kballard/go-shellquote"
)

// OneLiner creates a shell command in one line
func OneLiner(c *exec.Cmd) string {
	return shellquote.Join(Slice(c)...)
}

// Pretty creates a shell command in multiple lines
func Pretty(c *exec.Cmd) string {
	var sb strings.Builder
	cmd := Slice(c)

	sb.WriteString(shellquote.Join(cmd[:4]...))

	for _, s := range cmd[4:] {
		sb.WriteString(" \\\n  " + shellquote.Join(s))
	}

	return sb.String()
}

// Slice returns the arguments to be passed to the shell
func Slice(c *exec.Cmd) []string {
	s := []string{"env", "-i", "-C", c.Dir}

	if s[3] == "" {
		s[3], _ = os.Getwd()
	}

	for _, e := range c.Env {
		if !strings.Contains(e, "=") {
			// 'env' requires equal sign in NAME=VALUE format
			e = e + "="
		}
		s = append(s, e)
	}

	s = append(s, c.Path)

	if len(c.Args) > 1 {
		s = append(s, c.Args[1:]...)
	}

	return s
}
