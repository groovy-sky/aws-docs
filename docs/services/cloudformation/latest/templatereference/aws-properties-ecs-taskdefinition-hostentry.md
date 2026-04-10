This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskDefinition HostEntry

The `HostEntry` property specifies a hostname and an IP address that are
added to the `/etc/hosts` file of a container through the
`extraHosts` parameter of its `ContainerDefinition`
resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Hostname" : String,
  "IpAddress" : String
}

```

### YAML

```yaml

  Hostname: String
  IpAddress: String

```

## Properties

`Hostname`

The hostname to use in the `/etc/hosts` entry.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IpAddress`

The IP address to use in the `/etc/hosts` entry.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HealthCheck

HostVolumeProperties

All content copied from https://docs.aws.amazon.com/.
