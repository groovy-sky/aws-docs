![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/weak-random-number-generation) [Missing Default in Switch](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-default-in-switch) [Unsafe File Extension](https://docs.aws.amazon.com/amazonq/detector-library/cpp/unsafe-file-extension) [Incorrect Order Of setuid and setgid](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-order-setuid-setgid) [Out Of Bounds Read](https://docs.aws.amazon.com/amazonq/detector-library/cpp/out-of-bounds-read) [Out Of Bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/cpp/out-of-bounds-write) [Thread safety violation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/thread-safety-violation) [Incorrect Pointer Subtraction](https://docs.aws.amazon.com/amazonq/detector-library/cpp/pointer-subtraction) [File System Access](https://docs.aws.amazon.com/amazonq/detector-library/cpp/file-system-access) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-buffer-access) [Incorrect Use of Sizeof](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-use-of-sizeof) [Incorrect Pointer Scaling](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-pointer-scaling) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/cpp/loose-file-permissions) [Sensitive information leak](https://docs.aws.amazon.com/amazonq/detector-library/cpp/sensitive-information-leak) [Missing Authorization](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-authorization) [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/cpp/return-stack-address) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/os-command-injection) [Use After Free](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-after-free) [Incorrect Comparison](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-comparison) [off by one error](https://docs.aws.amazon.com/amazonq/detector-library/cpp/off-by-one-error) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/cpp/path-traversal) [Insecure temporary file or directory](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-temporary-file-or-directory) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-cryptography) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-connection) [Unchecked Null Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/unchecked-null-dereference) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/sql-injection) [Missing check on method output](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-check-on-method-output) [Improper Restriction on Memory Buffer](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-restriction-on-memory-buffer) [Multiple Locks](https://docs.aws.amazon.com/amazonq/detector-library/cpp/multiple-locks) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-input-validation) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/null-pointer-dereference) [Use Of Redundant Code](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-of-redundant-code) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-certificate-validation) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-authentication)

# Weak pseudorandom number generation [High](https://docs.aws.amazon.com/amazonq/detector-library/cpp/severity/high)

Insufficiently random generators or hardcoded seeds can make pseudorandom sequences predictable, which may lead to security vulnerabilities.

**Detector ID**

cpp/weak-random-number-generation@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/cpp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-330](https://cwe.mitre.org/data/definitions/330.html)

**Tags**

[\# cryptography](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags/cryptography) [\# security-context](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags/security-context) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags/owasp-top10)

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
