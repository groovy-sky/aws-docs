This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ResourceDefinitionVersion ResourceInstance

A
local resource, machine learning resource, or secret resource. For more information, see
[Access Local\
Resources with Lambda Functions](https://docs.aws.amazon.com/greengrass/v1/developerguide/access-local-resources.html), [Perform Machine Learning\
Inference](https://docs.aws.amazon.com/greengrass/v1/developerguide/ml-inference.html), and [Deploy Secrets to the AWS IoT Greengrass Core](https://docs.aws.amazon.com/greengrass/v1/developerguide/secrets.html) in the _AWS IoT Greengrass Version 1 Developer Guide_.

In an CloudFormation template, the `Resources` property of the [`AWS::Greengrass::ResourceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html) resource contains a list of `ResourceInstance` property types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "Name" : String,
  "ResourceDataContainer" : ResourceDataContainer
}

```

### YAML

```yaml

  Id: String
  Name: String
  ResourceDataContainer:
    ResourceDataContainer

```

## Properties

`Id`

A descriptive or arbitrary ID for the resource. This value must be unique within the
resource definition version. Maximum length is 128 characters with pattern
`[a-zA-Z0-9:_-]+`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The descriptive resource name, which is displayed on the AWS IoT Greengrass console.
Maximum length 128 characters with pattern \[a-zA-Z0-9:\_-\]+. This must be unique within a
Greengrass group.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceDataContainer`

A container for resource data. The container takes only one of the following supported
resource data types: `LocalDeviceResourceData`,
`LocalVolumeResourceData`,
`SageMakerMachineLearningModelResourceData`,
`S3MachineLearningModelResourceData`, or
`SecretsManagerSecretResourceData`.

###### Note

Only one resource type can be defined for a `ResourceDataContainer`
instance.

_Required_: Yes

_Type_: [ResourceDataContainer](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [Resource](https://docs.aws.amazon.com/greengrass/v1/apireference/definitions-resource.html)
in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResourceDownloadOwnerSetting

S3MachineLearningModelResourceData
