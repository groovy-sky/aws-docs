# PointInTimeRecoveryDescription

The description of the point in time settings applied to the table.

## Contents

###### Note

In the following list, the required parameters are described first.

**EarliestRestorableDateTime**

Specifies the earliest point in time you can restore your table to. You can restore
your table to any point in time during the last 35 days.

Type: Timestamp

Required: No

**LatestRestorableDateTime**

`LatestRestorableDateTime` is typically 5 minutes before the current time.

Type: Timestamp

Required: No

**PointInTimeRecoveryStatus**

The current state of point in time recovery:

- `ENABLED` \- Point in time recovery is enabled.

- `DISABLED` \- Point in time recovery is disabled.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**RecoveryPeriodInDays**

The number of preceding days for which continuous backups are taken and maintained.
Your table data is only recoverable to any point-in-time from within the configured
recovery period. This parameter is optional.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 35.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/pointintimerecoverydescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/pointintimerecoverydescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/pointintimerecoverydescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParameterizedStatement

PointInTimeRecoverySpecification

All content copied from https://docs.aws.amazon.com/.
