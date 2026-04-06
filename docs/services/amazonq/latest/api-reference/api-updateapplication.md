# UpdateApplication

Updates an existing Amazon Q Business application.

###### Note

Amazon Q Business applications may securely transmit data for processing across
AWS Regions within your geography. For more information, see
[Cross region\
inference in Amazon Q Business](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/cross-region-inference.html).

###### Note

An Amazon Q Apps service-linked role will be created if it's absent in the
AWS account when `QAppsConfiguration` is enabled in
the request. For more information, see [Using\
service-linked roles for Q Apps](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/using-service-linked-roles-qapps.html).

## Request Syntax

```nohighlight

PUT /applications/applicationId HTTP/1.1
Content-type: application/json

{
   "attachmentsConfiguration": {
      "attachmentsControlMode": "string"
   },
   "autoSubscriptionConfiguration": {
      "autoSubscribe": "string",
      "defaultSubscriptionType": "string"
   },
   "description": "string",
   "displayName": "string",
   "identityCenterInstanceArn": "string",
   "personalizationConfiguration": {
      "personalizationControlMode": "string"
   },
   "qAppsConfiguration": {
      "qAppsControlMode": "string"
   },
   "roleArn": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_UpdateApplication_RequestSyntax)**

The identifier of the Amazon Q Business application.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[attachmentsConfiguration](#API_UpdateApplication_RequestSyntax)**

An option to allow end users to upload files directly during chat.

Type: [AttachmentsConfiguration](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_AttachmentsConfiguration.html) object

Required: No

**[autoSubscriptionConfiguration](#API_UpdateApplication_RequestSyntax)**

An option to enable updating the default subscription type assigned to an Amazon Q Business application using IAM identity federation for user
management.

Type: [AutoSubscriptionConfiguration](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_AutoSubscriptionConfiguration.html) object

Required: No

**[description](#API_UpdateApplication_RequestSyntax)**

A description for the Amazon Q Business application.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1000.

Pattern: `[\s\S]*`

Required: No

**[displayName](#API_UpdateApplication_RequestSyntax)**

A name for the Amazon Q Business application.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`

Required: No

**[identityCenterInstanceArn](#API_UpdateApplication_RequestSyntax)**

The Amazon Resource Name (ARN) of the IAM Identity Center instance you are either
creating for—or connecting to—your Amazon Q Business application.

Type: String

Length Constraints: Minimum length of 10. Maximum length of 1224.

Pattern: `arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}`

Required: No

**[personalizationConfiguration](#API_UpdateApplication_RequestSyntax)**

Configuration information about chat response personalization. For more information,
see [Personalizing chat responses](../qbusiness-ug/personalizing-chat-responses.md).

Type: [PersonalizationConfiguration](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_PersonalizationConfiguration.html) object

Required: No

**[qAppsConfiguration](#API_UpdateApplication_RequestSyntax)**

An option to allow end users to create and use Amazon Q Apps in the web
experience.

Type: [QAppsConfiguration](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_QAppsConfiguration.html) object

Required: No

**[roleArn](#API_UpdateApplication_RequestSyntax)**

An AWS Identity and Access Management (IAM) role that
gives Amazon Q Business permission to access Amazon CloudWatch logs and
metrics.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

Required: No

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have access to perform this action. Make sure you have the required
permission policies and user accounts and try again.

HTTP Status Code: 403

**ConflictException**

You are trying to perform an action that conflicts with the current status of your
resource. Fix any inconsistencies with your resources and try again.

**message**

The message describing a `ConflictException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 409

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/qbusiness-2023-11-27/UpdateApplication)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/qbusiness-2023-11-27/UpdateApplication)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/UpdateApplication)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/qbusiness-2023-11-27/UpdateApplication)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/UpdateApplication)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/qbusiness-2023-11-27/UpdateApplication)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/qbusiness-2023-11-27/UpdateApplication)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/qbusiness-2023-11-27/UpdateApplication)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/qbusiness-2023-11-27/UpdateApplication)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/UpdateApplication)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UntagResource

UpdateChatControlsConfiguration
