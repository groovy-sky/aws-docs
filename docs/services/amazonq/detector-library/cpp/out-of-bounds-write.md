---
title: "Out Of Bounds Write High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](loose-file-permissions.md) [Sensitive information leak](sensitive-information-leak.md) [Missing Authorization](missing-authorization.md) [Return Stack Address](return-stack-address.md) [OS Command Injection](os-command-injection.md) [Use After Free](use-after-free.md) [Incorrect Comparison](incorrect-comparison.md) [off by one error](off-by-one-error.md) [Path traversal](path-traversal.md) [Insecure temporary file or directory](insecure-temporary-file-or-directory.md) [Insecure Cryptography](insecure-cryptography.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Unchecked Null Dereference](unchecked-null-dereference.md) [SQL injection](sql-injection.md) [Missing check on method output](missing-check-on-method-output.md) [Improper Restriction on Memory Buffer](improper-restriction-on-memory-buffer.md) [Multiple Locks](multiple-locks.md) [Improper Input Validation](improper-input-validation.md) [Null Pointer Dereference](null-pointer-dereference.md) [Use Of Redundant Code](use-of-redundant-code.md) [Improper Certificate Validation](improper-certificate-validation.md) [Improper Authentication](improper-authentication.md)

# Out Of Bounds Write [High](severity/high.md)

When software attempts to write data beyond the confines of a fixed-size buffer, array, or allocated memory region. To address and mitigate this vulnerability effectively, implementing robust input validation and leveraging memory safety mechanisms are crucial steps.

**Detector ID**

cpp/out-of-bounds-write@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-787](https://cwe.mitre.org/data/definitions/787.html)

**Tags**

[\# top25-cwes](tags/top25-cwes.md)

* * *

#### Noncompliant example

```cpp
1void outOfBoundWriteNoncompliant()
2{
3    // Declaring an array named id_sequence with a size of 3 integers
4    int id_sequence[3] = {1, 2, 3};
5    // Noncompliant: Attempting to assign a value to the fourth element.
6    id_sequence[4] = 456;
7}

```

#### Compliant example

```cpp
1void outOfBoundWriteCompliant() {
2    int arr[3] = {1, 2, 3};
3    // Compliant: Accessing indices within array bounds
4    for (int i = 0; i < 3; ++i)
5    {
6        arr[i] = i * 2;
7    }
8}

```

All content copied from https://docs.aws.amazon.com/.
