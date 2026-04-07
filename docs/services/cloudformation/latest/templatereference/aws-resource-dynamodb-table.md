This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::Table

The `AWS::DynamoDB::Table` resource creates a DynamoDB table. For
more information, see [CreateTable](../../../../reference/amazondynamodb/latest/apireference/api-createtable.md) in the _Amazon DynamoDB API Reference_.

You should be aware of the following behaviors when working with DynamoDB
tables:

- AWS CloudFormation typically creates DynamoDB tables in parallel.
However, if your template includes multiple DynamoDB tables with indexes,
you must declare dependencies so that the tables are created sequentially. Amazon DynamoDB limits the number of tables with secondary indexes that are in
the creating state. If you create multiple tables with indexes at the same time,
DynamoDB returns an error and the stack operation fails. For an
example, see [DynamoDB Table with a DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#aws-resource-dynamodb-table--examples--DynamoDB_Table_with_a_DependsOn_Attribute).

###### Important

Our guidance is to use the latest schema documented for your AWS CloudFormation templates. This schema supports the provisioning of all table settings below. When
using this schema in your AWS CloudFormation templates, please ensure that your
Identity and Access Management (IAM) policies are updated with
appropriate permissions to allow for the authorization of these setting changes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DynamoDB::Table",
  "Properties" : {
      "AttributeDefinitions" : [ AttributeDefinition, ... ],
      "BillingMode" : String,
      "ContributorInsightsSpecification" : ContributorInsightsSpecification,
      "DeletionProtectionEnabled" : Boolean,
      "GlobalSecondaryIndexes" : [ GlobalSecondaryIndex, ... ],
      "ImportSourceSpecification" : ImportSourceSpecification,
      "KeySchema" : [ KeySchema, ... ],
      "KinesisStreamSpecification" : KinesisStreamSpecification,
      "LocalSecondaryIndexes" : [ LocalSecondaryIndex, ... ],
      "OnDemandThroughput" : OnDemandThroughput,
      "PointInTimeRecoverySpecification" : PointInTimeRecoverySpecification,
      "ProvisionedThroughput" : ProvisionedThroughput,
      "ResourcePolicy" : ResourcePolicy,
      "SSESpecification" : SSESpecification,
      "StreamSpecification" : StreamSpecification,
      "TableClass" : String,
      "TableName" : String,
      "Tags" : [ Tag, ... ],
      "TimeToLiveSpecification" : TimeToLiveSpecification,
      "WarmThroughput" : WarmThroughput
    }
}

```

### YAML

```yaml

Type: AWS::DynamoDB::Table
Properties:
  AttributeDefinitions:
    - AttributeDefinition
  BillingMode: String
  ContributorInsightsSpecification:
    ContributorInsightsSpecification
  DeletionProtectionEnabled: Boolean
  GlobalSecondaryIndexes:
    - GlobalSecondaryIndex
  ImportSourceSpecification:
    ImportSourceSpecification
  KeySchema:
    - KeySchema
  KinesisStreamSpecification:
    KinesisStreamSpecification
  LocalSecondaryIndexes:
    - LocalSecondaryIndex
  OnDemandThroughput:
    OnDemandThroughput
  PointInTimeRecoverySpecification:
    PointInTimeRecoverySpecification
  ProvisionedThroughput:
    ProvisionedThroughput
  ResourcePolicy:
    ResourcePolicy
  SSESpecification:
    SSESpecification
  StreamSpecification:
    StreamSpecification
  TableClass: String
  TableName: String
  Tags:
    - Tag
  TimeToLiveSpecification:
    TimeToLiveSpecification
  WarmThroughput:
    WarmThroughput

```

## Properties

`AttributeDefinitions`

A list of attributes that describe the key schema for the table and indexes.

This property is required to create a DynamoDB table.

Update requires: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt). Replacement if you edit an existing
AttributeDefinition.

_Required_: Conditional

_Type_: Array of [AttributeDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-attributedefinition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BillingMode`

Specify how you are charged for read and write throughput and how you manage
capacity.

Valid values include:

- `PAY_PER_REQUEST` \- We recommend using `PAY_PER_REQUEST` for
most DynamoDB workloads. `PAY_PER_REQUEST` sets the billing mode to [On-demand\
capacity mode](../../../dynamodb/latest/developerguide/on-demand-capacity-mode.md).

