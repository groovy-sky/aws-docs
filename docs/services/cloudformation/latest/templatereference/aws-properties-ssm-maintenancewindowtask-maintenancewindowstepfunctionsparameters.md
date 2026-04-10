This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::MaintenanceWindowTask MaintenanceWindowStepFunctionsParameters

The `MaintenanceWindowStepFunctionsParameters` property type specifies the
parameters for the execution of a `STEP_FUNCTIONS` task in a Systems Manager maintenance window.

`MaintenanceWindowStepFunctionsParameters` is a property of the [TaskInvocationParameters](../userguide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Input" : String,
  "Name" : String
}

```

### YAML

```yaml

  Input: String
  Name: String

```

## Properties

`Input`

The inputs for the `STEP_FUNCTIONS` task.

_Required_: No

_Type_: String

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the `STEP_FUNCTIONS` task.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `80`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MaintenanceWindowRunCommandParameters

NotificationConfig

All content copied from https://docs.aws.amazon.com/.
