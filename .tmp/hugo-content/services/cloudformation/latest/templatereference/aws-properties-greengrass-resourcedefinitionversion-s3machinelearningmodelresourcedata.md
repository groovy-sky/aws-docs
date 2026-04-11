This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ResourceDefinitionVersion S3MachineLearningModelResourceData

Settings for an Amazon S3 machine learning resource. For more information, see
[Perform Machine Learning Inference](../../../greengrass/v1/developerguide/ml-inference.md) in the _AWS IoT Greengrass Version 1 Developer Guide_.

In an CloudFormation template, `S3MachineLearningModelResourceData` can
be used in the [`ResourceDataContainer`](../userguide/aws-properties-greengrass-resourcedefinitionversion-resourcedatacontainer.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationPath" : String,
  "OwnerSetting" : ResourceDownloadOwnerSetting,
  "S3Uri" : String
}

```

### YAML

```yaml

  DestinationPath: String
  OwnerSetting:
    ResourceDownloadOwnerSetting
  S3Uri: String

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

_Type_: [ResourceDownloadOwnerSetting](aws-properties-greengrass-resourcedefinitionversion-resourcedownloadownersetting.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Uri`

The URI of the source model in an Amazon S3 bucket. The model package must be
in `tar.gz` or `.zip` format.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [S3MachineLearningModelResourceData](../../../../reference/greengrass/v1/apireference/definitions-s3machinelearningmodelresourcedata.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceInstance

SageMakerMachineLearningModelResourceData

All content copied from https://docs.aws.amazon.com/.
