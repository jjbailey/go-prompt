# go-prompt

This repository is a fork of github.com/c-bata/go-prompt, written to add
a minor change to the prompt.go file.

In the c-bata implementation, to exit a prompt without taking any action, a
user must press [ENTER] with no other input at the prompt, in other words,
enter a blank line. This change to prompt.go enables a user to press [BACKSPACE]
on an empty line and achieve the same result. This UX behavior is similar to
other Linux programs, such as Vim, which enable users to backspace out of prompts.

### Other Changes

 - Upgrade Golang from version 1.14 to 1.24.7
 - In history.go, change copy loop to copy()
 - In output_vt100.go, attempt to not clobber an xterm title with an empty string
