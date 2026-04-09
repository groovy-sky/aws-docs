# OpenSearchWorkspace

This structure contains information about the OpenSearch Service workspace used for this
integration. An OpenSearch Service workspace is the collection of dashboards along with other
OpenSearch Service tools. This workspace was created automatically as part of the
integration setup. For more information, see [Centralized OpenSearch user\
interface (Dashboards) with OpenSearch Service](../../../../services/opensearch-service/latest/developerguide/application.md).

## Contents

**status**

This structure contains information about the status of an OpenSearch Service
resource.

Type: [OpenSearchResourceStatus](api-opensearchresourcestatus.md) object

Required: No

**workspaceId**

The ID of this workspace.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/opensearchworkspace.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/opensearchworkspace.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/opensearchworkspace.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenSearchResourceStatus

OutputLogEvent

All content copied from https://docs.aws.amazon.com/.
