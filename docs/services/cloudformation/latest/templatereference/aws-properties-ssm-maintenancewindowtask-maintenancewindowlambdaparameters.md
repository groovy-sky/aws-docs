This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::MaintenanceWindowTask MaintenanceWindowLambdaParameters

The `MaintenanceWindowLambdaParameters` property type specifies the
parameters for a `LAMBDA` task type for a maintenance window task in AWS Systems Manager.

`MaintenanceWindowLambdaParameters` is a property of the [TaskInvocationParameters](../userguide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientContext" : String,
  "Payload" : String,
  "Qualifier" : String
}

```

### YAML

```yaml

  ClientContext: String
  Payload: String
  Qualifier: String

```

## Properties

`ClientContext`

Client-specific information to pass to the AWS Lambda function that
you're invoking. You can then use the `context` variable to process the
client information in your AWS Lambda function.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `8000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Payload`

JSON to provide to your AWS Lambda function as input.

###### Important

Although `Type` is listed as "String" for this property, the payload
content must be formatted as a Base64-encoded binary data object.

_Length Constraint:_ 4096

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Qualifier`

An AWS Lambda function version or alias name. If you specify a function
version, the action uses the qualified function Amazon Resource Name (ARN) to invoke a
specific Lambda function. If you specify an alias name, the action uses
the alias ARN to invoke the Lambda function version that the alias points
to.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MaintenanceWindowAutomationParameters

MaintenanceWindowRunCommandParameters

All content copied from https://docs.aws.amazon.com/.
