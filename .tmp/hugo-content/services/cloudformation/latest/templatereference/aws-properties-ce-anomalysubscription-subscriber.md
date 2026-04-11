This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CE::AnomalySubscription Subscriber

The recipient of `AnomalySubscription` notifications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Address" : String,
  "Status" : String,
  "Type" : String
}

```

### YAML

```yaml

  Address: String
  Status: String
  Type: String

```

## Properties

`Address`

The email address or SNS Topic Amazon Resource Name (ARN), depending on the
`Type`.

_Required_: Yes

_Type_: String

_Pattern_: `(^[a-zA-Z0-9.!#$%&'*+=?^_‘{|}~-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$)|(^arn:(aws[a-zA-Z-]*):sns:[a-zA-Z0-9-]+:[0-9]{12}:[a-zA-Z0-9_-]+(\.fifo)?$)`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Indicates if the subscriber accepts the notifications.

_Required_: No

_Type_: String

_Allowed values_: `CONFIRMED | DECLINED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The notification delivery channel.

_Required_: Yes

_Type_: String

_Allowed values_: `EMAIL | SNS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceTag

AWS::CE::CostCategory

All content copied from https://docs.aws.amazon.com/.
