![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Shell detectors(16/16)

[Avoid LS-Grep](avoid-ls-grep.md) [Sudo Redirect Misuse](sudo-redirect-misuse.md) [Incorrectly used escape sequences](incorrectly-used-escape-sequences.md) [Expr Modernization](expr-modernization.md) [Unquoted Special Parameters](unquoted-special-parameters.md) [Single-Iteration Loop](single-iteration-loop.md) [Prefer Find Over LS](prefer-find-over-ls.md) [Avoid Complex Logical Expressions](avoid-complex-logical-expressions.md) [Unquoted Array Expansion](unquoted-array-expansion.md) [PS Grep Alternative](ps-grep-alternative.md) [Unquoted Find Patterns](unquoted-find-patterns.md) [Unnecessary Variable Expansion](https://docs.aws.amazon.com/amazonq/detector-library/shell/unnecessary-variable-expansion) [Use of if and then](https://docs.aws.amazon.com/amazonq/detector-library/shell/use-of-if-and-then) [Read Lines with While Loop](https://docs.aws.amazon.com/amazonq/detector-library/shell/read-lines-with-while-loop) [Incorrect Quoting in Trap Commands](https://docs.aws.amazon.com/amazonq/detector-library/shell/incorrect-quoting-in-trap-commands) [Command Substitution Syntax](https://docs.aws.amazon.com/amazonq/detector-library/shell/command-substitution-syntax)

# Unnecessary Variable Expansion [Low](https://docs.aws.amazon.com/amazonq/detector-library/shell/severity/low)

Dollar signs or curly braces are typically redundant when referencing variables within arithmetic expressions. Omitting them simplifies the expression and can prevent potential issues. However, there are exceptions for special variables or complex scenarios.

**Detector ID**

shell/unnecessary-variable-expansion@v1.0

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
2# Noncompliant: Unnecessary use of $ in arithmetic context.
3total=$((${count} + ${offset}))
```

#### Compliant example

```bash
1
2# Compliant: Simplified arithmetic expression without $
3total=$((count + offset))

```
