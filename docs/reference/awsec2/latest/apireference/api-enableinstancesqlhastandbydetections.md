# EnableInstanceSqlHaStandbyDetections

Enable Amazon EC2 instances running in an SQL Server High Availability cluster for SQL Server High Availability
instance standby detection monitoring. Once enabled, AWS monitors the metadata for
the instances to determine whether they are active or standby nodes in the SQL Server High Availability cluster.
If the instances are determined to be standby failover nodes, AWS
automatically applies SQL Server licensing fee waiver for those instances.

To register an instance, it must be running a Windows SQL Server license-included
AMI and have the AWS Systems Manager agent installed and running. Only Windows Server 2019 and later
and SQL Server (Standard and Enterprise editions) 2017 and later are supported. For more
information, see [Prerequisites for using SQL Server High Availability instance standby detection](../../../../services/sql-server-ec2/latest/userguide/prerequisites-and-requirements.md).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action,
without actually making the request, and provides an error response. If you have the
required permissions, the error response is `DryRunOperation`. Otherwise,
it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceId.N**

The IDs of the instances to enable for SQL Server High Availability standby detection monitoring.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 30 items.

Required: Yes

**SqlServerCredentials**

The ARN of the AWS Secrets Manager secret containing the SQL Server access credentials. The specified
secret must contain valid SQL Server credentials for the specified instances. If not specified,
deafult local user credentials will be used by the AWS Systems Manager agent. To enable
instances with different credentials, you must make separate requests.

Type: String

Pattern: `^(?=.{20,2048}$)arn:aws[a-z-]*:secretsmanager:[a-z0-9-]+:\d{12}:secret:[a-zA-Z0-9/_+=.@-]+`

Required: No

## Response Elements

The following elements are returned by the service.

**instanceSet**

Information about the instances that were enabled for SQL Server High Availability standby
detection monitoring.

Type: Array of [RegisteredInstance](api-registeredinstance.md) objects

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/EnableInstanceSqlHaStandbyDetections)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/EnableInstanceSqlHaStandbyDetections)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/enableinstancesqlhastandbydetections.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/enableinstancesqlhastandbydetections.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/enableinstancesqlhastandbydetections.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/enableinstancesqlhastandbydetections.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/enableinstancesqlhastandbydetections.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/enableinstancesqlhastandbydetections.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/EnableInstanceSqlHaStandbyDetections)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/enableinstancesqlhastandbydetections.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EnableImageDeregistrationProtection

EnableIpamOrganizationAdminAccount
