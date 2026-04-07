This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SqlHaStandbyDetectedInstance

Specifies an Amazon EC2 instance for SQL Server High Availability
standby detection monitoring. Once enabled, AWS monitors the metadata for
the instance to determine whether it is an active or standby node in a SQL Server High Availability cluster.
If the instance is determined to be a standby failover node, AWS
automatically applies SQL Server licensing fee waiver for this instance.

To register an instance, it must be running a Windows SQL Server license-included
AMI and have the AWS Systems Manager agent installed and running. Only Windows Server 2019 and later
and SQL Server (Standard and Enterprise editions) 2017 and later are supported. For more
information, see [Prerequisites for using SQL Server High Availability instance standby detection](https://docs.aws.amazon.com/sql-server-ec2/latest/userguide/prerequisites-and-requirements.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::SqlHaStandbyDetectedInstance",
  "Properties" : {
      "InstanceId" : String,
      "SqlServerCredentials" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::SqlHaStandbyDetectedInstance
Properties:
  InstanceId: String
  SqlServerCredentials: String

```

## Properties

`InstanceId`

The ID of the SQL Server High Availability instance.

_Required_: Yes

_Type_: String

_Pattern_: `^i-[0-9a-f]{8,17}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SqlServerCredentials`

The ARN of the AWS Secrets Manager secret containing the SQL Server access credentials for the SQL Server High Availability instance.
If not specified, deafult local user credentials will be used by the AWS Systems Manager agent.

_Required_: No

_Type_: String

_Pattern_: `^(?=.{20,2048}$)arn:aws[a-z-]*:secretsmanager:[a-z0-9-]+:\d{12}:secret:[a-zA-Z0-9/_+=.@-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the SQL Server High Availability instance.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`HaStatus`

The SQL Server High Availability status of the instance. Valid values are:

- `processing` \- The SQL Server High Availability status for the SQL Server High Availability instance is being updated.

- `active` \- The SQL Server High Availability instance is an active node in an SQL Server High Availability cluster.

- `standby` \- The SQL Server High Availability instance is a standby failover node in an SQL Server High Availability
cluster.

- `invalid` \- An error occurred due to misconfigured permissions, or unable
to dertemine SQL Server High Availability status for the SQL Server High Availability instance.

`LastUpdatedTime`

The date and time when the instance's SQL Server High Availability status was last updated, in the ISO 8601 format
in the UTC time zone ( `YYYY-MM-DDThh:mm:ss.sssZ`).

`SqlServerLicenseUsage`

The license type for the SQL Server license. Valid values include:

- `full` \- The SQL Server High Availability instance is using a full SQL Server license.

- `waived` \- The SQL Server High Availability instance is waived from the SQL Server license.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VCpuCountRangeRequest

AWS::EC2::Subnet
