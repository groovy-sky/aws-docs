---
title: "Unsafe File Extension Critical"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](loose-file-permissions.md) [Sensitive information leak](sensitive-information-leak.md) [Missing Authorization](missing-authorization.md) [Return Stack Address](return-stack-address.md) [OS Command Injection](os-command-injection.md) [Use After Free](use-after-free.md) [Incorrect Comparison](incorrect-comparison.md) [off by one error](off-by-one-error.md) [Path traversal](path-traversal.md) [Insecure temporary file or directory](insecure-temporary-file-or-directory.md) [Insecure Cryptography](insecure-cryptography.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Unchecked Null Dereference](unchecked-null-dereference.md) [SQL injection](sql-injection.md) [Missing check on method output](missing-check-on-method-output.md) [Improper Restriction on Memory Buffer](improper-restriction-on-memory-buffer.md) [Multiple Locks](multiple-locks.md) [Improper Input Validation](improper-input-validation.md) [Null Pointer Dereference](null-pointer-dereference.md) [Use Of Redundant Code](use-of-redundant-code.md) [Improper Certificate Validation](improper-certificate-validation.md) [Improper Authentication](improper-authentication.md)

# Unsafe File Extension [Critical](severity/critical.md)

Unsafe file extensions like `.exe` or `.vbs` can execute code without consent. Especially from untrusted sources, risks allowing viruses, malware, or hackers to compromise your device security.

**Detector ID**

cpp/unsafe-file-extension@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-434](https://cwe.mitre.org/data/definitions/434.html)

**Tags**

[\# owasp-top10](tags/owasp-top10.md)

* * *

#### Noncompliant example

```cpp
1#include <iostream>
2#include <cstdio>
3#include <fstream>
4
5void unsafeFileExtensionNoncompliant() {
6    // Noncompliant: `fopen` opens a file with unsafe extension
7    FILE* fileFopen = fopen("example.bat", "rb");
8    if (fileFopen != nullptr) {
9        std::cout << "File opened successfully using fopen." << std::endl;
10        fclose(fileFopen);
11    } else {
12        std::cout << "Error: Failed to open the file using fopen." << std::endl;
13    }
14}

```

#### Compliant example

```cpp
1#include <iostream>
2#include <cstdio>
3#include <fstream>
4
5void unsafeFileExtensionCompliant() {
6    // Compliant: `fopen` opens a file with safe extension
7    FILE* fileFopen = fopen("example.txt", "r");
8    if (fileFopen != nullptr) {
9        std::cout << "File opened successfully using fopen." << std::endl;
10        fclose(fileFopen);
11    } else {
12        std::cout << "Error: Failed to open the file using fopen." << std::endl;
13    }
14}

```

All content copied from https://docs.aws.amazon.com/.
