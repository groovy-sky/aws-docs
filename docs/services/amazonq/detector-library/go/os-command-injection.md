![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Go detectors(45/45)

[Useless if Body](useless-if-body.md) [Channel Guarded With Mutex](channel-guarded-with-mutex.md) [Improper Certificate Validation](improper-certificate-validation.md) [Unvalidated S3 Bucket Ownership](s3-verify-bucket-owner.md) [Resource Leak](resource-leak.md) [Insecure Cookie](insecure-cookie.md) [Weak Random Number Generation](weak-random-number-generation.md) [Redundant Equality Check](redundant-equality-check.md) [Insecure Ignore Host Key](insecure-ignore-host-key.md) [Unsafe Reflection](unsafe-reflection.md) [Unchecked Batch Operation Failures](aws-unchecked-batch-failures.md) [Lambda Client Reuse](lambda-client-reuse.md) [Os Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/os-command-injection) [Useless if Conditional](https://docs.aws.amazon.com/amazonq/detector-library/go/useless-if-conditional) [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/log-injection) [Httptrace FileServer As Handler](https://docs.aws.amazon.com/amazonq/detector-library/go/http-trace-file-server-as-handler) [Pprof Endpoint](https://docs.aws.amazon.com/amazonq/detector-library/go/pprof-endpoint) [Cross Site Scripting (XSS)](https://docs.aws.amazon.com/amazonq/detector-library/go/cross-site-scripting) [Not Recommended API Usage](https://docs.aws.amazon.com/amazonq/detector-library/go/not-recommended-apis) [Hidden Goroutine](https://docs.aws.amazon.com/amazonq/detector-library/go/hidden-goroutine) [Channel Accessible By Non Endpoint](https://docs.aws.amazon.com/amazonq/detector-library/go/channel-accessible-by-non-endpoint) [Decompression Bomb](https://docs.aws.amazon.com/amazonq/detector-library/go/decompression-bomb) [Cross-Site Request Forgery (CSRF)](https://docs.aws.amazon.com/amazonq/detector-library/go/cross-site-request-forgery) [Thread Safety Violation](https://docs.aws.amazon.com/amazonq/detector-library/go/thread-safety-violation) [Insecure Connection](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-connection) [SQL Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/sql-injection) [Deprecated Key Generator](https://docs.aws.amazon.com/amazonq/detector-library/go/deprecated-key-generator) [Exported Loop Pointer](https://docs.aws.amazon.com/amazonq/detector-library/go/exported-loop-pointer) [Server Side Request Forgery (SSRF)](https://docs.aws.amazon.com/amazonq/detector-library/go/server-side-request-forgery) [Sensitive Information Leak](https://docs.aws.amazon.com/amazonq/detector-library/go/sensitive-information-leak) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/go/integer-overflow) [Missing Pagination](https://docs.aws.amazon.com/amazonq/detector-library/go/missing-pagination) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-cryptography) [Protection Mechanism Failure](https://docs.aws.amazon.com/amazonq/detector-library/go/protection-mechanism-failure) [Nil Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/go/nil-pointer-dereference) [Temporary Files](https://docs.aws.amazon.com/amazonq/detector-library/go/temporary-files) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/go/xml-external-entity) [Insecure File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-file-permissions) [Authentication Bypass By Alternate Name](https://docs.aws.amazon.com/amazonq/detector-library/go/authentication-bypass-by-alternate-name) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/code-injection) [Improper authentication](https://docs.aws.amazon.com/amazonq/detector-library/go/improper-authentication) [Use Filepath Join](https://docs.aws.amazon.com/amazonq/detector-library/go/use-filepath-join) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/go/path-traversal) [Write Pprof Profile Output](https://docs.aws.amazon.com/amazonq/detector-library/go/write-pprof-profile-output) [Hardcoded true or false](https://docs.aws.amazon.com/amazonq/detector-library/go/hardcoded-eq-true-or-false)

# Os Command Injection [High](https://docs.aws.amazon.com/amazonq/detector-library/go/severity/high)

OS command injection is a critical vulnerability that can lead to a full system compromise as it may allow an adversary to pass in arbitrary commands or arguments to be executed. User input should never be used in constructing commands or command arguments to functions which execute OS commands. This includes filenames supplied by user uploads or downloads. The application should have a hardcoded set of arguments that are to be passed to OS commands. If filenames are being passed to these functions, it is recommended that a hash of the filename be used instead, or some other unique identifier. It is strongly recommended that a native library that implements the same functionality be used instead of using OS system commands, due to the risk of unknown attacks against third party commands.

**Detector ID**

go/os-command-injection@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/go/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-78](https://cwe.mitre.org/data/definitions/78.html)

**Tags**

[\# injection](https://docs.aws.amazon.com/amazonq/detector-library/go/tags/injection) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/go/tags/owasp-top10) [\# top25-cwes](https://docs.aws.amazon.com/amazonq/detector-library/go/tags/top25-cwes)

* * *
