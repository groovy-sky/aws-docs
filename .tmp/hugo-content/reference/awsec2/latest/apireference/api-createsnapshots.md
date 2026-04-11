# CreateSnapshots

Creates crash-consistent snapshots of multiple EBS volumes attached to an Amazon EC2 instance.
Volumes are chosen by specifying an instance. Each volume attached to the specified instance
will produce one snapshot that is crash-consistent across the instance. You can include all of
the volumes currently attached to the instance, or you can exclude the root volume or specific
data (non-root) volumes from the multi-volume snapshot set.

The location of the source instance determines where you can create the snapshots.

- If the source instance is in a Region, you must create the snapshots in the same
Region as the instance.

- If the source instance is in a Local Zone, you can create the snapshots in the same
Local Zone or in its parent AWS Region.

- If the source instance is on an Outpost, you can create the snapshots on the same
Outpost or in its parent AWS Region.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CopyTagsFromSource**

Copies the tags from the specified volume to corresponding snapshot.

Type: String

Valid Values: `volume`

Required: No

**Description**

A description propagated to every snapshot specified by the instance.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceSpecification**

The instance to specify which volumes should be included in the snapshots.

Type: [InstanceSpecification](api-instancespecification.md) object

Required: Yes

**Location**

###### Note

Only supported for instances in Local Zones. If the source instance is not in a Local Zone,
omit this parameter.

- To create local snapshots in the same Local Zone as the source instance, specify
`local`.

- To create regional snapshots in the parent Region of the Local Zone, specify
`regional` or omit this parameter.

Default value: `regional`

Type: String

Valid Values: `regional | local`

Required: No

**OutpostArn**

###### Note

Only supported for instances on Outposts. If the source instance is not on an Outpost,
omit this parameter.

- To create the snapshots on the same Outpost as the source instance, specify the
ARN of that Outpost. The snapshots must be created on the same Outpost as the instance.

- To create the snapshots in the parent Region of the Outpost, omit this parameter.

For more information, see [Create local snapshots from volumes on an Outpost](../../../../services/ebs/latest/userguide/snapshots-outposts.md#create-snapshot) in the _Amazon EBS User Guide_.

Type: String

Required: No

**TagSpecification.N**

Tags to apply to every snapshot specified by the instance.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**snapshotSet**

List of snapshots.

Type: Array of [SnapshotInfo](api-snapshotinfo.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createsnapshots.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createsnapshots.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createsnapshots.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createsnapshots.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createsnapshots.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createsnapshots.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createsnapshots.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createsnapshots.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createsnapshots.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createsnapshots.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateSnapshot

CreateSpotDatafeedSubscription
