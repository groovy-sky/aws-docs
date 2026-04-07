# ListConfigurationProfiles

Lists the configuration profiles for an application.

## Request Syntax

```nohighlight

GET /applications/ApplicationId/configurationprofiles?max_results=MaxResults&next_token=NextToken&type=Type HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[ApplicationId](#API_ListConfigurationProfiles_RequestSyntax)**

The application ID.

Pattern: `[a-z0-9]{4,7}`

Required: Yes

**[MaxResults](#API_ListConfigurationProfiles_RequestSyntax)**

The maximum number of items to return for this call. The call also returns a token that
you can specify in a subsequent call to get the next set of results.

Valid Range: Minimum value of 1. Maximum value of 50.

**[NextToken](#API_ListConfigurationProfiles_RequestSyntax)**

A token to start the list. Use this token to get the next set of results.

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[Type](#API_ListConfigurationProfiles_RequestSyntax)**

A filter based on the type of configurations that the configuration profile contains. A
configuration can be a feature flag or a freeform configuration.

Pattern: `^[a-zA-Z\.]+`

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "Items": [
      {
         "ApplicationId": "string",
         "Id": "string",
         "LocationUri": "string",
         "Name": "string",
         "Type": "string",
         "ValidatorTypes": [ "string" ]
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Items](#API_ListConfigurationProfiles_ResponseSyntax)**

The elements from this collection.

Type: Array of [ConfigurationProfileSummary](https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_ConfigurationProfileSummary.html) objects

**[NextToken](#API_ListConfigurationProfiles_ResponseSyntax)**

The token for the next set of items to return. Use this token to get the next set of
results.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The input fails to satisfy the constraints specified by an AWS service.

**Details**

Detailed information about the input that failed to satisfy the constraints specified by
a call.

HTTP Status Code: 400

**InternalServerException**

There was an internal failure in the AWS AppConfig service.

HTTP Status Code: 500

**ResourceNotFoundException**

The requested resource could not be found.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of ListConfigurationProfiles.

#### Sample Request

```

GET /applications/abc1234/configurationprofiles HTTP/1.1
Host: appconfig.us-east-1.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.2.4 Python/3.8.8 Linux/5.4.134-73.228.amzn2int.x86_64 exe/x86_64.amzn.2 prompt/off command/appconfig.list-configuration-profiles
X-Amz-Date: 20210920T174422Z
Authorization: AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20210920/us-east-1/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
```

#### Sample Response

```

{
    "Items": [
        {
            "ApplicationId": "abc1234",
            "Id": "ur8hx2f",
            "Name": "Example-Configuration-Profile",
            "LocationUri": "ssm-parameter://Example-Parameter"
        }
    ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/appconfig-2019-10-09/ListConfigurationProfiles)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/appconfig-2019-10-09/ListConfigurationProfiles)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appconfig-2019-10-09/ListConfigurationProfiles)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/appconfig-2019-10-09/ListConfigurationProfiles)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appconfig-2019-10-09/ListConfigurationProfiles)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/appconfig-2019-10-09/ListConfigurationProfiles)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/appconfig-2019-10-09/ListConfigurationProfiles)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/appconfig-2019-10-09/ListConfigurationProfiles)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/appconfig-2019-10-09/ListConfigurationProfiles)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appconfig-2019-10-09/ListConfigurationProfiles)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListApplications

ListDeployments
