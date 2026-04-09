# S3ComputeObjectChecksumOperation

Directs the specified job to invoke the ComputeObjectChecksum operation on every object listed in the job's manifest.

## Contents

**ChecksumAlgorithm**

Indicates the algorithm that you want Amazon S3 to use to create the checksum. For more information, see [Checking object integrity](../userguide/checking-object-integrity.md) in the Amazon S3 User Guide.

Type: String

Valid Values: `CRC32 | CRC32C | CRC64NVME | MD5 | SHA1 | SHA256`

Required: No

**ChecksumType**

Indicates the checksum type that you want Amazon S3 to use to calculate the object's checksum value. For more information, see [Checking object integrity](../userguide/checking-object-integrity.md) in the Amazon S3 User Guide.

Type: String

Valid Values: `FULL_OBJECT | COMPOSITE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/s3computeobjectchecksumoperation.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/s3computeobjectchecksumoperation.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/s3computeobjectchecksumoperation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3BucketDestination

S3CopyObjectOperation

All content copied from https://docs.aws.amazon.com/.
