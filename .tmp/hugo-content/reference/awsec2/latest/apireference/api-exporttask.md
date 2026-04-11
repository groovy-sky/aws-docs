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

Type: [ExportToS3Task](api-exporttos3task.md) object

Required: No

**instanceExport**

Information about the instance to export.

Type: [InstanceExportDetails](api-instanceexportdetails.md) object

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

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/exporttask.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/exporttask.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/exporttask.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ExportImageTask

ExportTaskS3Location
