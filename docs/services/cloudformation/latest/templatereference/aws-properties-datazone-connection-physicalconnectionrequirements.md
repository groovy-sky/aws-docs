This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection PhysicalConnectionRequirements

Physical connection requirements of a connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AvailabilityZone" : String,
  "SecurityGroupIdList" : [ String, ... ],
  "SubnetId" : String,
  "SubnetIdList" : [ String, ... ]
}

```

### YAML

```yaml

  AvailabilityZone: String
  SecurityGroupIdList:
    - String
  SubnetId: String
  SubnetIdList:
    - String

```

## Properties

`AvailabilityZone`

The availability zone of the physical connection requirements of a connection.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIdList`

The group ID list of the physical connection requirements of a connection.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `255 | 50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetId`

The subnet ID of the physical connection requirements of a connection.

_Required_: No

_Type_: String

_Pattern_: `^subnet-[a-z0-9]+$`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIdList`

The subnet ID list of the physical connection requirements of a connection.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `32 | 50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OAuth2Properties

RedshiftCredentials

All content copied from https://docs.aws.amazon.com/.
