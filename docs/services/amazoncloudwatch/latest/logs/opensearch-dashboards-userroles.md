# IAM policies for users

CloudWatch Logs has created two IAM policies,
**CloudWatchOpenSearchDashboardsFullAccess** and
**CloudWatchOpenSearchDashboardAccess**.

The following table lists which actions each of these policies enables.

ActionIAM policyAdditional permissions needed

Create integration

**CloudWatchOpenSearchDashboardsFullAccess**

Delete integration

**CloudWatchOpenSearchDashboardsFullAccess**

Create dashboard

**CloudWatchOpenSearchDashboardsFullAccess**

Edit dashboard

**CloudWatchOpenSearchDashboardsFullAccess**

Delete dashboard

**CloudWatchOpenSearchDashboardsFullAccess**

Refresh dashboard using **Synchronize**
**now**

**CloudWatchOpenSearchDashboardsFullAccess**

View integration in **Settings**

**CloudWatchOpenSearchDashboardAccess** or
**CloudWatchOpenSearchDashboardsFullAccess**

View dashboard

**CloudWatchOpenSearchDashboardAccess** or
**CloudWatchOpenSearchDashboardsFullAccess**

Specify the role or user when you create the integration, or
edit the data access policy for the collection to add these roles or
users. For more information, see [Data access control for Amazon OpenSearch Service Serverless](../../../opensearch-service/latest/developerguide/serverless-data-access.md) in the
OpenSearch Service Developer Guide.

View dashboard in OpenSearch Service console

**CloudWatchOpenSearchDashboardAccess** or
**CloudWatchOpenSearchDashboardsFullAccess**

Specify the role or user when you create the integration, or
edit the data access policy for the collection to add these roles or
users. For more information, see [Data access control for Amazon OpenSearch Service Serverless](../../../opensearch-service/latest/developerguide/serverless-data-access.md) in the
OpenSearch Service Developer Guide.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

View, edit, or delete vended logs dashboards

Permissions that the integration needs

All content copied from https://docs.aws.amazon.com/.
