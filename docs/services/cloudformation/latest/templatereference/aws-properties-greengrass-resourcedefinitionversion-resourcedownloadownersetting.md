This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ResourceDefinitionVersion ResourceDownloadOwnerSetting

The owner setting for a downloaded machine learning resource. For more information, see
[Access Machine Learning\
Resources from Lambda Functions](../../../greengrass/v1/developerguide/access-ml-resources.md) in the _AWS IoT Greengrass Version 1 Developer Guide_.

In an CloudFormation template, `ResourceDownloadOwnerSetting` is the
property type of the `OwnerSetting` property for the [`S3MachineLearningModelResourceData`](../userguide/aws-properties-greengrass-resourcedefinitionversion-s3machinelearningmodelresourcedata.md) and [`SageMakerMachineLearningModelResourceData`](../userguide/aws-properties-greengrass-resourcedefinitionversion-sagemakermachinelearningmodelresourcedata.md) property types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GroupOwner" : String,
  "GroupPermission" : String
}

```

### YAML

```yaml

  GroupOwner: String
  GroupPermission: String

```

## Properties

`GroupOwner`

The group owner of the machine learning resource. This is the group ID (GID) of an
existing Linux OS group on the system. The group's permissions are added to the Lambda process.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroupPermission`

The permissions that the group owner has to the machine learning resource. Valid values
are `rw` (read-write) or `ro` (read-only).

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceDataContainer

ResourceInstance

All content copied from https://docs.aws.amazon.com/.
