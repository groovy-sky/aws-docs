![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](https://docs.aws.amazon.com/amazonq/detector-library/c/os-command-injection) [Incorrect Use Of Free](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-of-free) [Use Of Uninitialized Variable](https://docs.aws.amazon.com/amazonq/detector-library/c/use-of-uninitialized-variable) [Insecure Use strcat fn](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strcat-fn) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/c/sql-injection) [Bitwise Operator On Signed Operand](https://docs.aws.amazon.com/amazonq/detector-library/c/bitwise-operator-on-signed-operand) [Insecure use gets fn](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-gets-fn) [Random fd exhaustion](https://docs.aws.amazon.com/amazonq/detector-library/c/random-fd-exhaustion) [Redundant Free Usage](https://docs.aws.amazon.com/amazonq/detector-library/c/redundant-free-usage) [Insecure Use Memset](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-memset) [Divide By Zero.](https://docs.aws.amazon.com/amazonq/detector-library/c/divide-by-zero) [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/c/return-stack-address) [Unchecked Return Value](https://docs.aws.amazon.com/amazonq/detector-library/c/unchecked-return-value) [Incorrect Format Specifier](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-format-specifier) [Unhandled Expression Result](https://docs.aws.amazon.com/amazonq/detector-library/c/unhandled-expression-result) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/c/path-traversal) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-input-validation) [Out Of Bounds Read](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-read) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/c/integer-overflow) [Insecure use strtok function](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strtok-fn) [Improper size of a memory buffer](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-size-of-a-memory-buffer) [incomplete-cleanup](https://docs.aws.amazon.com/amazonq/detector-library/c/incomplete-cleanup) [Null pointer dereference](https://docs.aws.amazon.com/amazonq/detector-library/c/null-pointer-dereference) [Insecure Temporary File Or Directory](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-temporary-file-or-directory) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-buffer-access) [Incorrect Use Ato Fn](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-ato-fn) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/c/loose-file-permissions) [Exposure of Sensitive Information](https://docs.aws.amazon.com/amazonq/detector-library/c/exposure-of-sensitive-information) [Out-of-bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-write) [String Equality](https://docs.aws.amazon.com/amazonq/detector-library/c/string-equality)

# OS command injection [High](https://docs.aws.amazon.com/amazonq/detector-library/c/severity/high)

Constructing operating system or shell commands with unsanitized user input can lead to inadvertently running malicious code. Ensure proper sanitation / validation of user inputs before passing them ahead.

**Detector ID**

c/os-command-injection@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/c/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-78](https://cwe.mitre.org/data/definitions/78.html) [CWE-77](https://cwe.mitre.org/data/definitions/77.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/c/tags/injection) [\# subprocess](https://docs.aws.amazon.com/amazonq/detector-library/c/tags/subprocess) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/c/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/c/tags/top25-cwes)

* * *

#### Noncompliant example

```c
1#include <stdio.h>
2#include <unistd.h>
3#include <string.h>
4
5int osCommandInjectionNonCompliant(int argc, char **argv) {
6    char cat[] = "cat ";
7    char *command;
8    size_t commandLength;
9
10    commandLength = strlen(cat) + strlen(argv[1]) + 1;
11    command = (char *) malloc(commandLength);
12    strncpy(command, cat, commandLength);
13
14    // Noncompliant: argv[1] is concatenated into command without validation
15    strncat(command, argv[1], (commandLength - strlen(cat)) );
16
17    // A potentially untrusted input is passed into `system` function
18    system(command);
19    return (0);
20}

```

#### Compliant example

```c
1#include <stdio.h>
2#include <unistd.h>
3#include <string.h>
4
5int osCommandInjectionCompliant(int argc, char** argv) {
6    char cat[] = "cat ";
7    char *command;
8    size_t commandLength;
9
10    commandLength = strlen(cat) + 1;
11    command = (char *) malloc(commandLength);
12    strncpy(command, cat, commandLength);
13
14    // Compliant: The (hardcoded) cat command will be executed
15    system(command);
16    return (0);
17}

```
