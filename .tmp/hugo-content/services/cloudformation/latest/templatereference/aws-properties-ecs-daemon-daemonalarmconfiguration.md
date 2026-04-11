This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Daemon DaemonAlarmConfiguration

The CloudWatch alarm configuration for a daemon. When enabled, CloudWatch alarms
determine whether a daemon deployment has failed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlarmNames" : [ String, ... ],
  "Enable" : Boolean
}

```

### YAML

```yaml

  AlarmNames:
    - String
  Enable: Boolean

```

## Properties

`AlarmNames`

The CloudWatch alarm names to monitor during a daemon deployment.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enable`

Determines whether to use the CloudWatch alarm option in the daemon deployment
process. The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ECS::Daemon

DaemonDeploymentConfiguration

All content copied from https://docs.aws.amazon.com/.
