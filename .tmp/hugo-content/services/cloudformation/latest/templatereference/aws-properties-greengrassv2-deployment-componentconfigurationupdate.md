This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::Deployment ComponentConfigurationUpdate

Contains information about a deployment's update to a component's configuration on AWS IoT Greengrass core devices. For more information, see [Update component\
configurations](../../../greengrass/v2/developerguide/update-component-configurations.md) in the _AWS IoT Greengrass V2 Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Merge" : String,
  "Reset" : [ String, ... ]
}

```

### YAML

```yaml

  Merge: String
  Reset:
    - String

```

## Properties

`Merge`

A serialized JSON string that contains the configuration object to merge to target
devices. The core device merges this configuration with the component's existing
configuration. If this is the first time a component deploys on a device, the core device
merges this configuration with the component's default configuration. This means that the core
device keeps it's existing configuration for keys and values that you don't specify in this
object. For more information, see [Merge configuration updates](../../../greengrass/v2/developerguide/update-component-configurations.md#merge-configuration-update) in the _AWS IoT Greengrass V2 Developer_
_Guide_.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `10485760`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Reset`

The list of configuration nodes to reset to default values on target devices. Use JSON
pointers to specify each node to reset. JSON pointers start with a forward slash
( `/`) and use forward slashes to separate the key for each level in the object.
For more information, see the [JSON pointer\
specification](https://tools.ietf.org/html/rfc6901) and [Reset configuration updates](../../../greengrass/v2/developerguide/update-component-configurations.md#reset-configuration-update) in the _AWS IoT Greengrass V2 Developer_
_Guide_.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::GreengrassV2::Deployment

ComponentDeploymentSpecification

All content copied from https://docs.aws.amazon.com/.
