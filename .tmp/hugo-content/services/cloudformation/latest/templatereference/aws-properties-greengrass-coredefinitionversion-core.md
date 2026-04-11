This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::CoreDefinitionVersion Core

A core is an
AWS IoT device that runs the AWS IoT Greengrass core software and manages
local processes for a Greengrass group. For more information, see [What Is\
AWS IoT Greengrass?](../../../greengrass/v1/developerguide/what-is-gg.md) in the _AWS IoT Greengrass Version 1 Developer Guide_.

In an CloudFormation template, the `Cores` property of the [`AWS::Greengrass::CoreDefinitionVersion`](../userguide/aws-resource-greengrass-coredefinitionversion.md) resource contains a list of `Core` property types. Currently, the list
can contain only one core.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateArn" : String,
  "Id" : String,
  "SyncShadow" : Boolean,
  "ThingArn" : String
}

```

### YAML

```yaml

  CertificateArn: String
  Id: String
  SyncShadow: Boolean
  ThingArn: String

```

## Properties

`CertificateArn`

The ARN of the device certificate for the core. This X.509 certificate is used to
authenticate the core with AWS IoT and AWS IoT Greengrass services.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Id`

A descriptive or arbitrary ID for the core. This value must be unique within the core
definition version. Maximum length is 128 characters with pattern
`[a-zA-Z0-9:_-]+`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SyncShadow`

Indicates whether the core's local shadow is synced with the cloud automatically. The
default is false.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ThingArn`

The Amazon Resource Name (ARN) of the core, which is an AWS IoT device
(thing).

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [Core](../../../../reference/greengrass/v1/apireference/definitions-core.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Greengrass::CoreDefinitionVersion

AWS::Greengrass::DeviceDefinition

All content copied from https://docs.aws.amazon.com/.
