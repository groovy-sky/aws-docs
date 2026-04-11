This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConvert::Queue

The AWS::MediaConvert::Queue resource is an AWS Elemental MediaConvert resource type
that you can use to manage the resources that are available to your account for parallel
processing of jobs. For more information about queues, see [Working with AWS Elemental MediaConvert Queues](../../../mediaconvert/latest/ug/working-with-queues.md) in the _AWS Elemental MediaConvert User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaConvert::Queue",
  "Properties" : {
      "ConcurrentJobs" : Integer,
      "Description" : String,
      "Name" : String,
      "PricingPlan" : String,
      "Status" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaConvert::Queue
Properties:
  ConcurrentJobs: Integer
  Description: String
  Name: String
  PricingPlan: String
  Status: String
  Tags:
    - Tag

```

## Properties

`ConcurrentJobs`

Specify the maximum number of jobs your queue can process concurrently. For on-demand queues, the value you enter is constrained by your service quotas for Maximum concurrent jobs, per on-demand queue and Maximum concurrent jobs, per account. For reserved queues, specify the number of jobs you can process concurrently in your reservation plan instead.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

Optional. A description of the queue that you are creating.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the queue that you are creating.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PricingPlan`

When you use CloudFormation, you can create only on-demand queues. Therefore,
always set `PricingPlan` to the value "ON\_DEMAND" when declaring an
AWS::MediaConvert::Queue in your CloudFormation template.

To create a reserved queue, use the AWS Elemental MediaConvert console at
https://console.aws.amazon.com/mediaconvert to set up a contract. For more information,
see [Working with AWS Elemental MediaConvert Queues](../../../mediaconvert/latest/ug/working-with-queues.md) in the _AWS Elemental MediaConvert User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Initial state of the queue. Queues can be either ACTIVE or PAUSED. If you create a
paused queue, then jobs that you send to that queue won't begin.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of an `AWS::MediaConvert::Queue` resource to
the intrinsic `Ref` function, the function returns the name of the queue,
such as `Queue 2`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the queue, such as
`arn:aws:mediaconvert:us-west-2:123456789012`.

`Name`

The name of the queue, such as `Queue 2`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaConvert::Preset

Next

All content copied from https://docs.aws.amazon.com/.
