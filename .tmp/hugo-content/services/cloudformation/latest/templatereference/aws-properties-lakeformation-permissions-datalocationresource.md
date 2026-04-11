This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::Permissions DataLocationResource

A structure for a data location object where permissions are granted or revoked.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CatalogId" : String,
  "S3Resource" : String
}

```

### YAML

```yaml

  CatalogId: String
  S3Resource: String

```

## Properties

`CatalogId`

The identifier for the Data Catalog. By default, it is the account ID of the caller.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Resource`

The Amazon Resource Name (ARN) that uniquely identifies the data location resource.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataLakePrincipal

Resource

All content copied from https://docs.aws.amazon.com/.
