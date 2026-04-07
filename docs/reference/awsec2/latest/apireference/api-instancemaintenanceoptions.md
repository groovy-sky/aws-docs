# InstanceMaintenanceOptions

The maintenance options for the instance.

## Contents

**autoRecovery**

Provides information on the current automatic recovery behavior of your
instance.

Type: String

Valid Values: `disabled | default`

Required: No

**rebootMigration**

Specifies whether to attempt reboot migration during a user-initiated reboot of an
instance that has a scheduled `system-reboot` event:

- `default` \- Amazon EC2 attempts to migrate the instance to
new hardware (reboot migration). If successful, the `system-reboot`
event is cleared. If unsuccessful, an in-place reboot occurs and the event
remains scheduled.

- `disabled` \- Amazon EC2 keeps the instance on the same
hardware (in-place reboot). The `system-reboot` event remains
scheduled.

This setting only applies to supported instances that have a scheduled reboot event.
For more information, see [Enable or disable reboot migration](../../../../services/ec2/latest/userguide/schedevents-actions-reboot.md#reboot-migration) in the
_Amazon EC2 User Guide_.

Type: String

Valid Values: `disabled | default`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/instancemaintenanceoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/instancemaintenanceoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/instancemaintenanceoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InstanceIpv6Prefix

InstanceMaintenanceOptionsRequest
