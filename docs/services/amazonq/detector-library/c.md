---
title: "C detectors"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](c/logging-of-sensitive-information.md) [Insecure Use Of Chroot](c/insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](c/deadlock-and-lock-inconsistency.md) [Unsafe File Extension](c/unsafe-file-extension.md) [OS command injection](c/os-command-injection.md) [Incorrect Use Of Free](c/incorrect-use-of-free.md) [Use Of Uninitialized Variable](c/use-of-uninitialized-variable.md) [Insecure Use strcat fn](c/insecure-use-strcat-fn.md) [SQL injection](c/sql-injection.md) [Bitwise Operator On Signed Operand](c/bitwise-operator-on-signed-operand.md) [Insecure use gets fn](c/insecure-use-gets-fn.md) [Random fd exhaustion](c/random-fd-exhaustion.md) [Redundant Free Usage](c/redundant-free-usage.md) [Insecure Use Memset](c/insecure-use-memset.md) [Divide By Zero.](c/divide-by-zero.md) [Return Stack Address](c/return-stack-address.md) [Unchecked Return Value](c/unchecked-return-value.md) [Incorrect Format Specifier](c/incorrect-format-specifier.md) [Unhandled Expression Result](c/unhandled-expression-result.md) [Path traversal](c/path-traversal.md) [Improper Input Validation](c/improper-input-validation.md) [Out Of Bounds Read](c/out-of-bounds-read.md) [Integer Overflow](c/integer-overflow.md) [Insecure use strtok function](c/insecure-use-strtok-fn.md) [Improper size of a memory buffer](c/improper-size-of-a-memory-buffer.md) [incomplete-cleanup](c/incomplete-cleanup.md) [Null pointer dereference](c/null-pointer-dereference.md) [Insecure Temporary File Or Directory](c/insecure-temporary-file-or-directory.md) [Insecure Buffer Access](c/insecure-buffer-access.md) [Incorrect Use Ato Fn](c/incorrect-use-ato-fn.md) [Loose File Permissions](c/loose-file-permissions.md) [Exposure of Sensitive Information](c/exposure-of-sensitive-information.md) [Out-of-bounds Write](c/out-of-bounds-write.md) [String Equality](c/string-equality.md)

# C detectors

Showing all detectors for the C language.

##### Browse by tags

Browse all detectors by tags.

[Click here→](c/tags.md)

##### Browse by severity

Browse all detectors by severity.

[Click here→](c/severity.md)

##### Browse by category

Browse all detectors by category.

[Click here→](c/categories.md)

* * *

### Browse all detectors

### [Logging of sensitive information](c/logging-of-sensitive-information.md)

Sensitive information has been logged in your code which may leads to sensitive information leak.

### [Insecure Use Of Chroot](c/insecure-use-of-chroot.md)

Improper use of `chroot` function may allow attackers to escape the chroot jail due to relative paths still referencing resources outside the intended jail after `chroot` function is called.

### [Deadlock And Lock Inconsistency](c/deadlock-and-lock-inconsistency.md)

This code contains a potential deadlock and violates lock consistency due to incorrect lock ordering or nested locking.

### [Unsafe File Extension](c/unsafe-file-extension.md)

Insufficiently restrictive file uploads can lead to inadvertently running malicious code.

### [OS command injection](c/os-command-injection.md)

Constructing operating system or shell commands with unsanitized user input can lead to inadvertently running malicious code.

### [Incorrect Use Of Free](c/incorrect-use-of-free.md)

Using memory after it has been freed can lead to unexpected behavior or exploitation.

### [Use Of Uninitialized Variable](c/use-of-uninitialized-variable.md)

The code uses a variable that has not been initialized, leading to unpredictable or unintended results.

### [Insecure Use strcat fn](c/insecure-use-strcat-fn.md)

strcat or strncat can lead to buffer overflow vulnerabilities.

