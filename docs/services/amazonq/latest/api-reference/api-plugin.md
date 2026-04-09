# Plugin

Information about an Amazon Q Business plugin and its configuration.

## Contents

**buildStatus**

The status of the plugin.

Type: String

Valid Values: `READY | CREATE_IN_PROGRESS | CREATE_FAILED | UPDATE_IN_PROGRESS | UPDATE_FAILED | DELETE_IN_PROGRESS | DELETE_FAILED`

Required: No

**createdAt**

The timestamp for when the plugin was created.

Type: Timestamp

Required: No

**displayName**

The name of the plugin.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`

Required: No

**pluginId**

The identifier of the plugin.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: No

**serverUrl**

The plugin server URL used for configuration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `(https?|ftp|file)://([^\s]*)`

Required: No

**state**

The current status of the plugin.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**type**

The type of the plugin.

Type: String

Valid Values: `SERVICE_NOW | SALESFORCE | JIRA | ZENDESK | CUSTOM | QUICKSIGHT | SERVICENOW_NOW_PLATFORM | JIRA_CLOUD | SALESFORCE_CRM | ZENDESK_SUITE | ATLASSIAN_CONFLUENCE | GOOGLE_CALENDAR | MICROSOFT_TEAMS | MICROSOFT_EXCHANGE | PAGERDUTY_ADVANCE | SMARTSHEET | ASANA`

Required: No

**updatedAt**

The timestamp for when the plugin was last updated.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/plugin.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/plugin.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/plugin.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PersonalizationConfiguration

PluginAuthConfiguration

All content copied from https://docs.aws.amazon.com/.
