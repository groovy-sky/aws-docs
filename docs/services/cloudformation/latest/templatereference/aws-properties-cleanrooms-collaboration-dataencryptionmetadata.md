This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Collaboration DataEncryptionMetadata

The settings for client-side encryption for cryptographic computing.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowCleartext" : Boolean,
  "AllowDuplicates" : Boolean,
  "AllowJoinsOnColumnsWithDifferentNames" : Boolean,
  "PreserveNulls" : Boolean
}

```

### YAML

```yaml

  AllowCleartext: Boolean
  AllowDuplicates: Boolean
  AllowJoinsOnColumnsWithDifferentNames: Boolean
  PreserveNulls: Boolean

```

## Properties

`AllowCleartext`

Indicates whether encrypted tables can contain cleartext data ( `TRUE`) or are
to cryptographically process every column ( `FALSE`).

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AllowDuplicates`

Indicates whether Fingerprint columns can contain duplicate entries ( `TRUE`)
or are to contain only non-repeated values ( `FALSE`).

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AllowJoinsOnColumnsWithDifferentNames`

Indicates whether Fingerprint columns can be joined on any other Fingerprint column with
a different name ( `TRUE`) or can only be joined on Fingerprint columns of the
same name ( `FALSE`).

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PreserveNulls`

Indicates whether NULL values are to be copied as NULL to encrypted tables
( `TRUE`) or cryptographically processed ( `FALSE`).

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CleanRooms::Collaboration

JobComputePaymentConfig
