![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](../do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](../weak-random-number-generation.md) [Missing Default in Switch](../missing-default-in-switch.md) [Unsafe File Extension](../unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](../incorrect-order-setuid-setgid.md) [Out Of Bounds Read](../out-of-bounds-read.md) [Out Of Bounds Write](../out-of-bounds-write.md) [Thread safety violation](../thread-safety-violation.md) [Incorrect Pointer Subtraction](../pointer-subtraction.md) [File System Access](../file-system-access.md) [Insecure Buffer Access](../insecure-buffer-access.md) [Incorrect Use of Sizeof](../incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](../incorrect-pointer-scaling.md) [Loose File Permissions](../loose-file-permissions.md) [Sensitive information leak](../sensitive-information-leak.md) [Missing Authorization](../missing-authorization.md) [Return Stack Address](../return-stack-address.md) [OS Command Injection](../os-command-injection.md) [Use After Free](../use-after-free.md) [Incorrect Comparison](../incorrect-comparison.md) [off by one error](../off-by-one-error.md) [Path traversal](../path-traversal.md) [Insecure temporary file or directory](../insecure-temporary-file-or-directory.md) [Insecure Cryptography](../insecure-cryptography.md) [Insecure connection using unencrypted protocol](../insecure-connection.md) [Unchecked Null Dereference](../unchecked-null-dereference.md) [SQL injection](../sql-injection.md) [Missing check on method output](../missing-check-on-method-output.md) [Improper Restriction on Memory Buffer](../improper-restriction-on-memory-buffer.md) [Multiple Locks](../multiple-locks.md) [Improper Input Validation](../improper-input-validation.md) [Null Pointer Dereference](../null-pointer-dereference.md) [Use Of Redundant Code](../use-of-redundant-code.md) [Improper Certificate Validation](../improper-certificate-validation.md) [Improper Authentication](../improper-authentication.md)

# Tag: top25-cwes

### [Out Of Bounds Read](../out-of-bounds-read.md)

Out of bounds read can allow attackers to read sensitive information from other memory locations or cause a crash.

### [Out Of Bounds Write](../out-of-bounds-write.md)

Out-of-bounds write vulnerability occurs when software attempts to write data beyond the allocated memory bounds, potentially leading to memory corruption and security risks.

### [Incorrect Pointer Subtraction](../pointer-subtraction.md)

Pointer subtraction allows unintended behavior.

### [File System Access](../file-system-access.md)

Concurrent execution using shared resource with improper synchronization.

### [Incorrect Use of Sizeof](../incorrect-use-of-sizeof.md)

Use of sizeof on a malloced pointer type is incorrect.

### [Loose File Permissions](../loose-file-permissions.md)

Weak file permissions can lead to privilege escalation.

### [Sensitive information leak](../sensitive-information-leak.md)

Sensitive information should not be exposed through log files or stack traces.

### [Return Stack Address](../return-stack-address.md)

A function returns the address of a stack variable will cause unintended program behavior, typically in the form of a crash.

### [Use After Free](../use-after-free.md)

Using memory after it has been freed can lead to unexpected behavior or exploitation.

### [Path traversal](../path-traversal.md)

Creating file paths from untrusted input might give a malicious actor access to sensitive files.

### [SQL injection](../sql-injection.md)

Use of untrusted inputs in SQL database query can enable attackers to read, modify, or delete sensitive data in the database.
