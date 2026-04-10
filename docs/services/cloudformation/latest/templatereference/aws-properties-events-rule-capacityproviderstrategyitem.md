This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule CapacityProviderStrategyItem

The details of a capacity provider strategy. To learn more, see [CapacityProviderStrategyItem](../../../../reference/amazonecs/latest/apireference/api-capacityproviderstrategyitem.md) in the Amazon ECS API Reference.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Base" : Integer,
  "CapacityProvider" : String,
  "Weight" : Integer
}

```

### YAML

```yaml

  Base: Integer
  CapacityProvider: String
  Weight: Integer

```

## Properties

`Base`

The base value designates how many tasks, at a minimum, to run on the specified capacity
provider. Only one capacity provider in a capacity provider strategy can have a base defined.
If no value is specified, the default value of 0 is used.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CapacityProvider`

The short name of the capacity provider.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Weight`

The weight value designates the relative percentage of the total number of tasks launched
that should use the specified capacity provider. The weight value is taken into consideration
after the base value, if defined, is satisfied.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Set the CapacityProviderStrategyItem property

The following example sets the `CapacityProviderStrategyItem` property to run a minimum of 10 tasks on CapacityProvider1.

#### JSON

```json

"CapacityProviderStrategy": {
  "Base": 10,
  "CapacityProvider": "CapacityProvider1",
  "Weight": 0
}
```

#### YAML

```yaml

CapacityProviderStrategy
  Base: 10
  CapacityProvider: CapacityProvider1
  Weight: 0
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchRetryStrategy

DeadLetterConfig

All content copied from https://docs.aws.amazon.com/.
