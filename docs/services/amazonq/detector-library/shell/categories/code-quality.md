![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Shell detectors(16/16)

[Avoid LS-Grep](../avoid-ls-grep.md) [Sudo Redirect Misuse](../sudo-redirect-misuse.md) [Incorrectly used escape sequences](../incorrectly-used-escape-sequences.md) [Expr Modernization](../expr-modernization.md) [Unquoted Special Parameters](../unquoted-special-parameters.md) [Single-Iteration Loop](../single-iteration-loop.md) [Prefer Find Over LS](../prefer-find-over-ls.md) [Avoid Complex Logical Expressions](../avoid-complex-logical-expressions.md) [Unquoted Array Expansion](../unquoted-array-expansion.md) [PS Grep Alternative](../ps-grep-alternative.md) [Unquoted Find Patterns](../unquoted-find-patterns.md) [Unnecessary Variable Expansion](../unnecessary-variable-expansion.md) [Use of if and then](../use-of-if-and-then.md) [Read Lines with While Loop](../read-lines-with-while-loop.md) [Incorrect Quoting in Trap Commands](../incorrect-quoting-in-trap-commands.md) [Command Substitution Syntax](../command-substitution-syntax.md)

# Code Quality detectors

### [Avoid LS-Grep](../avoid-ls-grep.md)

Prefer globs or loops over ls-grep.

### [Incorrectly used escape sequences](../incorrectly-used-escape-sequences.md)

The problem involves replacing literal '\\t' with actual tab characters using the command $(printf '\\t').

### [Expr Modernization](../expr-modernization.md)

Replace 'expr' with modern shell constructs.

### [Unquoted Special Parameters](../unquoted-special-parameters.md)

Use quoted special parameters to prevent word splitting and globbing.

### [Single-Iteration Loop](../single-iteration-loop.md)

Ensure for loops iterate over multiple values.

### [Prefer Find Over LS](../prefer-find-over-ls.md)

Use 'find' for file operations instead of 'ls'.

### [Avoid Complex Logical Expressions](../avoid-complex-logical-expressions.md)

Don't use' A && B \|\| C' as a substitute for if-then-else constructs.

### [Unquoted Array Expansion](../unquoted-array-expansion.md)

Missing quotes around array expansion.

### [PS Grep Alternative](../ps-grep-alternative.md)

Use pgrep for process lookup.

### [Unquoted Find Patterns](../unquoted-find-patterns.md)

Lack of quotes around pattern arguments in 'find' command.

### [Unnecessary Variable Expansion](../unnecessary-variable-expansion.md)

Avoid unnecessary dollar signs in arithmetic expressions.

### [Use of if and then](../use-of-if-and-then.md)

The problem involves checking the exit code or output of a command in a shell script using different conditional constructs.

### [Read Lines with While Loop](../read-lines-with-while-loop.md)

Use while loop for line-by-line processing.

### [Incorrect Quoting in Trap Commands](../incorrect-quoting-in-trap-commands.md)

Using single quotes instead of double quotes in trap commands to prevent premature expansion of variables and commands.

### [Command Substitution Syntax](../command-substitution-syntax.md)

Use '$(...)' instead of backticks for command substitution.
