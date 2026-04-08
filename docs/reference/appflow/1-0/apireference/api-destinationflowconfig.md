# DestinationFlowConfig

Contains information about the configuration of destination connectors present in the
flow.

## Contents

**connectorType**

The type of connector, such as Salesforce, Amplitude, and so on.

Type: String

Valid Values: `Salesforce | Singular | Slack | Redshift | S3 | Marketo | Googleanalytics | Zendesk | Servicenow | Datadog | Trendmicro | Snowflake | Dynatrace | Infornexus | Amplitude | Veeva | EventBridge | LookoutMetrics | Upsolver | Honeycode | CustomerProfiles | SAPOData | CustomConnector | Pardot`

Required: Yes

**destinationConnectorProperties**

This stores the information that is required to query a particular connector.

Type: [DestinationConnectorProperties](api-destinationconnectorproperties.md) object

Required: Yes

**apiVersion**

The API version that the destination connector uses.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: No

**connectorProfileName**

The name of the connector profile. This name must be unique for each connector profile in
the AWS account.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[\w/!@#+=.-]+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/destinationflowconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/destinationflowconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/destinationflowconfig.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DestinationFieldProperties

DynatraceConnectorProfileCredentials
