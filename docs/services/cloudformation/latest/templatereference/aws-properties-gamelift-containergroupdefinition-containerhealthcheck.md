This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerGroupDefinition ContainerHealthCheck

Instructions on when and how to check the health of a support container in a container
fleet. These properties override any Docker health checks that are set in the container image.
For more information on container health checks, see [HealthCheck command](../../../../reference/amazonecs/latest/apireference/api-healthcheck.md#ECS-Type-HealthCheck-command) in the _Amazon Elastic Container Service API_. Game server
containers don't have a health check parameter; Amazon GameLift Servers automatically handles health checks
for these containers.

The following example instructs the container to initiate a health check command every 60
seconds and wait 10 seconds for it to succeed. If it fails, retry the command 3 times before
flagging the container as unhealthy. It also tells the container to wait 100 seconds after
launch before counting failed health checks.

`{"Command": [ "CMD-SHELL", "ps cax | grep "processmanager" || exit 1" ], "Interval":
        60, "Timeout": 10, "Retries": 3, "StartPeriod": 100 }`

**Part of:** [SupportContainerDefinition](../../../../reference/gamelift/latest/apireference/api-supportcontainerdefinition.md), [SupportContainerDefinitionInput](../../../../reference/gamelift/latest/apireference/api-supportcontainerdefinitioninput.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Command" : [ String, ... ],
  "Interval" : Integer,
  "Retries" : Integer,
  "StartPeriod" : Integer,
  "Timeout" : Integer
}

```

### YAML

```yaml

  Command:
    - String
  Interval: Integer
  Retries: Integer
  StartPeriod: Integer
  Timeout: Integer

```

## Properties

`Command`

A string array that specifies the command that the container runs to determine if it's
healthy.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `255 | 20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interval`

The time period (in seconds) between each health check.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Retries`

The number of times to retry a failed health check before flagging the container
unhealthy. The first run of the command does not count as a retry.

_Required_: No

_Type_: Integer

_Minimum_: `5`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartPeriod`

The optional grace period (in seconds) to give a container time to bootstrap before the
first failed health check counts toward the number of retries.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timeout`

The time period (in seconds) to wait for a health check to succeed before counting a
failed health check.

_Required_: No

_Type_: Integer

_Minimum_: `30`

_Maximum_: `60`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContainerEnvironment

ContainerMountPoint

All content copied from https://docs.aws.amazon.com/.
