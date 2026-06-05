---
title: "Go detectors"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Go detectors(45/45)

[Useless if Body](go/useless-if-body.md) [Channel Guarded With Mutex](go/channel-guarded-with-mutex.md) [Improper Certificate Validation](go/improper-certificate-validation.md) [Unvalidated S3 Bucket Ownership](go/s3-verify-bucket-owner.md) [Resource Leak](go/resource-leak.md) [Insecure Cookie](go/insecure-cookie.md) [Weak Random Number Generation](go/weak-random-number-generation.md) [Redundant Equality Check](go/redundant-equality-check.md) [Insecure Ignore Host Key](go/insecure-ignore-host-key.md) [Unsafe Reflection](go/unsafe-reflection.md) [Unchecked Batch Operation Failures](go/aws-unchecked-batch-failures.md) [Lambda Client Reuse](go/lambda-client-reuse.md) [Os Command Injection](go/os-command-injection.md) [Useless if Conditional](go/useless-if-conditional.md) [Log Injection](go/log-injection.md) [Httptrace FileServer As Handler](go/http-trace-file-server-as-handler.md) [Pprof Endpoint](go/pprof-endpoint.md) [Cross Site Scripting (XSS)](go/cross-site-scripting.md) [Not Recommended API Usage](go/not-recommended-apis.md) [Hidden Goroutine](go/hidden-goroutine.md) [Channel Accessible By Non Endpoint](go/channel-accessible-by-non-endpoint.md) [Decompression Bomb](go/decompression-bomb.md) [Cross-Site Request Forgery (CSRF)](go/cross-site-request-forgery.md) [Thread Safety Violation](go/thread-safety-violation.md) [Insecure Connection](go/insecure-connection.md) [SQL Injection](go/sql-injection.md) [Deprecated Key Generator](go/deprecated-key-generator.md) [Exported Loop Pointer](go/exported-loop-pointer.md) [Server Side Request Forgery (SSRF)](go/server-side-request-forgery.md) [Sensitive Information Leak](go/sensitive-information-leak.md) [Integer Overflow](go/integer-overflow.md) [Missing Pagination](go/missing-pagination.md) [Insecure Cryptography](go/insecure-cryptography.md) [Protection Mechanism Failure](go/protection-mechanism-failure.md) [Nil Pointer Dereference](go/nil-pointer-dereference.md) [Temporary Files](go/temporary-files.md) [XML External Entity](go/xml-external-entity.md) [Insecure File Permissions](go/insecure-file-permissions.md) [Authentication Bypass By Alternate Name](go/authentication-bypass-by-alternate-name.md) [Code Injection](go/code-injection.md) [Improper authentication](go/improper-authentication.md) [Use Filepath Join](go/use-filepath-join.md) [Path Traversal](go/path-traversal.md) [Write Pprof Profile Output](go/write-pprof-profile-output.md) [Hardcoded true or false](go/hardcoded-eq-true-or-false.md)

# Go detectors

Showing all detectors for the Go language.

##### Browse by tags

Browse all detectors by tags.

[Click here→](go/tags.md)

##### Browse by severity

Browse all detectors by severity.

[Click here→](go/severity.md)

##### Browse by category

Browse all detectors by category.

[Click here→](go/categories.md)

* * *

### Browse all detectors

### [Useless if Body](go/useless-if-body.md)

If statement with identical bodies in if and else blocks

### [Channel Guarded With Mutex](go/channel-guarded-with-mutex.md)

Redundant mutex guards on channels in Go

### [Improper Certificate Validation](go/improper-certificate-validation.md)

Disabled TLS certificate validation

### [Unvalidated S3 Bucket Ownership](go/s3-verify-bucket-owner.md)

S3 bucket operations without owner validation

### [Resource Leak](go/resource-leak.md)

Improper resource handling leading to resource exhaustion or arbitrary code execution

### [Insecure Cookie](go/insecure-cookie.md)

Cookies created without HttpOnly and Secure flags

### [Weak Random Number Generation](go/weak-random-number-generation.md)

Use of insecure math/rand for random number generation

### [Redundant Equality Check](go/redundant-equality-check.md)

Redundant equality checks affect code quality and return predictable results

