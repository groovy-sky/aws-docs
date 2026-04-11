# ModifyInstanceMaintenanceOptions

Modifies the recovery behavior of your instance to disable simplified automatic
recovery or set the recovery behavior to default. The default configuration will not
enable simplified automatic recovery for an unsupported instance type. For more
information, see [Simplified automatic recovery](../../../../services/ec2/latest/userguide/ec2-instance-recover.md#instance-configuration-recovery).

Modifies the reboot migration behavior during a user-initiated reboot of an instance
that has a pending `system-reboot` event. For more information, see [Enable or disable reboot migration](../../../../services/ec2/latest/userguide/schedevents-actions-reboot.md#reboot-migration).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AutoRecovery**

Disables the automatic recovery behavior of your instance or sets it to
default.

Type: String

Valid Values: `disabled | default`

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceId**

The ID of the instance.

Type: String

Required: Yes

**RebootMigration**

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

## Response Elements

The following elements are returned by the service.

**autoRecovery**

Provides information on the current automatic recovery behavior of your
instance.

Type: String

Valid Values: `disabled | default`

**instanceId**

The ID of the instance.

Type: String

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

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifyinstancemaintenanceoptions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifyinstancemaintenanceoptions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyinstancemaintenanceoptions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyinstancemaintenanceoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyinstancemaintenanceoptions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyinstancemaintenanceoptions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyinstancemaintenanceoptions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyinstancemaintenanceoptions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifyinstancemaintenanceoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyinstancemaintenanceoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyInstanceEventWindow

ModifyInstanceMetadataDefaults
