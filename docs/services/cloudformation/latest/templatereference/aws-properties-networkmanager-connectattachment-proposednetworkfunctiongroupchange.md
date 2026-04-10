This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::ConnectAttachment ProposedNetworkFunctionGroupChange

Describes proposed changes to a network function group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttachmentPolicyRuleNumber" : Integer,
  "NetworkFunctionGroupName" : String,
  "Tags" : [ Tag, ... ]
}

```

### YAML

```yaml

  AttachmentPolicyRuleNumber: Integer
  NetworkFunctionGroupName: String
  Tags:
    - Tag

```

## Properties

`AttachmentPolicyRuleNumber`

The proposed new attachment policy rule number for the network function group.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkFunctionGroupName`

The proposed name change for the network function group name.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The list of proposed changes to the key-value tags associated with the network function group.

_Required_: No

_Type_: Array of [Tag](aws-properties-networkmanager-connectattachment-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectAttachmentOptions

ProposedSegmentChange

All content copied from https://docs.aws.amazon.com/.
