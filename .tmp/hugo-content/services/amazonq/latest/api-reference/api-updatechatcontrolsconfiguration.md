# UpdateChatControlsConfiguration

Updates a set of chat controls configured for an existing Amazon Q Business
application.

## Request Syntax

```nohighlight

PATCH /applications/applicationId/chatcontrols HTTP/1.1
Content-type: application/json

{
   "blockedPhrasesConfigurationUpdate": {
      "blockedPhrasesToCreateOrUpdate": [ "string" ],
      "blockedPhrasesToDelete": [ "string" ],
      "systemMessageOverride": "string"
   },
   "clientToken": "string",
   "creatorModeConfiguration": {
      "creatorModeControl": "string"
   },
   "hallucinationReductionConfiguration": {
      "hallucinationReductionControl": "string"
   },
   "orchestrationConfiguration": {
      "control": "string"
   },
   "responseScope": "string",
   "topicConfigurationsToCreateOrUpdate": [
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
   ],
   "topicConfigurationsToDelete": [
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

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_UpdateChatControlsConfiguration_RequestSyntax)**

The identifier of the application for which the chat controls are configured.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[blockedPhrasesConfigurationUpdate](#API_UpdateChatControlsConfiguration_RequestSyntax)**

The phrases blocked from chat by your chat control configuration.

Type: [BlockedPhrasesConfigurationUpdate](api-blockedphrasesconfigurationupdate.md) object

Required: No

**[clientToken](#API_UpdateChatControlsConfiguration_RequestSyntax)**

A token that you provide to identify the request to update a Amazon Q Business application
chat configuration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Required: No

**[creatorModeConfiguration](#API_UpdateChatControlsConfiguration_RequestSyntax)**

The configuration details for `CREATOR_MODE`.

Type: [CreatorModeConfiguration](api-creatormodeconfiguration.md) object

Required: No

**[hallucinationReductionConfiguration](#API_UpdateChatControlsConfiguration_RequestSyntax)**

The hallucination reduction settings for your application.

Type: [HallucinationReductionConfiguration](api-hallucinationreductionconfiguration.md) object

Required: No

**[orchestrationConfiguration](#API_UpdateChatControlsConfiguration_RequestSyntax)**

The chat response orchestration settings for your application.

Type: [OrchestrationConfiguration](api-orchestrationconfiguration.md) object

Required: No

**[responseScope](#API_UpdateChatControlsConfiguration_RequestSyntax)**

The response scope configured for your application. This determines whether your
application uses its retrieval augmented generation (RAG) system to generate answers
only from your enterprise data, or also uses the large language models (LLM) knowledge
to respons to end user questions in chat.

Type: String

Valid Values: `ENTERPRISE_CONTENT_ONLY | EXTENDED_KNOWLEDGE_ENABLED`

Required: No

**[topicConfigurationsToCreateOrUpdate](#API_UpdateChatControlsConfiguration_RequestSyntax)**

The configured topic specific chat controls you want to update.

Type: Array of [TopicConfiguration](api-topicconfiguration.md) objects

Array Members: Minimum number of 0 items. Maximum number of 10 items.

Required: No

**[topicConfigurationsToDelete](#API_UpdateChatControlsConfiguration_RequestSyntax)**

The configured topic specific chat controls you want to delete.

Type: Array of [TopicConfiguration](api-topicconfiguration.md) objects

Array Members: Minimum number of 0 items. Maximum number of 10 items.

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

**ServiceQuotaExceededException**

You have exceeded the set limits for your Amazon Q Business service.

**message**

The message describing a `ServiceQuotaExceededException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 402

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

- [AWS Command Line Interface V2](../../../goto/cli2/qbusiness-2023-11-27/updatechatcontrolsconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qbusiness-2023-11-27/updatechatcontrolsconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/updatechatcontrolsconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qbusiness-2023-11-27/updatechatcontrolsconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/updatechatcontrolsconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qbusiness-2023-11-27/updatechatcontrolsconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qbusiness-2023-11-27/updatechatcontrolsconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qbusiness-2023-11-27/updatechatcontrolsconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/qbusiness-2023-11-27/updatechatcontrolsconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/updatechatcontrolsconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateApplication

UpdateChatResponseConfiguration

All content copied from https://docs.aws.amazon.com/.
