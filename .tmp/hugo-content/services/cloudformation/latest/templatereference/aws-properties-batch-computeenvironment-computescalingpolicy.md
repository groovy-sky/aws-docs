This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::ComputeEnvironment ComputeScalingPolicy

An object that represents a scaling policy for a compute environment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MinScaleDownDelayMinutes" : Integer
}

```

### YAML

```yaml

  MinScaleDownDelayMinutes: Integer

```

## Properties

`MinScaleDownDelayMinutes`

The minimum time (in minutes) that AWS Batch keeps instances running in the compute environment
after their jobs complete. For each instance, the delay period begins when the last job finishes.
If no new jobs are placed on the instance during this delay, AWS Batch terminates the instance once
the delay expires.

Valid Range: Minimum value of 20. Maximum value of 10080. Use 0 to unset and disable the scale down delay.

###### Note

The scale down delay does not apply to:

- Instances being replaced during infrastructure updates

- Newly launched instances that have not yet run any jobs

- Spot instances reclaimed due to interruption

_Required_: No

_Type_: Integer

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComputeResources

Ec2ConfigurationObject

All content copied from https://docs.aws.amazon.com/.
