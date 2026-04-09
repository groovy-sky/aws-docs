# GetApplication

Gets information about an existing Amazon Q Business application.

## Request Syntax

```nohighlight

GET /applications/applicationId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_GetApplication_RequestSyntax)**

The identifier of the Amazon Q Business application.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "applicationArn": "string",
   "applicationId": "string",
   "attachmentsConfiguration": {
      "attachmentsControlMode": "string"
   },
   "autoSubscriptionConfiguration": {
      "autoSubscribe": "string",
      "defaultSubscriptionType": "string"
   },
   "clientIdsForOIDC": [ "string" ],
   "createdAt": number,
   "description": "string",
   "displayName": "string",
   "encryptionConfiguration": {
      "kmsKeyId": "string"
   },
   "error": {
      "errorCode": "string",
      "errorMessage": "string"
   },
   "iamIdentityProviderArn": "string",
   "identityCenterApplicationArn": "string",
   "identityType": "string",
   "personalizationConfiguration": {
      "personalizationControlMode": "string"
   },
   "qAppsConfiguration": {
      "qAppsControlMode": "string"
   },
   "quickSightConfiguration": {
      "clientNamespace": "string"
   },
   "roleArn": "string",
   "status": "string",
   "updatedAt": number
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[applicationArn](#API_GetApplication_ResponseSyntax)**

The Amazon Resource Name (ARN) of the Amazon Q Business application.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

**[applicationId](#API_GetApplication_ResponseSyntax)**

The identifier of the Amazon Q Business application.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

**[attachmentsConfiguration](#API_GetApplication_ResponseSyntax)**

Settings for whether end users can upload files directly during chat.

Type: [AppliedAttachmentsConfiguration](api-appliedattachmentsconfiguration.md) object

**[autoSubscriptionConfiguration](#API_GetApplication_ResponseSyntax)**

Settings for auto-subscription behavior for this application. This is only applicable
to SAML and OIDC applications.

Type: [AutoSubscriptionConfiguration](api-autosubscriptionconfiguration.md) object

**[clientIdsForOIDC](#API_GetApplication_ResponseSyntax)**

The OIDC client ID for a Amazon Q Business application.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.:/()*?=-]*`

**[createdAt](#API_GetApplication_ResponseSyntax)**

The Unix timestamp when the Amazon Q Business application was last updated.

Type: Timestamp

**[description](#API_GetApplication_ResponseSyntax)**

A description for the Amazon Q Business application.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1000.

Pattern: `[\s\S]*`

**[displayName](#API_GetApplication_ResponseSyntax)**

The name of the Amazon Q Business application.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`

**[encryptionConfiguration](#API_GetApplication_ResponseSyntax)**

The identifier of the AWS
AWS KMS key that is used to encrypt your data. Amazon Q Business doesn't support
asymmetric keys.

Type: [EncryptionConfiguration](api-encryptionconfiguration.md) object

**[error](#API_GetApplication_ResponseSyntax)**

If the `Status` field is set to `ERROR`, the
`ErrorMessage` field contains a description of the error that caused the
synchronization to fail.

Type: [ErrorDetail](api-errordetail.md) object

**[iamIdentityProviderArn](#API_GetApplication_ResponseSyntax)**

The Amazon Resource Name (ARN) of an identity provider being used by an Amazon Q Business application.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:aws:iam::\d{12}:(oidc-provider|saml-provider)/[a-zA-Z0-9_\.\/@\-]+`

**[identityCenterApplicationArn](#API_GetApplication_ResponseSyntax)**

The Amazon Resource Name (ARN) of the AWS IAM Identity Center instance attached to
your Amazon Q Business application.

Type: String

Length Constraints: Minimum length of 10. Maximum length of 1224.

Pattern: `arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso::\d{12}:application/(sso)?ins-[a-zA-Z0-9-.]{16}/apl-[a-zA-Z0-9]{16}`

**[identityType](#API_GetApplication_ResponseSyntax)**

The authentication type being used by a Amazon Q Business application.

Type: String

Valid Values: `AWS_IAM_IDP_SAML | AWS_IAM_IDP_OIDC | AWS_IAM_IDC | AWS_QUICKSIGHT_IDP | ANONYMOUS`

**[personalizationConfiguration](#API_GetApplication_ResponseSyntax)**

Configuration information about chat response personalization. For more information,
see [Personalizing chat responses](../qbusiness-ug/personalizing-chat-responses.md).

Type: [PersonalizationConfiguration](api-personalizationconfiguration.md) object

**[qAppsConfiguration](#API_GetApplication_ResponseSyntax)**

Settings for whether end users can create and use Amazon Q Apps in the web
experience.

Type: [QAppsConfiguration](api-qappsconfiguration.md) object

**[quickSightConfiguration](#API_GetApplication_ResponseSyntax)**

The Amazon Quick authentication configuration for the Amazon Q Business application.

Type: [QuickSightConfiguration](api-quicksightconfiguration.md) object

**[roleArn](#API_GetApplication_ResponseSyntax)**

The Amazon Resource Name (ARN) of the IAM with permissions to access
your CloudWatch logs and metrics.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

**[status](#API_GetApplication_ResponseSyntax)**

The status of the Amazon Q Business application.

Type: String

Valid Values: `CREATING | ACTIVE | DELETING | FAILED | UPDATING`

**[updatedAt](#API_GetApplication_ResponseSyntax)**

The Unix timestamp when the Amazon Q Business application was last updated.

Type: Timestamp

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have access to perform this action. Make sure you have the required
permission policies and user accounts and try again.

HTTP Status Code: 403

**InternalServerException**

An issue occurred with the internal server used for your Amazon Q Business service. Wait
some minutes and try again, or contact [Support](http://aws.amazon.com/contact-us) for help.

HTTP Status Code: 500

**ResourceNotFoundException**

The application or plugin resource you want to use doesn’t exist. Make sure you have
provided the correct resource and try again.

**message**

The message describing a `ResourceNotFoundException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 404

**ThrottlingException**

The request was denied due to throttling. Reduce the number of requests and try
again.

HTTP Status Code: 429

**ValidationException**

The input doesn't meet the constraints set by the Amazon Q Business service. Provide the
correct input and try again.

**fields**

The input field(s) that failed validation.

**message**

The message describing the `ValidationException`.

**reason**

The reason for the `ValidationException`.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/qbusiness-2023-11-27/getapplication.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qbusiness-2023-11-27/getapplication.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/getapplication.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qbusiness-2023-11-27/getapplication.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/getapplication.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qbusiness-2023-11-27/getapplication.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qbusiness-2023-11-27/getapplication.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qbusiness-2023-11-27/getapplication.md)

- [AWS SDK for Python](../../../goto/boto3/qbusiness-2023-11-27/getapplication.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/getapplication.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DisassociatePermission

GetChatControlsConfiguration

All content copied from https://docs.aws.amazon.com/.
