This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelExplainabilityJobDefinition ModelExplainabilityAppSpecification

Docker container image configuration object for the model explainability job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConfigUri" : String,
  "Environment" : {Key: Value, ...},
  "ImageUri" : String
}

```

### YAML

```yaml

  ConfigUri: String
  Environment:
    Key: Value
  ImageUri: String

```

## Properties

`ConfigUri`

JSON formatted Amazon S3 file that defines explainability parameters. For more
information on this JSON configuration file, see [Configure model explainability parameters](https://docs.aws.amazon.com/sagemaker/latest/dg/clarify-config-json-monitor-model-explainability-parameters.html).

_Required_: Yes

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Environment`

Sets the environment variables in the Docker container.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z_][a-zA-Z0-9_]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ImageUri`

The container image to be run by the model explainability job.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Json

ModelExplainabilityBaselineConfig
