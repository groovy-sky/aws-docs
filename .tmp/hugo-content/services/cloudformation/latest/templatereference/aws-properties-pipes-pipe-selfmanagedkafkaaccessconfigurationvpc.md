This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe SelfManagedKafkaAccessConfigurationVpc

This structure specifies the VPC subnets and security groups for the stream, and whether
a public IP address is to be used.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroup" : [ String, ... ],
  "Subnets" : [ String, ... ]
}

```

### YAML

```yaml

  SecurityGroup:
    - String
  Subnets:
    - String

```

## Properties

`SecurityGroup`

Specifies the security groups associated with the stream. These security groups must all
be in the same VPC. You can specify as many as five security groups.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `1024 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subnets`

Specifies the subnets associated with the stream. These subnets must all be in the same
VPC. You can specify as many as 16 subnets.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `1024 | 16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SelfManagedKafkaAccessConfigurationCredentials

SingleMeasureMapping

All content copied from https://docs.aws.amazon.com/.
