This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Panorama::Package StorageLocation

A storage location.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BinaryPrefixLocation" : String,
  "Bucket" : String,
  "GeneratedPrefixLocation" : String,
  "ManifestPrefixLocation" : String,
  "RepoPrefixLocation" : String
}

```

### YAML

```yaml

  BinaryPrefixLocation: String
  Bucket: String
  GeneratedPrefixLocation: String
  ManifestPrefixLocation: String
  RepoPrefixLocation: String

```

## Properties

`BinaryPrefixLocation`

The location's binary prefix.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Bucket`

The location's bucket.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GeneratedPrefixLocation`

The location's generated prefix.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestPrefixLocation`

The location's manifest prefix.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepoPrefixLocation`

The location's repo prefix.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Panorama::Package

Tag
