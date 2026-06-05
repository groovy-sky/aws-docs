---
title: "File System Access High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](loose-file-permissions.md) [Sensitive information leak](sensitive-information-leak.md) [Missing Authorization](missing-authorization.md) [Return Stack Address](return-stack-address.md) [OS Command Injection](os-command-injection.md) [Use After Free](use-after-free.md) [Incorrect Comparison](incorrect-comparison.md) [off by one error](off-by-one-error.md) [Path traversal](path-traversal.md) [Insecure temporary file or directory](insecure-temporary-file-or-directory.md) [Insecure Cryptography](insecure-cryptography.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Unchecked Null Dereference](unchecked-null-dereference.md) [SQL injection](sql-injection.md) [Missing check on method output](missing-check-on-method-output.md) [Improper Restriction on Memory Buffer](improper-restriction-on-memory-buffer.md) [Multiple Locks](multiple-locks.md) [Improper Input Validation](improper-input-validation.md) [Null Pointer Dereference](null-pointer-dereference.md) [Use Of Redundant Code](use-of-redundant-code.md) [Improper Certificate Validation](improper-certificate-validation.md) [Improper Authentication](improper-authentication.md)

# File System Access [High](severity/high.md)

When multiple threads or processes attempt simultaneous access to a shared resource without coordination or synchronization, it can lead to race conditions. To mitigate the risk of race conditions and associated issues, proper synchronization mechanisms should be implemented.

**Detector ID**

cpp/file-system-access@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-362](https://cwe.mitre.org/data/definitions/362.html) [CWE-367](https://cwe.mitre.org/data/definitions/367.html)

**Tags**

[\# top25-cwes](tags/top25-cwes.md)

* * *

#### Noncompliant example

```cpp
1#include <fstream>
2#include <mutex>
3#include <cstdio>
4
5
6void fileSystemAccessNoncompliant(char* filename)
7{
8    std::ifstream fileIn(filename);
9    std::string data;
10    fileIn >> data;
11    std::ofstream fileOut(filename, std::ios::app);
12    // Noncompliant: Simultaneous read and write without proper synchronization.
13    fileOut << "New data appended: " << data << std::endl;
14    fileOut.close();
15    fileIn.close();
16}

```

#### Compliant example

```cpp
1#include <fstream>
2#include <mutex>
3#include <cstdio>
4
5void fileSystemAccessCompliant(char* filename, char* content)
6{
7    std::mutex mtx;
8    // Compliant: Using `std::unique_lock` for synchronization.
9    std::unique_lock<std::mutex> lock(mtx);
10    std::ofstream file(filename, std::ios::app);
11    file << content << std::endl;
12    file.close();
13}

```

All content copied from https://docs.aws.amazon.com/.
