# PendingMaintenanceAction

Provides information about a pending maintenance action for a resource.

## Contents

###### Note

In the following list, the required parameters are described first.

**Action**

The type of pending maintenance action that is available for the resource.

For more information about maintenance actions, see [Maintaining a DB instance](../../../../services/amazonrds/latest/userguide/user-upgradedbinstance-maintenance.md).

Valid Values:

- `ca-certificate-rotation`

- `db-upgrade`

- `hardware-maintenance`

- `os-upgrade`

- `system-update`

For more information about these actions, see
[Maintenance actions for Amazon Aurora](../../../../services/amazonrds/latest/aurorauserguide/user-upgradedbinstance-maintenance.md#maintenance-actions-aurora) or
[Maintenance actions for Amazon RDS](../../../../services/amazonrds/latest/userguide/user-upgradedbinstance-maintenance.md#maintenance-actions-rds).

Type: String

Required: No

**AutoAppliedAfterDate**

The date of the maintenance window when the action is applied.
The maintenance action is applied to the resource during
its first maintenance window after this date.

Type: Timestamp

Required: No

**CurrentApplyDate**

The effective date when the pending maintenance action is applied
to the resource. This date takes into account opt-in requests received from
the `ApplyPendingMaintenanceAction` API, the `AutoAppliedAfterDate`,
and the `ForcedApplyDate`. This value is blank if an
opt-in request has not been received and nothing has been specified as
`AutoAppliedAfterDate` or `ForcedApplyDate`.

Type: Timestamp

Required: No

**Description**

A description providing more detail about the maintenance action.

Type: String

Required: No

**ForcedApplyDate**

The date when the maintenance action is automatically applied.

On this date, the maintenance action is applied to the resource as soon as possible,
regardless of the maintenance window for the resource. There might be a delay of
one or more days from this date before the maintenance action is applied.

Type: Timestamp

Required: No

**OptInStatus**

Indicates the type of opt-in request that has been received for the resource.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/pendingmaintenanceaction.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/pendingmaintenanceaction.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/pendingmaintenanceaction.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PendingCloudwatchLogsExports

PendingModifiedValues

All content copied from https://docs.aws.amazon.com/.
