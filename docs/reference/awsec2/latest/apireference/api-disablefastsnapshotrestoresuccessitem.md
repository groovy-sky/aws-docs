# DisableFastSnapshotRestoreSuccessItem

Describes fast snapshot restores that were successfully disabled.

## Contents

**availabilityZone**

The Availability Zone.

Type: String

Required: No

**availabilityZoneId**

The ID of the Availability Zone.

Type: String

Required: No

**disabledTime**

The time at which fast snapshot restores entered the `disabled` state.

Type: Timestamp

Required: No

**disablingTime**

The time at which fast snapshot restores entered the `disabling` state.

Type: Timestamp

Required: No

**enabledTime**

The time at which fast snapshot restores entered the `enabled` state.

Type: Timestamp

Required: No

**enablingTime**

The time at which fast snapshot restores entered the `enabling` state.

Type: Timestamp

Required: No

**optimizingTime**

The time at which fast snapshot restores entered the `optimizing` state.

Type: Timestamp

Required: No

**ownerAlias**

The AWS owner alias that enabled fast snapshot restores on the snapshot. This is intended for future use.

Type: String

Required: No

**ownerId**

The ID of the AWS account that enabled fast snapshot restores on the snapshot.

Type: String

Required: No

**snapshotId**

The ID of the snapshot.

Type: String

Required: No

**state**

The state of fast snapshot restores for the snapshot.

Type: String

Valid Values: `enabling | optimizing | enabled | disabling | disabled`

Required: No

**stateTransitionReason**

The reason for the state transition. The possible values are as follows:

- `Client.UserInitiated` \- The state successfully transitioned to `enabling` or
`disabling`.

- `Client.UserInitiated - Lifecycle state transition` \- The state successfully transitioned
to `optimizing`, `enabled`, or `disabled`.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disablefastsnapshotrestoresuccessitem.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disablefastsnapshotrestoresuccessitem.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disablefastsnapshotrestoresuccessitem.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisableFastSnapshotRestoreStateErrorItem

DiskImage