### [Insecure Ignore Host Key](go/insecure-ignore-host-key.md)

Disabling SSH host key validation

### [Unsafe Reflection](go/unsafe-reflection.md)

Use of adversary-controlled input in reflection

### [Unchecked Batch Operation Failures](go/aws-unchecked-batch-failures.md)

Unhandled failures in AWS batch operations

### [Lambda Client Reuse](go/lambda-client-reuse.md)

AWS client re-creation in Lambda handlers

### [Os Command Injection](go/os-command-injection.md)

OS command injection from untrusted input

### [Useless if Conditional](go/useless-if-conditional.md)

Redundant conditional checks

### [Log Injection](go/log-injection.md)

Log injection from untrusted input

### [Httptrace FileServer As Handler](go/http-trace-file-server-as-handler.md)

Using http.FileServer as handler

### [Pprof Endpoint](go/pprof-endpoint.md)

Exposed pprof endpoints enable information leaks

### [Cross Site Scripting (XSS)](go/cross-site-scripting.md)

XSS from untrusted input in web outputs

### [Not Recommended API Usage](go/not-recommended-apis.md)

Security risks and quality issues from deprecated AWS APIs and clients

### [Hidden Goroutine](go/hidden-goroutine.md)

Asynchronous hidden goroutine function invocations

### [Channel Accessible By Non Endpoint](go/channel-accessible-by-non-endpoint.md)

Insecure gRPC client and server connections in Go enable data tampering

### [Decompression Bomb](go/decompression-bomb.md)

Decompression of untrusted data without size limits

### [Cross-Site Request Forgery (CSRF)](go/cross-site-request-forgery.md)

Insecure validation and lack of restrictions enable cross-site request forgery

### [Thread Safety Violation](go/thread-safety-violation.md)

Unsynchronized concurrent access to shared data

### [Insecure Connection](go/insecure-connection.md)

Plain HTTP traffic enables eavesdropping and tampering

### [SQL Injection](go/sql-injection.md)

Improper Neutralization of Special Elements used in an SQL Command

### [Deprecated Key Generator](go/deprecated-key-generator.md)

Use of weak RSA key generation function

### [Exported Loop Pointer](go/exported-loop-pointer.md)

Loop pointers exported directly can cause unintended behavior

### [Server Side Request Forgery (SSRF)](go/server-side-request-forgery.md)

User input used unsanitized in outbound requests

### [Sensitive Information Leak](go/sensitive-information-leak.md)

Unprotected sensitive data in network services and client alerts

### [Integer Overflow](go/integer-overflow.md)

Integer overflow from improper input validation in conversions

### [Missing Pagination](go/missing-pagination.md)

Missing pagination in paginated API calls

### [Insecure Cryptography](go/insecure-cryptography.md)

Use of insecure cryptography

### [Protection Mechanism Failure](go/protection-mechanism-failure.md)

Disabled or incorrectly used protection mechanism can lead to security vulnerabilities

### [Nil Pointer Dereference](go/nil-pointer-dereference.md)

Dereferencing a nil pointer can lead to unexpected nil pointer exceptions.

### [Temporary Files](go/temporary-files.md)

Insecure temporary file creation

### [XML External Entity](go/xml-external-entity.md)

XXE vulnerability from XML

### [Insecure File Permissions](go/insecure-file-permissions.md)

Overly permissive file permissions

### [Authentication Bypass By Alternate Name](go/authentication-bypass-by-alternate-name.md)

Inconsistent variable assignment from multiple sources

### [Code Injection](go/code-injection.md)

Code injection from untrusted input

### [Improper authentication](go/improper-authentication.md)

Improper authentication from insufficient identity verification

### [Use Filepath Join](go/use-filepath-join.md)

File path compatibility with different systems path separators risks from path.Join

### [Path Traversal](go/path-traversal.md)

Path traversal from untrusted input

### [Write Pprof Profile Output](go/write-pprof-profile-output.md)

Identified the presence of stack traces within HTTP response, posing a potential security risk if deployed in a user-facing manner in a production environment.

### [Hardcoded true or false](go/hardcoded-eq-true-or-false.md)

Redundant true/false conditions in if statements

All content copied from https://docs.aws.amazon.com/.
