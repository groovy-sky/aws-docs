![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Go detectors(45/45)

[Useless if Body](useless-if-body.md) [Channel Guarded With Mutex](channel-guarded-with-mutex.md) [Improper Certificate Validation](improper-certificate-validation.md) [Unvalidated S3 Bucket Ownership](s3-verify-bucket-owner.md) [Resource Leak](resource-leak.md) [Insecure Cookie](insecure-cookie.md) [Weak Random Number Generation](weak-random-number-generation.md) [Redundant Equality Check](redundant-equality-check.md) [Insecure Ignore Host Key](insecure-ignore-host-key.md) [Unsafe Reflection](unsafe-reflection.md) [Unchecked Batch Operation Failures](aws-unchecked-batch-failures.md) [Lambda Client Reuse](lambda-client-reuse.md) [Os Command Injection](os-command-injection.md) [Useless if Conditional](useless-if-conditional.md) [Log Injection](log-injection.md) [Httptrace FileServer As Handler](http-trace-file-server-as-handler.md) [Pprof Endpoint](pprof-endpoint.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Not Recommended API Usage](not-recommended-apis.md) [Hidden Goroutine](hidden-goroutine.md) [Channel Accessible By Non Endpoint](channel-accessible-by-non-endpoint.md) [Decompression Bomb](decompression-bomb.md) [Cross-Site Request Forgery (CSRF)](cross-site-request-forgery.md) [Thread Safety Violation](thread-safety-violation.md) [Insecure Connection](insecure-connection.md) [SQL Injection](sql-injection.md) [Deprecated Key Generator](deprecated-key-generator.md) [Exported Loop Pointer](exported-loop-pointer.md) [Server Side Request Forgery (SSRF)](server-side-request-forgery.md) [Sensitive Information Leak](sensitive-information-leak.md) [Integer Overflow](integer-overflow.md) [Missing Pagination](missing-pagination.md) [Insecure Cryptography](insecure-cryptography.md) [Protection Mechanism Failure](protection-mechanism-failure.md) [Nil Pointer Dereference](nil-pointer-dereference.md) [Temporary Files](temporary-files.md) [XML External Entity](xml-external-entity.md) [Insecure File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-file-permissions) [Authentication Bypass By Alternate Name](https://docs.aws.amazon.com/amazonq/detector-library/go/authentication-bypass-by-alternate-name) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/code-injection) [Improper authentication](https://docs.aws.amazon.com/amazonq/detector-library/go/improper-authentication) [Use Filepath Join](https://docs.aws.amazon.com/amazonq/detector-library/go/use-filepath-join) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/go/path-traversal) [Write Pprof Profile Output](https://docs.aws.amazon.com/amazonq/detector-library/go/write-pprof-profile-output) [Hardcoded true or false](https://docs.aws.amazon.com/amazonq/detector-library/go/hardcoded-eq-true-or-false)

# Insecure File Permissions [High](https://docs.aws.amazon.com/amazonq/detector-library/go/severity/high)

Setting file permissions to world-readable/writable values allows unauthorized access. File permissions should be restricted to the minimum required - 0400 for read, 0200 for write, 0600 for read/write if only the app user should access the file. Overly permissive modes can allow attackers to access sensitive data. Setting permissions appropriately prevents malicious access.

**Detector ID**

go/insecure-file-permissions@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/go/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-276](https://cwe.mitre.org/data/definitions/276.html)

**Tags**

[\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/go/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/go/tags/top25-cwes)

* * *
