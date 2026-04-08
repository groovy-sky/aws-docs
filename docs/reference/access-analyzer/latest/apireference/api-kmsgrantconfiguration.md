# KmsGrantConfiguration

A proposed grant configuration for a KMS key. For more information, see [CreateGrant](../../../kms/latest/apireference/api-creategrant.md).

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
operations](../../../../services/kms/latest/developerguide/concepts-cryptographic-operations.md) in the grant only when the operation request includes the specified
[encryption\
context](../../../../services/kms/latest/developerguide/concepts-encrypt-context.md).

Type: [KmsGrantConstraints](api-kmsgrantconstraints.md) object

Required: No

**retiringPrincipal**

The principal that is given permission to retire the grant by using [RetireGrant](../../../kms/latest/apireference/api-retiregrant.md) operation.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/kmsgrantconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/kmsgrantconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/kmsgrantconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

JobError

KmsGrantConstraints
