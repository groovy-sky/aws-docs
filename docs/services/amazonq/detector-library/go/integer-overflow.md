![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Go detectors(45/45)

[Useless if Body](useless-if-body.md) [Channel Guarded With Mutex](channel-guarded-with-mutex.md) [Improper Certificate Validation](improper-certificate-validation.md) [Unvalidated S3 Bucket Ownership](s3-verify-bucket-owner.md) [Resource Leak](resource-leak.md) [Insecure Cookie](insecure-cookie.md) [Weak Random Number Generation](weak-random-number-generation.md) [Redundant Equality Check](redundant-equality-check.md) [Insecure Ignore Host Key](insecure-ignore-host-key.md) [Unsafe Reflection](unsafe-reflection.md) [Unchecked Batch Operation Failures](aws-unchecked-batch-failures.md) [Lambda Client Reuse](lambda-client-reuse.md) [Os Command Injection](os-command-injection.md) [Useless if Conditional](useless-if-conditional.md) [Log Injection](log-injection.md) [Httptrace FileServer As Handler](http-trace-file-server-as-handler.md) [Pprof Endpoint](pprof-endpoint.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Not Recommended API Usage](not-recommended-apis.md) [Hidden Goroutine](hidden-goroutine.md) [Channel Accessible By Non Endpoint](channel-accessible-by-non-endpoint.md) [Decompression Bomb](decompression-bomb.md) [Cross-Site Request Forgery (CSRF)](cross-site-request-forgery.md) [Thread Safety Violation](thread-safety-violation.md) [Insecure Connection](insecure-connection.md) [SQL Injection](sql-injection.md) [Deprecated Key Generator](deprecated-key-generator.md) [Exported Loop Pointer](exported-loop-pointer.md) [Server Side Request Forgery (SSRF)](server-side-request-forgery.md) [Sensitive Information Leak](sensitive-information-leak.md) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/go/integer-overflow) [Missing Pagination](https://docs.aws.amazon.com/amazonq/detector-library/go/missing-pagination) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-cryptography) [Protection Mechanism Failure](https://docs.aws.amazon.com/amazonq/detector-library/go/protection-mechanism-failure) [Nil Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/go/nil-pointer-dereference) [Temporary Files](https://docs.aws.amazon.com/amazonq/detector-library/go/temporary-files) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/go/xml-external-entity) [Insecure File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-file-permissions) [Authentication Bypass By Alternate Name](https://docs.aws.amazon.com/amazonq/detector-library/go/authentication-bypass-by-alternate-name) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/code-injection) [Improper authentication](https://docs.aws.amazon.com/amazonq/detector-library/go/improper-authentication) [Use Filepath Join](https://docs.aws.amazon.com/amazonq/detector-library/go/use-filepath-join) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/go/path-traversal) [Write Pprof Profile Output](https://docs.aws.amazon.com/amazonq/detector-library/go/write-pprof-profile-output) [Hardcoded true or false](https://docs.aws.amazon.com/amazonq/detector-library/go/hardcoded-eq-true-or-false)

# Integer Overflow [High](https://docs.aws.amazon.com/amazonq/detector-library/go/severity/high)

Converting untrusted input strings directly to smaller integer types like int16 or int32 can cause integer overflow if the input represents a number outside the target type's range. Using strconv.Atoi on very large strings and type converting the result to a small integer type allows creating integers that wrap around the type's maximum value. This can result in data loss, unexpected behavior, and errors from the wrapped value. Input requires validation before such conversions to prevent overflow. Checks should ensure the input string can be converted and the resulting integer will fit into the target type's range. Adding input checks helps prevent overflow from large inputs.

**Detector ID**

go/integer-overflow@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/go/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-190](https://cwe.mitre.org/data/definitions/190.html)

**Tags**

[\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/go/tags/top25-cwes)

* * *
