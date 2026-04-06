![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](loose-file-permissions.md) [Sensitive information leak](sensitive-information-leak.md) [Missing Authorization](missing-authorization.md) [Return Stack Address](return-stack-address.md) [OS Command Injection](os-command-injection.md) [Use After Free](use-after-free.md) [Incorrect Comparison](incorrect-comparison.md) [off by one error](off-by-one-error.md) [Path traversal](path-traversal.md) [Insecure temporary file or directory](insecure-temporary-file-or-directory.md) [Insecure Cryptography](insecure-cryptography.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Unchecked Null Dereference](unchecked-null-dereference.md) [SQL injection](sql-injection.md) [Missing check on method output](missing-check-on-method-output.md) [Improper Restriction on Memory Buffer](improper-restriction-on-memory-buffer.md) [Multiple Locks](https://docs.aws.amazon.com/amazonq/detector-library/cpp/multiple-locks) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-input-validation) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/null-pointer-dereference) [Use Of Redundant Code](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-of-redundant-code) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-certificate-validation) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-authentication)

# Multiple Locks [High](https://docs.aws.amazon.com/amazonq/detector-library/cpp/severity/high)

In a concurrent environment, repeated locking of critical resources yields varied consequences. For instance, with pooled resources like semaphores, excessive locking calls can degrade performance or trigger denial of service. Conversely, binary locks may stall progress by rejecting duplicate lock attempts. Understanding these implications is crucial for mitigating performance issues and security vulnerabilities.

**Detector ID**

cpp/multiple-locks@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/cpp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-764](https://cwe.mitre.org/data/definitions/764.html)

**Tags**

-

* * *

#### Noncompliant example

```cpp
1#include <iostream>
2#include <mutex>
3using namespace std;
4
5std::mutex mtx1;
6std::mutex mtx2;
7
8void multipleLocksNoncompliant() {
9    // Noncompliant: Attempts to acquire locks on mtx1 and mtx2 sequentially, which may lead to a deadlock scenario.
10    mtx1.lock();
11    mtx2.lock();
12    printf("f1");
13    mtx2.unlock();
14    mtx1.unlock();
15}

```

#### Compliant example

```cpp
1#include <iostream>
2#include <mutex>
3using namespace std;
4
5std::mutex mtx1;
6std::mutex mtx2;
7
8void multipleLocksCompliant() {
9    // Compliant: Two mutexes mtx1 and mtx2 are acquired and released sequentially, ensuring that they are locked and unlocked correctly.
10    mtx1.lock();
11    mtx1.unlock();
12    mtx2.lock();
13    printf("f1");
14    mtx2.unlock();
15}

```
