# ListPluginTypeActions

Lists configured Amazon Q Business actions for any plugin type—both
built-in and custom.

## Request Syntax

```nohighlight

GET /pluginTypes/pluginType/actions?maxResults=maxResults&nextToken=nextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[maxResults](#API_ListPluginTypeActions_RequestSyntax)**

The maximum number of plugins to return.

Valid Range: Minimum value of 1. Maximum value of 50.

**[nextToken](#API_ListPluginTypeActions_RequestSyntax)**

If the number of plugins returned exceeds `maxResults`, Amazon Q Business returns a next token as a pagination token to retrieve the next set
of plugins.

Length Constraints: Minimum length of 1. Maximum length of 800.

**[pluginType](#API_ListPluginTypeActions_RequestSyntax)**

The type of the plugin.

Valid Values: `SERVICE_NOW | SALESFORCE | JIRA | ZENDESK | CUSTOM | QUICKSIGHT | SERVICENOW_NOW_PLATFORM | JIRA_CLOUD | SALESFORCE_CRM | ZENDESK_SUITE | ATLASSIAN_CONFLUENCE | GOOGLE_CALENDAR | MICROSOFT_TEAMS | MICROSOFT_EXCHANGE | PAGERDUTY_ADVANCE | SMARTSHEET | ASANA`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "items": [
      {
         "actionIdentifier": "string",
         "description": "string",
         "displayName": "string",
         "instructionExample": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[items](#API_ListPluginTypeActions_ResponseSyntax)**

An array of information on one or more plugins.

Type: Array of [ActionSummary](api-actionsummary.md) objects

**[nextToken](#API_ListPluginTypeActions_ResponseSyntax)**

If the response is truncated, Amazon Q Business returns this token, which you
can use in a later request to list the next set of plugins.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 800.

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

- [AWS Command Line Interface V2](../../../goto/cli2/qbusiness-2023-11-27/listplugintypeactions.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qbusiness-2023-11-27/listplugintypeactions.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/listplugintypeactions.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qbusiness-2023-11-27/listplugintypeactions.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/listplugintypeactions.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qbusiness-2023-11-27/listplugintypeactions.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qbusiness-2023-11-27/listplugintypeactions.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qbusiness-2023-11-27/listplugintypeactions.md)

- [AWS SDK for Python](../../../goto/boto3/qbusiness-2023-11-27/listplugintypeactions.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/listplugintypeactions.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListPlugins

ListPluginTypeMetadata

All content copied from https://docs.aws.amazon.com/.
