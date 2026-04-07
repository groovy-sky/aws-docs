This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ImagePipeline Schedule

A schedule configures when and how often a pipeline will automatically create a new
image.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoDisablePolicy" : AutoDisablePolicy,
  "PipelineExecutionStartCondition" : String,
  "ScheduleExpression" : String
}

```

### YAML

```yaml

  AutoDisablePolicy:
    AutoDisablePolicy
  PipelineExecutionStartCondition: String
  ScheduleExpression: String

```

## Properties

`AutoDisablePolicy`

The policy that configures when Image Builder should automatically disable a pipeline that
is failing.

_Required_: No

_Type_: [AutoDisablePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-imagepipeline-autodisablepolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PipelineExecutionStartCondition`

The condition configures when the pipeline should trigger a new image build. When the
`pipelineExecutionStartCondition` is set to
`EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE`, and you use semantic version
filters on the base image or components in your image recipe, Image Builder will build a
new image only when there are new versions of the image or components in your recipe that
match the semantic version filter. When it is set to `EXPRESSION_MATCH_ONLY`, it
will build a new image every time the CRON expression matches the current time. For semantic
version syntax, see [CreateComponent](https://docs.aws.amazon.com/imagebuilder/latest/APIReference/API_CreateComponent.html)
in the _Image Builder API Reference_.

_Required_: No

_Type_: String

_Allowed values_: `EXPRESSION_MATCH_ONLY | EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleExpression`

The cron expression determines how often EC2 Image Builder evaluates your
`pipelineExecutionStartCondition`.

For information on how to format a cron expression in Image Builder, see [Use\
cron expressions in EC2 Image Builder](https://docs.aws.amazon.com/imagebuilder/latest/userguide/image-builder-cron.html).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Create \
an image pipeline](https://docs.aws.amazon.com/imagebuilder/latest/userguide/managing-image-builder-cli.html#image-builder-cli-create-image-pipeline) in the _Image Builder User Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PipelineLoggingConfiguration

WorkflowConfiguration
