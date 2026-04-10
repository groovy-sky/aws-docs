This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::FunctionConfiguration LambdaConflictHandlerConfig

The `LambdaConflictHandlerConfig` object when configuring `LAMBDA` as the Conflict
Handler.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LambdaConflictHandlerArn" : String
}

```

### YAML

```yaml

  LambdaConflictHandlerArn: String

```

## Properties

`LambdaConflictHandlerArn`

The Amazon Resource Name (ARN) for the Lambda function to use as the Conflict Handler.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AppSyncRuntime

SyncConfig

All content copied from https://docs.aws.amazon.com/.
