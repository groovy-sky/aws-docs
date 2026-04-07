This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Table

You can use the `AWS::Cassandra::Table` resource to create a new table in
Amazon Keyspaces (for Apache Cassandra). For more information, see [Create a table](https://docs.aws.amazon.com/keyspaces/latest/devguide/getting-started.tables.html) in the _Amazon Keyspaces Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cassandra::Table",
  "Properties" : {
      "AutoScalingSpecifications" : AutoScalingSpecification,
      "BillingMode" : BillingMode,
      "CdcSpecification" : CdcSpecification,
      "ClientSideTimestampsEnabled" : Boolean,
      "ClusteringKeyColumns" : [ ClusteringKeyColumn, ... ],
      "DefaultTimeToLive" : Integer,
      "EncryptionSpecification" : EncryptionSpecification,
      "KeyspaceName" : String,
      "PartitionKeyColumns" : [ Column, ... ],
      "PointInTimeRecoveryEnabled" : Boolean,
      "RegularColumns" : [ Column, ... ],
      "ReplicaSpecifications" : [ ReplicaSpecification, ... ],
      "TableName" : String,
      "Tags" : [ Tag, ... ],
      "WarmThroughput" : WarmThroughput
    }
}

```

### YAML

```yaml

Type: AWS::Cassandra::Table
Properties:
  AutoScalingSpecifications:
    AutoScalingSpecification
  BillingMode:
    BillingMode
  CdcSpecification:
    CdcSpecification
  ClientSideTimestampsEnabled: Boolean
  ClusteringKeyColumns:
    - ClusteringKeyColumn
  DefaultTimeToLive: Integer
  EncryptionSpecification:
    EncryptionSpecification
  KeyspaceName: String
  PartitionKeyColumns:
    - Column
  PointInTimeRecoveryEnabled: Boolean
  RegularColumns:
    - Column
  ReplicaSpecifications:
    - ReplicaSpecification
  TableName: String
  Tags:
    - Tag
  WarmThroughput:
    WarmThroughput

```

## Properties

`AutoScalingSpecifications`

The optional auto scaling capacity settings for a table in provisioned capacity mode.

_Required_: No

_Type_: [AutoScalingSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cassandra-table-autoscalingspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BillingMode`

The billing mode for the table, which determines how you'll be charged for reads and writes:

- **On-demand mode** (default) - You pay based on
the actual reads and writes your application performs.

- **Provisioned mode** \- Lets you specify the
number of reads and writes per second that you need for your application.

If you don't specify a value for this property, then the table will use on-demand mode.

_Required_: No

