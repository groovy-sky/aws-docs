This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceFleetConfig Configuration

###### Note

Used only with Amazon EMR release 4.0 and later.

`Configuration` specifies optional configurations for customizing open-source big data applications and environment parameters. A configuration consists of a classification, properties, and optional nested configurations. A classification refers to an application-specific configuration file. Properties are the settings you want to change in that file. For more information, see [Configuring Applications](../../../emr/latest/releaseguide/emr-configure-apps.md) in the _Amazon EMR Release Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Classification" : String,
  "ConfigurationProperties" : {Key: Value, ...},
  "Configurations" : [ Configuration, ... ]
}

```

### YAML

```yaml

  Classification: String
  ConfigurationProperties:
    Key: Value
  Configurations:
    - Configuration

```

## Properties

`Classification`

The classification within a configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfigurationProperties`

Within a configuration classification, a set of properties that represent the settings that you want to change in the configuration file. Duplicates not allowed.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Configurations`

A list of additional configurations to apply within a configuration object.

_Required_: No

_Type_: Array of [Configuration](aws-properties-emr-instancefleetconfig-configuration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EMR::InstanceFleetConfig

EbsBlockDeviceConfig

All content copied from https://docs.aws.amazon.com/.
