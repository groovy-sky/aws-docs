This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Version FunctionScalingConfig

Configuration that defines the scaling behavior for a Lambda Managed Instances function, including the minimum and maximum number of execution environments that can be provisioned.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxExecutionEnvironments" : Integer,
  "MinExecutionEnvironments" : Integer
}

```

### YAML

```yaml

  MaxExecutionEnvironments: Integer
  MinExecutionEnvironments: Integer

```

## Properties

`MaxExecutionEnvironments`

The maximum number of execution environments that can be provisioned for the function.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinExecutionEnvironments`

The minimum number of execution environments to maintain for the function.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Lambda::Version

ProvisionedConcurrencyConfiguration

All content copied from https://docs.aws.amazon.com/.
