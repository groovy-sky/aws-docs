# S3ManifestOutputLocation

Location details for where the generated manifest should be written.

## Contents

**Bucket**

The bucket ARN the generated manifest should be written to.

###### Note

**Directory buckets** \- Directory buckets aren't supported
as the buckets to store the generated manifest.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `arn:[^:]+:s3:.*`

Required: Yes

**ManifestFormat**

The format of the generated manifest.

Type: String

Valid Values: `S3InventoryReport_CSV_20211130`

Required: Yes

**ExpectedManifestBucketOwner**

The Account ID that owns the bucket the generated manifest is written to.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: No

**ManifestEncryption**

Specifies what encryption should be used when the generated manifest objects are
written.

Type: [GeneratedManifestEncryption](api-control-generatedmanifestencryption.md) data type

Required: No

**ManifestPrefix**

Prefix identifying one or more objects to which the manifest applies.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/s3manifestoutputlocation.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/s3manifestoutputlocation.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/s3manifestoutputlocation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3JobManifestGenerator

S3ObjectLockLegalHold

All content copied from https://docs.aws.amazon.com/.
