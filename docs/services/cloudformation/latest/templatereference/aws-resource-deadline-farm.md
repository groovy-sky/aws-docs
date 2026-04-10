This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Farm

Creates a farm to allow space for queues and fleets. Farms are the space where the
components of your renders gather and are pieced together in the cloud. Farms contain
budgets and allow you to enforce permissions. Deadline Cloud farms are a useful container for
large projects.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Deadline::Farm",
  "Properties" : {
      "CostScaleFactor" : Number,
      "Description" : String,
      "DisplayName" : String,
      "KmsKeyArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Deadline::Farm
Properties:
  CostScaleFactor: Number
  Description: String
  DisplayName: String
  KmsKeyArn: String
  Tags:
    - Tag

```

## Properties

`CostScaleFactor`

Property description not available.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the farm that helps identify what the farm is used for.

###### Important

This field can store any content. Escape or encode this content before displaying it
on a webpage or any other system that might interpret the content of this field.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The display name of the farm.

###### Important

This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The ARN for the KMS key.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[-a-z]*:kms:.*:key/.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to add to your farm. Each tag consists of a tag key and a tag value. Tag keys and values are both required, but tag values can be empty strings.

_Required_: No

_Type_: Array of [Tag](aws-properties-deadline-farm-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the farm.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) assigned to the farm.

`FarmId`

The farm ID.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Deadline Cloud

Tag

All content copied from https://docs.aws.amazon.com/.
