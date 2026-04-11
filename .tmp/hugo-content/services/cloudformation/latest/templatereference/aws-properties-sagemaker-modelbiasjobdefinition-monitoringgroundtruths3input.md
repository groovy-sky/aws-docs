This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelBiasJobDefinition MonitoringGroundTruthS3Input

The ground truth labels for the dataset used for the monitoring job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Uri" : String
}

```

### YAML

```yaml

  S3Uri: String

```

## Properties

`S3Uri`

The address of the Amazon S3 location of the ground truth labels.

_Required_: Yes

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelBiasJobInput

MonitoringOutput

All content copied from https://docs.aws.amazon.com/.
