# ConnectorDetail

Information about the registered connector.

## Contents

**applicationType**

The application type of the connector.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: No

**connectorDescription**

A description about the registered connector.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `[\w!@#\-.?,\s]*`

Required: No

**connectorLabel**

A label used for the connector.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[a-zA-Z0-9][\w!@#.-]+`

Required: No

**connectorModes**

The connection mode that the connector supports.

Type: Array of strings

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: No

**connectorName**

The name of the connector.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `.*`

Required: No

**connectorOwner**

The owner of the connector.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `.*`

Required: No

**connectorProvisioningType**

The provisioning type that the connector uses.

Type: String

Valid Values: `LAMBDA`

Required: No

**connectorType**

The connector type.

Type: String

Valid Values: `Salesforce | Singular | Slack | Redshift | S3 | Marketo | Googleanalytics | Zendesk | Servicenow | Datadog | Trendmicro | Snowflake | Dynatrace | Infornexus | Amplitude | Veeva | EventBridge | LookoutMetrics | Upsolver | Honeycode | CustomerProfiles | SAPOData | CustomConnector | Pardot`

Required: No

**connectorVersion**

The connector version.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: No

**registeredAt**

The time at which the connector was registered.

Type: Timestamp

Required: No

**registeredBy**

The user who registered the connector.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: No

**supportedDataTransferTypes**

The data transfer types that the connector supports.

RECORD

Structured records.

FILE

Files or binary data.

Type: Array of strings

Valid Values: `RECORD | FILE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/connectordetail.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/connectordetail.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/connectordetail.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectorConfiguration

ConnectorEntity

All content copied from https://docs.aws.amazon.com/.
