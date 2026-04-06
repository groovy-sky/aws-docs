![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/cpp/do-not-disable-html-autoescape) [Weak pseudorandom number generation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/weak-random-number-generation) [Missing Default in Switch](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-default-in-switch) [Unsafe File Extension](https://docs.aws.amazon.com/amazonq/detector-library/cpp/unsafe-file-extension) [Incorrect Order Of setuid and setgid](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-order-setuid-setgid) [Out Of Bounds Read](https://docs.aws.amazon.com/amazonq/detector-library/cpp/out-of-bounds-read) [Out Of Bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/cpp/out-of-bounds-write) [Thread safety violation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/thread-safety-violation) [Incorrect Pointer Subtraction](https://docs.aws.amazon.com/amazonq/detector-library/cpp/pointer-subtraction) [File System Access](https://docs.aws.amazon.com/amazonq/detector-library/cpp/file-system-access) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-buffer-access) [Incorrect Use of Sizeof](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-use-of-sizeof) [Incorrect Pointer Scaling](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-pointer-scaling) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/cpp/loose-file-permissions) [Sensitive information leak](https://docs.aws.amazon.com/amazonq/detector-library/cpp/sensitive-information-leak) [Missing Authorization](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-authorization) [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/cpp/return-stack-address) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/os-command-injection) [Use After Free](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-after-free) [Incorrect Comparison](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-comparison) [off by one error](https://docs.aws.amazon.com/amazonq/detector-library/cpp/off-by-one-error) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/cpp/path-traversal) [Insecure temporary file or directory](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-temporary-file-or-directory) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-cryptography) [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-connection) [Unchecked Null Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/unchecked-null-dereference) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/sql-injection) [Missing check on method output](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-check-on-method-output) [Improper Restriction on Memory Buffer](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-restriction-on-memory-buffer) [Multiple Locks](https://docs.aws.amazon.com/amazonq/detector-library/cpp/multiple-locks) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-input-validation) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/null-pointer-dereference) [Use Of Redundant Code](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-of-redundant-code) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-certificate-validation) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-authentication)

# C++ detectors

Showing all detectors for the C++ language.

##### Browse by tags

Browse all detectors by tags.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags)

##### Browse by severity

Browse all detectors by severity.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/cpp/severity)

##### Browse by category

Browse all detectors by category.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/cpp/categories)

* * *

### Browse all detectors

### [Disabled HTML autoescape](https://docs.aws.amazon.com/amazonq/detector-library/cpp/do-not-disable-html-autoescape)

Disabling the HTML autoescape mechanism exposes your web applications to attacks.

### [Weak pseudorandom number generation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/weak-random-number-generation)

Insufficiently random generators (or hardcoded seeds) can make pseudorandom sequences predictable.

### [Missing Default in Switch](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-default-in-switch)

Missing `default` in `switch` can lead to unintended bugs.

### [Unsafe File Extension](https://docs.aws.amazon.com/amazonq/detector-library/cpp/unsafe-file-extension)

Unsafe file extensions like `.exe` or `.vbs` can execute code without consent.

### [Incorrect Order Of setuid and setgid](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-order-setuid-setgid)

if set(e)gid() is called after set(e)uid(), it can regain elevated group privileges.

### [Out Of Bounds Read](https://docs.aws.amazon.com/amazonq/detector-library/cpp/out-of-bounds-read)

Out of bounds read can allow attackers to read sensitive information from other memory locations or cause a crash.

### [Out Of Bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/cpp/out-of-bounds-write)

Out-of-bounds write vulnerability occurs when software attempts to write data beyond the allocated memory bounds, potentially leading to memory corruption and security risks.

### [Thread safety violation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/thread-safety-violation)

A thread safety violation might indicate a data race which can put the system into an inconsistent state.

### [Incorrect Pointer Subtraction](https://docs.aws.amazon.com/amazonq/detector-library/cpp/pointer-subtraction)

Pointer subtraction allows unintended behavior.

### [File System Access](https://docs.aws.amazon.com/amazonq/detector-library/cpp/file-system-access)

Concurrent execution using shared resource with improper synchronization.

### [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-buffer-access)

Use of insecure functions can lead to buffer overflow.

### [Incorrect Use of Sizeof](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-use-of-sizeof)

Use of sizeof on a malloced pointer type is incorrect.

### [Incorrect Pointer Scaling](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-pointer-scaling)

Instances of incorrect pointer scaling detected, potentially leading to unexpected behavior and vulnerabilities.

### [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/cpp/loose-file-permissions)

Weak file permissions can lead to privilege escalation.

### [Sensitive information leak](https://docs.aws.amazon.com/amazonq/detector-library/cpp/sensitive-information-leak)

Sensitive information should not be exposed through log files or stack traces.

### [Missing Authorization](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-authorization)

Missing authorization checks can lead to unauthorized access to a resource or performance of an action.

### [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/cpp/return-stack-address)

A function returns the address of a stack variable will cause unintended program behavior, typically in the form of a crash.

### [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/os-command-injection)

Improper neutralization of special elements used in an OS command ('OS Command Injection')

### [Use After Free](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-after-free)

Using memory after it has been freed can lead to unexpected behavior or exploitation.

### [Incorrect Comparison](https://docs.aws.amazon.com/amazonq/detector-library/cpp/incorrect-comparison)

Secure comparison type safety recommendation.

### [off by one error](https://docs.aws.amazon.com/amazonq/detector-library/cpp/off-by-one-error)

Off-by-one errors occur when a loop or array index is incorrectly incremented or decremented by one, leading to unintended behavior.

### [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/cpp/path-traversal)

Creating file paths from untrusted input might give a malicious actor access to sensitive files.

### [Insecure temporary file or directory](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-temporary-file-or-directory)

Insecure ways of creating temporary files and directories can lead to race conditions, privilege escalation, and other security vulnerabilities.

### [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-cryptography)

Use of insecure cryptography.

### [Insecure connection using unencrypted protocol](https://docs.aws.amazon.com/amazonq/detector-library/cpp/insecure-connection)

Connections that use insecure protocols transmit data in cleartext, which can leak sensitive information.

### [Unchecked Null Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/unchecked-null-dereference)

Dereferencing an unchecked value without verification, leading to potential runtime errors or undefined behavior.

### [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/sql-injection)

Use of untrusted inputs in SQL database query can enable attackers to read, modify, or delete sensitive data in the database.

### [Missing check on method output](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-check-on-method-output)

Missing checks might cause silent failures that are harder to debug.

### [Improper Restriction on Memory Buffer](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-restriction-on-memory-buffer)

Improper restriction of operations within the bounds of a memory buffer.

### [Multiple Locks](https://docs.aws.amazon.com/amazonq/detector-library/cpp/multiple-locks)

Repeatedly locking critical resources in a concurrent environment can lead to varied consequences.

### [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-input-validation)

Improper input validation can enable attacks and lead to unwanted behavior.

### [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/null-pointer-dereference)

Dereferencing a null pointer can lead to unexpected null pointer exceptions.

### [Use Of Redundant Code](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-of-redundant-code)

Redundant data copying operations in the code lead to performance and memory inefficiencies.

### [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-certificate-validation)

Improper certificate validation might allow an attacker to spoof a trusted entity by interfering in the communication path between the host and client.

### [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-authentication)

Improper authentication from insufficient identity verification.
