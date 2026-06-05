---
title: "off by one error High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](loose-file-permissions.md) [Sensitive information leak](sensitive-information-leak.md) [Missing Authorization](missing-authorization.md) [Return Stack Address](return-stack-address.md) [OS Command Injection](os-command-injection.md) [Use After Free](use-after-free.md) [Incorrect Comparison](incorrect-comparison.md) [off by one error](off-by-one-error.md) [Path traversal](path-traversal.md) [Insecure temporary file or directory](insecure-temporary-file-or-directory.md) [Insecure Cryptography](insecure-cryptography.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Unchecked Null Dereference](unchecked-null-dereference.md) [SQL injection](sql-injection.md) [Missing check on method output](missing-check-on-method-output.md) [Improper Restriction on Memory Buffer](improper-restriction-on-memory-buffer.md) [Multiple Locks](multiple-locks.md) [Improper Input Validation](improper-input-validation.md) [Null Pointer Dereference](null-pointer-dereference.md) [Use Of Redundant Code](use-of-redundant-code.md) [Improper Certificate Validation](improper-certificate-validation.md) [Improper Authentication](improper-authentication.md)

# off by one error [High](severity/high.md)

Off-by-one errors are programming mistakes where loops or array indices are improperly incremented or decremented by one, resulting in unintended behavior. These errors often lead to out-of-bounds memory access, incorrect data processing, or program crashes. Mitigation involves ensuring correct loop boundaries, validating indices, and implementing thorough boundary checks to prevent such issues.

**Detector ID**

cpp/off-by-one-error@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-193](https://cwe.mitre.org/data/definitions/193.html)

**Tags**

-

* * *

#### Noncompliant example

```cpp
1#include <stdio.h>
2
3void offByOneErrorNoncompliant() {
4    int id_sequence[3];
5
6    id_sequence[0] = 123;
7    id_sequence[1] = 234;
8    id_sequence[2] = 345;
9    // Noncompliant: Attempting to access index 3 out of bound
10    id_sequence[3] = 456;
11}

```

#### Compliant example

```cpp
1#include <stdio.h>
2
3void offByOneErrorCompliant() {
4
5    int id_sequence[4]; // Increase array size to accommodate the additional element
6
7    id_sequence[0] = 123;
8    id_sequence[1] = 234;
9    id_sequence[2] = 345;
10    // Compliant: This is a valid index within the bounds of the array.
11    id_sequence[3] = 456;
12}

```

All content copied from https://docs.aws.amazon.com/.
