# StartDeployment

Starts a deployment.

## Request Syntax

```nohighlight

POST /applications/ApplicationId/environments/EnvironmentId/deployments HTTP/1.1
Content-type: application/json

{
   "ConfigurationProfileId": "string",
   "ConfigurationVersion": "string",
   "DeploymentStrategyId": "string",
   "Description": "string",
   "DynamicExtensionParameters": {
      "string" : "string"
   },
   "KmsKeyIdentifier": "string",
   "Tags": {
      "string" : "string"
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[ApplicationId](#API_StartDeployment_RequestSyntax)**

The application ID.

Pattern: `[a-z0-9]{4,7}`

Required: Yes

**[EnvironmentId](#API_StartDeployment_RequestSyntax)**

The environment ID.

Pattern: `[a-z0-9]{4,7}`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[ConfigurationProfileId](#API_StartDeployment_RequestSyntax)**

The configuration profile ID.

Type: String

Pattern: `[a-z0-9]{4,7}`

Required: Yes

**[ConfigurationVersion](#API_StartDeployment_RequestSyntax)**

The configuration version to deploy. If deploying an AWS AppConfig hosted
configuration version, you can specify either the version number or version label. For all
other configurations, you must specify the version number.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**[DeploymentStrategyId](#API_StartDeployment_RequestSyntax)**

The deployment strategy ID.

Type: String

Pattern: `(^[a-z0-9]{4,7}$|^AppConfig\.[A-Za-z0-9]{9,40}$)`

Required: Yes

**[Description](#API_StartDeployment_RequestSyntax)**

A description of the deployment.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**[DynamicExtensionParameters](#API_StartDeployment_RequestSyntax)**

A map of dynamic extension parameter names to values to pass to associated extensions
with `PRE_START_DEPLOYMENT` actions.

Type: String to string map

Map Entries: Maximum number of 10 items.

Key Pattern: `^([^#\n]{1,96})#([^\/#\n]{1,64})$`

Value Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**[KmsKeyIdentifier](#API_StartDeployment_RequestSyntax)**

The AWS KMS key identifier (key ID, key alias, or key ARN). AWS AppConfig uses this ID to encrypt the configuration data using a customer managed key.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**[Tags](#API_StartDeployment_RequestSyntax)**

Metadata to assign to the deployment. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which
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
   "AppliedExtensions": [
      {
         "ExtensionAssociationId": "string",
         "ExtensionId": "string",
         "Parameters": {
            "string" : "string"
         },
         "VersionNumber": number
      }
   ],
   "CompletedAt": "string",
   "ConfigurationLocationUri": "string",
   "ConfigurationName": "string",
   "ConfigurationProfileId": "string",
   "ConfigurationVersion": "string",
   "DeploymentDurationInMinutes": number,
   "DeploymentNumber": number,
   "DeploymentStrategyId": "string",
   "Description": "string",
   "EnvironmentId": "string",
   "EventLog": [
      {
         "ActionInvocations": [
            {
               "ActionName": "string",
               "ErrorCode": "string",
               "ErrorMessage": "string",
               "ExtensionIdentifier": "string",
               "InvocationId": "string",
               "RoleArn": "string",
               "Uri": "string"
            }
         ],
         "Description": "string",
         "EventType": "string",
         "OccurredAt": "string",
         "TriggeredBy": "string"
      }
   ],
   "FinalBakeTimeInMinutes": number,
   "GrowthFactor": number,
   "GrowthType": "string",
   "KmsKeyArn": "string",
   "KmsKeyIdentifier": "string",
   "PercentageComplete": number,
   "StartedAt": "string",
   "State": "string",
   "VersionLabel": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The following data is returned in JSON format by the service.

**[ApplicationId](#API_StartDeployment_ResponseSyntax)**

The ID of the application that was deployed.

Type: String

Pattern: `[a-z0-9]{4,7}`

**[AppliedExtensions](#API_StartDeployment_ResponseSyntax)**

A list of extensions that were processed as part of the deployment. The extensions that
were previously associated to the configuration profile, environment, or the application
when `StartDeployment` was called.

Type: Array of [AppliedExtension](api-appliedextension.md) objects

**[CompletedAt](#API_StartDeployment_ResponseSyntax)**

The time the deployment completed.

Type: Timestamp

**[ConfigurationLocationUri](#API_StartDeployment_ResponseSyntax)**

Information about the source location of the configuration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[ConfigurationName](#API_StartDeployment_ResponseSyntax)**

The name of the configuration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

**[ConfigurationProfileId](#API_StartDeployment_ResponseSyntax)**

The ID of the configuration profile that was deployed.

Type: String

Pattern: `[a-z0-9]{4,7}`

**[ConfigurationVersion](#API_StartDeployment_ResponseSyntax)**

The configuration version that was deployed.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**[DeploymentDurationInMinutes](#API_StartDeployment_ResponseSyntax)**

Total amount of time the deployment lasted.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 1440.

**[DeploymentNumber](#API_StartDeployment_ResponseSyntax)**

The sequence number of the deployment.

Type: Integer

**[DeploymentStrategyId](#API_StartDeployment_ResponseSyntax)**

The ID of the deployment strategy that was deployed.

Type: String

Pattern: `[a-z0-9]{4,7}`

**[Description](#API_StartDeployment_ResponseSyntax)**

The description of the deployment.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

**[EnvironmentId](#API_StartDeployment_ResponseSyntax)**

The ID of the environment that was deployed.

Type: String

Pattern: `[a-z0-9]{4,7}`

**[EventLog](#API_StartDeployment_ResponseSyntax)**

A list containing all events related to a deployment. The most recent events are
displayed first.

Type: Array of [DeploymentEvent](api-deploymentevent.md) objects

**[FinalBakeTimeInMinutes](#API_StartDeployment_ResponseSyntax)**

The amount of time that AWS AppConfig monitored for alarms before considering the
deployment to be complete and no longer eligible for automatic rollback.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 1440.

**[GrowthFactor](#API_StartDeployment_ResponseSyntax)**

The percentage of targets to receive a deployed configuration during each
interval.

Type: Float

Valid Range: Minimum value of 1.0. Maximum value of 100.0.

**[GrowthType](#API_StartDeployment_ResponseSyntax)**

The algorithm used to define how percentage grew over time.

Type: String

Valid Values: `LINEAR | EXPONENTIAL`

**[KmsKeyArn](#API_StartDeployment_ResponseSyntax)**

The Amazon Resource Name of the AWS Key Management Service key used to encrypt configuration
data. You can encrypt secrets stored in AWS Secrets Manager, Amazon Simple Storage Service
(Amazon S3) objects encrypted with SSE-KMS, or secure string parameters stored in AWS Systems Manager
Parameter Store.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

**[KmsKeyIdentifier](#API_StartDeployment_ResponseSyntax)**

The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when
the resource was created or updated.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[PercentageComplete](#API_StartDeployment_ResponseSyntax)**

The percentage of targets for which the deployment is available.

Type: Float

Valid Range: Minimum value of 1.0. Maximum value of 100.0.

**[StartedAt](#API_StartDeployment_ResponseSyntax)**

The time the deployment started.

Type: Timestamp

**[State](#API_StartDeployment_ResponseSyntax)**

The state of the deployment.

Type: String

Valid Values: `BAKING | VALIDATING | DEPLOYING | COMPLETE | ROLLING_BACK | ROLLED_BACK | REVERTED`

**[VersionLabel](#API_StartDeployment_ResponseSyntax)**

A user-defined label for an AWS AppConfig hosted configuration version.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `.*[^0-9].*`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The input fails to satisfy the constraints specified by an AWS service.

**Details**

Detailed information about the input that failed to satisfy the constraints specified by
a call.

HTTP Status Code: 400

**ConflictException**

The request could not be processed because of conflict in the current state of the
resource.

HTTP Status Code: 409

**InternalServerException**

There was an internal failure in the AWS AppConfig service.

HTTP Status Code: 500

**ResourceNotFoundException**

The requested resource could not be found.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of StartDeployment.

#### Sample Request

```

POST /applications/abc1234/environments/54j1r29/deployments HTTP/1.1
Host: appconfig.us-east-1.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.2.4 Python/3.8.8 Linux/5.4.134-73.228.amzn2int.x86_64 exe/x86_64.amzn.2 prompt/off command/appconfig.start-deployment
X-Amz-Date: 20210917T214353Z
Authorization: AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20210917/us-east-1/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
Content-Length: 101

{
    "DeploymentStrategyId": "1225qzk",
    "ConfigurationProfileId": "ur8hx2f",
    "ConfigurationVersion": "1"
}
```

#### Sample Response

```

{
    "ApplicationId": "abc1234",
    "EnvironmentId": "54j1r29",
    "DeploymentStrategyId": "1225qzk",
    "ConfigurationProfileId": "ur8hx2f",
    "DeploymentNumber": 1,
    "ConfigurationName": "Example-Configuration-Profile",
    "ConfigurationLocationUri": "ssm-parameter://Example-Parameter",
    "ConfigurationVersion": "1",
    "DeploymentDurationInMinutes": 15,
    "GrowthType": "LINEAR",
    "GrowthFactor": 25.0,
    "FinalBakeTimeInMinutes": 0,
    "State": "DEPLOYING",
    "EventLog": [
        {
            "EventType": "DEPLOYMENT_STARTED",
            "TriggeredBy": "USER",
            "Description": "Deployment started",
            "OccurredAt": "2021-09-17T21:43:54.205000+00:00"
        }
    ],
    "PercentageComplete": 0.0,
    "StartedAt": "2021-09-17T21:43:54.205000+00:00"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/startdeployment.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/startdeployment.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/startdeployment.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/startdeployment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/startdeployment.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/startdeployment.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/startdeployment.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/startdeployment.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/startdeployment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/startdeployment.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListTagsForResource

StopDeployment
