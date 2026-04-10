This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::Group GroupVersion

A group version in AWS IoT Greengrass, which references of a core definition version, device definition version,
subscription definition version, and other version types that contain the components you
want to deploy to a Greengrass core device. The group version must reference a core
definition version that contains one core. Other version types are optionally included,
depending on your business need.

In an CloudFormation template, `GroupVersion` is the property type of the
`InitialVersion` property in the [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectorDefinitionVersionArn" : String,
  "CoreDefinitionVersionArn" : String,
  "DeviceDefinitionVersionArn" : String,
  "FunctionDefinitionVersionArn" : String,
  "LoggerDefinitionVersionArn" : String,
  "ResourceDefinitionVersionArn" : String,
  "SubscriptionDefinitionVersionArn" : String
}

```

### YAML

```yaml

  ConnectorDefinitionVersionArn: String
  CoreDefinitionVersionArn: String
  DeviceDefinitionVersionArn: String
  FunctionDefinitionVersionArn: String
  LoggerDefinitionVersionArn: String
  ResourceDefinitionVersionArn: String
  SubscriptionDefinitionVersionArn: String

```

## Properties

`ConnectorDefinitionVersionArn`

The Amazon Resource Name (ARN) of the connector definition version that contains the
connectors you want to deploy with the group version.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CoreDefinitionVersionArn`

The ARN of the core definition version that contains the core you want to deploy with
the group version. Currently, the core definition version can contain only one core.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeviceDefinitionVersionArn`

The ARN of the device definition version that contains the devices you want to deploy
with the group version.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FunctionDefinitionVersionArn`

The ARN of the function definition version that contains the functions you want to
deploy with the group version.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LoggerDefinitionVersionArn`

The ARN of the logger definition version that contains the loggers you want to deploy
with the group version.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceDefinitionVersionArn`

The ARN of the resource definition version that contains the resources you want to
deploy with the group version.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubscriptionDefinitionVersionArn`

The ARN of the subscription definition version that contains the subscriptions you want
to deploy with the group version.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [GroupVersion](../../../../reference/greengrass/v1/apireference/definitions-groupversion.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Greengrass::Group

AWS::Greengrass::GroupVersion

All content copied from https://docs.aws.amazon.com/.
