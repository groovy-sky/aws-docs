This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Queue

Contains information about a queue.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::Queue",
  "Properties" : {
      "AdditionalEmailAddresses" : [ EmailAddress, ... ],
      "Description" : String,
      "HoursOfOperationArn" : String,
      "InstanceArn" : String,
      "MaxContacts" : Integer,
      "Name" : String,
      "OutboundCallerConfig" : OutboundCallerConfig,
      "OutboundEmailConfig" : OutboundEmailConfig,
      "QuickConnectArns" : [ String, ... ],
      "Status" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Connect::Queue
Properties:
  AdditionalEmailAddresses:
    - EmailAddress
  Description: String
  HoursOfOperationArn: String
  InstanceArn: String
  MaxContacts: Integer
  Name: String
  OutboundCallerConfig:
    OutboundCallerConfig
  OutboundEmailConfig:
    OutboundEmailConfig
  QuickConnectArns:
    - String
  Status: String
  Tags:
    - Tag

```

## Properties

`AdditionalEmailAddresses`

Property description not available.

_Required_: No

_Type_: Array of [EmailAddress](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-queue-emailaddress.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the queue.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `250`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HoursOfOperationArn`

The Amazon Resource Name (ARN) of the hours of operation.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/operating-hours/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The identifier of the Amazon Connect instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxContacts`

The maximum number of contacts that can be in the queue before it is considered full.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the queue.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutboundCallerConfig`

The outbound caller ID name, number, and outbound whisper flow.

_Required_: No

_Type_: [OutboundCallerConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-queue-outboundcallerconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutboundEmailConfig`

The outbound email address ID for a specified queue.

_Required_: No

_Type_: [OutboundEmailConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-queue-outboundemailconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QuickConnectArns`

The Amazon Resource Names (ARN) of the of the quick connects available to agents who
are working the queue.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the queue.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource. For example, { "Tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-queue-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the queue name. For example:

`{ "Ref": "myQueueName" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`QueueArn`

The Amazon Resource Name (ARN) of the queue.

`Type`

The type of queue.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

EmailAddress
