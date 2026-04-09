# ComputeCapacityStatus

Describes the capacity status for a fleet.

## Contents

**Desired**

The desired number of streaming instances.

Type: Integer

Required: Yes

**ActiveUserSessions**

The number of user sessions currently being used for streaming sessions. This only applies to multi-session fleets.

Type: Integer

Required: No

**ActualUserSessions**

The total number of session slots that are available for streaming or are currently streaming.

ActualUserSessionCapacity = AvailableUserSessionCapacity + ActiveUserSessions

This only applies to multi-session fleets.

Type: Integer

Required: No

**Available**

The number of currently available instances that can be used to stream
sessions.

Type: Integer

Required: No

**AvailableUserSessions**

The number of idle session slots currently available for user sessions.

AvailableUserSessionCapacity = ActualUserSessionCapacity - ActiveUserSessions

This only applies to multi-session fleets.

Type: Integer

Required: No

**DesiredUserSessions**

The total number of sessions slots that are either running or pending. This represents the total number of concurrent streaming sessions your fleet can support in a steady state.

DesiredUserSessionCapacity = ActualUserSessionCapacity + PendingUserSessionCapacity

This only applies to multi-session fleets.

Type: Integer

Required: No

**Draining**

The number of instances in drain mode. This only applies to multi-session fleets.

Type: Integer

Required: No

**DrainModeActiveUserSessions**

The number of active user sessions on instances in drain mode. This only applies to multi-session fleets.

Type: Integer

Required: No

**DrainModeUnusedUserSessions**

The number of unused session slots on instances in drain mode that cannot be used for user session provisioning. This only applies to multi-session fleets.

Type: Integer

Required: No

**InUse**

The number of instances in use for streaming.

Type: Integer

Required: No

**Running**

The total number of simultaneous streaming instances that are running.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/computecapacitystatus.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/computecapacitystatus.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/computecapacitystatus.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComputeCapacity

DirectoryConfig

All content copied from https://docs.aws.amazon.com/.
