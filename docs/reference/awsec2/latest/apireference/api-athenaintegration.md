# AthenaIntegration

Describes integration options for Amazon Athena.

## Contents

**IntegrationResultS3DestinationArn**

The location in Amazon S3 to store the generated CloudFormation template.

Type: String

Required: Yes

**PartitionLoadFrequency**

The schedule for adding new partitions to the table.

Type: String

Valid Values: `none | daily | weekly | monthly`

Required: Yes

**PartitionEndDate**

The end date for the partition.

Type: Timestamp

Required: No

**PartitionStartDate**

The start date for the partition.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/athenaintegration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/athenaintegration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/athenaintegration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AssociationStatus

AttachmentEnaSrdSpecification
