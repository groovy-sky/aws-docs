---
title: "AWS::NeptuneGraph::GraphSnapshot"
---

This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NeptuneGraph::GraphSnapshot

The `AWS::NeptuneGraph::GraphSnapshot` resource Property description not available. for NeptuneGraph.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NeptuneGraph::GraphSnapshot",
  "Properties" : {
      "GraphIdentifier" : String,
      "SnapshotName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::NeptuneGraph::GraphSnapshot
Properties:
  GraphIdentifier: String
  SnapshotName: String
  Tags:
    - Tag

```

## Properties

`GraphIdentifier`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^g-[a-z0-9]{10}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotName`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!gs-)[a-z][a-z0-9]*(-[a-z0-9]+)*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-neptunegraph-graphsnapshot-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The ARN of the graph snapshot.

`Id`

The unique identifier of the graph snapshot.

`KmsKeyIdentifier`

The ID of the KMS key used to encrypt and decrypt the snapshot.

`SnapshotCreateTime`

The time when the snapshot was created.

`Status`

The status of the graph snapshot.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VectorSearchConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
