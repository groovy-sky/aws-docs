This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::Image ImagePipelineExecutionSettings

Contains settings for starting an image pipeline execution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeploymentId" : String,
  "OnUpdate" : Boolean
}

```

### YAML

```yaml

  DeploymentId: String
  OnUpdate: Boolean

```

## Properties

`DeploymentId`

The deployment identifier of the pipeline, utilized to initiate new image pipeline executions.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`OnUpdate`

Defines whether the pipeline should be executed upon pipeline updates. False by default.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageLoggingConfiguration

ImageScanningConfiguration

All content copied from https://docs.aws.amazon.com/.
