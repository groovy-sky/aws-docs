---
title: "ListTagsOfResource"
---

# ListTagsOfResource

List all tags on an Amazon DynamoDB resource. You can call ListTagsOfResource up to 10
times per second, per account.

For an overview on tagging DynamoDB resources, see [Tagging for DynamoDB](../../../../services/dynamodb/latest/developerguide/tagging.md)
in the _Amazon DynamoDB Developer Guide_.

## Request Syntax

```nohighlight

{
   "NextToken": "string",
   "ResourceArn": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ResourceArn](#API_ListTagsOfResource_RequestSyntax)**

The Amazon DynamoDB resource with tags to be listed. This value is an Amazon Resource
Name (ARN).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: Yes

**[NextToken](#API_ListTagsOfResource_RequestSyntax)**

An optional string that, if supplied, must be copied from the output of a previous
call to ListTagOfResource. When provided in this manner, this API fetches the next page
of results.

Type: String

Required: No

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

**[NextToken](#API_ListTagsOfResource_ResponseSyntax)**

If this value is returned, there are additional results to be displayed. To retrieve
them, call ListTagsOfResource again, with NextToken set to this value.

Type: String

**[Tags](#API_ListTagsOfResource_ResponseSyntax)**

The tags currently associated with the Amazon DynamoDB resource.

Type: Array of [Tag](api-tag.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerError**

An error occurred on the server side.

**message**

The server encountered an internal error trying to fulfill the request.

HTTP Status Code: 500

**ResourceNotFoundException**

The operation tried to access a nonexistent table or index. The resource might not
be specified correctly, or its status might not be `ACTIVE`.

**message**

The resource which is being requested does not exist.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/dynamodb-2012-08-10/ListTagsOfResource)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/dynamodb-2012-08-10/ListTagsOfResource)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/ListTagsOfResource)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/dynamodb-2012-08-10/ListTagsOfResource)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/ListTagsOfResource)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/dynamodb-2012-08-10/ListTagsOfResource)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/dynamodb-2012-08-10/ListTagsOfResource)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/dynamodb-2012-08-10/ListTagsOfResource)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/dynamodb-2012-08-10/ListTagsOfResource)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/ListTagsOfResource)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListTables

PutItem

All content copied from https://docs.aws.amazon.com/.
