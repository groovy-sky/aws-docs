# Dimensions for Amazon WorkSpaces Applications Metrics

The `AWS/AppStream` namespace includes the following dimensions and dimension groups.

DimensionDescription`Fleet`Filters the metric data by name of the Fleet.`FleetName`Filters the metric data by name of the Fleet.`SessionId`Filters the metric data by session identifier.`InstanceId`Filters the metric data by instance identifier.`UserId`Filters the metric data by user identifier.

DimensionWhere Available in Amazon CloudWatch Metrics`[Fleet]`Fleet Metrics`[FleetName, InstanceId]`Fleet Instance Metrics`[FleetName, InstanceId, SessionId]`Fleet Session Metrics`[UserId]`UserId`[FleetName, InstanceId, SessionId, UserId]`FleetName, InstanceId, SessionId, UserId

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Instance and Session Performance Metrics for Single-session and Multi-session Fleets

Manage Multi-Session Fleet Instances

All content copied from https://docs.aws.amazon.com/.
