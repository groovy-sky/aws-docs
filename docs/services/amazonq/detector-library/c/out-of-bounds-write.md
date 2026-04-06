![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](divide-by-zero.md) [Return Stack Address](return-stack-address.md) [Unchecked Return Value](unchecked-return-value.md) [Incorrect Format Specifier](incorrect-format-specifier.md) [Unhandled Expression Result](unhandled-expression-result.md) [Path traversal](path-traversal.md) [Improper Input Validation](improper-input-validation.md) [Out Of Bounds Read](out-of-bounds-read.md) [Integer Overflow](integer-overflow.md) [Insecure use strtok function](insecure-use-strtok-fn.md) [Improper size of a memory buffer](improper-size-of-a-memory-buffer.md) [incomplete-cleanup](incomplete-cleanup.md) [Null pointer dereference](null-pointer-dereference.md) [Insecure Temporary File Or Directory](insecure-temporary-file-or-directory.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use Ato Fn](incorrect-use-ato-fn.md) [Loose File Permissions](loose-file-permissions.md) [Exposure of Sensitive Information](exposure-of-sensitive-information.md) [Out-of-bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-write) [String Equality](https://docs.aws.amazon.com/amazonq/detector-library/c/string-equality)

# Out-of-bounds Write [High](https://docs.aws.amazon.com/amazonq/detector-library/c/severity/high)

This is a type of memory access error that occurs when a program writes data from a memory address outside of the bounds of a buffer. This can result in the program writing data that does not belong to it, which can cause crashes, incorrect behavior, or even security vulnerabilities.

**Detector ID**

c/out-of-bounds-write@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/c/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-787](https://cwe.mitre.org/data/definitions/787.html) [CWE-119](https://cwe.mitre.org/data/definitions/119.html)

**Tags**

[\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/c/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/c/tags/top25-cwes)

* * *

#### Noncompliant example

```c
1void outOfBoundsWriteNonCompliant(){
2    // Declaring an array named id_sequence with a size of 3 integers
3    int id_sequence[3];
4    id_sequence[0] = 123;
5    id_sequence[1] = 234;
6    id_sequence[2] = 345;
7    // Noncompliant: Attempting to assign a value to the fourth element (out of bounds)
8    id_sequence[3] = 456;
9}

```

#### Compliant example

```c
1#include <stdio.h>
2
3void outOfBoundsWriteCompliant(){
4
5  // Compliant: Ensuring correct loop bounds
6  int arr[3] = {1, 2, 3};
7  for (int i = 0; i < 3; ++i) {
8    arr[i] = i * 2; // Accessing indices within array bounds
9  }
10}

```
