---
title: "Weak pseudorandom number generation High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](loose-file-permissions.md) [Sensitive information leak](sensitive-information-leak.md) [Missing Authorization](missing-authorization.md) [Return Stack Address](return-stack-address.md) [OS Command Injection](os-command-injection.md) [Use After Free](use-after-free.md) [Incorrect Comparison](incorrect-comparison.md) [off by one error](off-by-one-error.md) [Path traversal](path-traversal.md) [Insecure temporary file or directory](insecure-temporary-file-or-directory.md) [Insecure Cryptography](insecure-cryptography.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Unchecked Null Dereference](unchecked-null-dereference.md) [SQL injection](sql-injection.md) [Missing check on method output](missing-check-on-method-output.md) [Improper Restriction on Memory Buffer](improper-restriction-on-memory-buffer.md) [Multiple Locks](multiple-locks.md) [Improper Input Validation](improper-input-validation.md) [Null Pointer Dereference](null-pointer-dereference.md) [Use Of Redundant Code](use-of-redundant-code.md) [Improper Certificate Validation](improper-certificate-validation.md) [Improper Authentication](improper-authentication.md)

# Weak pseudorandom number generation [High](severity/high.md)

Insufficiently random generators or hardcoded seeds can make pseudorandom sequences predictable, which may lead to security vulnerabilities.

**Detector ID**

cpp/weak-random-number-generation@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-330](https://cwe.mitre.org/data/definitions/330.html)

**Tags**

[\# cryptography](tags/cryptography.md) [\# security-context](tags/security-context.md) [\# owasp-top10](tags/owasp-top10.md)

* * *

#### Noncompliant example

```cpp
1#include <stdio.h>
2#include <stdlib.h>
3
4enum { len = 12 };
5
6void weakRandomNumberGenerationNoncompliant(void) {
7    char id[len];
8	int r;
9	int num;
10
11    // Noncompliant: `rand()` is insecure to generate random numbers.
12    r = rand();
13	num = snprintf(id, len, "ID%-d", r);
14}

```

#### Compliant example

```cpp
1#include <stdio.h>
2#include <stdlib.h>
3
4enum { len = 12 };
5
6void weakRandomNumberGenerationCompliant() {
7    char id[len];
8	int r;
9	int num;
10
11    // Compliant: `random()` is secure to generate random numbers.
12    r = random();
13	num = snprintf(id, len, "ID%-d", r);
14}

```

All content copied from https://docs.aws.amazon.com/.
