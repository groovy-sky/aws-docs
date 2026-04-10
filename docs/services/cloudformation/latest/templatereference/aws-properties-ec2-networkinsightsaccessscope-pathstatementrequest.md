This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkInsightsAccessScope PathStatementRequest

Describes a path statement.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PacketHeaderStatement" : PacketHeaderStatementRequest,
  "ResourceStatement" : ResourceStatementRequest
}

```

### YAML

```yaml

  PacketHeaderStatement:
    PacketHeaderStatementRequest
  ResourceStatement:
    ResourceStatementRequest

```

## Properties

`PacketHeaderStatement`

The packet header statement.

_Required_: No

_Type_: [PacketHeaderStatementRequest](aws-properties-ec2-networkinsightsaccessscope-packetheaderstatementrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceStatement`

The resource statement.

_Required_: No

_Type_: [ResourceStatementRequest](aws-properties-ec2-networkinsightsaccessscope-resourcestatementrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PacketHeaderStatementRequest

ResourceStatementRequest

All content copied from https://docs.aws.amazon.com/.
