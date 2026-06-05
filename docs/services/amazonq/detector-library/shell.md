---
title: "Shell detectors"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Shell detectors(16/16)

[Avoid LS-Grep](shell/avoid-ls-grep.md) [Sudo Redirect Misuse](shell/sudo-redirect-misuse.md) [Incorrectly used escape sequences](shell/incorrectly-used-escape-sequences.md) [Expr Modernization](shell/expr-modernization.md) [Unquoted Special Parameters](shell/unquoted-special-parameters.md) [Single-Iteration Loop](shell/single-iteration-loop.md) [Prefer Find Over LS](shell/prefer-find-over-ls.md) [Avoid Complex Logical Expressions](shell/avoid-complex-logical-expressions.md) [Unquoted Array Expansion](shell/unquoted-array-expansion.md) [PS Grep Alternative](shell/ps-grep-alternative.md) [Unquoted Find Patterns](shell/unquoted-find-patterns.md) [Unnecessary Variable Expansion](shell/unnecessary-variable-expansion.md) [Use of if and then](shell/use-of-if-and-then.md) [Read Lines with While Loop](shell/read-lines-with-while-loop.md) [Incorrect Quoting in Trap Commands](shell/incorrect-quoting-in-trap-commands.md) [Command Substitution Syntax](shell/command-substitution-syntax.md)

# Shell detectors

Showing all detectors for the Shell language.

##### Browse by tags

Browse all detectors by tags.

[Click here→](shell/tags.md)

##### Browse by severity

Browse all detectors by severity.

[Click here→](shell/severity.md)

##### Browse by category

Browse all detectors by category.

[Click here→](shell/categories.md)

* * *

### Browse all detectors

### [Avoid LS-Grep](shell/avoid-ls-grep.md)

Prefer globs or loops over ls-grep.

### [Sudo Redirect Misuse](shell/sudo-redirect-misuse.md)

Shell script incorrectly uses sudo with I/O redirection, which doesn't work as intended.

### [Incorrectly used escape sequences](shell/incorrectly-used-escape-sequences.md)

The problem involves replacing literal '\\t' with actual tab characters using the command $(printf '\\t').

### [Expr Modernization](shell/expr-modernization.md)

Replace 'expr' with modern shell constructs.

### [Unquoted Special Parameters](shell/unquoted-special-parameters.md)

Use quoted special parameters to prevent word splitting and globbing.

### [Single-Iteration Loop](shell/single-iteration-loop.md)

Ensure for loops iterate over multiple values.

### [Prefer Find Over LS](shell/prefer-find-over-ls.md)

Use 'find' for file operations instead of 'ls'.

### [Avoid Complex Logical Expressions](shell/avoid-complex-logical-expressions.md)

Don't use' A && B \|\| C' as a substitute for if-then-else constructs.

### [Unquoted Array Expansion](shell/unquoted-array-expansion.md)

Missing quotes around array expansion.

### [PS Grep Alternative](shell/ps-grep-alternative.md)

Use pgrep for process lookup.

### [Unquoted Find Patterns](shell/unquoted-find-patterns.md)

Lack of quotes around pattern arguments in 'find' command.

### [Unnecessary Variable Expansion](shell/unnecessary-variable-expansion.md)

Avoid unnecessary dollar signs in arithmetic expressions.

### [Use of if and then](shell/use-of-if-and-then.md)

The problem involves checking the exit code or output of a command in a shell script using different conditional constructs.

### [Read Lines with While Loop](shell/read-lines-with-while-loop.md)

Use while loop for line-by-line processing.

### [Incorrect Quoting in Trap Commands](shell/incorrect-quoting-in-trap-commands.md)

Using single quotes instead of double quotes in trap commands to prevent premature expansion of variables and commands.

### [Command Substitution Syntax](shell/command-substitution-syntax.md)

Use '$(...)' instead of backticks for command substitution.

All content copied from https://docs.aws.amazon.com/.