- `PROVISIONED` \- We recommend using `PROVISIONED` for steady
workloads with predictable growth where capacity requirements can be reliably
forecasted. `PROVISIONED` sets the billing mode to [Provisioned capacity mode](../../../dynamodb/latest/developerguide/provisioned-capacity-mode.md).

If not specified, the default is `PROVISIONED`.

_Required_: No

_Type_: String

_Allowed values_: `PROVISIONED | PAY_PER_REQUEST`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContributorInsightsSpecification`

The settings used to specify whether to enable CloudWatch Contributor Insights for the
table and define which events to monitor.

_Required_: No

_Type_: [ContributorInsightsSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-contributorinsightsspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeletionProtectionEnabled`

Determines if a table is protected from deletion. When enabled, the table cannot be
deleted by any user or process. This setting is disabled by default. For more information,
see [Using deletion protection](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html#WorkingWithTables.Basics.DeletionProtection) in the _Amazon DynamoDBDeveloper Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalSecondaryIndexes`

Global secondary indexes to be created on the table. You can create up to 20 global
secondary indexes.

###### Important

If you update a table to include a new global secondary index, AWS CloudFormation initiates the index creation and then proceeds with the stack update. AWS CloudFormation doesn't wait for the index to complete creation because the
backfilling phase can take a long time, depending on the size of the table. You can't
use the index or update the table until the index's status is `ACTIVE`. You
can track its status by using the DynamoDB [DescribeTable](https://docs.aws.amazon.com/cli/latest/reference/dynamodb/describe-table.html)
command.

If you add or delete an index during an update, we recommend that you don't update
any other resources. If your stack fails to update and is rolled back while adding a new
index, you must manually delete the index.

Updates are not supported. The following are exceptions:

- If you update either the contributor insights specification or the provisioned
throughput values of global secondary indexes, you can update the table without
interruption.

- You can delete or add one global secondary index without interruption. If you
do both in the same update (for example, by changing the index's logical ID), the
update fails.

_Required_: No

_Type_: Array of [GlobalSecondaryIndex](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-globalsecondaryindex.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImportSourceSpecification`

Specifies the properties of data being imported from the S3 bucket source to the"
table.

###### Important

If you specify the `ImportSourceSpecification` property, and also specify
either the `StreamSpecification`, the `TableClass` property, the
`DeletionProtectionEnabled` property, or the `WarmThroughput`
property, the IAM entity creating/updating stack must have `UpdateTable`
permission.

_Required_: No

_Type_: [ImportSourceSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-importsourcespecification.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KeySchema`

Specifies the attributes that make up the primary key for the table. The attributes in
the `KeySchema` property must also be defined in the
`AttributeDefinitions` property.

_Required_: Yes

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-keyschema.html) of [KeySchema](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-keyschema.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`KinesisStreamSpecification`

The Kinesis Data Streams configuration for the specified table.

_Required_: No

_Type_: [KinesisStreamSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-kinesisstreamspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocalSecondaryIndexes`

Local secondary indexes to be created on the table. You can create up to 5 local
secondary indexes. Each index is scoped to a given hash key value. The size of each hash
key can be up to 10 gigabytes.

_Required_: No

_Type_: Array of [LocalSecondaryIndex](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-localsecondaryindex.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnDemandThroughput`

Sets the maximum number of read and write units for the specified on-demand table. If
you use this property, you must specify `MaxReadRequestUnits`,
`MaxWriteRequestUnits`, or both.

_Required_: No

_Type_: [OnDemandThroughput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-ondemandthroughput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PointInTimeRecoverySpecification`

The settings used to enable point in time recovery.

_Required_: No

_Type_: [PointInTimeRecoverySpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-pointintimerecoveryspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisionedThroughput`

Throughput for the specified table, which consists of values for
`ReadCapacityUnits` and `WriteCapacityUnits`. For more information
about the contents of a provisioned throughput structure, see [Amazon DynamoDB\
Table ProvisionedThroughput](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughput.html).

If you set `BillingMode` as `PROVISIONED`, you must specify this
property. If you set `BillingMode` as `PAY_PER_REQUEST`, you cannot
specify this property.

_Required_: Conditional

_Type_: [ProvisionedThroughput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-provisionedthroughput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourcePolicy`

An AWS resource-based policy document in JSON format that will be
attached to the table.

When you attach a resource-based policy while creating a table, the policy application
is _strongly consistent_.

The maximum size supported for a resource-based policy document is 20 KB. DynamoDB counts whitespaces when calculating the size of a policy against this
limit. For a full list of all considerations that apply for resource-based policies, see
[Resource-based\
policy considerations](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/rbac-considerations.html).

###### Note

You need to specify the `CreateTable` and
`PutResourcePolicy`
IAM actions for authorizing a user to create a table with a
resource-based policy.

_Required_: No

_Type_: [ResourcePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-resourcepolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SSESpecification`

Specifies the settings to enable server-side encryption.

_Required_: No

_Type_: [SSESpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-ssespecification.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`StreamSpecification`

The settings for the DynamoDB table stream, which captures changes to items stored
in the table. Including this property in your AWS CloudFormation template automatically enables streaming.

_Required_: No

_Type_: [StreamSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-streamspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableClass`

The table class of the new table. Valid values are `STANDARD` and
`STANDARD_INFREQUENT_ACCESS`.

_Required_: No

_Type_: String

_Allowed values_: `STANDARD | STANDARD_INFREQUENT_ACCESS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

A name for the table. If you don't specify a name, AWS CloudFormation generates a
unique physical ID and uses that ID for the table name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).

###### Important

If you specify a name, you cannot perform updates that require replacement of this
resource. You can perform updates that require no or some interruption. If you must
replace the resource, specify a new name.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeToLiveSpecification`

Specifies the Time to Live (TTL) settings for the table.

###### Note

For detailed information about the limits in DynamoDB, see [Limits in\
Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.

_Required_: No

_Type_: [TimeToLiveSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-timetolivespecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WarmThroughput`

Represents the warm throughput (in read units per second and write units per second)
for creating a table.

_Required_: No

_Type_: [WarmThroughput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-warmthroughput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myDynamoDBTable" }`

For the resource with the logical ID `myDynamoDBTable`, `Ref` will
return the DynamoDB table name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the DynamoDB table, such as
`arn:aws:dynamodb:us-east-2:123456789012:table/myDynamoDBTable`.

`StreamArn`

The ARN of the DynamoDB stream, such as
`arn:aws:dynamodb:us-east-1:123456789012:table/testddbstack-myDynamoDBTable-012A1SL7SMP5Q/stream/2015-11-30T20:10:00.000`.

###### Note

You must specify the `StreamSpecification` property to use this
attribute.

## Examples

- [DynamoDB Table with Local and Secondary Indexes](#aws-resource-dynamodb-table--examples--DynamoDB_Table_with_Local_and_Secondary_Indexes)

- [DynamoDB Table with a DependsOn Attribute](#aws-resource-dynamodb-table--examples--DynamoDB_Table_with_a_DependsOn_Attribute)

- [DynamoDB Table with Application Auto Scaling](#aws-resource-dynamodb-table--examples--DynamoDB_Table_with_Application_Auto_Scaling)

### DynamoDB Table with Local and Secondary Indexes

The following sample creates a DynamoDB table with `Album`,
`Artist`, `Sales`, `NumberOfSongs` as attributes.
The primary key includes the `Album` attribute as the hash key and
`Artist` attribute as the range key. The table also includes two global
and one secondary index. For querying the number of sales for a given artist, the
global secondary index uses the `Sales` attribute as the hash key and the
`Artist` attribute as the range key.

For querying the sales based on the number of songs, the global secondary index
uses the `NumberOfSongs` attribute as the hash key and the
`Sales` attribute as the range key.

For querying the sales of an album, the local secondary index uses the same hash
key as the table but uses the `Sales` attribute as the range key.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "myDynamoDBTable": {
            "Type": "AWS::DynamoDB::Table",
            "Properties": {
                "AttributeDefinitions": [
                    {
                        "AttributeName": "Album",
                        "AttributeType": "S"
                    },
                    {
                        "AttributeName": "Artist",
                        "AttributeType": "S"
                    },
                    {
                        "AttributeName": "Sales",
                        "AttributeType": "N"
                    },
                    {
                        "AttributeName": "NumberOfSongs",
                        "AttributeType": "N"
                    }
                ],
                "KeySchema": [
                    {
                        "AttributeName": "Album",
                        "KeyType": "HASH"
                    },
                    {
                        "AttributeName": "Artist",
                        "KeyType": "RANGE"
                    }
                ],
                "ProvisionedThroughput": {
                    "ReadCapacityUnits": "5",
                    "WriteCapacityUnits": "5"
                },
                "TableName": "myTableName",
                "GlobalSecondaryIndexes": [
                    {
                        "IndexName": "myGSI",
                        "KeySchema": [
                            {
                                "AttributeName": "Sales",
                                "KeyType": "HASH"
                            },
                            {
                                "AttributeName": "Artist",
                                "KeyType": "RANGE"
                            }
                        ],
                        "Projection": {
                            "NonKeyAttributes": [
                                "Album",
                                "NumberOfSongs"
                            ],
                            "ProjectionType": "INCLUDE"
                        },
                        "ProvisionedThroughput": {
                            "ReadCapacityUnits": "5",
                            "WriteCapacityUnits": "5"
                        }
                    },
                    {
                        "IndexName": "myGSI2",
                        "KeySchema": [
                            {
                                "AttributeName": "NumberOfSongs",
                                "KeyType": "HASH"
                            },
                            {
                                "AttributeName": "Sales",
                                "KeyType": "RANGE"
                            }
                        ],
                        "Projection": {
                            "NonKeyAttributes": [
                                "Album",
                                "Artist"
                            ],
                            "ProjectionType": "INCLUDE"
                        },
                        "ProvisionedThroughput": {
                            "ReadCapacityUnits": "5",
                            "WriteCapacityUnits": "5"
                        }
                    }
                ],
                "LocalSecondaryIndexes": [
                    {
                        "IndexName": "myLSI",
                        "KeySchema": [
                            {
                                "AttributeName": "Album",
                                "KeyType": "HASH"
                            },
                            {
                                "AttributeName": "Sales",
                                "KeyType": "RANGE"
                            }
                        ],
                        "Projection": {
                            "NonKeyAttributes": [
                                "Artist",
                                "NumberOfSongs"
                            ],
                            "ProjectionType": "INCLUDE"
                        }
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  myDynamoDBTable:
    Type: AWS::DynamoDB::Table
    Properties:
      AttributeDefinitions:
        - AttributeName: Album
          AttributeType: S
        - AttributeName: Artist
          AttributeType: S
        - AttributeName: Sales
          AttributeType: 'N'
        - AttributeName: NumberOfSongs
          AttributeType: 'N'
      KeySchema:
        - AttributeName: Album
          KeyType: HASH
        - AttributeName: Artist
          KeyType: RANGE
      ProvisionedThroughput:
        ReadCapacityUnits: '5'
        WriteCapacityUnits: '5'
      TableName: myTableName
      GlobalSecondaryIndexes:
        - IndexName: myGSI
          KeySchema:
            - AttributeName: Sales
              KeyType: HASH
            - AttributeName: Artist
              KeyType: RANGE
          Projection:
            NonKeyAttributes:
              - Album
              - NumberOfSongs
            ProjectionType: INCLUDE
          ProvisionedThroughput:
            ReadCapacityUnits: '5'
            WriteCapacityUnits: '5'
        - IndexName: myGSI2
          KeySchema:
            - AttributeName: NumberOfSongs
              KeyType: HASH
            - AttributeName: Sales
              KeyType: RANGE
          Projection:
            NonKeyAttributes:
              - Album
              - Artist
            ProjectionType: INCLUDE
          ProvisionedThroughput:
            ReadCapacityUnits: '5'
            WriteCapacityUnits: '5'
      LocalSecondaryIndexes:
        - IndexName: myLSI
          KeySchema:
            - AttributeName: Album
              KeyType: HASH
            - AttributeName: Sales
              KeyType: RANGE
          Projection:
            NonKeyAttributes:
              - Artist
              - NumberOfSongs
            ProjectionType: INCLUDE
```

### DynamoDB Table with a DependsOn Attribute

If you include multiple DynamoDB tables with indexes in a single template, you
must include dependencies so that the tables are created sequentially. DynamoDB
limits the number of tables with secondary indexes that are in the creating state. If
you create multiple tables with indexes at the same time, DynamoDB returns an error
and the stack operation fails.

The following sample assumes that the `myFirstDDBTable` table is
declared in the same template as the `mySecondDDBTable` table, and both
tables include a secondary index. The `mySecondDDBTable` table includes a
dependency on the `myFirstDDBTable` table so that AWS CloudFormation
creates the tables one at a time.

#### JSON

```json

{
    "mySecondDDBTable": {
        "Type": "AWS::DynamoDB::Table",
        "DependsOn": "myFirstDDBTable",
        "Properties": {
            "AttributeDefinitions": [
                {
                    "AttributeName": "ArtistId",
                    "AttributeType": "S"
                },
                {
                    "AttributeName": "Concert",
                    "AttributeType": "S"
                },
                {
                    "AttributeName": "TicketSales",
                    "AttributeType": "S"
                }
            ],
            "KeySchema": [
                {
                    "AttributeName": "ArtistId",
                    "KeyType": "HASH"
                },
                {
                    "AttributeName": "Concert",
                    "KeyType": "RANGE"
                }
            ],
            "ProvisionedThroughput": {
                "ReadCapacityUnits": {
                    "Ref": "ReadCapacityUnits"
                },
                "WriteCapacityUnits": {
                    "Ref": "WriteCapacityUnits"
                }
            },
            "GlobalSecondaryIndexes": [
                {
                    "IndexName": "myGSI",
                    "KeySchema": [
                        {
                            "AttributeName": "TicketSales",
                            "KeyType": "HASH"
                        }
                    ],
                    "Projection": {
                        "ProjectionType": "KEYS_ONLY"
                    },
                    "ProvisionedThroughput": {
                        "ReadCapacityUnits": {
                            "Ref": "ReadCapacityUnits"
                        },
                        "WriteCapacityUnits": {
                            "Ref": "WriteCapacityUnits"
                        }
                    }
                }
            ],
            "Tags": [
                {
                    "Key": "foo",
                    "Value": "bar"
                }
            ]
        }
    }
}
```

#### YAML

```yaml

mySecondDDBTable:
  Type: AWS::DynamoDB::Table
  DependsOn: myFirstDDBTable
  Properties:
    AttributeDefinitions:
      - AttributeName: ArtistId
        AttributeType: S
      - AttributeName: Concert
        AttributeType: S
      - AttributeName: TicketSales
        AttributeType: S
    KeySchema:
      - AttributeName: ArtistId
        KeyType: HASH
      - AttributeName: Concert
        KeyType: RANGE
    ProvisionedThroughput:
      ReadCapacityUnits: !Ref ReadCapacityUnits
      WriteCapacityUnits: !Ref WriteCapacityUnits
    GlobalSecondaryIndexes:
      - IndexName: myGSI
        KeySchema:
          - AttributeName: TicketSales
            KeyType: HASH
        Projection:
          ProjectionType: KEYS_ONLY
        ProvisionedThroughput:
          ReadCapacityUnits: !Ref ReadCapacityUnits
          WriteCapacityUnits: !Ref WriteCapacityUnits
    Tags:
      - Key: foo
        Value: bar
```

### DynamoDB Table with Application Auto Scaling

This example sets up Application Auto Scaling for a
`AWS::DynamoDB::Table` resource. The template defines a
`TargetTrackingScaling` scaling policy that scales up the
`WriteCapacityUnits` throughput for the table.

#### JSON

```json

{
  "Resources": {
    "DDBTable": {
      "Type": "AWS::DynamoDB::Table",
      "Properties": {
        "AttributeDefinitions": [
          {
            "AttributeName": "ArtistId",
            "AttributeType": "S"
          },
          {
            "AttributeName": "Concert",
            "AttributeType": "S"
          },
          {
            "AttributeName": "TicketSales",
            "AttributeType": "S"
          }
        ],
        "KeySchema": [
          {
            "AttributeName": "ArtistId",
            "KeyType": "HASH"
          },
          {
            "AttributeName": "Concert",
            "KeyType": "RANGE"
          }
        ],
        "GlobalSecondaryIndexes": [
          {
            "IndexName": "GSI",
            "KeySchema": [
              {
                "AttributeName": "TicketSales",
                "KeyType": "HASH"
              }
            ],
            "Projection": {
              "ProjectionType": "KEYS_ONLY"
            },
            "ProvisionedThroughput": {
              "ReadCapacityUnits": 5,
              "WriteCapacityUnits": 5
            }
          }
        ],
        "ProvisionedThroughput": {
          "ReadCapacityUnits": 5,
          "WriteCapacityUnits": 5
        }
      }
    },
    "WriteCapacityScalableTarget": {
      "Type": "AWS::ApplicationAutoScaling::ScalableTarget",
      "Properties": {
        "MaxCapacity": 15,
        "MinCapacity": 5,
        "ResourceId": { "Fn::Join": [
          "/",
          [
            "table",
            { "Ref": "DDBTable" }
          ]
        ] },
        "RoleARN": {
          "Fn::GetAtt": ["ScalingRole", "Arn"]
        },
        "ScalableDimension": "dynamodb:table:WriteCapacityUnits",
        "ServiceNamespace": "dynamodb"
      }
    },
    "ScalingRole": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Effect": "Allow",
              "Principal": {
                "Service": [
                  "application-autoscaling.amazonaws.com"
                ]
              },
              "Action": [
                "sts:AssumeRole"
              ]
            }
          ]
        },
        "Path": "/",
        "Policies": [
          {
            "PolicyName": "root",
            "PolicyDocument": {
              "Version": "2012-10-17",
              "Statement": [
                {
                  "Effect": "Allow",
                  "Action": [
                    "dynamodb:DescribeTable",
                    "dynamodb:UpdateTable",
                    "cloudwatch:PutMetricAlarm",
                    "cloudwatch:DescribeAlarms",
                    "cloudwatch:GetMetricStatistics",
                    "cloudwatch:SetAlarmState",
                    "cloudwatch:DeleteAlarms"
                  ],
                  "Resource": "*"
                }
              ]
            }
          }
        ]
      }
    },
    "WriteScalingPolicy": {
      "Type": "AWS::ApplicationAutoScaling::ScalingPolicy",
      "Properties": {
        "PolicyName": "WriteAutoScalingPolicy",
        "PolicyType": "TargetTrackingScaling",
        "ScalingTargetId": {
          "Ref": "WriteCapacityScalableTarget"
        },
        "TargetTrackingScalingPolicyConfiguration": {
          "TargetValue": 50.0,
          "ScaleInCooldown": 60,
          "ScaleOutCooldown": 60,
          "PredefinedMetricSpecification": {
            "PredefinedMetricType": "DynamoDBWriteCapacityUtilization"
          }
        }
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  DDBTable:
    Type: AWS::DynamoDB::Table
    Properties:
      AttributeDefinitions:
        -
          AttributeName: "ArtistId"
          AttributeType: "S"
        -
          AttributeName: "Concert"
          AttributeType: "S"
        -
          AttributeName: "TicketSales"
          AttributeType: "S"
      KeySchema:
        -
          AttributeName: "ArtistId"
          KeyType: "HASH"
        -
          AttributeName: "Concert"
          KeyType: "RANGE"
      GlobalSecondaryIndexes:
        -
          IndexName: "GSI"
          KeySchema:
            -
              AttributeName: "TicketSales"
              KeyType: "HASH"
          Projection:
            ProjectionType: "KEYS_ONLY"
          ProvisionedThroughput:
            ReadCapacityUnits: 5
            WriteCapacityUnits: 5
      ProvisionedThroughput:
        ReadCapacityUnits: 5
        WriteCapacityUnits: 5
  WriteCapacityScalableTarget:
    Type: AWS::ApplicationAutoScaling::ScalableTarget
    Properties:
      MaxCapacity: 15
      MinCapacity: 5
      ResourceId: !Join
        - /
        - - table
          - !Ref DDBTable
      RoleARN: !GetAtt ScalingRole.Arn
      ScalableDimension: dynamodb:table:WriteCapacityUnits
      ServiceNamespace: dynamodb
  ScalingRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
          -
            Effect: "Allow"
            Principal:
              Service:
                - application-autoscaling.amazonaws.com
            Action:
              - "sts:AssumeRole"
      Path: "/"
      Policies:
        -
          PolicyName: "root"
          PolicyDocument:
            Version: "2012-10-17"
            Statement:
              -
                Effect: "Allow"
                Action:
                  - "dynamodb:DescribeTable"
                  - "dynamodb:UpdateTable"
                  - "cloudwatch:PutMetricAlarm"
                  - "cloudwatch:DescribeAlarms"
                  - "cloudwatch:GetMetricStatistics"
                  - "cloudwatch:SetAlarmState"
                  - "cloudwatch:DeleteAlarms"
                Resource: "*"
  WriteScalingPolicy:
    Type: AWS::ApplicationAutoScaling::ScalingPolicy
    Properties:
      PolicyName: WriteAutoScalingPolicy
      PolicyType: TargetTrackingScaling
      ScalingTargetId: !Ref WriteCapacityScalableTarget
      TargetTrackingScalingPolicyConfiguration:
        TargetValue: 50.0
        ScaleInCooldown: 60
        ScaleOutCooldown: 60
        PredefinedMetricSpecification:
          PredefinedMetricType: DynamoDBWriteCapacityUtilization
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WriteProvisionedThroughputSettings

AttributeDefinition
