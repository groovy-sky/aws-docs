# SessionStatus

Contains information about the status of a session.

## Contents

**EndDateTime**

The date and time that the session ended.

Type: Timestamp

Required: No

**IdleSinceDateTime**

The date and time starting at which the session became idle. Can be empty if the
session is not currently idle.

Type: Timestamp

Required: No

**LastModifiedDateTime**

The most recent date and time that the session was modified.

Type: Timestamp

Required: No

**StartDateTime**

The date and time that the session started.

Type: Timestamp

Required: No

**State**

The state of the session. A description of each state follows.

`CREATING` \- The session is being started, including acquiring
resources.

`CREATED` \- The session has been started.

`IDLE` \- The session is able to accept a calculation.

`BUSY` \- The session is processing another task and is unable to accept a
calculation.

`TERMINATING` \- The session is in the process of shutting down.

`TERMINATED` \- The session and its resources are no longer running.

`DEGRADED` \- The session has no healthy coordinators.

`FAILED` \- Due to a failure, the session and its resources are no longer
running.

Type: String

Valid Values: `CREATING | CREATED | IDLE | BUSY | TERMINATING | TERMINATED | DEGRADED | FAILED`

Required: No

**StateChangeReason**

The reason for the session state change (for example, canceled because the session was
terminated).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/sessionstatus.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/sessionstatus.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/sessionstatus.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SessionStatistics

SessionSummary
