---
title: "Loose File Permissions High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](loose-file-permissions.md) [Sensitive information leak](sensitive-information-leak.md) [Missing Authorization](missing-authorization.md) [Return Stack Address](return-stack-address.md) [OS Command Injection](os-command-injection.md) [Use After Free](use-after-free.md) [Incorrect Comparison](incorrect-comparison.md) [off by one error](off-by-one-error.md) [Path traversal](path-traversal.md) [Insecure temporary file or directory](insecure-temporary-file-or-directory.md) [Insecure Cryptography](insecure-cryptography.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Unchecked Null Dereference](unchecked-null-dereference.md) [SQL injection](sql-injection.md) [Missing check on method output](missing-check-on-method-output.md) [Improper Restriction on Memory Buffer](improper-restriction-on-memory-buffer.md) [Multiple Locks](multiple-locks.md) [Improper Input Validation](improper-input-validation.md) [Null Pointer Dereference](null-pointer-dereference.md) [Use Of Redundant Code](use-of-redundant-code.md) [Improper Certificate Validation](improper-certificate-validation.md) [Improper Authentication](improper-authentication.md)

# Loose File Permissions [High](severity/high.md)

File and directory permissions should be granted to specific users and groups. Granting permissions to wildcards, such as everyone or others, can lead to privilege escalations, leakage of sensitive information, and inadvertently running malicious code.

**Detector ID**

cpp/loose-file-permissions@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-732](https://cwe.mitre.org/data/definitions/732.html) [CWE-266](https://cwe.mitre.org/data/definitions/266.html)

**Tags**

[\# access-control](tags/access-control.md) [\# owasp-top10](tags/owasp-top10.md) [\# top25-cwes](tags/top25-cwes.md)

* * *

#### Noncompliant example

```cpp
1#include <stdio.h>
2
3void looseFilePermissionsNoncompliant() {
4    // Noncompliant: `S_IRWXU | S_IRWXG | S_IRWXO` will grant read, write, and execute permissions to the owner, group, and others to this newly created file.
5    open("myfile.txt", O_CREAT, S_IRWXU | S_IRWXG | S_IRWXO);
6}

```

#### Compliant example

```cpp
1#include <stdio.h>
2
3void looseFilePermissionsCompliant() {
4    // Compliant: `S_IRWXU | S_IRWXG` will grant read, write, and execute permissions to the owner and group to this newly created file.
5    open("myfile.txt", O_CREAT, S_IRWXU | S_IRWXG);
6}

```

All content copied from https://docs.aws.amazon.com/.
