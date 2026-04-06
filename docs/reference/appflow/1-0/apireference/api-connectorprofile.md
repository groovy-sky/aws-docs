# ConnectorProfile

Describes an instance of a connector. This includes the provided name, credentials ARN,
connection-mode, and so on. To keep the API intuitive and extensible, the fields that are
common to all types of connector profiles are explicitly specified at the top level. The rest
of the connector-specific properties are available via the
`connectorProfileProperties` field.

## Contents

**connectionMode**

Indicates the connection mode and if it is public or private.

Type: String

Valid Values: `Public | Private`

Required: No

**connectorLabel**

The label for the connector profile being created.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[a-zA-Z0-9][\w!@#.-]+`

Required: No

**connectorProfileArn**

The Amazon Resource Name (ARN) of the connector profile.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `arn:aws:appflow:.*:[0-9]+:.*`

Required: No

**connectorProfileName**

The name of the connector profile. The name is unique for each
`ConnectorProfile` in the AWS account.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[\w/!@#+=.-]+`

Required: No

**connectorProfileProperties**

The connector-specific properties of the profile configuration.

Type: [ConnectorProfileProperties](api-connectorprofileproperties.md) object

Required: No

**connectorType**

The type of connector, such as Salesforce, Amplitude, and so on.

Type: String

Valid Values: `Salesforce | Singular | Slack | Redshift | S3 | Marketo | Googleanalytics | Zendesk | Servicenow | Datadog | Trendmicro | Snowflake | Dynatrace | Infornexus | Amplitude | Veeva | EventBridge | LookoutMetrics | Upsolver | Honeycode | CustomerProfiles | SAPOData | CustomConnector | Pardot`

Required: No

**createdAt**

Specifies when the connector profile was created.

Type: Timestamp

Required: No

**credentialsArn**

The Amazon Resource Name (ARN) of the connector profile credentials.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `arn:aws:.*:.*:[0-9]+:.*`

Required: No

**lastUpdatedAt**

Specifies when the connector profile was last updated.

Type: Timestamp

Required: No

**privateConnectionProvisioningState**

Specifies the private connection provisioning state.

Type: [PrivateConnectionProvisioningState](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_PrivateConnectionProvisioningState.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/ConnectorProfile)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/ConnectorProfile)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/ConnectorProfile)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConnectorOperator

ConnectorProfileConfig
