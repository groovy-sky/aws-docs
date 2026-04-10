This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain NodeToNodeEncryptionOptions

Specifies options for node-to-node encryption.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean
}

```

### YAML

```yaml

  Enabled: Boolean

```

## Properties

`Enabled`

Specifies to enable or disable node-to-node encryption on the domain. Required if you
enable fine-grained access control in [AdvancedSecurityOptionsInput](../userguide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.md).

_Required_: Conditional

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NodeOption

OffPeakWindow

All content copied from https://docs.aws.amazon.com/.
