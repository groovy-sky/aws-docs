This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskSet Scale

A floating-point percentage of the desired number of tasks to place and keep running
in the task set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Unit" : String,
  "Value" : Number
}

```

### YAML

```yaml

  Unit: String
  Value: Number

```

## Properties

`Unit`

The unit of measure for the scale value.

_Required_: No

_Type_: String

_Allowed values_: `PERCENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value, specified as a percent total of a service's `desiredCount`, to
scale the task set. Accepted values are numbers between 0 and 100.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NetworkConfiguration

ServiceRegistry
