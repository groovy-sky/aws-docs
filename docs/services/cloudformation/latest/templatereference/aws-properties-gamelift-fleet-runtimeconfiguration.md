This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::Fleet RuntimeConfiguration

A collection of server process configurations that describe the set of processes to
run on each instance in a fleet. Server processes run either an executable in a custom
game build or a Realtime Servers script. GameLift launches the configured processes, manages their
life cycle, and replaces them as needed. Each instance checks regularly for an updated
runtime configuration.

A GameLift instance is limited to 50 processes running concurrently. To calculate the
total number of processes in a runtime configuration, add the values of the
`ConcurrentExecutions` parameter for each ServerProcess. Learn more about
[Running Multiple\
Processes on a Fleet](../../../gamelift/latest/developerguide/fleets-multiprocess.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GameSessionActivationTimeoutSeconds" : Integer,
  "MaxConcurrentGameSessionActivations" : Integer,
  "ServerProcesses" : [ ServerProcess, ... ]
}

```

### YAML

```yaml

  GameSessionActivationTimeoutSeconds: Integer
  MaxConcurrentGameSessionActivations: Integer
  ServerProcesses:
    - ServerProcess

```

## Properties

`GameSessionActivationTimeoutSeconds`

The maximum amount of time (in seconds) allowed to launch a new game session and have
it report ready to host players. During this time, the game session is in status
`ACTIVATING`. If the game session does not become active before the
timeout, it is ended and the game session status is changed to
`TERMINATED`.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxConcurrentGameSessionActivations`

The number of game sessions in status `ACTIVATING` to allow on an
instance or compute. This setting limits the instance resources that can be
used for new game activations at any one time.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerProcesses`

A collection of server process configurations that identify what server processes to
run on fleet computes.

_Required_: No

_Type_: Array of [ServerProcess](aws-properties-gamelift-fleet-serverprocess.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Create GameLift resources using Amazon CloudFront](../../../gamelift/latest/developerguide/resources-cloudformation.md) in the _Amazon_
_GameLift Developer Guide_

- [Deploy a GameLift fleet for a custom game build](../../../gamelift/latest/developerguide/fleets-creating.md) in the _Amazon_
_GameLift Developer Guide_

- [Deploy a Realtime\
Servers fleet](../../../gamelift/latest/developerguide/realtime-fleets-creating.md) in the _Amazon GameLift Developer Guide_

- [Run multiple processes\
on a fleet](../../../gamelift/latest/developerguide/fleets-multiprocess.md) in the _Amazon GameLift Developer Guide_

- [RuntimeConfiguration](../../../../reference/gamelift/latest/apireference/api-runtimeconfiguration.md) in the _Amazon GameLift API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceCreationLimitPolicy

ScalingPolicy

All content copied from https://docs.aws.amazon.com/.
