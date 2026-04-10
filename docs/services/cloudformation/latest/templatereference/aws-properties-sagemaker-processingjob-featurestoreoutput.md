This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ProcessingJob FeatureStoreOutput

Configuration for processing job outputs in Amazon SageMaker Feature Store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FeatureGroupName" : String
}

```

### YAML

```yaml

  FeatureGroupName: String

```

## Properties

`FeatureGroupName`

The name of the Amazon SageMaker FeatureGroup to use as the destination for processing job output.
Note that your processing script is responsible for putting records into your Feature
Store.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9]([_-]*[a-zA-Z0-9]){0,63}`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExperimentConfig

NetworkConfig

All content copied from https://docs.aws.amazon.com/.
