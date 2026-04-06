# KmsGrantConstraints

Use this structure to propose allowing [cryptographic\
operations](../../../../services/kms/latest/developerguide/concepts.md#cryptographic-operations) in the grant only when the operation request includes the specified
[encryption\
context](../../../../services/kms/latest/developerguide/concepts.md#encrypt_context). You can specify only one type of encryption context. An empty map is
treated as not specified. For more information, see [GrantConstraints](../../../kms/latest/apireference/api-grantconstraints.md).

## Contents

**encryptionContextEquals**

A list of key-value pairs that must match the encryption context in the [cryptographic\
operation](../../../../services/kms/latest/developerguide/concepts.md#cryptographic-operations) request. The grant allows the operation only when the encryption
context in the request is the same as the encryption context specified in this
constraint.

Type: String to string map

Required: No

**encryptionContextSubset**

A list of key-value pairs that must be included in the encryption context of the [cryptographic\
operation](../../../../services/kms/latest/developerguide/concepts.md#cryptographic-operations) request. The grant allows the cryptographic operation only when the
encryption context in the request includes the key-value pairs specified in this
constraint, although it can include additional key-value pairs.

Type: String to string map

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/kmsgrantconstraints.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/kmsgrantconstraints.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/kmsgrantconstraints.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

KmsGrantConfiguration

KmsKeyConfiguration
