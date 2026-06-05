---
title: "Unsafe File Extension Critical"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](divide-by-zero.md) [Return Stack Address](return-stack-address.md) [Unchecked Return Value](unchecked-return-value.md) [Incorrect Format Specifier](incorrect-format-specifier.md) [Unhandled Expression Result](unhandled-expression-result.md) [Path traversal](path-traversal.md) [Improper Input Validation](improper-input-validation.md) [Out Of Bounds Read](out-of-bounds-read.md) [Integer Overflow](integer-overflow.md) [Insecure use strtok function](insecure-use-strtok-fn.md) [Improper size of a memory buffer](improper-size-of-a-memory-buffer.md) [incomplete-cleanup](incomplete-cleanup.md) [Null pointer dereference](null-pointer-dereference.md) [Insecure Temporary File Or Directory](insecure-temporary-file-or-directory.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use Ato Fn](incorrect-use-ato-fn.md) [Loose File Permissions](loose-file-permissions.md) [Exposure of Sensitive Information](exposure-of-sensitive-information.md) [Out-of-bounds Write](out-of-bounds-write.md) [String Equality](string-equality.md)

# Unsafe File Extension [Critical](severity/critical.md)

Insufficiently restricted file uploads can allow a file to be uploaded that runs malicious code. For example, a website that doesn't check the file extension of an image can be exploited by uploading a script with an extension, such as `.php` or `.asp`, that can be run on the server.

**Detector ID**

c/unsafe-file-extension@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-434](https://cwe.mitre.org/data/definitions/434.html)

**Tags**

[\# owasp-top10](tags/owasp-top10.md)

* * *

#### Compliant example

```c
1#include <stdio.h>
2
3void unsafeFileExtensionCompliant() {
4    // Compliant: Safe extension used with fopen example
5    FILE* fileFopen = fopen("example.txt", "r");
6    if (fileFopen != NULL) {
7        printf("File opened successfully using fopen.\n");
8        fclose(fileFopen);
9    } else {
10        printf("Error: Failed to open the file using fopen.\n");
11    }
12}

```

#### Noncompliant example

```c
1#include <stdio.h>
2
3void unsafeFileExtensionNonCompliant() {
4    // Noncompliant: Unsafe file extension used with fopen
5    FILE* fileFopen = fopen("example.bat", "rb");
6    if (fileFopen != NULL) {
7        printf("File opened successfully using fopen.\n");
8        fclose(fileFopen);
9    } else {
10        printf("Error: Failed to open the file using fopen.\n");
11    }
12}

```

All content copied from https://docs.aws.amazon.com/.
