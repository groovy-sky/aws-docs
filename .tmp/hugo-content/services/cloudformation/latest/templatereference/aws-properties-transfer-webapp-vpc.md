This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::WebApp Vpc

The `Vpc` property type specifies Property description not available. for an [AWS::Transfer::WebApp](aws-resource-transfer-webapp.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroupIds" : [ String, ... ],
  "SubnetIds" : [ String, ... ],
  "VpcId" : String
}

```

### YAML

```yaml

  SecurityGroupIds:
    - String
  SubnetIds:
    - String
  VpcId: String

```

## Properties

`SecurityGroupIds`

Property description not available.

_Required_: No

_Type_: Array of String

_Minimum_: `11`

_Maximum_: `20 | 10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

Property description not available.

_Required_: No

_Type_: Array of String

_Minimum_: `15`

_Maximum_: `24 | 10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^vpc-[0-9a-f]{8,17}$`

_Minimum_: `12`

_Maximum_: `21`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

WebAppCustomization

All content copied from https://docs.aws.amazon.com/.
