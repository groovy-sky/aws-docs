# UpdateStack

Updates the specified fields for the specified stack.

## Request Syntax

```nohighlight

{
   "AccessEndpoints": [
      {
         "EndpointType": "string",
         "VpceId": "string"
      }
   ],
   "ApplicationSettings": {
      "Enabled": boolean,
      "SettingsGroup": "string"
   },
   "AttributesToDelete": [ "string" ],
   "DeleteStorageConnectors": boolean,
   "Description": "string",
   "DisplayName": "string",
   "EmbedHostDomains": [ "string" ],
   "FeedbackURL": "string",
   "Name": "string",
   "RedirectURL": "string",
   "StorageConnectors": [
      {
         "ConnectorType": "string",
         "Domains": [ "string" ],
         "DomainsRequireAdminConsent": [ "string" ],
         "ResourceIdentifier": "string"
      }
   ],
   "StreamingExperienceSettings": {
      "PreferredProtocol": "string"
   },
   "UserSettings": [
      {
         "Action": "string",
         "MaximumLength": number,
         "Permission": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[AccessEndpoints](#API_UpdateStack_RequestSyntax)**

The list of interface VPC endpoint (interface endpoint) objects. Users of the stack can connect to WorkSpaces Applications only through the specified endpoints.

Type: Array of [AccessEndpoint](api-accessendpoint.md) objects

Array Members: Minimum number of 1 item. Maximum number of 4 items.

Required: No

**[ApplicationSettings](#API_UpdateStack_RequestSyntax)**

The persistent application settings for users of a stack. When these settings are enabled, changes that users make to applications and Windows settings are automatically saved after each session and applied to the next session.

Type: [ApplicationSettings](api-applicationsettings.md) object

Required: No

**[AttributesToDelete](#API_UpdateStack_RequestSyntax)**

The stack attributes to delete.

Type: Array of strings

Valid Values: `STORAGE_CONNECTORS | STORAGE_CONNECTOR_HOMEFOLDERS | STORAGE_CONNECTOR_GOOGLE_DRIVE | STORAGE_CONNECTOR_ONE_DRIVE | REDIRECT_URL | FEEDBACK_URL | THEME_NAME | USER_SETTINGS | EMBED_HOST_DOMAINS | IAM_ROLE_ARN | ACCESS_ENDPOINTS | STREAMING_EXPERIENCE_SETTINGS`

Required: No

**[DeleteStorageConnectors](#API_UpdateStack_RequestSyntax)**

_This parameter has been deprecated._

Deletes the storage connectors currently enabled for the stack.

Type: Boolean

Required: No

**[Description](#API_UpdateStack_RequestSyntax)**

The description to display.

Type: String

Length Constraints: Maximum length of 256.

Required: No

**[DisplayName](#API_UpdateStack_RequestSyntax)**

The stack name to display.

Type: String

Length Constraints: Maximum length of 100.

Required: No

**[EmbedHostDomains](#API_UpdateStack_RequestSyntax)**

The domains where WorkSpaces Applications streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded WorkSpaces Applications streaming sessions.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 20 items.

Length Constraints: Maximum length of 128.

Pattern: `(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]`

Required: No

**[FeedbackURL](#API_UpdateStack_RequestSyntax)**

The URL that users are redirected to after they choose the Send Feedback link. If no URL is specified, no Send Feedback link is displayed.

Type: String

Length Constraints: Maximum length of 1000.

Required: No

**[Name](#API_UpdateStack_RequestSyntax)**

The name of the stack.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**[RedirectURL](#API_UpdateStack_RequestSyntax)**

The URL that users are redirected to after their streaming session ends.

Type: String

Length Constraints: Maximum length of 1000.

Required: No

**[StorageConnectors](#API_UpdateStack_RequestSyntax)**

The storage connectors to enable.

Type: Array of [StorageConnector](api-storageconnector.md) objects

Required: No

**[StreamingExperienceSettings](#API_UpdateStack_RequestSyntax)**

The streaming protocol you want your stack to prefer. This can be UDP or TCP. Currently, UDP is only supported in the Windows native client.

Type: [StreamingExperienceSettings](api-streamingexperiencesettings.md) object

Required: No

**[UserSettings](#API_UpdateStack_RequestSyntax)**

The actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled.

Type: Array of [UserSetting](api-usersetting.md) objects

Array Members: Minimum number of 1 item.

Required: No

## Response Syntax

```nohighlight

{
   "Stack": {
      "AccessEndpoints": [
         {
            "EndpointType": "string",
            "VpceId": "string"
         }
      ],
      "ApplicationSettings": {
         "Enabled": boolean,
         "S3BucketName": "string",
         "SettingsGroup": "string"
      },
      "Arn": "string",
      "CreatedTime": number,
      "Description": "string",
      "DisplayName": "string",
      "EmbedHostDomains": [ "string" ],
      "FeedbackURL": "string",
      "Name": "string",
      "RedirectURL": "string",
      "StackErrors": [
         {
            "ErrorCode": "string",
            "ErrorMessage": "string"
         }
      ],
      "StorageConnectors": [
         {
            "ConnectorType": "string",
            "Domains": [ "string" ],
            "DomainsRequireAdminConsent": [ "string" ],
            "ResourceIdentifier": "string"
         }
      ],
      "StreamingExperienceSettings": {
         "PreferredProtocol": "string"
      },
      "UserSettings": [
         {
            "Action": "string",
            "MaximumLength": number,
            "Permission": "string"
         }
      ]
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Stack](#API_UpdateStack_ResponseSyntax)**

Information about the stack.

Type: [Stack](api-stack.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConcurrentModificationException**

An API error occurred. Wait a few minutes and try again.

**Message**

The error message in the exception.

HTTP Status Code: 400

**IncompatibleImageException**

The image can't be updated because it's not compatible for updates.

**Message**

The error message in the exception.

HTTP Status Code: 400

**InvalidAccountStatusException**

The resource cannot be created because your AWS account is suspended. For assistance, contact AWS Support.

**Message**

The error message in the exception.

HTTP Status Code: 400

**InvalidParameterCombinationException**

Indicates an incorrect combination of parameters, or a missing parameter.

**Message**

The error message in the exception.

HTTP Status Code: 400

**InvalidRoleException**

The specified role is invalid.

**Message**

The error message in the exception.

HTTP Status Code: 400

**LimitExceededException**

The requested limit exceeds the permitted limit for an account.

**Message**

The error message in the exception.

HTTP Status Code: 400

**OperationNotPermittedException**

The attempted operation is not permitted.

**Message**

The error message in the exception.

HTTP Status Code: 400

**ResourceInUseException**

The specified resource is in use.

**Message**

The error message in the exception.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource was not found.

**Message**

The error message in the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/updatestack.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/updatestack.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/updatestack.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/updatestack.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/updatestack.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/updatestack.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/updatestack.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/updatestack.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/updatestack.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/updatestack.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateImagePermissions

UpdateThemeForStack

All content copied from https://docs.aws.amazon.com/.
