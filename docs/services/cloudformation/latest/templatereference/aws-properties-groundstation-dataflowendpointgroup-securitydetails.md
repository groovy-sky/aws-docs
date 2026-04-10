This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::DataflowEndpointGroup SecurityDetails

Information about IAM roles, subnets, and security groups needed for this DataflowEndpointGroup.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RoleArn" : String,
  "SecurityGroupIds" : [ String, ... ],
  "SubnetIds" : [ String, ... ]
}

```

### YAML

```yaml

  RoleArn: String
  SecurityGroupIds:
    - String
  SubnetIds:
    - String

```

## Properties

`RoleArn`

The ARN of a role which Ground Station has permission to assume, such as
`arn:aws:iam::012345678910:role/DataDeliveryServiceRole`.

Ground Station will assume this role and create an ENI in your VPC on the specified subnet upon creation of a dataflow endpoint group. This ENI is used as the ingress/egress point for data streamed during a satellite contact.

_Required_: No

_Type_: String

_Pattern_: `^(arn:(aws[a-zA-Z-]*)?:[a-z0-9-.]+:.*)|()$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupIds`

The security group Ids of the security role, such as
`sg-1234567890abcdef0`.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

The subnet Ids of the security details, such as
`subnet-12345678`.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

### Create SecurityDetails

The following example creates Ground Station `SecurityDetails`

#### JSON

```json

{
  "SecurityDetails": {
    "SubnetIds": [
      "subnet-6782e71e"
    ],
    "SecurityGroupIds": [
      "sg-6979fe18"
    ],
    "RoleArn": "arn:aws:iam::012345678910:role/groundstation-service-role-AWSServiceRoleForAmazonGroundStation-EXAMPLEBQ4PI"
  }
}
```

#### YAML

```yaml

SecurityDetails:
  SubnetIds:
    - subnet-12345678
  SecurityGroupIds:
    - sg-87654321
  RoleArn: arn:aws:iam::012345678910:role/groundstation-service-role-AWSServiceRoleForAmazonGroundStation-EXAMPLEABCDE
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RangedSocketAddress

SocketAddress

All content copied from https://docs.aws.amazon.com/.
