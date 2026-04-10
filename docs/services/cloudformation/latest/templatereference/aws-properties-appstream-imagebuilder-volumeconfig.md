This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::ImageBuilder VolumeConfig

Configuration for the root volume of fleet instances and image builders. This allows you to customize the storage capacity beyond the default 200 GB.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VolumeSizeInGb" : Integer
}

```

### YAML

```yaml

  VolumeSizeInGb: Integer

```

## Properties

`VolumeSizeInGb`

The size of the root volume in GB. Valid range is 200-500 GB. The default is 200 GB, which is included in the hourly instance rate. Additional storage beyond 200 GB incurs extra charges and applies to instances regardless of their running state.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

VpcConfig

All content copied from https://docs.aws.amazon.com/.
