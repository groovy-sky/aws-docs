---
title: "Os Command Injection High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Go detectors(45/45)

[Useless if Body](useless-if-body.md) [Channel Guarded With Mutex](channel-guarded-with-mutex.md) [Improper Certificate Validation](improper-certificate-validation.md) [Unvalidated S3 Bucket Ownership](s3-verify-bucket-owner.md) [Resource Leak](resource-leak.md) [Insecure Cookie](insecure-cookie.md) [Weak Random Number Generation](weak-random-number-generation.md) [Redundant Equality Check](redundant-equality-check.md) [Insecure Ignore Host Key](insecure-ignore-host-key.md) [Unsafe Reflection](unsafe-reflection.md) [Unchecked Batch Operation Failures](aws-unchecked-batch-failures.md) [Lambda Client Reuse](lambda-client-reuse.md) [Os Command Injection](os-command-injection.md) [Useless if Conditional](useless-if-conditional.md) [Log Injection](log-injection.md) [Httptrace FileServer As Handler](http-trace-file-server-as-handler.md) [Pprof Endpoint](pprof-endpoint.md) [Cross Site Scripting (XSS)](cross-site-scripting.md) [Not Recommended API Usage](not-recommended-apis.md) [Hidden Goroutine](hidden-goroutine.md) [Channel Accessible By Non Endpoint](channel-accessible-by-non-endpoint.md) [Decompression Bomb](decompression-bomb.md) [Cross-Site Request Forgery (CSRF)](cross-site-request-forgery.md) [Thread Safety Violation](thread-safety-violation.md) [Insecure Connection](insecure-connection.md) [SQL Injection](sql-injection.md) [Deprecated Key Generator](deprecated-key-generator.md) [Exported Loop Pointer](exported-loop-pointer.md) [Server Side Request Forgery (SSRF)](server-side-request-forgery.md) [Sensitive Information Leak](sensitive-information-leak.md) [Integer Overflow](integer-overflow.md) [Missing Pagination](missing-pagination.md) [Insecure Cryptography](insecure-cryptography.md) [Protection Mechanism Failure](protection-mechanism-failure.md) [Nil Pointer Dereference](nil-pointer-dereference.md) [Temporary Files](temporary-files.md) [XML External Entity](xml-external-entity.md) [Insecure File Permissions](insecure-file-permissions.md) [Authentication Bypass By Alternate Name](authentication-bypass-by-alternate-name.md) [Code Injection](code-injection.md) [Improper authentication](improper-authentication.md) [Use Filepath Join](use-filepath-join.md) [Path Traversal](path-traversal.md) [Write Pprof Profile Output](write-pprof-profile-output.md) [Hardcoded true or false](hardcoded-eq-true-or-false.md)

# Os Command Injection [High](severity/high.md)

OS command injection is a critical vulnerability that can lead to a full system compromise as it may allow an adversary to pass in arbitrary commands or arguments to be executed. User input should never be used in constructing commands or command arguments to functions which execute OS commands. This includes filenames supplied by user uploads or downloads. The application should have a hardcoded set of arguments that are to be passed to OS commands. If filenames are being passed to these functions, it is recommended that a hash of the filename be used instead, or some other unique identifier. It is strongly recommended that a native library that implements the same functionality be used instead of using OS system commands, due to the risk of unknown attacks against third party commands.

**Detector ID**

go/os-command-injection@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-78](https://cwe.mitre.org/data/definitions/78.html)

**Tags**

[\# injection](tags/injection.md) [\# owasp-top10](tags/owasp-top10.md) [\# top25-cwes](tags/top25-cwes.md)

* * *

All content copied from https://docs.aws.amazon.com/.
