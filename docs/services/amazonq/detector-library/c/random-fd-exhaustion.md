---
title: "Random fd exhaustion High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](divide-by-zero.md) [Return Stack Address](return-stack-address.md) [Unchecked Return Value](unchecked-return-value.md) [Incorrect Format Specifier](incorrect-format-specifier.md) [Unhandled Expression Result](unhandled-expression-result.md) [Path traversal](path-traversal.md) [Improper Input Validation](improper-input-validation.md) [Out Of Bounds Read](out-of-bounds-read.md) [Integer Overflow](integer-overflow.md) [Insecure use strtok function](insecure-use-strtok-fn.md) [Improper size of a memory buffer](improper-size-of-a-memory-buffer.md) [incomplete-cleanup](incomplete-cleanup.md) [Null pointer dereference](null-pointer-dereference.md) [Insecure Temporary File Or Directory](insecure-temporary-file-or-directory.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use Ato Fn](incorrect-use-ato-fn.md) [Loose File Permissions](loose-file-permissions.md) [Exposure of Sensitive Information](exposure-of-sensitive-information.md) [Out-of-bounds Write](out-of-bounds-write.md) [String Equality](string-equality.md)

# Random fd exhaustion [High](severity/high.md)

We noticed your failure to limit and close open file descriptors allows uncontrolled resource consumption which can crash programs or degrade system performance by exhausting the operating system's capacity.

**Detector ID**

c/random-fd-exhaustion@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-400](https://cwe.mitre.org/data/definitions/400.html)

**Tags**

-

* * *

#### Noncompliant example

```c
1
2#include <fcntl.h>
3#include <stdio.h>
4#include <string.h>
5#include <sys/stat.h>
6#include <sys/types.h>
7#include <unistd.h>
8#include <stdlib.h>
9
10int randomFdExhaustionNonCompliant() {
11    int fd;
12    char buf[16];
13    // Noncompliant: Does not handle resource allocation
14    fd = open("/dev/urandom", 0);
15    memset(buf, 0, sizeof(buf));
16    read(fd, buf, sizeof(buf));
17    return 0;
18}

```

#### Compliant example

```c
1#include <fcntl.h>
2#include <stdio.h>
3#include <string.h>
4#include <sys/stat.h>
5#include <sys/types.h>
6#include <unistd.h>
7#include <stdlib.h>
8
9int randomFdExhaustionCompliant() {
10    int fd;
11    int bytes_read;
12    char buf[16];
13    // Compliant: Limits the file descriptor use handling resource allocation
14    fd = open("/dev/urandom", 0);
15    memset(buf, 0, sizeof(buf));
16    bytes_read = read(fd, buf, sizeof(buf));
17    if (bytes_read != sizeof(buf)) {
18        return -1;
19    }
20    return 0;
21}

```

All content copied from https://docs.aws.amazon.com/.
