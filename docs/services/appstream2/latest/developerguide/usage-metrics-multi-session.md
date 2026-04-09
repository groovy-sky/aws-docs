# Fleet Usage Metrics for Multi-session Fleets

The following are fleet usage metrics for multi-session fleets.

MetricDescriptionDimensionsStatisticsUnits`CapacityUtilization`

The percentage of sessions in a fleet that are being used,
using the following formula.

```nohighlight

CapacityUtilization = (ActiveUserSessionCapacity / ActualUserSessionCapacity) * 100
```

Monitoring this metric helps with decisions about increasing
or decreasing the value of a fleet's desired
capacity.

\[Fleet\]Average, Minimum, MaximumPercent`ActualUserSessionCapacity`

The total number of session slots that are available for
streaming or are currently streaming.

```nohighlight

ActualUserSessionCapacity = AvailableUserSessionCapacity + ActiveUserSessionCapacity
```

\[Fleet\]Average, Minimum, MaximumCount`AvailableUserSessionCapacity`

The number of idle session slots currently available for user
sessions.

```nohighlight

AvailableUserSessionCapacity = ActualUserSessionCapacity - ActiveUserSessions
```

\[Fleet\]Average, Minimum, MaximumCount`DesiredUserSessionCapacity`

The total number of session slots that are either running or
pending. This represents the total number of concurrent
streaming sessions your fleet can support in a steady
state.

```nohighlight

DesiredUserSessionCapacity = ActualUserSessionCapacity + PendingUserSessionCapacity
```

\[Fleet\]Average, Minimum, MaximumCount`ActiveUserSessionCapacity`

The number of user sessions currently being used for streaming
sessions.

\[Fleet\]Average, Minimum, MaximumCount`PendingUserSessionCapacity`

The number of session slots being provisioned by WorkSpaces Applications.
Represents the additional number of streaming sessions the fleet
can support after provisioning is complete. When provisioning
starts, it usually takes 10-20 minutes for an instance to become
available for streaming.

\[Fleet\]Average, Minimum, MaximumCount`RunningUserSessionCapacity`

The total number of session slots currently that are available
for streaming or are currently streaming. Represents the number
of concurrent streaming sessions that can be supported by the
fleet in its current state.

This metric is provided for Always-On fleets only, and has the
same value as the `ActualUserSessionCapacity`
metric.

\[Fleet\]Average, Minimum, MaximumCount`DrainingCapacity`

The number of instances in Drain Mode.

\[Fleet\]Average, Minimum, MaximumCount`DrainModeActiveUserSessionCapacity`

The number of active user sessions on instances in Drain Mode.

\[Fleet\]Average, Minimum, MaximumCount`DrainModeUnusedUserSessionCapacity`

The number of unused session slots on instances in Drain Mode which cannot be used for user session provisioning.

\[Fleet\]Average, Minimum, MaximumCount

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Fleet Usage Metrics for Single-session Fleets

Instance and Session Performance Metrics for Single-session and Multi-session Fleets

All content copied from https://docs.aws.amazon.com/.
