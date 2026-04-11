This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream

The `AWS::KinesisFirehose::DeliveryStream` resource specifies an Amazon
Kinesis Data Firehose (Kinesis Data Firehose) delivery stream that delivers real-time
streaming data to an Amazon Simple Storage Service (Amazon S3), Amazon Redshift, or Amazon
Elasticsearch Service (Amazon ES) destination. For more information, see [Creating an Amazon\
Kinesis Data Firehose Delivery Stream](../../../firehose/latest/dev/basic-create.md) in the _Amazon Kinesis Data_
_Firehose Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::KinesisFirehose::DeliveryStream",
  "Properties" : {
      "AmazonOpenSearchServerlessDestinationConfiguration" : AmazonOpenSearchServerlessDestinationConfiguration,
      "AmazonopensearchserviceDestinationConfiguration" : AmazonopensearchserviceDestinationConfiguration,
      "DatabaseSourceConfiguration" : DatabaseSourceConfiguration,
      "DeliveryStreamEncryptionConfigurationInput" : DeliveryStreamEncryptionConfigurationInput,
      "DeliveryStreamName" : String,
      "DeliveryStreamType" : String,
      "DirectPutSourceConfiguration" : DirectPutSourceConfiguration,
      "ElasticsearchDestinationConfiguration" : ElasticsearchDestinationConfiguration,
      "ExtendedS3DestinationConfiguration" : ExtendedS3DestinationConfiguration,
      "HttpEndpointDestinationConfiguration" : HttpEndpointDestinationConfiguration,
      "IcebergDestinationConfiguration" : IcebergDestinationConfiguration,
      "KinesisStreamSourceConfiguration" : KinesisStreamSourceConfiguration,
      "MSKSourceConfiguration" : MSKSourceConfiguration,
      "RedshiftDestinationConfiguration" : RedshiftDestinationConfiguration,
      "S3DestinationConfiguration" : S3DestinationConfiguration,
      "SnowflakeDestinationConfiguration" : SnowflakeDestinationConfiguration,
      "SplunkDestinationConfiguration" : SplunkDestinationConfiguration,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::KinesisFirehose::DeliveryStream
Properties:
  AmazonOpenSearchServerlessDestinationConfiguration:
    AmazonOpenSearchServerlessDestinationConfiguration
  AmazonopensearchserviceDestinationConfiguration:
    AmazonopensearchserviceDestinationConfiguration
  DatabaseSourceConfiguration:
    DatabaseSourceConfiguration
  DeliveryStreamEncryptionConfigurationInput:
    DeliveryStreamEncryptionConfigurationInput
  DeliveryStreamName: String
  DeliveryStreamType: String
  DirectPutSourceConfiguration:
    DirectPutSourceConfiguration
  ElasticsearchDestinationConfiguration:
    ElasticsearchDestinationConfiguration
  ExtendedS3DestinationConfiguration:
    ExtendedS3DestinationConfiguration
  HttpEndpointDestinationConfiguration:
    HttpEndpointDestinationConfiguration
  IcebergDestinationConfiguration:
    IcebergDestinationConfiguration
  KinesisStreamSourceConfiguration:
    KinesisStreamSourceConfiguration
  MSKSourceConfiguration:
    MSKSourceConfiguration
  RedshiftDestinationConfiguration:
    RedshiftDestinationConfiguration
  S3DestinationConfiguration:
    S3DestinationConfiguration
  SnowflakeDestinationConfiguration:
    SnowflakeDestinationConfiguration
  SplunkDestinationConfiguration:
    SplunkDestinationConfiguration
  Tags:
    - Tag

```

## Properties

`AmazonOpenSearchServerlessDestinationConfiguration`

Describes the configuration of a destination in the Serverless offering for Amazon
OpenSearch Service.

_Required_: No

_Type_: [AmazonOpenSearchServerlessDestinationConfiguration](aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AmazonopensearchserviceDestinationConfiguration`

The destination in Amazon OpenSearch Service. You can specify only one
destination.

_Required_: Conditional

_Type_: [AmazonopensearchserviceDestinationConfiguration](aws-properties-kinesisfirehose-deliverystream-amazonopensearchservicedestinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseSourceConfiguration`

The top level object for configuring streams with database as a source.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: No

_Type_: [DatabaseSourceConfiguration](aws-properties-kinesisfirehose-deliverystream-databasesourceconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeliveryStreamEncryptionConfigurationInput`

Specifies the type and Amazon Resource Name (ARN) of the CMK to use for Server-Side
Encryption (SSE).

_Required_: No

_Type_: [DeliveryStreamEncryptionConfigurationInput](aws-properties-kinesisfirehose-deliverystream-deliverystreamencryptionconfigurationinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeliveryStreamName`

The name of the Firehose stream.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9._-]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeliveryStreamType`

The Firehose stream type. This can be one of the following values:

- `DirectPut`: Provider applications access the Firehose stream
directly.

- `KinesisStreamAsSource`: The Firehose stream uses a Kinesis data
stream as a source.

_Required_: No

_Type_: String

_Allowed values_: `DatabaseAsSource | DirectPut | KinesisStreamAsSource | MSKAsSource`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DirectPutSourceConfiguration`

The structure that configures parameters such as `ThroughputHintInMBs` for a stream configured with
Direct PUT as a source.

_Required_: No

_Type_: [DirectPutSourceConfiguration](aws-properties-kinesisfirehose-deliverystream-directputsourceconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ElasticsearchDestinationConfiguration`

An Amazon ES destination for the delivery stream.

Conditional. You must specify only one destination configuration.

If you change the delivery stream destination from an Amazon ES destination to an
Amazon S3 or Amazon Redshift destination, update requires [some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt).

_Required_: Conditional

_Type_: [ElasticsearchDestinationConfiguration](aws-properties-kinesisfirehose-deliverystream-elasticsearchdestinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExtendedS3DestinationConfiguration`

An Amazon S3 destination for the delivery stream.

Conditional. You must specify only one destination configuration.

If you change the delivery stream destination from an Amazon Extended S3 destination
to an Amazon ES destination, update requires [some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt).

_Required_: Conditional

_Type_: [ExtendedS3DestinationConfiguration](aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpEndpointDestinationConfiguration`

Enables configuring Kinesis Firehose to deliver data to any HTTP endpoint
destination. You can specify only one destination.

_Required_: No

_Type_: [HttpEndpointDestinationConfiguration](aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IcebergDestinationConfiguration`

Specifies the destination configure settings for Apache Iceberg Table.

_Required_: No

_Type_: [IcebergDestinationConfiguration](aws-properties-kinesisfirehose-deliverystream-icebergdestinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisStreamSourceConfiguration`

When a Kinesis stream is used as the source for the delivery stream, a [KinesisStreamSourceConfiguration](../userguide/aws-properties-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration.md) containing the Kinesis stream ARN and the role
ARN for the source stream.

_Required_: No

_Type_: [KinesisStreamSourceConfiguration](aws-properties-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MSKSourceConfiguration`

The configuration for the Amazon MSK cluster to be used as the source for a delivery
stream.

_Required_: No

_Type_: [MSKSourceConfiguration](aws-properties-kinesisfirehose-deliverystream-msksourceconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RedshiftDestinationConfiguration`

An Amazon Redshift destination for the delivery stream.

Conditional. You must specify only one destination configuration.

If you change the delivery stream destination from an Amazon Redshift destination to
an Amazon ES destination, update requires [some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt).

_Required_: Conditional

_Type_: [RedshiftDestinationConfiguration](aws-properties-kinesisfirehose-deliverystream-redshiftdestinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3DestinationConfiguration`

The `S3DestinationConfiguration` property type specifies an Amazon Simple
Storage Service (Amazon S3) destination to which Amazon Kinesis Data Firehose (Kinesis Data
Firehose) delivers data.

Conditional. You must specify only one destination configuration.

If you change the delivery stream destination from an Amazon S3 destination to an
Amazon ES destination, update requires [some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt).

_Required_: Conditional

_Type_: [S3DestinationConfiguration](aws-properties-kinesisfirehose-deliverystream-s3destinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnowflakeDestinationConfiguration`

Configure Snowflake destination

_Required_: No

_Type_: [SnowflakeDestinationConfiguration](aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SplunkDestinationConfiguration`

The configuration of a destination in Splunk for the delivery stream.

_Required_: No

_Type_: [SplunkDestinationConfiguration](aws-properties-kinesisfirehose-deliverystream-splunkdestinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A set of tags to assign to the Firehose stream. A tag is a key-value pair that you can
define and assign to AWS resources. Tags are metadata. For example, you can
add friendly names and descriptions or other types of information that can help you
distinguish the Firehose stream. For more information about tags, see [Using\
Cost Allocation Tags](../../../awsaccountbilling/latest/aboutv2/cost-alloc-tags.md) in the AWS Billing and Cost Management User
Guide.

You can specify up to 50 tags when creating a Firehose stream.

If you specify tags in the `CreateDeliveryStream` action, Amazon Data
Firehose performs an additional authorization on the
`firehose:TagDeliveryStream` action to verify if users have permissions to
create tags. If you do not provide this permission, requests to create new Firehose streams
with IAM resource tags will fail with an `AccessDeniedException` such as
following.

**AccessDeniedException**

User: arn:aws:sts::x:assumed-role/x/x is not authorized to perform: firehose:TagDeliveryStream on resource: arn:aws:firehose:us-east-1:x:deliverystream/x with an explicit deny in an identity-based policy.

For an example IAM policy, see [Tag example.](../../../../reference/firehose/latest/apireference/api-createdeliverystream.md#API_CreateDeliveryStream_Examples)

_Required_: No

_Type_: Array of [Tag](aws-properties-kinesisfirehose-deliverystream-tag.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function, Ref
returns the delivery stream name, such as
`mystack-deliverystream-1ABCD2EF3GHIJ`.

For more information about using the Ref function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

Fn::GetAtt returns a value for a specified attribute of this type. The following are
the available attributes and sample return values.

For more information about using Fn::GetAtt, see [Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the delivery stream, such as
`arn:aws:firehose:us-east-2:123456789012:deliverystream/delivery-stream-name`.

## Examples

- [Create a Kinesis Data Firehose Delivery Stream](#aws-resource-kinesisfirehose-deliverystream--examples--Create_a_Kinesis_Data_Firehose_Delivery_Stream)

- [Convert Record Format](#aws-resource-kinesisfirehose-deliverystream--examples--Convert_Record_Format)

- [Specify an Amazon S3 Destination for the Delivery Stream](#aws-resource-kinesisfirehose-deliverystream--examples--Specify_an_Amazon_S3_Destination_for_the_Delivery_Stream)

- [Specify a Kinesis Stream as the Source for the Delivery Stream](#aws-resource-kinesisfirehose-deliverystream--examples--Specify_a_Kinesis_Stream_as_the_Source_for_the_Delivery_Stream)

### Create a Kinesis Data Firehose Delivery Stream

The following example creates a Kinesis Data Firehose delivery stream that
delivers data to an Amazon ES destination. Kinesis Data Firehose backs up all data
sent to the destination in an Amazon S3 bucket.

#### JSON

```json

"ElasticSearchDeliveryStream": {
   "Type": "AWS::KinesisFirehose::DeliveryStream",
   "Properties": {
      "ElasticsearchDestinationConfiguration": {
         "BufferingHints": {
            "IntervalInSeconds": 60,
            "SizeInMBs": 50
      },
      "CloudWatchLoggingOptions": {
         "Enabled": true,
         "LogGroupName": "deliverystream",
         "LogStreamName": "elasticsearchDelivery"
      },
      "DomainARN": { "Ref" : "MyDomainARN" },
      "IndexName": { "Ref" : "MyIndexName" },
      "IndexRotationPeriod": "NoRotation",
      "TypeName" : "fromFirehose",
      "RetryOptions": {
         "DurationInSeconds": "60"
      },
      "RoleARN": { "Fn::GetAtt" : ["ESdeliveryRole", "Arn"] },
      "S3BackupMode": "AllDocuments",
      "S3Configuration": {
         "BucketARN": { "Ref" : "MyBackupBucketARN" },
         "BufferingHints": {
            "IntervalInSeconds": "60",
            "SizeInMBs": "50"
         },
         "CompressionFormat": "UNCOMPRESSED",
         "Prefix": "firehose/",
         "RoleARN": { "Fn::GetAtt" : ["S3deliveryRole", "Arn"] },
         "CloudWatchLoggingOptions" : {
            "Enabled" : true,
            "LogGroupName" : "deliverystream",
            "LogStreamName" : "s3Backup"
         }
      }
    }
  }
}
```

#### YAML

```yaml

ElasticSearchDeliveryStream:
   Type: AWS::KinesisFirehose::DeliveryStream
   Properties:
      ElasticsearchDestinationConfiguration:
         BufferingHints:
            IntervalInSeconds: 60
            SizeInMBs: 50
         CloudWatchLoggingOptions:
            Enabled: true
            LogGroupName: "deliverystream"
            LogStreamName: "elasticsearchDelivery"
         DomainARN:
            Ref: "MyDomainARN"
         IndexName:
            Ref: "MyIndexName"
         IndexRotationPeriod: "NoRotation"
         TypeName: "fromFirehose"
         RetryOptions:
            DurationInSeconds: "60"
         RoleARN:
            Fn::GetAtt:
               - "ESdeliveryRole"
               - "Arn"
         S3BackupMode: "AllDocuments"
         S3Configuration:
            BucketARN:
               Ref: "MyBackupBucketARN"
            BufferingHints:
               IntervalInSeconds: "60"
               SizeInMBs: "50"
            CompressionFormat: "UNCOMPRESSED"
            Prefix: "firehose/"
            RoleARN:
               Fn::GetAtt:
                  - "S3deliveryRole"
                  - "Arn"
            CloudWatchLoggingOptions:
               Enabled: true
               LogGroupName: "deliverystream"
               LogStreamName: "s3Backup"
```

### Convert Record Format

The following example shows record format conversion.

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Stack for Firehose DeliveryStream S3 Destination.
Resources:

  GlueDatabase:
    Type: AWS::Glue::Database
    Properties:
      CatalogId: !Ref AWS::AccountId
      DatabaseInput: {}

  GlueTable:
    Type: AWS::Glue::Table
    Properties:
      CatalogId: !Ref AWS::AccountId
      DatabaseName: !Ref GlueDatabase
      TableInput:
        Owner: owner
        Retention: 0
        StorageDescriptor:
          Columns:
          - Name: pickup_latitude
            Type: double
          - Name: pickup_longitude
            Type: double
          - Name: dropoff_latitude
            Type: double
          - Name: dropoff_longitude
            Type: double
          - Name: trip_id
            Type: int
          - Name: trip_distance
            Type: double
          - Name: passenger_count
            Type: int
          - Name: pickup_datetime
            Type: timestamp
          - Name: dropoff_datetime
            Type: timestamp
          - Name: total_amount
            Type: double
          InputFormat: org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat
          OutputFormat: org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat
          Compressed: false
          NumberOfBuckets: -1
          SerdeInfo:
            SerializationLibrary: org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe
            Parameters:
              serialization.format: '1'
          BucketColumns: []
          SortColumns: []
          StoredAsSubDirectories: false
        PartitionKeys:
        - Name: year
          Type: string
        - Name: month
          Type: string
        - Name: day
          Type: string
        - Name: hour
          Type: string
        TableType: EXTERNAL_TABLE

  deliverystream:
    Type: AWS::KinesisFirehose::DeliveryStream
    Properties:
      DeliveryStreamType: DirectPut
      ExtendedS3DestinationConfiguration:
        RoleARN: !GetAtt deliveryRole.Arn
        BucketARN: !Join
          - ''
          - - 'arn:aws:s3:::'
            - !Ref s3bucket
        Prefix: !Join
          - ''
          - - !Ref GlueTable
            -  '/year=!{timestamp:YYYY}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/'
        ErrorOutputPrefix: !Join
          - ''
          - - !Ref GlueTable
            -  'error/!{firehose:error-output-type}/year=!{timestamp:YYYY}/month=!{timestamp:MM}/day=!{timestamp:dd}/hour=!{timestamp:HH}/'
        BufferingHints:
          SizeInMBs: 128
          IntervalInSeconds: 300
        CompressionFormat: UNCOMPRESSED
        EncryptionConfiguration:
          NoEncryptionConfig: NoEncryption
        CloudWatchLoggingOptions:
          Enabled: true
          LogGroupName: !Join
            - ''
            - - 'KDF-'
              - !Ref GlueTable
          LogStreamName: S3Delivery
        S3BackupMode: Disabled
        DataFormatConversionConfiguration:
          SchemaConfiguration:
            CatalogId: !Ref AWS::AccountId
            RoleARN: !GetAtt deliveryRole.Arn
            DatabaseName: !Ref GlueDatabase
            TableName: !Ref GlueTable
            Region: !Ref AWS::Region
            VersionId: LATEST
          InputFormatConfiguration:
            Deserializer:
              OpenXJsonSerDe: {}
          OutputFormatConfiguration:
            Serializer:
              ParquetSerDe: {}
          Enabled: True

  s3bucket:
    Type: AWS::S3::Bucket
    Properties:
      VersioningConfiguration:
        Status: Enabled

  deliveryRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2012-10-17
        Statement:
          - Sid: ''
            Effect: Allow
            Principal:
              Service: firehose.amazonaws.com
            Action: 'sts:AssumeRole'
            Condition:
              StringEquals:
                'sts:ExternalId': !Ref 'AWS::AccountId'
      Path: "/"
      Policies:
        - PolicyName: firehose_delivery_policy
          PolicyDocument:
            Version: 2012-10-17
            Statement:
              - Effect: Allow
                Action:
                  - 's3:AbortMultipartUpload'
                  - 's3:GetBucketLocation'
                  - 's3:GetObject'
                  - 's3:ListBucket'
                  - 's3:ListBucketMultipartUploads'
                  - 's3:PutObject'
                Resource:
                  - !Join
                    - ''
                    - - 'arn:aws:s3:::'
                      - !Ref s3bucket
                  - !Join
                    - ''
                    - - 'arn:aws:s3:::'
                      - !Ref s3bucket
                      - '/*'
              - Effect: Allow
                Action: 'glue:GetTableVersions'
                Resource: '*'
              - Effect: Allow
                Action: 'logs:PutLogEvents'
                Resource:
                - !Join
                    - ''
                    - - 'arn:aws:logs:'
                      - !Ref 'AWS::Region'
                      - ':'
                      - !Ref 'AWS::AccountId'
                      - 'log-group:/aws/kinesisfirehose/KDF-'
                      - !Ref GlueTable
                      - ':log-stream:*'
Outputs:
  deliverysreamARN:
    Description: The ARN of the firehose delivery stream
    Value: !GetAtt deliverystream.Arn
```

### Specify an Amazon S3 Destination for the Delivery Stream

The following example uses the `ExtendedS3DestinationConfiguration`
property to specify an Amazon S3 destination for the delivery stream.

#### JSON

```json

{
  "Resources":{
    "Firehose":{
      "Type" : "AWS::KinesisFirehose::DeliveryStream",
      "Properties" : {
          "DeliveryStreamName" : "tester-partitioning-delimiter",
          "DeliveryStreamType" : "DirectPut",
          "ExtendedS3DestinationConfiguration":
            {
              "BucketARN": "arn:aws:s3:::dp-firehose-test",
              "BufferingHints": {
                "SizeInMBs": 128,
                "IntervalInSeconds": 900
              },
              "CompressionFormat": "UNCOMPRESSED",
              "ErrorOutputPrefix": "table/error/!{firehose:error-output-type}/dt=!{timestamp:yyyy'-'MM'-'dd}/h=!{timestamp:HH}/",
              "Prefix": "YYYY=!{partitionKeyFromQuery:YYYY}/MM=!{partitionKeyFromQuery:MM}//DD=!{partitionKeyFromQuery:DD}/HH=!{partitionKeyFromQuery:HH}/REGION=!{partitionKeyFromQuery:REGION}/SITEID=!{partitionKeyFromQuery:SITEID}/",
              "RoleARN": "arn:aws:iam::012345678912:role/service-role/KinesisFirehoseServiceRole-dp-kinesis-f-us-east-1-012345678912",
              "DynamicPartitioningConfiguration":
              {
                "Enabled": true,
                "RetryOptions": {
                  "DurationInSeconds": 300
                  }
                },
                "ProcessingConfiguration": {
                  "Enabled": true,
                  "Processors": [
                  {
                    "Type": "MetadataExtraction",
                    "Parameters": [
                      {
                      "ParameterName": "MetadataExtractionQuery",
                      "ParameterValue": "{YYYY : (.ts/1000) | strftime(\"%Y\"), MM : (.ts/1000) | strftime(\"%m\"), DD : (.ts/1000) | strftime(\"%d\"), HH: (.ts/1000) | strftime(\"%H\")}"
                      },
                      {
                        "ParameterName": "JsonParsingEngine",
                        "ParameterValue": "JQ-1.6"
                      }
                    ]
                  },
				    {
                    "Type": "AppendDelimiterToRecord",
                    "Parameters": [
                        {
                        "ParameterName": "Delimiter",
                        "ParameterValue": "\\n"
                      }
                    ]
                  }
                ]
                }
              }
            }
          }
        }
      }
```

#### YAML

```yaml

---
Resources:
  Firehose:
    Type: AWS::KinesisFirehose::DeliveryStream
    Properties:
      DeliveryStreamName: tester-partitioning-delimiter
      DeliveryStreamType: DirectPut
      ExtendedS3DestinationConfiguration:
        BucketARN: arn:aws:s3:::dp-firehose-test
        BufferingHints:
          SizeInMBs: 128
          IntervalInSeconds: 900
        CompressionFormat: UNCOMPRESSED
        ErrorOutputPrefix: table/error/!{firehose:error-output-type}/dt=!{timestamp:yyyy'-'MM'-'dd}/h=!{timestamp:HH}/
        Prefix: YYYY=!{partitionKeyFromQuery:YYYY}/MM=!{partitionKeyFromQuery:MM}//DD=!{partitionKeyFromQuery:DD}/HH=!{partitionKeyFromQuery:HH}/REGION=!{partitionKeyFromQuery:REGION}/SITEID=!{partitionKeyFromQuery:SITEID}/
        RoleARN: arn:aws:iam::012345678912:role/service-role/KinesisFirehoseServiceRole-dp-kinesis-f-us-east-1-012345678912
        DynamicPartitioningConfiguration:
          Enabled: true
          RetryOptions:
            DurationInSeconds: 300
        ProcessingConfiguration:
          Enabled: true
          Processors:
          - Type: MetadataExtraction
            Parameters:
            - ParameterName: MetadataExtractionQuery
              ParameterValue: '{YYYY : (.ts/1000) | strftime("%Y"), MM : (.ts/1000)
                | strftime("%m"), DD : (.ts/1000) | strftime("%d"), HH: (.ts/1000)
                | strftime("%H")}'
            - ParameterName: JsonParsingEngine
              ParameterValue: JQ-1.6
          - Type: AppendDelimiterToRecord
            Parameters:
            - ParameterName: Delimiter
              ParameterValue: "\\n"

```

### Specify a Kinesis Stream as the Source for the Delivery Stream

The following example uses the `KinesisStreamSourceConfiguration`
property to specify a Kinesis stream as the source for the delivery stream.

#### JSON

```json

{
    "Parameters": {
        "deliveryRoleArn": {
            "Type": "String"
        },
        "deliveryStreamName": {
            "Type": "String"
        },
        "kinesisStreamARN": {
            "Type": "String"
        },
        "kinesisStreamRoleArn": {
            "Type": "String"
        },
        "s3bucketArn": {
            "Type": "String"
        }
    },
    "Resources": {
        "Deliverystream": {
            "Type": "AWS::KinesisFirehose::DeliveryStream",
            "Properties": {
                "DeliveryStreamName": {
                    "Ref": "deliveryStreamName"
                },
                "DeliveryStreamType": "KinesisStreamAsSource",
                "KinesisStreamSourceConfiguration": {
                    "KinesisStreamARN": {
                        "Ref": "kinesisStreamARN"
                    },
                    "RoleARN": {
                        "Ref": "kinesisStreamRoleArn"
                    }
                },
                "ExtendedS3DestinationConfiguration": {
                    "BucketARN": {
                        "Ref": "s3bucketArn"
                    },
                    "BufferingHints": {
                        "IntervalInSeconds": 60,
                        "SizeInMBs": 50
                    },
                    "CompressionFormat": "UNCOMPRESSED",
                    "Prefix": "firehose/",
                    "RoleARN": {
                        "Ref": "deliveryRoleArn"
                    }
                }
            }
        }
    }
}
```

#### YAML

```yaml

Parameters:
    deliveryRoleArn:
        Type: String
    deliveryStreamName:
        Type: String
    kinesisStreamARN :
        Type : String
    kinesisStreamRoleArn:
        Type : String
    s3bucketArn:
        Type: String
Resources :
    Deliverystream:
        Type: AWS::KinesisFirehose::DeliveryStream
        Properties:
            DeliveryStreamName: !Ref deliveryStreamName
            DeliveryStreamType: KinesisStreamAsSource
            KinesisStreamSourceConfiguration:
                KinesisStreamARN: !Ref kinesisStreamARN
                RoleARN: !Ref kinesisStreamRoleArn
            ExtendedS3DestinationConfiguration:
                BucketARN: !Ref s3bucketArn
                BufferingHints:
                    IntervalInSeconds: 60
                    SizeInMBs: 50
                CompressionFormat: UNCOMPRESSED
                Prefix: firehose/
                RoleARN: !Ref deliveryRoleArn
```

## See also

- [CreateDeliveryStream](../../../../reference/firehose/latest/apireference/api-createdeliverystream.md) in the _Amazon Kinesis Data Firehose API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Data Firehose

AmazonOpenSearchServerlessBufferingHints

All content copied from https://docs.aws.amazon.com/.
