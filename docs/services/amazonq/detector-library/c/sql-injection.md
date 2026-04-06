![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/c/sql-injection) [Bitwise Operator On Signed Operand](https://docs.aws.amazon.com/amazonq/detector-library/c/bitwise-operator-on-signed-operand) [Insecure use gets fn](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-gets-fn) [Random fd exhaustion](https://docs.aws.amazon.com/amazonq/detector-library/c/random-fd-exhaustion) [Redundant Free Usage](https://docs.aws.amazon.com/amazonq/detector-library/c/redundant-free-usage) [Insecure Use Memset](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-memset) [Divide By Zero.](https://docs.aws.amazon.com/amazonq/detector-library/c/divide-by-zero) [Return Stack Address](https://docs.aws.amazon.com/amazonq/detector-library/c/return-stack-address) [Unchecked Return Value](https://docs.aws.amazon.com/amazonq/detector-library/c/unchecked-return-value) [Incorrect Format Specifier](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-format-specifier) [Unhandled Expression Result](https://docs.aws.amazon.com/amazonq/detector-library/c/unhandled-expression-result) [Path traversal](https://docs.aws.amazon.com/amazonq/detector-library/c/path-traversal) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-input-validation) [Out Of Bounds Read](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-read) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/c/integer-overflow) [Insecure use strtok function](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-use-strtok-fn) [Improper size of a memory buffer](https://docs.aws.amazon.com/amazonq/detector-library/c/improper-size-of-a-memory-buffer) [incomplete-cleanup](https://docs.aws.amazon.com/amazonq/detector-library/c/incomplete-cleanup) [Null pointer dereference](https://docs.aws.amazon.com/amazonq/detector-library/c/null-pointer-dereference) [Insecure Temporary File Or Directory](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-temporary-file-or-directory) [Insecure Buffer Access](https://docs.aws.amazon.com/amazonq/detector-library/c/insecure-buffer-access) [Incorrect Use Ato Fn](https://docs.aws.amazon.com/amazonq/detector-library/c/incorrect-use-ato-fn) [Loose File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/c/loose-file-permissions) [Exposure of Sensitive Information](https://docs.aws.amazon.com/amazonq/detector-library/c/exposure-of-sensitive-information) [Out-of-bounds Write](https://docs.aws.amazon.com/amazonq/detector-library/c/out-of-bounds-write) [String Equality](https://docs.aws.amazon.com/amazonq/detector-library/c/string-equality)

# SQL injection [High](https://docs.aws.amazon.com/amazonq/detector-library/c/severity/high)

User-provided inputs must be sanitized before being used to generate a SQL database query. An attacker can create and use untrusted input to run query statements that read, modify, or delete database content.

**Detector ID**

c/sql-injection@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/c/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-89](https://cwe.mitre.org/data/definitions/89.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/c/tags/injection) [\# sql](https://docs.aws.amazon.com/amazonq/detector-library/c/tags/sql) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/c/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/c/tags/top25-cwes)

* * *

#### Noncompliant example

```c
1#include <stdio.h>
2#include <mysql.h>
3#include <stdlib.h>
4#include <sqlite3.h>
5
6void sqlInjectionNonCompliant(int argc, char** argv) {
7    MYSQL *connection = mysql_init(NULL);
8    if (mysql_real_connect(connection, "localhost", "root", "root_passwd", NULL, 0, NULL, 0) == NULL) {
9        fprintf(stderr, "%s\n", mysql_error(connection));
10        mysql_close(connection);
11        exit(1);
12    }
13    char query[200];
14    // Noncompliant: Untrusted argv passed into query
15    sprintf(query, "SELECT * FROM users WHERE name = '%s'", argv[1]);
16    mysql_query(connection, query);
17}

```

#### Compliant example

```c
1#include <stdio.h>
2#include <mysql.h>
3#include <stdlib.h>
4#include <sqlite3.h>
5
6void sqlInjectionCompliant(int argc, char** argv) {
7    MYSQL *connection = mysql_init(NULL);
8    if (mysql_real_connect(connection, "localhost", "root", "root_passwd", NULL, 0, NULL, 0) == NULL) {
9        fprintf(stderr, "%s\n", mysql_error(connection));
10        mysql_close(connection);
11        exit(1);
12    }
13    char query[200];
14    char* name = argv[1];
15    char escaped_name[100];
16    mysql_real_escape_string(connection, escaped_name, name, strlen(name));
17    // Compliant: This is safe as `mysql_real_escape_string` escapes potentially malicious characters
18    sprintf(query, "SELECT * FROM users WHERE name = '%s'", escaped_name);
19    mysql_query(connection, query);
20}

```
