# Stack

Describes a stack.

## Contents

**Name**

The name of the stack.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**AccessEndpoints**

The list of virtual private cloud (VPC) interface endpoint objects. Users of the stack can connect to WorkSpaces Applications only through the specified endpoints.

Type: Array of [AccessEndpoint](api-accessendpoint.md) objects

Array Members: Minimum number of 1 item. Maximum number of 4 items.

Required: No

**ApplicationSettings**

The persistent application settings for users of the stack.

Type: [ApplicationSettingsResponse](api-applicationsettingsresponse.md) object

Required: No

**Arn**

The ARN of the stack.

Type: String

Pattern: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

Required: No

**CreatedTime**

The time the stack was created.

Type: Timestamp

Required: No

**Description**

The description to display.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**DisplayName**

The stack name to display.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**EmbedHostDomains**

The domains where WorkSpaces Applications streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded WorkSpaces Applications streaming sessions.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 20 items.

Length Constraints: Maximum length of 128.

Pattern: `(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]`

Required: No

**FeedbackURL**

The URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed.

Type: String

Length Constraints: Maximum length of 1000.

Required: No

**RedirectURL**

The URL that users are redirected to after their streaming session ends.

Type: String

Length Constraints: Maximum length of 1000.

Required: No

**StackErrors**

The errors for the stack.

Type: Array of [StackError](api-stackerror.md) objects

Required: No

**StorageConnectors**

The storage connectors to enable.

Type: Array of [StorageConnector](api-storageconnector.md) objects

Required: No

**StreamingExperienceSettings**

The streaming protocol you want your stack to prefer. This can be UDP or TCP. Currently, UDP is only supported in the Windows native client.

Type: [StreamingExperienceSettings](api-streamingexperiencesettings.md) object

Required: No

**UserSettings**

The actions that are enabled or disabled for users during their streaming sessions. By default these actions are enabled.

Type: Array of [UserSetting](api-usersetting.md) objects

Array Members: Minimum number of 1 item.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/stack.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/stack.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/stack.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SoftwareAssociations

StackError

All content copied from https://docs.aws.amazon.com/.
