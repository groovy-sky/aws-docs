![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Go detectors(45/45)

[Useless if Body](https://docs.aws.amazon.com/amazonq/detector-library/go/useless-if-body) [Channel Guarded With Mutex](https://docs.aws.amazon.com/amazonq/detector-library/go/channel-guarded-with-mutex) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/go/improper-certificate-validation) [Unvalidated S3 Bucket Ownership](https://docs.aws.amazon.com/amazonq/detector-library/go/s3-verify-bucket-owner) [Resource Leak](https://docs.aws.amazon.com/amazonq/detector-library/go/resource-leak) [Insecure Cookie](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-cookie) [Weak Random Number Generation](https://docs.aws.amazon.com/amazonq/detector-library/go/weak-random-number-generation) [Redundant Equality Check](https://docs.aws.amazon.com/amazonq/detector-library/go/redundant-equality-check) [Insecure Ignore Host Key](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-ignore-host-key) [Unsafe Reflection](https://docs.aws.amazon.com/amazonq/detector-library/go/unsafe-reflection) [Unchecked Batch Operation Failures](https://docs.aws.amazon.com/amazonq/detector-library/go/aws-unchecked-batch-failures) [Lambda Client Reuse](https://docs.aws.amazon.com/amazonq/detector-library/go/lambda-client-reuse) [Os Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/os-command-injection) [Useless if Conditional](https://docs.aws.amazon.com/amazonq/detector-library/go/useless-if-conditional) [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/log-injection) [Httptrace FileServer As Handler](https://docs.aws.amazon.com/amazonq/detector-library/go/http-trace-file-server-as-handler) [Pprof Endpoint](https://docs.aws.amazon.com/amazonq/detector-library/go/pprof-endpoint) [Cross Site Scripting (XSS)](https://docs.aws.amazon.com/amazonq/detector-library/go/cross-site-scripting) [Not Recommended API Usage](https://docs.aws.amazon.com/amazonq/detector-library/go/not-recommended-apis) [Hidden Goroutine](https://docs.aws.amazon.com/amazonq/detector-library/go/hidden-goroutine) [Channel Accessible By Non Endpoint](https://docs.aws.amazon.com/amazonq/detector-library/go/channel-accessible-by-non-endpoint) [Decompression Bomb](https://docs.aws.amazon.com/amazonq/detector-library/go/decompression-bomb) [Cross-Site Request Forgery (CSRF)](https://docs.aws.amazon.com/amazonq/detector-library/go/cross-site-request-forgery) [Thread Safety Violation](https://docs.aws.amazon.com/amazonq/detector-library/go/thread-safety-violation) [Insecure Connection](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-connection) [SQL Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/sql-injection) [Deprecated Key Generator](https://docs.aws.amazon.com/amazonq/detector-library/go/deprecated-key-generator) [Exported Loop Pointer](https://docs.aws.amazon.com/amazonq/detector-library/go/exported-loop-pointer) [Server Side Request Forgery (SSRF)](https://docs.aws.amazon.com/amazonq/detector-library/go/server-side-request-forgery) [Sensitive Information Leak](https://docs.aws.amazon.com/amazonq/detector-library/go/sensitive-information-leak) [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/go/integer-overflow) [Missing Pagination](https://docs.aws.amazon.com/amazonq/detector-library/go/missing-pagination) [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-cryptography) [Protection Mechanism Failure](https://docs.aws.amazon.com/amazonq/detector-library/go/protection-mechanism-failure) [Nil Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/go/nil-pointer-dereference) [Temporary Files](https://docs.aws.amazon.com/amazonq/detector-library/go/temporary-files) [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/go/xml-external-entity) [Insecure File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-file-permissions) [Authentication Bypass By Alternate Name](https://docs.aws.amazon.com/amazonq/detector-library/go/authentication-bypass-by-alternate-name) [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/code-injection) [Improper authentication](https://docs.aws.amazon.com/amazonq/detector-library/go/improper-authentication) [Use Filepath Join](https://docs.aws.amazon.com/amazonq/detector-library/go/use-filepath-join) [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/go/path-traversal) [Write Pprof Profile Output](https://docs.aws.amazon.com/amazonq/detector-library/go/write-pprof-profile-output) [Hardcoded true or false](https://docs.aws.amazon.com/amazonq/detector-library/go/hardcoded-eq-true-or-false)

# Go detectors

Showing all detectors for the Go language.

##### Browse by tags

Browse all detectors by tags.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/go/tags)

##### Browse by severity

Browse all detectors by severity.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/go/severity)

