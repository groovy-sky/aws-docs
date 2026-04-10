This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Grafana::Workspace VpcConfiguration

The configuration settings for an Amazon VPC that contains data sources for
your Grafana workspace to connect to.

###### Note

Provided `securityGroupIds` and `subnetIds` must be part of
the same VPC.

Connecting to a private VPC is not yet available in the Asia Pacific (Seoul)
Region (ap-northeast-2).

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

The list of Amazon EC2 security group IDs attached to the Amazon VPC
for your Grafana workspace to connect. Duplicates not allowed.

_Array Members_: Minimum number of 1 items. Maximum number of 5 items.

_Length_: Minimum length of 0. Maximum length of 255.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `255 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

The list of Amazon EC2 subnet IDs created in the Amazon VPC for
your Grafana workspace to connect. Duplicates not allowed.

_Array Members_: Minimum number of 2 items. Maximum number of 6 items.

_Length_: Minimum length of 0. Maximum length of 255.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 2`

_Maximum_: `255 | 6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SamlConfiguration

Next

All content copied from https://docs.aws.amazon.com/.
