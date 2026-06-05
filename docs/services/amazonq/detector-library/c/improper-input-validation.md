---
title: "Improper Input Validation High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](divide-by-zero.md) [Return Stack Address](return-stack-address.md) [Unchecked Return Value](unchecked-return-value.md) [Incorrect Format Specifier](incorrect-format-specifier.md) [Unhandled Expression Result](unhandled-expression-result.md) [Path traversal](path-traversal.md) [Improper Input Validation](improper-input-validation.md) [Out Of Bounds Read](out-of-bounds-read.md) [Integer Overflow](integer-overflow.md) [Insecure use strtok function](insecure-use-strtok-fn.md) [Improper size of a memory buffer](improper-size-of-a-memory-buffer.md) [incomplete-cleanup](incomplete-cleanup.md) [Null pointer dereference](null-pointer-dereference.md) [Insecure Temporary File Or Directory](insecure-temporary-file-or-directory.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use Ato Fn](incorrect-use-ato-fn.md) [Loose File Permissions](loose-file-permissions.md) [Exposure of Sensitive Information](exposure-of-sensitive-information.md) [Out-of-bounds Write](out-of-bounds-write.md) [String Equality](string-equality.md)

# Improper Input Validation [High](severity/high.md)

Improper input validation can enable attacks and lead to unwanted behavior. Parts of the system may receive unintended input, which could result in altered control flow, arbitrary control of a resource, or arbitrary code execution.

**Detector ID**

c/improper-input-validation@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-20](https://cwe.mitre.org/data/definitions/20.html)

**Tags**

[\# top10-cwes](tags/top10-cwes.md)

* * *

#### Noncompliant example

```c
1#include <stdio.h>
2#include <string.h>
3
4void improperInputValidationNonCompliant(const char* username) {
5    printf("Enter username: ");
6    fgets(username, sizeof(username), stdin);
7    // Noncompliant: Input validation is needed to prevent user input from exceeding the allocated memory for `username`.
8    printf("Hello, %s!\n", username);
9}

```

#### Compliant example

```c
1#include <stdio.h>
2#include <string.h>
3
4void improperInputValidationCompliant(const char* input) {
5    char buffer[100]; // Assuming a maximum length of 100 characters
6
7    printf("Enter input: ");
8    scanf("%99s", buffer); // Limit input to 99 characters to leave space for null terminator
9
10    if(strlen(buffer) > 99) {
11        printf("Input exceeds maximum length\n");
12        return;
13    }
14    // Compliant: Validated input is copied to the provided const char* input
15    strcpy(input, buffer);
16
17    printf("You entered: %s\n", input);
18}

```

All content copied from https://docs.aws.amazon.com/.
