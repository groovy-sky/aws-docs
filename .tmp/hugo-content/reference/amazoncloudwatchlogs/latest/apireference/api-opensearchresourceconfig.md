# OpenSearchResourceConfig

This structure contains configuration details about an integration between CloudWatch Logs and OpenSearch Service.

## Contents

**dashboardViewerPrincipals**

Specify the ARNs of IAM roles and IAM users who you want to
grant permission to for viewing the dashboards.

###### Important

In addition to specifying these users here, you must also grant them the **CloudWatchOpenSearchDashboardAccess**
IAM policy. For more information, see [IAM policies for\
users](../../../../services/amazoncloudwatch/latest/logs/opensearch-dashboards-userroles.md).

Type: Array of strings

Required: Yes

**dataSourceRoleArn**

Specify the ARN of an IAM role that CloudWatch Logs will use to create
the integration. This role must have the permissions necessary to access the OpenSearch Service collection to be able to create the dashboards. For more information about the permissions
needed, see [Permissions that\
the integration needs](../../../../services/amazoncloudwatch/latest/logs/opensearch-dashboards-createrole.md) in the CloudWatch Logs User Guide.

Type: String

Required: Yes

**retentionDays**

Specify how many days that you want the data derived by OpenSearch Service to be retained
in the index that the dashboard refers to. This also sets the maximum time period that you can
choose when viewing data in the dashboard. Choosing a longer time frame will incur additional
costs.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 30.

Required: Yes

**applicationArn**

If you want to use an existing OpenSearch Service application for your integration with
OpenSearch Service, specify it here. If you omit this, a new application will be
created.

Type: String

Required: No

**kmsKeyArn**

To have the vended dashboard data encrypted with AWS KMS instead of the CloudWatch Logs default encryption method, specify the ARN of the AWS KMS key that you
want to use.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/opensearchresourceconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/opensearchresourceconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/opensearchresourceconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenSearchNetworkPolicy

OpenSearchResourceStatus

All content copied from https://docs.aws.amazon.com/.
