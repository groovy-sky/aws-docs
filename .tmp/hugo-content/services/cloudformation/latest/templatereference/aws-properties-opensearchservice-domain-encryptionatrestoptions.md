This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain EncryptionAtRestOptions

Whether the domain should encrypt data at rest, and if so, the AWS Key Management Service key to use.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "KmsKeyId" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  KmsKeyId: String

```

## Properties

`Enabled`

Specify `true` to enable encryption at rest. Required if you enable
fine-grained access control in [AdvancedSecurityOptionsInput](../userguide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.md).

If no encryption at rest options were initially specified in the template, updating this property by adding it causes no interruption. However, if you change this property after it's already been set within a template,
the domain is deleted and recreated in order to modify the property.

_Required_: Conditional

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`KmsKeyId`

The KMS key ID. Takes the form `1a2a3a4-1a2a-3a4a-5a6a-1a2a3a4a5a6a`. Required
if you enable encryption at rest.

You can also use `keyAlias` as a value.

If no encryption at rest options were initially specified in the template, updating this property by adding it causes no interruption. However, if you change this property after it's already been set within a template,
the domain is deleted and recreated in order to modify the property.

_Required_: Conditional

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## See also

- [CreateDomain](../../../opensearch-service/latest/developerguide/configuration-api.md#configuration-api-actions-createdomain) in the _Amazon OpenSearch Service Developer_
_Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EBSOptions

IAMFederationOptions

All content copied from https://docs.aws.amazon.com/.
