# EventSubscription

Contains the results of a successful invocation of the `DescribeEventSubscriptions` action.

## Contents

###### Note

In the following list, the required parameters are described first.

**CustomerAwsId**

The AWS customer account associated with the RDS event notification subscription.

Type: String

Required: No

**CustSubscriptionId**

The RDS event notification subscription Id.

Type: String

Required: No

**Enabled**

Specifies whether the subscription is enabled. True indicates the subscription is enabled.

Type: Boolean

Required: No

**EventCategoriesList.EventCategory.N**

A list of event categories for the RDS event notification subscription.

Type: Array of strings

Required: No

**EventSubscriptionArn**

The Amazon Resource Name (ARN) for the event subscription.

Type: String

Required: No

**SnsTopicArn**

The topic ARN of the RDS event notification subscription.

Type: String

Required: No

**SourceIdsList.SourceId.N**

A list of source IDs for the RDS event notification subscription.

Type: Array of strings

Required: No

**SourceType**

The source type for the RDS event notification subscription.

Type: String

Required: No

**Status**

The status of the RDS event notification subscription.

Constraints:

Can be one of the following: creating \| modifying \| deleting \| active \| no-permission \| topic-not-exist

The status "no-permission" indicates that RDS no longer has permission to post to the SNS topic. The status "topic-not-exist" indicates that the topic was deleted after the subscription was created.

Type: String

Required: No

**SubscriptionCreationTime**

The time the RDS event notification subscription was created.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/eventsubscription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/eventsubscription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/eventsubscription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventCategoriesMap

ExportTask

All content copied from https://docs.aws.amazon.com/.
