![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C++ detectors(35/35)

[Disabled HTML autoescape](do-not-disable-html-autoescape.md) [Weak pseudorandom number generation](weak-random-number-generation.md) [Missing Default in Switch](missing-default-in-switch.md) [Unsafe File Extension](unsafe-file-extension.md) [Incorrect Order Of setuid and setgid](incorrect-order-setuid-setgid.md) [Out Of Bounds Read](out-of-bounds-read.md) [Out Of Bounds Write](out-of-bounds-write.md) [Thread safety violation](thread-safety-violation.md) [Incorrect Pointer Subtraction](pointer-subtraction.md) [File System Access](file-system-access.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use of Sizeof](incorrect-use-of-sizeof.md) [Incorrect Pointer Scaling](incorrect-pointer-scaling.md) [Loose File Permissions](loose-file-permissions.md) [Sensitive information leak](sensitive-information-leak.md) [Missing Authorization](missing-authorization.md) [Return Stack Address](return-stack-address.md) [OS Command Injection](os-command-injection.md) [Use After Free](use-after-free.md) [Incorrect Comparison](incorrect-comparison.md) [off by one error](off-by-one-error.md) [Path traversal](path-traversal.md) [Insecure temporary file or directory](insecure-temporary-file-or-directory.md) [Insecure Cryptography](insecure-cryptography.md) [Insecure connection using unencrypted protocol](insecure-connection.md) [Unchecked Null Dereference](unchecked-null-dereference.md) [SQL injection](sql-injection.md) [Missing check on method output](missing-check-on-method-output.md) [Improper Restriction on Memory Buffer](improper-restriction-on-memory-buffer.md) [Multiple Locks](multiple-locks.md) [Improper Input Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-input-validation) [Null Pointer Dereference](https://docs.aws.amazon.com/amazonq/detector-library/cpp/null-pointer-dereference) [Use Of Redundant Code](https://docs.aws.amazon.com/amazonq/detector-library/cpp/use-of-redundant-code) [Improper Certificate Validation](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-certificate-validation) [Improper Authentication](https://docs.aws.amazon.com/amazonq/detector-library/cpp/improper-authentication)

# Improper Input Validation [High](https://docs.aws.amazon.com/amazonq/detector-library/cpp/severity/high)

Improper input validation can enable attacks and lead to unwanted behavior. Parts of the system may receive unintended input, which could result in altered control flow, arbitrary control of a resource, or arbitrary code execution.

**Detector ID**

cpp/improper-input-validation@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/cpp/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-20](https://cwe.mitre.org/data/definitions/20.html)

**Tags**

[\# top10-cwes](https://docs.aws.amazon.com/amazonq/detector-library/cpp/tags/top10-cwes)

* * *

#### Noncompliant example

```cpp
1#include <iostream>
2
3void improperInputValidationNoncompliant()
4{
5    std::string userInput;
6    std::cout << "Enter input: ";
7    std::getline(std::cin, userInput);
8    // Noncompliant: Use user input without validation.
9    std::cout << " input is: " << userInput << std::endl;
10}

```

#### Compliant example

```cpp
1#include <iostream>
2
3void improperInputValidationCompliant()
4{
5    std::string userInput;
6    std::cout << "Enter input: ";
7    std::getline(std::cin, userInput);
8
9    // Compliant: Sanitize input by trimming leading and trailing whitespace
10    userInput.erase(0, userInput.find_first_not_of(" \t\r\n"));
11    userInput.erase(userInput.find_last_not_of(" \t\r\n") + 1);
12    std::cout << "Sanitized input is: " << userInput << std::endl;
13}

```
