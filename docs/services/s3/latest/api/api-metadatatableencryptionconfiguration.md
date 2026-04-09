# MetadataTableEncryptionConfiguration

The encryption settings for an S3 Metadata journal table or inventory table configuration.

## Contents

**SseAlgorithm**

The encryption type specified for a metadata table. To specify server-side encryption with
AWS Key Management Service (AWS KMS) keys (SSE-KMS), use the `aws:kms` value. To specify server-side
encryption with Amazon S3 managed keys (SSE-S3), use the `AES256` value.

Type: String

Valid Values: `aws:kms | AES256`

Required: Yes

**KmsKeyArn**

If server-side encryption with AWS Key Management Service (AWS KMS) keys (SSE-KMS) is specified, you must also
specify the KMS key Amazon Resource Name (ARN). You must specify a customer-managed KMS key
that's located in the same Region as the general purpose bucket that corresponds to the metadata
table configuration.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/metadatatableencryptionconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/metadatatableencryptionconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/metadatatableencryptionconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetadataTableConfigurationResult

Metrics

All content copied from https://docs.aws.amazon.com/.
