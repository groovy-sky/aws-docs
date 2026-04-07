This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Comprehend::DocumentClassifier VpcConfig

Configuration parameters for an optional private Virtual Private Cloud (VPC) containing
the resources you are using for the job. For more information, see [Amazon\
VPC](https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroupIds" : [ String, ... ],
  "Subnets" : [ String, ... ]
}

```

### YAML

```yaml

  SecurityGroupIds:
    - String
  Subnets:
    - String

```

## Properties

`SecurityGroupIds`

The ID number for a security group on an instance of your private VPC. Security groups on
your VPC function serve as a virtual firewall to control inbound and outbound traffic and
provides security for the resources that you’ll be accessing on the VPC. This ID number is
preceded by "sg-", for instance: "sg-03b388029b0a285ea". For more information, see [Security\
Groups for your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html).

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `32 | 5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Subnets`

The ID for each subnet being used in your private VPC. This subnet is a subset of the a
range of IPv4 addresses used by the VPC and is specific to a given availability zone in the
VPC’s Region. This ID number is preceded by "subnet-", for instance:
"subnet-04ccf456919e69055". For more information, see [VPCs and\
Subnets](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html).

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `32 | 16`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::Comprehend::Flywheel
