# ImageUsageReport

The configuration and status of an image usage report.

## Contents

**AccountIdSet.N**

The IDs of the AWS accounts that were specified when the report was created.

Type: Array of strings

Required: No

**creationTime**

The date and time when the report was created.

Type: Timestamp

Required: No

**expirationTime**

The date and time when Amazon EC2 will delete the report (30 days after the report was
created).

Type: Timestamp

Required: No

**imageId**

The ID of the image that was specified when the report was created.

Type: String

Required: No

**reportId**

The ID of the report.

Type: String

Required: No

**ResourceTypeSet.N**

The resource types that were specified when the report was created.

Type: Array of [ImageUsageResourceType](api-imageusageresourcetype.md) objects

Required: No

**state**

The current state of the report. Possible values:

- `available` \- The report is available to view.

- `pending` \- The report is being created and not available to view.

- `error` \- The report could not be created.

Type: String

Required: No

**stateReason**

Provides additional details when the report is in an `error` state.

Type: String

Required: No

**TagSet.N**

Any tags assigned to the report.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/imageusagereport.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/imageusagereport.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/imageusagereport.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ImageReference

ImageUsageReportEntry
