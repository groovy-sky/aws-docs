This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::Table

The CreateTable operation adds a new table to an existing database in your account. In an
AWS account, table names must be at least unique within each Region if they
are in the same database. You may have identical table names in the same Region if the tables
are in separate databases. While creating the table, you must specify the table name, database
name, and the retention properties. [Service quotas apply](https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html). See
[code sample](https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.create-table.html)
for details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Timestream::Table",
  "Properties" : {
      "DatabaseName" : String,
      "MagneticStoreWriteProperties" : MagneticStoreWriteProperties,
      "RetentionProperties" : RetentionProperties,
      "Schema" : Schema,
      "TableName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Timestream::Table
Properties:
  DatabaseName: String
  MagneticStoreWriteProperties:
    MagneticStoreWriteProperties
  RetentionProperties:
    RetentionProperties
  Schema:
    Schema
  TableName: String
  Tags:
    - Tag

```

## Properties

`DatabaseName`

The name of the Timestream database that contains this table.

_Length Constraints_: Minimum length of 3 bytes. Maximum length of 256
bytes.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]{3,256}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MagneticStoreWriteProperties`

Contains properties to set on the table when enabling magnetic store writes.

This object has the following attributes:

- _EnableMagneticStoreWrites_: A `boolean` flag to enable
magnetic store writes.

- _MagneticStoreRejectedDataLocation_: The location to write error
reports for records rejected, asynchronously, during magnetic store writes. Only
`S3Configuration` objects are allowed. The `S3Configuration`
object has the following attributes:

- _BucketName_: The name of the S3 bucket.

- _EncryptionOption_: The encryption option for the S3 location.
Valid values are S3 server-side encryption with an S3 managed key
( `SSE_S3`) or AWS managed key ( `
                SSE_KMS`).

- _KmsKeyId_: The AWS KMS key ID to use when
encrypting with an AWS managed key.

- _ObjectKeyPrefix_: The prefix to use option for the objects
stored in S3.

Both `BucketName` and `EncryptionOption` are **required** when `S3Configuration` is specified. If you
specify ` SSE_KMS` as your `EncryptionOption` then
`KmsKeyId` is **required**.

`EnableMagneticStoreWrites` attribute is **required**
when `MagneticStoreWriteProperties` is specified.
`MagneticStoreRejectedDataLocation` attribute is **required** when `EnableMagneticStoreWrites` is set to
`true`.

See the following examples:

**JSON**

```json

{
   "Type" : AWS::Timestream::Table",
   "Properties":{
      "DatabaseName":"TestDatabase",
      "TableName":"TestTable",
      "MagneticStoreWriteProperties":{
         "EnableMagneticStoreWrites":true,
         "MagneticStoreRejectedDataLocation":{
            "S3Configuration":{
               "BucketName":"amzn-s3-demo-bucket",
               "EncryptionOption":"SSE_KMS",
               "KmsKeyId":"1234abcd-12ab-34cd-56ef-1234567890ab",
               "ObjectKeyPrefix":"prefix"
            }
         }
      }
   }
}
```

**YAML**

```

Type: AWS::Timestream::Table
DependsOn: TestDatabase
Properties:
  TableName: "TestTable"
  DatabaseName: "TestDatabase"
  MagneticStoreWriteProperties:
    EnableMagneticStoreWrites: true
    MagneticStoreRejectedDataLocation:
      S3Configuration:
        BucketName: "amzn-s3-demo-bucket"
        EncryptionOption: "SSE_KMS"
        KmsKeyId: "1234abcd-12ab-34cd-56ef-1234567890ab"
        ObjectKeyPrefix: "prefix"

```

_Required_: No

_Type_: [MagneticStoreWriteProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-timestream-table-magneticstorewriteproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetentionProperties`

The retention duration for the memory store and magnetic store. This object has the
following attributes:

- _MemoryStoreRetentionPeriodInHours_: Retention duration for memory
store, in hours.

- _MagneticStoreRetentionPeriodInDays_: Retention duration for
magnetic store, in days.

Both attributes are of type `string`. Both attributes are **required** when `RetentionProperties` is specified.

See the following examples:

**JSON**

```

{
    "Type" : AWS::Timestream::Table",
    "Properties" : {
        "DatabaseName" : "TestDatabase",
        "TableName" : "TestTable",
        "RetentionProperties" : {
            "MemoryStoreRetentionPeriodInHours": "24",
            "MagneticStoreRetentionPeriodInDays": "7"
        }
    }
}
```

**YAML**

```

Type: AWS::Timestream::Table
DependsOn: TestDatabase
Properties:
    TableName: "TestTable"
    DatabaseName: "TestDatabase"
    RetentionProperties:
        MemoryStoreRetentionPeriodInHours: "24"
        MagneticStoreRetentionPeriodInDays: "7"
```

_Required_: No

_Type_: [RetentionProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-timestream-table-retentionproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schema`

The schema of the table.

_Required_: No

_Type_: [Schema](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-timestream-table-schema.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The name of the Timestream table.

_Length Constraints_: Minimum length of 3 bytes. Maximum length of 256
bytes.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]{3,256}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to add to the table

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-timestream-table-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the table name `TABLE_NAME` in the form
`DATABASE_NAME|TABLE_NAME`. `DATABASE_NAME` is the name of the Timestream database that the table is contained in.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` returns a value for the specified attribute of this type. The
following are the available attributes:

`Arn`

The `arn` of the table.

`Name`

The name of the table.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TimestreamConfiguration

MagneticStoreRejectedDataLocation
