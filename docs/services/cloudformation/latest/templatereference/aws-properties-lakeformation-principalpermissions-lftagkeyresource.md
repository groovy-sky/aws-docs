This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::PrincipalPermissions LFTagKeyResource

A structure containing an LF-tag key and values for a resource.

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

The identifier for the Data Catalog where the location is registered with Data Catalog.

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

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TagValues`

A list of possible values for the corresponding `TagKey` of an LF-tag key-value pair.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LFTag

LFTagPolicyResource

All content copied from https://docs.aws.amazon.com/.
