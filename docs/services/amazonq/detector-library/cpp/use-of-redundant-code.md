![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](loose-file-permissions.md) [Sensitive information leak](sensitive-information-leak.md) [Missing Authorization](missing-authorization.md) [Return Stack Address](return-stack-address.md) [OS Command Injection](os-command-injection.md) [Use After Free](use-after-free.md) [Incorrect Comparison](incorrect-comparison.md) [off by one error](off-by-one-error.md) [Path traversal](path-traversal.md) [Insecure temporary file or directory](insecure-temporary-file-or-directory.md) [Insecure Cryptography](insecure-cryptography.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Unchecked Null Dereference](unchecked-null-dereference.md) [SQL injection](sql-injection.md) [Missing check on method output](missing-check-on-method-output.md) [Improper Restriction on Memory Buffer](improper-restriction-on-memory-buffer.md) [Multiple Locks](multiple-locks.md) [Improper Input Validation](improper-input-validation.md) [Null Pointer Dereference](null-pointer-dereference.md) [Use Of Redundant Code](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-of-redundant-code) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-certificate-validation) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-authentication)

# Use Of Redundant Code [High](https://docs.aws.amazon.com/amazonq/detector-library/cpp/severity/high)

The use of redundant code encompasses instances where code contains unnecessary, redundant, or superfluous elements that serve no practical purpose, potentially leading to confusion, increased code size, and maintenance overhead. To address this vulnerability, it is essential to regularly review and refactor the codebase to remove redundant elements, ensuring clarity, simplicity, and efficiency in the software design.

**Detector ID**

cpp/use-of-redundant-code@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/cpp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-1041](https://cwe.mitre.org/data/definitions/1041.html)

**Tags**

-

* * *

#### Noncompliant example

```cpp
1#include <stdlib.h>
2
3int useOfRedundantCodeNoncompliant(bool b) {
4    if (b)
5    {
6        Arr a;
7        // Noncompliant: Unnecessary copy operation because it creates a copy of a conditionally based on the value of `b`.
8        auto cpy = a;
9    }
10}

```

#### Compliant example

```cpp
1#include <stdlib.h>
2
3int useOfRedundantCodeCompliant(bool b) {
4    if (b)
5    {
6        Arr a;
7        auto cpy = a;
8        // Compliant: Modification ensures that the copy operation is necessary, as the copied instance `cpy` is used independently of the original instance `a`.
9        cpy.arr[0] = 8;
10    }
11}

```
