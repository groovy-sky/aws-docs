This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule PlacementStrategy

The task placement strategy for a task or service. To learn more, see [Task Placement Strategies](../../../amazonecs/latest/developerguide/task-placement-strategies.md) in the Amazon Elastic Container Service Service Developer
Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Field" : String,
  "Type" : String
}

```

### YAML

```yaml

  Field: String
  Type: String

```

## Properties

`Field`

The field to apply the placement strategy against. For the spread placement strategy,
valid values are instanceId (or host, which has the same effect), or any platform or custom
attribute that is applied to a container instance, such as attribute:ecs.availability-zone.
For the binpack placement strategy, valid values are cpu and memory. For the random placement
strategy, this field is not used.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of placement strategy. The random placement strategy randomly places tasks on
available candidates. The spread placement strategy spreads placement across available
candidates evenly based on the field parameter. The binpack strategy places tasks on available
candidates that have the least available amount of the resource that is specified with the
field parameter. For example, if you binpack on memory, a task is placed on the instance with
the least amount of remaining memory (but still enough to run the task).

_Required_: No

_Type_: String

_Allowed values_: `random | spread | binpack`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PlacementConstraint

RedshiftDataParameters

All content copied from https://docs.aws.amazon.com/.
