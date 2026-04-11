This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Type

The `CreateType` operation creates a new user-defined type in the specified keyspace.

To configure the required permissions, see [Permissions to create a UDT](../../../keyspaces/latest/devguide/configure-udt-permissions.md#udt-permissions-create)
in the _Amazon Keyspaces Developer Guide_.

For more information, see [User-defined types (UDTs)](../../../keyspaces/latest/devguide/udts.md) in the _Amazon Keyspaces Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cassandra::Type",
  "Properties" : {
      "Fields" : [ Field, ... ],
      "KeyspaceName" : String,
      "TypeName" : String
    }
}

```

### YAML

```yaml

Type: AWS::Cassandra::Type
Properties:
  Fields:
    - Field
  KeyspaceName: String
  TypeName: String

```

## Properties

`Fields`

A list of fields that define this type.

_Required_: Yes

_Type_: Array of [Field](aws-properties-cassandra-type-field.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KeyspaceName`

The name of the keyspace to create the type in. The keyspace must already exist.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TypeName`

The name of the user-defined type. UDT names must contain 48 characters or less, must begin with
an alphabetic character, and can only contain alpha-numeric characters and underscores. Amazon Keyspaces converts upper case characters automatically into
lower case characters. For more information,
see [Create a user-defined type (UDT) in Amazon Keyspaces](../../../keyspaces/latest/devguide/keyspaces-create-udt.md) in the _Amazon Keyspaces Developer_
_Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the type and the keyspace where the type exists (delimited by '\|'). For example:

`{ "Ref": "myKeyspace|myType" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DirectParentTypes`

A list of user-defined types that use this type.

`DirectReferringTables`

A list of tables that use this type.

`KeyspaceArn`

The unique identifier of the keyspace that contains this type
in the format of Amazon Resource Name (ARN)

`LastModifiedTimestamp`

The last time this type was modified.

`MaxNestingDepth`

The maximum nesting depth of this type. For more information,
see [Amazon Keyspaces UDT quotas and default values](../../../keyspaces/latest/devguide/quotas.md#quotas-udts) in the _Amazon Keyspaces Developer_
_Guide_.

## Examples

This section includes code examples that demonstrate how to create and use user-defined types (UDTs).

- [Create a new UDT](#aws-resource-cassandra-type--examples--Create_a_new_UDT)

- [Create a new UDT with a nested UDT](#aws-resource-cassandra-type--examples--Create_a_new_UDT_with_a_nested_UDT)

- [Create a new table that uses UDTs](#aws-resource-cassandra-type--examples--Create_a_new_table_that_uses_UDTs)

### Create a new UDT

The following example creates a new type named
`my_udt`.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "MyNewKeyspaceResource": {
      "Type": "AWS::Cassandra::Keyspace",
      "Properties": {
        "KeyspaceName": "my_keyspace"
      }
    },
    "MyNewTypeResource": {
      "Type": "AWS::Cassandra::Type",
      "Properties": {
        "KeyspaceName": "my_keyspace",
        "TypeName": "my_udt",
        "Fields": [
          {
            "FieldName": "field_1",
            "FieldType": "int"
          }
        ]
      },
      "DependsOn": "MyNewKeyspaceResource"
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  MyNewKeyspaceResource:
    Type: 'AWS::Cassandra::Keyspace'
    Properties:
      KeyspaceName: my_keyspace
  MyNewTypeResource:
    Type: 'AWS::Cassandra::Type'
    Properties:
      KeyspaceName: my_keyspace
      TypeName: my_udt
      Fields:
        - FieldName: field_1
          FieldType: int
    DependsOn: "MyNewKeyspaceResource"
```

### Create a new UDT with a nested UDT

The following example creates a new UDT named
`parent_udt` that uses the nested type `child_udt`.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "MyNewKeyspaceResource": {
      "Type": "AWS::Cassandra::Keyspace",
      "Properties": {
        "KeyspaceName": "my_ks"
      }
    },
    "MyNewChildTypeResource": {
      "Type": "AWS::Cassandra::Type",
      "Properties": {
        "KeyspaceName": "my_ks",
        "TypeName": "child_udt",
        "Fields": [
          {
            "FieldName": "field_a",
            "FieldType": "int"
          }
        ]
      },
      "DependsOn": "MyNewKeyspaceResource"
    },
    "MyNewParentTypeResource": {
      "Type": "AWS::Cassandra::Type",
      "Properties": {
        "KeyspaceName": "my_ks",
        "TypeName": "parent_udt",
        "Fields": [
          {
            "FieldName": "child_type_field",
            "FieldType": "frozen<child_udt>"
          },
          {
            "FieldName": "int_field",
            "FieldType": "int"
          }
        ]
      },
      "DependsOn": ["MyNewChildTypeResource", "MyNewKeyspaceResource"]
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  MyNewKeyspaceResource:
    Type: 'AWS::Cassandra::Keyspace'
    Properties:
      KeyspaceName: my_ks
  MyNewChildTypeResource:
    Type: 'AWS::Cassandra::Type'
    Properties:
      KeyspaceName: my_ks
      TypeName: child_udt
      Fields:
        - FieldName: field_a
          FieldType: int
    DependsOn: "MyNewKeyspaceResource"
  MyNewParentTypeResource:
    Type: 'AWS::Cassandra::Type'
    Properties:
      KeyspaceName: my_ks
      TypeName: parent_udt
      Fields:
        - FieldName: child_type_field
          FieldType: frozen<child_udt>
        - FieldName: int_field
          FieldType: int
    DependsOn: ["MyNewChildTypeResource", "MyNewKeyspaceResource"]
```

### Create a new table that uses UDTs

The following example creates a new table called `my_table` that uses a UDT named
`my_udt`.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "MyNewKeyspaceResource": {
      "Type": "AWS::Cassandra::Keyspace",
      "Properties": {
        "KeyspaceName": "my_ks"
      }
    },
    "MyNewTypeResource": {
      "Type": "AWS::Cassandra::Type",
      "Properties": {
        "KeyspaceName": "my_ks",
        "TypeName": "my_udt",
        "Fields": [
          {
            "FieldName": "field_1",
            "FieldType": "int"
          }
        ]
      },
      "DependsOn": "MyNewKeyspaceResource"
    },
    "MyNewTableResource": {
      "Type": "AWS::Cassandra::Table",
      "Properties": {
        "KeyspaceName": "my_ks",
        "TableName": "my_table",
        "PartitionKeyColumns": [
          {
            "ColumnName": "frozen_udt",
            "ColumnType": "frozen<my_udt>"
          }
        ],
        "RegularColumns": [
          {
            "ColumnName": "udt_field",
            "ColumnType": "my_udt"
          },
          {
            "ColumnName": "other_field",
            "ColumnType": "int"
          }
        ]
      },
      "DependsOn": ["MyNewTypeResource", "MyNewKeyspaceResource"]
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  MyNewKeyspaceResource:
    Type: 'AWS::Cassandra::Keyspace'
    Properties:
      KeyspaceName: my_ks
  MyNewTypeResource:
    Type: 'AWS::Cassandra::Type'
    Properties:
      KeyspaceName: my_ks
      TypeName: my_udt
      Fields:
        - FieldName: field_1
          FieldType: int
    DependsOn: "MyNewKeyspaceResource"
  MyNewTableResource:
    Type: 'AWS::Cassandra::Table'
    Properties:
      KeyspaceName: my_ks
      TableName: my_table
      PartitionKeyColumns:
        - ColumnName: frozen_udt
          ColumnType: frozen<my_udt>
      RegularColumns:
        - ColumnName: udt_field
          ColumnType: my_udt
        - ColumnName: other_field
          ColumnType: int
    DependsOn: ["MyNewTypeResource", "MyNewKeyspaceResource"]
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WarmThroughput

Field

All content copied from https://docs.aws.amazon.com/.
