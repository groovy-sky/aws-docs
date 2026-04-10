This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage TransformJobDefinition

Defines the input needed to run a transform job using the inference specification
specified in the algorithm.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BatchStrategy" : String,
  "Environment" : {Key: Value, ...},
  "MaxConcurrentTransforms" : Integer,
  "MaxPayloadInMB" : Integer,
  "TransformInput" : TransformInput,
  "TransformOutput" : TransformOutput,
  "TransformResources" : TransformResources
}

```

### YAML

```yaml

  BatchStrategy: String
  Environment:
    Key: Value
  MaxConcurrentTransforms: Integer
  MaxPayloadInMB: Integer
  TransformInput:
    TransformInput
  TransformOutput:
    TransformOutput
  TransformResources:
    TransformResources

```

## Properties

`BatchStrategy`

A string that determines the number of records included in a single mini-batch.

`SingleRecord` means only one record is used per mini-batch.
`MultiRecord` means a mini-batch is set to contain as many records that
can fit within the `MaxPayloadInMB` limit.

_Required_: No

_Type_: String

_Allowed values_: `MultiRecord | SingleRecord`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Environment`

The environment variables to set in the Docker container. We support up to 16 key and
values entries in the map.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z_][a-zA-Z0-9_]*`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxConcurrentTransforms`

The maximum number of parallel requests that can be sent to each instance in a
transform job. The default value is 1.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxPayloadInMB`

The maximum payload size allowed, in MB. A payload is the data portion of a record
(without metadata).

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransformInput`

A description of the input source and the way the transform job consumes it.

_Required_: Yes

_Type_: [TransformInput](aws-properties-sagemaker-modelpackage-transforminput.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransformOutput`

Identifies the Amazon S3 location where you want Amazon SageMaker to save the results from the
transform job.

_Required_: Yes

_Type_: [TransformOutput](aws-properties-sagemaker-modelpackage-transformoutput.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransformResources`

Identifies the ML compute instances for the transform job.

_Required_: Yes

_Type_: [TransformResources](aws-properties-sagemaker-modelpackage-transformresources.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TransformInput

TransformOutput

All content copied from https://docs.aws.amazon.com/.
