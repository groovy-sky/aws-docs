This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::Configuration VpcConfig

VPC configuration for workflow runs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroupIds" : [ String, ... ],
  "SubnetIds" : [ String, ... ]
}

```

### YAML

```yaml

  SecurityGroupIds:
    - String
  SubnetIds:
    - String

```

## Properties

`SecurityGroupIds`

List of security group IDs. Maximum of 5 security groups allowed.

_Required_: No

_Type_: Array of String

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

List of subnet IDs. Maximum of 16 subnets allowed.

_Required_: No

_Type_: Array of String

_Maximum_: `16`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RunConfigurations

AWS::Omics::ReferenceStore

All content copied from https://docs.aws.amazon.com/.
