This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application ApplicationEncryptionConfiguration

Specifies the configuration to manage encryption at rest.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KeyId" : String,
  "KeyType" : String
}

```

### YAML

```yaml

  KeyId: String
  KeyType: String

```

## Properties

`KeyId`

The key ARN, key ID, alias ARN, or alias name of the KMS key used for encryption at rest.

_Required_: No

_Type_: String

_Pattern_: `^(?:arn:.*:kms:.*:.*:(?:key\/.*|alias\/.*)|alias\/.*|(?i)[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyType`

Specifies the type of key used for encryption at rest.

_Required_: Yes

_Type_: String

_Allowed values_: `AWS_OWNED_KEY | CUSTOMER_MANAGED_KEY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApplicationConfiguration

ApplicationMaintenanceConfiguration

All content copied from https://docs.aws.amazon.com/.
