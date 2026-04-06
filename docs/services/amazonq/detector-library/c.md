![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](https://docs.aws.amazon.com/amazonq/detector-library/c/logging-of-sensitive-information) [Insecure Use Of Chroot](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-of-chroot) [Deadlock And Lock Inconsistency](https://docs.aws.amazon.com/amazonq/detector-library/c/deadlock-and-lock-inconsistency) [Unsafe File Extension](https://docs.aws.amazon.com/amazonq/detector-library/c/unsafe-file-extension) [OS command injection](https://docs.aws.amazon.com/amazonq/detector-library/c/os-command-injection) [Incorrect Use Of Free](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-of-free) [Use Of Uninitialized Variable](https://docs.aws.amazon.com/amazonq/detector-library/c/use-of-uninitialized-variable) [Insecure Use strcat fn](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strcat-fn) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/c/sql-injection) [Bitwise Operator On Signed Operand](https://docs.aws.amazon.com/amazonq/detector-library/c/bitwise-operator-on-signed-operand) [Insecure use gets fn](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-gets-fn) [Random fd exhaustion](https://docs.aws.amazon.com/amazonq/detector-library/c/random-fd-exhaustion) [Redundant Free Usage](https://docs.aws.amazon.com/amazonq/detector-library/c/redundant-free-usage) [Insecure Use Memset](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-memset) [Divide By Zero.](https://docs.aws.amazon.com/amazonq/detector-library/c/divide-by-zero) [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/c/return-stack-address) [Unchecked Return Value](https://docs.aws.amazon.com/amazonq/detector-library/c/unchecked-return-value) [Incorrect Format Specifier](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-format-specifier) [Unhandled Expression Result](https://docs.aws.amazon.com/amazonq/detector-library/c/unhandled-expression-result) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/c/path-traversal) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-input-validation) [Out Of Bounds Read](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-read) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/c/integer-overflow) [Insecure use strtok function](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strtok-fn) [Improper size of a memory buffer](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-size-of-a-memory-buffer) [incomplete-cleanup](https://docs.aws.amazon.com/amazonq/detector-library/c/incomplete-cleanup) [Null pointer dereference](https://docs.aws.amazon.com/amazonq/detector-library/c/null-pointer-dereference) [Insecure Temporary File Or Directory](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-temporary-file-or-directory) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-buffer-access) [Incorrect Use Ato Fn](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-ato-fn) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/c/loose-file-permissions) [Exposure of Sensitive Information](https://docs.aws.amazon.com/amazonq/detector-library/c/exposure-of-sensitive-information) [Out-of-bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-write) [String Equality](https://docs.aws.amazon.com/amazonq/detector-library/c/string-equality)

# C detectors

Showing all detectors for the C language.

##### Browse by tags

Browse all detectors by tags.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/c/tags)

##### Browse by severity

Browse all detectors by severity.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/c/severity)

##### Browse by category

Browse all detectors by category.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/c/categories)

* * *

### Browse all detectors

### [Logging of sensitive information](https://docs.aws.amazon.com/amazonq/detector-library/c/logging-of-sensitive-information)

Sensitive information has been logged in your code which may leads to sensitive information leak.

### [Insecure Use Of Chroot](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-of-chroot)

Improper use of `chroot` function may allow attackers to escape the chroot jail due to relative paths still referencing resources outside the intended jail after `chroot` function is called.

### [Deadlock And Lock Inconsistency](https://docs.aws.amazon.com/amazonq/detector-library/c/deadlock-and-lock-inconsistency)

This code contains a potential deadlock and violates lock consistency due to incorrect lock ordering or nested locking.

### [Unsafe File Extension](https://docs.aws.amazon.com/amazonq/detector-library/c/unsafe-file-extension)

Insufficiently restrictive file uploads can lead to inadvertently running malicious code.

### [OS command injection](https://docs.aws.amazon.com/amazonq/detector-library/c/os-command-injection)

Constructing operating system or shell commands with unsanitized user input can lead to inadvertently running malicious code.

### [Incorrect Use Of Free](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-of-free)

Using memory after it has been freed can lead to unexpected behavior or exploitation.

### [Use Of Uninitialized Variable](https://docs.aws.amazon.com/amazonq/detector-library/c/use-of-uninitialized-variable)

