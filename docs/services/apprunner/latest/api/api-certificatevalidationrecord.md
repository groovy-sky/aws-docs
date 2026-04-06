# CertificateValidationRecord

Describes a certificate CNAME record to add to your DNS. For more information, see [AssociateCustomDomain](api-associatecustomdomain.md).

## Contents

**Name**

The certificate CNAME record name.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

Required: No

**Status**

The current state of the certificate CNAME record validation. It should change to `SUCCESS` after App Runner completes validation with your
DNS.

Type: String

Valid Values: `PENDING_VALIDATION | SUCCESS | FAILED`

Required: No

**Type**

The record type, always `CNAME`.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

Required: No

**Value**

The certificate CNAME record value.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apprunner-2020-05-15/CertificateValidationRecord)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apprunner-2020-05-15/CertificateValidationRecord)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apprunner-2020-05-15/CertificateValidationRecord)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AutoScalingConfigurationSummary

CodeConfiguration
