This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelBiasJobDefinition ConstraintsResource

The constraints resource for a monitoring job.

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

The Amazon S3 URI for the constraints resource.

_Required_: No

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterConfig

Csv

All content copied from https://docs.aws.amazon.com/.
