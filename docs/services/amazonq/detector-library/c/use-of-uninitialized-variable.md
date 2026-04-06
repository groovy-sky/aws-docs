![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](https://docs.aws.amazon.com/amazonq/detector-library/c/use-of-uninitialized-variable) [Insecure Use strcat fn](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strcat-fn) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/c/sql-injection) [Bitwise Operator On Signed Operand](https://docs.aws.amazon.com/amazonq/detector-library/c/bitwise-operator-on-signed-operand) [Insecure use gets fn](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-gets-fn) [Random fd exhaustion](https://docs.aws.amazon.com/amazonq/detector-library/c/random-fd-exhaustion) [Redundant Free Usage](https://docs.aws.amazon.com/amazonq/detector-library/c/redundant-free-usage) [Insecure Use Memset](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-memset) [Divide By Zero.](https://docs.aws.amazon.com/amazonq/detector-library/c/divide-by-zero) [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/c/return-stack-address) [Unchecked Return Value](https://docs.aws.amazon.com/amazonq/detector-library/c/unchecked-return-value) [Incorrect Format Specifier](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-format-specifier) [Unhandled Expression Result](https://docs.aws.amazon.com/amazonq/detector-library/c/unhandled-expression-result) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/c/path-traversal) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-input-validation) [Out Of Bounds Read](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-read) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/c/integer-overflow) [Insecure use strtok function](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strtok-fn) [Improper size of a memory buffer](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-size-of-a-memory-buffer) [incomplete-cleanup](https://docs.aws.amazon.com/amazonq/detector-library/c/incomplete-cleanup) [Null pointer dereference](https://docs.aws.amazon.com/amazonq/detector-library/c/null-pointer-dereference) [Insecure Temporary File Or Directory](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-temporary-file-or-directory) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-buffer-access) [Incorrect Use Ato Fn](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-ato-fn) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/c/loose-file-permissions) [Exposure of Sensitive Information](https://docs.aws.amazon.com/amazonq/detector-library/c/exposure-of-sensitive-information) [Out-of-bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-write) [String Equality](https://docs.aws.amazon.com/amazonq/detector-library/c/string-equality)

# Use Of Uninitialized Variable [High](https://docs.aws.amazon.com/amazonq/detector-library/c/severity/high)

Your code contains an uninitialized variable, which can lead to unpredictable behavior and bugs in your program. Make sure to initialize all variables before using them to ensure consistent and reliable execution of your code. Addressing this issue will help maintain the stability and correctness of your application.

**Detector ID**

c/use-of-uninitialized-variable@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/c/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-457](https://cwe.mitre.org/data/definitions/457.html)

**Tags**

-

* * *

#### Noncompliant example

```c
1#include <stdio.h>
2
3int useOfUninitializedVariableNonCompliant() {
4  int x;  // x is not initialized
5  // Noncompliant: x has grabage value
6  return x + 10;
7}

```

#### Compliant example

```c
1#include <stdio.h>
2
3int useOfUninitializedVariableCompliant(int flag, int b) {
4  int a;
5  if (flag) {
6    a = b;
7  } else {
8    a = 10;
9  }
10  // Compliant: Variable is used
11  return a;
12}

```
