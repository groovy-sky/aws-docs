![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](loose-file-permissions.md) [Sensitive information leak](sensitive-information-leak.md) [Missing Authorization](missing-authorization.md) [Return Stack Address](return-stack-address.md) [OS Command Injection](os-command-injection.md) [Use After Free](use-after-free.md) [Incorrect Comparison](incorrect-comparison.md) [off by one error](off-by-one-error.md) [Path traversal](path-traversal.md) [Insecure temporary file or directory](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-temporary-file-or-directory) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-cryptography) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-connection) [Unchecked Null Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/unchecked-null-dereference) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/sql-injection) [Missing check on method output](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-check-on-method-output) [Improper Restriction on Memory Buffer](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-restriction-on-memory-buffer) [Multiple Locks](https://docs.aws.amazon.com/amazonq/detector-library/cpp/multiple-locks) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-input-validation) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/null-pointer-dereference) [Use Of Redundant Code](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-of-redundant-code) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-certificate-validation) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-authentication)

# Insecure temporary file or directory [Medium](https://docs.aws.amazon.com/amazonq/detector-library/cpp/severity/medium)

Insecure creation of temporary files and directories can introduce race condition vulnerabilities. Attackers can leverage these race conditions to carry out exploits like denial-of-service attacks or escalating their privileges. Proper security practices are required when generating temp files to mitigate these risks.

**Detector ID**

cpp/insecure-temporary-file-or-directory@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/cpp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-377](https://cwe.mitre.org/data/definitions/377.html)

**Tags**

[\# availability](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags/availability) [\# race-condition](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags/race-condition) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags/owasp-top10)

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
