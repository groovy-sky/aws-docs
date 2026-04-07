This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::StackSet StackInstances

Stack instances in some specific accounts and Regions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeploymentTargets" : DeploymentTargets,
  "ParameterOverrides" : [ Parameter, ... ],
  "Regions" : [ String, ... ]
}

```

### YAML

```yaml

  DeploymentTargets:
    DeploymentTargets
  ParameterOverrides:
    - Parameter
  Regions:
    - String

```

## Properties

`DeploymentTargets`

The AWS Organizations accounts or AWS accounts to deploy stacks to
in the specified Regions.

_Required_: Yes

_Type_: [DeploymentTargets](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-stackset-deploymenttargets.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterOverrides`

A list of StackSet parameters whose values you want to override in the selected stack
instances.

_Required_: No

_Type_: Array of [Parameter](aws-properties-cloudformation-stackset-parameter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Regions`

The names of one or more Regions where you want to create stack instances using the
specified AWS accounts.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Parameter

Tag
