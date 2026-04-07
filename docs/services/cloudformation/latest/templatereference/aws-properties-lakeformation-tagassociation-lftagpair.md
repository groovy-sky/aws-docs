This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::TagAssociation LFTagPair

A structure containing the catalog ID, tag key, and tag values of an LF-tag key-value pair.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CatalogId" : String,
  "TagKey" : String,
  "TagValues" : [ String, ... ]
}

```

### YAML

```yaml

  CatalogId: String
  TagKey: String
  TagValues:
    - String

```

## Properties

`CatalogId`

The identifier for the Data Catalog. By default, it is the account ID of the caller.

_Required_: Yes

_Type_: String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TagKey`

The key-name for the LF-tag.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TagValues`

A list of possible values of the corresponding `TagKey` of an LF-tag key-value pair.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DatabaseResource

Resource
