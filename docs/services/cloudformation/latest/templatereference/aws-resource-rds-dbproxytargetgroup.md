This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBProxyTargetGroup

The `AWS::RDS::DBProxyTargetGroup` resource represents a set of RDS DB
instances, Aurora DB clusters, or both that a proxy can connect to. Currently, each
target group is associated with exactly one RDS DB instance or Aurora DB cluster.

This data type is used as a response element in the
`DescribeDBProxyTargetGroups` action.

For information about RDS Proxy for Amazon RDS, see [Managing Connections with Amazon\
RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy.html) in the _Amazon RDS User Guide_.

For information about RDS Proxy for Amazon Aurora, see [Managing Connections with\
Amazon RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-proxy.html) in the _Amazon Aurora User Guide_.

For a sample template that creates a DB proxy and registers a DB instance, see
[Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#aws-resource-rds-dbproxy--examples) in AWS::RDS::DBProxy.

###### Note

Limitations apply to RDS Proxy, including DB engine version limitations and AWS
Region limitations.

For information about limitations that apply to RDS Proxy for Amazon RDS, see
[Limitations for RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy.html#rds-proxy.limitations) in the _Amazon RDS User Guide_.

For information about that apply to RDS Proxy for Amazon Aurora, see [Limitations for RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-proxy.html#rds-proxy.limitations) in the _Amazon Aurora User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RDS::DBProxyTargetGroup",
  "Properties" : {
      "ConnectionPoolConfigurationInfo" : ConnectionPoolConfigurationInfoFormat,
      "DBClusterIdentifiers" : [ String, ... ],
      "DBInstanceIdentifiers" : [ String, ... ],
      "DBProxyName" : String,
      "TargetGroupName" : String
    }
}

```

### YAML

```yaml

Type: AWS::RDS::DBProxyTargetGroup
Properties:
  ConnectionPoolConfigurationInfo:
    ConnectionPoolConfigurationInfoFormat
  DBClusterIdentifiers:
    - String
  DBInstanceIdentifiers:
    - String
  DBProxyName: String
  TargetGroupName: String

```

## Properties

`ConnectionPoolConfigurationInfo`

Displays the settings that control the size and behavior of the connection pool associated with a `DBProxyTarget`.

_Required_: No

_Type_: [ConnectionPoolConfigurationInfoFormat](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rds-dbproxytargetgroup-connectionpoolconfigurationinfoformat.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBClusterIdentifiers`

One or more DB cluster identifiers.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBInstanceIdentifiers`

One or more DB instance identifiers.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBProxyName`

The identifier of the `DBProxy` that is associated with the `DBProxyTargetGroup`.

_Required_: Yes

_Type_: String

_Pattern_: `[A-z][0-z]*`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetGroupName`

The identifier for the target group.

###### Note

Currently, this property must be set to `default`.

_Required_: Yes

_Type_: String

_Allowed values_: `default`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the target group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`TargetGroupArn`

The Amazon Resource Name (ARN) representing the target group.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TagFormat

ConnectionPoolConfigurationInfoFormat
