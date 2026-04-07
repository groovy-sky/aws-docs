# InstanceStatusEvent

Describes a scheduled event for an instance.

## Contents

**code**

The event code.

Type: String

Valid Values: `instance-reboot | system-reboot | system-maintenance | instance-retirement | instance-stop`

Required: No

**description**

A description of the event.

After a scheduled event is completed, it can still be described for up to a week. If
the event has been completed, this description starts with the following text:
\[Completed\].

Type: String

Required: No

**instanceEventId**

The ID of the event.

Type: String

Required: No

**notAfter**

The latest scheduled end time for the event.

Type: Timestamp

Required: No

**notBefore**

The earliest scheduled start time for the event.

Type: Timestamp

Required: No

**notBeforeDeadline**

The deadline for starting the event.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/instancestatusevent.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/instancestatusevent.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/instancestatusevent.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InstanceStatusDetails

InstanceStatusSummary
