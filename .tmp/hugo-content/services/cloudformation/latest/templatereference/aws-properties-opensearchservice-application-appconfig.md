This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Application AppConfig

Configuration settings for an OpenSearch application. For more information, see [Using the\
OpenSearch user interface in Amazon OpenSearch Service](../../../opensearch-service/latest/developerguide/application.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The configuration item to set, such as the admin role for the OpenSearch
application.

_Required_: Yes

_Type_: String

_Allowed values_: `opensearchDashboards.dashboardAdmin.users | opensearchDashboards.dashboardAdmin.groups`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value assigned to the configuration key, such as an IAM user ARN.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::OpenSearchService::Application

DataSource

All content copied from https://docs.aws.amazon.com/.
