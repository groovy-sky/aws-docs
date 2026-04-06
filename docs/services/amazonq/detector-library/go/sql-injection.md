![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Go detectors(45/45)

[Useless if Body](useless-if-body.md) [Channel Guarded With Mutex](channel-guarded-with-mutex.md) [Improper Certificate Validation](improper-certificate-validation.md) [Unvalidated S3 Bucket Ownership](s3-verify-bucket-owner.md) [Resource Leak](resource-leak.md) [Insecure Cookie](insecure-cookie.md) [Weak Random Number Generation](weak-random-number-generation.md) [Redundant Equality Check](redundant-equality-check.md) [Insecure Ignore Host Key](insecure-ignore-host-key.md) [Unsafe Reflection](unsafe-reflection.md) [Unchecked Batch Operation Failures](aws-unchecked-batch-failures.md) [Lambda Client Reuse](lambda-client-reuse.md) [Os Command Injection](os-command-injection.md) [Useless if Conditional](useless-if-conditional.md) [Log Injection](log-injection.md) [Httptrace FileServer As Handler](http-trace-file-server-as-handler.md) [Pprof Endpoint](pprof-endpoint.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Not Recommended API Usage](not-recommended-apis.md) [Hidden Goroutine](hidden-goroutine.md) [Channel Accessible By Non Endpoint](channel-accessible-by-non-endpoint.md) [Decompression Bomb](decompression-bomb.md) [Cross-Site Request Forgery (CSRF)](cross-site-request-forgery.md) [Thread Safety Violation](thread-safety-violation.md) [Insecure Connection](insecure-connection.md) [SQL Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/sql-injection) [Deprecated Key Generator](https://docs.aws.amazon.com/amazonq/detector-library/go/deprecated-key-generator) [Exported Loop Pointer](https://docs.aws.amazon.com/amazonq/detector-library/go/exported-loop-pointer) [Server Side Request Forgery (SSRF)](https://docs.aws.amazon.com/amazonq/detector-library/go/server-side-request-forgery) [Sensitive Information Leak](https://docs.aws.amazon.com/amazonq/detector-library/go/sensitive-information-leak) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/go/integer-overflow) [Missing Pagination](https://docs.aws.amazon.com/amazonq/detector-library/go/missing-pagination) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-cryptography) [Protection Mechanism Failure](https://docs.aws.amazon.com/amazonq/detector-library/go/protection-mechanism-failure) [Nil Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/go/nil-pointer-dereference) [Temporary Files](https://docs.aws.amazon.com/amazonq/detector-library/go/temporary-files) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/go/xml-external-entity) [Insecure File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-file-permissions) [Authentication Bypass By Alternate Name](https://docs.aws.amazon.com/amazonq/detector-library/go/authentication-bypass-by-alternate-name) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/code-injection) [Improper authentication](https://docs.aws.amazon.com/amazonq/detector-library/go/improper-authentication) [Use Filepath Join](https://docs.aws.amazon.com/amazonq/detector-library/go/use-filepath-join) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/go/path-traversal) [Write Pprof Profile Output](https://docs.aws.amazon.com/amazonq/detector-library/go/write-pprof-profile-output) [Hardcoded true or false](https://docs.aws.amazon.com/amazonq/detector-library/go/hardcoded-eq-true-or-false)

# SQL Injection [High](https://docs.aws.amazon.com/amazonq/detector-library/go/severity/high)

SQL Injection is a critical vulnerability that can lead to data or system compromise. By dynamically generating SQL query strings, user input may be able to influence the logic of the SQL statement. This could lead to an adversary accessing information they should not have access to or in some circumstances, being able to execute OS functionality or code. Replace all dynamically generated SQL queries with parameterized queries. In situations where dynamic queries must be created, never use direct user input, but instead use a map or dictionary of valid values and resolve them using a user supplied key.

**Detector ID**

go/sql-injection@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/go/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-89](https://cwe.mitre.org/data/definitions/89.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/go/tags/injection) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/go/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/go/tags/top25-cwes)

* * *
