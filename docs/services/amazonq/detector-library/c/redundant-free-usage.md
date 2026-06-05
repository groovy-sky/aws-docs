---
title: "Redundant Free Usage High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](divide-by-zero.md) [Return Stack Address](return-stack-address.md) [Unchecked Return Value](unchecked-return-value.md) [Incorrect Format Specifier](incorrect-format-specifier.md) [Unhandled Expression Result](unhandled-expression-result.md) [Path traversal](path-traversal.md) [Improper Input Validation](improper-input-validation.md) [Out Of Bounds Read](out-of-bounds-read.md) [Integer Overflow](integer-overflow.md) [Insecure use strtok function](insecure-use-strtok-fn.md) [Improper size of a memory buffer](improper-size-of-a-memory-buffer.md) [incomplete-cleanup](incomplete-cleanup.md) [Null pointer dereference](null-pointer-dereference.md) [Insecure Temporary File Or Directory](insecure-temporary-file-or-directory.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use Ato Fn](incorrect-use-ato-fn.md) [Loose File Permissions](loose-file-permissions.md) [Exposure of Sensitive Information](exposure-of-sensitive-information.md) [Out-of-bounds Write](out-of-bounds-write.md) [String Equality](string-equality.md)

# Redundant Free Usage [High](severity/high.md)

The product calls free method twice on the same memory address, which leads to modification of unexpected memory locations.

**Detector ID**

c/redundant-free-usage@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-415](https://cwe.mitre.org/data/definitions/415.html)

**Tags**

-

* * *

#### Noncompliant example

```c
1#include <stdlib.h>
2#include <string.h>
3
4int redundantFreeUsageNonCompliant(char *argv[]) {
5    char *buf1;
6    char *buf2;
7    buf1 = (char *) malloc(sizeof(char) * 10);
8    free(buf1);
9    buf2 = (char *) malloc(sizeof(char) * 5);
10    strncpy(buf2, argv[1], 1);
11    // Noncompliant: Redundent use of `free` on buf1 without memory reallocation
12    free(buf1);
13    free(buf2);
14}

```

#### Compliant example

```c
1#include <stdlib.h>
2#include <string.h>
3
4int redundantFreeUsageCompliant() {
5    char *var = malloc(sizeof(char) * 10);
6    free(var);
7    var = malloc(sizeof(char) * 10);
8    // Compliant: Use of free on variable after memory reallocation
9    free(var);
10    return 0;
11}

```

All content copied from https://docs.aws.amazon.com/.
