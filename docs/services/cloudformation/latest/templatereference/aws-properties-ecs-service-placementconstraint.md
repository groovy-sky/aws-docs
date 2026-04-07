This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service PlacementConstraint

An object representing a constraint on task placement. For more information, see
[Task placement constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html) in the _Amazon_
_Elastic Container Service Developer Guide_.

###### Note

If you're using the Fargate launch type, task placement constraints aren't
supported.

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

A cluster query language expression to apply to the constraint. The expression can
have a maximum length of 2000 characters. You can't specify an expression if the
constraint type is `distinctInstance`. For more information, see [Cluster query language](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html) in the _Amazon Elastic_
_Container Service Developer Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of constraint. Use `distinctInstance` to ensure that each task in
a particular group is running on a different container instance. Use
`memberOf` to restrict the selection to a group of valid
candidates.

_Required_: Yes

_Type_: String

_Allowed values_: `distinctInstance | memberOf`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Associate an Application Load Balancer with a service](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#aws-resource-ecs-service--examples--Associate_an_Application_Load_Balancer_with_a_service)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NetworkConfiguration

PlacementStrategy
