# OpenSearchApplication

This structure contains information about the OpenSearch Service application used for this
integration. An OpenSearch Service application is the web application created by the
integration with CloudWatch Logs. It hosts the vended logs dashboards.

## Contents

**applicationArn**

The Amazon Resource Name (ARN) of the application.

Type: String

Required: No

**applicationEndpoint**

The endpoint of the application.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `^https://[\.\-_/#:A-Za-z0-9]+\.com$`

Required: No

**applicationId**

The ID of the application.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: No

**status**

This structure contains information about the status of this OpenSearch Service
resource.

Type: [OpenSearchResourceStatus](api-opensearchresourcestatus.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/opensearchapplication.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/opensearchapplication.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/opensearchapplication.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MoveKeys

OpenSearchCollection

All content copied from https://docs.aws.amazon.com/.
