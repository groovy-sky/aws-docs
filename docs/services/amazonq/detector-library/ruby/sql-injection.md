![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Ruby detectors(21/21)

[SQL Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sql-injection) [Divide by Zero](https://docs.aws.amazon.com/amazonq/detector-library/ruby/divide-by-zero) [Sensitive HTTP Action](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sensitive-http-action) [Insufficient Protected Credentials](https://docs.aws.amazon.com/amazonq/detector-library/ruby/insufficiently-protected-credentials) [Sensitive Information Leak](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sensitive-information-leak) [Untrusted Deserialization](https://docs.aws.amazon.com/amazonq/detector-library/ruby/untrusted-deserialization) [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/log-injection) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/ruby/xml-external-entity) [Path Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/path-traversal) [Http to File Access](https://docs.aws.amazon.com/amazonq/detector-library/ruby/http-to-file-access) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/code-injection) [OS Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/os-command-injection) [Resource leak](https://docs.aws.amazon.com/amazonq/detector-library/ruby/resource-leak) [Cross Site Scripting (XSS)](https://docs.aws.amazon.com/amazonq/detector-library/ruby/cross-site-scripting) [Untrusted Open](https://docs.aws.amazon.com/amazonq/detector-library/ruby/untrusted-file-open) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-input-validation) [Stack Trace Exposure](https://docs.aws.amazon.com/amazonq/detector-library/ruby/stack-trace-exposure) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/ruby/improper-certificate-validation) [send\_file Injection](https://docs.aws.amazon.com/amazonq/detector-library/ruby/sendfile-injection) [Unsafe File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/ruby/loose-file-permissions) [Tainted Format](https://docs.aws.amazon.com/amazonq/detector-library/ruby/tainted-format)

# SQL Injection [High](https://docs.aws.amazon.com/amazonq/detector-library/ruby/severity/high)

User input is fed into an SQL command. This allows a user to execute an SQL command injection and run custom actions, which could leak sensitive data or delete data in the database. Ensure the user does not have direct influence on the command.

**Detector ID**

ruby/sql-injection@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/ruby/categories/security)

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
