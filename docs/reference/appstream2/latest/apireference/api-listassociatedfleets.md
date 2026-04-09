# ListAssociatedFleets

Retrieves the name of the fleet that is associated with the specified stack.

## Request Syntax

```nohighlight

{
   "NextToken": "string",
   "StackName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[NextToken](#API_ListAssociatedFleets_RequestSyntax)**

The pagination token to use to retrieve the next page of results for this operation. If this value is null, it retrieves the first page.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[StackName](#API_ListAssociatedFleets_RequestSyntax)**

The name of the stack.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

## Response Syntax

```nohighlight

{
   "Names": [ "string" ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Names](#API_ListAssociatedFleets_ResponseSyntax)**

The name of the fleet.

Type: Array of strings

Length Constraints: Minimum length of 1.

**[NextToken](#API_ListAssociatedFleets_ResponseSyntax)**

The pagination token to use to retrieve the next page of results for this operation. If there are no more pages, this value is null.

Type: String

Length Constraints: Minimum length of 1.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/listassociatedfleets.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/listassociatedfleets.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/listassociatedfleets.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/listassociatedfleets.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/listassociatedfleets.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/listassociatedfleets.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/listassociatedfleets.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/listassociatedfleets.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/listassociatedfleets.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/listassociatedfleets.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetExportImageTask

ListAssociatedStacks

All content copied from https://docs.aws.amazon.com/.
