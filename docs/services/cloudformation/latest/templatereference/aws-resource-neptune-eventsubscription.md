This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Neptune::EventSubscription

Creates an event notification subscription. This action requires a topic ARN (Amazon
Resource Name) created by either the Neptune console, the SNS console, or the SNS API. To
obtain an ARN with SNS, you must create a topic in Amazon SNS and subscribe to the topic. The
ARN is displayed in the SNS console.

You can specify the type of source (SourceType) you want to be notified of, provide a list
of Neptune sources (SourceIds) that triggers the events, and provide a list of event
categories (EventCategories) for events you want to be notified of. For example, you can
specify SourceType = db-instance, SourceIds = mydbinstance1, mydbinstance2 and EventCategories
= Availability, Backup.

If you specify both the SourceType and SourceIds, such as SourceType = db-instance and
SourceIdentifier = myDBInstance1, you are notified of all the db-instance events for the
specified source. If you specify a SourceType but do not specify a SourceIdentifier, you
receive notice of the events for that source type for all your Neptune sources. If you do not
specify either the SourceType nor the SourceIdentifier, you are notified of events generated
from all Neptune sources belonging to your customer account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Neptune::EventSubscription",
  "Properties" : {
      "Enabled" : Boolean,
      "EventCategories" : [ String, ... ],
      "SnsTopicArn" : String,
      "SourceIds" : [ String, ... ],
      "SourceType" : String,
      "SubscriptionName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Neptune::EventSubscription
Properties:
  Enabled: Boolean
  EventCategories:
    - String
  SnsTopicArn: String
  SourceIds:
    - String
  SourceType: String
  SubscriptionName: String
  Tags:
    - Tag

```

## Properties

`Enabled`

A Boolean value indicating if the subscription is enabled. True indicates the subscription
is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventCategories`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnsTopicArn`

The topic ARN of the event notification subscription.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceIds`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceType`

The source type for the event notification subscription.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubscriptionName`

Property description not available.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-neptune-eventsubscription-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
