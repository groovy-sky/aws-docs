---
title: "Out Of Bounds Read High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](loose-file-permissions.md) [Sensitive information leak](sensitive-information-leak.md) [Missing Authorization](missing-authorization.md) [Return Stack Address](return-stack-address.md) [OS Command Injection](os-command-injection.md) [Use After Free](use-after-free.md) [Incorrect Comparison](incorrect-comparison.md) [off by one error](off-by-one-error.md) [Path traversal](path-traversal.md) [Insecure temporary file or directory](insecure-temporary-file-or-directory.md) [Insecure Cryptography](insecure-cryptography.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Unchecked Null Dereference](unchecked-null-dereference.md) [SQL injection](sql-injection.md) [Missing check on method output](missing-check-on-method-output.md) [Improper Restriction on Memory Buffer](improper-restriction-on-memory-buffer.md) [Multiple Locks](multiple-locks.md) [Improper Input Validation](improper-input-validation.md) [Null Pointer Dereference](null-pointer-dereference.md) [Use Of Redundant Code](use-of-redundant-code.md) [Improper Certificate Validation](improper-certificate-validation.md) [Improper Authentication](improper-authentication.md)

# Out Of Bounds Read [High](severity/high.md)

This is a type of memory access error that occurs when a program reads data from a memory address outside of the bounds of a buffer. This can result in the program reading data that does not belong to it, which can cause crashes, incorrect behavior, or even security vulnerabilities.

**Detector ID**

cpp/out-of-bounds-read@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-125](https://cwe.mitre.org/data/definitions/125.html)

**Tags**

[\# owasp-top10](tags/owasp-top10.md) [\# top25-cwes](tags/top25-cwes.md)

* * *

#### Noncompliant example

```cpp
1#include <cstring>
2
3void outOfBoundReadsNoncompliant() {
4   int MAX = 10;
5   char array1[MAX];
6   int  array2[MAX];
7   // Noncompliant: The call to `memcpy()` reads memory from outside the allocated bounds of character array, which contains MAX elements of type char, while integer array contains MAX elements of type int.
8   memcpy(array2, array1, sizeof(array2));
9}

```

#### Compliant example

```cpp
1#include <cstring>
2
3void outOfBoundReadscompliant() {
4   int MAX = 10;
5   int array1[MAX];
6   int  array2[MAX];
7   // Compliant: Both arrays are of same data type.
8   memcpy(array2, array1, sizeof(array2));
9}

```

All content copied from https://docs.aws.amazon.com/.
