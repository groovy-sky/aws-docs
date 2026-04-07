# ExportImageTask

Describes an export image task.

## Contents

**description**

A description of the image being exported.

Type: String

Required: No

**exportImageTaskId**

The ID of the export image task.

Type: String

Required: No

**imageId**

The ID of the image.

Type: String

Required: No

**progress**

The percent complete of the export image task.

Type: String

Required: No

**s3ExportLocation**

Information about the destination Amazon S3 bucket.

Type: [ExportTaskS3Location](api-exporttasks3location.md) object

Required: No

**status**

The status of the export image task. The possible values are `active`, `completed`,
`deleting`, and `deleted`.

Type: String

Required: No

**statusMessage**

The status message for the export image task.

Type: String

Required: No

**TagSet.N**

Any tags assigned to the export image task.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/exportimagetask.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/exportimagetask.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/exportimagetask.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Explanation

ExportTask
