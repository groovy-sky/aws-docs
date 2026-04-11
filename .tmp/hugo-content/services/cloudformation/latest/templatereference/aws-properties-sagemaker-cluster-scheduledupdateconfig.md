This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster ScheduledUpdateConfig

The configuration object of the schedule that SageMaker follows when updating the AMI.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeploymentConfig" : DeploymentConfig,
  "ScheduleExpression" : String
}

```

### YAML

```yaml

  DeploymentConfig:
    DeploymentConfig
  ScheduleExpression: String

```

## Properties

`DeploymentConfig`

The configuration to use when updating the AMI versions.

_Required_: No

_Type_: [DeploymentConfig](aws-properties-sagemaker-cluster-deploymentconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleExpression`

A cron expression that specifies the schedule that SageMaker follows when updating the
AMI.

_Required_: Yes

_Type_: String

_Pattern_: `cron\((?:[0-5][0-9]|[0-9]|) (?:[01][0-9]|2[0-3]|[0-9]) (?:[1-9]|0[1-9]|[12][0-9]|3[01]|\?) (?:[1-9]|0[1-9]|1[0-2]|\*|\*/(?:[1-9]|1[0-2])|) (?:MON|TUE|WED|THU|FRI|SAT|SUN|[1-7]|\?|L|(?:[1-7]#[1-5])|(?:[1-7]L)) (?:20[2-9][0-9]|\*|)\)`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RollingUpdatePolicy

Tag

All content copied from https://docs.aws.amazon.com/.
