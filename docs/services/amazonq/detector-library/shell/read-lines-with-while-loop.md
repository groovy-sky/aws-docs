![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Shell detectors(16/16)

[Avoid LS-Grep](avoid-ls-grep.md) [Sudo Redirect Misuse](sudo-redirect-misuse.md) [Incorrectly used escape sequences](incorrectly-used-escape-sequences.md) [Expr Modernization](expr-modernization.md) [Unquoted Special Parameters](unquoted-special-parameters.md) [Single-Iteration Loop](single-iteration-loop.md) [Prefer Find Over LS](prefer-find-over-ls.md) [Avoid Complex Logical Expressions](avoid-complex-logical-expressions.md) [Unquoted Array Expansion](unquoted-array-expansion.md) [PS Grep Alternative](ps-grep-alternative.md) [Unquoted Find Patterns](unquoted-find-patterns.md) [Unnecessary Variable Expansion](unnecessary-variable-expansion.md) [Use of if and then](use-of-if-and-then.md) [Read Lines with While Loop](https://docs.aws.amazon.com/amazonq/detector-library/shell/read-lines-with-while-loop) [Incorrect Quoting in Trap Commands](https://docs.aws.amazon.com/amazonq/detector-library/shell/incorrect-quoting-in-trap-commands) [Command Substitution Syntax](https://docs.aws.amazon.com/amazonq/detector-library/shell/command-substitution-syntax)

# Read Lines with While Loop [Medium](https://docs.aws.amazon.com/amazonq/detector-library/shell/severity/medium)

Using for loops to process file contents can lead to unexpected word splitting and glob expansion, potentially altering the original content. Implement while-read loops to accurately process files line-by-line, preserving the original structure and content of each line.

**Detector ID**

shell/read-lines-with-while-loop@v1.0

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
2# Noncompliant: Word splitting and glob expansion can break on special characters.
3echo "Processing lines from file.txt:"
4for line in $(cat file.txt | grep -v '^#')
5do
6    echo "Processing: $line"
7done
```

#### Compliant example

```bash
1
2# Compliant: while read loop preserves whitespace and handles special characters.
3echo "Processing lines from file.txt:"
4while IFS= read -r line
5do
6    [[ $line =~ ^# ]] && continue
7    echo "Processing: $line"
8done < file.txt
```
