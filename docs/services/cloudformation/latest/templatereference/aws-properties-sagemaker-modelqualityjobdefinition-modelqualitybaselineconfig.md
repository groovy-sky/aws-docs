This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelQualityJobDefinition ModelQualityBaselineConfig

Configuration for monitoring constraints and monitoring statistics. These baseline resources are
compared against the results of the current job from the series of jobs scheduled to collect data
periodically.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BaseliningJobName" : String,
  "ConstraintsResource" : ConstraintsResource
}

```

### YAML

```yaml

  BaseliningJobName: String
  ConstraintsResource:
    ConstraintsResource

```

## Properties

`BaseliningJobName`

The name of the job that performs baselining for the monitoring job.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConstraintsResource`

The constraints resource for a monitoring job.

_Required_: No

_Type_: [ConstraintsResource](aws-properties-sagemaker-modelqualityjobdefinition-constraintsresource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelQualityAppSpecification

ModelQualityJobInput

All content copied from https://docs.aws.amazon.com/.
