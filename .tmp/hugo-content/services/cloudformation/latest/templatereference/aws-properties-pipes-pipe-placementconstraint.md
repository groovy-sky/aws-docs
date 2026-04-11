This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PlacementConstraint

An object representing a constraint on task placement. To learn more, see [Task Placement\
Constraints](../../../amazonecs/latest/developerguide/task-placement-constraints.md) in the Amazon Elastic Container Service Developer Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Expression" : String,
  "Type" : String
}

```

### YAML

```yaml

  Expression: String
  Type: String

```

## Properties

`Expression`

A cluster query language expression to apply to the constraint. You cannot specify an
expression if the constraint type is `distinctInstance`. To learn more, see
[Cluster Query\
Language](../../../amazonecs/latest/developerguide/cluster-query-language.md) in the Amazon Elastic Container Service Developer Guide.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of constraint. Use distinctInstance to ensure that each task in a particular
group is running on a different container instance. Use memberOf to restrict the selection
to a group of valid candidates.

_Required_: No

_Type_: String

_Allowed values_: `distinctInstance | memberOf`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeTargetTimestreamParameters

PlacementStrategy

All content copied from https://docs.aws.amazon.com/.
