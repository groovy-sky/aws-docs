![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/cpp/out-of-bounds-write) [Thread safety violation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/thread-safety-violation) [Incorrect Pointer Subtraction](https://docs.aws.amazon.com/amazonq/detector-library/cpp/pointer-subtraction) [File System Access](https://docs.aws.amazon.com/amazonq/detector-library/cpp/file-system-access) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-buffer-access) [Incorrect Use of Sizeof](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-use-of-sizeof) [Incorrect Pointer Scaling](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-pointer-scaling) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/cpp/loose-file-permissions) [Sensitive information leak](https://docs.aws.amazon.com/amazonq/detector-library/cpp/sensitive-information-leak) [Missing Authorization](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-authorization) [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/cpp/return-stack-address) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/os-command-injection) [Use After Free](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-after-free) [Incorrect Comparison](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-comparison) [off by one error](https://docs.aws.amazon.com/amazonq/detector-library/cpp/off-by-one-error) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/cpp/path-traversal) [Insecure temporary file or directory](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-temporary-file-or-directory) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-cryptography) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-connection) [Unchecked Null Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/unchecked-null-dereference) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/sql-injection) [Missing check on method output](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-check-on-method-output) [Improper Restriction on Memory Buffer](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-restriction-on-memory-buffer) [Multiple Locks](https://docs.aws.amazon.com/amazonq/detector-library/cpp/multiple-locks) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-input-validation) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/null-pointer-dereference) [Use Of Redundant Code](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-of-redundant-code) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-certificate-validation) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-authentication)

# Out Of Bounds Write [High](https://docs.aws.amazon.com/amazonq/detector-library/cpp/severity/high)

When software attempts to write data beyond the confines of a fixed-size buffer, array, or allocated memory region. To address and mitigate this vulnerability effectively, implementing robust input validation and leveraging memory safety mechanisms are crucial steps.

**Detector ID**

cpp/out-of-bounds-write@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/cpp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-787](https://cwe.mitre.org/data/definitions/787.html)

**Tags**

[\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags/top25-cwes)

* * *

#### Noncompliant example

```cpp
1void outOfBoundWriteNoncompliant()
2{
3    // Declaring an array named id_sequence with a size of 3 integers
4    int id_sequence[3] = {1, 2, 3};
5    // Noncompliant: Attempting to assign a value to the fourth element.
6    id_sequence[4] = 456;
7}

```

#### Compliant example

```cpp
1void outOfBoundWriteCompliant() {
2    int arr[3] = {1, 2, 3};
3    // Compliant: Accessing indices within array bounds
4    for (int i = 0; i < 3; ++i)
5    {
6        arr[i] = i * 2;
7    }
8}

```
