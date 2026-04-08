# UpdateConfigurationProfile

Updates a configuration profile.

## Request Syntax

```nohighlight

PATCH /applications/ApplicationId/configurationprofiles/ConfigurationProfileId HTTP/1.1
Content-type: application/json

{
   "Description": "string",
   "KmsKeyIdentifier": "string",
   "Name": "string",
   "RetrievalRoleArn": "string",
   "Validators": [
      {
         "Content": "string",
         "Type": "string"
      }
   ]
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[ApplicationId](#API_UpdateConfigurationProfile_RequestSyntax)**

The application ID.

Pattern: `[a-z0-9]{4,7}`

Required: Yes

**[ConfigurationProfileId](#API_UpdateConfigurationProfile_RequestSyntax)**

The ID of the configuration profile.

Pattern: `[a-z0-9]{4,7}`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[Description](#API_UpdateConfigurationProfile_RequestSyntax)**

A description of the configuration profile.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**[KmsKeyIdentifier](#API_UpdateConfigurationProfile_RequestSyntax)**

The identifier for a AWS Key Management Service key to encrypt new configuration data
versions in the AWS AppConfig hosted configuration store. This attribute is only used
for `hosted` configuration types. The identifier can be an AWS KMS
key ID, alias, or the Amazon Resource Name (ARN) of the key ID or alias. To encrypt data
managed in other configuration stores, see the documentation for how to specify an AWS KMS key for that particular service.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 2048.

Required: No

**[Name](#API_UpdateConfigurationProfile_RequestSyntax)**

The name of the configuration profile.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**[RetrievalRoleArn](#API_UpdateConfigurationProfile_RequestSyntax)**

The ARN of an IAM role with permission to access the configuration at the specified
`LocationUri`.

###### Important

A retrieval role ARN is not required for configurations stored in AWS CodePipeline or the AWS AppConfig hosted configuration store. It is required for all other sources that
store your configuration.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^((arn):(aws|aws-cn|aws-iso|aws-iso-[a-z]{1}|aws-us-gov|aws-eusc):(iam)::\d{12}:role[/].*)$`

Required: No

**[Validators](#API_UpdateConfigurationProfile_RequestSyntax)**

A list of methods for validating the configuration.

Type: Array of [Validator](api-validator.md) objects

Array Members: Minimum number of 0 items. Maximum number of 2 items.

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "ApplicationId": "string",
   "Description": "string",
   "Id": "string",
   "KmsKeyArn": "string",
   "KmsKeyIdentifier": "string",
   "LocationUri": "string",
   "Name": "string",
   "RetrievalRoleArn": "string",
   "Type": "string",
   "Validators": [
      {
         "Content": "string",
         "Type": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ApplicationId](#API_UpdateConfigurationProfile_ResponseSyntax)**

The application ID.

Type: String

Pattern: `[a-z0-9]{4,7}`

**[Description](#API_UpdateConfigurationProfile_ResponseSyntax)**

The configuration profile description.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

**[Id](#API_UpdateConfigurationProfile_ResponseSyntax)**

The configuration profile ID.

Type: String

Pattern: `[a-z0-9]{4,7}`

**[KmsKeyArn](#API_UpdateConfigurationProfile_ResponseSyntax)**

The Amazon Resource Name of the AWS Key Management Service key to encrypt new configuration
data versions in the AWS AppConfig hosted configuration store. This attribute is only
used for `hosted` configuration types. To encrypt data managed in other
configuration stores, see the documentation for how to specify an AWS KMS key
for that particular service.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

**[KmsKeyIdentifier](#API_UpdateConfigurationProfile_ResponseSyntax)**

The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when
the resource was created or updated.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[LocationUri](#API_UpdateConfigurationProfile_ResponseSyntax)**

The URI location of the configuration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[Name](#API_UpdateConfigurationProfile_ResponseSyntax)**

The name of the configuration profile.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

**[RetrievalRoleArn](#API_UpdateConfigurationProfile_ResponseSyntax)**

The ARN of an IAM role with permission to access the configuration at the specified
`LocationUri`.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^((arn):(aws|aws-cn|aws-iso|aws-iso-[a-z]{1}|aws-us-gov|aws-eusc):(iam)::\d{12}:role[/].*)$`

**[Type](#API_UpdateConfigurationProfile_ResponseSyntax)**

The type of configurations contained in the profile. AWS AppConfig supports
`feature flags` and `freeform` configurations. We recommend you
create feature flag configurations to enable or disable new features and freeform
configurations to distribute configurations to an application. When calling this API, enter
one of the following values for `Type`:

`AWS.AppConfig.FeatureFlags`

`AWS.Freeform`

Type: String

Pattern: `^[a-zA-Z\.]+`

**[Validators](#API_UpdateConfigurationProfile_ResponseSyntax)**

A list of methods for validating the configuration.

Type: Array of [Validator](api-validator.md) objects

Array Members: Minimum number of 0 items. Maximum number of 2 items.

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

This example illustrates one usage of UpdateConfigurationProfile.

#### Sample Request

```

PATCH /applications/abc1234/configurationprofiles/ur8hx2f HTTP/1.1
Host: appconfig.us-east-1.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.2.4 Python/3.8.8 Linux/5.4.134-73.228.amzn2int.x86_64 exe/x86_64.amzn.2 prompt/off command/appconfig.update-configuration-profile
X-Amz-Date: 20210920T213335Z
Authorization: AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20210920/us-east-1/appconfig/aws4_request, SignedHeaders=host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
Content-Length: 59

{
    "Description": "Configuration profile used for examples."
}
```

#### Sample Response

```

{
    "ApplicationId": "abc1234",
    "Id": "ur8hx2f",
    "Name": "Example-Configuration-Profile",
    "Description": "Configuration profile used for examples.",
    "LocationUri": "ssm-parameter://Example-Parameter",
    "RetrievalRoleArn": "arn:aws:iam::682428703967:role/Example-App-Config-Role"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/updateconfigurationprofile.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/updateconfigurationprofile.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/updateconfigurationprofile.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/updateconfigurationprofile.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/updateconfigurationprofile.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/updateconfigurationprofile.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/updateconfigurationprofile.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/updateconfigurationprofile.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/updateconfigurationprofile.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/updateconfigurationprofile.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UpdateApplication

UpdateDeploymentStrategy
