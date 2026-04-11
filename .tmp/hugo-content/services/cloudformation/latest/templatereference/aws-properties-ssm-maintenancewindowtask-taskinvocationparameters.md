This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::MaintenanceWindowTask TaskInvocationParameters

The `TaskInvocationParameters` property type specifies the task execution
parameters for a maintenance window task in AWS Systems Manager.

`TaskInvocationParameters` is a property of the [AWS::SSM::MaintenanceWindowTask](../userguide/aws-resource-ssm-maintenancewindowtask.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaintenanceWindowAutomationParameters" : MaintenanceWindowAutomationParameters,
  "MaintenanceWindowLambdaParameters" : MaintenanceWindowLambdaParameters,
  "MaintenanceWindowRunCommandParameters" : MaintenanceWindowRunCommandParameters,
  "MaintenanceWindowStepFunctionsParameters" : MaintenanceWindowStepFunctionsParameters
}

```

### YAML

```yaml

  MaintenanceWindowAutomationParameters:
    MaintenanceWindowAutomationParameters
  MaintenanceWindowLambdaParameters:
    MaintenanceWindowLambdaParameters
  MaintenanceWindowRunCommandParameters:
    MaintenanceWindowRunCommandParameters
  MaintenanceWindowStepFunctionsParameters:
    MaintenanceWindowStepFunctionsParameters

```

## Properties

`MaintenanceWindowAutomationParameters`

The parameters for an `AUTOMATION` task type.

_Required_: No

_Type_: [MaintenanceWindowAutomationParameters](aws-properties-ssm-maintenancewindowtask-maintenancewindowautomationparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaintenanceWindowLambdaParameters`

The parameters for a `LAMBDA` task type.

_Required_: No

_Type_: [MaintenanceWindowLambdaParameters](aws-properties-ssm-maintenancewindowtask-maintenancewindowlambdaparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaintenanceWindowRunCommandParameters`

The parameters for a `RUN_COMMAND` task type.

_Required_: No

_Type_: [MaintenanceWindowRunCommandParameters](aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaintenanceWindowStepFunctionsParameters`

The parameters for a `STEP_FUNCTIONS` task type.

_Required_: No

_Type_: [MaintenanceWindowStepFunctionsParameters](aws-properties-ssm-maintenancewindowtask-maintenancewindowstepfunctionsparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Target

AWS::SSM::Parameter

All content copied from https://docs.aws.amazon.com/.
