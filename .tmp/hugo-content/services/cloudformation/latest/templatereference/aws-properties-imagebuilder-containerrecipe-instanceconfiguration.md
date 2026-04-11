This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ContainerRecipe InstanceConfiguration

Defines a custom base AMI and block device mapping configurations of an instance
used for building and testing container images.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlockDeviceMappings" : [ InstanceBlockDeviceMapping, ... ],
  "Image" : String
}

```

### YAML

```yaml

  BlockDeviceMappings:
    - InstanceBlockDeviceMapping
  Image: String

```

## Properties

`BlockDeviceMappings`

Defines the block devices to attach for building an instance from this Image Builder
AMI.

_Required_: No

_Type_: Array of [InstanceBlockDeviceMapping](aws-properties-imagebuilder-containerrecipe-instanceblockdevicemapping.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Image`

The base image for a container build and test instance. This can contain an AMI ID
or it can specify an AWS Systems Manager (SSM) Parameter Store Parameter, prefixed by `ssm:`,
followed by the parameter name or ARN.

If not specified, Image Builder uses the appropriate ECS-optimized AMI as a base image.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceBlockDeviceMapping

LatestVersion

All content copied from https://docs.aws.amazon.com/.
