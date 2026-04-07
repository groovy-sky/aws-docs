This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConvert::JobTemplate AccelerationSettings

Accelerated transcoding can significantly speed up jobs with long, visually complex
content. Outputs that use this feature incur pro-tier pricing. For information about
feature limitations, For more information, see
[Job\
Limitations for Accelerated Transcoding in AWS Elemental MediaConvert](https://docs.aws.amazon.com/mediaconvert/latest/ug/job-requirements.html) in the _AWS Elemental MediaConvert User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mode" : String
}

```

### YAML

```yaml

  Mode: String

```

## Properties

`Mode`

Specify the conditions when the service will run your job with accelerated transcoding.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::MediaConvert::JobTemplate

HopDestination
