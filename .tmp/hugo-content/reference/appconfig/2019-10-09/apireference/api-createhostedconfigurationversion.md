# CreateHostedConfigurationVersion

Creates a new configuration in the AWS AppConfig hosted configuration store. If
you're creating a feature flag, we recommend you familiarize yourself with the JSON schema
for feature flag data. For more information, see [Type reference for AWS.AppConfig.FeatureFlags](../../../../services/appconfig/latest/userguide/appconfig-type-reference-feature-flags.md) in the
_AWS AppConfig User Guide_.

## Request Syntax

```nohighlight

POST /applications/ApplicationId/configurationprofiles/ConfigurationProfileId/hostedconfigurationversions HTTP/1.1
Description: Description
Content-Type: ContentType
Latest-Version-Number: LatestVersionNumber
VersionLabel: VersionLabel

Content
```

## URI Request Parameters

The request uses the following URI parameters.

**[ApplicationId](#API_CreateHostedConfigurationVersion_RequestSyntax)**

The application ID.

Pattern: `[a-z0-9]{4,7}`

Required: Yes

**[ConfigurationProfileId](#API_CreateHostedConfigurationVersion_RequestSyntax)**

The configuration profile ID.

Pattern: `[a-z0-9]{4,7}`

Required: Yes

**[ContentType](#API_CreateHostedConfigurationVersion_RequestSyntax)**

A standard MIME type describing the format of the configuration content. For more
information, see [Content-Type](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**[Description](#API_CreateHostedConfigurationVersion_RequestSyntax)**

A description of the configuration.

###### Note

Due to HTTP limitations, this field only supports ASCII characters.

Length Constraints: Minimum length of 0. Maximum length of 1024.

**[LatestVersionNumber](#API_CreateHostedConfigurationVersion_RequestSyntax)**

An optional locking token used to prevent race conditions from overwriting configuration
updates when creating a new version. To ensure your data is not overwritten when creating
multiple hosted configuration versions in rapid succession, specify the version number of
the latest hosted configuration version.

**[VersionLabel](#API_CreateHostedConfigurationVersion_RequestSyntax)**

An optional, user-defined label for the AWS AppConfig hosted configuration
version. This value must contain at least one non-numeric character. For example,
"v2.2.0".

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `.*[^0-9].*`

## Request Body

The request accepts the following binary data.

**[Content](#API_CreateHostedConfigurationVersion_RequestSyntax)**

The configuration data, as bytes.

###### Note

AWS AppConfig accepts any type of data, including text formats like JSON or
TOML, or binary formats like protocol buffers or compressed data.

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 201
Application-Id: ApplicationId
Configuration-Profile-Id: ConfigurationProfileId
Version-Number: VersionNumber
Description: Description
Content-Type: ContentType
VersionLabel: VersionLabel
KmsKeyArn: KmsKeyArn

Content
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The response returns the following HTTP headers.

**[ApplicationId](#API_CreateHostedConfigurationVersion_ResponseSyntax)**

The application ID.

Pattern: `[a-z0-9]{4,7}`

**[ConfigurationProfileId](#API_CreateHostedConfigurationVersion_ResponseSyntax)**

The configuration profile ID.

Pattern: `[a-z0-9]{4,7}`

**[ContentType](#API_CreateHostedConfigurationVersion_ResponseSyntax)**

A standard MIME type describing the format of the configuration content. For more
information, see [Content-Type](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).

Length Constraints: Minimum length of 1. Maximum length of 255.

**[Description](#API_CreateHostedConfigurationVersion_ResponseSyntax)**

A description of the configuration.

Length Constraints: Minimum length of 0. Maximum length of 1024.

**[KmsKeyArn](#API_CreateHostedConfigurationVersion_ResponseSyntax)**

The Amazon Resource Name of the AWS Key Management Service key that was used to encrypt this
specific version of the configuration data in the AWS AppConfig hosted configuration
store.

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

**[VersionLabel](#API_CreateHostedConfigurationVersion_ResponseSyntax)**

A user-defined label for an AWS AppConfig hosted configuration version.

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `.*[^0-9].*`

**[VersionNumber](#API_CreateHostedConfigurationVersion_ResponseSyntax)**

The configuration version.

The response returns the following as the HTTP body.

**[Content](#API_CreateHostedConfigurationVersion_ResponseSyntax)**

The content of the configuration or the configuration data.

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

**PayloadTooLargeException**

The configuration size is too large.

HTTP Status Code: 413

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

This example illustrates one usage of CreateHostedConfigurationVersion.

#### Sample Request

```

POST /applications/abc1234/configurationprofiles/ur8hx2f/hostedconfigurationversions HTTP/1.1
Host: appconfig.us-east-1.amazonaws.com
Accept-Encoding: identity
Content-Type: application/json
User-Agent: aws-cli/2.2.4 Python/3.8.8 Linux/5.4.134-73.228.amzn2int.x86_64 exe/x86_64.amzn.2 prompt/off command/appconfig.create-hosted-configuration-version
X-Amz-Date: 20210917T184857Z
Authorization: AWS4-HMAC-SHA256 Credential=AWS_ACCESS_KEY_ID_REDACTED/20210917/us-east-1/appconfig/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
Content-Length: 27

{ "Name": "ExampleApplication", "Id": ExampleID, "Rank": 7 }
```

#### Sample Response

```

{
    "ApplicationId": "abc1234",
    "ConfigurationProfileId": "ur8hx2f",
    "VersionNumber": "1",
    "ContentType": "application/json"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/createhostedconfigurationversion.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/createhostedconfigurationversion.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/createhostedconfigurationversion.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/createhostedconfigurationversion.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/createhostedconfigurationversion.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/createhostedconfigurationversion.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/createhostedconfigurationversion.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/createhostedconfigurationversion.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/createhostedconfigurationversion.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/createhostedconfigurationversion.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateExtensionAssociation

DeleteApplication

All content copied from https://docs.aws.amazon.com/.
