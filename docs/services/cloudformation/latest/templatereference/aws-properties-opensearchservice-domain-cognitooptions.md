This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain CognitoOptions

Configures OpenSearch Service to use Amazon Cognito authentication for OpenSearch
Dashboards.

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

Whether to enable or disable Amazon Cognito authentication for OpenSearch Dashboards. See
[Amazon Cognito\
authentication for OpenSearch Dashboards](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityPoolId`

The Amazon Cognito identity pool ID that you want OpenSearch Service to use for OpenSearch
Dashboards authentication.

Required if you enabled Cognito Authentication for OpenSearch Dashboards.

_Required_: Conditional

_Type_: String

_Pattern_: `[\w-]+:[0-9a-f-]+`

_Minimum_: `1`

_Maximum_: `55`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The `AmazonOpenSearchServiceCognitoAccess` role that allows OpenSearch Service
to configure your user pool and identity pool.

Required if you enabled Cognito Authentication for OpenSearch Dashboards.

_Required_: Conditional

_Type_: String

_Pattern_: `arn:(aws|aws\-cn|aws\-us\-gov|aws\-iso|aws\-iso\-b):iam::[0-9]+:role\/.*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolId`

The Amazon Cognito user pool ID that you want OpenSearch Service to use for OpenSearch
Dashboards authentication.

Required if you enabled Cognito Authentication for OpenSearch Dashboards.

_Required_: Conditional

_Type_: String

_Pattern_: `[\w-]+_[0-9a-zA-Z]+`

_Minimum_: `1`

_Maximum_: `55`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ClusterConfig

ColdStorageOptions
