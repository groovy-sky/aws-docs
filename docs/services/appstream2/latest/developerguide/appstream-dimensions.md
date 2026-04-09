# Fleet Usage Metrics for Single-session Fleets

The following are fleet usage metrics for single-session fleets.

MetricDescriptionDimensionsStatisticsUnits`ActualCapacity`

The total number of instances that are available for streaming or are currently streaming.

```nohighlight

ActualCapacity = AvailableCapacity + InUseCapacity
```

\[Fleet\]Average, Minimum, MaximumCount`AvailableCapacity`

The number of idle instances currently available for user sessions.

```nohighlight

AvailableCapacity = ActualCapacity - InUseCapacity
```

\[Fleet\]Average, Minimum, MaximumCount`CapacityUtilization`

The percentage of instances in a fleet that are being used, using
the following formula.

```nohighlight

CapacityUtilization = (InUseCapacity/ActualCapacity) * 100
```

Monitoring this metric helps with decisions about increasing or
decreasing the value of a fleet's desired capacity.

\[Fleet\]Average, Minimum, MaximumPercent`DesiredCapacity`

The total number of instances that are either running or pending. This
represents the total number of concurrent streaming sessions your
fleet can support in a steady state.

```nohighlight

DesiredCapacity = ActualCapacity + PendingCapacity
```

\[Fleet\]Average, Minimum, MaximumCount`InUseCapacity`

The number of instances currently being used for streaming sessions.
One `InUseCapacity` count represents one streaming session.

\[Fleet\]Average, Minimum, MaximumCount`PendingCapacity`

The number of instances being provisioned by WorkSpaces Applications. Represents the
additional number of streaming sessions the fleet can support after
provisioning is complete. When provisioning starts, it usually takes
10-20 minutes for an instance to become available for streaming.

\[Fleet\]Average, Minimum, MaximumCount`RunningCapacity`

The total number of instances currently running. Represents the number
of concurrent streaming sessions that can be supported by the fleet
in its current state.

This metric is provided for Always-On fleets only, and has the same value
as the `ActualCapacity` metric.

\[Fleet\]Average, Minimum, MaximumCount`InsufficientCapacityError`

The number of session requests rejected due to lack of capacity.

You can set alarms to use this metric to be notified of users
waiting for streaming sessions.

\[Fleet\]Average, Minimum, Maximum, SumCount`InsufficientConcurrencyLimitError`

The number of Elastic fleet session requests rejected due to
reaching max concurrent streaming capacity.

You can set alarms to use this metric to be notified of users
waiting for streaming sessions.

\[Fleet\]Average, Minimum, Maximum, SumCount

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkSpaces Applications Metrics and Dimensions

Fleet Usage Metrics for Multi-session Fleets

All content copied from https://docs.aws.amazon.com/.
