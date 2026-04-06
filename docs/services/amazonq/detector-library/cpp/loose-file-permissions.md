![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/cpp/loose-file-permissions) [Sensitive information leak](https://docs.aws.amazon.com/amazonq/detector-library/cpp/sensitive-information-leak) [Missing Authorization](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-authorization) [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/cpp/return-stack-address) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/os-command-injection) [Use After Free](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-after-free) [Incorrect Comparison](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-comparison) [off by one error](https://docs.aws.amazon.com/amazonq/detector-library/cpp/off-by-one-error) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/cpp/path-traversal) [Insecure temporary file or directory](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-temporary-file-or-directory) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-cryptography) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-connection) [Unchecked Null Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/unchecked-null-dereference) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/sql-injection) [Missing check on method output](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-check-on-method-output) [Improper Restriction on Memory Buffer](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-restriction-on-memory-buffer) [Multiple Locks](https://docs.aws.amazon.com/amazonq/detector-library/cpp/multiple-locks) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-input-validation) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/null-pointer-dereference) [Use Of Redundant Code](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-of-redundant-code) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-certificate-validation) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-authentication)

# Loose File Permissions [High](https://docs.aws.amazon.com/amazonq/detector-library/cpp/severity/high)

File and directory permissions should be granted to specific users and groups. Granting permissions to wildcards, such as everyone or others, can lead to privilege escalations, leakage of sensitive information, and inadvertently running malicious code.

**Detector ID**

cpp/loose-file-permissions@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/cpp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-732](https://cwe.mitre.org/data/definitions/732.html) [CWE-266](https://cwe.mitre.org/data/definitions/266.html)

**Tags**

[\# access-control](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags/access-control) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags/top25-cwes)

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
