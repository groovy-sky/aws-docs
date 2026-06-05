---
title: "Insecure Use strcat fn High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](divide-by-zero.md) [Return Stack Address](return-stack-address.md) [Unchecked Return Value](unchecked-return-value.md) [Incorrect Format Specifier](incorrect-format-specifier.md) [Unhandled Expression Result](unhandled-expression-result.md) [Path traversal](path-traversal.md) [Improper Input Validation](improper-input-validation.md) [Out Of Bounds Read](out-of-bounds-read.md) [Integer Overflow](integer-overflow.md) [Insecure use strtok function](insecure-use-strtok-fn.md) [Improper size of a memory buffer](improper-size-of-a-memory-buffer.md) [incomplete-cleanup](incomplete-cleanup.md) [Null pointer dereference](null-pointer-dereference.md) [Insecure Temporary File Or Directory](insecure-temporary-file-or-directory.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use Ato Fn](incorrect-use-ato-fn.md) [Loose File Permissions](loose-file-permissions.md) [Exposure of Sensitive Information](exposure-of-sensitive-information.md) [Out-of-bounds Write](out-of-bounds-write.md) [String Equality](string-equality.md)

# Insecure Use strcat fn [High](severity/high.md)

strcat/strncat that can lead to buffer overflow vulnerabilities because it does not affirm the size of the destination array and do not automatically NULL-terminate strings.

**Detector ID**

c/insecure-use-strcat-fn@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-676](https://cwe.mitre.org/data/definitions/676.html) [CWE-120](https://cwe.mitre.org/data/definitions/120.html)

**Tags**

-

* * *

#### Noncompliant example

```c
1#include <strings.h>
2
3int DST_BUFFER_SIZE = 120;
4
5void insecureUseStrcatNonCompliant(char* src, char* dst) {
6    int n = DST_BUFFER_SIZE;
7    if ((dst != NULL) && (src != NULL) && (strlen(dst)+strlen(src)+1 <= n)) {
8        // Noncompliant: Does not affirm length
9        strcat(dst, src);
10        // Noncompliant: Hardcoded length passed
11        strncat(dst, src, 100);
12    }
13}

```

#### Compliant example

```c
1#include <strings.h>
2
3void insecureUseStrcatCompliant(char* src, char* dest, int dest_size) {
4    // Compliant: No hardcoded length
5    strncat(dest, src, dest_size - 1);
6    dest[dest_size - 1] = '\0';
7}

```

All content copied from https://docs.aws.amazon.com/.
