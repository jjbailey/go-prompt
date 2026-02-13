## go-prompt

This repository is a fork of github.com/c-bata/go-prompt, providing minor changes
to go-prompt's behavior.

- In the c-bata implementation, to exit a prompt without taking any action,
users must press [ENTER] with no other input at the prompt, in other words, enter
a blank line. The new variable prompt.BackedOut is a global boolean flag that
communicates to the caller how the prompt was exited. Specifically, it is set to
true if and only if the user presses [BACKSPACE] and the input buffer is currently
empty. When these conditions are met, the prompt loop terminates immediately,
allowing the calling code to detect this specific exit condition and distinguish
it from pressing [ENTER] on an empty line. This UX behavior is similar to other
Linux programs, such as Vim, which enable users to backspace out of prompts.

- In the c-bata implementation, renderPrefix() has two issues: it does not allow
escape sequences in the prefix, and it sets the prefix color, potentially to a
faint or invisible color. The flag prompt.RawPrefix, set to true, allows escape
sequences in the prompt and avoids setting the prefix color.

- Add OptionDisableTitle to disable setting the terminal title. Unset, go-prompt
sets the title to an empty string if it is not explicitly set with OptionTitle.

### Other Changes

- Upgrade Golang from version 1.14 to 1.25.7

- Cline found several minor issues
