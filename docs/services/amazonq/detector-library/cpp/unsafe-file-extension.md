![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](https://docs.aws.amazon.com/amazonq/detector-library/cpp/unsafe-file-extension) [Incorrect Order Of setuid and setgid](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-order-setuid-setgid) [Out Of Bounds Read](https://docs.aws.amazon.com/amazonq/detector-library/cpp/out-of-bounds-read) [Out Of Bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/cpp/out-of-bounds-write) [Thread safety violation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/thread-safety-violation) [Incorrect Pointer Subtraction](https://docs.aws.amazon.com/amazonq/detector-library/cpp/pointer-subtraction) [File System Access](https://docs.aws.amazon.com/amazonq/detector-library/cpp/file-system-access) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-buffer-access) [Incorrect Use of Sizeof](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-use-of-sizeof) [Incorrect Pointer Scaling](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-pointer-scaling) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/cpp/loose-file-permissions) [Sensitive information leak](https://docs.aws.amazon.com/amazonq/detector-library/cpp/sensitive-information-leak) [Missing Authorization](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-authorization) [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/cpp/return-stack-address) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/os-command-injection) [Use After Free](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-after-free) [Incorrect Comparison](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-comparison) [off by one error](https://docs.aws.amazon.com/amazonq/detector-library/cpp/off-by-one-error) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/cpp/path-traversal) [Insecure temporary file or directory](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-temporary-file-or-directory) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-cryptography) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-connection) [Unchecked Null Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/unchecked-null-dereference) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/sql-injection) [Missing check on method output](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-check-on-method-output) [Improper Restriction on Memory Buffer](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-restriction-on-memory-buffer) [Multiple Locks](https://docs.aws.amazon.com/amazonq/detector-library/cpp/multiple-locks) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-input-validation) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/null-pointer-dereference) [Use Of Redundant Code](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-of-redundant-code) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-certificate-validation) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-authentication)

# Unsafe File Extension [Critical](https://docs.aws.amazon.com/amazonq/detector-library/cpp/severity/critical)

Unsafe file extensions like `.exe` or `.vbs` can execute code without consent. Especially from untrusted sources, risks allowing viruses, malware, or hackers to compromise your device security.

**Detector ID**

cpp/unsafe-file-extension@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/cpp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-434](https://cwe.mitre.org/data/definitions/434.html)

**Tags**

[\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags/owasp-top10)

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
