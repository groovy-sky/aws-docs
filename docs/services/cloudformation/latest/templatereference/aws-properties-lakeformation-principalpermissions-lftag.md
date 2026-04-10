This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::PrincipalPermissions LFTag

The LF-tag key and values attached to a resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TagKey" : String,
  "TagValues" : [ String, ... ]
}

```

### YAML

```yaml

  TagKey: String
  TagValues:
    - String

```

## Properties

`TagKey`

The key-name for the LF-tag.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TagValues`

A list of possible values of the corresponding `TagKey` of an LF-tag key-value pair.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

### Permissons on an LF-tag

The following example demonstrates how to grant permissions on a `LFTag` resource:

#### JSON

```json

{
    "SamplePermission": {
        "LFTag": {
            "CatalogId": "12345678910",
            "TagKey": "sample_key",
            "TagValues": ["sample_value"]
        }
    },
    "Permissions": ["DESCRIBE"],
    "PermissionsWithGrantOption": ["DESCRIBE"]
}
```

#### YAML

```yaml

SamplePermission:
  Type: AWS::LakeFormation::PrincipalPermissions
  Properties:
    Principal:
      DataLakePrincipalIdentifier: "arn:sample_principal"
    Resource:
      LFTag:
        CatalogId: "12345678910"
        TagKey: "sample_key"
        TagValues:
           - "sample_value"
    Permissions:
          - "DESCRIBE"
    PermissionsWithGrantOption:
         - "DESCRIBE"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataLocationResource

LFTagKeyResource

All content copied from https://docs.aws.amazon.com/.
