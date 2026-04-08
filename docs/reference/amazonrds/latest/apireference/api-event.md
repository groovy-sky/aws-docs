# Event

This data type is used as a response element in the [DescribeEvents](api-describeevents.md) action.

## Contents

###### Note

In the following list, the required parameters are described first.

**Date**

Specifies the date and time of the event.

Type: Timestamp

Required: No

**EventCategories.EventCategory.N**

Specifies the category for the event.

Type: Array of strings

Required: No

**Message**

Provides the text of this event.

Type: String

Required: No

**SourceArn**

The Amazon Resource Name (ARN) for the event.

Type: String

Required: No

**SourceIdentifier**

Provides the identifier for the source of the event.

Type: String

Required: No

**SourceType**

Specifies the source type for this event.

Type: String

Valid Values: `db-instance | db-parameter-group | db-security-group | db-snapshot | db-cluster | db-cluster-snapshot | custom-engine-version | db-proxy | blue-green-deployment | db-shard-group | zero-etl`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/event.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/event.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/event.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EngineDefaults

EventCategoriesMap

All content copied from https://docs.aws.amazon.com/.
