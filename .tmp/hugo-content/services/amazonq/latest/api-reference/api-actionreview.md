# ActionReview

An output event that Amazon Q Business returns to an user who wants to perform a plugin
action during a non-streaming chat conversation. It contains information about the
selected action with a list of possible user input fields, some pre-populated by
Amazon Q Business.

## Contents

**payload**

Field values that an end user needs to provide to Amazon Q Business for Amazon Q Business to
perform the requested plugin action.

Type: String to [ActionReviewPayloadField](api-actionreviewpayloadfield.md) object map

Key Length Constraints: Minimum length of 1.

Required: No

**payloadFieldNameSeparator**

A string used to retain information about the hierarchical contexts within an action
review payload.

Type: String

Length Constraints: Fixed length of 1.

Required: No

**pluginId**

The identifier of the plugin associated with the action review.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: No

**pluginType**

The type of plugin.

Type: String

Valid Values: `SERVICE_NOW | SALESFORCE | JIRA | ZENDESK | CUSTOM | QUICKSIGHT | SERVICENOW_NOW_PLATFORM | JIRA_CLOUD | SALESFORCE_CRM | ZENDESK_SUITE | ATLASSIAN_CONFLUENCE | GOOGLE_CALENDAR | MICROSOFT_TEAMS | MICROSOFT_EXCHANGE | PAGERDUTY_ADVANCE | SMARTSHEET | ASANA`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/actionreview.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/actionreview.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/actionreview.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActionFilterConfiguration

ActionReviewEvent

All content copied from https://docs.aws.amazon.com/.
