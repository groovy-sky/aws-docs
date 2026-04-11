This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::PrincipalPermissions DataLocationResource

A structure for a data location object where permissions are granted or revoked.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CatalogId" : String,
  "ResourceArn" : String
}

```

### YAML

```yaml

  CatalogId: String
  ResourceArn: String

```

## Properties

`CatalogId`

The identifier for the Data Catalog where the location is registered with AWS Lake Formation.

_Required_: Yes

_Type_: String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceArn`

The Amazon Resource Name (ARN) that uniquely identifies the data location resource.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataLakePrincipal

LFTag

All content copied from https://docs.aws.amazon.com/.
