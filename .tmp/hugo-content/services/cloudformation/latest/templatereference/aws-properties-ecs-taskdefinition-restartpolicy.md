This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskDefinition RestartPolicy

You can enable a restart policy for each container defined in your task definition, to
overcome transient failures faster and maintain task availability. When you enable a
restart policy for a container, Amazon ECS can restart the container if it exits,
without needing to replace the task. For more information, see [Restart\
individual containers in Amazon ECS tasks with container restart policies](../../../amazonecs/latest/developerguide/container-restart-policy.md) in
the _Amazon Elastic Container Service Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "IgnoredExitCodes" : [ Integer, ... ],
  "RestartAttemptPeriod" : Integer
}

```

### YAML

```yaml

  Enabled: Boolean
  IgnoredExitCodes:
    - Integer
  RestartAttemptPeriod: Integer

```

## Properties

`Enabled`

Specifies whether a restart policy is enabled for the container.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IgnoredExitCodes`

A list of exit codes that Amazon ECS will ignore and not attempt a restart on. You can
specify a maximum of 50 container exit codes. By default, Amazon ECS does not ignore any
exit codes.

_Required_: No

_Type_: Array of Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RestartAttemptPeriod`

A period of time (in seconds) that the container must run for before a restart can be
attempted. A container can be restarted only once every
`restartAttemptPeriod` seconds. If a container isn't able to run for this
time period and exits early, it will not be restarted. You can set a minimum
`restartAttemptPeriod` of 60 seconds and a maximum
`restartAttemptPeriod` of 1800 seconds. By default, a container must run
for 300 seconds before it can be restarted.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceRequirement

RuntimePlatform

All content copied from https://docs.aws.amazon.com/.
