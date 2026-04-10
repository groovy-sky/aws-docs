This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::DatabaseSnapshot

Creates a snapshot of your database in Amazon Lightsail. You can use snapshots for backups,
to make copies of a database, and to save data before deleting a database.

The `create relational database snapshot` operation supports tag-based access
control via request tags. For more information, see the [Amazon Lightsail Developer Guide](../../../lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lightsail::DatabaseSnapshot",
  "Properties" : {
      "RelationalDatabaseName" : String,
      "RelationalDatabaseSnapshotName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Lightsail::DatabaseSnapshot
Properties:
  RelationalDatabaseName: String
  RelationalDatabaseSnapshotName: String
  Tags:
    - Tag

```

## Properties

`RelationalDatabaseName`

The unique name of the database resource in Lightsail.

_Required_: Yes

_Type_: String

_Pattern_: `^\w[\w\-]*\w$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RelationalDatabaseSnapshotName`

The name of the database snapshot.

_Required_: Yes

_Type_: String

_Pattern_: `^\w[\w\-]*\w$`

_Minimum_: `2`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tag keys and optional values for the resource. For more information about tags in
Lightsail, see the [Amazon Lightsail Developer Guide](../../../lightsail/latest/userguide/amazon-lightsail-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-lightsail-databasesnapshot-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the database snapshot.

`CreatedAt`

The timestamp when the database snapshot was created.

`Engine`

The software of the database snapshot (for example, `MySQL`)

`EngineVersion`

The database engine version for the database snapshot (for example,
`5.7.23`).

`FromRelationalDatabaseArn`

The Amazon Resource Name (ARN) of the database from which the database snapshot was
created.

`FromRelationalDatabaseBlueprintId`

The blueprint ID of the database from which the database snapshot was created. A blueprint
describes the major engine version of a database.

`FromRelationalDatabaseBundleId`

The bundle ID of the database from which the database snapshot was created.

`FromRelationalDatabaseName`

The name of the source database from which the database snapshot was created.

`Name`

The name of the database snapshot.

`ResourceType`

The Lightsail resource type.

`SizeInGb`

The size of the disk in GB (for example, `32`) for the database
snapshot.

`State`

The state of the database snapshot.

`SupportCode`

The support code for the database snapshot. Include this code in your email to support
when you have questions about a database snapshot in Lightsail. This code enables our
support team to look up your Lightsail information more easily.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Location

All content copied from https://docs.aws.amazon.com/.
