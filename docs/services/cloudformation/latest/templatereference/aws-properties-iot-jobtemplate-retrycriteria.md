This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::JobTemplate RetryCriteria

The criteria that determines how many retries are allowed for each failure type for a
job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FailureType" : String,
  "NumberOfRetries" : Integer
}

```

### YAML

```yaml

  FailureType: String
  NumberOfRetries: Integer

```

## Properties

`FailureType`

The type of job execution failures that can initiate a job retry.

_Required_: No

_Type_: String

_Allowed values_: `FAILED | TIMED_OUT | ALL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NumberOfRetries`

The number of retries allowed for a failure type for the job.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RateIncreaseCriteria

Tag

All content copied from https://docs.aws.amazon.com/.
