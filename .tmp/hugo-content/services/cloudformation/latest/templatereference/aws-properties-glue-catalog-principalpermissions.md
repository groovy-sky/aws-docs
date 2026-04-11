This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Catalog PrincipalPermissions

Permissions granted to a principal.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Permissions" : [ String, ... ],
  "Principal" : DataLakePrincipal
}

```

### YAML

```yaml

  Permissions:
    - String
  Principal:
    DataLakePrincipal

```

## Properties

`Permissions`

The permissions that are granted to the principal.

_Required_: No

_Type_: Array of String

_Allowed values_: `ALL | SELECT | ALTER | DROP | DELETE | INSERT | CREATE_DATABASE | CREATE_TABLE | DATA_LOCATION_ACCESS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Principal`

The principal who is granted permissions.

_Required_: No

_Type_: [DataLakePrincipal](aws-properties-glue-catalog-datalakeprincipal.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FederatedCatalog

Tag

All content copied from https://docs.aws.amazon.com/.