##### Browse by category

Browse all detectors by category.

[Click here→](https://docs.aws.amazon.com/amazonq/detector-library/go/categories)

* * *

### Browse all detectors

### [Useless if Body](https://docs.aws.amazon.com/amazonq/detector-library/go/useless-if-body)

If statement with identical bodies in if and else blocks

### [Channel Guarded With Mutex](https://docs.aws.amazon.com/amazonq/detector-library/go/channel-guarded-with-mutex)

Redundant mutex guards on channels in Go

### [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/go/improper-certificate-validation)

Disabled TLS certificate validation

### [Unvalidated S3 Bucket Ownership](https://docs.aws.amazon.com/amazonq/detector-library/go/s3-verify-bucket-owner)

S3 bucket operations without owner validation

### [Resource Leak](https://docs.aws.amazon.com/amazonq/detector-library/go/resource-leak)

Improper resource handling leading to resource exhaustion or arbitrary code execution

### [Insecure Cookie](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-cookie)

Cookies created without HttpOnly and Secure flags

### [Weak Random Number Generation](https://docs.aws.amazon.com/amazonq/detector-library/go/weak-random-number-generation)

Use of insecure math/rand for random number generation

### [Redundant Equality Check](https://docs.aws.amazon.com/amazonq/detector-library/go/redundant-equality-check)

Redundant equality checks affect code quality and return predictable results

### [Insecure Ignore Host Key](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-ignore-host-key)

Disabling SSH host key validation

### [Unsafe Reflection](https://docs.aws.amazon.com/amazonq/detector-library/go/unsafe-reflection)

Use of adversary-controlled input in reflection

### [Unchecked Batch Operation Failures](https://docs.aws.amazon.com/amazonq/detector-library/go/aws-unchecked-batch-failures)

Unhandled failures in AWS batch operations

### [Lambda Client Reuse](https://docs.aws.amazon.com/amazonq/detector-library/go/lambda-client-reuse)

AWS client re-creation in Lambda handlers

### [Os Command Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/os-command-injection)

OS command injection from untrusted input

### [Useless if Conditional](https://docs.aws.amazon.com/amazonq/detector-library/go/useless-if-conditional)

Redundant conditional checks

### [Log Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/log-injection)

Log injection from untrusted input

### [Httptrace FileServer As Handler](https://docs.aws.amazon.com/amazonq/detector-library/go/http-trace-file-server-as-handler)

Using http.FileServer as handler

### [Pprof Endpoint](https://docs.aws.amazon.com/amazonq/detector-library/go/pprof-endpoint)

Exposed pprof endpoints enable information leaks

### [Cross Site Scripting (XSS)](https://docs.aws.amazon.com/amazonq/detector-library/go/cross-site-scripting)

XSS from untrusted input in web outputs

### [Not Recommended API Usage](https://docs.aws.amazon.com/amazonq/detector-library/go/not-recommended-apis)

Security risks and quality issues from deprecated AWS APIs and clients

### [Hidden Goroutine](https://docs.aws.amazon.com/amazonq/detector-library/go/hidden-goroutine)

Asynchronous hidden goroutine function invocations

### [Channel Accessible By Non Endpoint](https://docs.aws.amazon.com/amazonq/detector-library/go/channel-accessible-by-non-endpoint)

Insecure gRPC client and server connections in Go enable data tampering

### [Decompression Bomb](https://docs.aws.amazon.com/amazonq/detector-library/go/decompression-bomb)

Decompression of untrusted data without size limits

### [Cross-Site Request Forgery (CSRF)](https://docs.aws.amazon.com/amazonq/detector-library/go/cross-site-request-forgery)

Insecure validation and lack of restrictions enable cross-site request forgery

### [Thread Safety Violation](https://docs.aws.amazon.com/amazonq/detector-library/go/thread-safety-violation)

Unsynchronized concurrent access to shared data

### [Insecure Connection](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-connection)

Plain HTTP traffic enables eavesdropping and tampering

### [SQL Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/sql-injection)

Improper Neutralization of Special Elements used in an SQL Command

### [Deprecated Key Generator](https://docs.aws.amazon.com/amazonq/detector-library/go/deprecated-key-generator)

Use of weak RSA key generation function

### [Exported Loop Pointer](https://docs.aws.amazon.com/amazonq/detector-library/go/exported-loop-pointer)

Loop pointers exported directly can cause unintended behavior

### [Server Side Request Forgery (SSRF)](https://docs.aws.amazon.com/amazonq/detector-library/go/server-side-request-forgery)

User input used unsanitized in outbound requests

### [Sensitive Information Leak](https://docs.aws.amazon.com/amazonq/detector-library/go/sensitive-information-leak)

Unprotected sensitive data in network services and client alerts

### [Integer Overflow](https://docs.aws.amazon.com/amazonq/detector-library/go/integer-overflow)

Integer overflow from improper input validation in conversions

### [Missing Pagination](https://docs.aws.amazon.com/amazonq/detector-library/go/missing-pagination)

Missing pagination in paginated API calls

### [Insecure Cryptography](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-cryptography)

Use of insecure cryptography

### [Protection Mechanism Failure](https://docs.aws.amazon.com/amazonq/detector-library/go/protection-mechanism-failure)

Disabled or incorrectly used protection mechanism can lead to security vulnerabilities

### [Nil Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/go/nil-pointer-dereference)

Dereferencing a nil pointer can lead to unexpected nil pointer exceptions.

### [Temporary Files](https://docs.aws.amazon.com/amazonq/detector-library/go/temporary-files)

Insecure temporary file creation

### [XML External Entity](https://docs.aws.amazon.com/amazonq/detector-library/go/xml-external-entity)

XXE vulnerability from XML

### [Insecure File Permissions](https://docs.aws.amazon.com/amazonq/detector-library/go/insecure-file-permissions)

Overly permissive file permissions

### [Authentication Bypass By Alternate Name](https://docs.aws.amazon.com/amazonq/detector-library/go/authentication-bypass-by-alternate-name)

Inconsistent variable assignment from multiple sources

### [Code Injection](https://docs.aws.amazon.com/amazonq/detector-library/go/code-injection)

Code injection from untrusted input

### [Improper authentication](https://docs.aws.amazon.com/amazonq/detector-library/go/improper-authentication)

Improper authentication from insufficient identity verification

### [Use Filepath Join](https://docs.aws.amazon.com/amazonq/detector-library/go/use-filepath-join)

File path compatibility with different systems path separators risks from path.Join

### [Path Traversal](https://docs.aws.amazon.com/amazonq/detector-library/go/path-traversal)

Path traversal from untrusted input

### [Write Pprof Profile Output](https://docs.aws.amazon.com/amazonq/detector-library/go/write-pprof-profile-output)

Identified the presence of stack traces within HTTP response, posing a potential security risk if deployed in a user-facing manner in a production environment.

### [Hardcoded true or false](https://docs.aws.amazon.com/amazonq/detector-library/go/hardcoded-eq-true-or-false)

Redundant true/false conditions in if statements
