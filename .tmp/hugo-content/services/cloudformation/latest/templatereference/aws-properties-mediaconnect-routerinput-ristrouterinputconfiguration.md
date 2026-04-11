This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterInput RistRouterInputConfiguration

The configuration settings for a router input using the RIST (Reliable Internet Stream Transport) protocol, including the port and recovery latency.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Port" : Integer,
  "RecoveryLatencyMilliseconds" : Integer
}

```

### YAML

```yaml

  Port: Integer
  RecoveryLatencyMilliseconds: Integer

```

## Properties

`Port`

The port number used for the RIST protocol in the router input configuration.

_Required_: Yes

_Type_: Integer

_Minimum_: `3000`

_Maximum_: `30000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecoveryLatencyMilliseconds`

The recovery latency in milliseconds for the RIST protocol in the router input configuration.

_Required_: Yes

_Type_: Integer

_Minimum_: `10`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PreferredDayTimeMaintenanceConfiguration

RouterInputConfiguration

All content copied from https://docs.aws.amazon.com/.
