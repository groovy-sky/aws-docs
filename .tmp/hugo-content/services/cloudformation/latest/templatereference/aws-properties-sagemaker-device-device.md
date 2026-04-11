This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Device Device

Information of a particular device.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "DeviceName" : String,
  "IotThingName" : String
}

```

### YAML

```yaml

  Description: String
  DeviceName: String
  IotThingName: String

```

## Properties

`Description`

Description of the device.

_Required_: No

_Type_: String

_Pattern_: `[\S\s]+`

_Minimum_: `1`

_Maximum_: `40`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceName`

The name of the device.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IotThingName`

AWS Internet of Things (IoT) object name.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SageMaker::Device

Tag

All content copied from https://docs.aws.amazon.com/.
