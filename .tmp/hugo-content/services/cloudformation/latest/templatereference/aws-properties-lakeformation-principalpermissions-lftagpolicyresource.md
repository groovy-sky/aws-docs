This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::PrincipalPermissions LFTagPolicyResource

A list of LF-tag conditions that define a resource's LF-tag policy.

A structure that allows an admin to grant user permissions on certain conditions. For example, granting a role access to all columns that do not have the LF-tag 'PII' in tables that have the LF-tag 'Prod'.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CatalogId" : String,
  "Expression" : [ LFTag, ... ],
  "ResourceType" : String
}

```

### YAML

```yaml

  CatalogId: String
  Expression:
    - LFTag
  ResourceType: String

```

## Properties

`CatalogId`

The identifier for the Data Catalog. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your AWS Lake Formation environment.

_Required_: Yes

_Type_: String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Expression`

A list of LF-tag conditions that apply to the resource's LF-tag policy.

_Required_: Yes

_Type_: Array of [LFTag](aws-properties-lakeformation-principalpermissions-lftag.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceType`

The resource type for which the LF-tag policy applies.

_Required_: Yes

_Type_: String

_Allowed values_: `DATABASE | TABLE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

### Permissions on LF-tag policy resource

The following example demonstrates how to grant permissions on a `LFTagPolicy` resource.

#### JSON

```json

{
  "SamplePermission": {
      "LFTagPolicy": {
          "CatalogId": "12345678910",
          "ResourceType": "TABLE",
          "Expression": [
              {
                  "TagKey": "sample_key",
                  "TagValues": ["sample_value"]
              }
          ]
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
      LFTagPolicy:
        CatalogId: "12345678910"
        ResourceType: "TABLE"
        Expression:
          - TagKey: "sample_key"
            TagValues: "sample_value"
    Permissions:
          - "DESCRIBE"
    PermissionsWithGrantOption:
         - "DESCRIBE"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LFTagKeyResource

Resource

All content copied from https://docs.aws.amazon.com/.
