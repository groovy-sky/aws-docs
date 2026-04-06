# KmsGrantConfiguration

A proposed grant configuration for a KMS key. For more information, see [CreateGrant](https://docs.aws.amazon.com/kms/latest/APIReference/API_CreateGrant.html).

## Contents

**granteePrincipal**

The principal that is given permission to perform the operations that the grant
permits.

Type: String

Required: Yes

**issuingAccount**

The AWS account under which the grant was issued. The account is used to propose
AWS KMS grants issued by accounts other than the owner of the key.

Type: String

Required: Yes

**operations**

A list of operations that the grant permits.

Type: Array of strings

Valid Values: `CreateGrant | Decrypt | DescribeKey | Encrypt | GenerateDataKey | GenerateDataKeyPair | GenerateDataKeyPairWithoutPlaintext | GenerateDataKeyWithoutPlaintext | GetPublicKey | ReEncryptFrom | ReEncryptTo | RetireGrant | Sign | Verify`

Required: Yes

**constraints**

Use this structure to propose allowing [cryptographic\
operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) in the grant only when the operation request includes the specified
[encryption\
context](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context).

Type: [KmsGrantConstraints](api-kmsgrantconstraints.md) object

Required: No

**retiringPrincipal**

The principal that is given permission to retire the grant by using [RetireGrant](https://docs.aws.amazon.com/kms/latest/APIReference/API_RetireGrant.html) operation.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/KmsGrantConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/KmsGrantConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/KmsGrantConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

JobError

KmsGrantConstraints
