This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream RedshiftDestinationConfiguration

The `RedshiftDestinationConfiguration` property type specifies an Amazon
Redshift cluster to which Amazon Kinesis Data Firehose (Kinesis Data Firehose) delivers
data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLoggingOptions" : CloudWatchLoggingOptions,
  "ClusterJDBCURL" : String,
  "CopyCommand" : CopyCommand,
  "Password" : String,
  "ProcessingConfiguration" : ProcessingConfiguration,
  "RetryOptions" : RedshiftRetryOptions,
  "RoleARN" : String,
  "S3BackupConfiguration" : S3DestinationConfiguration,
  "S3BackupMode" : String,
  "S3Configuration" : S3DestinationConfiguration,
  "SecretsManagerConfiguration" : SecretsManagerConfiguration,
  "Username" : String
}

```

### YAML

```yaml

  CloudWatchLoggingOptions:
    CloudWatchLoggingOptions
  ClusterJDBCURL: String
  CopyCommand:
    CopyCommand
  Password: String
  ProcessingConfiguration:
    ProcessingConfiguration
  RetryOptions:
    RedshiftRetryOptions
  RoleARN: String
  S3BackupConfiguration:
    S3DestinationConfiguration
  S3BackupMode: String
  S3Configuration:
    S3DestinationConfiguration
  SecretsManagerConfiguration:
    SecretsManagerConfiguration
  Username: String

```

## Properties

`CloudWatchLoggingOptions`

The CloudWatch logging options for your Firehose stream.

_Required_: No

_Type_: [CloudWatchLoggingOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-cloudwatchloggingoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterJDBCURL`

The connection string that Kinesis Data Firehose uses to connect to the Amazon Redshift
cluster.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CopyCommand`

Configures the Amazon Redshift `COPY` command that Kinesis Data Firehose uses
to load data into the cluster from the Amazon S3 bucket.

_Required_: Yes

_Type_: [CopyCommand](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-copycommand.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Password`

The password for the Amazon Redshift user that you specified in the
`Username` property.

_Required_: No

_Type_: String

_Minimum_: `6`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessingConfiguration`

The data processing configuration for the Kinesis Data Firehose delivery
stream.

_Required_: No

_Type_: [ProcessingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-processingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryOptions`

The retry behavior in case Firehose is unable to deliver documents to
Amazon Redshift. Default value is 3600 (60 minutes).

_Required_: No

_Type_: [RedshiftRetryOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-redshiftretryoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleARN`

The ARN of the AWS Identity and Access Management (IAM) role that
grants Kinesis Data Firehose access to your Amazon S3 bucket and AWS KMS
(if you enable data encryption). For more information, see [Grant Kinesis Data\
Firehose Access to an Amazon Redshift Destination](https://docs.aws.amazon.com/firehose/latest/dev/controlling-access.html#using-iam-rs) in the _Amazon_
_Kinesis Data Firehose Developer Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BackupConfiguration`

The configuration for backup in Amazon S3.

_Required_: No

_Type_: [S3DestinationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-s3destinationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BackupMode`

The Amazon S3 backup mode. After you create a Firehose stream, you can update it to
enable Amazon S3 backup if it is disabled. If backup is enabled, you can't update the
Firehose stream to disable it.

_Required_: No

_Type_: String

_Allowed values_: `Disabled | Enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Configuration`

The S3 bucket where Kinesis Data Firehose first delivers data. After the data is in the
bucket, Kinesis Data Firehose uses the `COPY` command to load the data into the
Amazon Redshift cluster. For the Amazon S3 bucket's compression format, don't specify
`SNAPPY` or `ZIP` because the Amazon Redshift `COPY`
command doesn't support them.

_Required_: Yes

_Type_: [S3DestinationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-s3destinationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerConfiguration`

The configuration that defines how you access secrets for Amazon Redshift.

_Required_: No

_Type_: [SecretsManagerConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-secretsmanagerconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

The Amazon Redshift user that has permission to access the Amazon Redshift cluster.
This user must have `INSERT` privileges for copying data from the Amazon S3
bucket to the cluster.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ProcessorParameter

RedshiftRetryOptions
