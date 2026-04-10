This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Endpoint AutoRollbackConfig

Automatic rollback configuration for handling endpoint deployment failures and
recovery.

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

List of CloudWatch alarms in your account that are configured to monitor metrics on an
endpoint. If any alarms are tripped during a deployment, SageMaker rolls back the
deployment.

_Required_: Yes

_Type_: Array of [Alarm](aws-properties-sagemaker-endpoint-alarm.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Alarm

BlueGreenUpdatePolicy

All content copied from https://docs.aws.amazon.com/.
