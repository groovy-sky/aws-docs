This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::ContactFlowModuleAlias

Creates a named alias that points to a specific version of a contact flow module.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::ContactFlowModuleAlias",
  "Properties" : {
      "ContactFlowModuleId" : String,
      "ContactFlowModuleVersion" : Integer,
      "Description" : String,
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::Connect::ContactFlowModuleAlias
Properties:
  ContactFlowModuleId: String
  ContactFlowModuleVersion: Integer
  Description: String
  Name: String

```

## Properties

`ContactFlowModuleId`

The identifier of the flow module.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]+:[0-9]{12}:instance/[-a-zA-Z0-9]+/flow-module/[-a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContactFlowModuleVersion`

The version of the flow module.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the alias.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the alias.

_Required_: Yes

_Type_: String

_Pattern_: `^([$0-9a-zA-Z][_-]?)+$`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`AliasId`

The identifier of the alias.

`ContactFlowModuleAliasARN`

Property description not available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::Connect::ContactFlowModuleVersion
