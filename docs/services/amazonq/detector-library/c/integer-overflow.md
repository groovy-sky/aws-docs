---
title: "Integer Overflow High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](divide-by-zero.md) [Return Stack Address](return-stack-address.md) [Unchecked Return Value](unchecked-return-value.md) [Incorrect Format Specifier](incorrect-format-specifier.md) [Unhandled Expression Result](unhandled-expression-result.md) [Path traversal](path-traversal.md) [Improper Input Validation](improper-input-validation.md) [Out Of Bounds Read](out-of-bounds-read.md) [Integer Overflow](integer-overflow.md) [Insecure use strtok function](insecure-use-strtok-fn.md) [Improper size of a memory buffer](improper-size-of-a-memory-buffer.md) [incomplete-cleanup](incomplete-cleanup.md) [Null pointer dereference](null-pointer-dereference.md) [Insecure Temporary File Or Directory](insecure-temporary-file-or-directory.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use Ato Fn](incorrect-use-ato-fn.md) [Loose File Permissions](loose-file-permissions.md) [Exposure of Sensitive Information](exposure-of-sensitive-information.md) [Out-of-bounds Write](out-of-bounds-write.md) [String Equality](string-equality.md)

# Integer Overflow [High](severity/high.md)

An integer overflow might occur when the input or resulting value is too large to store in associated representation. This can result in a critical security issue when it is used to make security decisions.

**Detector ID**

c/integer-overflow@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-190](https://cwe.mitre.org/data/definitions/190.html)

**Tags**

-

* * *

#### Noncompliant example

```c
1#include <stdlib.h>
2
3void integerOverflowNoncompliant(int *ptr, size_t offset) {
4    // Noncompliant: Perform pointer arithmetic without checking for potential integer overflow.
5    int *result = ptr + offset;
6}

```

#### Compliant example

```c
1#include <stdlib.h>
2#include <stdio.h>
3#include <stdint.h>
4
5void integerOverflowCompliant(int *ptr, size_t offset) {
6    // Compliant: Safer pointer arithmetic with proper check
7    if (offset <= SIZE_MAX / sizeof(int))
8    {
9        int *result = ptr + offset;
10        // Use 'result'
11    }
12    else
13    {
14        fprintf(stderr, "Overflow detected in pointer arithmetic.\n");
15    }
16}

```

All content copied from https://docs.aws.amazon.com/.
