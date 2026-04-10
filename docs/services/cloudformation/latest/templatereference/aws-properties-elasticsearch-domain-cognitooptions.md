This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Elasticsearch::Domain CognitoOptions

Configures OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.

###### Important

The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](../userguide/aws-resource-opensearchservice-domain.md) resource. While the legacy Elasticsearch resource
and options are still supported, we recommend modifying your existing Cloudformation
templates to use the new OpenSearch Service resource, which supports both OpenSearch and
Elasticsearch. For more information about the service rename, see [New resource\
types](../../../opensearch-service/latest/developerguide/rename.md#rename-resource) in the _Amazon OpenSearch Service Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "IdentityPoolId" : String,
  "RoleArn" : String,
  "UserPoolId" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  IdentityPoolId: String
  RoleArn: String
  UserPoolId: String

```

## Properties

`Enabled`

Whether to enable or disable Amazon Cognito authentication for OpenSearch Dashboards. See [Amazon\
Cognito authentication for OpenSearch Dashboards](../../../opensearch-service/latest/developerguide/cognito-auth.md).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityPoolId`

The Amazon Cognito identity pool ID that you want OpenSearch Service to use for OpenSearch
Dashboards authentication. Required if you enable Cognito authentication.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The `AmazonESCognitoAccess` role that allows OpenSearch Service to configure
your user pool and identity pool. Required if you enable Cognito authentication.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolId`

The Amazon Cognito user pool ID that you want OpenSearch Service to use for OpenSearch
Dashboards authentication. Required if you enable Cognito authentication.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AdvancedSecurityOptionsInput

ColdStorageOptions

All content copied from https://docs.aws.amazon.com/.
