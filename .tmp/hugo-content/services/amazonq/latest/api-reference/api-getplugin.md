# GetPlugin

Gets information about an existing Amazon Q Business plugin.

## Request Syntax

```nohighlight

GET /applications/applicationId/plugins/pluginId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_GetPlugin_RequestSyntax)**

The identifier of the application which contains the plugin.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[pluginId](#API_GetPlugin_RequestSyntax)**

The identifier of the plugin.

Length Constraints: Fixed length of 36.

Pattern: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "applicationId": "string",
   "authConfiguration": { ... },
   "buildStatus": "string",
   "createdAt": number,
   "customPluginConfiguration": {
      "apiSchema": { ... },
      "apiSchemaType": "string",
      "description": "string"
   },
   "displayName": "string",
   "pluginArn": "string",
   "pluginId": "string",
   "serverUrl": "string",
   "state": "string",
   "type": "string",
   "updatedAt": number
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[applicationId](#API_GetPlugin_ResponseSyntax)**

The identifier of the application which contains the plugin.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

**[authConfiguration](#API_GetPlugin_ResponseSyntax)**

Authentication configuration information for an Amazon Q Business plugin.

Type: [PluginAuthConfiguration](api-pluginauthconfiguration.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

**[buildStatus](#API_GetPlugin_ResponseSyntax)**

The current status of a plugin. A plugin is modified asynchronously.

Type: String

Valid Values: `READY | CREATE_IN_PROGRESS | CREATE_FAILED | UPDATE_IN_PROGRESS | UPDATE_FAILED | DELETE_IN_PROGRESS | DELETE_FAILED`

**[createdAt](#API_GetPlugin_ResponseSyntax)**

The timestamp for when the plugin was created.

Type: Timestamp

**[customPluginConfiguration](#API_GetPlugin_ResponseSyntax)**

Configuration information required to create a custom plugin.

Type: [CustomPluginConfiguration](api-custompluginconfiguration.md) object

**[displayName](#API_GetPlugin_ResponseSyntax)**

The name of the plugin.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`

**[pluginArn](#API_GetPlugin_ResponseSyntax)**

The Amazon Resource Name (ARN) of the role with permission to access resources needed
to create the plugin.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

**[pluginId](#API_GetPlugin_ResponseSyntax)**

The identifier of the plugin.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

**[serverUrl](#API_GetPlugin_ResponseSyntax)**

The source URL used for plugin configuration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `(https?|ftp|file)://([^\s]*)`

**[state](#API_GetPlugin_ResponseSyntax)**

The current state of the plugin.

Type: String

Valid Values: `ENABLED | DISABLED`

**[type](#API_GetPlugin_ResponseSyntax)**

The type of the plugin.

Type: String

Valid Values: `SERVICE_NOW | SALESFORCE | JIRA | ZENDESK | CUSTOM | QUICKSIGHT | SERVICENOW_NOW_PLATFORM | JIRA_CLOUD | SALESFORCE_CRM | ZENDESK_SUITE | ATLASSIAN_CONFLUENCE | GOOGLE_CALENDAR | MICROSOFT_TEAMS | MICROSOFT_EXCHANGE | PAGERDUTY_ADVANCE | SMARTSHEET | ASANA`

**[updatedAt](#API_GetPlugin_ResponseSyntax)**

The timestamp for when the plugin was last updated.

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

- [AWS Command Line Interface V2](../../../goto/cli2/qbusiness-2023-11-27/getplugin.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qbusiness-2023-11-27/getplugin.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/getplugin.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qbusiness-2023-11-27/getplugin.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/getplugin.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qbusiness-2023-11-27/getplugin.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qbusiness-2023-11-27/getplugin.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qbusiness-2023-11-27/getplugin.md)

- [AWS SDK for Python](../../../goto/boto3/qbusiness-2023-11-27/getplugin.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/getplugin.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetMedia

GetPolicy

All content copied from https://docs.aws.amazon.com/.
