This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::JobTemplate JobExecutionsRetryConfig

The configuration that determines how many retries are allowed for each failure type
for a job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RetryCriteriaList" : [ RetryCriteria, ... ]
}

```

### YAML

```yaml

  RetryCriteriaList:
    - RetryCriteria

```

## Properties

`RetryCriteriaList`

The list of criteria that determines how many retries are allowed for each failure
type for a job.

_Required_: No

_Type_: Array of [RetryCriteria](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-jobtemplate-retrycriteria.html)

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ExponentialRolloutRate

JobExecutionsRolloutConfig
