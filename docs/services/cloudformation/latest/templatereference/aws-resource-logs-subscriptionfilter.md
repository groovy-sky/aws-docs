This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::SubscriptionFilter

The `AWS::Logs::SubscriptionFilter` resource specifies a subscription
filter and associates it with the specified log group. Subscription filters allow you to
subscribe to a real-time stream of log events and have them delivered to a specific
destination. Currently, the supported destinations are:

- An Amazon Kinesis data stream belonging to the same account as the
subscription filter, for same-account delivery.

- A logical destination that belongs to a different account, for
cross-account delivery.

- An Amazon Kinesis Firehose delivery stream that belongs to the same account
as the subscription filter, for same-account delivery.

- An AWS Lambda function that belongs to the same account as
the subscription filter, for same-account delivery.

There can be as many as two subscription filters associated with a log
group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Logs::SubscriptionFilter",
  "Properties" : {
      "ApplyOnTransformedLogs" : Boolean,
      "DestinationArn" : String,
      "Distribution" : String,
      "EmitSystemFields" : [ String, ... ],
      "FieldSelectionCriteria" : String,
      "FilterName" : String,
      "FilterPattern" : String,
      "LogGroupName" : String,
      "RoleArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Logs::SubscriptionFilter
Properties:
  ApplyOnTransformedLogs: Boolean
  DestinationArn: String
  Distribution: String
  EmitSystemFields:
    - String
  FieldSelectionCriteria: String
  FilterName: String
  FilterPattern: String
  LogGroupName: String
  RoleArn: String

```

## Properties

`ApplyOnTransformedLogs`

This parameter is valid only for log groups that have an active log transformer. For
more information about log transformers, see [PutTransformer](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutTransformer.html).

If this value is `true`, the subscription filter is applied on the
transformed version of the log events instead of the original ingested log
events.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationArn`

The Amazon Resource Name (ARN) of the destination.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Distribution`

The method used to distribute log data to the destination, which can be either random
or grouped by log stream.

_Required_: No

_Type_: String

_Allowed values_: `Random | ByLogStream`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmitSystemFields`

The list of system fields that are included in the log events sent to the subscription
destination. Returns the `emitSystemFields` value if it was specified when the
subscription filter was created.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldSelectionCriteria`

The filter expression that specifies which log events are processed by this subscription
filter based on system fields. Returns the `fieldSelectionCriteria` value if it was
specified when the subscription filter was created.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterName`

The name of the subscription filter.

_Required_: No

_Type_: String

_Pattern_: `[^:*]*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FilterPattern`

The filtering expressions that restrict what gets delivered to the destination
AWS resource. For more information about the filter pattern syntax,
see [Filter and Pattern\
Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroupName`

The log group to associate with the subscription filter. All log events that are
uploaded to this log group are filtered and delivered to the specified AWS resource if the filter pattern matches the log events.

_Required_: Yes

_Type_: String

_Pattern_: `[\.\-_/#A-Za-z0-9]+`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The ARN of an IAM role that grants CloudWatch Logs permissions to deliver
ingested log events to the destination stream. You don't need to provide the ARN when
you are working with a logical destination for cross-account delivery.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Create a Subscription Filter

The following example sends log events that are associated with the
`Root` user to a Kinesis data stream.

#### JSON

```json

"SubscriptionFilter" : {
  "Type" : "AWS::Logs::SubscriptionFilter",
  "Properties" : {
    "RoleArn" : { "Fn::GetAtt" : [ "CloudWatchIAMRole", "Arn" ] },
    "LogGroupName" : { "Ref" : "LogGroup" },
    "Distribution" : "Random",
    "FilterName" : "filterNameString",
    "FilterPattern" : "{$.userIdentity.type = Root}",
    "DestinationArn" : { "Fn::GetAtt" : [ "KinesisStream", "Arn" ] }
  }
}
```

#### YAML

```yaml

SubscriptionFilter:
  Type: AWS::Logs::SubscriptionFilter
  Properties:
    RoleArn:
      Fn::GetAtt:
        - "CloudWatchIAMRole"
        - "Arn"
    LogGroupName:
      Ref: "LogGroup"
    Distribution: "Random"
    FilterName: "filterNameString"
    FilterPattern: "{$.userIdentity.type = Root}"
    DestinationArn:
      Fn::GetAtt:
        - "KinesisStream"
        - "Arn"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TagsItems

AWS::Logs::Transformer
