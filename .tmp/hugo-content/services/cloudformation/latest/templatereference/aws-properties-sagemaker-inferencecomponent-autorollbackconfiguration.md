This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceComponent AutoRollbackConfiguration

Configuration for automatic rollback of the inference component deployment if issues are detected.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alarms" : [ Alarm, ... ]
}

```

### YAML

```yaml

  Alarms:
    - Alarm

```

## Properties

`Alarms`

List of CloudWatch alarms that trigger automatic rollback if they enter the ALARM state during
deployment.

_Required_: Yes

_Type_: Array of [Alarm](aws-properties-sagemaker-inferencecomponent-alarm.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Alarm

DeployedImage

All content copied from https://docs.aws.amazon.com/.
