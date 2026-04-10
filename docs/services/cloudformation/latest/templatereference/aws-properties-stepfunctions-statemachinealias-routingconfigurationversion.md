This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::StepFunctions::StateMachineAlias RoutingConfigurationVersion

The state machine version to which you want to route the execution traffic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StateMachineVersionArn" : String,
  "Weight" : Integer
}

```

### YAML

```yaml

  StateMachineVersionArn: String
  Weight: Integer

```

## Properties

`StateMachineVersionArn`

The Amazon Resource Name (ARN) that identifies one or two state machine versions defined in the routing configuration.

If you specify the ARN of a second version, it must belong to the same state machine as the first version.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Weight`

The percentage of traffic you want to route to the state machine version. The sum of the weights in the routing configuration must be equal to 100.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeploymentPreference

AWS::StepFunctions::StateMachineVersion

All content copied from https://docs.aws.amazon.com/.
