This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Database FederatedDatabase

A `FederatedDatabase` structure that references an entity outside the AWS Glue Data Catalog.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectionName" : String,
  "Identifier" : String
}

```

### YAML

```yaml

  ConnectionName: String
  Identifier: String

```

## Properties

`ConnectionName`

The name of the connection to the external metastore.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Identifier`

A unique identifier for the federated database.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataLakePrincipal

PrincipalPrivileges
