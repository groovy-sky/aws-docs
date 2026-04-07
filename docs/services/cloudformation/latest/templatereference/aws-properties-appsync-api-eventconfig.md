This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::Api EventConfig

Describes the authorization configuration for connections, message publishing, message
subscriptions, and logging for an Event API.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthProviders" : [ AuthProvider, ... ],
  "ConnectionAuthModes" : [ AuthMode, ... ],
  "DefaultPublishAuthModes" : [ AuthMode, ... ],
  "DefaultSubscribeAuthModes" : [ AuthMode, ... ],
  "LogConfig" : EventLogConfig
}

```

### YAML

```yaml

  AuthProviders:
    - AuthProvider
  ConnectionAuthModes:
    - AuthMode
  DefaultPublishAuthModes:
    - AuthMode
  DefaultSubscribeAuthModes:
    - AuthMode
  LogConfig:
    EventLogConfig

```

## Properties

`AuthProviders`

A list of authorization providers.

_Required_: Yes

_Type_: Array of [AuthProvider](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-api-authprovider.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionAuthModes`

A list of valid authorization modes for the Event API connections.

_Required_: Yes

_Type_: Array of [AuthMode](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-api-authmode.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultPublishAuthModes`

A list of valid authorization modes for the Event API publishing.

_Required_: Yes

_Type_: Array of [AuthMode](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-api-authmode.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultSubscribeAuthModes`

A list of valid authorization modes for the Event API subscriptions.

_Required_: Yes

_Type_: Array of [AuthMode](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-api-authmode.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogConfig`

The CloudWatch Logs configuration for the Event API.

_Required_: No

_Type_: [EventLogConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appsync-api-eventlogconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DnsMap

EventLogConfig
