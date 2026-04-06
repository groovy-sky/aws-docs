This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataSource DataSourceVpcConfiguration

Provides configuration information needed to connect to an Amazon VPC (Virtual
Private Cloud).

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

A list of identifiers of security groups within your Amazon VPC. The security
groups should enable Amazon Q Business to connect to the data source.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `200 | 10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

A list of identifiers for subnets within your Amazon VPC. The subnets should
be able to connect to each other in the VPC, and they should have outgoing access to the
Internet through a NAT device.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AudioExtractionConfiguration

DocumentAttributeCondition
