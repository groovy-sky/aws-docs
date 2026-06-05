---
title: "Insecure temporary file or directory Medium"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](loose-file-permissions.md) [Sensitive information leak](sensitive-information-leak.md) [Missing Authorization](missing-authorization.md) [Return Stack Address](return-stack-address.md) [OS Command Injection](os-command-injection.md) [Use After Free](use-after-free.md) [Incorrect Comparison](incorrect-comparison.md) [off by one error](off-by-one-error.md) [Path traversal](path-traversal.md) [Insecure temporary file or directory](insecure-temporary-file-or-directory.md) [Insecure Cryptography](insecure-cryptography.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Unchecked Null Dereference](unchecked-null-dereference.md) [SQL injection](sql-injection.md) [Missing check on method output](missing-check-on-method-output.md) [Improper Restriction on Memory Buffer](improper-restriction-on-memory-buffer.md) [Multiple Locks](multiple-locks.md) [Improper Input Validation](improper-input-validation.md) [Null Pointer Dereference](null-pointer-dereference.md) [Use Of Redundant Code](use-of-redundant-code.md) [Improper Certificate Validation](improper-certificate-validation.md) [Improper Authentication](improper-authentication.md)

# Insecure temporary file or directory [Medium](severity/medium.md)

Insecure creation of temporary files and directories can introduce race condition vulnerabilities. Attackers can leverage these race conditions to carry out exploits like denial-of-service attacks or escalating their privileges. Proper security practices are required when generating temp files to mitigate these risks.

**Detector ID**

cpp/insecure-temporary-file-or-directory@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-377](https://cwe.mitre.org/data/definitions/377.html)

**Tags**

[\# availability](tags/availability.md) [\# race-condition](tags/race-condition.md) [\# owasp-top10](tags/owasp-top10.md)

* * *

#### Noncompliant example

```cpp
1#include <fstream>
2
3void insecureTemporaryFileOrDirectoryNoncompliant()
4{
5    char templateName[] = "/tmp/fileXXXXXX";
6	FILE* file = fopen(templateName, "w");
7    // Noncompliant: Used insecure temporary file.
8    mktemp(templateName);
9    fprintf(file, "This is unsafe content.\n");
10    fclose(file);
11}

```

#### Compliant example

```cpp
1#include <fstream>
2
3void insecureTemporaryFileOrDirectoryCompliant()
4{
5    char templateName[] = "fileXXXXXX";
6    // Compliant: `mkstemp` creates a unique file and returns a file descriptor.
7    int fileDescriptor = mkstemp(templateName);
8    FILE* file = fdopen(fileDescriptor, "w");
9    fprintf(file, "This is safe content.\n");
10    fclose(file);
11}

```

All content copied from https://docs.aws.amazon.com/.
