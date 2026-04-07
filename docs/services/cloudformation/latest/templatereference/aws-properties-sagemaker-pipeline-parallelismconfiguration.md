This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Pipeline ParallelismConfiguration

Configuration that controls the parallelism of the pipeline.
By default, the parallelism configuration specified applies to all
executions of the pipeline unless overridden.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxParallelExecutionSteps" : Integer
}

```

### YAML

```yaml

  MaxParallelExecutionSteps: Integer

```

## Properties

`MaxParallelExecutionSteps`

The max number of steps that can be executed in parallel.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SageMaker::Pipeline

PipelineDefinition
