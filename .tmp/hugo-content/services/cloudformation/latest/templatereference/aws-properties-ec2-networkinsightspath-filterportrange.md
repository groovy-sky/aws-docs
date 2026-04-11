This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkInsightsPath FilterPortRange

Describes a port range.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FromPort" : Integer,
  "ToPort" : Integer
}

```

### YAML

```yaml

  FromPort: Integer
  ToPort: Integer

```

## Properties

`FromPort`

The first port in the range.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ToPort`

The last port in the range.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::NetworkInsightsPath

PathFilter

All content copied from https://docs.aws.amazon.com/.
