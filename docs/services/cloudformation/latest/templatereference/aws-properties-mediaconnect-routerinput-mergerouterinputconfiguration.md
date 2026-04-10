This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterInput MergeRouterInputConfiguration

Configuration settings for a merge router input that combines two input sources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MergeRecoveryWindowMilliseconds" : Integer,
  "NetworkInterfaceArn" : String,
  "ProtocolConfigurations" : [ MergeRouterInputProtocolConfiguration, ... ]
}

```

### YAML

```yaml

  MergeRecoveryWindowMilliseconds: Integer
  NetworkInterfaceArn: String
  ProtocolConfigurations:
    - MergeRouterInputProtocolConfiguration

```

## Properties

`MergeRecoveryWindowMilliseconds`

The time window in milliseconds for merging the two input sources.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkInterfaceArn`

The ARN of the network interface to use for this merge router input.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):mediaconnect:[a-z0-9-]+:[0-9]{12}:routerNetworkInterface:[a-z0-9]{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtocolConfigurations`

A list of exactly two protocol configurations for the merge input sources. Both must use the same protocol type.

_Required_: Yes

_Type_: Array of [MergeRouterInputProtocolConfiguration](aws-properties-mediaconnect-routerinput-mergerouterinputprotocolconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MediaConnectFlowRouterInputConfiguration

MergeRouterInputProtocolConfiguration

All content copied from https://docs.aws.amazon.com/.
