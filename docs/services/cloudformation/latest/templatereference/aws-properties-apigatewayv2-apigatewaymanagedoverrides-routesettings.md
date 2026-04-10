This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::ApiGatewayManagedOverrides RouteSettings

The `RouteSettings` property overrides the route settings for an API Gateway-managed route.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataTraceEnabled" : Boolean,
  "DetailedMetricsEnabled" : Boolean,
  "LoggingLevel" : String,
  "ThrottlingBurstLimit" : Integer,
  "ThrottlingRateLimit" : Number
}

```

### YAML

```yaml

  DataTraceEnabled: Boolean
  DetailedMetricsEnabled: Boolean
  LoggingLevel: String
  ThrottlingBurstLimit: Integer
  ThrottlingRateLimit: Number

```

## Properties

`DataTraceEnabled`

Specifies whether ( `true`) or not ( `false`) data trace logging is enabled for this route. This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DetailedMetricsEnabled`

Specifies whether detailed metrics are enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoggingLevel`

Specifies the logging level for this route: `INFO`, `ERROR`, or `OFF`. This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThrottlingBurstLimit`

Specifies the throttling burst limit.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThrottlingRateLimit`

Specifies the throttling rate limit.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RouteOverrides

StageOverrides

All content copied from https://docs.aws.amazon.com/.
