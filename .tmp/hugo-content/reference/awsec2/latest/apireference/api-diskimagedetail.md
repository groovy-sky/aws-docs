# DiskImageDetail

Describes a disk image.

## Contents

**Bytes**

The size of the disk image, in GiB.

Type: Long

Required: Yes

**Format**

The disk image format.

Type: String

Valid Values: `VMDK | RAW | VHD`

Required: Yes

**ImportManifestUrl**

A presigned URL for the import manifest stored in Amazon S3 and presented here as an Amazon S3 presigned URL.
For information about creating a presigned URL for an Amazon S3 object, read the "Query String Request Authentication
Alternative" section of the [Authenticating REST Requests](../../../../services/s3/latest/dev/restauthentication.md) topic in the _Amazon Simple Storage Service Developer_
_Guide_.

For information about the import manifest referenced by this API action, see [VM Import Manifest](manifest.md).

Type: String

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/diskimagedetail.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/diskimagedetail.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/diskimagedetail.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DiskImageDescription

DiskImageVolumeDescription
