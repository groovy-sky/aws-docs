This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ResourceDefinition ResourceDataContainer

A
container for resource data, which defines the resource type. The container takes only one
of the following supported resource data types: `LocalDeviceResourceData`,
`LocalVolumeResourceData`,
`SageMakerMachineLearningModelResourceData`,
`S3MachineLearningModelResourceData`, or
`SecretsManagerSecretResourceData`.

###### Note

Only one resource type can be defined for a `ResourceDataContainer`
instance.

In
an CloudFormation template, `ResourceDataContainer` is a property of the
[`ResourceInstance`](../userguide/aws-properties-greengrass-resourcedefinition-resourceinstance.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LocalDeviceResourceData" : LocalDeviceResourceData,
  "LocalVolumeResourceData" : LocalVolumeResourceData,
  "S3MachineLearningModelResourceData" : S3MachineLearningModelResourceData,
  "SageMakerMachineLearningModelResourceData" : SageMakerMachineLearningModelResourceData,
  "SecretsManagerSecretResourceData" : SecretsManagerSecretResourceData
}

```

### YAML

```yaml

  LocalDeviceResourceData:
    LocalDeviceResourceData
  LocalVolumeResourceData:
    LocalVolumeResourceData
  S3MachineLearningModelResourceData:
    S3MachineLearningModelResourceData
  SageMakerMachineLearningModelResourceData:
    SageMakerMachineLearningModelResourceData
  SecretsManagerSecretResourceData:
    SecretsManagerSecretResourceData

```

## Properties

`LocalDeviceResourceData`

Settings for a local device resource.

_Required_: No

_Type_: [LocalDeviceResourceData](aws-properties-greengrass-resourcedefinition-localdeviceresourcedata.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LocalVolumeResourceData`

Settings for a local volume resource.

_Required_: No

_Type_: [LocalVolumeResourceData](aws-properties-greengrass-resourcedefinition-localvolumeresourcedata.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3MachineLearningModelResourceData`

Settings for a machine learning resource stored in Amazon S3.

_Required_: No

_Type_: [S3MachineLearningModelResourceData](aws-properties-greengrass-resourcedefinition-s3machinelearningmodelresourcedata.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SageMakerMachineLearningModelResourceData`

Settings for a machine learning resource saved as an SageMaker AI training
job.

_Required_: No

_Type_: [SageMakerMachineLearningModelResourceData](aws-properties-greengrass-resourcedefinition-sagemakermachinelearningmodelresourcedata.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecretsManagerSecretResourceData`

Settings for a secret resource.

_Required_: No

_Type_: [SecretsManagerSecretResourceData](aws-properties-greengrass-resourcedefinition-secretsmanagersecretresourcedata.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [ResourceDataContainer](../../../../reference/greengrass/v1/apireference/definitions-resourcedatacontainer.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LocalVolumeResourceData

ResourceDefinitionVersion

All content copied from https://docs.aws.amazon.com/.
