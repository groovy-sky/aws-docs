This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::ReplicationInstance

The `AWS::DMS::ReplicationInstance` resource creates an AWS DMS replication instance.
To create a ReplicationInstance, you need permissions to create instances. You'll need similar permissions to terminate
instances when you delete stacks with instances.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DMS::ReplicationInstance",
  "Properties" : {
      "AllocatedStorage" : Integer,
      "AllowMajorVersionUpgrade" : Boolean,
      "AutoMinorVersionUpgrade" : Boolean,
      "AvailabilityZone" : String,
      "DnsNameServers" : String,
      "EngineVersion" : String,
      "KmsKeyId" : String,
      "MultiAZ" : Boolean,
      "NetworkType" : String,
      "PreferredMaintenanceWindow" : String,
      "PubliclyAccessible" : Boolean,
      "ReplicationInstanceClass" : String,
      "ReplicationInstanceIdentifier" : String,
      "ReplicationSubnetGroupIdentifier" : String,
      "ResourceIdentifier" : String,
      "Tags" : [ Tag, ... ],
      "VpcSecurityGroupIds" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DMS::ReplicationInstance
Properties:
  AllocatedStorage: Integer
  AllowMajorVersionUpgrade: Boolean
  AutoMinorVersionUpgrade: Boolean
  AvailabilityZone: String
  DnsNameServers: String
  EngineVersion: String
  KmsKeyId: String
  MultiAZ: Boolean
  NetworkType: String
  PreferredMaintenanceWindow: String
  PubliclyAccessible: Boolean
  ReplicationInstanceClass: String
  ReplicationInstanceIdentifier: String
  ReplicationSubnetGroupIdentifier: String
  ResourceIdentifier: String
  Tags:
    - Tag
  VpcSecurityGroupIds:
    - String

```

## Properties

`AllocatedStorage`

The amount of storage (in gigabytes) to be initially allocated for the replication instance.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowMajorVersionUpgrade`

Indicates that major version upgrades are allowed. Changing this parameter does not
result in an outage, and the change is asynchronously applied as soon as possible.

This parameter must be set to `true` when specifying a value for the
`EngineVersion` parameter that is a different major version than the
replication instance's current version.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoMinorVersionUpgrade`

A value that indicates whether minor engine upgrades are applied automatically to the
replication instance during the maintenance window. This parameter defaults to `true`.

Default: `true`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZone`

The Availability Zone that the replication instance will be created in.

The default value is a random, system-chosen Availability Zone in the endpoint's AWS Region,
for example `us-east-1d`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DnsNameServers`

A list of custom DNS name servers supported for the replication instance to access your
on-premise source or target database. This list overrides the default name servers
supported by the replication instance. You can specify a comma-separated list of internet
addresses for up to four on-premise DNS name servers. For example:
`"1.1.1.1,2.2.2.2,3.3.3.3,4.4.4.4"`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EngineVersion`

The engine version number of the replication instance.

If an engine version number is not specified when a replication
instance is created, the default is the latest engine version available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

An AWS KMS key identifier that is used to encrypt the data on the replication instance.

If you don't specify a value for the `KmsKeyId` parameter, AWS DMS uses your default encryption key.

AWS KMS creates the default encryption key for your AWS account. Your AWS account
has a different default encryption key for each AWS Region.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MultiAZ`

Specifies whether the replication instance is a Multi-AZ deployment. You can't set the
`AvailabilityZone` parameter if the Multi-AZ parameter is set to `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkType`

The type of IP address protocol used by a replication instance, such as IPv4 only or
Dual-stack that supports both IPv4 and IPv6 addressing. IPv6 only is not yet
supported.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredMaintenanceWindow`

The weekly time range during which system maintenance can occur, in UTC.

_Format_: `ddd:hh24:mi-ddd:hh24:mi`

_Default_: A 30-minute window selected at random from an 8-hour block of time per AWS Region,
occurring on a random day of the week.

_Valid days_ ( `ddd`): `Mon` \| `Tue` \|
`Wed` \| `Thu` \| `Fri` \| `Sat` \| `Sun`

_Constraints_: Minimum 30-minute window.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PubliclyAccessible`

Specifies the accessibility options for the replication instance. A value of
`true` represents an instance with a public IP address. A value of
`false` represents an instance with a private IP address.
The default value is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReplicationInstanceClass`

The compute and memory capacity of the replication instance as defined for the specified
replication instance class. For example, to specify the instance class dms.c4.large, set this parameter to `"dms.c4.large"`.
For more information on the settings and capacities for the available replication instance classes, see
[Selecting the right AWS DMS replication instance for your migration](../../../dms/latest/userguide/chap-replicationinstance.md#CHAP_ReplicationInstance.InDepth)
in the _AWS Database Migration Service User Guide_.

_Required_: Yes

_Type_: String

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationInstanceIdentifier`

The replication instance identifier. This parameter is stored as a lowercase string.

Constraints:

- Must contain 1-63 alphanumeric characters or hyphens.

- First character must be a letter.

- Can't end with a hyphen or contain two consecutive hyphens.

Example: `myrepinstance`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationSubnetGroupIdentifier`

A subnet group to associate with the replication instance.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceIdentifier`

A display name for the resource identifier at the end of the `EndpointArn`
response parameter that is returned in the created `Endpoint` object. The value
for this parameter can have up to 31 characters. It can contain only ASCII letters, digits,
and hyphen ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens,
and can only begin with a letter, such as `Example-App-ARN1`. For example, this
value might result in the `EndpointArn` value
`arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1`. If you don't
specify a `ResourceIdentifier` value, AWS DMS generates a default identifier
value for the end of `EndpointArn`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

One or more tags to be assigned to the replication instance.

_Required_: No

_Type_: Array of [Tag](aws-properties-dms-replicationinstance-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcSecurityGroupIds`

Specifies the virtual private cloud (VPC) security group to be used with the replication instance. The VPC
security group must work with the VPC containing the replication instance.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the replication instance ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ReplicationInstancePrivateIpAddresses`

One or more private IP addresses for the replication instance.

`ReplicationInstancePublicIpAddresses`

One or more public IP addresses for the replication instance.

## Examples

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "BasicReplicationInstance": {
            "Type": "AWS::DMS::ReplicationInstance",
            "Properties": {
                "ReplicationInstanceClass": "dms.t2.small"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  BasicReplicationInstance:
    Type: AWS::DMS::ReplicationInstance
    Properties:
      ReplicationInstanceClass: dms.t2.small
```

## See also

- [CreateReplicationInstance](../../../../reference/dms/latest/apireference/api-createreplicationinstance.md)
in the _AWS Database Migration Service API Reference_

- [Managing AWS resources as a single unit with AWS CloudFormation stacks](../userguide/stacks.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
