# GetConfiguration

(Deprecated) Retrieves the latest deployed configuration.

###### Important

Note the following important information.

- This API action is deprecated. Calls to receive configuration data should use
the [StartConfigurationSession](api-appconfigdata-startconfigurationsession.md) and [GetLatestConfiguration](api-appconfigdata-getlatestconfiguration.md) APIs instead.

- GetConfiguration is a priced call. For more information, see
[Pricing](https://aws.amazon.com/systems-manager/pricing).

## Request Syntax

```nohighlight

GET /applications/Application/environments/Environment/configurations/Configuration?client_configuration_version=ClientConfigurationVersion&client_id=ClientId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Application](#API_GetConfiguration_RequestSyntax)**

The application to get. Specify either the application name or the application
ID.

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[ClientConfigurationVersion](#API_GetConfiguration_RequestSyntax)**

The configuration version returned in the most recent [GetConfiguration](api-getconfiguration.md)
response.

###### Important

AWS AppConfig uses the value of the `ClientConfigurationVersion`
parameter to identify the configuration version on your clients. If you don’t send
`ClientConfigurationVersion` with each call to [GetConfiguration](api-getconfiguration.md), your clients receive the current configuration. You are
charged each time your clients receive a configuration.

To avoid excess charges, we recommend you use the [StartConfigurationSession](startconfigurationsession.md) and [GetLatestConfiguration](getlatestconfiguration.md) APIs, which track the client configuration version on
your behalf. If you choose to continue using [GetConfiguration](api-getconfiguration.md), we
recommend that you include the `ClientConfigurationVersion` value with every
call to [GetConfiguration](api-getconfiguration.md). The value to use for
`ClientConfigurationVersion` comes from the
`ConfigurationVersion` attribute returned by [GetConfiguration](api-getconfiguration.md) when there is new or updated data, and should be saved
for subsequent calls to [GetConfiguration](api-getconfiguration.md).

For more information about working with configurations, see [Retrieving feature flags and\
configuration data in AWS AppConfig](../../../../services/appconfig/latest/userguide/retrieving-feature-flags.md) in the _AWS AppConfig_
_User Guide_.

Length Constraints: Minimum length of 1. Maximum length of 1024.

**[ClientId](#API_GetConfiguration_RequestSyntax)**

The clientId parameter in the following command is a unique, user-specified ID to
identify the client for the configuration. This ID enables AWS AppConfig to deploy
the configuration in intervals, as defined in the deployment strategy.

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[Configuration](#API_GetConfiguration_RequestSyntax)**

The configuration to get. Specify either the configuration name or the configuration
ID.

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[Environment](#API_GetConfiguration_RequestSyntax)**

The environment to get. Specify either the environment name or the environment
ID.

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Configuration-Version: ConfigurationVersion
Content-Type: ContentType

Content
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The response returns the following HTTP headers.

**[ConfigurationVersion](#API_GetConfiguration_ResponseSyntax)**

The configuration version.

Length Constraints: Minimum length of 1. Maximum length of 1024.

**[ContentType](#API_GetConfiguration_ResponseSyntax)**

A standard MIME type describing the format of the configuration content. For more
information, see [Content-Type](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).

The response returns the following as the HTTP body.

**[Content](#API_GetConfiguration_ResponseSyntax)**

The content of the configuration or the configuration data.

###### Important

The `Content` attribute only contains data if the system finds new or
updated configuration data. If there is no new or updated data and
`ClientConfigurationVersion` matches the version of the current
configuration, AWS AppConfig returns a `204 No Content` HTTP response
code and the `Content` value will be empty.

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

This example illustrates one usage of GetConfiguration.

#### Sample Request

```

GET /applications/test-application/environments/Example-Environment/configurations/Example-Configuration-Profile?client_id=test-id HTTP/1.1
Host: appconfig.us-east-1.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.2.4 Python/3.8.8 Linux/5.4.134-73.228.amzn2int.x86_64 exe/x86_64.amzn.2 prompt/off command/appconfig.get-configuration
X-Amz-Date: 20210917T215745Z
Authorization: AWS4-HMAC-SHA256 Credential=AWS_ACCESS_KEY_ID_REDACTED/20210917/us-east-1/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
```

#### Sample Response

```

{
    "ConfigurationVersion": "1",
    "ContentType": "application/octet-stream"
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/getconfiguration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/getconfiguration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/getconfiguration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/getconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/getconfiguration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/getconfiguration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/getconfiguration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/getconfiguration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/getconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/getconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetApplication

GetConfigurationProfile

All content copied from https://docs.aws.amazon.com/.