The code uses a variable that has not been initialized, leading to unpredictable or unintended results.

### [Insecure Use strcat fn](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strcat-fn)

strcat or strncat can lead to buffer overflow vulnerabilities.

### [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/c/sql-injection)

The use of untrusted inputs in a SQL database query can enable attackers to read, modify, or delete sensitive data in the database.

### [Bitwise Operator On Signed Operand](https://docs.aws.amazon.com/amazonq/detector-library/c/bitwise-operator-on-signed-operand)

Bitwise operator applied on signed operand.

### [Insecure use gets fn](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-gets-fn)

gets can lead to buffer overflow vulnerabilities.

### [Random fd exhaustion](https://docs.aws.amazon.com/amazonq/detector-library/c/random-fd-exhaustion)

Failure to limit and close open file descriptors allows uncontrolled resource consumption which can crash programs or degrade system performance.

### [Redundant Free Usage](https://docs.aws.amazon.com/amazonq/detector-library/c/redundant-free-usage)

Calling free method twice can be vulnerable to memory location.

### [Insecure Use Memset](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-memset)

Calling memset method can leave sensitive information behind.

### [Divide By Zero.](https://docs.aws.amazon.com/amazonq/detector-library/c/divide-by-zero)

Software flaws related to dividing by zero.

### [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/c/return-stack-address)

A function returns the address of a stack variable will cause unintended program behavior, typically in the form of a crash.

### [Unchecked Return Value](https://docs.aws.amazon.com/amazonq/detector-library/c/unchecked-return-value)

Check the return value from a method or function.

### [Incorrect Format Specifier](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-format-specifier)

This code contains a potential format vulnerability due to the use of function without specifying correct format specifier.

### [Unhandled Expression Result](https://docs.aws.amazon.com/amazonq/detector-library/c/unhandled-expression-result)

Encourage Purposeful Operator Usage in Codebase.

### [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/c/path-traversal)

Creating file paths from untrusted input might give a malicious actor access to sensitive files.

### [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-input-validation)

Improper input validation can enable attacks and lead to unwanted behavior.

### [Out Of Bounds Read](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-read)

Out of bounds read can allow attackers to read sensitive information from other memory locations or cause a crash.

### [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/c/integer-overflow)

An integer overflow might might cause security issues when it is used for resource management or execution control.

### [Insecure use strtok function](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strtok-fn)

strtok() can cause unintended consequences and security issues.

### [Improper size of a memory buffer](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-size-of-a-memory-buffer)

The product performs operations on a memory buffer, but it can read from or write to a memory location that is outside of the intended boundary of the buffer.

### [incomplete-cleanup](https://docs.aws.amazon.com/amazonq/detector-library/c/incomplete-cleanup)

This concept often emphasizes identifying instances in code where resources, like file descriptors, aren't properly `released` or `closed`, particularly in C programming.

### [Null pointer dereference](https://docs.aws.amazon.com/amazonq/detector-library/c/null-pointer-dereference)

We observed that a NULL pointer dereference occurs when a program attempts to access the value referenced by a pointer that is currently set to NULL.

### [Insecure Temporary File Or Directory](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-temporary-file-or-directory)

Securely create temporary files using functions like mkstemp() or tmpfile(), ensuring proper permissions with open() or fopen() during creation or via chmod() afterward.

### [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-buffer-access)

Use of insecure function can lead to buffer overflows.

### [Incorrect Use Ato Fn](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-ato-fn)

Use `strtol()` instead of `atoi()` for string to number conversions to avoid undefined behavior from invalid inputs.

### [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/c/loose-file-permissions)

Providing permissions for a security-critical resource in a way that allows that resource to be read or modified by unintended actors.

### [Exposure of Sensitive Information](https://docs.aws.amazon.com/amazonq/detector-library/c/exposure-of-sensitive-information)

This code transmits sensitive information over a network or communication channel in cleartext, making it vulnerable to interception by attackers.

### [Out-of-bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-write)

Out of bounds write can allow attackers to write sensitive information from other memory locations or cause a crash.

### [String Equality](https://docs.aws.amazon.com/amazonq/detector-library/c/string-equality)

Use `strcmp()` or `strncmp()` instead of `==` and `!=` for character content comparison.
