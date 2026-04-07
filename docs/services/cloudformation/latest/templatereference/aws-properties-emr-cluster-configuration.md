This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster Configuration

###### Note

Used only with Amazon EMR release 4.0 and later.

`Configuration` is a subproperty of `InstanceFleetConfig` or `InstanceGroupConfig`. `Configuration` specifies optional configurations for customizing open-source big data applications and environment parameters. A configuration consists of a classification, properties, and optional nested configurations. A classification refers to an application-specific configuration file. Properties are the settings you want to change in that file. For more information, see [Configuring Applications](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-configure-apps.html) in the _Amazon EMR Release Guide_.

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

A list of additional configurations to apply within a configuration object.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Configurations`

A list of additional configurations to apply within a configuration object.

_Required_: No

_Type_: Array of [Configuration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emr-cluster-configuration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ComputeLimits

EbsBlockDeviceConfig
