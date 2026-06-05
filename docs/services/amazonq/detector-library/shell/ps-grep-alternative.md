---
title: "PS Grep Alternative Medium"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Shell detectors(16/16)

[Avoid LS-Grep](avoid-ls-grep.md) [Sudo Redirect Misuse](sudo-redirect-misuse.md) [Incorrectly used escape sequences](incorrectly-used-escape-sequences.md) [Expr Modernization](expr-modernization.md) [Unquoted Special Parameters](unquoted-special-parameters.md) [Single-Iteration Loop](single-iteration-loop.md) [Prefer Find Over LS](prefer-find-over-ls.md) [Avoid Complex Logical Expressions](avoid-complex-logical-expressions.md) [Unquoted Array Expansion](unquoted-array-expansion.md) [PS Grep Alternative](ps-grep-alternative.md) [Unquoted Find Patterns](unquoted-find-patterns.md) [Unnecessary Variable Expansion](unnecessary-variable-expansion.md) [Use of if and then](use-of-if-and-then.md) [Read Lines with While Loop](read-lines-with-while-loop.md) [Incorrect Quoting in Trap Commands](incorrect-quoting-in-trap-commands.md) [Command Substitution Syntax](command-substitution-syntax.md)

# PS Grep Alternative [Medium](severity/medium.md)

Grepping 'ps' output can be inefficient and potentially error-prone when searching for specific processes. Using 'pgrep' instead provides a more robust and efficient solution for process identification and management.

**Detector ID**

shell/ps-grep-alternative@v1.0

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
2# Noncompliant: Grepping `ps` output.
3if ps aux | grep -v grep | grep "apache2" > /dev/null; then
4    echo "Apache is running"
5else
6    echo "Apache is not running"
7fi
```

#### Compliant example

```bash
1
2# Compliant: Using `pgrep` instead of grepping `ps` output.
3if pgrep -x "apache2" > /dev/null; then
4    echo "Apache is running"
5else
6    echo "Apache is not running"
7fi
```

All content copied from https://docs.aws.amazon.com/.
