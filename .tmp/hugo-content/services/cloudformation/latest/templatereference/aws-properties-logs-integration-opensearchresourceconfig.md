This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Integration OpenSearchResourceConfig

This structure contains configuration details about an integration between CloudWatch
Logs and OpenSearch Service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationARN" : String,
  "DashboardViewerPrincipals" : [ String, ... ],
  "DataSourceRoleArn" : String,
  "KmsKeyArn" : String,
  "RetentionDays" : Integer
}

```

### YAML

```yaml

  ApplicationARN: String
  DashboardViewerPrincipals:
    - String
  DataSourceRoleArn: String
  KmsKeyArn: String
  RetentionDays: Integer

```

## Properties

`ApplicationARN`

If you want to use an existing OpenSearch Service application for your integration
with OpenSearch Service, specify it here. If you omit this, a new application will be
created.

_Required_: No

_Type_: String

_Pattern_: `[\w#+=/:,.@-]*\*?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DashboardViewerPrincipals`

Specify the ARNs of IAM roles and IAM users who you want to
grant permission to for viewing the dashboards.

###### Important

In addition to specifying these users here, you must also grant them the **CloudWatchOpenSearchDashboardAccess**
IAM policy. For more information, see [IAM policies for\
users](../../../amazoncloudwatch/latest/logs/opensearch-dashboards-userroles.md).

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataSourceRoleArn`

Specify the ARN of an IAM role that CloudWatch Logs will use to create the
integration. This role must have the permissions necessary to access the OpenSearch
Service collection to be able to create the dashboards. For more information about the
permissions needed, see [Permissions\
that the integration needs](../../../amazoncloudwatch/latest/logs/opensearch-dashboards-createrole.md) in the CloudWatch Logs User Guide.

_Required_: Yes

_Type_: String

_Pattern_: `[\w#+=/:,.@-]*\*?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyArn`

To have the vended dashboard data encrypted with AWS KMS instead of the
CloudWatch Logs default encryption method, specify the ARN of the AWS KMS key
that you want to use.

_Required_: No

_Type_: String

_Pattern_: `[\w#+=/:,.@-]*\*?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RetentionDays`

Specify how many days that you want the data derived by OpenSearch Service to be
retained in the index that the dashboard refers to. This also sets the maximum time
period that you can choose when viewing data in the dashboard. Choosing a longer time
frame will incur additional costs.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `3650`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Logs::Integration

ResourceConfig

All content copied from https://docs.aws.amazon.com/.
