This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Plugin

Information about an Amazon Q Business plugin and its configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QBusiness::Plugin",
  "Properties" : {
      "ApplicationId" : String,
      "AuthConfiguration" : PluginAuthConfiguration,
      "CustomPluginConfiguration" : CustomPluginConfiguration,
      "DisplayName" : String,
      "ServerUrl" : String,
      "State" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::QBusiness::Plugin
Properties:
  ApplicationId: String
  AuthConfiguration:
    PluginAuthConfiguration
  CustomPluginConfiguration:
    CustomPluginConfiguration
  DisplayName: String
  ServerUrl: String
  State: String
  Tags:
    - Tag
  Type: String

```

## Properties

`ApplicationId`

The identifier of the application that will contain the plugin.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9-]{35}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AuthConfiguration`

Authentication configuration information for an Amazon Q Business plugin.

_Required_: Yes

_Type_: [PluginAuthConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-plugin-pluginauthconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomPluginConfiguration`

Configuration information required to create a custom plugin.

_Required_: No

_Type_: [CustomPluginConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-plugin-custompluginconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The name of the plugin.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerUrl`

The plugin server URL used for configuration.

_Required_: No

_Type_: String

_Pattern_: `^(https?|ftp|file)://([^\s]*)$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

The current status of the plugin.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs that identify or categorize the data source connector. You
can also use tags to help control access to the data source connector. Tag keys and
values can consist of Unicode letters, digits, white space, and any of the following
symbols: \_ . : / = + - @.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-plugin-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the plugin.

_Required_: Yes

_Type_: String

_Allowed values_: `SERVICE_NOW | SALESFORCE | JIRA | ZENDESK | CUSTOM | QUICKSIGHT | SERVICENOW_NOW_PLATFORM | JIRA_CLOUD | SALESFORCE_CRM | ZENDESK_SUITE | ATLASSIAN_CONFLUENCE | GOOGLE_CALENDAR | MICROSOFT_TEAMS | MICROSOFT_EXCHANGE | PAGERDUTY_ADVANCE | SMARTSHEET | ASANA`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the application ID and plugin ID. For example:

`{"Ref": "ApplicationId|PluginId"}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`BuildStatus`

The current status of a plugin. A plugin is modified asynchronously.

`CreatedAt`

The timestamp for when the plugin was created.

`PluginArn`

The Amazon Resource Name (ARN) of a plugin.

`PluginId`

The identifier of the plugin.

`UpdatedAt`

The timestamp for when the plugin was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Condition

APISchema
