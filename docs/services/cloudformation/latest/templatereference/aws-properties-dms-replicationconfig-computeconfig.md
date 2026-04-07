This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::ReplicationConfig ComputeConfig

Configuration parameters for provisioning an AWS DMS Serverless replication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AvailabilityZone" : String,
  "DnsNameServers" : String,
  "KmsKeyId" : String,
  "MaxCapacityUnits" : Integer,
  "MinCapacityUnits" : Integer,
  "MultiAZ" : Boolean,
  "PreferredMaintenanceWindow" : String,
  "ReplicationSubnetGroupId" : String,
  "VpcSecurityGroupIds" : [ String, ... ]
}

```

### YAML

```yaml

  AvailabilityZone: String
  DnsNameServers: String
  KmsKeyId: String
  MaxCapacityUnits: Integer
  MinCapacityUnits: Integer
  MultiAZ: Boolean
  PreferredMaintenanceWindow: String
  ReplicationSubnetGroupId: String
  VpcSecurityGroupIds:
    - String

```

## Properties

`AvailabilityZone`

The Availability Zone where the AWS DMS Serverless replication using this configuration
will run. The default value is a random, system-chosen Availability Zone in the
configuration's AWS Region, for example, `"us-west-2"`. You can't set this
parameter if the `MultiAZ` parameter is set to `true`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DnsNameServers`

A list of custom DNS name servers supported for the AWS DMS Serverless replication to
access your source or target database. This list overrides the default name servers
supported by the AWS DMS Serverless replication. You can specify a comma-separated list of
internet addresses for up to four DNS name servers. For example:
`"1.1.1.1,2.2.2.2,3.3.3.3,4.4.4.4"`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

An AWS Key Management Service (AWS KMS) key Amazon Resource Name (ARN) that is used to encrypt the data
during AWS DMS Serverless replication.

If you don't specify a value for the `KmsKeyId` parameter, AWS DMS uses your
default encryption key.

AWS KMS creates the default encryption key for your Amazon Web Services account. Your AWS account
has a different default encryption key for each AWS Region.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxCapacityUnits`

Specifies the maximum value of the AWS DMS capacity units (DCUs) for which a given AWS DMS
Serverless replication can be provisioned. A single DCU is 2GB of RAM, with 1 DCU as the
minimum value allowed. The list of valid DCU values includes 1, 2, 4, 8, 16, 32, 64, 128,
192, 256, and 384. So, the maximum value that you can specify for AWS DMS Serverless is 384.
The `MaxCapacityUnits` parameter is the only DCU parameter you are required to
specify.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinCapacityUnits`

Specifies the minimum value of the AWS DMS capacity units (DCUs) for which a given AWS DMS
Serverless replication can be provisioned. A single DCU is 2GB of RAM, with 1 DCU as the
minimum value allowed. The list of valid DCU values includes 1, 2, 4, 8, 16, 32, 64, 128,
192, 256, and 384. So, the minimum DCU value that you can specify for AWS DMS Serverless is
1\. If you don't set this value, AWS DMS sets this parameter to the minimum DCU value allowed,
1\. If there is no current source activity, AWS DMS scales down your replication until it
reaches the value specified in `MinCapacityUnits`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultiAZ`

Specifies whether the AWS DMS Serverless replication is a Multi-AZ deployment. You can't
set the `AvailabilityZone` parameter if the `MultiAZ` parameter is
set to `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredMaintenanceWindow`

The weekly time range during which system maintenance can occur for the AWS DMS Serverless
replication, in Universal Coordinated Time (UTC). The format is
`ddd:hh24:mi-ddd:hh24:mi`.

The default is a 30-minute window selected at random from an 8-hour block of time per
AWS Region. This maintenance occurs on a random day of the week. Valid values for days of
the week include `Mon`, `Tue`, `Wed`, `Thu`,
`Fri`, `Sat`, and `Sun`.

Constraints include a minimum 30-minute window.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationSubnetGroupId`

Specifies a subnet group identifier to associate with the AWS DMS Serverless
replication.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcSecurityGroupIds`

Specifies the virtual private cloud (VPC) security group to use with the AWS DMS
Serverless replication. The VPC security group must work with the VPC containing the
replication.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DMS::ReplicationConfig

Tag
