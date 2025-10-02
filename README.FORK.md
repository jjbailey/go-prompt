# go-prompt

This repository is a fork of github.com/c-bata/go-prompt, with the intention of adding a small change to `prompt.go`.

In the c-bata implementation, to exit a prompt without taking any action, a user must press `[ENTER]` with no other input at the prompt, in other words, enter a blank line. The changes to `prompt.go` enable a user to type `[BACKSPACE]` on an empty line and achieve the same result. This UX behavior is similar to that in other Linux programs that enable users to backspace out of prompts, for example, Vim.

