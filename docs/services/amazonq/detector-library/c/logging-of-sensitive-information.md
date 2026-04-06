![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](https://docs.aws.amazon.com/amazonq/detector-library/c/logging-of-sensitive-information) [Insecure Use Of Chroot](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-of-chroot) [Deadlock And Lock Inconsistency](https://docs.aws.amazon.com/amazonq/detector-library/c/deadlock-and-lock-inconsistency) [Unsafe File Extension](https://docs.aws.amazon.com/amazonq/detector-library/c/unsafe-file-extension) [OS command injection](https://docs.aws.amazon.com/amazonq/detector-library/c/os-command-injection) [Incorrect Use Of Free](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-of-free) [Use Of Uninitialized Variable](https://docs.aws.amazon.com/amazonq/detector-library/c/use-of-uninitialized-variable) [Insecure Use strcat fn](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strcat-fn) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/c/sql-injection) [Bitwise Operator On Signed Operand](https://docs.aws.amazon.com/amazonq/detector-library/c/bitwise-operator-on-signed-operand) [Insecure use gets fn](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-gets-fn) [Random fd exhaustion](https://docs.aws.amazon.com/amazonq/detector-library/c/random-fd-exhaustion) [Redundant Free Usage](https://docs.aws.amazon.com/amazonq/detector-library/c/redundant-free-usage) [Insecure Use Memset](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-memset) [Divide By Zero.](https://docs.aws.amazon.com/amazonq/detector-library/c/divide-by-zero) [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/c/return-stack-address) [Unchecked Return Value](https://docs.aws.amazon.com/amazonq/detector-library/c/unchecked-return-value) [Incorrect Format Specifier](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-format-specifier) [Unhandled Expression Result](https://docs.aws.amazon.com/amazonq/detector-library/c/unhandled-expression-result) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/c/path-traversal) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-input-validation) [Out Of Bounds Read](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-read) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/c/integer-overflow) [Insecure use strtok function](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strtok-fn) [Improper size of a memory buffer](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-size-of-a-memory-buffer) [incomplete-cleanup](https://docs.aws.amazon.com/amazonq/detector-library/c/incomplete-cleanup) [Null pointer dereference](https://docs.aws.amazon.com/amazonq/detector-library/c/null-pointer-dereference) [Insecure Temporary File Or Directory](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-temporary-file-or-directory) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-buffer-access) [Incorrect Use Ato Fn](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-ato-fn) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/c/loose-file-permissions) [Exposure of Sensitive Information](https://docs.aws.amazon.com/amazonq/detector-library/c/exposure-of-sensitive-information) [Out-of-bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-write) [String Equality](https://docs.aws.amazon.com/amazonq/detector-library/c/string-equality)

# Logging of sensitive information [High](https://docs.aws.amazon.com/amazonq/detector-library/c/severity/high)

We Observed that sensitive information has been logged in your code which may leads to sensitive information leak. Mitigate this issue by reviewing logging practices, minimizing the logging of sensitive data, using secure logging libraries, and implementing data masking techniques.

**Detector ID**

c/logging-of-sensitive-information@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/c/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-532](https://cwe.mitre.org/data/definitions/532.html)

**Tags**

-

* * *

#### Noncompliant example

```c
1#include <stdio.h>
2
3int loggingOfSensitiveInformationNonCompliant(int argc, char *argv[]) {
4    // Noncompliant: Logging sensitive information
5    printf(argv[1]);
6    return 0;
7}

```

#### Compliant example

```c
1#include <stdio.h>
2
3void loggingOfSensitiveInformationCompliant(const char *data) {
4    FILE *file = fopen("log.txt", "a");
5    if (file != NULL) {
6        // Redact sensitive information before logging
7        char redactedData[strlen(data) + 1];
8        strcpy(redactedData, data);
9        // Compliant: Replace sensitive information with placeholders or tokens
10        // For example, replace credit card numbers with "************"
11        // Modify this based on the type of sensitive data
12        redactCreditCardNumbers(redactedData); // Function to replace credit card numbers with ****
13
14        fputs(redactedData, file);
15        fputs("\n", file);
16        fclose(file);
17    }
18}

```
