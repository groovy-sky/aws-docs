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

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/certificatevalidationrecord.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/certificatevalidationrecord.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/certificatevalidationrecord.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoScalingConfigurationSummary

CodeConfiguration

All content copied from https://docs.aws.amazon.com/.
