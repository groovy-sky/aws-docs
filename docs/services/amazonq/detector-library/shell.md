![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Shell detectors(16/16)

[Avoid LS-Grep](https://docs.aws.amazon.com/amazonq/detector-library/shell/avoid-ls-grep) [Sudo Redirect Misuse](https://docs.aws.amazon.com/amazonq/detector-library/shell/sudo-redirect-misuse) [Incorrectly used escape sequences](https://docs.aws.amazon.com/amazonq/detector-library/shell/incorrectly-used-escape-sequences) [Expr Modernization](https://docs.aws.amazon.com/amazonq/detector-library/shell/expr-modernization) [Unquoted Special Parameters](https://docs.aws.amazon.com/amazonq/detector-library/shell/unquoted-special-parameters) [Single-Iteration Loop](https://docs.aws.amazon.com/amazonq/detector-library/shell/single-iteration-loop) [Prefer Find Over LS](https://docs.aws.amazon.com/amazonq/detector-library/shell/prefer-find-over-ls) [Avoid Complex Logical Expressions](https://docs.aws.amazon.com/amazonq/detector-library/shell/avoid-complex-logical-expressions) [Unquoted Array Expansion](https://docs.aws.amazon.com/amazonq/detector-library/shell/unquoted-array-expansion) [PS Grep Alternative](https://docs.aws.amazon.com/amazonq/detector-library/shell/ps-grep-alternative) [Unquoted Find Patterns](https://docs.aws.amazon.com/amazonq/detector-library/shell/unquoted-find-patterns) [Unnecessary Variable Expansion](https://docs.aws.amazon.com/amazonq/detector-library/shell/unnecessary-variable-expansion) [Use of if and then](https://docs.aws.amazon.com/amazonq/detector-library/shell/use-of-if-and-then) [Read Lines with While Loop](https://docs.aws.amazon.com/amazonq/detector-library/shell/read-lines-with-while-loop) [Incorrect Quoting in Trap Commands](https://docs.aws.amazon.com/amazonq/detector-library/shell/incorrect-quoting-in-trap-commands) [Command Substitution Syntax](https://docs.aws.amazon.com/amazonq/detector-library/shell/command-substitution-syntax)

# Shell detectors

Showing all detectors for the Shell language.

##### Browse by tags

Browse all detectors by tags.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/shell/tags)

##### Browse by severity

Browse all detectors by severity.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/shell/severity)

##### Browse by category

Browse all detectors by category.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/shell/categories)

* * *

### Browse all detectors

### [Avoid LS-Grep](https://docs.aws.amazon.com/amazonq/detector-library/shell/avoid-ls-grep)

Prefer globs or loops over ls-grep.

### [Sudo Redirect Misuse](https://docs.aws.amazon.com/amazonq/detector-library/shell/sudo-redirect-misuse)

Shell script incorrectly uses sudo with I/O redirection, which doesn't work as intended.

### [Incorrectly used escape sequences](https://docs.aws.amazon.com/amazonq/detector-library/shell/incorrectly-used-escape-sequences)

The problem involves replacing literal '\\t' with actual tab characters using the command $(printf '\\t').

### [Expr Modernization](https://docs.aws.amazon.com/amazonq/detector-library/shell/expr-modernization)

Replace 'expr' with modern shell constructs.

### [Unquoted Special Parameters](https://docs.aws.amazon.com/amazonq/detector-library/shell/unquoted-special-parameters)

Use quoted special parameters to prevent word splitting and globbing.

### [Single-Iteration Loop](https://docs.aws.amazon.com/amazonq/detector-library/shell/single-iteration-loop)

Ensure for loops iterate over multiple values.

### [Prefer Find Over LS](https://docs.aws.amazon.com/amazonq/detector-library/shell/prefer-find-over-ls)

Use 'find' for file operations instead of 'ls'.

### [Avoid Complex Logical Expressions](https://docs.aws.amazon.com/amazonq/detector-library/shell/avoid-complex-logical-expressions)

Don't use' A && B \|\| C' as a substitute for if-then-else constructs.

### [Unquoted Array Expansion](https://docs.aws.amazon.com/amazonq/detector-library/shell/unquoted-array-expansion)

Missing quotes around array expansion.

### [PS Grep Alternative](https://docs.aws.amazon.com/amazonq/detector-library/shell/ps-grep-alternative)

Use pgrep for process lookup.

### [Unquoted Find Patterns](https://docs.aws.amazon.com/amazonq/detector-library/shell/unquoted-find-patterns)

Lack of quotes around pattern arguments in 'find' command.

### [Unnecessary Variable Expansion](https://docs.aws.amazon.com/amazonq/detector-library/shell/unnecessary-variable-expansion)

Avoid unnecessary dollar signs in arithmetic expressions.

### [Use of if and then](https://docs.aws.amazon.com/amazonq/detector-library/shell/use-of-if-and-then)

The problem involves checking the exit code or output of a command in a shell script using different conditional constructs.

### [Read Lines with While Loop](https://docs.aws.amazon.com/amazonq/detector-library/shell/read-lines-with-while-loop)

Use while loop for line-by-line processing.

### [Incorrect Quoting in Trap Commands](https://docs.aws.amazon.com/amazonq/detector-library/shell/incorrect-quoting-in-trap-commands)

Using single quotes instead of double quotes in trap commands to prevent premature expansion of variables and commands.

### [Command Substitution Syntax](https://docs.aws.amazon.com/amazonq/detector-library/shell/command-substitution-syntax)

Use '$(...)' instead of backticks for command substitution.
