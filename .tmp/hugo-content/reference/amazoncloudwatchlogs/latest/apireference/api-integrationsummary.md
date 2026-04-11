# IntegrationSummary

This structure contains information about one CloudWatch Logs integration. This
structure is returned by a [ListIntegrations](api-listintegrations.md) operation.

## Contents

**integrationName**

The name of this integration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 50.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: No

**integrationStatus**

The current status of this integration.

Type: String

Valid Values: `PROVISIONING | ACTIVE | FAILED`

Required: No

**integrationType**

The type of integration. Integrations with OpenSearch Service have the type
`OPENSEARCH`.

Type: String

Valid Values: `OPENSEARCH`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/integrationsummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/integrationsummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/integrationsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntegrationDetails

ListToMap

All content copied from https://docs.aws.amazon.com/.
