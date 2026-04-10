This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ResourceDefinition SageMakerMachineLearningModelResourceData

Settings for an Secrets Manager machine learning resource. For more information, see
[Perform Machine Learning Inference](../../../greengrass/v1/developerguide/ml-inference.md) in the _AWS IoT Greengrass Version 1 Developer Guide_.

In an CloudFormation template,
`SageMakerMachineLearningModelResourceData` can be used in the [`ResourceDataContainer`](../userguide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationPath" : String,
  "OwnerSetting" : ResourceDownloadOwnerSetting,
  "SageMakerJobArn" : String
}

```

### YAML

```yaml

  DestinationPath: String
  OwnerSetting:
    ResourceDownloadOwnerSetting
  SageMakerJobArn: String

```

## Properties

`DestinationPath`

The absolute local path of the resource inside the Lambda
environment.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OwnerSetting`

The owner setting for the downloaded machine learning resource. For more information,
see [Access Machine Learning\
Resources from Lambda Functions](../../../greengrass/v1/developerguide/access-ml-resources.md) in the _AWS IoT Greengrass Version 1 Developer Guide_.

_Required_: No

_Type_: [ResourceDownloadOwnerSetting](aws-properties-greengrass-resourcedefinition-resourcedownloadownersetting.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SageMakerJobArn`

The Amazon Resource Name (ARN) of the Amazon SageMaker AI training job that
represents the source model.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [SageMakerMachineLearningModelResourceData](../../../../reference/greengrass/v1/apireference/definitions-sagemakermachinelearningmodelresourcedata.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3MachineLearningModelResourceData

SecretsManagerSecretResourceData

All content copied from https://docs.aws.amazon.com/.
