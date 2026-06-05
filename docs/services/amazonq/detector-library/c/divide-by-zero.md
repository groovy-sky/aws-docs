---
title: "Divide By Zero. High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](divide-by-zero.md) [Return Stack Address](return-stack-address.md) [Unchecked Return Value](unchecked-return-value.md) [Incorrect Format Specifier](incorrect-format-specifier.md) [Unhandled Expression Result](unhandled-expression-result.md) [Path traversal](path-traversal.md) [Improper Input Validation](improper-input-validation.md) [Out Of Bounds Read](out-of-bounds-read.md) [Integer Overflow](integer-overflow.md) [Insecure use strtok function](insecure-use-strtok-fn.md) [Improper size of a memory buffer](improper-size-of-a-memory-buffer.md) [incomplete-cleanup](incomplete-cleanup.md) [Null pointer dereference](null-pointer-dereference.md) [Insecure Temporary File Or Directory](insecure-temporary-file-or-directory.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use Ato Fn](incorrect-use-ato-fn.md) [Loose File Permissions](loose-file-permissions.md) [Exposure of Sensitive Information](exposure-of-sensitive-information.md) [Out-of-bounds Write](out-of-bounds-write.md) [String Equality](string-equality.md)

# Divide By Zero. [High](severity/high.md)

Software flaws related to dividing by zero or performing other arithmetic operations that result in a divide-by-zero condition, can lead to unexpected behavior, application crashes, or security vulnerabilities if not properly handled.

**Detector ID**

c/divide-by-zero@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-369](https://cwe.mitre.org/data/definitions/369.html)

**Tags**

-

* * *

#### Noncompliant example

```c
1struct OptionalInt divideByZeroNonCompliant(int a, int b) {
2   struct OptionalInt result;
3  // While the following check correctly prevents signed integer overflows,
4  // it fails to prevent divide-by-zero errors. If `b` is equal to `0`, the
5  // application emits undefined behavior.
6  if ((a == INT_MIN) && (b == -1)) {
7    result.has_value = 0;
8    return result;
9  }
10  result.has_value = 1;
11  // Noncompliant: Performing division without checking if the denominator is zero will lead to division by zero errors
12  result.value = a / b;
13  return result; // causes undefined behavior if `b` is zero
14}

```

#### Compliant example

```c
1struct OptionalInt divideByZeroCompliant(int a, int b) {
2    struct OptionalInt result;
3
4    if ((b == 0) || ((a == INT_MIN) && (b == -1))) {
5        result.has_value = 0; // Indicates failure
6        return result;
7    }
8
9    result.has_value = 1;
10    // Compliant: Checking if the denominator is zero before dividing to avoid division by zero errors
11    result.value = a / b;
12    return result; // Check correctly prevents divide-by-zero and signed integer overflows
13}

```

All content copied from https://docs.aws.amazon.com/.
