This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::InfluxDBInstance

A DB instance is an isolated database environment running in the cloud. It is the basic building block of Amazon Timestream for InfluxDB. A DB instance can contain multiple user-created databases (or organizations and buckets for the case of InfluxDb 2.x databases), and can be accessed using the same client tools and applications you might use to access a standalone self-managed InfluxDB instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Timestream::InfluxDBInstance",
  "Properties" : {
      "AllocatedStorage" : Integer,
      "Bucket" : String,
      "DbInstanceType" : String,
      "DbParameterGroupIdentifier" : String,
      "DbStorageType" : String,
      "DeploymentType" : String,
      "LogDeliveryConfiguration" : LogDeliveryConfiguration,
      "Name" : String,
      "NetworkType" : String,
      "Organization" : String,
      "Password" : String,
      "Port" : Integer,
      "PubliclyAccessible" : Boolean,
      "Tags" : [ Tag, ... ],
      "Username" : String,
      "VpcSecurityGroupIds" : [ String, ... ],
      "VpcSubnetIds" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Timestream::InfluxDBInstance
Properties:
  AllocatedStorage: Integer
  Bucket: String
  DbInstanceType: String
  DbParameterGroupIdentifier: String
  DbStorageType: String
  DeploymentType: String
  LogDeliveryConfiguration:
    LogDeliveryConfiguration
  Name: String
  NetworkType: String
  Organization: String
  Password: String
  Port: Integer
  PubliclyAccessible: Boolean
  Tags:
    - Tag
  Username: String
  VpcSecurityGroupIds:
    - String
  VpcSubnetIds:
    - String

```

## Properties

`AllocatedStorage`

The amount of storage to allocate for your DB storage type in GiB (gibibytes).

_Required_: No

_Type_: Integer

_Minimum_: `20`

_Maximum_: `16384`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Bucket`

The name of the initial InfluxDB bucket. All InfluxDB data is stored in a bucket. A bucket combines the concept of a database and a retention period (the duration of time that each data point persists). A bucket belongs to an organization.

_Required_: No

_Type_: String

_Pattern_: `^[^_][^"]*$`

_Minimum_: `2`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DbInstanceType`

The Timestream for InfluxDB DB instance type to run on.

_Required_: No

_Type_: String

_Allowed values_: `db.influx.medium | db.influx.large | db.influx.xlarge | db.influx.2xlarge | db.influx.4xlarge | db.influx.8xlarge | db.influx.12xlarge | db.influx.16xlarge | db.influx.24xlarge`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DbParameterGroupIdentifier`

The name or id of the DB parameter group to assign to your DB instance. DB parameter groups specify how the database is configured. For example, DB parameter groups can specify the limit for query concurrency.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `3`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DbStorageType`

The Timestream for InfluxDB DB storage type to read and write InfluxDB data.

You can choose between 3 different types of provisioned Influx IOPS included storage according to your workloads requirements:

- Influx IO Included 3000 IOPS

- Influx IO Included 12000 IOPS

- Influx IO Included 16000 IOPS

_Required_: No

_Type_: String

_Allowed values_: `InfluxIOIncludedT1 | InfluxIOIncludedT2 | InfluxIOIncludedT3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeploymentType`

Specifies whether the Timestream for InfluxDB is deployed as Single-AZ or with a MultiAZ Standby for High availability.

_Required_: No

_Type_: String

_Allowed values_: `SINGLE_AZ | WITH_MULTIAZ_STANDBY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogDeliveryConfiguration`

Configuration for sending InfluxDB engine logs to a specified S3 bucket.

_Required_: No

_Type_: [LogDeliveryConfiguration](aws-properties-timestream-influxdbinstance-logdeliveryconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name that uniquely identifies the DB instance when interacting with the Amazon Timestream for InfluxDB API and CLI commands. This name will also be a prefix included in the endpoint. DB instance names must be unique per customer and per region.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9]*(-[a-zA-Z0-9]+)*$`

_Minimum_: `3`

_Maximum_: `40`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkType`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | DUAL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Organization`

The name of the initial organization for the initial admin user in InfluxDB. An InfluxDB organization is a workspace for a group of users.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Password`

The password of the initial admin user created in InfluxDB. This password will allow you to access the InfluxDB UI to perform various administrative tasks and also use the InfluxDB CLI to create an operator token. These attributes will be stored in a Secret created in Amazon SecretManager in your account.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `8`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Port`

Property description not available.

_Required_: No

_Type_: Integer

_Minimum_: `1024`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PubliclyAccessible`

Configures the DB instance with a public IP to facilitate access.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of key-value pairs to associate with the DB instance.

_Required_: No

_Type_: Array of [Tag](aws-properties-timestream-influxdbinstance-tag.md)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

The username of the initial admin user created in InfluxDB. Must start with a letter and can't end with a hyphen or contain two consecutive hyphens. For example, my-user1. This username will allow you to access the InfluxDB UI to perform various administrative tasks and also use the InfluxDB CLI to create an operator token. These attributes will be stored in a Secret created in Amazon Secrets Manager in your account.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcSecurityGroupIds`

A list of VPC security group IDs to associate with the DB instance.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcSubnetIds`

A list of VPC subnet IDs to associate with the DB instance. Provide at least two VPC subnet IDs in different availability zones when deploying with a Multi-AZ standby.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `3`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the DB instance

`AvailabilityZone`

The Availability Zone in which the DB instance resides.

`Endpoint`

The endpoint used to connect to InfluxDB. The default InfluxDB port is 8086.

`Id`

A service-generated unique identifier

`InfluxAuthParametersSecretArn`

The Amazon Resource Name (ARN) of the Amazon Secrets Manager secret containing the initial InfluxDB authorization parameters. The secret value is a JSON formatted key-value pair holding InfluxDB authorization values: organization, bucket, username, and password.

`SecondaryAvailabilityZone`

Describes an Availability Zone in which the InfluxDB instance is located

`Status`

The status of the DB instance.

Valid Values: `CREATING` \| `AVAILABLE` \| `DELETING` \| `MODIFYING` \| `UPDATING` \| `DELETED` \| `FAILED`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

LogDeliveryConfiguration

All content copied from https://docs.aws.amazon.com/.
