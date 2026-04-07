This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Elasticsearch::Domain AdvancedSecurityOptionsInput

Specifies options for fine-grained access control.

###### Important

The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource
and options are still supported, we recommend modifying your existing Cloudformation
templates to use the new OpenSearch Service resource, which supports both OpenSearch and
Elasticsearch. For more information about the service rename, see [New resource\
types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the _Amazon OpenSearch Service Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnonymousAuthEnabled" : Boolean,
  "Enabled" : Boolean,
  "InternalUserDatabaseEnabled" : Boolean,
  "MasterUserOptions" : MasterUserOptions
}

```

### YAML

```yaml

  AnonymousAuthEnabled: Boolean
  Enabled: Boolean
  InternalUserDatabaseEnabled: Boolean
  MasterUserOptions:
    MasterUserOptions

```

## Properties

`AnonymousAuthEnabled`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

True to enable fine-grained access control. You must also enable encryption of data at rest
and node-to-node encryption.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InternalUserDatabaseEnabled`

True to enable the internal user database.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUserOptions`

Specifies information about the master user.

_Required_: No

_Type_: [MasterUserOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticsearch-domain-masteruseroptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Elasticsearch::Domain

CognitoOptions
