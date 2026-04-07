This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::JobTemplate TimeoutConfig

Specifies the amount of time each device has to finish its execution of the job. A
timer is started when the job execution status is set to `IN_PROGRESS`. If
the job execution status is not set to another terminal state before the timer expires,
it will be automatically set to `TIMED_OUT`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InProgressTimeoutInMinutes" : Integer
}

```

### YAML

```yaml

  InProgressTimeoutInMinutes: Integer

```

## Properties

`InProgressTimeoutInMinutes`

Specifies the amount of time, in minutes, this device has to finish execution of this
job. The timeout interval can be anywhere between 1 minute and 7 days (1 to 10080 minutes).
The in progress timer can't be updated and will apply to all job executions for the job.
Whenever a job execution remains in the IN\_PROGRESS status for longer than this interval,
the job execution will fail and switch to the terminal `TIMED_OUT`
status.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10080`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::IoT::Logging
