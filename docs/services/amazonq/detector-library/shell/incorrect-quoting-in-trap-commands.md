---
title: "Incorrect Quoting in Trap Commands Medium"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Shell detectors(16/16)

[Avoid LS-Grep](avoid-ls-grep.md) [Sudo Redirect Misuse](sudo-redirect-misuse.md) [Incorrectly used escape sequences](incorrectly-used-escape-sequences.md) [Expr Modernization](expr-modernization.md) [Unquoted Special Parameters](unquoted-special-parameters.md) [Single-Iteration Loop](single-iteration-loop.md) [Prefer Find Over LS](prefer-find-over-ls.md) [Avoid Complex Logical Expressions](avoid-complex-logical-expressions.md) [Unquoted Array Expansion](unquoted-array-expansion.md) [PS Grep Alternative](ps-grep-alternative.md) [Unquoted Find Patterns](unquoted-find-patterns.md) [Unnecessary Variable Expansion](unnecessary-variable-expansion.md) [Use of if and then](use-of-if-and-then.md) [Read Lines with While Loop](read-lines-with-while-loop.md) [Incorrect Quoting in Trap Commands](incorrect-quoting-in-trap-commands.md) [Command Substitution Syntax](command-substitution-syntax.md)

# Incorrect Quoting in Trap Commands [Medium](severity/medium.md)

The issue with using double quotes in trap commands is that it can cause the variables and commands to expand immediately, rather than when the trap is executed. This can lead to unexpected behavior, as the values may not be what you expect at the time the trap is triggered. To avoid this, use single quotes to prevent the expansion until the trap is executed.

**Detector ID**

shell/incorrect-quoting-in-trap-commands@v1.0

**Category**

[Code Quality](categories/code-quality.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

-

**Tags**

-

* * *

#### Noncompliant example

```bash
1
2# Noncompliant: Double quotes cause immediate expansion of the date command.
3trap "echo \"Script finished at $(date)\"" EXIT
```

#### Compliant example

```bash
1
2# Compliant: Single quotes delay expansion until the `trap` is triggered
3trap 'echo "Script finished at $(date)"' EXIT
```

All content copied from https://docs.aws.amazon.com/.
