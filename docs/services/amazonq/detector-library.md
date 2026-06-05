---
title: "Amazon Q  Detector Library"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

# Amazon Q  Detector Library

## Trained on decades of knowledge and experience across millions of code reviews

[![Java](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/java.3b781738.svg)Java](detector-library/java.md) [![Python](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/python.40e7bbb3.svg)Python](detector-library/python.md) [![JavaScript](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/javascript.f62769e6.svg)JavaScript](detector-library/javascript.md) [![JSX](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/jsx.041b6532.svg)JSX](detector-library/jsx.md) [![TypeScript](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/typescript.17fda397.svg)TypeScript](detector-library/typescript.md) [![C#](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/csharp.ee9c49fb.svg)C#](detector-library/csharp.md) [![CloudFormation](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/cloudformation.0125b063.svg)CloudFormation](detector-library/cloudformation.md) [![Terraform](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/terraform.2e7e71e7.svg)Terraform](detector-library/terraform.md) [![Go](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/go.dad1aa01.svg)Go](detector-library/go.md) [![Ruby](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/ruby.6896219f.svg)Ruby](detector-library/ruby.md) [![C](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/c.c9a38a87.svg)C](detector-library/c.md) [![C++](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/cpp.c908ae3d.svg)C++](detector-library/cpp.md) [![PHP](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/php.0357268e.svg)PHP](detector-library/php.md) [![Kotlin](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/kotlin.1bb7dd47.svg)Kotlin](detector-library/kotlin.md) [![Scala](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/scala.f104853e.svg)Scala](detector-library/scala.md) [![Shell](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/shell.c58a5629.svg)Shell](detector-library/shell.md)

## About Amazon Q Detector Library

The Amazon Q Detector Library describes the detectors used during code reviews to identify security and quality issues in code. Detectors contain rules that are used to identify critical security vulnerabilities like OWASP Top 10 and CWE Top 25 issues, including secrets exposure and package dependency vulnerabilities. They also detect code quality concerns such as IaC best practices and inefficient AWS API usage patterns, helping developers maintain secure and high-quality applications.

You can use Amazon Q Developer to review code and receive intelligent recommendations that improve code security and quality. Code reviewing capabilities are available through [Amazon Q code reviews](latest/qdeveloper-ug/security-scans.md) and [Amazon Inspector code scanning](../inspector/latest/user/scanning-lambda.md).

View the [Amazon Q Detector Library change log](detector-library/changelogs.md).

## Frequently asked questions

##### What is the Amazon Q Detector Library?

The Amazon QDetector Library is a resource that contains detailed information about Amazon Q's security and code quality detectors to help developers build secure and efficient applications on AWS. Each detection page within the Detector Library contains descriptions, noncompliant and compliant example code snippets, severities, and additional information that helps developers mitigate their risks (such as [CWE numbers](https://cwe.mitre.org/index.html)). The materials presented in the Amazon QDetector Library are intended to be a high-level summary of the service's capabilities but may not be inclusive of all detectors or their functionality.

##### How should I use the Amazon Q Detector Library?

You should review the Detector Library to get a deeper understanding of the capabilities of Amazon Q. Additionally, after reviewing your code resources in Amazon Q, you can use the detailed detection pages to help mitigate the findings you receive. You can also use this as an educational resource to help improve the overall security posture of your application and help ensure you are following AWS best practices.

##### How can I see Amazon Q code reviews in action?

You can see Amazon Q in action by using the [Amazon Q example detection repository](https://github.com/aws-samples/amazon-q-detectors) to review code. You can review code using Amazon Q's code reviewing capabilities available through [Amazon Q code review](latest/qdeveloper-ug/security-scans.md) or [Amazon Inspector code scanning](../inspector/latest/user/scanning-lambda.md). The repository contains many of the noncompliant code examples that appear in this Detector Library.

##### How often are the detectors updated?

The Amazon Q team is continually adding new detectors to help you keep your applications free from new, potentially harmful security vulnerabilities.

##### Do the detectors only find the specific example within each detection page?

No. Each detector can detect a wide range of different code defects. We included one noncompliant and compliant code example on each detection page (such as insecure cryptography) to help clarify the detection. However, each detector can find a range of defects in addition to the explicit code example shown on the detection page.

##### Which detectors are available in Amazon Q and Amazon Inspector Code Scanning?

Amazon Q code reviews features hundreds of Amazon Q's code security detectors, as well as hundreds of code quality detectors that can be enabled through additional configuration. For a list of supported languages, see the [Amazon Q Developer User Guide](latest/qdeveloper-ug/q-language-ide-support.md#security-scans-language-support).

For a list of Amazon Inspector code scanning supported languages, see the [Amazon Inspector User Guide](../inspector/latest/user/supported.md#supported-programming-languages-lambda-code). Amazon Inspector code scanning uses these detectors to scan Lambda functions.

##### How does Amazon Q determine what to include or exclude in a scan?

Before commencing a code review, Amazon Q applies filtering to ensure that only relevant customer code is reviewed. This ensures that the detected code issues are valuable to customers. As part of the filtering process, Amazon Q excludes unsupported languages, test code, and open source code.

##### Where can I find a change log with updates about Amazon Q detectors?

A history of changes to Amazon Q detectors, including additions and improvements, is published to theAmazon Q Detector Library change log. You can view the change log [here](detector-library/changelogs.md).

All content copied from https://docs.aws.amazon.com/.
