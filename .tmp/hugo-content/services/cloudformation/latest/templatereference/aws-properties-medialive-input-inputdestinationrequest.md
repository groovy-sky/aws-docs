This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Input InputDestinationRequest

Settings that apply only if the input is a push type of input.

The parent of this entity is Input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Network" : String,
  "NetworkRoutes" : [ InputRequestDestinationRoute, ... ],
  "StaticIpAddress" : String,
  "StreamName" : String
}

```

### YAML

```yaml

  Network: String
  NetworkRoutes:
    - InputRequestDestinationRoute
  StaticIpAddress: String
  StreamName: String

```

## Properties

`Network`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkRoutes`

Property description not available.

_Required_: No

_Type_: Array of [InputRequestDestinationRoute](aws-properties-medialive-input-inputrequestdestinationroute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StaticIpAddress`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamName`

The stream name (application name/application instance) for the location the RTMP source content will be pushed
to in MediaLive.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaLive::Input

InputDeviceSettings

All content copied from https://docs.aws.amazon.com/.
