# ConnectorConfiguration

The configuration settings related to a given connector.

## Contents

**authenticationConfig**

The authentication config required for the connector.

Type: [AuthenticationConfig](api-authenticationconfig.md) object

Required: No

**canUseAsDestination**

Specifies whether the connector can be used as a destination.

Type: Boolean

Required: No

**canUseAsSource**

Specifies whether the connector can be used as a source.

Type: Boolean

Required: No

**connectorArn**

The Amazon Resource Name (ARN) for the registered connector.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `arn:aws:.*:.*:[0-9]+:.*`

Required: No

**connectorDescription**

A description about the connector.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `[\w!@#\-.?,\s]*`

Required: No

**connectorLabel**

The label used for registering the connector.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[a-zA-Z0-9][\w!@#.-]+`

Required: No

**connectorMetadata**

Specifies connector-specific metadata such as `oAuthScopes`,
`supportedRegions`, `privateLinkServiceUrl`, and so on.

Type: [ConnectorMetadata](api-connectormetadata.md) object

Required: No

**connectorModes**

The connection modes that the connector supports.

Type: Array of strings

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: No

**connectorName**

The connector name.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `.*`

Required: No

**connectorOwner**

The owner who developed the connector.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `.*`

Required: No

**connectorProvisioningConfig**

The configuration required for registering the connector.

Type: [ConnectorProvisioningConfig](api-connectorprovisioningconfig.md) object

Required: No

**connectorProvisioningType**

The provisioning type used to register the connector.

Type: String

Valid Values: `LAMBDA`

Required: No

**connectorRuntimeSettings**

The required connector runtime settings.

Type: Array of [ConnectorRuntimeSetting](api-connectorruntimesetting.md) objects

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

**isPrivateLinkEnabled**

Specifies if PrivateLink is enabled for that connector.

Type: Boolean

Required: No

**isPrivateLinkEndpointUrlRequired**

Specifies if a PrivateLink endpoint URL is required.

Type: Boolean

Required: No

**logoURL**

Logo URL of the connector.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `^(https?|ftp|file)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]`

Required: No

**registeredAt**

The date on which the connector was registered.

Type: Timestamp

Required: No

**registeredBy**

Information about who registered the connector.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: No

**supportedApiVersions**

A list of API versions that are supported by the connector.

Type: Array of strings

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: No

**supportedDataTransferApis**

The APIs of the connector application that Amazon AppFlow can use to transfer your
data.

Type: Array of [DataTransferApi](api-datatransferapi.md) objects

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

**supportedDestinationConnectors**

Lists the connectors that are available for use as destinations.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 100 items.

Valid Values: `Salesforce | Singular | Slack | Redshift | S3 | Marketo | Googleanalytics | Zendesk | Servicenow | Datadog | Trendmicro | Snowflake | Dynatrace | Infornexus | Amplitude | Veeva | EventBridge | LookoutMetrics | Upsolver | Honeycode | CustomerProfiles | SAPOData | CustomConnector | Pardot`

Required: No

**supportedOperators**

A list of operators supported by the connector.

Type: Array of strings

Valid Values: `PROJECTION | LESS_THAN | GREATER_THAN | CONTAINS | BETWEEN | LESS_THAN_OR_EQUAL_TO | GREATER_THAN_OR_EQUAL_TO | EQUAL_TO | NOT_EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

Required: No

**supportedSchedulingFrequencies**

Specifies the supported flow frequency for that connector.

Type: Array of strings

Valid Values: `BYMINUTE | HOURLY | DAILY | WEEKLY | MONTHLY | ONCE`

Required: No

**supportedTriggerTypes**

Specifies the supported trigger types for the flow.

Type: Array of strings

Valid Values: `Scheduled | Event | OnDemand`

Required: No

**supportedWriteOperations**

A list of write operations supported by the connector.

Type: Array of strings

Valid Values: `INSERT | UPSERT | UPDATE | DELETE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/connectorconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/connectorconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/connectorconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BasicAuthCredentials

ConnectorDetail

All content copied from https://docs.aws.amazon.com/.
