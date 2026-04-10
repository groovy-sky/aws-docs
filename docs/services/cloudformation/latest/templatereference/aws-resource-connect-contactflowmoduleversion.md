This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::ContactFlowModuleVersion

Creates an immutable snapshot of a contact flow module, preserving its content and settings at a specific point
in time for version control and rollback capabilities.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::ContactFlowModuleVersion",
  "Properties" : {
      "ContactFlowModuleId" : String,
      "Description" : String
    }
}

```

### YAML

```yaml

Type: AWS::Connect::ContactFlowModuleVersion
Properties:
  ContactFlowModuleId: String
  Description: String

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

`Description`

Property description not available.

_Required_: No

_Type_: String

_Maximum_: `500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`ContactFlowModuleVersionARN`

Property description not available.

`FlowModuleContentSha256`

Property description not available.

`Version`

The version of the flow module.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Connect::ContactFlowModuleAlias

AWS::Connect::ContactFlowVersion

All content copied from https://docs.aws.amazon.com/.
