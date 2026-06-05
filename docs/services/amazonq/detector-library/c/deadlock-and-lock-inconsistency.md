---
title: "Deadlock And Lock Inconsistency High"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### C detectors(34/34)

[Logging of sensitive information](logging-of-sensitive-information.md) [Insecure Use Of Chroot](insecure-use-of-chroot.md) [Deadlock And Lock Inconsistency](deadlock-and-lock-inconsistency.md) [Unsafe File Extension](unsafe-file-extension.md) [OS command injection](os-command-injection.md) [Incorrect Use Of Free](incorrect-use-of-free.md) [Use Of Uninitialized Variable](use-of-uninitialized-variable.md) [Insecure Use strcat fn](insecure-use-strcat-fn.md) [SQL injection](sql-injection.md) [Bitwise Operator On Signed Operand](bitwise-operator-on-signed-operand.md) [Insecure use gets fn](insecure-use-gets-fn.md) [Random fd exhaustion](random-fd-exhaustion.md) [Redundant Free Usage](redundant-free-usage.md) [Insecure Use Memset](insecure-use-memset.md) [Divide By Zero.](divide-by-zero.md) [Return Stack Address](return-stack-address.md) [Unchecked Return Value](unchecked-return-value.md) [Incorrect Format Specifier](incorrect-format-specifier.md) [Unhandled Expression Result](unhandled-expression-result.md) [Path traversal](path-traversal.md) [Improper Input Validation](improper-input-validation.md) [Out Of Bounds Read](out-of-bounds-read.md) [Integer Overflow](integer-overflow.md) [Insecure use strtok function](insecure-use-strtok-fn.md) [Improper size of a memory buffer](improper-size-of-a-memory-buffer.md) [incomplete-cleanup](incomplete-cleanup.md) [Null pointer dereference](null-pointer-dereference.md) [Insecure Temporary File Or Directory](insecure-temporary-file-or-directory.md) [Insecure Buffer Access](insecure-buffer-access.md) [Incorrect Use Ato Fn](incorrect-use-ato-fn.md) [Loose File Permissions](loose-file-permissions.md) [Exposure of Sensitive Information](exposure-of-sensitive-information.md) [Out-of-bounds Write](out-of-bounds-write.md) [String Equality](string-equality.md)

# Deadlock And Lock Inconsistency [High](severity/high.md)

We observed that your code contains either incorrect lock ordering, nested locking, which can potentially lead to a deadlock or incorrectly handles locking by not following proper conventions, potentially introducing vulnerabilities through lock consistency violations. To mitigate this ensure correct lock ordering, initialize the mutex with recursive attributes and Follow standardized locking conventions.

**Detector ID**

c/deadlock-and-lock-inconsistency@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-362](https://cwe.mitre.org/data/definitions/362.html)

**Tags**

-

* * *

#### Noncompliant example

```c
1#include <limits.h>
2#include <pthread.h>
3#include <stdio.h>
4
5void deadlockAndLockInconsistencyNonComplaint() {
6    pthread_mutex_init(&lock, NULL);
7    // Noncompliant: Lock never acquired
8    pthread_mutex_unlock(&lock);
9}

```

#### Compliant example

```c
1#include <limits.h>
2#include <pthread.h>
3#include <stdio.h>
4
5pthread_mutex_t lock;
6
7void deadlockAndLockInconsistencyCompliant()
8{
9    // Compliant: This code dose not contains a potential deadlock or violate lock consistency due to incorrect lock ordering or nested locking.
10    pthread_mutex_init(&lock, NULL); // initialize mutex first
11    pthread_mutex_lock(&lock); // okay to lock now
12    // critical section
13    pthread_mutex_unlock(&lock);
14}

```

All content copied from https://docs.aws.amazon.com/.
