This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service ServiceConnectAccessLogConfiguration

Configuration for Service Connect access logging. Access logs provide detailed
information about requests made to your service, including request patterns, response
codes, and timing data for debugging and monitoring purposes.

###### Note

To enable access logs, you must also specify a `logConfiguration` in the `serviceConnectConfiguration`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Format" : String,
  "IncludeQueryParameters" : String
}

```

### YAML

```yaml

  Format: String
  IncludeQueryParameters: String

```

## Properties

`Format`

The format for Service Connect access log output. Choose TEXT for human-readable logs
or JSON for structured data that integrates well with log analysis tools.

_Required_: Yes

_Type_: String

_Allowed values_: `TEXT | JSON`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeQueryParameters`

Specifies whether to include query parameters in Service Connect access logs.

When enabled, query parameters from HTTP requests are included in the access logs.
Consider security and privacy implications when enabling this feature, as query
parameters may contain sensitive information such as request IDs and tokens. By default,
this parameter is `DISABLED`.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | ENABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Secret

ServiceConnectClientAlias

All content copied from https://docs.aws.amazon.com/.
