![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-of-chroot) [Deadlock And Lock Inconsistency](https://docs.aws.amazon.com/amazonq/detector-library/c/deadlock-and-lock-inconsistency) [Unsafe File Extension](https://docs.aws.amazon.com/amazonq/detector-library/c/unsafe-file-extension) [OS command injection](https://docs.aws.amazon.com/amazonq/detector-library/c/os-command-injection) [Incorrect Use Of Free](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-of-free) [Use Of Uninitialized Variable](https://docs.aws.amazon.com/amazonq/detector-library/c/use-of-uninitialized-variable) [Insecure Use strcat fn](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strcat-fn) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/c/sql-injection) [Bitwise Operator On Signed Operand](https://docs.aws.amazon.com/amazonq/detector-library/c/bitwise-operator-on-signed-operand) [Insecure use gets fn](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-gets-fn) [Random fd exhaustion](https://docs.aws.amazon.com/amazonq/detector-library/c/random-fd-exhaustion) [Redundant Free Usage](https://docs.aws.amazon.com/amazonq/detector-library/c/redundant-free-usage) [Insecure Use Memset](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-memset) [Divide By Zero.](https://docs.aws.amazon.com/amazonq/detector-library/c/divide-by-zero) [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/c/return-stack-address) [Unchecked Return Value](https://docs.aws.amazon.com/amazonq/detector-library/c/unchecked-return-value) [Incorrect Format Specifier](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-format-specifier) [Unhandled Expression Result](https://docs.aws.amazon.com/amazonq/detector-library/c/unhandled-expression-result) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/c/path-traversal) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-input-validation) [Out Of Bounds Read](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-read) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/c/integer-overflow) [Insecure use strtok function](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strtok-fn) [Improper size of a memory buffer](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-size-of-a-memory-buffer) [incomplete-cleanup](https://docs.aws.amazon.com/amazonq/detector-library/c/incomplete-cleanup) [Null pointer dereference](https://docs.aws.amazon.com/amazonq/detector-library/c/null-pointer-dereference) [Insecure Temporary File Or Directory](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-temporary-file-or-directory) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-buffer-access) [Incorrect Use Ato Fn](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-ato-fn) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/c/loose-file-permissions) [Exposure of Sensitive Information](https://docs.aws.amazon.com/amazonq/detector-library/c/exposure-of-sensitive-information) [Out-of-bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-write) [String Equality](https://docs.aws.amazon.com/amazonq/detector-library/c/string-equality)

# Insecure Use Of Chroot [High](https://docs.aws.amazon.com/amazonq/detector-library/c/severity/high)

Your code uses `chroot` function insecurely, risking security vulnerabilities. Ensure secure directory paths, proper error handling, and permissions to prevent unauthorized access. Review and enhance code for safer `chroot` usage.

**Detector ID**

c/insecure-use-of-chroot@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/c/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-243](https://cwe.mitre.org/data/definitions/243.html)

**Tags**

-

* * *

#### Noncompliant example

```c
1#include <stdlib.h>
2#include <unistd.h>
3#include <stdio.h>
4
5void insecureUseofChrootNoncompliant(){
6
7    const char* root_dir = "/jail/";
8    // Noncompliant: No chdir before or after chroot, and missing check of return value
9    chroot(root_dir);
10}

```

#### Compliant example

```c
1#include <stdio.h>
2void insecureUseofChrootCompliant(){
3
4    const char* root_dir = "/jail/";
5
6    if(chdir(root_dir) == -1) {
7      exit(-1);
8    }
9    // Compliant: the current dir is changed to the jail and the results of both functions are checked
10    if(chroot(root_dir) == -1) {
11      exit(-1);
12    }
13
14}

```
