---
title: "Incorrect Pointer Scaling High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](loose-file-permissions.md) [Sensitive information leak](sensitive-information-leak.md) [Missing Authorization](missing-authorization.md) [Return Stack Address](return-stack-address.md) [OS Command Injection](os-command-injection.md) [Use After Free](use-after-free.md) [Incorrect Comparison](incorrect-comparison.md) [off by one error](off-by-one-error.md) [Path traversal](path-traversal.md) [Insecure temporary file or directory](insecure-temporary-file-or-directory.md) [Insecure Cryptography](insecure-cryptography.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Unchecked Null Dereference](unchecked-null-dereference.md) [SQL injection](sql-injection.md) [Missing check on method output](missing-check-on-method-output.md) [Improper Restriction on Memory Buffer](improper-restriction-on-memory-buffer.md) [Multiple Locks](multiple-locks.md) [Improper Input Validation](improper-input-validation.md) [Null Pointer Dereference](null-pointer-dereference.md) [Use Of Redundant Code](use-of-redundant-code.md) [Improper Certificate Validation](improper-certificate-validation.md) [Improper Authentication](improper-authentication.md)

# Incorrect Pointer Scaling [High](severity/high.md)

We observed that your code contains instances of incorrect pointer scaling, potentially leading to unexpected behavior and vulnerabilities. To mitigate this issue, ensure proper pointer arithmetic by using the correct pointer types and letting the language handle scaling automatically. Avoid casting pointers to narrower types unnecessarily.

**Detector ID**

cpp/incorrect-pointer-scaling@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-468](https://cwe.mitre.org/data/definitions/468.html)

**Tags**

-

* * *

#### Noncompliant example

```cpp
1#include <iostream>
2
3inline char incorrectPointerScalingNoncompliant() {
4      int intArray[5] = { 1, 2, 3, 4, 5 };
5      char *charPointer = (char *)intArray;
6      // Noncompliant: The pointer arithmetic uses type char*, so the offset is not scaled by sizeof(int).
7      return *(charPointer + i);
8}

```

#### Compliant example

```cpp
1#include <iostream>
2
3inline char incorrectPointerScalingCompliant() {
4      int intArray[10] = { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 };
5      int *intPointer = intArray;
6      // Compliant: The offset is automatically scaled by sizeof(int).
7      return *(intPointer + i);
8}

```

All content copied from https://docs.aws.amazon.com/.
