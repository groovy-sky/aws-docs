# DisableSnapshotBlockPublicAccess

Disables the _block public access for snapshots_ setting at
the account level for the specified AWS Region. After you disable block public
access for snapshots in a Region, users can publicly share snapshots in that Region.

###### Important

Enabling block public access for snapshots in _block-all-sharing_
mode does not change the permissions for snapshots that are already publicly shared.
Instead, it prevents these snapshots from be publicly visible and publicly accessible.
Therefore, the attributes for these snapshots still indicate that they are publicly
shared, even though they are not publicly available.

If you disable block public access , these snapshots will become publicly available
again.

For more information, see [Block public access for snapshots](../../../../services/ebs/latest/userguide/block-public-access-snapshots.md) in the _Amazon EBS User Guide_ .

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

**requestId**

The ID of the request.

Type: String

**state**

Returns `unblocked` if the request succeeds.

Type: String

Valid Values: `block-all-sharing | block-new-sharing | unblocked`

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/disablesnapshotblockpublicaccess.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/disablesnapshotblockpublicaccess.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disablesnapshotblockpublicaccess.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/disablesnapshotblockpublicaccess.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disablesnapshotblockpublicaccess.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/disablesnapshotblockpublicaccess.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/disablesnapshotblockpublicaccess.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/disablesnapshotblockpublicaccess.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/disablesnapshotblockpublicaccess.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disablesnapshotblockpublicaccess.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisableSerialConsoleAccess

DisableTransitGatewayRouteTablePropagation
