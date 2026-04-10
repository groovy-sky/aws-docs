This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::WorkGroup ManagedQueryResultsConfiguration

The configuration for storing results in Athena owned storage, which includes whether this feature is enabled; whether encryption configuration, if any, is used for encrypting query results.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "EncryptionConfiguration" : ManagedStorageEncryptionConfiguration
}

```

### YAML

```yaml

  Enabled: Boolean
  EncryptionConfiguration:
    ManagedStorageEncryptionConfiguration

```

## Properties

`Enabled`

If set to true, allows you to store query results in Athena owned storage. If set to
false, workgroup member stores query results in location specified under
`ResultConfiguration$OutputLocation`. The default is false. A workgroup
cannot have the `ResultConfiguration$OutputLocation` parameter when you set
this field to true.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfiguration`

If you encrypt query and calculation results in Athena owned storage, this field
indicates the encryption option (for example, SSE\_KMS or CSE\_KMS) and key
information.

_Required_: No

_Type_: [ManagedStorageEncryptionConfiguration](aws-properties-athena-workgroup-managedstorageencryptionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedLoggingConfiguration

ManagedStorageEncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
