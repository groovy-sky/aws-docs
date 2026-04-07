This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Database

The `AWS::Lightsail::Database` resource specifies an Amazon Lightsail database.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lightsail::Database",
  "Properties" : {
      "AvailabilityZone" : String,
      "BackupRetention" : Boolean,
      "CaCertificateIdentifier" : String,
      "MasterDatabaseName" : String,
      "MasterUsername" : String,
      "MasterUserPassword" : String,
      "PreferredBackupWindow" : String,
      "PreferredMaintenanceWindow" : String,
      "PubliclyAccessible" : Boolean,
      "RelationalDatabaseBlueprintId" : String,
      "RelationalDatabaseBundleId" : String,
      "RelationalDatabaseName" : String,
      "RelationalDatabaseParameters" : [ RelationalDatabaseParameter, ... ],
      "RotateMasterUserPassword" : Boolean,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Lightsail::Database
Properties:
  AvailabilityZone: String
  BackupRetention: Boolean
  CaCertificateIdentifier: String
  MasterDatabaseName: String
  MasterUsername: String
  MasterUserPassword: String
  PreferredBackupWindow: String
  PreferredMaintenanceWindow: String
  PubliclyAccessible: Boolean
  RelationalDatabaseBlueprintId: String
  RelationalDatabaseBundleId: String
  RelationalDatabaseName: String
  RelationalDatabaseParameters:
    - RelationalDatabaseParameter
  RotateMasterUserPassword: Boolean
  Tags:
    - Tag

```

## Properties

`AvailabilityZone`

The Availability Zone for the database.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: Updates are not supported.

`BackupRetention`

A Boolean value indicating whether automated backup retention is enabled for the
database. Data Import Mode is enabled when `BackupRetention` is set to `false`, and is disabled when `BackupRetention` is set to `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`CaCertificateIdentifier`

The certificate associated with the database.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterDatabaseName`

The meaning of this parameter differs according to the database engine you use.

**MySQL**

The name of the database to create when the Lightsail database resource
is created. If this parameter isn't specified, no database is created in the database
resource.

Constraints:

- Must contain 1-64 letters or numbers.

- Must begin with a letter. Subsequent characters can be letters, underscores, or
numbers (0-9).

- Can't be a word reserved by the specified database engine.

For more information about reserved words in MySQL, see the Keywords and Reserved
Words articles for [MySQL 5.6](https://dev.mysql.com/doc/refman/5.6/en/keywords.html), [MySQL 5.7](https://dev.mysql.com/doc/refman/5.7/en/keywords.html), and
[MySQL\
8.0](https://dev.mysql.com/doc/refman/8.0/en/keywords.html).

**PostgreSQL**

The name of the database to create when the Lightsail database resource
is created. If this parameter isn't specified, a database named `postgres` is
created in the database resource.

Constraints:

- Must contain 1-63 letters or numbers.

- Must begin with a letter. Subsequent characters can be letters, underscores, or
numbers (0-9).

- Can't be a word reserved by the specified database engine.

For more information about reserved words in PostgreSQL, see the SQL Key Words
articles for [PostgreSQL\
9.6](https://www.postgresql.org/docs/9.6/sql-keywords-appendix.html), [PostgreSQL\
10](https://www.postgresql.org/docs/10/sql-keywords-appendix.html), [PostgreSQL\
11](https://www.postgresql.org/docs/11/sql-keywords-appendix.html), and [PostgreSQL\
12](https://www.postgresql.org/docs/12/sql-keywords-appendix.html).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: Updates are not supported.

`MasterUsername`

The name for the primary user.

**MySQL**

Constraints:

- Required for MySQL.

- Must be 1-16 letters or numbers. Can contain underscores.

- First character must be a letter.

- Can't be a reserved word for the chosen database engine.

For more information about reserved words in MySQL 5.6 or 5.7, see the Keywords
and Reserved Words articles for [MySQL 5.6](https://dev.mysql.com/doc/refman/5.6/en/keywords.html),
[MySQL\
5.7](https://dev.mysql.com/doc/refman/5.7/en/keywords.html), or [MySQL 8.0](https://dev.mysql.com/doc/refman/8.0/en/keywords.html).

**PostgreSQL**

Constraints:

- Required for PostgreSQL.

- Must be 1-63 letters or numbers. Can contain underscores.

- First character must be a letter.

- Can't be a reserved word for the chosen database engine.

For more information about reserved words in MySQL 5.6 or 5.7, see the Keywords
and Reserved Words articles for [PostgreSQL\
9.6](https://www.postgresql.org/docs/9.6/sql-keywords-appendix.html), [PostgreSQL\
10](https://www.postgresql.org/docs/10/sql-keywords-appendix.html), [PostgreSQL\
11](https://www.postgresql.org/docs/11/sql-keywords-appendix.html), and [PostgreSQL\
12](https://www.postgresql.org/docs/12/sql-keywords-appendix.html).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: Updates are not supported.

`MasterUserPassword`

The password for the primary user of the database. The password can include any
printable ASCII character except the following: /, ", or @. It cannot contain
spaces.

###### Note

The `MasterUserPassword` and `RotateMasterUserPassword`
parameters cannot be used together in the same template.

**MySQL**

Constraints: Must contain 8-41 characters.

**PostgreSQL**

Constraints: Must contain 8-128 characters.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredBackupWindow`

The daily time range during which automated backups are created for the database (for
example, `16:00-16:30`).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredMaintenanceWindow`

The weekly time range during which system maintenance can occur for the database,
formatted as follows: `ddd:hh24:mi-ddd:hh24:mi`. For example,
`Tue:17:00-Tue:17:30`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PubliclyAccessible`

A Boolean value indicating whether the database is accessible to anyone on the
internet.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelationalDatabaseBlueprintId`

The blueprint ID for the database (for example, `mysql_8_0`).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: Updates are not supported.

`RelationalDatabaseBundleId`

The bundle ID for the database (for example, `medium_1_0`).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: Updates are not supported.

`RelationalDatabaseName`

The name of the instance.

_Required_: Yes

_Type_: String

_Pattern_: `\w[\w\-]*\w`

_Minimum_: `2`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RelationalDatabaseParameters`

An array of parameters for the database.

_Required_: No

_Type_: Array of [RelationalDatabaseParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-database-relationaldatabaseparameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RotateMasterUserPassword`

A Boolean value indicating whether to change the primary user password to a new, strong
password generated by Lightsail.

###### Note

The `RotateMasterUserPassword` and `MasterUserPassword`
parameters cannot be used together in the same template.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html)
in the _AWS CloudFormation User Guide_.

###### Note

The `Value` of `Tags` is optional for Lightsail resources.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-database-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DatabaseArn`

The Amazon Resource Name (ARN) of the database (for example,
`arn:aws:lightsail:us-east-2:123456789101:RelationalDatabase/244ad76f-8aad-4741-809f-12345EXAMPLE`).

## Remarks

_Availability Zone_

You can specify an Availability Zone when you perform a create database request. If
you don’t specify one, the database is created in the same Availability Zone as the last
Lightsail resource you created.

_Database updates_

All database update operations are performed immediately. However, parameter updates
are applied based on the `ApplyMethod` value for the specific parameter being
updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

RelationalDatabaseParameter
