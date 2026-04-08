# PlacementConstraint

An object representing a constraint on task placement. For more information, see
[Task placement constraints](../../../../services/amazonecs/latest/developerguide/task-placement-constraints.md) in the _Amazon_
_Elastic Container Service Developer Guide_.

###### Note

If you're using the Fargate launch type, task placement constraints aren't
supported.

## Contents

**expression**

A cluster query language expression to apply to the constraint. The expression can
have a maximum length of 2000 characters. You can't specify an expression if the
constraint type is `distinctInstance`. For more information, see [Cluster query language](../../../../services/amazonecs/latest/developerguide/cluster-query-language.md) in the _Amazon Elastic_
_Container Service Developer Guide_.

Type: String

Required: No

**type**

The type of constraint. Use `distinctInstance` to ensure that each task in
a particular group is running on a different container instance. Use
`memberOf` to restrict the selection to a group of valid
candidates.

Type: String

Valid Values: `distinctInstance | memberOf`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/placementconstraint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/placementconstraint.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/placementconstraint.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

NetworkInterfaceCountRequest

PlacementStrategy
