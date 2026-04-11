# ChatSync

Starts or continues a non-streaming Amazon Q Business conversation.

## Request Syntax

```nohighlight

POST /applications/applicationId/conversations?sync&userGroups=userGroups&userId=userId HTTP/1.1
Content-type: application/json

{
   "actionExecution": {
      "payload": {
         "string" : {
            "value": JSON value
         }
      },
      "payloadFieldNameSeparator": "string",
      "pluginId": "string"
   },
   "attachments": [
      {
         "copyFrom": { ... },
         "data": blob,
         "name": "string"
      }
   ],
   "attributeFilter": {
      "andAllFilters": [
         "AttributeFilter"
      ],
      "containsAll": {
         "name": "string",
         "value": { ... }
      },
      "containsAny": {
         "name": "string",
         "value": { ... }
      },
      "equalsTo": {
         "name": "string",
         "value": { ... }
      },
      "greaterThan": {
         "name": "string",
         "value": { ... }
      },
      "greaterThanOrEquals": {
         "name": "string",
         "value": { ... }
      },
      "lessThan": {
         "name": "string",
         "value": { ... }
      },
      "lessThanOrEquals": {
         "name": "string",
         "value": { ... }
      },
      "notFilter": "AttributeFilter",
      "orAllFilters": [
         "AttributeFilter"
      ]
   },
   "authChallengeResponse": {
      "responseMap": {
         "string" : "string"
      }
   },
   "chatMode": "string",
   "chatModeConfiguration": { ... },
   "clientToken": "string",
   "conversationId": "string",
   "parentMessageId": "string",
   "userMessage": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_ChatSync_RequestSyntax)**

The identifier of the Amazon Q Business application linked to the Amazon Q Business
conversation.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[userGroups](#API_ChatSync_RequestSyntax)**

The group names that a user associated with the chat input belongs to.

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[userId](#API_ChatSync_RequestSyntax)**

The identifier of the user attached to the chat input.

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `\P{C}*`

## Request Body

The request accepts the following data in JSON format.

**[actionExecution](#API_ChatSync_RequestSyntax)**

A request from an end user to perform an Amazon Q Business plugin action.

Type: [ActionExecution](api-actionexecution.md) object

Required: No

**[attachments](#API_ChatSync_RequestSyntax)**

A list of files uploaded directly during chat. You can upload a maximum of 5 files of
upto 10 MB each.

Type: Array of [AttachmentInput](api-attachmentinput.md) objects

Array Members: Minimum number of 1 item.

Required: No

**[attributeFilter](#API_ChatSync_RequestSyntax)**

Enables filtering of Amazon Q Business web experience responses based on document
attributes or metadata fields.

Type: [AttributeFilter](api-attributefilter.md) object

Required: No

**[authChallengeResponse](#API_ChatSync_RequestSyntax)**

An authentication verification event response by a third party authentication server
to Amazon Q Business.

Type: [AuthChallengeResponse](api-authchallengeresponse.md) object

Required: No

**[chatMode](#API_ChatSync_RequestSyntax)**

The `chatMode` parameter determines the chat modes available to
Amazon Q Business users:

- `RETRIEVAL_MODE` \- If you choose this mode, Amazon Q generates responses solely from the data sources connected and
indexed by the application. If an answer is not found in the data sources or
there are no data sources available, Amazon Q will respond with a
" _No Answer Found_" message, unless LLM knowledge has
been enabled. In that case, Amazon Q will generate a response from
the LLM knowledge

- `CREATOR_MODE` \- By selecting this mode, you can choose to generate
responses only from the LLM knowledge. You can also attach files and have Amazon Q
generate a response based on the data in those files.
If the attached files do not contain an answer for the query, Amazon Q
will automatically fall back to generating a response from the LLM knowledge.

- `PLUGIN_MODE` \- By selecting this mode, users can choose to
use plugins in chat to get their responses.

###### Note

If none of the modes are selected, Amazon Q will only respond using the information
from the attached files.

For more information, see [Admin controls and guardrails](../qbusiness-ug/guardrails.md), [Plugins](../qbusiness-ug/plugins.md),
and [Response sources](../business-use-dg/using-web-experience.md#chat-source-scope).

Type: String

Valid Values: `RETRIEVAL_MODE | CREATOR_MODE | PLUGIN_MODE`

Required: No

**[chatModeConfiguration](#API_ChatSync_RequestSyntax)**

The chat mode configuration for an Amazon Q Business application.

Type: [ChatModeConfiguration](api-chatmodeconfiguration.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**[clientToken](#API_ChatSync_RequestSyntax)**

A token that you provide to identify a chat request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Required: No

**[conversationId](#API_ChatSync_RequestSyntax)**

The identifier of the Amazon Q Business conversation.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**[parentMessageId](#API_ChatSync_RequestSyntax)**

The identifier of the previous system message in a conversation.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**[userMessage](#API_ChatSync_RequestSyntax)**

A end user message in a conversation.

Type: String

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "actionReview": {
      "payload": {
         "string" : {
            "allowedFormat": "string",
            "allowedValues": [
               {
                  "displayValue": JSON value,
                  "value": JSON value
               }
            ],
            "arrayItemJsonSchema": JSON value,
            "displayDescription": "string",
            "displayName": "string",
            "displayOrder": number,
            "required": boolean,
            "type": "string",
            "value": JSON value
         }
      },
      "payloadFieldNameSeparator": "string",
      "pluginId": "string",
      "pluginType": "string"
   },
   "authChallengeRequest": {
      "authorizationUrl": "string"
   },
   "conversationId": "string",
   "failedAttachments": [
      {
         "attachmentId": "string",
         "conversationId": "string",
         "error": {
            "errorCode": "string",
            "errorMessage": "string"
         },
         "name": "string",
         "status": "string"
      }
   ],
   "sourceAttributions": [
      {
         "citationNumber": number,
         "datasourceId": "string",
         "documentId": "string",
         "indexId": "string",
         "snippet": "string",
         "textMessageSegments": [
            {
               "beginOffset": number,
               "endOffset": number,
               "mediaId": "string",
               "mediaMimeType": "string",
               "snippetExcerpt": {
                  "text": "string"
               },
               "sourceDetails": { ... }
            }
         ],
         "title": "string",
         "updatedAt": number,
         "url": "string"
      }
   ],
   "systemMessage": "string",
   "systemMessageId": "string",
   "userMessageId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[actionReview](#API_ChatSync_ResponseSyntax)**

A request from Amazon Q Business to the end user for information Amazon Q Business needs to
successfully complete a requested plugin action.

Type: [ActionReview](api-actionreview.md) object

**[authChallengeRequest](#API_ChatSync_ResponseSyntax)**

An authentication verification event activated by an end user request to use a custom
plugin.

Type: [AuthChallengeRequest](api-authchallengerequest.md) object

**[conversationId](#API_ChatSync_ResponseSyntax)**

The identifier of the Amazon Q Business conversation.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

**[failedAttachments](#API_ChatSync_ResponseSyntax)**

A list of files which failed to upload during chat.

Type: Array of [AttachmentOutput](api-attachmentoutput.md) objects

**[sourceAttributions](#API_ChatSync_ResponseSyntax)**

The source documents used to generate the conversation response.

Type: Array of [SourceAttribution](api-sourceattribution.md) objects

**[systemMessage](#API_ChatSync_ResponseSyntax)**

An AI-generated message in a conversation.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[systemMessageId](#API_ChatSync_ResponseSyntax)**

The identifier of an Amazon Q Business AI generated message within the
conversation.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

**[userMessageId](#API_ChatSync_ResponseSyntax)**

The identifier of an Amazon Q Business end user text input message within the
conversation.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

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

**ExternalResourceException**

An external resource that you configured with your application is returning errors and
preventing this operation from succeeding. Fix those errors and try again.

HTTP Status Code: 424

**InternalServerException**

An issue occurred with the internal server used for your Amazon Q Business service. Wait
some minutes and try again, or contact [Support](http://aws.amazon.com/contact-us) for help.

HTTP Status Code: 500

**LicenseNotFoundException**

You don't have permissions to perform the action because your license is inactive. Ask
your admin to activate your license and try again after your licence is active.

HTTP Status Code: 400

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

- [AWS Command Line Interface V2](../../../goto/cli2/qbusiness-2023-11-27/chatsync.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qbusiness-2023-11-27/chatsync.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/chatsync.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qbusiness-2023-11-27/chatsync.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/chatsync.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qbusiness-2023-11-27/chatsync.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qbusiness-2023-11-27/chatsync.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qbusiness-2023-11-27/chatsync.md)

- [AWS SDK for Python](../../../goto/boto3/qbusiness-2023-11-27/chatsync.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/chatsync.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Chat

CheckDocumentAccess

All content copied from https://docs.aws.amazon.com/.
