# ExportTask

Describes an export instance task.

## Contents

**description**

A description of the resource being exported.

Type: String

Required: No

**exportTaskId**

The ID of the export task.

Type: String

Required: No

**exportToS3**

Information about the export task.

Type: [ExportToS3Task](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ExportToS3Task.html) object

Required: No

**instanceExport**

Information about the instance to export.

Type: [InstanceExportDetails](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceExportDetails.html) object

Required: No

**state**

The state of the export task.

Type: String

Valid Values: `active | cancelling | cancelled | completed`

Required: No

**statusMessage**

The status message related to the export task.

Type: String

Required: No

**TagSet.N**

The tags for the export task.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ExportTask)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ExportTask)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ExportTask)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ExportImageTask

ExportTaskS3Location
