# ContinuousBackupsDescription

Represents the continuous backups and point in time recovery settings on the
table.

## Contents

###### Note

In the following list, the required parameters are described first.

**ContinuousBackupsStatus**

`ContinuousBackupsStatus` can be one of the following states: ENABLED,
DISABLED

Type: String

Valid Values: `ENABLED | DISABLED`

Required: Yes

**PointInTimeRecoveryDescription**

The description of the point in time recovery settings applied to the table.

Type: [PointInTimeRecoveryDescription](api-pointintimerecoverydescription.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/continuousbackupsdescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/continuousbackupsdescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/continuousbackupsdescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConsumedCapacity

ContributorInsightsSummary

All content copied from https://docs.aws.amazon.com/.
