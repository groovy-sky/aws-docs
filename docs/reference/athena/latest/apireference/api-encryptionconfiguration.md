# EncryptionConfiguration

If query and calculation results are encrypted in Amazon S3, indicates the
encryption option used (for example, `SSE_KMS` or `CSE_KMS`) and
key information.

## Contents

**EncryptionOption**

Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys ( `SSE_S3`), server-side encryption with KMS-managed keys
( `SSE_KMS`), or client-side encryption with KMS-managed keys
( `CSE_KMS`) is used.

If a query runs in a workgroup and the workgroup overrides client-side settings, then
the workgroup's setting for encryption is used. It specifies whether query results must
be encrypted, for all queries that run in this workgroup.

Type: String

Valid Values: `SSE_S3 | SSE_KMS | CSE_KMS`

Required: Yes

**KmsKey**

For `SSE_KMS` and `CSE_KMS`, this is the KMS key ARN or
ID.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/encryptionconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/encryptionconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/encryptionconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Datum

EngineConfiguration

All content copied from https://docs.aws.amazon.com/.
