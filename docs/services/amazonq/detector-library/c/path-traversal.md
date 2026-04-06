![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](divide-by-zero.md) [Return Stack Address](return-stack-address.md) [Unchecked Return Value](unchecked-return-value.md) [Incorrect Format Specifier](incorrect-format-specifier.md) [Unhandled Expression Result](unhandled-expression-result.md) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/c/path-traversal) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-input-validation) [Out Of Bounds Read](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-read) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/c/integer-overflow) [Insecure use strtok function](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strtok-fn) [Improper size of a memory buffer](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-size-of-a-memory-buffer) [incomplete-cleanup](https://docs.aws.amazon.com/amazonq/detector-library/c/incomplete-cleanup) [Null pointer dereference](https://docs.aws.amazon.com/amazonq/detector-library/c/null-pointer-dereference) [Insecure Temporary File Or Directory](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-temporary-file-or-directory) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-buffer-access) [Incorrect Use Ato Fn](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-ato-fn) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/c/loose-file-permissions) [Exposure of Sensitive Information](https://docs.aws.amazon.com/amazonq/detector-library/c/exposure-of-sensitive-information) [Out-of-bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-write) [String Equality](https://docs.aws.amazon.com/amazonq/detector-library/c/string-equality)

# Path traversal [High](https://docs.aws.amazon.com/amazonq/detector-library/c/severity/high)

Creating file paths from untrusted input could allow a malicious actor to access arbitrary files on a disk by manipulating the file name in the path.

**Detector ID**

c/path-traversal@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/c/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-22](https://cwe.mitre.org/data/definitions/22.html) [CWE-23](https://cwe.mitre.org/data/definitions/23.html)

**Tags**

[\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/c/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/c/tags/top25-cwes)

* * *

#### Noncompliant example

```c
1#include <stdio.h>
2#include <string.h>
3#include <stdlib.h>
4
5void pathTraversalNonComplaint(int argc, char *argv[]) {
6  char filename[100];
7
8  strcpy(filename, argv[1]);
9  // Noncompliant: user input is used to construct file path
10  FILE *fp = fopen(filename, "r");
11  if(fp == NULL) {
12    printf("Error opening file\n");
13    return 1;
14  }
15  // Read file contents
16  fclose(fp);
17}

```

#### Compliant example

```c
1#include <stdio.h>
2#include <string.h>
3#include <stdlib.h>
4
5void pathTraversalComplaint() {
6    const char* zip_filename = "example.zip";
7    const char* output_dir = "output";
8
9    size_t len = strlen(zip_filename);
10    // Compliant: checking if the provided zip_filename ends with the ".zip" extension and if the output_dir exists before calling the extract_all_files function
11    if ((len < 4 || strcmp(zip_filename + len - 4, ".zip") == 0) && (access(output_dir, F_OK) == 0)) {
12        extract_all_files(zip_filename, output_dir);
13    }
14}

```
