# ListResourceScanRelatedResources

Lists the related resources for a list of resources from a resource scan. The response
indicates whether each returned resource is already managed by CloudFormation.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**MaxResults**

If the number of available results exceeds this maximum, the response includes a
`NextToken` value that you can use for the `NextToken` parameter to
get the next set of results. By default the `ListResourceScanRelatedResources` API
action will return up to 100 results in each response. The maximum value is 100.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**NextToken**

The token for the next set of items to return. (You received this token from a previous
call.)

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**Resources.member.N**

The list of resources for which you want to get the related resources. Up to 100 resources
can be provided.

Type: Array of [ScannedResourceIdentifier](api-scannedresourceidentifier.md) objects

Required: Yes

**ResourceScanId**

The Amazon Resource Name (ARN) of the resource scan.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**NextToken**

If the request doesn't return all the remaining results, `NextToken` is set to
a token. To retrieve the next set of results, call
`ListResourceScanRelatedResources` again and use that value for the
`NextToken` parameter. If the request returns all results, `NextToken`
is set to an empty string.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**RelatedResources.member.N**

List of up to `MaxResults` resources in the specified resource scan related to
the specified resources.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/listresourcescanrelatedresources.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/listresourcescanrelatedresources.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/listresourcescanrelatedresources.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/listresourcescanrelatedresources.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/listresourcescanrelatedresources.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/listresourcescanrelatedresources.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/listresourcescanrelatedresources.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/listresourcescanrelatedresources.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/listresourcescanrelatedresources.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/listresourcescanrelatedresources.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListImports

ListResourceScanResources

All content copied from https://docs.aws.amazon.com/.
