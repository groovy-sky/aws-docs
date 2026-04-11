This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::MaintenanceWindowTask MaintenanceWindowAutomationParameters

The `MaintenanceWindowAutomationParameters` property type specifies the
parameters for an `AUTOMATION` task type for a maintenance window task in
AWS Systems Manager.

`MaintenanceWindowAutomationParameters` is a property of the [TaskInvocationParameters](../userguide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.md) property type.

For information about available parameters in Automation runbooks, you can view the
content of the runbook itself in the Systems Manager console. For information, see
[View runbook content](../../../systems-manager/latest/userguide/automation-documents-reference-details.md#view-automation-json) in the _AWS Systems Manager User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DocumentVersion" : String,
  "Parameters" : Json
}

```

### YAML

```yaml

  DocumentVersion: String
  Parameters: Json

```

## Properties

`DocumentVersion`

The version of an Automation runbook to use during task execution.

_Required_: No

_Type_: String

_Pattern_: `([$]LATEST|[$]DEFAULT|^[1-9][0-9]*$)`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

The parameters for the `AUTOMATION` type task.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoggingInfo

MaintenanceWindowLambdaParameters

All content copied from https://docs.aws.amazon.com/.
