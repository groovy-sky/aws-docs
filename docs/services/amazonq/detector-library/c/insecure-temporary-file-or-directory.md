![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](divide-by-zero.md) [Return Stack Address](return-stack-address.md) [Unchecked Return Value](unchecked-return-value.md) [Incorrect Format Specifier](incorrect-format-specifier.md) [Unhandled Expression Result](unhandled-expression-result.md) [Path traversal](path-traversal.md) [Improper Input Validation](improper-input-validation.md) [Out Of Bounds Read](out-of-bounds-read.md) [Integer Overflow](integer-overflow.md) [Insecure use strtok function](insecure-use-strtok-fn.md) [Improper size of a memory buffer](improper-size-of-a-memory-buffer.md) [incomplete-cleanup](incomplete-cleanup.md) [Null pointer dereference](null-pointer-dereference.md) [Insecure Temporary File Or Directory](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-temporary-file-or-directory) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-buffer-access) [Incorrect Use Ato Fn](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-ato-fn) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/c/loose-file-permissions) [Exposure of Sensitive Information](https://docs.aws.amazon.com/amazonq/detector-library/c/exposure-of-sensitive-information) [Out-of-bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-write) [String Equality](https://docs.aws.amazon.com/amazonq/detector-library/c/string-equality)

# Insecure Temporary File Or Directory [High](https://docs.aws.amazon.com/amazonq/detector-library/c/severity/high)

For secure creation of temporary files, it is advisable to use functions such as `mkstemp()` or `tmpfile()`, and ensure secure permissions by either setting appropriate file modes during creation with `open()` or `fopen()`, or by using `chmod()` afterward.

**Detector ID**

c/insecure-temporary-file-or-directory@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/c/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-377](https://cwe.mitre.org/data/definitions/377.html)

**Tags**

-

* * *

#### Noncompliant example

```c
1#include <stdio.h>
2#include <stdlib.h>
3#include <fcntl.h>
4#include <sys/stat.h>
5#include <string.h>
6#include <unistd.h>
7
8int insecureTemporaryFileorDirectoryNonCompliant(char *tempData) {
9  // Noncompliant: Insecure function used
10  char *path = tmpnam(NULL);
11  FILE* f = fopen(path, "w");
12  fputs(tempData, f);
13  fclose(f);
14}

```

#### Compliant example

```c
1#include <stdio.h>
2#include <stdlib.h>
3#include <fcntl.h>
4#include <sys/stat.h>
5#include <string.h>
6#include <unistd.h>
7
8int insecureTemporaryFileorDirectoryCompliant(char *tempData) {
9  // Compliant: The file will be opened in "wb+" mode, and will be automatically removed on normal program exit
10  FILE* f = tmpfile();
11  fputs(tempData, f);
12  fclose(f);
13  return 0;
14}

```
