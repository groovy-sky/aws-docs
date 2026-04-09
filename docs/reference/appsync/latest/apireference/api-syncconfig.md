# SyncConfig

Describes a Sync configuration for a resolver.

Specifies which Conflict Detection strategy and Resolution strategy to use when the
resolver is invoked.

## Contents

**conflictDetection**

The Conflict Detection strategy to use.

- **VERSION**: Detect conflicts based on object
versions for this resolver.

- **NONE**: Do not detect conflicts when invoking
this resolver.

Type: String

Valid Values: `VERSION | NONE`

Required: No

**conflictHandler**

The Conflict Resolution strategy to perform in the event of a conflict.

- **OPTIMISTIC\_CONCURRENCY**: Resolve conflicts by
rejecting mutations when versions don't match the latest version at the
server.

- **AUTOMERGE**: Resolve conflicts with the
Automerge conflict resolution strategy.

- **LAMBDA**: Resolve conflicts with an AWS Lambda function supplied in the
`LambdaConflictHandlerConfig`.

Type: String

Valid Values: `OPTIMISTIC_CONCURRENCY | LAMBDA | AUTOMERGE | NONE`

Required: No

**lambdaConflictHandlerConfig**

The `LambdaConflictHandlerConfig` when configuring `LAMBDA` as the
Conflict Handler.

Type: [LambdaConflictHandlerConfig](api-lambdaconflicthandlerconfig.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appsync-2017-07-25/syncconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appsync-2017-07-25/syncconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appsync-2017-07-25/syncconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceApiAssociationSummary

Type

All content copied from https://docs.aws.amazon.com/.
