This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskDefinition TaskDefinitionPlacementConstraint

The constraint on task placement in the task definition. For more information, see
[Task placement constraints](../../../amazonecs/latest/developerguide/task-placement-constraints.md) in the _Amazon_
_Elastic Container Service Developer Guide_.

###### Note

Task placement constraints aren't supported for tasks run on AWS Fargate.

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

A cluster query language expression to apply to the constraint. For more information,
see [Cluster query language](../../../amazonecs/latest/developerguide/cluster-query-language.md) in the _Amazon Elastic_
_Container Service Developer Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of constraint. The `MemberOf` constraint restricts selection to be
from a group of valid candidates.

_Required_: Yes

_Type_: String

_Allowed values_: `memberOf`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tmpfs

All content copied from https://docs.aws.amazon.com/.
