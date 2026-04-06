![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](loose-file-permissions.md) [Sensitive information leak](sensitive-information-leak.md) [Missing Authorization](missing-authorization.md) [Return Stack Address](return-stack-address.md) [OS Command Injection](os-command-injection.md) [Use After Free](use-after-free.md) [Incorrect Comparison](incorrect-comparison.md) [off by one error](off-by-one-error.md) [Path traversal](path-traversal.md) [Insecure temporary file or directory](insecure-temporary-file-or-directory.md) [Insecure Cryptography](insecure-cryptography.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Unchecked Null Dereference](unchecked-null-dereference.md) [SQL injection](sql-injection.md) [Missing check on method output](missing-check-on-method-output.md) [Improper Restriction on Memory Buffer](improper-restriction-on-memory-buffer.md) [Multiple Locks](multiple-locks.md) [Improper Input Validation](improper-input-validation.md) [Null Pointer Dereference](null-pointer-dereference.md) [Use Of Redundant Code](use-of-redundant-code.md) [Improper Certificate Validation](improper-certificate-validation.md) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-authentication)

# Improper Authentication [High](https://docs.aws.amazon.com/amazonq/detector-library/cpp/severity/high)

Failing to properly verify user identities and authenticate against strong credentials enables attackers to bypass authentication controls. Weaknesses like hardcoded, empty, or missing credential checks allow unauthorized system and data access. User identities must be verified against secure credentials retrieved from env vars, vaults etc. before granting access. Proper authentication controls including credential strength verification are essential to prevent malicious login and account compromise.

**Detector ID**

cpp/improper-authentication@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/cpp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-287](https://cwe.mitre.org/data/definitions/287.html)

**Tags**

-

* * *

#### Noncompliant example

```cpp
1#include <iostream>
2
3void improperAuthenticationNoncompliant() {
4    std::string secret = "your_secret_key";
5
6    std::string correctToken = create()
7                                .set_issuer("your_issuer")
8                                .set_type("JWT")
9                                .set_payload_claim("user_id", claim("123"))
10                                .sign(algorithm::hs256{ secret });
11
12    std::string inCorrectToken = "invalid_token";
13
14    try {
15        // Noncompliant: Insecure Token has been used.
16        auto decoded_token = decode(inCorrectToken, algorithms({ algorithm::hs256{ secret } }));
17    } catch (const std::exception& e) {
18        std::cerr << "Error decoding or verifying token: " << e.what() << std::endl;
19    }
20}

```

#### Compliant example

```cpp
1#include <iostream>
2
3void improperAuthenticationCompliant() {
4    std::string secret = "your_secret_key";
5
6    std::string correctToken = create()
7                                .set_issuer("your_issuer")
8                                .set_type("JWT")
9                                .set_payload_claim("user_id", claim("123"))
10                                .sign(algorithm::hs256{ secret });
11
12    try {
13        // Compliant: Secure generated Token has been used.
14        auto decoded_token = decode(correctToken, algorithms({ algorithm::hs256{ secret } }));
15    } catch (const std::exception& e) {
16        std::cerr << "Error decoding or verifying token: " << e.what() << std::endl;
17    }
18}

```
