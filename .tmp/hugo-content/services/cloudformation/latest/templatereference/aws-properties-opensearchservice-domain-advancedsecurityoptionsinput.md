This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain AdvancedSecurityOptionsInput

Specifies options for fine-grained access control.

If you specify advanced security options,
you must also enable node-to-node encryption ( [NodeToNodeEncryptionOptions](../userguide/aws-properties-opensearchservice-domain-nodetonodeencryptionoptions.md)) and encryption at rest ( [EncryptionAtRestOptions](../userguide/aws-properties-opensearchservice-domain-encryptionatrestoptions.md)). You must also enable `EnforceHTTPS` within
[DomainEndpointOptions](../userguide/aws-properties-opensearchservice-domain-domainendpointoptions.md), which requires HTTPS for all traffic to the domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnonymousAuthDisableDate" : String,
  "AnonymousAuthEnabled" : Boolean,
  "Enabled" : Boolean,
  "IAMFederationOptions" : IAMFederationOptions,
  "InternalUserDatabaseEnabled" : Boolean,
  "JWTOptions" : JWTOptions,
  "MasterUserOptions" : MasterUserOptions,
  "SAMLOptions" : SAMLOptions
}

```

### YAML

```yaml

  AnonymousAuthDisableDate: String
  AnonymousAuthEnabled: Boolean
  Enabled: Boolean
  IAMFederationOptions:
    IAMFederationOptions
  InternalUserDatabaseEnabled: Boolean
  JWTOptions:
    JWTOptions
  MasterUserOptions:
    MasterUserOptions
  SAMLOptions:
    SAMLOptions

```

## Properties

`AnonymousAuthDisableDate`

Date and time when the migration period will be disabled. Only necessary when [enabling\
fine-grained access control on an existing domain](../../../opensearch-service/latest/developerguide/fgac.md#fgac-enabling-existing).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AnonymousAuthEnabled`

True to enable a 30-day migration period during which administrators can create role
mappings. Only necessary when [enabling fine-grained access control on an existing domain](../../../opensearch-service/latest/developerguide/fgac.md#fgac-enabling-existing).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

True to enable fine-grained access control. You must also enable encryption of data at rest
and node-to-node encryption. See [Fine-grained access control in\
Amazon OpenSearch Service](../../../opensearch-service/latest/developerguide/fgac.md).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IAMFederationOptions`

Input configuration for IAM identity federation within advanced security
options.

_Required_: No

_Type_: [IAMFederationOptions](aws-properties-opensearchservice-domain-iamfederationoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InternalUserDatabaseEnabled`

True to enable the internal user database.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JWTOptions`

Container for information about the JWT configuration of the Amazon OpenSearch
Service.

_Required_: No

_Type_: [JWTOptions](aws-properties-opensearchservice-domain-jwtoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUserOptions`

Specifies information about the master user.

_Required_: No

_Type_: [MasterUserOptions](aws-properties-opensearchservice-domain-masteruseroptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SAMLOptions`

Container for information about the SAML configuration for OpenSearch
Dashboards.

_Required_: No

_Type_: [SAMLOptions](aws-properties-opensearchservice-domain-samloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::OpenSearchService::Domain

AIMLOptions

All content copied from https://docs.aws.amazon.com/.
