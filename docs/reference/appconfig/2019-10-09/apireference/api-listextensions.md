# ListExtensions

Lists all custom and AWS authored AWS AppConfig extensions in the
account. For more information about extensions, see [Extending\
workflows](../../../../services/appconfig/latest/userguide/working-with-appconfig-extensions.md) in the _AWS AppConfig User Guide_.

## Request Syntax

```nohighlight

GET /extensions?max_results=MaxResults&name=Name&next_token=NextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[MaxResults](#API_ListExtensions_RequestSyntax)**

The maximum number of items to return for this call. The call also returns a token that
you can specify in a subsequent call to get the next set of results.

Valid Range: Minimum value of 1. Maximum value of 50.

**[Name](#API_ListExtensions_RequestSyntax)**

The extension name.

Length Constraints: Minimum length of 1. Maximum length of 64.

**[NextToken](#API_ListExtensions_RequestSyntax)**

A token to start the list. Use this token to get the next set of results.

Length Constraints: Minimum length of 1. Maximum length of 2048.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "Items": [
      {
         "Arn": "string",
         "Description": "string",
         "Id": "string",
         "Name": "string",
         "VersionNumber": number
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Items](#API_ListExtensions_ResponseSyntax)**

The list of available extensions. The list includes AWS authored and
user-created extensions.

Type: Array of [ExtensionSummary](api-extensionsummary.md) objects

**[NextToken](#API_ListExtensions_ResponseSyntax)**

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

## Examples

### Example

This example illustrates one usage of ListExtensions.

#### Sample Request

```

GET /extensions HTTP/1.1
Host: appconfig.us-west-2.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.7.19 Python/3.9.11 Windows/10 exe/AMD64 prompt/off command/appconfig.list-extensions
X-Amz-Date: 20220803T215314Z
Authorization: AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20220803/us-west-2/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
```

#### Sample Response

```

{
	"Items": [{
		"Arn": "arn:aws:appconfig:us-west-2:111122223333:extension/6czExample/1",
		"Description": null,
		"Id": "6czExampmle",
		"Name": "my-test-extension",
		"VersionNumber": 1
	}, {
		"Arn": "arn:aws:appconfig:us-west-2::extension/AWS.AppConfig.FeatureFlags/1",
		"Description": "Validates AppConfig feature flag data automatically against a JSON schema that includes structure and constraints. Also transforms feature flag data prior to sending to the client. This extension is automatically associated to configuration profiles with type \"AWS.AppConfig.FeatureFlags\".",
		"Id": "AWS.AppConfig.FeatureFlags",
		"Name": "AppConfig Feature Flags Helper",
		"VersionNumber": 1
	}, {
		"Arn": "arn:aws:appconfig:us-west-2::extension/AWS.AppConfig.JiraIntegration/1",
		"Description": "Exports feature flag data from AWS AppConfig into Jira. The lifecycle of each feature flag in AppConfig is tracked in Jira as an individual issue. Customers can see in Jira when flags are updated, turned on or off. Works in conjunction with the AppConfig app in the Atlassian Marketplace and is automatically associated to configuration profiles configured within that app.",
		"Id": "AWS.AppConfig.JiraIntegration",
		"Name": "AppConfig integration with Atlassian Jira",
		"VersionNumber": 1
	}, {
		"Arn": "arn:aws:appconfig:us-west-2::extension/AWS.AppConfig.DeploymentNotificationsToEventBridge/1",
		"Description": "Sends events to Amazon EventBridge when a deployment of configuration data in AppConfig is started, completed, or rolled back. Can be associated to the following resources in AppConfig: Application, Environment, Configuration Profile.",
		"Id": "AWS.AppConfig.DeploymentNotificationsToEventBridge",
		"Name": "AppConfig deployment events to Amazon EventBridge",
		"VersionNumber": 1
	}, {
		"Arn": "arn:aws:appconfig:us-west-2::extension/AWS.AppConfig.DeploymentNotificationsToSqs/1",
		"Description": "Sends messages to the configured Amazon SQS queue when a deployment of configuration data in AppConfig is started, completed, or rolled back. Can be associated to the following resources in AppConfig: Application, Environment, Configuration Profile.",
		"Id": "AWS.AppConfig.DeploymentNotificationsToSqs",
		"Name": "AppConfig deployment events to Amazon SQS",
		"VersionNumber": 1
	}, {
		"Arn": null,
		"Description": "Sends events to the configured Amazon SNS topic when a deployment of configuration data in AppConfig is started, completed, or rolled back. Can be associated to the following resources in AppConfig: Application, Environment, Configuration Profile.",
		"Id": "AWS.AppConfig.DeploymentNotificationsToSns",
		"Name": "AppConfig deployment events to Amazon SNS",
		"VersionNumber": 1
	}],
	"NextToken": null
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/listextensions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/listextensions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/listextensions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/listextensions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/listextensions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/listextensions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/listextensions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/listextensions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/listextensions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/listextensions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListExtensionAssociations

ListHostedConfigurationVersions
