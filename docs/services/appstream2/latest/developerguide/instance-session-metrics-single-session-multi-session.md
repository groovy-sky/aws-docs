# Instance and Session Performance Metrics for Single-session and Multi-session Fleets

The following are instance and session performance metrics for single-session and multi-session fleets.

MetricDescriptionDimensionsStatisticsUnits`CpuUtilizationInstance`

The percentage of allocated compute units that are currently
in use on the instance.

\[Fleet\]

\[UserId\]

\[FleetName, InstanceId\]

\[FleetName, InstanceId, SessionId, UserId\]

Average, Minimum, MaximumPercent`MemoryUtilizationInstance`

The percentage of allocated physical memory units that are
currently in use on the instance.

\[Fleet\]

\[UserId\]

\[FleetName, InstanceId\]

\[FleetName, InstanceId, SessionId, UserId\]

Average, Minimum, MaximumPercent`PagingFileUtilizationInstance`

The percentage of the paging file that is currently in use to
extend the Memory (RAM) capacity.

\[Fleet\]

\[UserId\]

\[FleetName, InstanceId\]

\[FleetName, InstanceId, SessionId, UserId\]

Average, Minimum, MaximumPercent`DiskUtilizationInstance`

The percentage of disk units that are currently in use to run
programs and carry out tasks on the instance.

\[Fleet\]

\[UserId\]

\[FleetName, InstanceId\]

\[FleetName, InstanceId, SessionId, UserId\]

Average, Minimum, MaximumPercent`CpuUtilizationSession`

The percentage of allocated compute units that are currently
in use by the session.

\[Fleet\]

\[UserId\]

\[FleetName, InstanceId, SessionId\]

\[FleetName, InstanceId, SessionId, UserId\]

Average, Minimum, MaximumPercent`MemoryUtilizationSession`

The percentage of allocated physical memory units that are
currently in use by the session.

\[Fleet\]

\[UserId\]

\[FleetName, InstanceId, SessionId\]

\[FleetName, InstanceId, SessionId, UserId\]

Average, Minimum, MaximumPercent`DiskReadOperations`

Amount of disk reads per instance

\[Fleet\]

\[UserId\]

\[FleetName, InstanceId\]

\[FleetName, InstanceId, SessionId, UserId\]

Average, Minimum, MaximumCount`DiskWriteOperations`

Amount of disk write per instance

\[Fleet\]

\[UserId\]

\[FleetName, InstanceId\]

\[FleetName, InstanceId, SessionId, UserId\]

Average, Minimum, MaximumCount`InSessionLatency`

Round trip time between WorkSpaces Application server and client measured at p90

\[Fleet\]

\[UserId\]

\[FleetName, InstanceId, SessionId\]

\[FleetName, InstanceId, SessionId, UserId\]

Average, Minimum, MaximumMilliseconds`FramesPerSecond`

Frames per second for the specific session

\[Fleet\]

\[UserId\]

\[FleetName, InstanceId, SessionId\]

\[FleetName, InstanceId, SessionId, UserId\]

Average, Minimum, MaximumCount`Bandwidth`

Amount of data transferred between WorkSpaces Applications service and Client during the session.

\[Fleet\]

\[UserId\]

\[FleetName, InstanceId, SessionId\]

\[FleetName, InstanceId, SessionId, UserId\]

Average, Minimum, MaximumKilobits/Second

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Fleet Usage Metrics for Multi-session Fleets

Dimensions for Amazon WorkSpaces Applications Metrics

All content copied from https://docs.aws.amazon.com/.
