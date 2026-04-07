# ExportToS3TaskSpecification

Describes an export instance task.

## Contents

**ContainerFormat**

The container format used to combine disk images with metadata (such as OVF). If absent, only the disk image is
exported.

Type: String

Valid Values: `ova`

Required: No

**DiskImageFormat**

The format for the exported image.

Type: String

Valid Values: `VMDK | RAW | VHD`

Required: No

**S3Bucket**

The Amazon S3 bucket for the destination image. The destination bucket must exist and have
an access control list (ACL) attached that specifies the Region-specific canonical account ID for
the `Grantee`. For more information about the ACL to your S3 bucket, see [Prerequisites](../../../../services/vm-import/latest/userguide/vmexport.md#vmexport-prerequisites) in the VM Import/Export User Guide.

Type: String

Required: No

**S3Prefix**

The image is written to a single object in the Amazon S3 bucket at the S3 key s3prefix +
exportTaskId + '.' + diskImageFormat.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/exporttos3taskspecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/exporttos3taskspecification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/exporttos3taskspecification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ExportToS3Task

ExternalAuthorityConfiguration
