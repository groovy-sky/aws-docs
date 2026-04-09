# CreateEnvironment

Creates an environment. For each application, you define one or more environments. An
environment is a deployment group of AWS AppConfig targets, such as applications in a
`Beta` or `Production` environment. You can also define
environments for application subcomponents such as the `Web`,
`Mobile` and `Back-end` components for your application. You can
configure Amazon CloudWatch alarms for each environment. The system monitors alarms during a
configuration deployment. If an alarm is triggered, the system rolls back the
configuration.

## Request Syntax

```nohighlight

POST /applications/ApplicationId/environments HTTP/1.1
Content-type: application/json

{
   "Description": "string",
   "Monitors": [
      {
         "AlarmArn": "string",
         "AlarmRoleArn": "string"
      }
   ],
   "Name": "string",
   "Tags": {
      "string" : "string"
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[ApplicationId](#API_CreateEnvironment_RequestSyntax)**

The application ID.

Pattern: `[a-z0-9]{4,7}`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[Description](#API_CreateEnvironment_RequestSyntax)**

A description of the environment.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**[Monitors](#API_CreateEnvironment_RequestSyntax)**

Amazon CloudWatch alarms to monitor during the deployment process.

Type: Array of [Monitor](api-monitor.md) objects

Array Members: Minimum number of 0 items. Maximum number of 5 items.

Required: No

**[Name](#API_CreateEnvironment_RequestSyntax)**

A name for the environment.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[Tags](#API_CreateEnvironment_RequestSyntax)**

Metadata to assign to the environment. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which
you define.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Value Length Constraints: Maximum length of 256.

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 201
Content-type: application/json

{
   "ApplicationId": "string",
   "Description": "string",
   "Id": "string",
   "Monitors": [
      {
         "AlarmArn": "string",
         "AlarmRoleArn": "string"
      }
   ],
   "Name": "string",
   "State": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The following data is returned in JSON format by the service.

**[ApplicationId](#API_CreateEnvironment_ResponseSyntax)**

The application ID.

Type: String

Pattern: `[a-z0-9]{4,7}`

**[Description](#API_CreateEnvironment_ResponseSyntax)**

The description of the environment.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

**[Id](#API_CreateEnvironment_ResponseSyntax)**

The environment ID.

Type: String

Pattern: `[a-z0-9]{4,7}`

**[Monitors](#API_CreateEnvironment_ResponseSyntax)**

Amazon CloudWatch alarms monitored during the deployment.

Type: Array of [Monitor](api-monitor.md) objects

Array Members: Minimum number of 0 items. Maximum number of 5 items.

**[Name](#API_CreateEnvironment_ResponseSyntax)**

The name of the environment.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

**[State](#API_CreateEnvironment_ResponseSyntax)**

The state of the environment. An environment can be in one of the following states:
`READY_FOR_DEPLOYMENT`, `DEPLOYING`, `ROLLING_BACK`, or
`ROLLED_BACK`

Type: String

Valid Values: `READY_FOR_DEPLOYMENT | DEPLOYING | ROLLING_BACK | ROLLED_BACK | REVERTED`

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

**ServiceQuotaExceededException**

The number of one more AWS AppConfig resources exceeds the maximum allowed. Verify that your
environment doesn't exceed the following service quotas:

Applications: 100 max

Deployment strategies: 20 max

Configuration profiles: 100 max per application

Environments: 20 max per application

To resolve this issue, you can delete one or more resources and try again. Or, you can
request a quota increase. For more information about quotas and to request an increase, see
[Service quotas for AWS AppConfig](../../../../general/latest/gr/appconfig.md#limits_appconfig) in the Amazon Web Services General Reference.

HTTP Status Code: 402

## Examples

### Example

This example illustrates one usage of CreateEnvironment.

#### Sample Request

```

POST /applications/abc1234/environments HTTP/1.1
Host: appconfig.us-east-1.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.2.4 Python/3.8.8 Linux/5.4.134-73.228.amzn2int.x86_64 exe/x86_64.amzn.2 prompt/off command/appconfig.create-environment
X-Amz-Date: 20210916T221023Z
Authorization: AWS4-HMAC-SHA256 Credential=AWS_ACCESS_KEY_ID_REDACTED/20210916/us-east-1/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
Content-Length: 31

{
	"Name": "Example-Environment"
}
```

#### Sample Response

```

{
	"ApplicationId": "abc1234",
	"Description": null,
	"Id": "54j1r29",
	"Monitors": null,
	"Name": "Example-Environment",
	"State": "ReadyForDeployment"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/createenvironment.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/createenvironment.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/createenvironment.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/createenvironment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/createenvironment.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/createenvironment.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/createenvironment.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/createenvironment.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/createenvironment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/createenvironment.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateDeploymentStrategy

CreateExtension

All content copied from https://docs.aws.amazon.com/.
