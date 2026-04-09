# LogGroup

Represents a log group.

## Contents

**arn**

The Amazon Resource Name (ARN) of the log group. This version of the ARN includes a
trailing `:*` after the log group name.

Use this version to refer to the ARN in IAM policies when specifying
permissions for most API actions. The exception is when specifying permissions for [TagResource](api-tagresource.md), [UntagResource](api-untagresource.md),
and [ListTagsForResource](api-listtagsforresource.md). The permissions for those three actions require the ARN
version that doesn't include a trailing `:*`.

Type: String

Required: No

**bearerTokenAuthenticationEnabled**

Indicates whether bearer token authentication is enabled for this log group. When enabled,
bearer token authentication is allowed on operations until it is explicitly disabled.

Type: Boolean

Required: No

**creationTime**

The creation time of the log group, expressed as the number of milliseconds after Jan
1, 1970 00:00:00 UTC.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**dataProtectionStatus**

Displays whether this log group has a protection policy, or whether it had one in the
past. For more information, see [PutDataProtectionPolicy](api-putdataprotectionpolicy.md).

Type: String

Valid Values: `ACTIVATED | DELETED | ARCHIVED | DISABLED`

Required: No

**deletionProtectionEnabled**

Indicates whether deletion protection is enabled for this log group. When enabled,
deletion protection blocks all deletion operations until it is explicitly disabled.

Type: Boolean

Required: No

**inheritedProperties**

Displays all the properties that this log group has inherited from account-level
settings.

Type: Array of strings

Valid Values: `ACCOUNT_DATA_PROTECTION`

Required: No

**kmsKeyId**

The Amazon Resource Name (ARN) of the AWS KMS key to use when
encrypting log data.

Type: String

Length Constraints: Maximum length of 256.

Required: No

**logGroupArn**

The Amazon Resource Name (ARN) of the log group. This version of the ARN doesn't
include a trailing `:*` after the log group name.

Use this version to refer to the ARN in the following situations:

- In the `logGroupIdentifier` input field in many CloudWatch Logs
APIs.

- In the `resourceArn` field in tagging APIs

- In IAM policies, when specifying permissions for [TagResource](api-tagresource.md), [UntagResource](api-untagresource.md), and [ListTagsForResource](api-listtagsforresource.md).

Type: String

Required: No

**logGroupClass**

This specifies the log group class for this log group. There are three classes:

- The `Standard` log class supports all CloudWatch Logs features.

- The `Infrequent Access` log class supports a subset of CloudWatch Logs
features and incurs lower costs.

- Use the `Delivery` log class only for delivering AWS Lambda
logs to store in Amazon S3 or Amazon Data Firehose. Log events in log groups in
the Delivery class are kept in CloudWatch Logs for only one day. This log class doesn't
offer rich CloudWatch Logs capabilities such as CloudWatch Logs Insights
queries.

For details about the features supported by the Standard and Infrequent Access classes,
see [Log classes](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-log-classes.md)

Type: String

Valid Values: `STANDARD | INFREQUENT_ACCESS | DELIVERY`

Required: No

**logGroupName**

The name of the log group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: No

**metricFilterCount**

The number of metric filters.

Type: Integer

Required: No

**retentionInDays**

The number of days to retain the log events in the specified log group. Possible values
are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827, 2192, 2557,
2922, 3288, and 3653.

To set a log group so that its log events do not expire, use [DeleteRetentionPolicy](api-deleteretentionpolicy.md).

Type: Integer

Required: No

**storedBytes**

The number of bytes stored.

Type: Long

Valid Range: Minimum value of 0.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/loggroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/loggroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/loggroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogFieldType

LogGroupField

All content copied from https://docs.aws.amazon.com/.
