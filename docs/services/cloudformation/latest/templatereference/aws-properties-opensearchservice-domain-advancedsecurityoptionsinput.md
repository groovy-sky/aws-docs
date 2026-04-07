This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain AdvancedSecurityOptionsInput

Specifies options for fine-grained access control.

If you specify advanced security options,
you must also enable node-to-node encryption ( [NodeToNodeEncryptionOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-nodetonodeencryptionoptions.html)) and encryption at rest ( [EncryptionAtRestOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-encryptionatrestoptions.html)). You must also enable `EnforceHTTPS` within
[DomainEndpointOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html), which requires HTTPS for all traffic to the domain.

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
fine-grained access control on an existing domain](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html#fgac-enabling-existing).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AnonymousAuthEnabled`

True to enable a 30-day migration period during which administrators can create role
mappings. Only necessary when [enabling fine-grained access control on an existing domain](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html#fgac-enabling-existing).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

True to enable fine-grained access control. You must also enable encryption of data at rest
and node-to-node encryption. See [Fine-grained access control in\
Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IAMFederationOptions`

Input configuration for IAM identity federation within advanced security
options.

_Required_: No

_Type_: [IAMFederationOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-opensearchservice-domain-iamfederationoptions.html)

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

_Type_: [JWTOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-opensearchservice-domain-jwtoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUserOptions`

Specifies information about the master user.

_Required_: No

_Type_: [MasterUserOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-opensearchservice-domain-masteruseroptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SAMLOptions`

Container for information about the SAML configuration for OpenSearch
Dashboards.

_Required_: No

_Type_: [SAMLOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-opensearchservice-domain-samloptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::OpenSearchService::Domain

AIMLOptions
