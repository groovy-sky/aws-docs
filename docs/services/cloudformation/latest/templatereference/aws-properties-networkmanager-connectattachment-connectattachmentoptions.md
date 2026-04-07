This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::ConnectAttachment ConnectAttachmentOptions

Describes a core network Connect attachment options.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Protocol" : String
}

```

### YAML

```yaml

  Protocol: String

```

## Properties

`Protocol`

The protocol used for the attachment connection.

_Required_: No

_Type_: String

_Allowed values_: `GRE | NO_ENCAP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::NetworkManager::ConnectAttachment

ProposedNetworkFunctionGroupChange
