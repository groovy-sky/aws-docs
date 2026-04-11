This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkInsightsAccessScope AccessScopePathRequest

Describes a path.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : PathStatementRequest,
  "Source" : PathStatementRequest,
  "ThroughResources" : [ ThroughResourcesStatementRequest, ... ]
}

```

### YAML

```yaml

  Destination:
    PathStatementRequest
  Source:
    PathStatementRequest
  ThroughResources:
    - ThroughResourcesStatementRequest

```

## Properties

`Destination`

The destination.

_Required_: No

_Type_: [PathStatementRequest](aws-properties-ec2-networkinsightsaccessscope-pathstatementrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Source`

The source.

_Required_: No

_Type_: [PathStatementRequest](aws-properties-ec2-networkinsightsaccessscope-pathstatementrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ThroughResources`

The through resources.

_Required_: No

_Type_: Array of [ThroughResourcesStatementRequest](aws-properties-ec2-networkinsightsaccessscope-throughresourcesstatementrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::NetworkInsightsAccessScope

PacketHeaderStatementRequest

All content copied from https://docs.aws.amazon.com/.
