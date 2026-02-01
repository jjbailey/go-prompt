## go-prompt

This repository is a fork of github.com/c-bata/go-prompt, providing minor changes
to go-prompt's behavior.

- In the c-bata implementation, to exit a prompt without taking any action,
users must press [ENTER] with no other input at the prompt, in other words, enter
a blank line. With this change to prompt.go, the flag prompt.BackedOut enables
users to press [BACKSPACE] on an empty line and achieve the same result. This UX
behavior is similar to other Linux programs, such as Vim, which enable users to
backspace out of prompts.

- In the c-bata implementation, renderPrefix() has two issues: it does not allow
escape sequences in the prefix, and it sets the prefix color, potentially to a
faint or invisible color. The flag prompt.RawPrefix, set to true, allows escape
sequences in the prompt and avoids setting the prefix color.

- Add OptionDisableTitle to disable setting the terminal title. Unset, go-prompt
sets the title to an empty string if it is not explicitly set with OptionTitle.

### Other Changes

- Upgrade Golang from version 1.14 to 1.24.12
- Cline found several minor issues
