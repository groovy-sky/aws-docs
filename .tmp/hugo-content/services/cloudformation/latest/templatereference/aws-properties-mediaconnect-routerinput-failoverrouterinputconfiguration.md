This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterInput FailoverRouterInputConfiguration

Configuration settings for a failover router input that allows switching between two input sources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NetworkInterfaceArn" : String,
  "PrimarySourceIndex" : Integer,
  "ProtocolConfigurations" : [ FailoverRouterInputProtocolConfiguration, ... ],
  "SourcePriorityMode" : String
}

```

### YAML

```yaml

  NetworkInterfaceArn: String
  PrimarySourceIndex: Integer
  ProtocolConfigurations:
    - FailoverRouterInputProtocolConfiguration
  SourcePriorityMode: String

```

## Properties

`NetworkInterfaceArn`

The ARN of the network interface to use for this failover router input.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):mediaconnect:[a-z0-9-]+:[0-9]{12}:routerNetworkInterface:[a-z0-9]{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimarySourceIndex`

The index (0 or 1) that specifies which source in the protocol configurations list is currently active. Used to control which of the two failover sources is currently selected. This field is ignored when sourcePriorityMode is set to NO\_PRIORITY

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtocolConfigurations`

A list of exactly two protocol configurations for the failover input sources. Both must use the same protocol type.

_Required_: Yes

_Type_: Array of [FailoverRouterInputProtocolConfiguration](aws-properties-mediaconnect-routerinput-failoverrouterinputprotocolconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourcePriorityMode`

The mode for determining source priority in failover configurations.

_Required_: Yes

_Type_: String

_Allowed values_: `NO_PRIORITY | PRIMARY_SECONDARY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaConnect::RouterInput

FailoverRouterInputProtocolConfiguration

All content copied from https://docs.aws.amazon.com/.
