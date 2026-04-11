# DescribeRepositoryCreationTemplates

Returns details about the repository creation templates in a registry. The
`prefixes` request parameter can be used to return the details for a
specific repository creation template.

## Request Syntax

```nohighlight

{
   "maxResults": number,
   "nextToken": "string",
   "prefixes": [ "string" ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[maxResults](#API_DescribeRepositoryCreationTemplates_RequestSyntax)**

The maximum number of repository results returned by
`DescribeRepositoryCreationTemplatesRequest` in paginated output. When
this parameter is used, `DescribeRepositoryCreationTemplatesRequest` only
returns `maxResults` results in a single page along with a
`nextToken` response element. The remaining results of the initial
request can be seen by sending another
`DescribeRepositoryCreationTemplatesRequest` request with the returned
`nextToken` value. This value can be between 1 and
1000\. If this parameter is not used, then
`DescribeRepositoryCreationTemplatesRequest` returns up to
100 results and a `nextToken` value, if applicable.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**[nextToken](#API_DescribeRepositoryCreationTemplates_RequestSyntax)**

The `nextToken` value returned from a previous paginated
`DescribeRepositoryCreationTemplates` request where
`maxResults` was used and the results exceeded the value of that
parameter. Pagination continues from the end of the previous results that returned the
`nextToken` value. This value is `null` when there are no more
results to return.

###### Note

This token should be treated as an opaque identifier that is only used to
retrieve the next items in a list and not for other programmatic purposes.

Type: String

Required: No

**[prefixes](#API_DescribeRepositoryCreationTemplates_RequestSyntax)**

The repository namespace prefixes associated with the repository creation templates to
describe. If this value is not specified, all repository creation templates are
returned.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `^([a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*\/?|ROOT)$`

Required: No

## Response Syntax

```nohighlight

{
   "nextToken": "string",
   "registryId": "string",
   "repositoryCreationTemplates": [
      {
         "appliedFor": [ "string" ],
         "createdAt": number,
         "customRoleArn": "string",
         "description": "string",
         "encryptionConfiguration": {
            "encryptionType": "string",
            "kmsKey": "string"
         },
         "imageTagMutability": "string",
         "imageTagMutabilityExclusionFilters": [
            {
               "filter": "string",
               "filterType": "string"
            }
         ],
         "lifecyclePolicy": "string",
         "prefix": "string",
         "repositoryPolicy": "string",
         "resourceTags": [
            {
               "Key": "string",
               "Value": "string"
            }
         ],
         "updatedAt": number
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[nextToken](#API_DescribeRepositoryCreationTemplates_ResponseSyntax)**

The `nextToken` value to include in a future
`DescribeRepositoryCreationTemplates` request. When the results of a
`DescribeRepositoryCreationTemplates` request exceed
`maxResults`, this value can be used to retrieve the next page of
results. This value is `null` when there are no more results to
return.

Type: String

**[registryId](#API_DescribeRepositoryCreationTemplates_ResponseSyntax)**

The registry ID associated with the request.

Type: String

Pattern: `[0-9]{12}`

**[repositoryCreationTemplates](#API_DescribeRepositoryCreationTemplates_ResponseSyntax)**

The details of the repository creation templates.

Type: Array of [RepositoryCreationTemplate](api-repositorycreationtemplate.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

The specified parameter is invalid. Review the available parameters for the API
request.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**ServerException**

These errors are usually caused by a server-side issue.

**message**

The error message associated with the exception.

HTTP Status Code: 500

**ValidationException**

There was an exception validating this request.

HTTP Status Code: 400

## Examples

In the following example or examples, the Authorization header contents
( `AUTHPARAMS`) must be replaced with an AWS Signature Version 4
signature. For more information about creating these signatures, see [Signature\
Version 4 Signing Process](../../../../general/latest/gr/signature-version-4.md) in the _AWS General_
_Reference_.

You only need to learn how to sign HTTP requests if you intend to manually
create them. When you use the [AWS Command Line Interface (AWS CLI)](http://aws.amazon.com/cli) or one of the [AWS SDKs](http://aws.amazon.com/tools) to make requests to AWS, these tools automatically sign the
requests for you with the access key that you specify when you configure the tools.
When you use these tools, you don't need to learn how to sign requests
yourself.

### Example

This example describes a repository creation template in the default registry
for an account.

#### Sample Request

```

PPOST / HTTP/1.1
Host: ecr.us-west-2.amazonaws.com
Accept-Encoding: identity
Content-Length:240
X-Amz-Target: AmazonEC2ContainerRegistry_V20150921.DescribeRepositoryCreationTemplates
X-Amz-Date: 20231216T195356Z
User-Agent: aws-cli/1.11.22 Python/2.7.12 Darwin/16.3.0 botocore/1.4.79
Content-Type: application/x-amz-json-1.1
Authorization: AUTHPARAMS

{
 "prefixes":
  [
    "eng"
  ]
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Server: Server
Date: Sat, 16 Dec 2023 19:54:56 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 1590
Connection: keep-alive
x-amzn-RequestId: 60dc1ea1-c3c9-11e6-aa04-25c3a5fb1b54
{
    "registryId": "012345678901",
    "repositoryCreationTemplates": [{
        "prefix": "eng/test",
        "appliedFor": ["PULL_THROUGH_CACHE", "REPLICATION", "CREATE_ON_PUSH"],
        "encryptionConfiguration": {
            "encryptionType": "AES256"
        },
        "imageTagMutability": "MUTABLE",
        "createdAt": "2023-12-16T17:29:02-07:00",
        "updatedAt": "2023-12-16T19:55:02-07:00"
    }, {
        "prefix": "eng/replication-test",
        "appliedFor": ["REPLICATION"],
        "encryptionConfiguration": {
            "encryptionType": "AES256"
        },
        "imageTagMutability": "IMMUTABLE",
        "createdAt": "2023-12-14T17:29:02-07:00",
        "updatedAt": "2023-12-14T19:55:02-07:00"
    }]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/describerepositorycreationtemplates.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/describerepositorycreationtemplates.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/describerepositorycreationtemplates.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/describerepositorycreationtemplates.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/describerepositorycreationtemplates.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/describerepositorycreationtemplates.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/describerepositorycreationtemplates.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/describerepositorycreationtemplates.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/describerepositorycreationtemplates.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/describerepositorycreationtemplates.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeRepositories

GetAccountSetting

All content copied from https://docs.aws.amazon.com/.
