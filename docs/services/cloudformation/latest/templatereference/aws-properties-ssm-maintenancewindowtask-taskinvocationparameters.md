This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::MaintenanceWindowTask TaskInvocationParameters

The `TaskInvocationParameters` property type specifies the task execution
parameters for a maintenance window task in AWS Systems Manager.

`TaskInvocationParameters` is a property of the [AWS::SSM::MaintenanceWindowTask](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html) property type.

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

_Type_: [MaintenanceWindowAutomationParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssm-maintenancewindowtask-maintenancewindowautomationparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaintenanceWindowLambdaParameters`

The parameters for a `LAMBDA` task type.

_Required_: No

_Type_: [MaintenanceWindowLambdaParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssm-maintenancewindowtask-maintenancewindowlambdaparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaintenanceWindowRunCommandParameters`

The parameters for a `RUN_COMMAND` task type.

_Required_: No

_Type_: [MaintenanceWindowRunCommandParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaintenanceWindowStepFunctionsParameters`

The parameters for a `STEP_FUNCTIONS` task type.

_Required_: No

_Type_: [MaintenanceWindowStepFunctionsParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssm-maintenancewindowtask-maintenancewindowstepfunctionsparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Target

AWS::SSM::Parameter
