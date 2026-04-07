# ListTagsForResource

Lists the tags associated with an Athena resource.

## Request Syntax

```nohighlight

{
   "MaxResults": number,
   "NextToken": "string",
   "ResourceARN": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[MaxResults](#API_ListTagsForResource_RequestSyntax)**

The maximum number of results to be returned per request that lists the tags for the
resource.

Type: Integer

Valid Range: Minimum value of 75.

Required: No

**[NextToken](#API_ListTagsForResource_RequestSyntax)**

The token for the next set of results, or null if there are no additional results for
this request, where the request lists the tags for the resource with the specified
ARN.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**[ResourceARN](#API_ListTagsForResource_RequestSyntax)**

Lists the tags for the resource with the specified ARN.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Required: Yes

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_ListTagsForResource_ResponseSyntax)**

A token to be used by the next request if this request is truncated.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**[Tags](#API_ListTagsForResource_ResponseSyntax)**

The list of tags associated with the specified resource.

Type: Array of [Tag](api-tag.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

Indicates a platform issue, which may be due to a transient condition or
outage.

HTTP Status Code: 500

**InvalidRequestException**

Indicates that something is wrong with the input to the request. For example, a
required parameter may be missing or out of range.

**AthenaErrorCode**

The error code returned when the query execution failed to process, or when the
processing request for the named query failed.

HTTP Status Code: 400

**ResourceNotFoundException**

A resource, such as a workgroup, was not found.

**ResourceName**

The name of the Amazon resource.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/athena-2017-05-18/ListTagsForResource)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/athena-2017-05-18/ListTagsForResource)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/ListTagsForResource)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/athena-2017-05-18/ListTagsForResource)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/ListTagsForResource)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/athena-2017-05-18/ListTagsForResource)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/athena-2017-05-18/ListTagsForResource)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/athena-2017-05-18/ListTagsForResource)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/athena-2017-05-18/ListTagsForResource)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/ListTagsForResource)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListTableMetadata

ListWorkGroups
