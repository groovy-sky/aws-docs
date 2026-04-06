![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](https://docs.aws.amazon.com/amazonq/detector-library) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

# Amazon Q  Detector Library

## Trained on decades of knowledge and experience across millions of code reviews

[![Java](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/java.3b781738.svg)Java](https://docs.aws.amazon.com/amazonq/detector-library/java) [![Python](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/python.40e7bbb3.svg)Python](https://docs.aws.amazon.com/amazonq/detector-library/python) [![JavaScript](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/javascript.f62769e6.svg)JavaScript](https://docs.aws.amazon.com/amazonq/detector-library/javascript) [![JSX](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/jsx.041b6532.svg)JSX](https://docs.aws.amazon.com/amazonq/detector-library/jsx) [![TypeScript](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/typescript.17fda397.svg)TypeScript](https://docs.aws.amazon.com/amazonq/detector-library/typescript) [![C#](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/csharp.ee9c49fb.svg)C#](https://docs.aws.amazon.com/amazonq/detector-library/csharp) [![CloudFormation](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/cloudformation.0125b063.svg)CloudFormation](https://docs.aws.amazon.com/amazonq/detector-library/cloudformation) [![Terraform](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/terraform.2e7e71e7.svg)Terraform](https://docs.aws.amazon.com/amazonq/detector-library/terraform) [![Go](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/go.dad1aa01.svg)Go](https://docs.aws.amazon.com/amazonq/detector-library/go) [![Ruby](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/ruby.6896219f.svg)Ruby](https://docs.aws.amazon.com/amazonq/detector-library/ruby) [![C](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/c.c9a38a87.svg)C](https://docs.aws.amazon.com/amazonq/detector-library/c) [![C++](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/cpp.c908ae3d.svg)C++](https://docs.aws.amazon.com/amazonq/detector-library/cpp) [![PHP](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/php.0357268e.svg)PHP](https://docs.aws.amazon.com/amazonq/detector-library/php) [![Kotlin](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/kotlin.1bb7dd47.svg)Kotlin](https://docs.aws.amazon.com/amazonq/detector-library/kotlin) [![Scala](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/scala.f104853e.svg)Scala](https://docs.aws.amazon.com/amazonq/detector-library/scala) [![Shell](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/shell.c58a5629.svg)Shell](https://docs.aws.amazon.com/amazonq/detector-library/shell)

## About Amazon Q Detector Library

The Amazon Q Detector Library describes the detectors used during code reviews to identify security and quality issues in code. Detectors contain rules that are used to identify critical security vulnerabilities like OWASP Top 10 and CWE Top 25 issues, including secrets exposure and package dependency vulnerabilities. They also detect code quality concerns such as IaC best practices and inefficient AWS API usage patterns, helping developers maintain secure and high-quality applications.

You can use Amazon Q Developer to review code and receive intelligent recommendations that improve code security and quality. Code reviewing capabilities are available through [Amazon Q code reviews](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/security-scans.html) and [Amazon Inspector code scanning](https://docs.aws.amazon.com/inspector/latest/user/scanning-lambda.html).

View the [Amazon Q Detector Library change log](https://docs.aws.amazon.com/amazonq/detector-library/changelogs).

## Frequently asked questions

##### What is the Amazon Q Detector Library?

The Amazon QDetector Library is a resource that contains detailed information about Amazon Q's security and code quality detectors to help developers build secure and efficient applications on AWS. Each detection page within the Detector Library contains descriptions, noncompliant and compliant example code snippets, severities, and additional information that helps developers mitigate their risks (such as [CWE numbers](https://cwe.mitre.org/index.html)). The materials presented in the Amazon QDetector Library are intended to be a high-level summary of the service's capabilities but may not be inclusive of all detectors or their functionality.

##### How should I use the Amazon Q Detector Library?

You should review the Detector Library to get a deeper understanding of the capabilities of Amazon Q. Additionally, after reviewing your code resources in Amazon Q, you can use the detailed detection pages to help mitigate the findings you receive. You can also use this as an educational resource to help improve the overall security posture of your application and help ensure you are following AWS best practices.

##### How can I see Amazon Q code reviews in action?

You can see Amazon Q in action by using the [Amazon Q example detection repository](https://github.com/aws-samples/amazon-q-detectors) to review code. You can review code using Amazon Q's code reviewing capabilities available through [Amazon Q code review](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/security-scans.html) or [Amazon Inspector code scanning](https://docs.aws.amazon.com/inspector/latest/user/scanning-lambda.html). The repository contains many of the noncompliant code examples that appear in this Detector Library.

##### How often are the detectors updated?

The Amazon Q team is continually adding new detectors to help you keep your applications free from new, potentially harmful security vulnerabilities.

##### Do the detectors only find the specific example within each detection page?

No. Each detector can detect a wide range of different code defects. We included one noncompliant and compliant code example on each detection page (such as insecure cryptography) to help clarify the detection. However, each detector can find a range of defects in addition to the explicit code example shown on the detection page.

##### Which detectors are available in Amazon Q and Amazon Inspector Code Scanning?

Amazon Q code reviews features hundreds of Amazon Q's code security detectors, as well as hundreds of code quality detectors that can be enabled through additional configuration. For a list of supported languages, see the [Amazon Q Developer User Guide](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-language-ide-support.html#security-scans-language-support).

For a list of Amazon Inspector code scanning supported languages, see the [Amazon Inspector User Guide](https://docs.aws.amazon.com/inspector/latest/user/supported.html#supported-programming-languages-lambda-code). Amazon Inspector code scanning uses these detectors to scan Lambda functions.

##### How does Amazon Q determine what to include or exclude in a scan?

Before commencing a code review, Amazon Q applies filtering to ensure that only relevant customer code is reviewed. This ensures that the detected code issues are valuable to customers. As part of the filtering process, Amazon Q excludes unsupported languages, test code, and open source code.

##### Where can I find a change log with updates about Amazon Q detectors?

A history of changes to Amazon Q detectors, including additions and improvements, is published to theAmazon Q Detector Library change log. You can view the change log [here](https://docs.aws.amazon.com/amazonq/detector-library/changelogs).
