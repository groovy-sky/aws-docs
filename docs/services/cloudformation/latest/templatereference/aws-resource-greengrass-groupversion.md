This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::GroupVersion

The `AWS::Greengrass::GroupVersion` resource represents a group version in
AWS IoT Greengrass. A group version references a core definition version, device
definition version, subscription definition version, and other version types that contain
the components you want to deploy to a Greengrass core device. The group version must
reference a core definition version that contains one core. Other version types are
optionally included, depending on your business need.

###### Note

To create a group version, you must specify the ID of the group that you want to
associate with the version. For information about creating a group, see [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Greengrass::GroupVersion",
  "Properties" : {
      "ConnectorDefinitionVersionArn" : String,
      "CoreDefinitionVersionArn" : String,
      "DeviceDefinitionVersionArn" : String,
      "FunctionDefinitionVersionArn" : String,
      "GroupId" : String,
      "LoggerDefinitionVersionArn" : String,
      "ResourceDefinitionVersionArn" : String,
      "SubscriptionDefinitionVersionArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Greengrass::GroupVersion
Properties:
  ConnectorDefinitionVersionArn: String
  CoreDefinitionVersionArn: String
  DeviceDefinitionVersionArn: String
  FunctionDefinitionVersionArn: String
  GroupId: String
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

`GroupId`

The ID of the group associated with this version. This value is a GUID.

_Required_: Yes

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

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the group version, such as
`arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/groups/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## See also

- [CreateGroupVersion](../../../../reference/greengrass/v1/apireference/creategroupversion-post.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GroupVersion

AWS::Greengrass::LoggerDefinition

All content copied from https://docs.aws.amazon.com/.
