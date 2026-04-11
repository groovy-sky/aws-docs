This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::PolicyStore SchemaDefinition

Contains a list of principal types, resource types, and actions that can be specified
in policies stored in the same policy store. If the validation mode for the policy store is set to
`STRICT`, then policies that can't be validated by this schema are
rejected by Verified Permissions and can't be stored in the policy store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CedarJson" : String
}

```

### YAML

```yaml

  CedarJson: String

```

## Properties

`CedarJson`

A JSON string representation of the schema supported by applications that use this
policy store. For more information, see [Policy store schema](../../../verifiedpermissions/latest/userguide/schema.md)
in the AVP User Guide.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KmsEncryptionState

Tag

All content copied from https://docs.aws.amazon.com/.
