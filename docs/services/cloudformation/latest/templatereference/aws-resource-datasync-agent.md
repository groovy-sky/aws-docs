This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::Agent

The `AWS::DataSync::Agent` resource activates an AWS DataSync
agent that you've deployed for storage discovery or data transfers. The activation
process associates the agent with your AWS account.

For more information, see the following topics in the _AWS DataSync User Guide_:

- [DataSync\
agent requirements](../../../datasync/latest/userguide/agent-requirements.md)

- [DataSync\
network requirements](../../../datasync/latest/userguide/datasync-network.md)

- [Create a DataSync agent](../../../datasync/latest/userguide/configure-agent.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataSync::Agent",
  "Properties" : {
      "ActivationKey" : String,
      "AgentName" : String,
      "SecurityGroupArns" : [ String, ... ],
      "SubnetArns" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "VpcEndpointId" : String
    }
}

```

### YAML

```yaml

Type: AWS::DataSync::Agent
Properties:
  ActivationKey: String
  AgentName: String
  SecurityGroupArns:
    - String
  SubnetArns:
    - String
  Tags:
    - Tag
  VpcEndpointId: String

```

## Properties

`ActivationKey`

Specifies your DataSync agent's activation key. If you don't have an
activation key, see [Activating your agent](../../../datasync/latest/userguide/activate-agent.md).

_Required_: No

_Type_: String

_Pattern_: `[A-Z0-9]{5}(-[A-Z0-9]{5}){4}`

_Maximum_: `29`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AgentName`

Specifies a name for your agent. We recommend specifying a name that you can
remember.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\s+=._:@/-]+$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupArns`

The Amazon Resource Names (ARNs) of the security groups used to protect your data
transfer task subnets. See [SecurityGroupArns](../../../datasync/latest/userguide/api-ec2config.md#DataSync-Type-Ec2Config-SecurityGroupArns).

_Pattern_:
`^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:security-group/.*$`

_Required_: No

_Type_: Array of String

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetArns`

Specifies the ARN of the subnet where your VPC service endpoint is located. You can only
specify one ARN.

_Required_: No

_Type_: Array of String

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Specifies labels that help you categorize, filter, and search for your AWS resources. We recommend creating at least one tag for your agent.

_Required_: No

_Type_: Array of [Tag](aws-properties-datasync-agent-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcEndpointId`

The ID of the virtual private cloud (VPC) endpoint that the agent has access to. This
is the client-side VPC endpoint, powered by AWS PrivateLink. If you don't
have an AWS PrivateLink VPC endpoint, see [AWS PrivateLink and VPC endpoints](../../../vpc/latest/userguide/endpoint-services-overview.md) in the _Amazon VPC User_
_Guide_.

For more information about activating your agent in a private network based on a VPC,
see [Using AWS DataSync in a Virtual Private Cloud](../../../datasync/latest/userguide/datasync-in-vpc.md) in the
_AWS DataSync User Guide._

A VPC endpoint ID looks like this: `vpce-01234d5aff67890e1`.

_Required_: No

_Type_: String

_Pattern_: `^vpce-[0-9a-f]{17}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the agent Amazon Resource Name (ARN). For
example:

`arn:aws:datasync:us-east-2:111222333444:agent/agent-0b0addbeef44baca3`

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified
attribute of this type. The following are the available attributes and sample return
values.

For more information about using the `Fn::GetAtt` intrinsic function, see
[Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md).

`AgentArn`

The Amazon Resource Name (ARN) of the agent. Use the `ListAgents` operation
to return a list of agents for your account and AWS Region.

`EndpointType`

The type of endpoint that your agent is connected to. If the endpoint is a VPC
endpoint, the agent is not accessible over the public internet.

## Examples

### DataSync Agent

The following example specifies a DataSync agent named `MyAgent`.
The agent activation key is included in the template.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "Specifies a DataSync agent",
  "Resources":
    {
      "Agent": {
        "Type": "AWS::DataSync::Agent",
        "Properties": {
          "ActivationKey": "AAAAA-7AAAA-GG7MC-3I9R3-27COD",
          "AgentName": "MyAgent"
          }
      }
    }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies a DataSync agent
Resources:
  Agent:
    Type: AWS::DataSync::Agent
    Properties:
      ActivationKey: AAAAA-7AAAA-GG7MC-3I9R3-27COD
      AgentName: MyAgent

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS DataSync

Tag

All content copied from https://docs.aws.amazon.com/.
