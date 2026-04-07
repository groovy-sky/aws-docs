This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Database PrincipalPrivileges

the permissions granted to a principal

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

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Principal`

The principal who is granted permissions.

_Required_: No

_Type_: [DataLakePrincipal](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-database-datalakeprincipal.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FederatedDatabase

AWS::Glue::DataCatalogEncryptionSettings
