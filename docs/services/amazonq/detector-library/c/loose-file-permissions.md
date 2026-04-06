![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](divide-by-zero.md) [Return Stack Address](return-stack-address.md) [Unchecked Return Value](unchecked-return-value.md) [Incorrect Format Specifier](incorrect-format-specifier.md) [Unhandled Expression Result](unhandled-expression-result.md) [Path traversal](path-traversal.md) [Improper Input Validation](improper-input-validation.md) [Out Of Bounds Read](out-of-bounds-read.md) [Integer Overflow](integer-overflow.md) [Insecure use strtok function](insecure-use-strtok-fn.md) [Improper size of a memory buffer](improper-size-of-a-memory-buffer.md) [incomplete-cleanup](incomplete-cleanup.md) [Null pointer dereference](null-pointer-dereference.md) [Insecure Temporary File Or Directory](insecure-temporary-file-or-directory.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use Ato Fn](incorrect-use-ato-fn.md) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/c/loose-file-permissions) [Exposure of Sensitive Information](https://docs.aws.amazon.com/amazonq/detector-library/c/exposure-of-sensitive-information) [Out-of-bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-write) [String Equality](https://docs.aws.amazon.com/amazonq/detector-library/c/string-equality)

# Loose File Permissions [High](https://docs.aws.amazon.com/amazonq/detector-library/c/severity/high)

Detects files with world write or read permissions, highlighting security risks and emphasizing the importance of restricting access for improved data protection.

**Detector ID**

c/loose-file-permissions@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/c/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-732](https://cwe.mitre.org/data/definitions/732.html) [CWE-266](https://cwe.mitre.org/data/definitions/266.html)

**Tags**

-

* * *

#### Noncompliant example

```c
1#include <sys/stat.h>
2#include <fcntl.h>
3
4void looseFilePermissionsNonCompliant(){
5
6  int fd;
7
8  // Noncompliant: The process set 777 permissions to this newly created file
9  open("myfile.txt", O_CREAT, S_IRWXU | S_IRWXG | S_IRWXO);
10
11  // Noncompliant: The process try to set 777 permissions to this newly created directory
12  mkdir("myfolder", S_IRWXU | S_IRWXG | S_IRWXO);
13
14  // Noncompliant: The process set 777 permissions to this file
15  chmod("myfile.txt", S_IRWXU | S_IRWXG | S_IRWXO);
16}

```

#### Compliant example

```c
1#include <sys/stat.h>
2#include <fcntl.h>
3
4void looseFilePermissionsCompliant(){
5
6  int fd;
7
8  // Compliant: The O_CREAT flag indicates that the file should be created if it doesn't exist.
9  open("myfile.txt", O_CREAT, S_IRWXU | S_IRWXG);
10  // Compliant: further created files or directories will not have permissions set for "other group"
11  umask(S_IRWXO);
12
13}

```
