This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerFleet ManagedCapacityConfiguration

Use ManagedCapacityConfiguration with the "SCALE\_TO\_AND\_FROM\_ZERO" ZeroCapacityStrategy to enable Amazon
GameLift Servers to fully manage the MinSize value, switching between 0 and 1 based on game session
activity. This is ideal for eliminating compute costs during periods of no game activity.
It is particularly beneficial during development when you're away from your desk, iterating on builds
for extended periods, in production environments serving low-traffic locations, or for games with long,
predictable downtime windows. By automatically managing capacity between 0 and 1 instances, you avoid paying
for idle instances while maintaining the ability to serve game sessions when demand arrives. Note that while
scale-out is triggered immediately upon receiving a game session request, actual game session availability
depends on your server process startup time, so this approach works best with multi-location Fleets where
cold-start latency is tolerable. With a "MANUAL" ZeroCapacityStrategy Amazon GameLift Servers will not
modify Fleet MinSize values automatically and will not scale out from zero instances in response to game
sessions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ScaleInAfterInactivityMinutes" : Integer,
  "ZeroCapacityStrategy" : String
}

```

### YAML

```yaml

  ScaleInAfterInactivityMinutes: Integer
  ZeroCapacityStrategy: String

```

## Properties

`ScaleInAfterInactivityMinutes`

Length of time, in minutes, that Amazon GameLift Servers will wait before scaling in your MinSize and
DesiredInstances to 0 after a period with no game session activity. Default: 30 minutes.

_Required_: No

_Type_: Integer

_Minimum_: `5`

_Maximum_: `1440`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ZeroCapacityStrategy`

The strategy Amazon GameLift Servers will use to automatically scale your capacity to and from zero
instances in response to game session activity. Game session activity refers to any active running sessions
or game session requests.

Possible ZeroCapacityStrategy types include:

- **MANUAL** \-\- (default value) Amazon GameLift Servers will not update
capacity to and from zero on your behalf.

- **SCALE\_TO\_AND\_FROM\_ZERO** \-\- Amazon GameLift Servers will automatically
scale out MinSize and DesiredInstances from 0 to 1 in response to a game session request, and will scale
in MinSize and DesiredInstances to 0 after a period with no game session activity. The duration of
this scale in period can be configured using ScaleInAfterInactivityMinutes.

_Required_: Yes

_Type_: String

_Allowed values_: `SCALE_TO_AND_FROM_ZERO | MANUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogConfiguration

ScalingPolicy

All content copied from https://docs.aws.amazon.com/.
