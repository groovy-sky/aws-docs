# ListTagsForResource

Displays the tags associated with a CloudWatch resource. Tags can be assigned to service level objectives.

## Request Syntax

```nohighlight

GET /tags?ResourceArn=ResourceArn HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[ResourceArn](#API_ListTagsForResource_RequestSyntax)**

The Amazon Resource Name (ARN) of the CloudWatch resource that you want to view tags for.

The ARN format of an Application Signals SLO is
`arn:aws:cloudwatch:Region:account-id:slo:slo-name
                  `

For more information about ARN format, see [Resource\
Types Defined by Amazon CloudWatch](https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazoncloudwatch.html#amazoncloudwatch-resources-for-iam-policies) in the _Amazon Web Services General_
_Reference_.

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
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

**[Tags](#API_ListTagsForResource_ResponseSyntax)**

The list of tag keys and values associated with the resource you specified.

Type: Array of [Tag](https://docs.aws.amazon.com/applicationsignals/latest/APIReference/API_Tag.html) objects

Array Members: Minimum number of 0 items. Maximum number of 200 items.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/applicationsignals/latest/APIReference/CommonErrors.html).

**ResourceNotFoundException**

Resource not found.

**ResourceId**

Can't find the resource id.

**ResourceType**

The resource type is not valid.

HTTP Status Code: 404

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 429

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/application-signals-2024-04-15/ListTagsForResource)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/application-signals-2024-04-15/ListTagsForResource)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/application-signals-2024-04-15/ListTagsForResource)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/application-signals-2024-04-15/ListTagsForResource)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/application-signals-2024-04-15/ListTagsForResource)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/application-signals-2024-04-15/ListTagsForResource)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/application-signals-2024-04-15/ListTagsForResource)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/application-signals-2024-04-15/ListTagsForResource)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/application-signals-2024-04-15/ListTagsForResource)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/application-signals-2024-04-15/ListTagsForResource)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListServiceStates

PutGroupingConfiguration
