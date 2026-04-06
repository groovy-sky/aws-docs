![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](https://docs.aws.amazon.com/amazonq/detector-library/c/random-fd-exhaustion) [Redundant Free Usage](https://docs.aws.amazon.com/amazonq/detector-library/c/redundant-free-usage) [Insecure Use Memset](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-memset) [Divide By Zero.](https://docs.aws.amazon.com/amazonq/detector-library/c/divide-by-zero) [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/c/return-stack-address) [Unchecked Return Value](https://docs.aws.amazon.com/amazonq/detector-library/c/unchecked-return-value) [Incorrect Format Specifier](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-format-specifier) [Unhandled Expression Result](https://docs.aws.amazon.com/amazonq/detector-library/c/unhandled-expression-result) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/c/path-traversal) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-input-validation) [Out Of Bounds Read](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-read) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/c/integer-overflow) [Insecure use strtok function](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strtok-fn) [Improper size of a memory buffer](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-size-of-a-memory-buffer) [incomplete-cleanup](https://docs.aws.amazon.com/amazonq/detector-library/c/incomplete-cleanup) [Null pointer dereference](https://docs.aws.amazon.com/amazonq/detector-library/c/null-pointer-dereference) [Insecure Temporary File Or Directory](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-temporary-file-or-directory) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-buffer-access) [Incorrect Use Ato Fn](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-ato-fn) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/c/loose-file-permissions) [Exposure of Sensitive Information](https://docs.aws.amazon.com/amazonq/detector-library/c/exposure-of-sensitive-information) [Out-of-bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-write) [String Equality](https://docs.aws.amazon.com/amazonq/detector-library/c/string-equality)

# Random fd exhaustion [High](https://docs.aws.amazon.com/amazonq/detector-library/c/severity/high)

We noticed your failure to limit and close open file descriptors allows uncontrolled resource consumption which can crash programs or degrade system performance by exhausting the operating system's capacity.

**Detector ID**

c/random-fd-exhaustion@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/c/categories/security)

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