### [SQL injection](c/sql-injection.md)

The use of untrusted inputs in a SQL database query can enable attackers to read, modify, or delete sensitive data in the database.

### [Bitwise Operator On Signed Operand](c/bitwise-operator-on-signed-operand.md)

Bitwise operator applied on signed operand.

### [Insecure use gets fn](c/insecure-use-gets-fn.md)

gets can lead to buffer overflow vulnerabilities.

### [Random fd exhaustion](c/random-fd-exhaustion.md)

Failure to limit and close open file descriptors allows uncontrolled resource consumption which can crash programs or degrade system performance.

### [Redundant Free Usage](c/redundant-free-usage.md)

Calling free method twice can be vulnerable to memory location.

### [Insecure Use Memset](c/insecure-use-memset.md)

Calling memset method can leave sensitive information behind.

### [Divide By Zero.](c/divide-by-zero.md)

Software flaws related to dividing by zero.

### [Return Stack Address](c/return-stack-address.md)

A function returns the address of a stack variable will cause unintended program behavior, typically in the form of a crash.

### [Unchecked Return Value](c/unchecked-return-value.md)

Check the return value from a method or function.

### [Incorrect Format Specifier](c/incorrect-format-specifier.md)

This code contains a potential format vulnerability due to the use of function without specifying correct format specifier.

### [Unhandled Expression Result](c/unhandled-expression-result.md)

Encourage Purposeful Operator Usage in Codebase.

### [Path traversal](c/path-traversal.md)

Creating file paths from untrusted input might give a malicious actor access to sensitive files.

### [Improper Input Validation](c/improper-input-validation.md)

Improper input validation can enable attacks and lead to unwanted behavior.

### [Out Of Bounds Read](c/out-of-bounds-read.md)

Out of bounds read can allow attackers to read sensitive information from other memory locations or cause a crash.

### [Integer Overflow](c/integer-overflow.md)

An integer overflow might might cause security issues when it is used for resource management or execution control.

### [Insecure use strtok function](c/insecure-use-strtok-fn.md)

strtok() can cause unintended consequences and security issues.

### [Improper size of a memory buffer](c/improper-size-of-a-memory-buffer.md)

The product performs operations on a memory buffer, but it can read from or write to a memory location that is outside of the intended boundary of the buffer.

### [incomplete-cleanup](c/incomplete-cleanup.md)

This concept often emphasizes identifying instances in code where resources, like file descriptors, aren't properly `released` or `closed`, particularly in C programming.

### [Null pointer dereference](c/null-pointer-dereference.md)

We observed that a NULL pointer dereference occurs when a program attempts to access the value referenced by a pointer that is currently set to NULL.

### [Insecure Temporary File Or Directory](c/insecure-temporary-file-or-directory.md)

Securely create temporary files using functions like mkstemp() or tmpfile(), ensuring proper permissions with open() or fopen() during creation or via chmod() afterward.

### [Insecure Buffer Access](c/insecure-buffer-access.md)

Use of insecure function can lead to buffer overflows.

### [Incorrect Use Ato Fn](c/incorrect-use-ato-fn.md)

Use `strtol()` instead of `atoi()` for string to number conversions to avoid undefined behavior from invalid inputs.

### [Loose File Permissions](c/loose-file-permissions.md)

Providing permissions for a security-critical resource in a way that allows that resource to be read or modified by unintended actors.

### [Exposure of Sensitive Information](c/exposure-of-sensitive-information.md)

This code transmits sensitive information over a network or communication channel in cleartext, making it vulnerable to interception by attackers.

### [Out-of-bounds Write](c/out-of-bounds-write.md)

Out of bounds write can allow attackers to write sensitive information from other memory locations or cause a crash.

### [String Equality](c/string-equality.md)

Use `strcmp()` or `strncmp()` instead of `==` and `!=` for character content comparison.

All content copied from https://docs.aws.amazon.com/.
