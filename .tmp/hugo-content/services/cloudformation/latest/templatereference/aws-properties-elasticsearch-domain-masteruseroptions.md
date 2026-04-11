This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Elasticsearch::Domain MasterUserOptions

Specifies information about the master user. Required if you enabled the internal user
database.

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
  "MasterUserARN" : String,
  "MasterUserName" : String,
  "MasterUserPassword" : String
}

```

### YAML

```yaml

  MasterUserARN: String
  MasterUserName: String
  MasterUserPassword: String

```

## Properties

`MasterUserARN`

ARN for the master user. Only specify if `InternalUserDatabaseEnabled` is false in `AdvancedSecurityOptions`.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUserName`

Username for the master user. Only specify if `InternalUserDatabaseEnabled` is true in `AdvancedSecurityOptions`.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUserPassword`

Password for the master user. Only specify if `InternalUserDatabaseEnabled` is true in `AdvancedSecurityOptions`.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogPublishingOption

NodeToNodeEncryptionOptions

All content copied from https://docs.aws.amazon.com/.
