This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::Fleet ServerProcess

A set of instructions for launching server processes on each instance in a fleet.
Server processes run either an executable in a custom game build or a Realtime Servers script.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConcurrentExecutions" : Integer,
  "LaunchPath" : String,
  "Parameters" : String
}

```

### YAML

```yaml

  ConcurrentExecutions: Integer
  LaunchPath: String
  Parameters: String

```

## Properties

`ConcurrentExecutions`

The number of server processes using this configuration that run concurrently on each
instance or compute.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchPath`

The location of a game build executable or Realtime script. Game builds and Realtime
scripts are installed on instances at the root:

- Windows (custom game builds only): `C:\game`. Example:
" `C:\game\MyGame\server.exe`"

- Linux: `/local/game`. Examples:
" `/local/game/MyGame/server.exe`" or
" `/local/game/MyRealtimeScript.js`"

###### Note

Amazon GameLift Servers doesn't support the use of setup scripts that launch the game executable.
For custom game builds, this parameter must indicate the executable that calls the
server SDK operations `initSDK()` and `ProcessReady()`.

_Required_: Yes

_Type_: String

_Pattern_: `^([Cc]:\\game\S+|/local/game/\S+)`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

An optional list of parameters to pass to the server executable or Realtime script on launch.

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: \[A-Za-z0-9\_:.+\\/\\\\\- =@{},?'\\\[\\\]"\]+

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

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

- [ServerProcess](../../../../reference/gamelift/latest/apireference/api-serverprocess.md) in the _Amazon GameLift API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScalingPolicy

Tag

All content copied from https://docs.aws.amazon.com/.
