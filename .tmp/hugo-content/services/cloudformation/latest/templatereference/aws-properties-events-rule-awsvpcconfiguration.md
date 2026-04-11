This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule AwsVpcConfiguration

This structure specifies the VPC subnets and security groups for the task, and whether a
public IP address is to be used. This structure is relevant only for ECS tasks that use the
`awsvpc` network mode.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssignPublicIp" : String,
  "SecurityGroups" : [ String, ... ],
  "Subnets" : [ String, ... ]
}

```

### YAML

```yaml

  AssignPublicIp: String
  SecurityGroups:
    - String
  Subnets:
    - String

```

## Properties

`AssignPublicIp`

Specifies whether the task's elastic network interface receives a public IP address. You
can specify `ENABLED` only when `LaunchType` in
`EcsParameters` is set to `FARGATE`.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroups`

Specifies the security groups associated with the task. These security groups must all be
in the same VPC. You can specify as many as five security groups. If you do not specify a
security group, the default security group for the VPC is used.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subnets`

Specifies the subnets associated with the task. These subnets must all be in the same VPC.
You can specify as many as 16 subnets.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Set the AwsVpcConfiguration parameter

The following example sets the `AwsVpcConfiguration` parameter to not assign a public IP and set the security groups for Vpc01.

#### JSON

```json

"AwsVpcConfiguration": {
  "AssignPublicIp": "DISABLED",
  "SecurityGroups": [
    {
      "Fn: : GetAtt": [
        "ScheduledFargateTaskScheduledTaskDefSecurityGroupE075BC19",
        "GroupId"
      ]
    }
  ],
  "Subnets": [
    {
      "Ref": "Vpc01"
    }
  ]
}
```

#### YAML

```yaml

AwsVpcConfiguration:
  AssignPublicIp: "DISABLED"
  SecurityGroups:
    Fn: : GetAtt:
      "ScheduledFargateTaskScheduledTaskDefSecurityGroupE075BC19",
      "GroupId"
  Subnets:
    Ref:
      "Vpc01"
```

## See also

- [AwsVpcConfiguration](../../../../reference/eventbridge/latest/apireference/api-awsvpcconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AppSyncParameters

BatchArrayProperties

All content copied from https://docs.aws.amazon.com/.
