# DeleteEnvironment

Deletes an environment.

To prevent users from unintentionally deleting actively-used environments, enable [deletion\
protection](../../../../services/appconfig/latest/userguide/deletion-protection.md).

## Request Syntax

```nohighlight

DELETE /applications/ApplicationId/environments/EnvironmentId HTTP/1.1
x-amzn-deletion-protection-check: DeletionProtectionCheck

```

## URI Request Parameters

The request uses the following URI parameters.

**[ApplicationId](#API_DeleteEnvironment_RequestSyntax)**

The application ID that includes the environment that you want to delete.

Pattern: `[a-z0-9]{4,7}`

Required: Yes

**[DeletionProtectionCheck](#API_DeleteEnvironment_RequestSyntax)**

A parameter to configure deletion protection. Deletion protection prevents a user from
deleting an environment if your application called either [GetLatestConfiguration](api-appconfigdata-getlatestconfiguration.md) or [GetConfiguration](api-getconfiguration.md) in the
environment during the specified interval.

This parameter supports the following values:

- `BYPASS`: Instructs AWS AppConfig to bypass the deletion
protection check and delete a configuration profile even if deletion protection would
have otherwise prevented it.

- `APPLY`: Instructs the deletion protection check to run, even if
deletion protection is disabled at the account level. `APPLY` also forces
the deletion protection check to run against resources created in the past hour,
which are normally excluded from deletion protection checks.

- `ACCOUNT_DEFAULT`: The default setting, which instructs AWS AppConfig to implement the deletion protection value specified in the
`UpdateAccountSettings` API.

Valid Values: `ACCOUNT_DEFAULT | APPLY | BYPASS`

**[EnvironmentId](#API_DeleteEnvironment_RequestSyntax)**

The ID of the environment that you want to delete.

Pattern: `[a-z0-9]{4,7}`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 204

```

## Response Elements

If the action is successful, the service sends back an HTTP 204 response with an empty HTTP body.

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

This example illustrates one usage of DeleteEnvironment.

#### Sample Request

```

DELETE /applications/abc1234/environments/54j1r29 HTTP/1.1
Host: appconfig.us-east-1.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.2.4 Python/3.8.8 Linux/5.4.134-73.228.amzn2int.x86_64 exe/x86_64.amzn.2 prompt/off command/appconfig.delete-environment
X-Amz-Date: 20210920T220756Z
Authorization: AWS4-HMAC-SHA256 Credential=AWS_ACCESS_KEY_ID_REDACTED/20210920/us-east-1/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
Content-Length: 0
```

#### Sample Response

```

{}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/deleteenvironment.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/deleteenvironment.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/deleteenvironment.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/deleteenvironment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/deleteenvironment.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/deleteenvironment.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/deleteenvironment.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/deleteenvironment.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/deleteenvironment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/deleteenvironment.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteDeploymentStrategy

DeleteExtension

All content copied from https://docs.aws.amazon.com/.
