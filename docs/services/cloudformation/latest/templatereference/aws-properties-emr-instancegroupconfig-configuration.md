This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceGroupConfig Configuration

`Configurations` is a property of the `AWS::EMR::Cluster` resource that specifies the configuration of applications on an Amazon EMR cluster.

Configurations are optional. You can use them to have EMR customize applications and software bundled with Amazon EMR when a cluster is created. A configuration consists of a classification, properties, and optional nested configurations. A classification refers to an application-specific configuration file. Properties are the settings you want to change in that file. For more information, see [Configuring Applications](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-configure-apps.html).

###### Note

Applies only to Amazon EMR releases 4.0 and later.

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

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConfigurationProperties`

Within a configuration classification, a set of properties that represent the settings that you want to change in the configuration file. Duplicates not allowed.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Configurations`

A list of additional configurations to apply within a configuration object.

_Required_: No

_Type_: Array of [Configuration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emr-instancegroupconfig-configuration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CloudWatchAlarmDefinition

EbsBlockDeviceConfig
