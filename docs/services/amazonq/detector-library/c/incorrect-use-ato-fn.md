![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](divide-by-zero.md) [Return Stack Address](return-stack-address.md) [Unchecked Return Value](unchecked-return-value.md) [Incorrect Format Specifier](incorrect-format-specifier.md) [Unhandled Expression Result](unhandled-expression-result.md) [Path traversal](path-traversal.md) [Improper Input Validation](improper-input-validation.md) [Out Of Bounds Read](out-of-bounds-read.md) [Integer Overflow](integer-overflow.md) [Insecure use strtok function](insecure-use-strtok-fn.md) [Improper size of a memory buffer](improper-size-of-a-memory-buffer.md) [incomplete-cleanup](incomplete-cleanup.md) [Null pointer dereference](null-pointer-dereference.md) [Insecure Temporary File Or Directory](insecure-temporary-file-or-directory.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use Ato Fn](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-ato-fn) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/c/loose-file-permissions) [Exposure of Sensitive Information](https://docs.aws.amazon.com/amazonq/detector-library/c/exposure-of-sensitive-information) [Out-of-bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-write) [String Equality](https://docs.aws.amazon.com/amazonq/detector-library/c/string-equality)

# Incorrect Use Ato Fn [Critical](https://docs.aws.amazon.com/amazonq/detector-library/c/severity/critical)

The functions `atoi()`, `atol()` and `atoll()` do not allow specifying the numeric base or checking for conversion errors. If the input string is invalid or out of the supported range, they invoke undefined behavior instead of returning an error. `Strtol()` and its variants `strtoll()` and `strtoull()` are safer alternatives as they allow specifying the base and return a pointer to the next unparsed character. This pointer can be checked against the original string to ensure the entire string was parsed.

**Detector ID**

c/incorrect-use-ato-fn@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/c/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-648](https://cwe.mitre.org/data/definitions/648.html)

**Tags**

-

* * *

#### Noncompliant example

```c
1#include <stdio.h>
2#include <stdlib.h>
3#include <string.h>
4
5const char *buf = "";
6void incorrectUseAtoFnNonCompliant()
7{
8    // Noncompliant: Insecure function used
9    int i = atoi(buf);
10    printf("Converted integer: %d\n", i);
11}

```

#### Compliant example

```c
1#include <stdio.h>
2#include <stdlib.h>
3#include <string.h>
4
5char* endptr;
6const char *buf = "";
7int incorrectUseAtoFnCompliant()
8{
9    // Compliant: secure function used
10    long num = strtol(buf, &endptr, 10);
11    if(endptr == buf) {
12        printf("Conversion failed\n");
13        return 1;
14    }
15    printf("Converted number: %ld\n", num);
16    return 0;
17}

```
