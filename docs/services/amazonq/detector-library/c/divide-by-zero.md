![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](https://docs.aws.amazon.com/amazonq/detector-library/c/divide-by-zero) [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/c/return-stack-address) [Unchecked Return Value](https://docs.aws.amazon.com/amazonq/detector-library/c/unchecked-return-value) [Incorrect Format Specifier](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-format-specifier) [Unhandled Expression Result](https://docs.aws.amazon.com/amazonq/detector-library/c/unhandled-expression-result) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/c/path-traversal) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-input-validation) [Out Of Bounds Read](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-read) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/c/integer-overflow) [Insecure use strtok function](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strtok-fn) [Improper size of a memory buffer](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-size-of-a-memory-buffer) [incomplete-cleanup](https://docs.aws.amazon.com/amazonq/detector-library/c/incomplete-cleanup) [Null pointer dereference](https://docs.aws.amazon.com/amazonq/detector-library/c/null-pointer-dereference) [Insecure Temporary File Or Directory](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-temporary-file-or-directory) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-buffer-access) [Incorrect Use Ato Fn](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-ato-fn) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/c/loose-file-permissions) [Exposure of Sensitive Information](https://docs.aws.amazon.com/amazonq/detector-library/c/exposure-of-sensitive-information) [Out-of-bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-write) [String Equality](https://docs.aws.amazon.com/amazonq/detector-library/c/string-equality)

# Divide By Zero. [High](https://docs.aws.amazon.com/amazonq/detector-library/c/severity/high)

Software flaws related to dividing by zero or performing other arithmetic operations that result in a divide-by-zero condition, can lead to unexpected behavior, application crashes, or security vulnerabilities if not properly handled.

**Detector ID**

c/divide-by-zero@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/c/categories/security)

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
