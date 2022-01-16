# Go cmd2shell

Convert `exec.Cmd` to a string that can be used as a shell command.

Environment, workdir, path and arguments are passed as parameters to the `env`
command. The workdir is specified with the `-L` option which is available in
GNU/Linux, but not POSIX. The strings in `exec.Cmd.Env` should be in
`KEY=VALUE` format, otherwise an equal sign will be appended.

Shell quoting is handled by
[go-shellquote](https://github.com/kballard/go-shellquote).

## Install

    go get github.com/larschri/go-cmd2shell

## Example

```go
c := exec.Command("echo", "Hello, world")
fmt.Println(cmd2shell.OneLiner(c))
```

This prints

```shell
env -i -C /current/work/dir /bin/echo 'Hello, world'
```
