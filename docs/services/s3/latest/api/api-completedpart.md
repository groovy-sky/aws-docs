# CompletedPart

Details of the parts that were uploaded.

## Contents

**ChecksumCRC32**

The Base64 encoded, 32-bit `CRC32` checksum of the part. This checksum is present if the
multipart upload request was created with the `CRC32` checksum algorithm. For more
information, see [Checking object integrity](../userguide/checking-object-integrity.md) in
the _Amazon S3 User Guide_.

Type: String

Required: No

**ChecksumCRC32C**

The Base64 encoded, 32-bit `CRC32C` checksum of the part. This checksum is present if the
multipart upload request was created with the `CRC32C` checksum algorithm. For more
information, see [Checking object integrity](../userguide/checking-object-integrity.md) in
the _Amazon S3 User Guide_.

Type: String

Required: No

**ChecksumCRC64NVME**

The Base64 encoded, 64-bit `CRC64NVME` checksum of the part. This checksum is present if
the multipart upload request was created with the `CRC64NVME` checksum algorithm to the
uploaded object). For more information, see [Checking object integrity](../userguide/checking-object-integrity.md) in
the _Amazon S3 User Guide_.

Type: String

Required: No

**ChecksumSHA1**

The Base64 encoded, 160-bit `SHA1` checksum of the part. This checksum is present if the
multipart upload request was created with the `SHA1` checksum algorithm. For more
information, see [Checking object integrity](../userguide/checking-object-integrity.md) in
the _Amazon S3 User Guide_.

Type: String

Required: No

**ChecksumSHA256**

The Base64 encoded, 256-bit `SHA256` checksum of the part. This checksum is present if
the multipart upload request was created with the `SHA256` checksum algorithm. For more
information, see [Checking object integrity](../userguide/checking-object-integrity.md) in
the _Amazon S3 User Guide_.

Type: String

Required: No

**ETag**

Entity tag returned when the part was uploaded.

Type: String

Required: No

**PartNumber**

Part number that identifies the part. This is a positive integer between 1 and 10,000.

###### Note

- **General purpose buckets** \- In
`CompleteMultipartUpload`, when a additional checksum (including
`x-amz-checksum-crc32`, `x-amz-checksum-crc32c`,
`x-amz-checksum-sha1`, or `x-amz-checksum-sha256`) is applied to each
part, the `PartNumber` must start at 1 and the part numbers must be consecutive.
Otherwise, Amazon S3 generates an HTTP `400 Bad Request` status code and an
`InvalidPartOrder` error code.

- **Directory buckets** \- In
`CompleteMultipartUpload`, the `PartNumber` must start at 1 and the part
numbers must be consecutive.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/completedpart.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/completedpart.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/completedpart.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CompletedMultipartUpload

Condition

All content copied from https://docs.aws.amazon.com/.
