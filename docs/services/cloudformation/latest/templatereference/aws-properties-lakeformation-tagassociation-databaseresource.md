This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::TagAssociation DatabaseResource

A structure for the database object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CatalogId" : String,
  "Name" : String
}

```

### YAML

```yaml

  CatalogId: String
  Name: String

```

## Properties

`CatalogId`

The identifier for the Data Catalog. By default, it should be the account ID of the caller.

_Required_: Yes

_Type_: String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the database resource. Unique to the Data Catalog.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::LakeFormation::TagAssociation

LFTagPair

All content copied from https://docs.aws.amazon.com/.
