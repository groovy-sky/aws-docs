This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup Alarm

The `Alarm` property type specifies a CloudWatch alarm to use for an
AWS CodeDeploy deployment group. The `Alarm` property of the [CodeDeploy DeploymentGroup AlarmConfiguration](../userguide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.md) property contains a list of
`Alarm` property types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String
}

```

### YAML

```yaml

  Name: String

```

## Properties

`Name`

The name of the alarm. Maximum length is 255 characters. Each alarm name can be used
only once in a list of alarms.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CodeDeploy::DeploymentGroup

AlarmConfiguration

All content copied from https://docs.aws.amazon.com/.
