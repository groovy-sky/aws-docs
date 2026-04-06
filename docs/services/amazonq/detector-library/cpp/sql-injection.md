![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](loose-file-permissions.md) [Sensitive information leak](sensitive-information-leak.md) [Missing Authorization](missing-authorization.md) [Return Stack Address](return-stack-address.md) [OS Command Injection](os-command-injection.md) [Use After Free](use-after-free.md) [Incorrect Comparison](incorrect-comparison.md) [off by one error](off-by-one-error.md) [Path traversal](path-traversal.md) [Insecure temporary file or directory](insecure-temporary-file-or-directory.md) [Insecure Cryptography](insecure-cryptography.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Unchecked Null Dereference](unchecked-null-dereference.md) [SQL injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/sql-injection) [Missing check on method output](https://docs.aws.amazon.com/amazonq/detector-library/cpp/missing-check-on-method-output) [Improper Restriction on Memory Buffer](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-restriction-on-memory-buffer) [Multiple Locks](https://docs.aws.amazon.com/amazonq/detector-library/cpp/multiple-locks) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-input-validation) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/null-pointer-dereference) [Use Of Redundant Code](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-of-redundant-code) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-certificate-validation) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-authentication)

# SQL injection [High](https://docs.aws.amazon.com/amazonq/detector-library/cpp/severity/high)

User-provided inputs must be sanitized before being used to generate a SQL database query. An untrusted input can be intentionally built by an attacker in order to run unwanted query statements, possibly allowing the attacker to read, modify, or delete database content.

**Detector ID**

cpp/sql-injection@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/cpp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-89](https://cwe.mitre.org/data/definitions/89.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags/injection) [\# sql](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags/sql) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags/top25-cwes)

* * *

#### Noncompliant example

```cpp
1#include <iostream>
2#include <sqlite3.h>
3
4int sqlInjectionNoncompliant(const std::string& username, const std::string& password) {
5    sqlite3* db;
6    sqlite3_open("users.db", &db);
7
8    std::string query = "SELECT * FROM users WHERE username = '" + username + "' AND password = '" + password + "';";
9
10    sqlite3_stmt* stmt;
11    // Noncompliant: This code takes a username and password from the user, constructs an SQL query, and tries to authenticate the user against an SQLite database.
12    if (sqlite3_prepare_v2(db, query.c_str(), -1, &stmt, nullptr) == SQLITE_OK) {
13        if (sqlite3_step(stmt) == SQLITE_ROW) {
14            std::cout << "Authentication successful.\n";
15            return 1;
16        }
17    }
18
19    std::cout << "Authentication failed.\n";
20    return 0;
21}

```

#### Compliant example

```cpp
1#include <iostream>
2#include <sqlite3.h>
3
4int sqlInjectionCompliant(const std::string& username, const std::string& password) {
5    sqlite3* db;
6    sqlite3_open("users.db", &db);
7
8    std::string query = "SELECT * FROM users WHERE username = ? AND password = ?;";
9
10    sqlite3_stmt* stmt;
11    if (sqlite3_prepare_v2(db, query.c_str(), -1, &stmt, nullptr) == SQLITE_OK) {
12        // Compliant: Used parameterized queries or prepared statements to avoid sql injection
13        sqlite3_bind_text(stmt, 1, username.c_str(), -1, SQLITE_STATIC);
14        sqlite3_bind_text(stmt, 2, password.c_str(), -1, SQLITE_STATIC);
15        if (sqlite3_step(stmt) == SQLITE_ROW) {
16            std::cout << "Authentication successful.\n";
17            return 1;
18        }
19    }
20
21    std::cout << "Authentication failed.\n";
22    return 0;
23}

```
