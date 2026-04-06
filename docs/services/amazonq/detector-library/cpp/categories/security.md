![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](../do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](../weak-random-number-generation.md) [Missing Default in Switch](../missing-default-in-switch.md) [Unsafe File Extension](../unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](../incorrect-order-setuid-setgid.md) [Out Of Bounds Read](../out-of-bounds-read.md) [Out Of Bounds Write](../out-of-bounds-write.md) [Thread safety violation](../thread-safety-violation.md) [Incorrect Pointer Subtraction](../pointer-subtraction.md) [File System Access](../file-system-access.md) [Insecure Buffer Access](../insecure-buffer-access.md) [Incorrect Use of Sizeof](../incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](../incorrect-pointer-scaling.md) [Loose File Permissions](../loose-file-permissions.md) [Sensitive information leak](../sensitive-information-leak.md) [Missing Authorization](../missing-authorization.md) [Return Stack Address](../return-stack-address.md) [OS Command Injection](../os-command-injection.md) [Use After Free](../use-after-free.md) [Incorrect Comparison](../incorrect-comparison.md) [off by one error](../off-by-one-error.md) [Path traversal](../path-traversal.md) [Insecure temporary file or directory](../insecure-temporary-file-or-directory.md) [Insecure Cryptography](../insecure-cryptography.md) [Insecure connection using unencrypted protocol](../insecure-connection.md) [Unchecked Null Dereference](../unchecked-null-dereference.md) [SQL injection](../sql-injection.md) [Missing check on method output](../missing-check-on-method-output.md) [Improper Restriction on Memory Buffer](../improper-restriction-on-memory-buffer.md) [Multiple Locks](../multiple-locks.md) [Improper Input Validation](../improper-input-validation.md) [Null Pointer Dereference](../null-pointer-dereference.md) [Use Of Redundant Code](../use-of-redundant-code.md) [Improper Certificate Validation](../improper-certificate-validation.md) [Improper Authentication](../improper-authentication.md)

# Security detectors

### [Disabled HTML autoescape](../do-not-disable-html-autoescape.md)

Disabling the HTML autoescape mechanism exposes your web applications to attacks.

### [Weak pseudorandom number generation](../weak-random-number-generation.md)

Insufficiently random generators (or hardcoded seeds) can make pseudorandom sequences predictable.

### [Missing Default in Switch](../missing-default-in-switch.md)

Missing `default` in `switch` can lead to unintended bugs.

### [Unsafe File Extension](../unsafe-file-extension.md)

Unsafe file extensions like `.exe` or `.vbs` can execute code without consent.

### [Incorrect Order Of setuid and setgid](../incorrect-order-setuid-setgid.md)

if set(e)gid() is called after set(e)uid(), it can regain elevated group privileges.

### [Out Of Bounds Read](../out-of-bounds-read.md)

Out of bounds read can allow attackers to read sensitive information from other memory locations or cause a crash.

### [Out Of Bounds Write](../out-of-bounds-write.md)

Out-of-bounds write vulnerability occurs when software attempts to write data beyond the allocated memory bounds, potentially leading to memory corruption and security risks.

### [Thread safety violation](../thread-safety-violation.md)

A thread safety violation might indicate a data race which can put the system into an inconsistent state.

### [Incorrect Pointer Subtraction](../pointer-subtraction.md)

Pointer subtraction allows unintended behavior.

### [File System Access](../file-system-access.md)

Concurrent execution using shared resource with improper synchronization.

### [Insecure Buffer Access](../insecure-buffer-access.md)

Use of insecure functions can lead to buffer overflow.

### [Incorrect Use of Sizeof](../incorrect-use-of-sizeof.md)

Use of sizeof on a malloced pointer type is incorrect.

### [Incorrect Pointer Scaling](../incorrect-pointer-scaling.md)

Instances of incorrect pointer scaling detected, potentially leading to unexpected behavior and vulnerabilities.

### [Loose File Permissions](../loose-file-permissions.md)

Weak file permissions can lead to privilege escalation.

### [Sensitive information leak](../sensitive-information-leak.md)

Sensitive information should not be exposed through log files or stack traces.

### [Missing Authorization](../missing-authorization.md)

Missing authorization checks can lead to unauthorized access to a resource or performance of an action.

### [Return Stack Address](../return-stack-address.md)

A function returns the address of a stack variable will cause unintended program behavior, typically in the form of a crash.

### [OS Command Injection](../os-command-injection.md)

Improper neutralization of special elements used in an OS command ('OS Command Injection')

### [Use After Free](../use-after-free.md)

Using memory after it has been freed can lead to unexpected behavior or exploitation.

### [Incorrect Comparison](../incorrect-comparison.md)

Secure comparison type safety recommendation.

### [off by one error](../off-by-one-error.md)

Off-by-one errors occur when a loop or array index is incorrectly incremented or decremented by one, leading to unintended behavior.

### [Path traversal](../path-traversal.md)

Creating file paths from untrusted input might give a malicious actor access to sensitive files.

### [Insecure temporary file or directory](../insecure-temporary-file-or-directory.md)

Insecure ways of creating temporary files and directories can lead to race conditions, privilege escalation, and other security vulnerabilities.

### [Insecure Cryptography](../insecure-cryptography.md)

Use of insecure cryptography.

### [Insecure connection using unencrypted protocol](../insecure-connection.md)

Connections that use insecure protocols transmit data in cleartext, which can leak sensitive information.

### [Unchecked Null Dereference](../unchecked-null-dereference.md)

Dereferencing an unchecked value without verification, leading to potential runtime errors or undefined behavior.

### [SQL injection](../sql-injection.md)

Use of untrusted inputs in SQL database query can enable attackers to read, modify, or delete sensitive data in the database.

### [Missing check on method output](../missing-check-on-method-output.md)

Missing checks might cause silent failures that are harder to debug.

### [Improper Restriction on Memory Buffer](../improper-restriction-on-memory-buffer.md)

Improper restriction of operations within the bounds of a memory buffer.

### [Multiple Locks](../multiple-locks.md)

Repeatedly locking critical resources in a concurrent environment can lead to varied consequences.

### [Improper Input Validation](../improper-input-validation.md)

Improper input validation can enable attacks and lead to unwanted behavior.

### [Null Pointer Dereference](../null-pointer-dereference.md)

Dereferencing a null pointer can lead to unexpected null pointer exceptions.

### [Use Of Redundant Code](../use-of-redundant-code.md)

Redundant data copying operations in the code lead to performance and memory inefficiencies.

### [Improper Certificate Validation](../improper-certificate-validation.md)

Improper certificate validation might allow an attacker to spoof a trusted entity by interfering in the communication path between the host and client.

### [Improper Authentication](../improper-authentication.md)

Improper authentication from insufficient identity verification.
