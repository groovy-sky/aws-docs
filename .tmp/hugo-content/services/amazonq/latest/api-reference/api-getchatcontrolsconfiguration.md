# GetChatControlsConfiguration

Gets information about chat controls configured for an existing Amazon Q Business
application.

## Request Syntax

```nohighlight

GET /applications/applicationId/chatcontrols?maxResults=maxResults&nextToken=nextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_GetChatControlsConfiguration_RequestSyntax)**

The identifier of the application for which the chat controls are configured.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[maxResults](#API_GetChatControlsConfiguration_RequestSyntax)**

The maximum number of configured chat controls to return.

Valid Range: Minimum value of 1. Maximum value of 50.

**[nextToken](#API_GetChatControlsConfiguration_RequestSyntax)**

If the `maxResults` response was incomplete because there is more data to
retrieve, Amazon Q Business returns a pagination token in the response. You can use this
pagination token to retrieve the next set of Amazon Q Business chat controls
configured.

Length Constraints: Minimum length of 1. Maximum length of 800.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "blockedPhrases": {
      "blockedPhrases": [ "string" ],
      "systemMessageOverride": "string"
   },
   "creatorModeConfiguration": {
      "creatorModeControl": "string"
   },
   "hallucinationReductionConfiguration": {
      "hallucinationReductionControl": "string"
   },
   "nextToken": "string",
   "orchestrationConfiguration": {
      "control": "string"
   },
   "responseScope": "string",
   "topicConfigurations": [
      {
         "description": "string",
         "exampleChatMessages": [ "string" ],
         "name": "string",
         "rules": [
            {
               "excludedUsersAndGroups": {
                  "userGroups": [ "string" ],
                  "userIds": [ "string" ]
               },
               "includedUsersAndGroups": {
                  "userGroups": [ "string" ],
                  "userIds": [ "string" ]
               },
               "ruleConfiguration": { ... },
               "ruleType": "string"
            }
         ]
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[blockedPhrases](#API_GetChatControlsConfiguration_ResponseSyntax)**

The phrases blocked from chat by your chat control configuration.

Type: [BlockedPhrasesConfiguration](api-blockedphrasesconfiguration.md) object

**[creatorModeConfiguration](#API_GetChatControlsConfiguration_ResponseSyntax)**

The configuration details for `CREATOR_MODE`.

Type: [AppliedCreatorModeConfiguration](api-appliedcreatormodeconfiguration.md) object

**[hallucinationReductionConfiguration](#API_GetChatControlsConfiguration_ResponseSyntax)**

The hallucination reduction settings for your application.

Type: [HallucinationReductionConfiguration](api-hallucinationreductionconfiguration.md) object

**[nextToken](#API_GetChatControlsConfiguration_ResponseSyntax)**

If the `maxResults` response was incomplete because there is more data to
retrieve, Amazon Q Business returns a pagination token in the response. You can use this
pagination token to retrieve the next set of Amazon Q Business chat controls
configured.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 800.

**[orchestrationConfiguration](#API_GetChatControlsConfiguration_ResponseSyntax)**

The chat response orchestration settings for your application.

###### Note

Chat orchestration is optimized to work for English language content. For more
details on language support in Amazon Q Business, see [Supported languages](../qbusiness-ug/supported-languages.md).

Type: [AppliedOrchestrationConfiguration](api-appliedorchestrationconfiguration.md) object

**[responseScope](#API_GetChatControlsConfiguration_ResponseSyntax)**

The response scope configured for a Amazon Q Business application. This determines whether
your application uses its retrieval augmented generation (RAG) system to generate
answers only from your enterprise data, or also uses the large language models (LLM)
knowledge to respons to end user questions in chat.

Type: String

Valid Values: `ENTERPRISE_CONTENT_ONLY | EXTENDED_KNOWLEDGE_ENABLED`

**[topicConfigurations](#API_GetChatControlsConfiguration_ResponseSyntax)**

The topic specific controls configured for a Amazon Q Business application.

Type: Array of [TopicConfiguration](api-topicconfiguration.md) objects

Array Members: Minimum number of 0 items. Maximum number of 10 items.

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

- [AWS Command Line Interface V2](../../../goto/cli2/qbusiness-2023-11-27/getchatcontrolsconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qbusiness-2023-11-27/getchatcontrolsconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/getchatcontrolsconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qbusiness-2023-11-27/getchatcontrolsconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/getchatcontrolsconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qbusiness-2023-11-27/getchatcontrolsconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qbusiness-2023-11-27/getchatcontrolsconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qbusiness-2023-11-27/getchatcontrolsconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/qbusiness-2023-11-27/getchatcontrolsconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/getchatcontrolsconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetApplication

GetChatResponseConfiguration

All content copied from https://docs.aws.amazon.com/.
