![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Go detectors(45/45)

[Useless if Body](../useless-if-body.md) [Channel Guarded With Mutex](../channel-guarded-with-mutex.md) [Improper Certificate Validation](../improper-certificate-validation.md) [Unvalidated S3 Bucket Ownership](../s3-verify-bucket-owner.md) [Resource Leak](../resource-leak.md) [Insecure Cookie](../insecure-cookie.md) [Weak Random Number Generation](../weak-random-number-generation.md) [Redundant Equality Check](../redundant-equality-check.md) [Insecure Ignore Host Key](../insecure-ignore-host-key.md) [Unsafe Reflection](../unsafe-reflection.md) [Unchecked Batch Operation Failures](../aws-unchecked-batch-failures.md) [Lambda Client Reuse](../lambda-client-reuse.md) [Os Command Injection](../os-command-injection.md) [Useless if Conditional](../useless-if-conditional.md) [Log Injection](../log-injection.md) [Httptrace FileServer As Handler](../http-trace-file-server-as-handler.md) [Pprof Endpoint](../pprof-endpoint.md) [Cross Site Scripting (XSS)](../cross-site-scripting.md) [Not Recommended API Usage](../not-recommended-apis.md) [Hidden Goroutine](../hidden-goroutine.md) [Channel Accessible By Non Endpoint](../channel-accessible-by-non-endpoint.md) [Decompression Bomb](../decompression-bomb.md) [Cross-Site Request Forgery (CSRF)](../cross-site-request-forgery.md) [Thread Safety Violation](../thread-safety-violation.md) [Insecure Connection](../insecure-connection.md) [SQL Injection](../sql-injection.md) [Deprecated Key Generator](../deprecated-key-generator.md) [Exported Loop Pointer](../exported-loop-pointer.md) [Server Side Request Forgery (SSRF)](../server-side-request-forgery.md) [Sensitive Information Leak](../sensitive-information-leak.md) [Integer Overflow](../integer-overflow.md) [Missing Pagination](../missing-pagination.md) [Insecure Cryptography](../insecure-cryptography.md) [Protection Mechanism Failure](../protection-mechanism-failure.md) [Nil Pointer Dereference](../nil-pointer-dereference.md) [Temporary Files](../temporary-files.md) [XML External Entity](../xml-external-entity.md) [Insecure File Permissions](../insecure-file-permissions.md) [Authentication Bypass By Alternate Name](../authentication-bypass-by-alternate-name.md) [Code Injection](../code-injection.md) [Improper authentication](../improper-authentication.md) [Use Filepath Join](../use-filepath-join.md) [Path Traversal](../path-traversal.md) [Write Pprof Profile Output](../write-pprof-profile-output.md) [Hardcoded true or false](../hardcoded-eq-true-or-false.md)

# Security detectors

### [Useless if Body](../useless-if-body.md)

If statement with identical bodies in if and else blocks

### [Channel Guarded With Mutex](../channel-guarded-with-mutex.md)

Redundant mutex guards on channels in Go

### [Improper Certificate Validation](../improper-certificate-validation.md)

Disabled TLS certificate validation

### [Unvalidated S3 Bucket Ownership](../s3-verify-bucket-owner.md)

S3 bucket operations without owner validation

### [Resource Leak](../resource-leak.md)

Improper resource handling leading to resource exhaustion or arbitrary code execution

### [Insecure Cookie](../insecure-cookie.md)

Cookies created without HttpOnly and Secure flags

### [Weak Random Number Generation](../weak-random-number-generation.md)

Use of insecure math/rand for random number generation

### [Redundant Equality Check](../redundant-equality-check.md)

Redundant equality checks affect code quality and return predictable results

### [Insecure Ignore Host Key](../insecure-ignore-host-key.md)

Disabling SSH host key validation

### [Unsafe Reflection](../unsafe-reflection.md)

Use of adversary-controlled input in reflection

### [Unchecked Batch Operation Failures](../aws-unchecked-batch-failures.md)

Unhandled failures in AWS batch operations

### [Lambda Client Reuse](../lambda-client-reuse.md)

AWS client re-creation in Lambda handlers

### [Os Command Injection](../os-command-injection.md)

OS command injection from untrusted input

### [Useless if Conditional](../useless-if-conditional.md)

Redundant conditional checks

### [Log Injection](../log-injection.md)

Log injection from untrusted input

### [Httptrace FileServer As Handler](../http-trace-file-server-as-handler.md)

Using http.FileServer as handler

### [Pprof Endpoint](../pprof-endpoint.md)

Exposed pprof endpoints enable information leaks

### [Cross Site Scripting (XSS)](../cross-site-scripting.md)

XSS from untrusted input in web outputs

### [Not Recommended API Usage](../not-recommended-apis.md)

Security risks and quality issues from deprecated AWS APIs and clients

### [Hidden Goroutine](../hidden-goroutine.md)

Asynchronous hidden goroutine function invocations

### [Channel Accessible By Non Endpoint](../channel-accessible-by-non-endpoint.md)

Insecure gRPC client and server connections in Go enable data tampering

### [Decompression Bomb](../decompression-bomb.md)

Decompression of untrusted data without size limits

### [Cross-Site Request Forgery (CSRF)](../cross-site-request-forgery.md)

Insecure validation and lack of restrictions enable cross-site request forgery

### [Thread Safety Violation](../thread-safety-violation.md)

Unsynchronized concurrent access to shared data

### [Insecure Connection](../insecure-connection.md)

Plain HTTP traffic enables eavesdropping and tampering

### [SQL Injection](../sql-injection.md)

Improper Neutralization of Special Elements used in an SQL Command

### [Deprecated Key Generator](../deprecated-key-generator.md)

Use of weak RSA key generation function

### [Exported Loop Pointer](../exported-loop-pointer.md)

Loop pointers exported directly can cause unintended behavior

### [Server Side Request Forgery (SSRF)](../server-side-request-forgery.md)

User input used unsanitized in outbound requests

### [Sensitive Information Leak](../sensitive-information-leak.md)

Unprotected sensitive data in network services and client alerts

### [Integer Overflow](../integer-overflow.md)

Integer overflow from improper input validation in conversions

### [Missing Pagination](../missing-pagination.md)

Missing pagination in paginated API calls

### [Insecure Cryptography](../insecure-cryptography.md)

Use of insecure cryptography

### [Protection Mechanism Failure](../protection-mechanism-failure.md)

Disabled or incorrectly used protection mechanism can lead to security vulnerabilities

### [Nil Pointer Dereference](../nil-pointer-dereference.md)

Dereferencing a nil pointer can lead to unexpected nil pointer exceptions.

### [Temporary Files](../temporary-files.md)

Insecure temporary file creation

### [XML External Entity](../xml-external-entity.md)

XXE vulnerability from XML

### [Insecure File Permissions](../insecure-file-permissions.md)

Overly permissive file permissions

### [Authentication Bypass By Alternate Name](../authentication-bypass-by-alternate-name.md)

Inconsistent variable assignment from multiple sources

### [Code Injection](../code-injection.md)

Code injection from untrusted input

### [Improper authentication](../improper-authentication.md)

Improper authentication from insufficient identity verification

### [Use Filepath Join](../use-filepath-join.md)

File path compatibility with different systems path separators risks from path.Join

### [Path Traversal](../path-traversal.md)

Path traversal from untrusted input

### [Write Pprof Profile Output](../write-pprof-profile-output.md)

Identified the presence of stack traces within HTTP response, posing a potential security risk if deployed in a user-facing manner in a production environment.

### [Hardcoded true or false](../hardcoded-eq-true-or-false.md)

Redundant true/false conditions in if statements
