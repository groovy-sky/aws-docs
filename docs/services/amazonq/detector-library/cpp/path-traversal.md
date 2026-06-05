---
title: "Path traversal High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](loose-file-permissions.md) [Sensitive information leak](sensitive-information-leak.md) [Missing Authorization](missing-authorization.md) [Return Stack Address](return-stack-address.md) [OS Command Injection](os-command-injection.md) [Use After Free](use-after-free.md) [Incorrect Comparison](incorrect-comparison.md) [off by one error](off-by-one-error.md) [Path traversal](path-traversal.md) [Insecure temporary file or directory](insecure-temporary-file-or-directory.md) [Insecure Cryptography](insecure-cryptography.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Unchecked Null Dereference](unchecked-null-dereference.md) [SQL injection](sql-injection.md) [Missing check on method output](missing-check-on-method-output.md) [Improper Restriction on Memory Buffer](improper-restriction-on-memory-buffer.md) [Multiple Locks](multiple-locks.md) [Improper Input Validation](improper-input-validation.md) [Null Pointer Dereference](null-pointer-dereference.md) [Use Of Redundant Code](use-of-redundant-code.md) [Improper Certificate Validation](improper-certificate-validation.md) [Improper Authentication](improper-authentication.md)

# Path traversal [High](severity/high.md)

Creating file paths from untrusted input could allow a malicious actor to access arbitrary files on a disk by manipulating the file name in the path.

**Detector ID**

cpp/path-traversal@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-22](https://cwe.mitre.org/data/definitions/22.html) [CWE-23](https://cwe.mitre.org/data/definitions/23.html) [CWE-24](https://cwe.mitre.org/data/definitions/24.html)

**Tags**

[\# owasp-top10](tags/owasp-top10.md) [\# top25-cwes](tags/top25-cwes.md)

* * *

#### Noncompliant example

```cpp
1using namespace std;
2#define CROW_MAIN
3#include <iostream>
4#include <fstream>
5#include <memory>
6#include <crow_all.h>
7
8int pathTraversalNoncompliant() {
9     crow::SimpleApp app;
10
11     CROW_ROUTE(app, "/download/<path>")
12         ([](const crow::request& req, crow::response& res, const std::string& filePath) {
13             ofstream file("uploads/" + filePath, ios::out);
14             // Noncompliant: filepath does not properly validate.
15             if (file.is_open()) {
16                 std::string content((std::istreambuf_iterator<char>(file)), std::istreambuf_iterator<char>());
17                 res.write(content);
18                 res.end();
19             } else {
20                 res.code = 404;
21                 res.write("File not found");
22                 res.end();
23             }
24         });
25
26     app.port(8080).multithreaded().run();
27     return 0;
28}

```

#### Compliant example

```cpp
1using namespace std;
2#define CROW_MAIN
3#include <iostream>
4#include <fstream>
5#include <memory>
6#include <crow_all.h>
7
8int pathTraversalCompliant() {
9     crow::SimpleApp app;
10
11     CROW_ROUTE(app, "/download/<path>")
12         ([](const crow::request& req, crow::response& res, const std::string& filePath) {
13             if (isValidFilePath(filePath)) {
14                 std::ifstream file("uploads/" + filePath);
15                 // Compliant: filepath is properly validated.
16                 if (file.is_open()) {
17                     std::string content((std::istreambuf_iterator<char>(file)), std::istreambuf_iterator<char>());
18                     res.write(content);
19                     res.end();
20                 } else {
21                     res.code = 404;
22                     res.write("File not found");
23                     res.end();
24                 }
25             } else {
26                 res.code = 400;
27                 res.write("Bad Request: Invalid file path");
28                 res.end();
29             }
30         });
31
32     app.port(8080).multithreaded().run();
33     return 0;
34 }

```

All content copied from https://docs.aws.amazon.com/.
