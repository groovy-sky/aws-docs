This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ResourceDefinitionVersion GroupOwnerSetting

Settings that define additional Linux OS group permissions to give to the Lambda function process. You can give the permissions of the Linux group that owns
the resource or choose another Linux group. These permissions are in addition to the
function's `RunAs` permissions.

In an CloudFormation template, `GroupOwnerSetting` is a property of the
[`LocalDeviceResourceData`](../userguide/aws-properties-greengrass-resourcedefinitionversion-localdeviceresourcedata.md) and [`LocalVolumeResourceData`](../userguide/aws-properties-greengrass-resourcedefinitionversion-localvolumeresourcedata.md) property types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoAddGroupOwner" : Boolean,
  "GroupOwner" : String
}

```

### YAML

```yaml

  AutoAddGroupOwner: Boolean
  GroupOwner: String

```

## Properties

`AutoAddGroupOwner`

Indicates whether to give the privileges of the Linux group that owns the resource to
the Lambda process. This gives the Lambda process the file
access permissions of the Linux group.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroupOwner`

The name of the Linux group whose privileges you want to add to the Lambda
process. This value is ignored if `AutoAddGroupOwner` is true.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [GroupOwnerSetting](../../../../reference/greengrass/v1/apireference/definitions-groupownersetting.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Greengrass::ResourceDefinitionVersion

LocalDeviceResourceData

All content copied from https://docs.aws.amazon.com/.
