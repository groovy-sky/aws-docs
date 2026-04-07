This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SNS::Subscription

The `AWS::SNS::Subscription` resource subscribes an endpoint to an Amazon SNS topic. For a subscription to be created, the owner of the endpoint must\`
confirm the subscription.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SNS::Subscription",
  "Properties" : {
      "DeliveryPolicy" : Json,
      "Endpoint" : String,
      "FilterPolicy" : Json,
      "FilterPolicyScope" : String,
      "Protocol" : String,
      "RawMessageDelivery" : Boolean,
      "RedrivePolicy" : Json,
      "Region" : String,
      "ReplayPolicy" : Json,
      "SubscriptionRoleArn" : String,
      "TopicArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::SNS::Subscription
Properties:
  DeliveryPolicy: Json
  Endpoint: String
  FilterPolicy: Json
  FilterPolicyScope: String
  Protocol: String
  RawMessageDelivery: Boolean
  RedrivePolicy: Json
  Region: String
  ReplayPolicy: Json
  SubscriptionRoleArn: String
  TopicArn: String

```

## Properties

`DeliveryPolicy`

The delivery policy JSON assigned to the subscription. Enables the subscriber to define
the message delivery retry strategy in the case of an HTTP/S endpoint subscribed to the
topic. For more information, see `
                            GetSubscriptionAttributes
                        ` in the _Amazon SNS API Reference_ and [Message delivery retries](https://docs.aws.amazon.com/sns/latest/dg/sns-message-delivery-retries.html)
in the _Amazon SNS Developer Guide_.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Endpoint`

The subscription's endpoint. The endpoint value depends on the protocol that you
specify. For more information, see the `Endpoint` parameter of the `
                            Subscribe
                        ` action in the _Amazon SNS API Reference_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FilterPolicy`

The filter policy JSON assigned to the subscription. Enables the subscriber to filter
out unwanted messages. For more information, see `
                            GetSubscriptionAttributes
                        ` in the _Amazon SNS API Reference_ and [Message filtering](https://docs.aws.amazon.com/sns/latest/dg/sns-message-filtering.html) in the _Amazon SNS Developer Guide_.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterPolicyScope`

This attribute lets you choose the filtering scope by using one of the following string
value types:

- `MessageAttributes` (default) - The filter is applied on the message
attributes.

- `MessageBody` \- The filter is applied on the message body.

###### Note

`Null` is not a valid value for `FilterPolicyScope`. To delete a
filter policy, delete the `FilterPolicy` property but keep
`FilterPolicyScope` property as is.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The subscription's protocol. For more information, see the `Protocol`
parameter of the `
                            Subscribe
                        ` action in the _Amazon SNS API Reference_.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RawMessageDelivery`

When set to `true`, enables raw message delivery. Raw messages don't contain
any JSON formatting and can be sent to Amazon SQS and HTTP/S endpoints. For more
information, see `
                            GetSubscriptionAttributes
                        ` in the _Amazon SNS API Reference_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedrivePolicy`

When specified, sends undeliverable messages to the specified Amazon SQS
dead-letter queue. Messages that can't be delivered due to client errors (for example, when
the subscribed endpoint is unreachable) or server errors (for example, when the service
that powers the subscribed endpoint becomes unavailable) are held in the dead-letter queue
for further analysis or reprocessing.

For more information about the redrive policy and dead-letter queues, see [Amazon\
SQS dead-letter queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html) in the _Amazon SQS Developer Guide_.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

For cross-region subscriptions, the region in which the topic resides.

If no region is specified, CloudFormation uses the region of the caller as the
default.

If you perform an update operation that only updates the `Region` property of
a `AWS::SNS::Subscription` resource, that operation will fail unless you are
either:

- Updating the `Region` from `NULL` to the caller
region.

- Updating the `Region` from the caller region to
`NULL`.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ReplayPolicy`

Specifies whether Amazon SNS resends the notification to the subscription when a
message's attribute changes.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubscriptionRoleArn`

This property applies only to Amazon Data Firehose delivery stream subscriptions.
Specify the ARN of the IAM role that has the following:

- Permission to write to the Amazon Data Firehose delivery stream

- Amazon SNS listed as a trusted entity

Specifying a valid ARN for this attribute is required for Firehose delivery
stream subscriptions. For more information, see [Fanout to Amazon Data Firehose\
delivery streams](https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html) in the _Amazon SNS Developer Guide._

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicArn`

The ARN of the topic to subscribe to.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the subscription's logical name. This attribute allows you to
reference the logical name of the subscription resource within the CloudFormation
template.

For example, if you have a subscription resource named 'MySubscription', you can use
'!Ref MySubscription' to refer to its logical name in other parts of the CloudFormation
template.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the ARN of the subscription created by the `AWS::SNS::Subscription`
resource.

## Examples

- [Create a subscription with mandatory attributes](#aws-resource-sns-subscription--examples--Create_a_subscription_with_mandatory_attributes)

- [Create a subscription with optional attributes](#aws-resource-sns-subscription--examples--Create_a_subscription_with_optional_attributes)

### Create a subscription with mandatory attributes

The following example creates a subscription with only an endpoint, protocol, and
topic ARN.

#### JSON

```json

"MySubscription" : {
  "Type" : "AWS::SNS::Subscription",
  "Properties" : {
    "Endpoint" : "test@example.com",
    "Protocol" : "email",
    "TopicArn" : { "Ref" : "MySNSTopic" }
  }
}
```

#### YAML

```yaml

MySubscription:
  Type: AWS::SNS::Subscription
  Properties:
    Endpoint: test@example.com
    Protocol: email
    TopicArn: !Ref 'MySNSTopic'
```

### Create a subscription with optional attributes

The following example creates a subscription with a filter policy, delivery
policy, and raw message delivery enabled.

###### Note

You can set subscription attributes only on standalone Amazon SNS
subscriptions (not on subscriptions nested in topics).

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  CarSalesTopic:
    Type: 'AWS::SNS::Topic'
  ERPSubscription:
    Type: 'AWS::SNS::Subscription'
    Properties:
      TopicArn: !Ref CarSalesTopic
      Endpoint: !GetAtt
        - ERPIntegrationQueue
        - Arn
      Protocol: sqs
      RawMessageDelivery: 'true'
  CRMSubscription:
    Type: 'AWS::SNS::Subscription'
    Properties:
      TopicArn: !Ref CarSalesTopic
      Endpoint: !GetAtt
        - CRMIntegrationQueue
        - Arn
      Protocol: sqs
      RawMessageDelivery: 'true'
      FilterPolicy:
        buyer-class:
          - vip
  SCMSubscription:
    Type: 'AWS::SNS::Subscription'
    Properties:
      TopicArn: !Ref CarSalesTopic
      Endpoint: !Ref myHttpEndpoint
      Protocol: https
      DeliveryPolicy:
        healthyRetryPolicy:
          numRetries: 20
          minDelayTarget: 10
          maxDelayTarget: 30
          numMinDelayRetries: 3
          numMaxDelayRetries: 17
          numNoDelayRetries: 0
          backoffFunction: exponential
  ERPIntegrationQueue:
    Type: 'AWS::SQS::Queue'
    Properties: {}
  CRMIntegrationQueue:
    Type: 'AWS::SQS::Queue'
    Properties: {}
Parameters:
  myHttpEndpoint:
    Type: String
```

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "CarSalesTopic": {
            "Type": "AWS::SNS::Topic"
        },
        "ERPSubscription": {
            "Type": "AWS::SNS::Subscription",
            "Properties": {
                "TopicArn": {
                    "Ref": "CarSalesTopic"
                },
                "Endpoint": {
                    "Fn::GetAtt": ["ERPIntegrationQueue", "Arn"]
                },
                "Protocol": "sqs",
                "RawMessageDelivery": "true"
            }
        },
        "CRMSubscription": {
            "Type": "AWS::SNS::Subscription",
            "Properties": {
                "TopicArn": {
                    "Ref": "CarSalesTopic"
                },
                "Endpoint": {
                    "Fn::GetAtt": ["CRMIntegrationQueue", "Arn"]
                },
                "Protocol": "sqs",
                "RawMessageDelivery": "true",
                "FilterPolicy": {
                    "buyer-class": [
                        "vip"
                    ]
                }
            }
        },
        "SCMSubscription": {
            "Type": "AWS::SNS::Subscription",
            "Properties": {
                "TopicArn": {
                    "Ref": "CarSalesTopic"
                },
                "Endpoint": {
                    "Ref": "myHttpEndpoint"
                },
                "Protocol": "https",
                "DeliveryPolicy": {
                    "healthyRetryPolicy": {
                        "numRetries": 20,
                        "minDelayTarget": 10,
                        "maxDelayTarget": 30,
                        "numMinDelayRetries": 3,
                        "numMaxDelayRetries": 17,
                        "numNoDelayRetries": 0,
                        "backoffFunction": "exponential"
                    }
                }
            }
        },
        "ERPIntegrationQueue": {
            "Type": "AWS::SQS::Queue",
            "Properties": {}
        },
        "CRMIntegrationQueue": {
            "Type": "AWS::SQS::Queue",
            "Properties": {}
        }
    },
    "Parameters": {
        "myHttpEndpoint": {
            "Type": "String"
        }
    }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Simple Notification Service

AWS::SNS::Topic
