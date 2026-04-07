This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::StepFunctions::StateMachineVersion

Represents a state machine [version](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html). A published version uses the latest state machine [_revision_](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html). A revision is an immutable, read-only snapshot of a state machine’s definition and configuration.

You can publish up to 1000 versions for each state machine.

###### Important

Before you delete a version, make sure that version's ARN isn't being referenced in any long-running workflows or application code outside of the stack.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::StepFunctions::StateMachineVersion",
  "Properties" : {
      "Description" : String,
      "StateMachineArn" : String,
      "StateMachineRevisionId" : String
    }
}

```

### YAML

```yaml

Type: AWS::StepFunctions::StateMachineVersion
Properties:
  Description: String
  StateMachineArn: String
  StateMachineRevisionId: String

```

## Properties

`Description`

An optional description of the state machine version.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StateMachineArn`

The Amazon Resource Name (ARN) of the state machine.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StateMachineRevisionId`

Identifier for a state machine revision, which is an immutable, read-only snapshot of a state machine’s definition and configuration.

Only publish the state machine version if the current state machine's revision ID matches the specified ID. Use this option to avoid publishing a version if the state machine has changed since you last updated it.

To specify the initial state machine revision, set the value as `INITIAL`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you provide the logical ID of this resource to the `Ref` intrinsic
function, `Ref` returns the ARN of the published state machine version. For example, `arn:aws:states:us-east-1:123456789012:stateMachine:myStateMachine:1.`

For more information about using `Ref`, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The following is the available attribute.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

`Arn`

Returns the ARN of the state machine version. For example, `arn:aws:states:us-east-1:123456789012:stateMachine:myStateMachine:1`.

## Examples

The following CloudFormation template examples show how you can create multiple versions of the same state machine and publish a version using the latest revision of a state machine.

- [Publish multiple state machine versions](#aws-resource-stepfunctions-statemachineversion--examples--Publish_multiple_state_machine_versions)

- [Publish a version for the latest revision of a state machine](#aws-resource-stepfunctions-statemachineversion--examples--Publish_a_version_for_the_latest_revision_of_a_state_machine)

### Publish multiple state machine versions

#### YAML

```yaml

MyStateMachineVersionA:
  Type: AWS::StepFunctions::StateMachineVersion
  Properties:
    Description: Version 1
    StateMachineArn: !Ref MyStateMachine

MyStateMachineVersionB:
  Type: AWS::StepFunctions::StateMachineVersion
  Properties:
    Description: Version 2
    StateMachineArn: !Ref MyStateMachine
```

### Publish a version for the latest revision of a state machine

The following example uses the [`StateMachineRevisionId`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachineversion.html#cfn-stepfunctions-statemachineversion-statemachinerevisionid) property, which returns a unique value for revision ID of the state machine resource whenever an update is made to the state machine. CloudFormation automatically detects a different value for the `StateMachineRevisionId` property compared to the property's value in previous stack and makes a replacement update to the `AWS::StepFunctions::StateMachineVersion` resource. This publishes a new version that points to the most recent revision of your state machine.

#### YAML

```yaml

MyLatestStateMachineVersion:
  Type: AWS::StepFunctions::StateMachineVersion
  Properties:
    Description: This version points to the most recent revision of a state machine.
    StateMachineArn: !Ref MyStateMachine
    StateMachineRevisionId: !GetAtt MyStateMachine.StateMachineRevisionId
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RoutingConfigurationVersion

Next
