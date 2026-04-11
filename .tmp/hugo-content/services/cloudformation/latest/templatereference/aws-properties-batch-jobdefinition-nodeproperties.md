This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition NodeProperties

An object that represents the node properties of a multi-node parallel job.

###### Note

Node properties can't be specified for Amazon EKS based job definitions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MainNode" : Integer,
  "NodeRangeProperties" : [ NodeRangeProperty, ... ],
  "NumNodes" : Integer
}

```

### YAML

```yaml

  MainNode: Integer
  NodeRangeProperties:
    - NodeRangeProperty
  NumNodes: Integer

```

## Properties

`MainNode`

Specifies the node index for the main node of a multi-node parallel job. This node index
value must be fewer than the number of nodes.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeRangeProperties`

A list of node ranges and their properties that are associated with a multi-node parallel
job.

_Required_: Yes

_Type_: Array of [NodeRangeProperty](aws-properties-batch-jobdefinition-noderangeproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumNodes`

The number of nodes that are associated with a multi-node parallel job.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Multi-node Parallel\
Jobs](../../../batch/latest/userguide/multi-node-parallel-jobs.md) in the _AWS Batch User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkConfiguration

NodeRangeProperty

All content copied from https://docs.aws.amazon.com/.
