This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::ContactFlowVersion

Creates a version for the specified customer-managed flow within the specified
instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::ContactFlowVersion",
  "Properties" : {
      "ContactFlowId" : String,
      "Description" : String
    }
}

```

### YAML

```yaml

Type: AWS::Connect::ContactFlowVersion
Properties:
  ContactFlowId: String
  Description: String

```

## Properties

`ContactFlowId`

The identifier of the flow.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]+:[0-9]{12}:instance/[-a-zA-Z0-9]+/contact-flow/[-a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the flow version.

_Required_: No

_Type_: String

_Maximum_: `500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`ContactFlowVersionARN`

The Amazon Resource Name (ARN) of the flow version.

`FlowContentSha256`

Indicates the checksum value of the flow content.

`Version`

The identifier of the flow version.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Connect::ContactFlowModuleVersion

AWS::Connect::DataTable
