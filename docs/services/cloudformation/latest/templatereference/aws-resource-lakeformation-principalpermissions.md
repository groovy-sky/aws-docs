This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::PrincipalPermissions

The `AWS::LakeFormation::PrincipalPermissions` resource represents the permissions that a principal has on a Data Catalog resource (such as AWS Glue databases or AWS Glue tables).
When you create a `PrincipalPermissions` resource, the permissions are granted via the AWS Lake Formation `GrantPermissions` API operation. When you delete a `PrincipalPermissions` resource, the permissions on principal-resource pair are revoked via the AWS Lake Formation `RevokePermissions` API operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::LakeFormation::PrincipalPermissions",
  "Properties" : {
      "Catalog" : String,
      "Permissions" : [ String, ... ],
      "PermissionsWithGrantOption" : [ String, ... ],
      "Principal" : DataLakePrincipal,
      "Resource" : Resource
    }
}

```

### YAML

```yaml

Type: AWS::LakeFormation::PrincipalPermissions
Properties:
  Catalog: String
  Permissions:
    - String
  PermissionsWithGrantOption:
    - String
  Principal:
    DataLakePrincipal
  Resource:
    Resource

```

## Properties

`Catalog`

The identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store.
It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.

_Required_: No

_Type_: String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Permissions`

The permissions granted or revoked.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PermissionsWithGrantOption`

Indicates the ability to grant permissions (as a subset of permissions granted).

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Principal`

The principal to be granted a permission.

_Required_: Yes

_Type_: [DataLakePrincipal](aws-properties-lakeformation-principalpermissions-datalakeprincipal.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Resource`

The resource to be granted or revoked permissions.

_Required_: Yes

_Type_: [Resource](aws-properties-lakeformation-principalpermissions-resource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the primary identifier of the resource. The primary identifier of the resource is a combination of `ResourceIdentifier` and `PrincipalIdentifier`
separated by a pipe.

For example:
`{"DataLakePrincipalIdentifier":"arn:aws:iam::123456789012:role/ExampleRole"}|{"Catalog":null,"Database":null,"Table":null,"TableWithColumns":null,"DataLocation":null,"DataCellsFilter":{"TableCatalogId":"123456789012","DatabaseName":"ExampleDatabase","TableName":"ExampleTable","Name":"ExampleFilter"},"LFTag":null,"LFTagPolicy":null}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`PrincipalIdentifier`

Json encoding of the input principal. For example: `{"DataLakePrincipalIdentifier":"arn:aws:iam::123456789012:role/ExampleRole"}`

`ResourceIdentifier`

Json encoding of the input resource. For example: `{"Catalog":null,"Database":null,"Table":null,"TableWithColumns":null,"DataLocation":null,"DataCellsFilter":{"TableCatalogId":"123456789012","DatabaseName":"ExampleDatabase","TableName":"ExampleTable","Name":"ExampleFilter"},"LFTag":null,"LFTagPolicy":null}`

## Remarks

When you delete a `PrincipalPermissions` resource, AWS Lake Formation
revokes all permissions (even manually granted ones) that the principal has on the
resource.

CloudFormation resources must have a one-to-one mapping between a defined resource (an AWS Lake Formation permission) and the primary identifier (a combination of AWS Lake Formation resource and AWS Lake Formation principal for which this permission is being granted).
Due to this limitation, the current implementation of `PrincipalPermissions` has the following behavior for permissions on `TableWithColumns` resources:

- When you create a `TableWithColumn` permissions resource, CloudFormation will first check whether the principal already
has any permissions on the underlying resource. CloudFormation will create the resource only
if there is no previous permission associated to the specific principal
for the same resource identifier. If there exists a permission resource with the same
combination, the request will fail with the `AlreadyExistsException` error.

###### Note

This limitation is also applicable to having a `SELECT` permission on the table since that is effectively a `SELECT` permission on a `ColumnWildcard` in a `TableWithColumns` resource.

## Examples

- [Permissions on a database](#aws-resource-lakeformation-principalpermissions--examples--Permissions_on_a_database)

- [Permissions on a table](#aws-resource-lakeformation-principalpermissions--examples--Permissions_on_a_table)

- [Permissions on columns](#aws-resource-lakeformation-principalpermissions--examples--Permissions_on_columns)

### Permissions on a database

The following example demonstrates how to grant permissions on a `Database` resource:

#### JSON

```json

{
    "SamplePermission": {
    "Type": "AWS::LakeFormation::PrincipalPermissions",
        "Properties": {
            "Principal": {
                "DataLakePrincipalIdentifier": " "arn:sample_principal"
            },
            "Resource": {
                "Database": {
                    "CatalogId": "12345678910",
                    "Name": "sample_db"
                }
            },
            "Permissions": ["CREATE_TABLE", "ALTER", "DROP", "DESCRIBE"],
            "PermissionsWithGrantOption": ["CREATE_TABLE", "ALTER", "DROP", "DESCRIBE"]
        }
    }
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
        Database:
          CatalogId: "12345678910"
          Name: "sample_db"
      Permissions:
        - "CREATE_TABLE"
        - "ALTER"
        - "DROP"
        - "DESCRIBE"
      PermissionsWithGrantOption:
        - "CREATE_TABLE"
        - "ALTER"
        - "DROP"
        - "DESCRIBE"

```

### Permissions on a table

The following example demonstrates how to grant permissions on a `Table` resource:

#### JSON

```json

{
    "SamplePermission": {
    "Type": "AWS::LakeFormation::PrincipalPermissions",
        "Properties": {
            "Principal": {
                "DataLakePrincipalIdentifier": " "arn:sample_principal"
            },
            "Resource": {
                "Table": {
                    "CatalogId": "12345678910",
                    "DatabaseName": "sample_db",
                    "Name": "sample_tbl"
                }
            },
            "Permissions": ["SELECT", "INSERT", "DELETE", "ALTER", "DROP", "DESCRIBE"],
            "PermissionsWithGrantOption": ["SELECT", "INSERT", "DELETE", "ALTER", "DROP", "DESCRIBE"]
        }
    }
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
        Table:
          CatalogId: "12345678910"
          DatabaseName: "sample_db"
          Name: "sample_tbl"
      Permissions:
        - "SELECT"
        - "INSERT"
        - "DELETE"
        - "ALTER"
        - "DROP"
        - "DESCRIBE"
      PermissionsWithGrantOption:
        - "SELECT"
        - "INSERT"
        - "DELETE"
        - "ALTER"
        - "DROP"
        - "DESCRIBE"
```

### Permissions on columns

The following example demonstrates how to grant permissions on a `TableWithColumns`
resource:

#### JSON

```json

{
    "SamplePermission": {
    "Type": "AWS::LakeFormation::PrincipalPermissions",
        "Properties": {
            "Principal": {
                "DataLakePrincipalIdentifier": " "arn:sample_principal"
            },
            "Resource": {
                "TableWithColumns": {
                    "CatalogId": "12345678910",
                    "DatabaseName": "sample_db",
                    "Name": "sample_tbl",
                    "ColumnNames": ["sample_col1", "sample_col2"]
                }
            },
            "Permissions": ["SELECT"],
            "PermissionsWithGrantOption": ["SELECT"]
        }
    }
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
        TableWithColumns:
          CatalogId: "12345678910"
          DatabaseName: "sample_db"
          Name: "sample_tbl"
          ColumnNames:
            - "sample_col1"
      Permissions:
        - "SELECT"
      PermissionsWithGrantOption:
        - "SELECT"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableWithColumnsResource

ColumnWildcard

All content copied from https://docs.aws.amazon.com/.