_Type_: [BillingMode](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cassandra-table-billingmode.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CdcSpecification`

The settings for the CDC stream of a table. For more information about CDC streams, see [Working with change data capture (CDC) streams in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/cdc.html) in the _Amazon Keyspaces Developer_
_Guide_.

_Required_: No

_Type_: [CdcSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cassandra-table-cdcspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientSideTimestampsEnabled`

Enables client-side timestamps for the table. By default, the setting is disabled.
You can enable client-side timestamps with the following option:

- `status: "enabled"`

After client-side timestamps are enabled for a table, you can't disable this setting.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClusteringKeyColumns`

One or more columns that determine how the table data is sorted.

_Required_: No

_Type_: Array of [ClusteringKeyColumn](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cassandra-table-clusteringkeycolumn.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DefaultTimeToLive`

The default Time To Live (TTL) value for all rows in a table in seconds.
The maximum configurable value is 630,720,000 seconds, which is the equivalent of 20 years. By default, the TTL value for a table is 0, which means data does not expire.

For more information,
see [Setting the default TTL value for a table](https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL-how-it-works.html#ttl-howitworks_default_ttl)
in the _Amazon Keyspaces Developer Guide_.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionSpecification`

The encryption at rest options for the table.

- **AWS owned key** (default) - The key is owned by Amazon Keyspaces.

- **Customer managed key** \- The key is stored in your account and is created, owned, and
managed by you.

###### Note

If you choose encryption with a customer managed key, you must specify
a valid customer managed KMS key with permissions granted to Amazon
Keyspaces.

For more information,
see [Encryption at rest in Amazon Keyspaces](https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html)
in the _Amazon Keyspaces Developer Guide_.

_Required_: No

_Type_: [EncryptionSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cassandra-table-encryptionspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyspaceName`

The name of the keyspace to create the table in. The keyspace must already
exist.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_]{1,47}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PartitionKeyColumns`

One or more columns that uniquely identify every row in the table. Every table must
have a partition key.

_Required_: Yes

_Type_: Array of [Column](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cassandra-table-column.html)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PointInTimeRecoveryEnabled`

Specifies if point-in-time recovery is enabled or disabled for the table. The options are `PointInTimeRecoveryEnabled=true` and
`PointInTimeRecoveryEnabled=false`.
If not specified, the default is `PointInTimeRecoveryEnabled=false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegularColumns`

One or more columns that are not part of the primary key - that is, columns that are
_not_ defined as partition key columns or clustering key
columns.

You can add regular columns to existing tables by adding them to the template.

_Required_: No

_Type_: Array of [Column](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cassandra-table-column.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicaSpecifications`

The AWS Region specific settings of a multi-Region table.

For a multi-Region table, you can configure the table's read capacity differently per AWS Region. You can do this by configuring the following parameters.

- `region`: The Region where these settings are applied. (Required)

- `readCapacityUnits`: The provisioned read capacity units. (Optional)

- `readCapacityAutoScaling`: The read capacity auto scaling settings for the table. (Optional)

_Required_: No

_Type_: Array of [ReplicaSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cassandra-table-replicaspecification.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The name of the table to be created. The table name is case sensitive. If you don't specify a name, AWS CloudFormation
generates a unique ID and uses that ID for the table name. For more information, see
[Name\
type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).

###### Important

If you specify a name, you can't perform updates that require replacing this
resource. You can perform updates that require no interruption or some interruption. If you must
replace the resource, specify a new name.

_Length constraints:_ Minimum length of 3. Maximum length of
255.

_Pattern:_ `^[a-zA-Z0-9][a-zA-Z0-9_]{1,47}$`

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_]{1,47}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cassandra-table-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WarmThroughput`

Property description not available.

_Required_: No

_Type_: [WarmThroughput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cassandra-table-warmthroughput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the table and the keyspace where the table exists (delimited by '\|'). For example:

`{ "Ref": "myKeyspace|myTable" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

- [Create a table with minimal options](#aws-resource-cassandra-table--examples--Create_a_table_with_minimal_options)

- [Create a table with frozen collections](#aws-resource-cassandra-table--examples--Create_a_table_with_frozen_collections)

- [Create a table with static columns](#aws-resource-cassandra-table--examples--Create_a_table_with_static_columns)

- [Create a table with a stream](#aws-resource-cassandra-table--examples--Create_a_table_with_a_stream)

- [Create a table with a stream and tags](#aws-resource-cassandra-table--examples--Create_a_table_with_a_stream_and_tags)

- [Create a table with client-side timestamps and other options](#aws-resource-cassandra-table--examples--Create_a_table_with_client-side_timestamps_and_other_options)

- [Create a single Region table in provisioned capacity mode with auto scaling](#aws-resource-cassandra-table--examples--Create_a_single_Region_table_in_provisioned_capacity_mode_with_auto_scaling)

- [Create a multi-Region table in provisioned capacity mode with auto scaling](#aws-resource-cassandra-table--examples--Create_a_multi-Region_table_in_provisioned_capacity_mode_with_auto_scaling)

- [Create a table with customer managed keys and other options](#aws-resource-cassandra-table--examples--Create_a_table_with_customer_managed_keys_and_other_options)

- [Add new columns to an existing table](#aws-resource-cassandra-table--examples--Add_new_columns_to_an_existing_table)

### Create a table with minimal options

The following example creates a new table.
The table has the name `my_table`, and uses on-demand billing.

#### JSON

```JSON

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "myNewTable": {
      "Type": "AWS::Cassandra::Table",
      "Properties": {
        "KeyspaceName": "my_keyspace",
        "TableName":"my_table",
        "PartitionKeyColumns": [
          {
            "ColumnName": "Message",
            "ColumnType": "ASCII"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  myNewTable:
    Type: 'AWS::Cassandra::Table'
    Properties:
      KeyspaceName: my_keyspace
      TableName: my_table
      PartitionKeyColumns:
        - ColumnName: Message
          ColumnType: ASCII
```

### Create a table with frozen collections

The following example creates a new table with the name `my_table` where the columns `projects`, `addresses`,
and `org_members_by_dept` use frozen collections.

#### JSON

```JSON

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "myTable": {
      "Type": "AWS::Cassandra::Table",
      "Properties": {
        "KeyspaceName": "my_keyspace",
        "TableName":"my_table",
        "PartitionKeyColumns": [
          {
            "ColumnName": "message",
            "ColumnType": "ASCII"
          }
        ],
        "RegularColumns": [
          {
            "ColumnName": "name",
            "ColumnType": "TEXT"
          },
          {
            "ColumnName": "region",
            "ColumnType": "TEXT"
          },
          {
            "ColumnName": "projects",
            "ColumnType": "FROZEN<SET<TEXT>>"
          },
          {
            "ColumnName": "role",
            "ColumnType": "TEXT"
          },
          {
            "ColumnName": "addresses",
            "ColumnType": "FROZEN<MAP<TEXT, SET<TEXT>>>"
          },
          {
            "ColumnName": "pay_scale",
            "ColumnType": "TEXT"
          },
          {
            "ColumnName": "vacation_hrs",
            "ColumnType": "FLOAT"
          },
          {
            "ColumnName": "manager_id",
            "ColumnType": "TEXT"
          },
          {
            "ColumnName": "org_members_by_dept",
            "ColumnType": "MAP<INT, FROZEN<LIST<TEXT>>>"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  myTable:
    Type: 'AWS::Cassandra::Table'
    Properties:
      KeyspaceName: my_keyspace
      TableName: my_table
      PartitionKeyColumns:
        - ColumnName: message
          ColumnType: ASCII
      RegularColumns:
        - ColumnName: name
          ColumnType: TEXT
        - ColumnName: region
          ColumnType: TEXT
        - ColumnName: projects
          ColumnType: FROZEN<SET<TEXT>>
        - ColumnName: role
          ColumnType: TEXT
        - ColumnName: addresses
          ColumnType: FROZEN<MAP<TEXT, SET<TEXT>>>
        - ColumnName: pay_scale
          ColumnType: TEXT
        - ColumnName: vacation_hrs
          ColumnType: FLOAT
        - ColumnName: manager_id
          ColumnType: TEXT
        - ColumnName: org_members_by_dept
          ColumnType: MAP<INT, FROZEN<LIST<TEXT>>>
```

### Create a table with static columns

The following example creates a new table with the name `my_table`.
The columns `company_name` and `company_id` are static columns.

#### JSON

```JSON

{
   "AWSTemplateFormatVersion":"2010-09-09",
   "Resources":{
      "myNewTable":{
         "Type":"AWS::Cassandra::Table",
         "Properties":{
            "KeyspaceName":"my_keyspace",
            "TableName":"my_table",
            "PartitionKeyColumns":[
               {
                  "ColumnName":"id",
                  "ColumnType":"ASCII"
               }
            ],
            "ClusteringKeyColumns":[
               {
                  "Column":{
                     "ColumnName":"division",
                     "ColumnType":"ASCII"
                  },
                  "OrderBy":"ASC"
               }
            ],
            "RegularColumns":[
               {
                  "ColumnName":"name",
                  "ColumnType":"TEXT"
               },
               {
                  "ColumnName":"company_name",
                  "ColumnType":"TEXT STATIC"
               },
               {
                  "ColumnName":"company_id",
                  "ColumnType":"INT STATIC"
               },
               {
                  "ColumnName":"region",
                  "ColumnType":"TEXT"
               },
               {
                  "ColumnName":"project",
                  "ColumnType":"TEXT"
               },
               {
                  "ColumnName":"role",
                  "ColumnType":"TEXT"
               },
               {
                  "ColumnName":"pay_scale",
                  "ColumnType":"TEXT"
               },
               {
                  "ColumnName":"vacation_hrs",
                  "ColumnType":"FLOAT"
               },
               {
                  "ColumnName":"manager_id",
                  "ColumnType":"TEXT"
               }
            ],
            "BillingMode":{
               "Mode":"PROVISIONED",
               "ProvisionedThroughput":{
                  "ReadCapacityUnits":5,
                  "WriteCapacityUnits":5
               }
            },
            "ClientSideTimestampsEnabled":true,
            "DefaultTimeToLive":63072000,
            "PointInTimeRecoveryEnabled":true,
            "Tags":[
              {
                 "Key":"tag1",
                 "Value":"val1"
              },
              {
                 "Key":"tag2",
                 "Value":"val2"
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
  myNewTable:
    Type: AWS::Cassandra::Table
    Properties:
      KeyspaceName: my_keyspace
      TableName: my_table
      PartitionKeyColumns:
      - ColumnName: id
        ColumnType: ASCII
      ClusteringKeyColumns:
      - Column:
          ColumnName: division
          ColumnType: ASCII
        OrderBy: ASC
      RegularColumns:
      - ColumnName: name
        ColumnType: TEXT
      - ColumnName: company_name
        ColumnType: TEXT STATIC
      - ColumnName: company_id
        ColumnType: INT STATIC
      - ColumnName: region
        ColumnType: TEXT
      - ColumnName: project
        ColumnType: TEXT
      - ColumnName: role
        ColumnType: TEXT
      - ColumnName: pay_scale
        ColumnType: TEXT
      - ColumnName: vacation_hrs
        ColumnType: FLOAT
      - ColumnName: manager_id
        ColumnType: TEXT
      BillingMode:
        Mode: PROVISIONED
        ProvisionedThroughput:
          ReadCapacityUnits: 5
          WriteCapacityUnits: 5
      ClientSideTimestampsEnabled: true
      DefaultTimeToLive: 63072000
      PointInTimeRecoveryEnabled: true
      Tags:
        - Key: tag1
          Value: val1
        - Key: tag2
          Value: val2
```

### Create a table with a stream

The following example creates a new table.
The table has the name `my_table`, and the table has a stream.

#### JSON

```JSON

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "myNewTable": {
      "Type": "AWS::Cassandra::Table",
      "Properties": {
        "KeyspaceName": "my_keyspace",
        "TableName":"my_table",
        "PartitionKeyColumns": [
          {
            "ColumnName": "Message",
            "ColumnType": "ASCII"
          }
        ],
        "CdcSpecification": {
          "Status": "ENABLED",
          "ViewType": "NEW_AND_OLD_IMAGES"
        }
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  myNewTable:
    Type: 'AWS::Cassandra::Table'
    Properties:
      KeyspaceName: my_keyspace
      TableName: my_table
      PartitionKeyColumns:
        - ColumnName: Message
          ColumnType: ASCII
      CdcSpecification:
        Status: ENABLED
        ViewType: NEW_AND_OLD_IMAGES
```

### Create a table with a stream and tags

The following example creates a new table.
The table has the name `my_table`, and the table has a stream with two tags.

#### JSON

```JSON

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "myNewTable": {
      "Type": "AWS::Cassandra::Table",
      "Properties": {
        "KeyspaceName": "my_keyspace",
        "TableName":"my_table",
        "PartitionKeyColumns": [
          {
            "ColumnName": "Message",
            "ColumnType": "ASCII"
          }
        ],
        "CdcSpecification": {
          "Status": "ENABLED",
          "ViewType": "NEW_AND_OLD_IMAGES",
          "Tags":[
            {
               "Key":"tag1",
               "Value":"val1"
            },
            {
               "Key":"tag2",
               "Value":"val2"
            }
        }
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  myNewTable:
    Type: 'AWS::Cassandra::Table'
    Properties:
      KeyspaceName: my_keyspace
      TableName: my_table
      PartitionKeyColumns:
        - ColumnName: Message
          ColumnType: ASCII
      CdcSpecification:
        Status: ENABLED
        ViewType: NEW_AND_OLD_IMAGES
        Tags:
          - Key: tag1
            Value: val1
          - Key: tag2
            Value: val2
```

### Create a table with client-side timestamps and other options

The following example creates a table `my_table` with client-side
timestamps, provisioned read and write capacity,
default TTL, PITR, and tags.

#### JSON

```JSON

{
   "AWSTemplateFormatVersion":"2010-09-09",
   "Resources":{
      "myNewTable":{
         "Type":"AWS::Cassandra::Table",
         "Properties":{
            "KeyspaceName":"my_keyspace",
            "TableName":"my_table",
            "PartitionKeyColumns":[
               {
                  "ColumnName":"id",
                  "ColumnType":"ASCII"
               }
            ],
            "ClusteringKeyColumns":[
               {
                  "Column":{
                     "ColumnName":"division",
                     "ColumnType":"ASCII"
                  },
                  "OrderBy":"ASC"
               }
            ],
            "RegularColumns":[
               {
                  "ColumnName":"name",
                  "ColumnType":"TEXT"
               },
               {
                  "ColumnName":"region",
                  "ColumnType":"TEXT"
               },
               {
                  "ColumnName":"project",
                  "ColumnType":"TEXT"
               },
               {
                  "ColumnName":"role",
                  "ColumnType":"TEXT"
               },
               {
                  "ColumnName":"pay_scale",
                  "ColumnType":"TEXT"
               },
               {
                  "ColumnName":"vacation_hrs",
                  "ColumnType":"FLOAT"
               },
               {
                  "ColumnName":"manager_id",
                  "ColumnType":"TEXT"
               }
            ],
            "BillingMode":{
               "Mode":"PROVISIONED",
               "ProvisionedThroughput":{
                  "ReadCapacityUnits":5,
                  "WriteCapacityUnits":5
               }
            },
            "ClientSideTimestampsEnabled":true,
            "DefaultTimeToLive":63072000,
            "PointInTimeRecoveryEnabled":true,
            "Tags":[
              {
                 "Key":"tag1",
                 "Value":"val1"
              },
              {
                 "Key":"tag2",
                 "Value":"val2"
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
  myNewTable:
    Type: AWS::Cassandra::Table
    Properties:
      KeyspaceName: my_keyspace
      TableName: my_table
      PartitionKeyColumns:
      - ColumnName: id
        ColumnType: ASCII
      ClusteringKeyColumns:
      - Column:
          ColumnName: division
          ColumnType: ASCII
        OrderBy: ASC
      RegularColumns:
      - ColumnName: name
        ColumnType: TEXT
      - ColumnName: region
        ColumnType: TEXT
      - ColumnName: project
        ColumnType: TEXT
      - ColumnName: role
        ColumnType: TEXT
      - ColumnName: pay_scale
        ColumnType: TEXT
      - ColumnName: vacation_hrs
        ColumnType: FLOAT
      - ColumnName: manager_id
        ColumnType: TEXT
      BillingMode:
        Mode: PROVISIONED
        ProvisionedThroughput:
          ReadCapacityUnits: 5
          WriteCapacityUnits: 5
      ClientSideTimestampsEnabled: true
      DefaultTimeToLive: 63072000
      PointInTimeRecoveryEnabled: true
      Tags:
        - Key: tag1
          Value: val1
        - Key: tag2
          Value: val2
```

### Create a single Region table in provisioned capacity mode with auto scaling

The following example creates a table `my_table` with provisioned read and write capacity
and auto scaling enabled.

#### JSON

```JSON

{
  "AWSTemplateFormatVersion" : "2010-09-09",
  "Resources" : {
    "myNewTable" : {
      "Type" : "AWS::Cassandra::Table",
      "Properties" : {
        "KeyspaceName" : "my_keyspace",
        "TableName" : "my_table",
        "PartitionKeyColumns" : [ {
          "ColumnName" : "Message",
          "ColumnType" : "ASCII"
        } ],
        "BillingMode" : {
          "Mode" : "PROVISIONED",
          "ProvisionedThroughput" : {
            "ReadCapacityUnits" : 5,
            "WriteCapacityUnits" : 5
          }
        },
        "AutoScalingSpecifications" : {
          "WriteCapacityAutoScaling" : {
            "AutoScalingDisabled" : false,
            "MinimumUnits" : 5,
            "MaximumUnits" : 10,
            "ScalingPolicy" : {
              "TargetTrackingScalingPolicyConfiguration" : {
                "ScaleInCooldown" : 60,
                "ScaleOutCooldown" : 60,
                "TargetValue" : 70,
                "DisableScaleIn" : false
              }
            }
          },
          "ReadCapacityAutoScaling" : {
            "AutoScalingDisabled" : false,
            "MinimumUnits" : 5,
            "MaximumUnits" : 10,
            "ScalingPolicy" : {
              "TargetTrackingScalingPolicyConfiguration" : {
                "ScaleInCooldown" : 60,
                "ScaleOutCooldown" : 60,
                "TargetValue" : 70,
                "DisableScaleIn" : false
              }
            }
          }
        }
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  myNewTable:
    Type: 'AWS::Cassandra::Table'
    Properties:
      KeyspaceName: my_keyspace
      TableName: my_table
      PartitionKeyColumns:
        - ColumnName: Message
          ColumnType: ASCII
      BillingMode:
        Mode: PROVISIONED
        ProvisionedThroughput:
          ReadCapacityUnits: 5
          WriteCapacityUnits: 5
      AutoScalingSpecifications:
        WriteCapacityAutoScaling:
          AutoScalingDisabled: false
          MinimumUnits: 5
          MaximumUnits: 10
          ScalingPolicy:
            TargetTrackingScalingPolicyConfiguration:
              ScaleInCooldown: 60
              ScaleOutCooldown: 60
              TargetValue: 70
              DisableScaleIn: false
        ReadCapacityAutoScaling:
          AutoScalingDisabled: false
          MinimumUnits: 5
          MaximumUnits: 10
          ScalingPolicy:
            TargetTrackingScalingPolicyConfiguration:
              ScaleInCooldown: 60
              ScaleOutCooldown: 60
              TargetValue: 70
              DisableScaleIn: false
```

### Create a multi-Region table in provisioned capacity mode with auto scaling

The following example creates a table `my_table` with provisioned read and write capacity
and auto scaling enabled. In this example, first the default auto scaling settings are defined. Then using
`ReplicaSpecifications`, different auto scaling settings that override the default settings are defined for the Regions
`us-east-1` and `eu-west-1`.

#### JSON

```JSON

{
  "AWSTemplateFormatVersion" : "2010-09-09",
  "Resources" : {
    "myNewTable" : {
      "Type" : "AWS::Cassandra::Table",
      "Properties" : {
        "KeyspaceName" : "my_keyspace",
        "TableName" : "my_table",
        "PartitionKeyColumns" : [ {
          "ColumnName" : "Message",
          "ColumnType" : "ASCII"
        } ],
        "BillingMode" : {
          "Mode" : "PROVISIONED",
          "ProvisionedThroughput" : {
            "ReadCapacityUnits" : 5,
            "WriteCapacityUnits" : 5
          }
        },
        "AutoScalingSpecifications" : {
          "WriteCapacityAutoScaling" : {
            "AutoScalingDisabled" : false,
            "MinimumUnits" : 5,
            "MaximumUnits" : 10,
            "ScalingPolicy" : {
              "TargetTrackingScalingPolicyConfiguration" : {
                "ScaleInCooldown" : 60,
                "ScaleOutCooldown" : 60,
                "TargetValue" : 70
              }
            }
          },
          "ReadCapacityAutoScaling" : {
            "AutoScalingDisabled" : false,
            "MinimumUnits" : 5,
            "MaximumUnits" : 10,
            "ScalingPolicy" : {
              "TargetTrackingScalingPolicyConfiguration" : {
                "ScaleInCooldown" : 60,
                "ScaleOutCooldown" : 60,
                "TargetValue" : 70
              }
            }
          }
        },
        "ReplicaSpecifications" : [ {
          "Region" : "us-east-1",
          "ReadCapacityAutoScaling" : {
            "AutoScalingDisabled" : false,
            "MinimumUnits" : 5,
            "MaximumUnits" : 20,
            "ScalingPolicy" : {
              "TargetTrackingScalingPolicyConfiguration" : {
                "ScaleInCooldown" : 60,
                "ScaleOutCooldown" : 60,
                "TargetValue" : 75,
                "DisableScaleIn" : false
              }
            }
          }
        }, {
          "Region" : "eu-west-1",
          "ReadCapacityAutoScaling" : {
            "AutoScalingDisabled" : false,
            "MinimumUnits" : 5,
            "MaximumUnits" : 50,
            "ScalingPolicy" : {
              "TargetTrackingScalingPolicyConfiguration" : {
                "ScaleInCooldown" : 60,
                "ScaleOutCooldown" : 60,
                "TargetValue" : 65,
                "DisableScaleIn" : false
              }
            }
          }
        } ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  myNewTable:
    Type: 'AWS::Cassandra::Table'
    Properties:
      KeyspaceName: my_keyspace
      TableName: my_table
      PartitionKeyColumns:
        - ColumnName: Message
          ColumnType: ASCII
      BillingMode:
        Mode: PROVISIONED
        ProvisionedThroughput:
          ReadCapacityUnits: 5
          WriteCapacityUnits: 5
      AutoScalingSpecifications:
        WriteCapacityAutoScaling:
          AutoScalingDisabled: false
          MinimumUnits: 5
          MaximumUnits: 10
          ScalingPolicy:
            TargetTrackingScalingPolicyConfiguration:
              ScaleInCooldown: 60
              ScaleOutCooldown: 60
              TargetValue: 70
        ReadCapacityAutoScaling:
          AutoScalingDisabled: false
          MinimumUnits: 5
          MaximumUnits: 10
          ScalingPolicy:
            TargetTrackingScalingPolicyConfiguration:
              ScaleInCooldown: 60
              ScaleOutCooldown: 60
              TargetValue: 70
      ReplicaSpecifications:
        - Region: us-east-1
          ReadCapacityAutoScaling:
            AutoScalingDisabled: false
            MinimumUnits: 5
            MaximumUnits: 20
            ScalingPolicy:
              TargetTrackingScalingPolicyConfiguration:
                ScaleInCooldown: 60
                ScaleOutCooldown: 60
                TargetValue: 75
                DisableScaleIn: false
        - Region: eu-west-1
          ReadCapacityAutoScaling:
            AutoScalingDisabled: false
            MinimumUnits: 5
            MaximumUnits: 50
            ScalingPolicy:
              TargetTrackingScalingPolicyConfiguration:
                ScaleInCooldown: 60
                ScaleOutCooldown: 60
                TargetValue: 65
                DisableScaleIn: false
```

### Create a table with customer managed keys and other options

The following example creates a table `my_table` with customer managed
encryption keys and all other available options.
For example, provisioned read and write capacity,
default TTL, PITR, and tags. To use this sample, you must
replace the key ARN in the example with your own.

#### JSON

```JSON

{
   "AWSTemplateFormatVersion":"2010-09-09",
   "Resources":{
      "myNewTable":{
         "Type":"AWS::Cassandra::Table",
         "Properties":{
            "KeyspaceName":"my_keyspace",
            "TableName":"my_table",
            "PartitionKeyColumns":[
               {
                  "ColumnName":"id",
                  "ColumnType":"ASCII"
               }
            ],
            "ClusteringKeyColumns":[
               {
                  "Column":{
                     "ColumnName":"division",
                     "ColumnType":"ASCII"
                  },
                  "OrderBy":"ASC"
               }
            ],
            "RegularColumns":[
               {
                  "ColumnName":"name",
                  "ColumnType":"TEXT"
               },
               {
                  "ColumnName":"region",
                  "ColumnType":"TEXT"
               },
               {
                  "ColumnName":"project",
                  "ColumnType":"TEXT"
               },
               {
                  "ColumnName":"role",
                  "ColumnType":"TEXT"
               },
               {
                  "ColumnName":"pay_scale",
                  "ColumnType":"TEXT"
               },
               {
                  "ColumnName":"vacation_hrs",
                  "ColumnType":"FLOAT"
               },
               {
                  "ColumnName":"manager_id",
                  "ColumnType":"TEXT"
               }
            ],
            "BillingMode":{
               "Mode":"PROVISIONED",
               "ProvisionedThroughput":{
                  "ReadCapacityUnits":5,
                  "WriteCapacityUnits":5
               }
            },
            "DefaultTimeToLive":63072000,
            "EncryptionSpecification":{
               "EncryptionType":"CUSTOMER_MANAGED_KMS_KEY",
               "KmsKeyIdentifier":"arn:aws:kms:eu-west-1:5555555555555:key/11111111-1111-111-1111-111111111111"
            },
            "PointInTimeRecoveryEnabled":true,
            "Tags":[
               {
                  "Key":"tag1",
                  "Value":"val1"
               },
               {
                  "Key":"tag2",
                  "Value":"val2"
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
  myNewTable:
    Type: AWS::Cassandra::Table
    Properties:
      KeyspaceName: my_keyspace
      TableName: my_table
      PartitionKeyColumns:
      - ColumnName: id
        ColumnType: ASCII
      ClusteringKeyColumns:
      - Column:
          ColumnName: division
          ColumnType: ASCII
        OrderBy: ASC
      RegularColumns:
      - ColumnName: name
        ColumnType: TEXT
      - ColumnName: region
        ColumnType: TEXT
      - ColumnName: project
        ColumnType: TEXT
      - ColumnName: role
        ColumnType: TEXT
      - ColumnName: pay_scale
        ColumnType: TEXT
      - ColumnName: vacation_hrs
        ColumnType: FLOAT
      - ColumnName: manager_id
        ColumnType: TEXT
      BillingMode:
        Mode: PROVISIONED
        ProvisionedThroughput:
          ReadCapacityUnits: 5
          WriteCapacityUnits: 5
      DefaultTimeToLive: 63072000
      EncryptionSpecification:
        EncryptionType: CUSTOMER_MANAGED_KMS_KEY
        KmsKeyIdentifier: arn:aws:kms:eu-west-1:5555555555555:key/11111111-1111-111-1111-111111111111
      PointInTimeRecoveryEnabled: true
      Tags:
        - Key: tag1
          Value: val1
        - Key: tag2
          Value: val2
```

### Add new columns to an existing table

The following example shows how to add five new columns to the existing table `my_table`.
You can only add regular columns to a table.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "myTable": {
      "Type": "AWS::Cassandra::Table",
      "Properties": {
        "KeyspaceName": "my_keyspace",
        "TableName":"my_table",
        "PartitionKeyColumns": [
          {
            "ColumnName": "Message",
            "ColumnType": "ASCII"
          }
        ],
        "RegularColumns": [
          {
            "ColumnName": "name",
            "ColumnType": "TEXT"
          },
          {
            "ColumnName": "region",
            "ColumnType": "TEXT"
          }
        ]
      }
    }
  }
}
```

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "myTable": {
      "Type": "AWS::Cassandra::Table",
      "Properties": {
        "KeyspaceName": "my_keyspace",
        "TableName":"my_table",
        "PartitionKeyColumns": [
          {
            "ColumnName": "Message",
            "ColumnType": "ASCII"
          }
        ],
        "RegularColumns": [
          {
            "ColumnName": "name",
            "ColumnType": "TEXT"
          },
          {
            "ColumnName": "region",
            "ColumnType": "TEXT"
          },
          {
            "ColumnName": "project",
            "ColumnType": "TEXT"
          },
          {
            "ColumnName": "role",
            "ColumnType": "TEXT"
          },
          {
            "ColumnName": "pay_scale",
            "ColumnType": "TEXT"
          },
          {
            "ColumnName": "vacation_hrs",
            "ColumnType": "FLOAT"
          },
          {
            "ColumnName": "manager_id",
            "ColumnType": "TEXT"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  myTable:
    Type: 'AWS::Cassandra::Table'
    Properties:
      KeyspaceName: my_keyspace
      TableName: my_table
      PartitionKeyColumns:
        - ColumnName: Message
          ColumnType: ASCII
      RegularColumns:
        - ColumnName: name
          ColumnType: TEXT
        - ColumnName: region
          ColumnType: TEXT
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  myTable:
    Type: 'AWS::Cassandra::Table'
    Properties:
      KeyspaceName: my_keyspace
      TableName: my_table
      PartitionKeyColumns:
        - ColumnName: Message
          ColumnType: ASCII
      RegularColumns:
        - ColumnName: name
          ColumnType: TEXT
        - ColumnName: region
          ColumnType: TEXT
        - ColumnName: project
          ColumnType: TEXT
        - ColumnName: role
          ColumnType: TEXT
        - ColumnName: pay_scale
          ColumnType: TEXT
        - ColumnName: vacation_hrs
          ColumnType: FLOAT
        - ColumnName: manager_id
          ColumnType: TEXT
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AutoScalingSetting
