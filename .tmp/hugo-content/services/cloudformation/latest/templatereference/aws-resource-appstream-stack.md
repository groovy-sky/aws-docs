This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::Stack

The `AWS::AppStream::Stack` resource creates a stack to start streaming applications to Amazon WorkSpaces Applications users. A stack consists of an associated fleet, user access policies, and storage configurations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppStream::Stack",
  "Properties" : {
      "AccessEndpoints" : [ AccessEndpoint, ... ],
      "ApplicationSettings" : ApplicationSettings,
      "AttributesToDelete" : [ String, ... ],
      "DeleteStorageConnectors" : Boolean,
      "Description" : String,
      "DisplayName" : String,
      "EmbedHostDomains" : [ String, ... ],
      "FeedbackURL" : String,
      "Name" : String,
      "RedirectURL" : String,
      "StorageConnectors" : [ StorageConnector, ... ],
      "StreamingExperienceSettings" : StreamingExperienceSettings,
      "Tags" : [ Tag, ... ],
      "UserSettings" : [ UserSetting, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppStream::Stack
Properties:
  AccessEndpoints:
    - AccessEndpoint
  ApplicationSettings:
    ApplicationSettings
  AttributesToDelete:
    - String
  DeleteStorageConnectors: Boolean
  Description: String
  DisplayName: String
  EmbedHostDomains:
    - String
  FeedbackURL: String
  Name: String
  RedirectURL: String
  StorageConnectors:
    - StorageConnector
  StreamingExperienceSettings:
    StreamingExperienceSettings
  Tags:
    - Tag
  UserSettings:
    - UserSetting

```

## Properties

`AccessEndpoints`

The list of virtual private cloud (VPC) interface endpoint objects. Users of the stack can connect to WorkSpaces Applications only through the specified endpoints.

_Required_: No

_Type_: Array of [AccessEndpoint](aws-properties-appstream-stack-accessendpoint.md)

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationSettings`

The persistent application settings for users of the stack. When these settings are enabled, changes that users make to applications and Windows settings are automatically saved after each session and applied to the next session.

_Required_: No

_Type_: [ApplicationSettings](aws-properties-appstream-stack-applicationsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AttributesToDelete`

The stack attributes to delete.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeleteStorageConnectors`

_This parameter has been deprecated._

Deletes the storage connectors currently enabled for the stack.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description to display.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The stack name to display.

_Required_: No

_Type_: String

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmbedHostDomains`

The domains where WorkSpaces Applications streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded WorkSpaces Applications streaming sessions.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FeedbackURL`

The URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed.

_Required_: No

_Type_: String

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the stack.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RedirectURL`

The URL that users are redirected to after their streaming session ends.

_Required_: No

_Type_: String

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageConnectors`

The storage connectors to enable.

_Required_: No

_Type_: Array of [StorageConnector](aws-properties-appstream-stack-storageconnector.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamingExperienceSettings`

The streaming protocol that you want your stack to prefer. This can be UDP or TCP. Currently, UDP is only supported in the Windows native client.

_Required_: No

_Type_: [StreamingExperienceSettings](aws-properties-appstream-stack-streamingexperiencesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs.

_Required_: No

_Type_: Array of [Tag](aws-properties-appstream-stack-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserSettings`

The actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled.

_Required_: No

_Type_: Array of [UserSetting](aws-properties-appstream-stack-usersetting.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [CreateStack](../../../../reference/appstream2/latest/apireference/api-createstack.md) in the _Amazon WorkSpaces Applications API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConfig

AccessEndpoint

All content copied from https://docs.aws.amazon.com/.
