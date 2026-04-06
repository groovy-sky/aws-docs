![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Shell detectors(16/16)

[Avoid LS-Grep](https://docs.aws.amazon.com/amazonq/detector-library/shell/avoid-ls-grep) [Sudo Redirect Misuse](https://docs.aws.amazon.com/amazonq/detector-library/shell/sudo-redirect-misuse) [Incorrectly used escape sequences](https://docs.aws.amazon.com/amazonq/detector-library/shell/incorrectly-used-escape-sequences) [Expr Modernization](https://docs.aws.amazon.com/amazonq/detector-library/shell/expr-modernization) [Unquoted Special Parameters](https://docs.aws.amazon.com/amazonq/detector-library/shell/unquoted-special-parameters) [Single-Iteration Loop](https://docs.aws.amazon.com/amazonq/detector-library/shell/single-iteration-loop) [Prefer Find Over LS](https://docs.aws.amazon.com/amazonq/detector-library/shell/prefer-find-over-ls) [Avoid Complex Logical Expressions](https://docs.aws.amazon.com/amazonq/detector-library/shell/avoid-complex-logical-expressions) [Unquoted Array Expansion](https://docs.aws.amazon.com/amazonq/detector-library/shell/unquoted-array-expansion) [PS Grep Alternative](https://docs.aws.amazon.com/amazonq/detector-library/shell/ps-grep-alternative) [Unquoted Find Patterns](https://docs.aws.amazon.com/amazonq/detector-library/shell/unquoted-find-patterns) [Unnecessary Variable Expansion](https://docs.aws.amazon.com/amazonq/detector-library/shell/unnecessary-variable-expansion) [Use of if and then](https://docs.aws.amazon.com/amazonq/detector-library/shell/use-of-if-and-then) [Read Lines with While Loop](https://docs.aws.amazon.com/amazonq/detector-library/shell/read-lines-with-while-loop) [Incorrect Quoting in Trap Commands](https://docs.aws.amazon.com/amazonq/detector-library/shell/incorrect-quoting-in-trap-commands) [Command Substitution Syntax](https://docs.aws.amazon.com/amazonq/detector-library/shell/command-substitution-syntax)

# Avoid LS-Grep [Medium](https://docs.aws.amazon.com/amazonq/detector-library/shell/severity/medium)

Parsing the output of 'ls' with 'grep' is unreliable due to potential issues with filenames containing spaces, newlines, or special characters. Use globbing patterns or a for loop with conditional checks for robust file handling.

**Detector ID**

shell/avoid-ls-grep@v1.0

**Category**

[Code Quality](https://docs.aws.amazon.com/amazonq/detector-library/shell/categories/code-quality)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

-

**Tags**

-

* * *

#### Noncompliant example

```bash
1
2# Noncompliant: `ls | grep` can break with special characters in filenames.
3echo "Files containing 'config' in the current directory:"
4ls | grep config

```

#### Compliant example

```bash
1
2# Compliant: Using glob pattern matching handles special characters safely.
3echo "Files containing 'config' in the current directory:"
4for file in *config*; do
5    [ -e "$file" ] && echo "$file"
6done
```
