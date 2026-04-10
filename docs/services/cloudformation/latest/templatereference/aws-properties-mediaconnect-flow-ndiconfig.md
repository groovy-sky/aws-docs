This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow NdiConfig

Specifies the configuration settings for NDI sources and outputs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MachineName" : String,
  "NdiDiscoveryServers" : [ NdiDiscoveryServerConfig, ... ],
  "NdiState" : String
}

```

### YAML

```yaml

  MachineName: String
  NdiDiscoveryServers:
    - NdiDiscoveryServerConfig
  NdiState: String

```

## Properties

`MachineName`

A prefix for the names of the NDI sources that the flow creates. If a custom name isn't specified, MediaConnect generates a unique 12-character ID as the prefix.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NdiDiscoveryServers`

A list of up to three NDI discovery server configurations. While not required by the API, this configuration is necessary for NDI functionality to work properly.

_Required_: No

_Type_: Array of [NdiDiscoveryServerConfig](aws-properties-mediaconnect-flow-ndidiscoveryserverconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NdiState`

A
setting that controls whether NDI® sources or outputs can be used in the flow.

The default value is `DISABLED`. This value must be set as `ENABLED` for your flow to support NDI sources or outputs.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MediaStreamSourceConfiguration

NdiDiscoveryServerConfig

All content copied from https://docs.aws.amazon.com/.
