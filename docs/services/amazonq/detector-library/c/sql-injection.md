---
title: "SQL injection High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](divide-by-zero.md) [Return Stack Address](return-stack-address.md) [Unchecked Return Value](unchecked-return-value.md) [Incorrect Format Specifier](incorrect-format-specifier.md) [Unhandled Expression Result](unhandled-expression-result.md) [Path traversal](path-traversal.md) [Improper Input Validation](improper-input-validation.md) [Out Of Bounds Read](out-of-bounds-read.md) [Integer Overflow](integer-overflow.md) [Insecure use strtok function](insecure-use-strtok-fn.md) [Improper size of a memory buffer](improper-size-of-a-memory-buffer.md) [incomplete-cleanup](incomplete-cleanup.md) [Null pointer dereference](null-pointer-dereference.md) [Insecure Temporary File Or Directory](insecure-temporary-file-or-directory.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use Ato Fn](incorrect-use-ato-fn.md) [Loose File Permissions](loose-file-permissions.md) [Exposure of Sensitive Information](exposure-of-sensitive-information.md) [Out-of-bounds Write](out-of-bounds-write.md) [String Equality](string-equality.md)

# SQL injection [High](severity/high.md)

User-provided inputs must be sanitized before being used to generate a SQL database query. An attacker can create and use untrusted input to run query statements that read, modify, or delete database content.

**Detector ID**

c/sql-injection@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-89](https://cwe.mitre.org/data/definitions/89.html)

**Tags**

[\# injection](tags/injection.md) [\# sql](tags/sql.md) [\# owasp-top10](tags/owasp-top10.md) [\# top25-cwes](tags/top25-cwes.md)

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

All content copied from https://docs.aws.amazon.com/.
