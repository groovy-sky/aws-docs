![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-buffer-access) [Incorrect Use of Sizeof](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-use-of-sizeof) [Incorrect Pointer Scaling](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-pointer-scaling) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/cpp/loose-file-permissions) [Sensitive information leak](https://docs.aws.amazon.com/amazonq/detector-library/cpp/sensitive-information-leak) [Missing Authorization](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-authorization) [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/cpp/return-stack-address) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/os-command-injection) [Use After Free](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-after-free) [Incorrect Comparison](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-comparison) [off by one error](https://docs.aws.amazon.com/amazonq/detector-library/cpp/off-by-one-error) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/cpp/path-traversal) [Insecure temporary file or directory](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-temporary-file-or-directory) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-cryptography) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-connection) [Unchecked Null Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/unchecked-null-dereference) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/sql-injection) [Missing check on method output](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-check-on-method-output) [Improper Restriction on Memory Buffer](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-restriction-on-memory-buffer) [Multiple Locks](https://docs.aws.amazon.com/amazonq/detector-library/cpp/multiple-locks) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-input-validation) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/null-pointer-dereference) [Use Of Redundant Code](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-of-redundant-code) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-certificate-validation) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-authentication)

# Insecure Buffer Access [High](https://docs.aws.amazon.com/amazonq/detector-library/cpp/severity/high)

We recommend you to avoid using insecure functions in your code. This functions, when used improperly, does not consider buffer boundaries and can lead to buffer overflows.

**Detector ID**

cpp/insecure-buffer-access@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/cpp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-676](https://cwe.mitre.org/data/definitions/676.html)

**Tags**

-

* * *

#### Noncompliant example

```cpp
1#include <cstring>
2#include <stdio.h>
3
4void insecureBufferAccessNoncompliant(char *string) {
5    char buf[BUFSIZE];
6    size_t length;
7
8    // Noncompliant: `snprintf()`function returns the total length of the string they tried to create.
9    length = snprintf(buf, BUFSIZE, "%s", string);
10}

```

#### Compliant example

```cpp
1#include <cstring>
2#include <stdio.h>
3
4void insecureBufferAccessCompliant(char *string) {
5    char buf[BUFSIZE];
6    size_t length;
7    // Compliant: `snprintf_s` ensures that the formatted string is no longer than the size of the buffer minus one (for the null terminator).
8    length = snprintf_s(buf, BUFSIZE, "%s", string);
9}

```
