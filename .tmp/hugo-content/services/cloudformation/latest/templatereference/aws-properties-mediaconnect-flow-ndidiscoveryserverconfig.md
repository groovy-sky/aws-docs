This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow NdiDiscoveryServerConfig

Specifies the configuration settings for individual NDI® discovery servers. A maximum of 3 servers is allowed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DiscoveryServerAddress" : String,
  "DiscoveryServerPort" : Integer,
  "VpcInterfaceAdapter" : String
}

```

### YAML

```yaml

  DiscoveryServerAddress: String
  DiscoveryServerPort: Integer
  VpcInterfaceAdapter: String

```

## Properties

`DiscoveryServerAddress`

The unique network address of the NDI discovery server.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DiscoveryServerPort`

The port for the NDI discovery server. Defaults to 5959 if a custom port isn't specified.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcInterfaceAdapter`

The identifier for the Virtual Private Cloud (VPC) network interface used by the flow.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NdiConfig

NdiSourceSettings

All content copied from https://docs.aws.amazon.com/.
