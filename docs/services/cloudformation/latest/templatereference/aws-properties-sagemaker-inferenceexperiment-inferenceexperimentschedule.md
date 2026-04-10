This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceExperiment InferenceExperimentSchedule

The start and end times of an inference experiment.

The maximum duration that you can set for an inference experiment is 30 days.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndTime" : String,
  "StartTime" : String
}

```

### YAML

```yaml

  EndTime: String
  StartTime: String

```

## Properties

`EndTime`

The timestamp at which the inference experiment ended or will end.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTime`

The timestamp at which the inference experiment started or will start.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EndpointMetadata

ModelInfrastructureConfig

All content copied from https://docs.aws.amazon.com/.
