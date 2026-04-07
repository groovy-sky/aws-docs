# ImageUsageReportEntry

A single entry in an image usage report, detailing how an image is being used by a
specific AWS account and resource type.

## Contents

**accountId**

The ID of the account that uses the image.

Type: String

Required: No

**imageId**

The ID of the image.

Type: String

Required: No

**reportCreationTime**

The date and time the report creation was initiated.

Type: Timestamp

Required: No

**reportId**

The ID of the report.

Type: String

Required: No

**resourceType**

The type of resource ( `ec2:Instance` or
`ec2:LaunchTemplate`).

Type: String

Required: No

**usageCount**

The number of times resources of this type reference this image in the account.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/imageusagereportentry.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/imageusagereportentry.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/imageusagereportentry.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ImageUsageReport

ImageUsageResourceType
