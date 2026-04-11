# ListImagesInRecycleBin

Lists one or more AMIs that are currently in the Recycle Bin. For more information, see
[Recycle\
Bin](../../../../services/ec2/latest/userguide/recycle-bin.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**ImageId.N**

The IDs of the AMIs to list. Omit this parameter to list all of the AMIs that are in the
Recycle Bin. You can specify up to 20 IDs in a single request.

Type: Array of strings

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**imageSet**

Information about the AMIs.

Type: Array of [ImageRecycleBinInfo](api-imagerecyclebininfo.md) objects

**nextToken**

The token to include in another request to get the next page of items. This value is `null` when there
are no more items to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/listimagesinrecyclebin.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/listimagesinrecyclebin.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/listimagesinrecyclebin.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/listimagesinrecyclebin.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/listimagesinrecyclebin.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/listimagesinrecyclebin.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/listimagesinrecyclebin.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/listimagesinrecyclebin.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/listimagesinrecyclebin.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/listimagesinrecyclebin.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ImportVolume

ListSnapshotsInRecycleBin
