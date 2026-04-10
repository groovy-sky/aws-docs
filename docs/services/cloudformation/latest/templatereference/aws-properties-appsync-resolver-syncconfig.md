This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::Resolver SyncConfig

Describes a Sync configuration for a resolver.

Specifies which Conflict Detection strategy and Resolution strategy to use when the resolver is
invoked.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConflictDetection" : String,
  "ConflictHandler" : String,
  "LambdaConflictHandlerConfig" : LambdaConflictHandlerConfig
}

```

### YAML

```yaml

  ConflictDetection: String
  ConflictHandler: String
  LambdaConflictHandlerConfig:
    LambdaConflictHandlerConfig

```

## Properties

`ConflictDetection`

The Conflict Detection strategy to use.

- **VERSION**: Detect conflicts based on object versions for this
resolver.

- **NONE**: Do not detect conflicts when invoking this
resolver.

_Required_: Yes

_Type_: String

_Allowed values_: `VERSION | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConflictHandler`

The Conflict Resolution strategy to perform in the event of a conflict.

- **OPTIMISTIC\_CONCURRENCY**: Resolve conflicts by rejecting mutations
when versions don't match the latest version at the server.

- **AUTOMERGE**: Resolve conflicts with the Automerge conflict
resolution strategy.

- **LAMBDA**: Resolve conflicts with an AWS Lambda function
supplied in the `LambdaConflictHandlerConfig`.

_Required_: No

_Type_: String

_Allowed values_: `OPTIMISTIC_CONCURRENCY | LAMBDA | AUTOMERGE | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaConflictHandlerConfig`

The `LambdaConflictHandlerConfig` when configuring `LAMBDA` as the Conflict
Handler.

_Required_: No

_Type_: [LambdaConflictHandlerConfig](aws-properties-appsync-resolver-lambdaconflicthandlerconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipelineConfig

AWS::AppSync::SourceApiAssociation

All content copied from https://docs.aws.amazon.com/.
