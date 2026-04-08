# AwsVpcConfiguration

An object representing the networking details for a task or service. For example
`awsVpcConfiguration={subnets=["subnet-12344321"],securityGroups=["sg-12344321"]}`.

## Contents

**subnets**

The IDs of the subnets associated with the task or service. There's a limit of 16
subnets that can be specified.

###### Note

All specified subnets must be from the same VPC.

Type: Array of strings

Required: Yes

**assignPublicIp**

Whether the task's elastic network interface receives a public IP address.

Consider the following when you set this value:

- When you use `create-service` or `update-service`, the
default is `DISABLED`.

- When the service `deploymentController` is `ECS`, the
value must be `DISABLED`.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**securityGroups**

The IDs of the security groups associated with the task or service. If you don't
specify a security group, the default security group for the VPC is used. There's a
limit of 5 security groups that can be specified.

###### Note

All specified security groups must be from the same VPC.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/awsvpcconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/awsvpcconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/awsvpcconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AutoScalingGroupProviderUpdate

BaselineEbsBandwidthMbpsRequest
