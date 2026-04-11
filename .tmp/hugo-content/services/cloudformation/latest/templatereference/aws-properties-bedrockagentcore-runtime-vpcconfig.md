This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Runtime VpcConfig

VpcConfig for the Agent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroups" : [ String, ... ],
  "Subnets" : [ String, ... ]
}

```

### YAML

```yaml

  SecurityGroups:
    - String
  Subnets:
    - String

```

## Properties

`SecurityGroups`

The security groups associated with the VPC configuration.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subnets`

The subnets associated with the VPC configuration.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SessionStorageConfiguration

WorkloadIdentityDetails

All content copied from https://docs.aws.amazon.com/.
