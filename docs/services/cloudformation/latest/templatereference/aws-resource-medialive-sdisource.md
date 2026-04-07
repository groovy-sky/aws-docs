This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::SdiSource

The `AWS::MediaLive::SdiSource` resource Property description not available. for MediaLive.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaLive::SdiSource",
  "Properties" : {
      "Mode" : String,
      "Name" : String,
      "Tags" : [ Tags, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::MediaLive::SdiSource
Properties:
  Mode: String
  Name: String
  Tags:
    - Tags
  Type: String

```

## Properties

`Mode`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `QUADRANT | INTERLEAVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-sdisource-tags.html) of [Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-sdisource-tags.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Property description not available.

_Required_: Yes

_Type_: String

_Allowed values_: `SINGLE | QUAD`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The unique ARN of the SDI source.

`Id`

The unique identifier of the SDI source.

`Inputs`

The list of inputs currently using this SDI source.

`State`

The current state of the SDI source.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tags

Tags
