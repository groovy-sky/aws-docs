This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ResourceDefinitionVersion LocalDeviceResourceData

Settings for a local device resource, which represents a file under `/dev`. For
more information, see [Access Local\
Resources with Lambda Functions](../../../greengrass/v1/developerguide/access-local-resources.md) in the _AWS IoT Greengrass Version 1 Developer Guide_.

In an CloudFormation template, `LocalDeviceResourceData` can be used in
the [`ResourceDataContainer`](../userguide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GroupOwnerSetting" : GroupOwnerSetting,
  "SourcePath" : String
}

```

### YAML

```yaml

  GroupOwnerSetting:
    GroupOwnerSetting
  SourcePath: String

```

## Properties

`GroupOwnerSetting`

Settings that define additional Linux OS group permissions to give to the Lambda function process.

_Required_: No

_Type_: [GroupOwnerSetting](aws-properties-greengrass-resourcedefinitionversion-groupownersetting.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourcePath`

The local absolute path of the device resource. The source path for a device resource
can refer only to a character device or block device under `/dev`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [LocalDeviceResourceData](../../../../reference/greengrass/v1/apireference/definitions-localdeviceresourcedata.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GroupOwnerSetting

LocalVolumeResourceData

All content copied from https://docs.aws.amazon.com/.
