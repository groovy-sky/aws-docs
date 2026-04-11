# ListResourceScanResources

Lists the resources from a resource scan. The results can be filtered by resource
identifier, resource type prefix, tag key, and tag value. Only resources that match all
specified filters are returned. The response indicates whether each returned resource is
already managed by CloudFormation.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**MaxResults**

If the number of available results exceeds this maximum, the response includes a
`NextToken` value that you can use for the `NextToken` parameter to
get the next set of results. By default the `ListResourceScanResources` API action
will return at most 100 results in each response. The maximum value is 100.

Type: Integer

Required: No

**NextToken**

The token for the next set of items to return. (You received this token from a previous
call.)

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**ResourceIdentifier**

If specified, the returned resources will have the specified resource identifier (or one
of them in the case where the resource has multiple identifiers).

Type: String

Required: No

**ResourceScanId**

The Amazon Resource Name (ARN) of the resource scan.

Type: String

Required: Yes

**ResourceTypePrefix**

If specified, the returned resources will be of any of the resource types with the
specified prefix.

Type: String

Required: No

**TagKey**

If specified, the returned resources will have a matching tag key.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**TagValue**

If specified, the returned resources will have a matching tag value.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: No

## Response Elements

The following elements are returned by the service.

**NextToken**

If the request doesn't return all the remaining results, `NextToken` is set to
a token. To retrieve the next set of results, call `ListResourceScanResources`
again and use that value for the `NextToken` parameter. If the request returns all
results, `NextToken` is set to an empty string.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**Resources.member.N**

List of up to `MaxResults` resources in the specified resource scan that match
all of the specified filters.

Type: Array of [ScannedResource](api-scannedresource.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ResourceScanInProgress**

A resource scan is currently in progress. Only one can be run at a time for an account in a Region.

HTTP Status Code: 400

**ResourceScanNotFound**

The resource scan was not found.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/listresourcescanresources.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/listresourcescanresources.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/listresourcescanresources.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/listresourcescanresources.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/listresourcescanresources.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/listresourcescanresources.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/listresourcescanresources.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/listresourcescanresources.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/listresourcescanresources.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/listresourcescanresources.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListResourceScanRelatedResources

ListResourceScans

All content copied from https://docs.aws.amazon.com/.
