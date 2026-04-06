![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](divide-by-zero.md) [Return Stack Address](return-stack-address.md) [Unchecked Return Value](unchecked-return-value.md) [Incorrect Format Specifier](incorrect-format-specifier.md) [Unhandled Expression Result](unhandled-expression-result.md) [Path traversal](path-traversal.md) [Improper Input Validation](improper-input-validation.md) [Out Of Bounds Read](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-read) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/c/integer-overflow) [Insecure use strtok function](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strtok-fn) [Improper size of a memory buffer](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-size-of-a-memory-buffer) [incomplete-cleanup](https://docs.aws.amazon.com/amazonq/detector-library/c/incomplete-cleanup) [Null pointer dereference](https://docs.aws.amazon.com/amazonq/detector-library/c/null-pointer-dereference) [Insecure Temporary File Or Directory](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-temporary-file-or-directory) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-buffer-access) [Incorrect Use Ato Fn](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-ato-fn) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/c/loose-file-permissions) [Exposure of Sensitive Information](https://docs.aws.amazon.com/amazonq/detector-library/c/exposure-of-sensitive-information) [Out-of-bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-write) [String Equality](https://docs.aws.amazon.com/amazonq/detector-library/c/string-equality)

# Out Of Bounds Read [High](https://docs.aws.amazon.com/amazonq/detector-library/c/severity/high)

This is a type of memory access error that occurs when a program reads data from a memory address outside of the bounds of a buffer. This can result in the program reading data that does not belong to it, which can cause crashes, incorrect behavior, or even security vulnerabilities.

**Detector ID**

c/out-of-bounds-read@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/c/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-125](https://cwe.mitre.org/data/definitions/125.html)

**Tags**

[\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/c/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/c/tags/top25-cwes)

* * *

#### Noncompliant example

```c
1#include <stdio.h>
2#include <string.h>
3#include <stddef.h>
4#include <stdlib.h>
5
6int outOfBoundsReadNonCompliant() {
7    int arr[5] = {1, 2, 3, 4, 5};
8    int index = 5;
9    // Noncompliant: Array indexing out of bounds
10    int value = arr[index];
11    printf("Value: %d\n", value);
12    return 0;
13}

```

#### Compliant example

```c
1#include <stdio.h>
2#include <string.h>
3#include <stddef.h>
4#include <stdlib.h>
5
6int outOfBoundsReadCompliant() {
7    int arr[5] = {1, 2, 3, 4, 5};
8    int index = 2; // Choose a valid index within the array bounds
9    // Compliant: Ensure index is within bounds before accessing the array
10    if (index >= 0 && index < 5) {
11        int value = arr[index];
12        printf("Value: %d\n", value);
13    } else {
14        printf("Invalid index\n");
15    }
16    return 0;
17}

```
