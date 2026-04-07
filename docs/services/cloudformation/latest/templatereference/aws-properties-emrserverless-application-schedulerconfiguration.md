This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application SchedulerConfiguration

The scheduler configuration for batch and streaming jobs running on this application. Supported with release labels emr-7.0.0 and above.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxConcurrentRuns" : Integer,
  "QueueTimeoutMinutes" : Integer
}

```

### YAML

```yaml

  MaxConcurrentRuns: Integer
  QueueTimeoutMinutes: Integer

```

## Properties

`MaxConcurrentRuns`

The maximum concurrent job runs on this application. If scheduler configuration is enabled on your application, the default value is 15. The valid range is 1 to 1000.

_Required_: No

_Type_: Integer

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`QueueTimeoutMinutes`

The maximum duration in minutes for the job in QUEUED state. If scheduler configuration is enabled on your application, the default value is 360 minutes (6 hours). The valid range is from 15 to 720.

_Required_: No

_Type_: Integer

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3MonitoringConfiguration

Tag
