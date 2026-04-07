This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::GraphQLApi LogConfig

The `LogConfig` property type specifies the logging configuration when
writing GraphQL operations and tracing to Amazon CloudWatch for an AWS AppSync GraphQL API.

`LogConfig` is a property of the [AWS::AppSync::GraphQLApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogsRoleArn" : String,
  "ExcludeVerboseContent" : Boolean,
  "FieldLogLevel" : String
}

```

### YAML

```yaml

  CloudWatchLogsRoleArn: String
  ExcludeVerboseContent: Boolean
  FieldLogLevel: String

```

## Properties

`CloudWatchLogsRoleArn`

The service role that AWS AppSync will assume to publish to Amazon CloudWatch Logs in your account.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludeVerboseContent`

Set to TRUE to exclude sections that contain information such as headers, context, and evaluated mapping
templates, regardless of logging level.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldLogLevel`

The field logging level. Values can be NONE, ERROR, INFO, DEBUG, or ALL.

- **NONE**: No field-level logs are
captured.

- **ERROR**: Logs the following information
**only** for the fields that are in the error
category:

- The error section in the server response.

- Field-level errors.

- The generated request/response functions that got resolved for
error fields.

- **INFO**: Logs the following information
**only** for the fields that are in the info
and error categories:

- Info-level messages.

- The user messages sent through `$util.log.info` and
`console.log`.

- Field-level tracing and mapping logs are not shown.

- **DEBUG**: Logs the following information
**only** for the fields that are in the debug,
info, and error categories:

- Debug-level messages.

- The user messages sent through `$util.log.info`,
`$util.log.debug`, `console.log`, and
`console.debug`.

- Field-level tracing and mapping logs are not shown.

- **ALL**: The following information is logged
for all fields in the query:

- Field-level tracing information.

- The generated request/response functions that were resolved for
each field.

_Required_: Yes

_Type_: String

_Allowed values_: `NONE | ERROR | ALL | INFO | DEBUG`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LambdaAuthorizerConfig

OpenIDConnectConfig
