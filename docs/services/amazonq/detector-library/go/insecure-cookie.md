![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Go detectors(45/45)

[Useless if Body](useless-if-body.md) [Channel Guarded With Mutex](channel-guarded-with-mutex.md) [Improper Certificate Validation](improper-certificate-validation.md) [Unvalidated S3 Bucket Ownership](s3-verify-bucket-owner.md) [Resource Leak](resource-leak.md) [Insecure Cookie](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-cookie) [Weak Random Number Generation](https://docs.aws.amazon.com/amazonq/detector-library/go/weak-random-number-generation) [Redundant Equality Check](https://docs.aws.amazon.com/amazonq/detector-library/go/redundant-equality-check) [Insecure Ignore Host Key](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-ignore-host-key) [Unsafe Reflection](https://docs.aws.amazon.com/amazonq/detector-library/go/unsafe-reflection) [Unchecked Batch Operation Failures](https://docs.aws.amazon.com/amazonq/detector-library/go/aws-unchecked-batch-failures) [Lambda Client Reuse](https://docs.aws.amazon.com/amazonq/detector-library/go/lambda-client-reuse) [Os Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/os-command-injection) [Useless if Conditional](https://docs.aws.amazon.com/amazonq/detector-library/go/useless-if-conditional) [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/log-injection) [Httptrace FileServer As Handler](https://docs.aws.amazon.com/amazonq/detector-library/go/http-trace-file-server-as-handler) [Pprof Endpoint](https://docs.aws.amazon.com/amazonq/detector-library/go/pprof-endpoint) [Cross Site Scripting (XSS)](https://docs.aws.amazon.com/amazonq/detector-library/go/cross-site-scripting) [Not Recommended API Usage](https://docs.aws.amazon.com/amazonq/detector-library/go/not-recommended-apis) [Hidden Goroutine](https://docs.aws.amazon.com/amazonq/detector-library/go/hidden-goroutine) [Channel Accessible By Non Endpoint](https://docs.aws.amazon.com/amazonq/detector-library/go/channel-accessible-by-non-endpoint) [Decompression Bomb](https://docs.aws.amazon.com/amazonq/detector-library/go/decompression-bomb) [Cross-Site Request Forgery (CSRF)](https://docs.aws.amazon.com/amazonq/detector-library/go/cross-site-request-forgery) [Thread Safety Violation](https://docs.aws.amazon.com/amazonq/detector-library/go/thread-safety-violation) [Insecure Connection](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-connection) [SQL Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/sql-injection) [Deprecated Key Generator](https://docs.aws.amazon.com/amazonq/detector-library/go/deprecated-key-generator) [Exported Loop Pointer](https://docs.aws.amazon.com/amazonq/detector-library/go/exported-loop-pointer) [Server Side Request Forgery (SSRF)](https://docs.aws.amazon.com/amazonq/detector-library/go/server-side-request-forgery) [Sensitive Information Leak](https://docs.aws.amazon.com/amazonq/detector-library/go/sensitive-information-leak) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/go/integer-overflow) [Missing Pagination](https://docs.aws.amazon.com/amazonq/detector-library/go/missing-pagination) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-cryptography) [Protection Mechanism Failure](https://docs.aws.amazon.com/amazonq/detector-library/go/protection-mechanism-failure) [Nil Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/go/nil-pointer-dereference) [Temporary Files](https://docs.aws.amazon.com/amazonq/detector-library/go/temporary-files) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/go/xml-external-entity) [Insecure File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-file-permissions) [Authentication Bypass By Alternate Name](https://docs.aws.amazon.com/amazonq/detector-library/go/authentication-bypass-by-alternate-name) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/code-injection) [Improper authentication](https://docs.aws.amazon.com/amazonq/detector-library/go/improper-authentication) [Use Filepath Join](https://docs.aws.amazon.com/amazonq/detector-library/go/use-filepath-join) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/go/path-traversal) [Write Pprof Profile Output](https://docs.aws.amazon.com/amazonq/detector-library/go/write-pprof-profile-output) [Hardcoded true or false](https://docs.aws.amazon.com/amazonq/detector-library/go/hardcoded-eq-true-or-false)

# Insecure Cookie [Low](https://docs.aws.amazon.com/amazonq/detector-library/go/severity/low)

Cookies should always be created with the HttpOnly and Secure flags to prevent interception and theft. The HttpOnly flag disables client-side JavaScript access to mitigate XSS threats. The Secure flag restricts transmission to HTTPS connections to prevent MITM eavesdropping. Failing to set these flags when creating cookies with http.Cookie or gorilla/sessions leaves them vulnerable regardless of current contents. All cookies should be HttpOnly to prevent JavaScript access. Sensitive session and auth cookies should also be Secure to block MITM snooping. Enabling cookie security flags is essential to limit exposure and misuse.

**Detector ID**

go/insecure-cookie@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/go/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-614](https://cwe.mitre.org/data/definitions/614.html) [CWE-1004](https://cwe.mitre.org/data/definitions/1004.html)

**Tags**

[\# cookies](https://docs.aws.amazon.com/amazonq/detector-library/go/tags/cookies) [\# owasp-top10](https://docs.aws.amazon.com/amazonq/detector-library/go/tags/owasp-top10)

* * *
