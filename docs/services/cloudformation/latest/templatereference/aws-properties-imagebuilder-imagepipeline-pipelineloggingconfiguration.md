This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ImagePipeline PipelineLoggingConfiguration

The logging configuration that's defined for pipeline execution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ImageLogGroupName" : String,
  "PipelineLogGroupName" : String
}

```

### YAML

```yaml

  ImageLogGroupName: String
  PipelineLogGroupName: String

```

## Properties

`ImageLogGroupName`

The log group name that Image Builder uses for image creation. If not specified, the log group
name defaults to `/aws/imagebuilder/image-name`.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-_/\.]{1,512}$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PipelineLogGroupName`

The log group name that Image Builder uses for the log output during creation of a new pipeline.
If not specified, the pipeline log group name defaults to
`/aws/imagebuilder/pipeline/pipeline-name`.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-_/\.]{1,512}$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageTestsConfiguration

Schedule

All content copied from https://docs.aws.amazon.com/.
