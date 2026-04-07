This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DirectoryService::SimpleAD VpcSettings

Contains VPC information for the [CreateDirectory](https://docs.aws.amazon.com/directoryservice/latest/devguide/API_CreateDirectory.html) or
[CreateMicrosoftAD](https://docs.aws.amazon.com/directoryservice/latest/devguide/API_CreateMicrosoftAD.html)
operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SubnetIds" : [ String, ... ],
  "VpcId" : String
}

```

### YAML

```yaml

  SubnetIds:
    - String
  VpcId: String

```

## Properties

`SubnetIds`

The identifiers of the subnets for the directory servers. The two subnets must be in
different Availability Zones. Directory Service specifies a directory server and a DNS
server in each of these subnets.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcId`

The identifier of the VPC in which to create the directory.

_Required_: Yes

_Type_: String

_Pattern_: `^(vpc-[0-9a-f]{8}|vpc-[0-9a-f]{17})$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Next
