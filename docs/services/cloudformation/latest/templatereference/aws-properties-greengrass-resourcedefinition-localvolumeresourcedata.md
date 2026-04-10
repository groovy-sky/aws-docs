This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ResourceDefinition LocalVolumeResourceData

Settings for a local volume resource, which represents a file or directory on the root
file system. For more information, see [Access Local\
Resources with Lambda Functions](../../../greengrass/v1/developerguide/access-local-resources.md) in the _AWS IoT Greengrass Version 1 Developer Guide_.

In an CloudFormation template, `LocalVolumeResourceData` can be used in
the [`ResourceDataContainer`](../userguide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationPath" : String,
  "GroupOwnerSetting" : GroupOwnerSetting,
  "SourcePath" : String
}

```

### YAML

```yaml

  DestinationPath: String
  GroupOwnerSetting:
    GroupOwnerSetting
  SourcePath: String

```

## Properties

`DestinationPath`

The absolute local path of the resource in the Lambda environment.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroupOwnerSetting`

Settings that define additional Linux OS group permissions to give to the Lambda function process.

_Required_: No

_Type_: [GroupOwnerSetting](aws-properties-greengrass-resourcedefinition-groupownersetting.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourcePath`

The local absolute path of the volume resource on the host. The source path for a volume
resource type cannot start with `/sys`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [LocalVolumeResourceData](../../../../reference/greengrass/v1/apireference/definitions-localvolumeresourcedata.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LocalDeviceResourceData

ResourceDataContainer

All content copied from https://docs.aws.amazon.com/.
