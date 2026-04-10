This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Container PortInfo

`PortInfo` is a property of the [Container](../userguide/aws-properties-lightsail-container-container.md) property. It describes the ports to open and the protocols to use for
a container on a Amazon Lightsail container service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Port" : String,
  "Protocol" : String
}

```

### YAML

```yaml

  Port: String
  Protocol: String

```

## Properties

`Port`

The open firewall ports of the container.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocol name for the open ports.

_Allowed values_: `HTTP` \| `HTTPS` \| `TCP` \| `UDP`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HealthCheckConfig

PrivateRegistryAccess

All content copied from https://docs.aws.amazon.com/.
