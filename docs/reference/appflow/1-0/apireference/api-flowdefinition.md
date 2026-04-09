# FlowDefinition

The properties of the flow, such as its source, destination, trigger type, and so on.

## Contents

**createdAt**

Specifies when the flow was created.

Type: Timestamp

Required: No

**createdBy**

The ARN of the user who created the flow.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: No

**description**

A user-entered description of the flow.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `[\w!@#\-.?,\s]*`

Required: No

**destinationConnectorLabel**

The label of the destination connector in the flow.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[a-zA-Z0-9][\w!@#.-]+`

Required: No

**destinationConnectorType**

Specifies the destination connector type, such as Salesforce, Amazon S3,
Amplitude, and so on.

Type: String

Valid Values: `Salesforce | Singular | Slack | Redshift | S3 | Marketo | Googleanalytics | Zendesk | Servicenow | Datadog | Trendmicro | Snowflake | Dynatrace | Infornexus | Amplitude | Veeva | EventBridge | LookoutMetrics | Upsolver | Honeycode | CustomerProfiles | SAPOData | CustomConnector | Pardot`

Required: No

**flowArn**

The flow's Amazon Resource Name (ARN).

Type: String

Length Constraints: Maximum length of 512.

Pattern: `arn:aws:appflow:.*:[0-9]+:.*`

Required: No

**flowName**

The specified name of the flow. Spaces are not allowed. Use underscores (\_) or hyphens
(-) only.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[a-zA-Z0-9][\w!@#.-]+`

Required: No

**flowStatus**

Indicates the current status of the flow.

Type: String

Valid Values: `Active | Deprecated | Deleted | Draft | Errored | Suspended`

Required: No

**lastRunExecutionDetails**

Describes the details of the most recent flow run.

Type: [ExecutionDetails](api-executiondetails.md) object

Required: No

**lastUpdatedAt**

Specifies when the flow was last updated.

Type: Timestamp

Required: No

**lastUpdatedBy**

Specifies the account user name that most recently updated the flow.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: No

**sourceConnectorLabel**

The label of the source connector in the flow.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[a-zA-Z0-9][\w!@#.-]+`

Required: No

**sourceConnectorType**

Specifies the source connector type, such as Salesforce, Amazon S3, Amplitude,
and so on.

Type: String

Valid Values: `Salesforce | Singular | Slack | Redshift | S3 | Marketo | Googleanalytics | Zendesk | Servicenow | Datadog | Trendmicro | Snowflake | Dynatrace | Infornexus | Amplitude | Veeva | EventBridge | LookoutMetrics | Upsolver | Honeycode | CustomerProfiles | SAPOData | CustomConnector | Pardot`

Required: No

**tags**

The tags used to organize, track, or control access for your flow.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^(?!aws:)[a-zA-Z+-=._:/]+$`

Value Length Constraints: Maximum length of 256.

Value Pattern: `[\s\w+-=\.:/@]*`

Required: No

**triggerType**

Specifies the type of flow trigger. This can be `OnDemand`,
`Scheduled`, or `Event`.

Type: String

Valid Values: `Scheduled | Event | OnDemand`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/flowdefinition.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/flowdefinition.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/flowdefinition.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldTypeDetails

GlueDataCatalogConfig

All content copied from https://docs.aws.amazon.com/.
