This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::StepFunctions::StateMachineAlias

Represents a state machine [alias](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html). An alias routes traffic to one or two versions of the same state machine.

You can create up to 100 aliases for each state machine.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::StepFunctions::StateMachineAlias",
  "Properties" : {
      "DeploymentPreference" : DeploymentPreference,
      "Description" : String,
      "Name" : String,
      "RoutingConfiguration" : [ RoutingConfigurationVersion, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::StepFunctions::StateMachineAlias
Properties:
  DeploymentPreference:
    DeploymentPreference
  Description: String
  Name: String
  RoutingConfiguration:
    - RoutingConfigurationVersion

```

## Properties

`DeploymentPreference`

The settings that enable gradual state machine deployments. These settings include [Alarms](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html#cfn-stepfunctions-statemachinealias-deploymentpreference-alarms), [Interval](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html#cfn-stepfunctions-statemachinealias-deploymentpreference-interval), [Percentage](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html#cfn-stepfunctions-statemachinealias-deploymentpreference-percentage), [StateMachineVersionArn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html#cfn-stepfunctions-statemachinealias-deploymentpreference-statemachineversionarn), and [Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html#cfn-stepfunctions-statemachinealias-deploymentpreference-type).

CloudFormation
automatically shifts traffic from the version an alias currently points to, to a new state machine version that you specify.

###### Note

`RoutingConfiguration` and `DeploymentPreference` are mutually exclusive properties. You must define only one of these properties.

Based on the type of deployment you want to perform, you can specify one of the following settings:

- `LINEAR` \- Shifts traffic to the new version in equal increments with an equal number of minutes between each increment.

For example, if you specify the increment percent as `20` with an interval of `600` minutes, this deployment increases traffic by 20 percent every 600 minutes until the new version receives 100 percent of the traffic.

This deployment immediately rolls back the new version if any
Amazon CloudWatch
alarms are triggered.

- `ALL_AT_ONCE` \- Shifts 100 percent of traffic to the new version immediately.
CloudFormation
monitors the new version and rolls it back automatically to the previous version if any
CloudWatch
alarms are triggered.

- `CANARY` \- Shifts traffic in two increments.

In the first increment, a small percentage of traffic, for example, 10 percent is shifted to the new version. In the second increment, before a specified time interval in seconds gets over, the remaining traffic is shifted to the new version. The shift to the new version for the remaining traffic takes place only if no
CloudWatch
alarms are triggered during the specified time interval.

_Required_: No

_Type_: [DeploymentPreference](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-stepfunctions-statemachinealias-deploymentpreference.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

An optional description of the state machine alias.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the state machine alias. If you don't provide a name, it uses an automatically generated name based on the logical ID.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `80`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoutingConfiguration`

The routing configuration of an alias. Routing configuration splits [StartExecution](https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartExecution.html) requests between one or two versions of the same state machine.

Use `RoutingConfiguration` if you want to explicitly set the alias [weights](https://docs.aws.amazon.com/step-functions/latest/apireference/API_RoutingConfigurationListItem.html#StepFunctions-Type-RoutingConfigurationListItem-weight). Weight is the percentage of traffic you want to route to a state machine version.

###### Note

`RoutingConfiguration` and `DeploymentPreference` are mutually exclusive properties. You must define only one of these properties.

_Required_: No

_Type_: Array of [RoutingConfigurationVersion](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-stepfunctions-statemachinealias-routingconfigurationversion.html)

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you provide the logical ID of this resource to the `Ref` intrinsic
function, `Ref` returns the ARN of the created state machine alias. For example,

```json

{ "Ref": "PROD" }
```

Returns the ARN of the created state machine alias as shown in the following example.

`arn:aws:states:us-east-1:123456789012:stateMachine:myStateMachine:PROD`

For more information about using `Ref`, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The
following is the available attribute.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

`Arn`

Returns the ARN of the state machine alias. For example, `arn:aws:states:us-east-1:123456789012:stateMachine:myStateMachine:PROD`.

## Examples

The following are examples of `LINEAR` deployment that show how you can use an alias to shift traffic to a new version.

- [Shifting traffic to a new version with an alias](#aws-resource-stepfunctions-statemachinealias--examples--Shifting_traffic_to_a_new_version_with_an_alias)

- [Complete example to publish and deploy a new version with an alias](#aws-resource-stepfunctions-statemachinealias--examples--Complete_example_to_publish_and_deploy_a_new_version_with_an_alias)

- [Publish and deploy a version that always points to the most recent state machine revision](#aws-resource-stepfunctions-statemachinealias--examples--Publish_and_deploy_a_version_that_always_points_to_the_most_recent_state_machine_revision)

### Shifting traffic to a new version with an alias

The following example shows an alias named `PROD` that shifts 20 percent of traffic to the version B every five minutes.

#### YAML

```yaml

PROD:
  Type: AWS::StepFunctions::Alias
  Properties:
    Name: PROD
    Description: The PROD state machine alias taking production traffic.
    DeploymentPreference:
      StateMachineVersionArn: !Ref MyStateMachineVersionB
      Type: LINEAR
      Interval: 5
      Percentage: 20
      Alarms:
        # A list of alarms that you want to monitor
        - !Ref AliasErrorMetricGreaterThanZeroAlarm
        - !Ref LatestVersionErrorMetricGreaterThanZeroAlarm
```

### Complete example to publish and deploy a new version with an alias

The following example publishes multiple versions of the same state machine with the [`AWS::StepFunctions::StateMachineVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachineversion.html) resource. The example also creates an alias with the `AWS::StepFunctions::StateMachineAlias` resource and uses that alias to deploy a new state machine version.

#### YAML

```yaml

MyStateMachine:
  Type: AWS::StepFunctions::StateMachine
  Properties:
    StateMachineType: STANDARD
    StateMachineName: MyStateMachine
    RoleArn: arn:aws:iam::12345678912:role/myIamRole
    Definition:
      StartAt: PassState
      States:
        PassState:
          Type: Pass
          Result: Result
          End: true

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

PROD:
  Type: AWS::StepFunctions::StateMachineAlias
  Properties:
    Name: PROD
    Description: The PROD state machine alias taking production traffic.
    DeploymentPreference:
      StateMachineVersionArn: !Ref MyStateMachineVersionB
      Type: LINEAR
      Percentage: 20
      Interval: 5
      Alarms:
        - !Ref CloudWatchAlarm1
        - !Ref CloudWatchAlarm2
```

### Publish and deploy a version that always points to the most recent state machine revision

The following example demonstrates the use of [`StateMachineRevisionId`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachineversion.html#cfn-stepfunctions-statemachineversion-statemachinerevisionid) property to return the value of revision ID for the state machine resource. CloudFormation automatically detects if this property's value is different from the value in previous stack and publishes a new version that always points to the most recent revision of your state machine. The example then creates an alias named PROD to deploy this new version.

#### YAML

```yaml

MyStateMachine:
  Type: AWS::StepFunctions::StateMachine
  Properties:
    StateMachineType: STANDARD
    StateMachineName: MyStateMachine
    RoleArn: arn:aws:iam::12345678912:role/myIamRole
    Definition:
      StartAt: PassState
      States:
        PassState:
          Type: Pass
          Result: Result
          End: true

MyStateMachineVersion:
  Type: AWS::StepFunctions::StateMachineVersion
  Properties:
    Description: Version referencing
    StateMachineArn: !Ref MyStateMachine
    StateMachineRevisionId: !GetAtt MyStateMachine.StateMachineRevisionId

PROD:
  Type: AWS::StepFunctions::StateMachineAlias
  Properties:
    Name: PROD
    Description: The PROD state machine alias taking production traffic.
    DeploymentPreference:
      StateMachineVersionArn: !Ref MyStateMachineVersion
      Type: LINEAR
      Percentage: 20
      Interval: 5
      Alarms:
        - !Ref CloudWatchAlarm1
        - !Ref CloudWatchAlarm2
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TracingConfiguration

DeploymentPreference
