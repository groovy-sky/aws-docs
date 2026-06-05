---
title: "Improper size of a memory buffer High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](divide-by-zero.md) [Return Stack Address](return-stack-address.md) [Unchecked Return Value](unchecked-return-value.md) [Incorrect Format Specifier](incorrect-format-specifier.md) [Unhandled Expression Result](unhandled-expression-result.md) [Path traversal](path-traversal.md) [Improper Input Validation](improper-input-validation.md) [Out Of Bounds Read](out-of-bounds-read.md) [Integer Overflow](integer-overflow.md) [Insecure use strtok function](insecure-use-strtok-fn.md) [Improper size of a memory buffer](improper-size-of-a-memory-buffer.md) [incomplete-cleanup](incomplete-cleanup.md) [Null pointer dereference](null-pointer-dereference.md) [Insecure Temporary File Or Directory](insecure-temporary-file-or-directory.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use Ato Fn](incorrect-use-ato-fn.md) [Loose File Permissions](loose-file-permissions.md) [Exposure of Sensitive Information](exposure-of-sensitive-information.md) [Out-of-bounds Write](out-of-bounds-write.md) [String Equality](string-equality.md)

# Improper size of a memory buffer [High](severity/high.md)

The product performs operations on a memory buffer, but it can read from or write to a memory location that is outside of the intended boundary of the buffer. As a result, an attacker may be able to execute arbitrary code, alter the intended control flow, read sensitive information, or cause the system to crash.

**Detector ID**

c/improper-size-of-a-memory-buffer@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-119](https://cwe.mitre.org/data/definitions/119.html) [CWE-676](https://cwe.mitre.org/data/definitions/676.html) [CWE-120](https://cwe.mitre.org/data/definitions/120.html)

**Tags**

[\# top25-cwes](tags/top25-cwes.md)

* * *

#### Noncompliant example

```c
1#include <stdio.h>
2#include <unistd.h>
3#include <fcntl.h>
4
5void improperSizeOfAMemoryBufferNonCompliant() {
6   int fd;
7   char buff[1024];
8   char path[] = "Documents/example.txt";
9
10   fd = open(path, O_RDONLY);
11
12   int size = 1027;
13   // Noncompliant: size argument exceeds the actual size of the buffer.
14   read(fd, buff, size);
15
16   printf("\n\n%s\n\n",buff);
17}

```

#### Compliant example

```c
1int improperSizeOfAMemoryBufferCompliant()
2{
3    char array[10];
4    initialize(array);
5    // Compliant: size argument is same as the actual size of the buffer.
6    char *pos = memchr(array, '@', sizeof(array));
7
8    return 0;
9}

```

All content copied from https://docs.aws.amazon.com/.
