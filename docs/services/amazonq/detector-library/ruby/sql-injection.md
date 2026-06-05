---
title: "SQL Injection High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](sql-injection.md) [Divide by Zero](divide-by-zero.md) [Sensitive HTTP Action](sensitive-http-action.md) [Insufficient Protected Credentials](insufficiently-protected-credentials.md) [Sensitive Information Leak](sensitive-information-leak.md) [Untrusted Deserialization](untrusted-deserialization.md) [Log Injection](log-injection.md) [XML External Entity](xml-external-entity.md) [Path Injection](path-traversal.md) [Http to File Access](http-to-file-access.md) [Code Injection](code-injection.md) [OS Command Injection](os-command-injection.md) [Resource leak](resource-leak.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Untrusted Open](untrusted-file-open.md) [Improper Input Validation](improper-input-validation.md) [Stack Trace Exposure](stack-trace-exposure.md) [Improper Certificate Validation](improper-certificate-validation.md) [send\_file Injection](sendfile-injection.md) [Unsafe File Permissions](loose-file-permissions.md) [Tainted Format](tainted-format.md)

# SQL Injection [High](severity/high.md)

User input is fed into an SQL command. This allows a user to execute an SQL command injection and run custom actions, which could leak sensitive data or delete data in the database. Ensure the user does not have direct influence on the command.

**Detector ID**

ruby/sql-injection@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-89](https://cwe.mitre.org/data/definitions/89.html)

**Tags**

-

* * *

#### Noncompliant example

```ruby
1require 'pg'
2
3def sql_injection_noncompliant(event:, context:)
4  conn = PG::Connection.open(:dbname => 'test')
5
6  # Noncompliant: User-controlled parameter is used in SQL Statement.
7  res2 = conn.exec_params('SELECT * FROM foobar WHERE id = %{id}' % {id: event['id']})
8
9end
```

#### Compliant example

```ruby
1require 'pg'
2
3def sql_injection_compliant(event:, context:)
4  conn = PG::Connection.open(:dbname => 'test')
5
6  # Compliant: Parameterized SQL Statement.
7  res = conn.exec_params('SELECT $1 AS a, $2 AS b, $3 AS c', [event['id'], 2, nil])
8
9end
```

All content copied from https://docs.aws.amazon.com/.
