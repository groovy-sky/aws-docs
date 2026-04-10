This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::JobTemplate RateIncreaseCriteria

Allows you to define a criteria to initiate the increase in rate of rollout for a
job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NumberOfNotifiedThings" : Integer,
  "NumberOfSucceededThings" : Integer
}

```

### YAML

```yaml

  NumberOfNotifiedThings: Integer
  NumberOfSucceededThings: Integer

```

## Properties

`NumberOfNotifiedThings`

The threshold for number of notified things that will initiate the increase in rate
of rollout.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NumberOfSucceededThings`

The threshold for number of succeeded things that will initiate the increase in
rate of rollout.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PresignedUrlConfig

RetryCriteria

All content copied from https://docs.aws.amazon.com/.
