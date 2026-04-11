# UpdateWebExperience

Updates an Amazon Q Business web experience.

## Request Syntax

```nohighlight

PUT /applications/applicationId/experiences/webExperienceId HTTP/1.1
Content-type: application/json

{
   "authenticationConfiguration": { ... },
   "browserExtensionConfiguration": {
      "enabledBrowserExtensions": [ "string" ]
   },
   "customizationConfiguration": {
      "customCSSUrl": "string",
      "faviconUrl": "string",
      "fontUrl": "string",
      "logoUrl": "string"
   },
   "identityProviderConfiguration": { ... },
   "origins": [ "string" ],
   "roleArn": "string",
   "samplePromptsControlMode": "string",
   "subtitle": "string",
   "title": "string",
   "welcomeMessage": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_UpdateWebExperience_RequestSyntax)**

The identifier of the Amazon Q Business application attached to the web experience.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[webExperienceId](#API_UpdateWebExperience_RequestSyntax)**

The identifier of the Amazon Q Business web experience.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]*`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[authenticationConfiguration](#API_UpdateWebExperience_RequestSyntax)**

_This parameter has been deprecated._

The authentication configuration of the Amazon Q Business web experience.

Type: [WebExperienceAuthConfiguration](api-webexperienceauthconfiguration.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**[browserExtensionConfiguration](#API_UpdateWebExperience_RequestSyntax)**

The browser extension configuration for an Amazon Q Business web experience.

###### Note

For Amazon Q Business application using external OIDC-compliant identity providers (IdPs).
The IdP administrator must add the browser extension sign-in redirect URLs to the IdP application.
For more information, see [Configure external OIDC identity provider for your browser extensions.](../qbusiness-ug/browser-extensions.md).

Type: [BrowserExtensionConfiguration](api-browserextensionconfiguration.md) object

Required: No

**[customizationConfiguration](#API_UpdateWebExperience_RequestSyntax)**

Updates the custom logo, favicon, font, and color used in the Amazon Q web experience.

Type: [CustomizationConfiguration](api-customizationconfiguration.md) object

Required: No

**[identityProviderConfiguration](#API_UpdateWebExperience_RequestSyntax)**

Information about the identity provider (IdP) used to authenticate end users of an
Amazon Q Business web experience.

Type: [IdentityProviderConfiguration](api-identityproviderconfiguration.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**[origins](#API_UpdateWebExperience_RequestSyntax)**

Updates the website domain origins that
are allowed to embed the Amazon Q Business web experience.

The _domain origin_ refers to the
_base URL_ for accessing a website including the protocol
( `http/https`), the domain name, and the port number (if specified).

###### Note

- Any values except `null` submitted as part of this
update will replace all previous values.

- You must only submit a _base URL_ and
not a full path. For example, `https://docs.aws.amazon.com`.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 10 items.

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `(http://|https://)[a-zA-Z0-9-_.]+(?::[0-9]{1,5})?`

Required: No

**[roleArn](#API_UpdateWebExperience_RequestSyntax)**

The Amazon Resource Name (ARN) of the role with permission to access the Amazon Q Business
web experience and required resources.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

Required: No

**[samplePromptsControlMode](#API_UpdateWebExperience_RequestSyntax)**

Determines whether sample prompts are enabled in the web experience for an end
user.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**[subtitle](#API_UpdateWebExperience_RequestSyntax)**

The subtitle of the Amazon Q Business web experience.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 500.

Pattern: `[\s\S]*`

Required: No

**[title](#API_UpdateWebExperience_RequestSyntax)**

The title of the Amazon Q Business web experience.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 500.

Pattern: `[\s\S]*`

Required: No

**[welcomeMessage](#API_UpdateWebExperience_RequestSyntax)**

A customized welcome message for an end user in an Amazon Q Business web
experience.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 300.

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

- [AWS Command Line Interface V2](../../../goto/cli2/qbusiness-2023-11-27/updatewebexperience.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qbusiness-2023-11-27/updatewebexperience.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/updatewebexperience.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qbusiness-2023-11-27/updatewebexperience.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/updatewebexperience.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qbusiness-2023-11-27/updatewebexperience.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qbusiness-2023-11-27/updatewebexperience.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qbusiness-2023-11-27/updatewebexperience.md)

- [AWS SDK for Python](../../../goto/boto3/qbusiness-2023-11-27/updatewebexperience.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/updatewebexperience.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateUser

QApps

All content copied from https://docs.aws.amazon.com/.
