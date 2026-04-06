![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Shell detectors(16/16)

[Avoid LS-Grep](avoid-ls-grep.md) [Sudo Redirect Misuse](sudo-redirect-misuse.md) [Incorrectly used escape sequences](incorrectly-used-escape-sequences.md) [Expr Modernization](expr-modernization.md) [Unquoted Special Parameters](https://docs.aws.amazon.com/amazonq/detector-library/shell/unquoted-special-parameters) [Single-Iteration Loop](https://docs.aws.amazon.com/amazonq/detector-library/shell/single-iteration-loop) [Prefer Find Over LS](https://docs.aws.amazon.com/amazonq/detector-library/shell/prefer-find-over-ls) [Avoid Complex Logical Expressions](https://docs.aws.amazon.com/amazonq/detector-library/shell/avoid-complex-logical-expressions) [Unquoted Array Expansion](https://docs.aws.amazon.com/amazonq/detector-library/shell/unquoted-array-expansion) [PS Grep Alternative](https://docs.aws.amazon.com/amazonq/detector-library/shell/ps-grep-alternative) [Unquoted Find Patterns](https://docs.aws.amazon.com/amazonq/detector-library/shell/unquoted-find-patterns) [Unnecessary Variable Expansion](https://docs.aws.amazon.com/amazonq/detector-library/shell/unnecessary-variable-expansion) [Use of if and then](https://docs.aws.amazon.com/amazonq/detector-library/shell/use-of-if-and-then) [Read Lines with While Loop](https://docs.aws.amazon.com/amazonq/detector-library/shell/read-lines-with-while-loop) [Incorrect Quoting in Trap Commands](https://docs.aws.amazon.com/amazonq/detector-library/shell/incorrect-quoting-in-trap-commands) [Command Substitution Syntax](https://docs.aws.amazon.com/amazonq/detector-library/shell/command-substitution-syntax)

# Unquoted Special Parameters [Medium](https://docs.aws.amazon.com/amazonq/detector-library/shell/severity/medium)

Unquoted special parameters are subject to word splitting and globbing, potentially altering the intended arguments. To avoid this, always use '$@' with quotes to preserve the original argument structure.

**Detector ID**

shell/unquoted-special-parameters@v1.0

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
2# Noncompliant: Word splitting occurs, breaking arguments with spaces.
3copy_files() {
4    cp $* /backup/
5}

```

#### Compliant example

```bash
1
2# Compliant: Preserves arguments with spaces.
3copy_files() {
4    cp "$@" /backup/
5}

```
