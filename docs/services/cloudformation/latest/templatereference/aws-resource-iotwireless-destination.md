This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::Destination

Creates a new destination that maps a device message to an AWS IoT
rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTWireless::Destination",
  "Properties" : {
      "Description" : String,
      "Expression" : String,
      "ExpressionType" : String,
      "Name" : String,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTWireless::Destination
Properties:
  Description: String
  Expression: String
  ExpressionType: String
  Name: String
  RoleArn: String
  Tags:
    - Tag

```

## Properties

`Description`

The description of the new resource. Maximum length is 2048 characters.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expression`

The rule name to send messages to.

_Required_: Yes

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExpressionType`

The type of value in `Expression`.

_Required_: Yes

_Type_: String

_Allowed values_: `RuleName | MqttTopic | SnsTopic`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the new resource.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The ARN of the IAM Role that authorizes the destination.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags are an array of key-value pairs to attach to the specified resource. Tags can
have a minimum of 0 and a maximum of 50 items.

_Required_: No

_Type_: Array of [Tag](aws-properties-iotwireless-destination-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Destination name.

### Fn::GetAtt

`Arn`

The ARN of the destination created.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS IoT Wireless

Tag

All content copied from https://docs.aws.amazon.com/.
