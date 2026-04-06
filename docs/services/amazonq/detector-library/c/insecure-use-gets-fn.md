![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-gets-fn) [Random fd exhaustion](https://docs.aws.amazon.com/amazonq/detector-library/c/random-fd-exhaustion) [Redundant Free Usage](https://docs.aws.amazon.com/amazonq/detector-library/c/redundant-free-usage) [Insecure Use Memset](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-memset) [Divide By Zero.](https://docs.aws.amazon.com/amazonq/detector-library/c/divide-by-zero) [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/c/return-stack-address) [Unchecked Return Value](https://docs.aws.amazon.com/amazonq/detector-library/c/unchecked-return-value) [Incorrect Format Specifier](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-format-specifier) [Unhandled Expression Result](https://docs.aws.amazon.com/amazonq/detector-library/c/unhandled-expression-result) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/c/path-traversal) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-input-validation) [Out Of Bounds Read](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-read) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/c/integer-overflow) [Insecure use strtok function](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strtok-fn) [Improper size of a memory buffer](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-size-of-a-memory-buffer) [incomplete-cleanup](https://docs.aws.amazon.com/amazonq/detector-library/c/incomplete-cleanup) [Null pointer dereference](https://docs.aws.amazon.com/amazonq/detector-library/c/null-pointer-dereference) [Insecure Temporary File Or Directory](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-temporary-file-or-directory) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-buffer-access) [Incorrect Use Ato Fn](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-ato-fn) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/c/loose-file-permissions) [Exposure of Sensitive Information](https://docs.aws.amazon.com/amazonq/detector-library/c/exposure-of-sensitive-information) [Out-of-bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-write) [String Equality](https://docs.aws.amazon.com/amazonq/detector-library/c/string-equality)

# Insecure use gets fn [High](https://docs.aws.amazon.com/amazonq/detector-library/c/severity/high)

gets does not consider buffer boundaries and can lead to buffer overflows.

**Detector ID**

c/insecure-use-gets-fn@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/c/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-676](https://cwe.mitre.org/data/definitions/676.html)

**Tags**

-

* * *

#### Noncompliant example

```c
1#include <stdio.h>
2#include <string.h>
3
4int DST_BUFFER_SIZE = 120;
5
6int insecureUseGetsFnNonCompliant() {
7    char str[DST_BUFFER_SIZE];
8    // Noncompliant: gets is insecure
9    gets(str);
10    printf("%s", str);
11    return 0;
12}

```

#### Compliant example

```c
1#include <stdio.h>
2#include <string.h>
3
4int insecureUseGetsFnCompliant() {
5    char buf[10];
6    printf("Enter text: ");
7    // Compliant: Secure funcion used with specifying size limit
8    fgets(buf, sizeof(buf), stdin);
9    return 0;
10}

```
