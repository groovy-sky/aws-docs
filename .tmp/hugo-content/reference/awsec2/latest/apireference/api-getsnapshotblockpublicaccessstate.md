# GetSnapshotBlockPublicAccessState

Gets the current state of _block public access for snapshots_ setting
for the account and Region.

For more information, see [Block public access for snapshots](../../../../services/ebs/latest/userguide/block-public-access-snapshots.md) in the _Amazon EBS User Guide_.

## Request Parameters

For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**managedBy**

The entity that manages the state for block public access for snapshots. Possible
values include:

- `account` \- The state is managed by the account.

- `declarative-policy` \- The state is managed by a declarative policy and
can't be modified by the account.

Type: String

Valid Values: `account | declarative-policy`

**requestId**

The ID of the request.

Type: String

**state**

The current state of block public access for snapshots. Possible values include:

- `block-all-sharing` \- All public sharing of snapshots is blocked. Users in
the account can't request new public sharing. Additionally, snapshots that were already
publicly shared are treated as private and are not publicly available.

- `block-new-sharing` \- Only new public sharing of snapshots is blocked.
Users in the account can't request new public sharing. However, snapshots that were
already publicly shared, remain publicly available.

- `unblocked` \- Public sharing is not blocked. Users can publicly share
snapshots.

Type: String

Valid Values: `block-all-sharing | block-new-sharing | unblocked`

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/getsnapshotblockpublicaccessstate.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/getsnapshotblockpublicaccessstate.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getsnapshotblockpublicaccessstate.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getsnapshotblockpublicaccessstate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getsnapshotblockpublicaccessstate.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getsnapshotblockpublicaccessstate.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getsnapshotblockpublicaccessstate.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getsnapshotblockpublicaccessstate.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/getsnapshotblockpublicaccessstate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getsnapshotblockpublicaccessstate.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetSerialConsoleAccessStatus

GetSpotPlacementScores
