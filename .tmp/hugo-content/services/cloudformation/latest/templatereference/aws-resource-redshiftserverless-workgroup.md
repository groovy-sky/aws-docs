This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RedshiftServerless::Workgroup

The collection of compute resources in Amazon Redshift Serverless.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RedshiftServerless::Workgroup",
  "Properties" : {
      "BaseCapacity" : Integer,
      "ConfigParameters" : [ ConfigParameter, ... ],
      "EnhancedVpcRouting" : Boolean,
      "MaxCapacity" : Integer,
      "NamespaceName" : String,
      "Port" : Integer,
      "PricePerformanceTarget" : PerformanceTarget,
      "PubliclyAccessible" : Boolean,
      "RecoveryPointId" : String,
      "SecurityGroupIds" : [ String, ... ],
      "SnapshotArn" : String,
      "SnapshotName" : String,
      "SnapshotOwnerAccount" : String,
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "TrackName" : String,
      "Workgroup" : Workgroup,
      "WorkgroupName" : String
    }
}

```

### YAML

```yaml

Type: AWS::RedshiftServerless::Workgroup
Properties:
  BaseCapacity: Integer
  ConfigParameters:
    - ConfigParameter
  EnhancedVpcRouting: Boolean
  MaxCapacity: Integer
  NamespaceName: String
  Port: Integer
  PricePerformanceTarget:
    PerformanceTarget
  PubliclyAccessible: Boolean
  RecoveryPointId: String
  SecurityGroupIds:
    - String
  SnapshotArn: String
  SnapshotName: String
  SnapshotOwnerAccount: String
  SubnetIds:
    - String
  Tags:
    - Tag
  TrackName: String
  Workgroup:
    Workgroup
  WorkgroupName: String

```

## Properties

`BaseCapacity`

The base compute capacity of the workgroup in Redshift Processing Units (RPUs).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfigParameters`

The key of the parameter. The options are `auto_mv`, `datestyle`, `enable_case_sensitive_identifier`,
`enable_user_activity_logging`, `query_group`, `search_path`, `require_ssl`, `use_fips_ssl`,
and query monitoring metrics that let you define performance boundaries. For more information about query monitoring rules and available metrics, see
[Query monitoring metrics for Amazon Redshift Serverless](../../../redshift/latest/dg/cm-c-wlm-query-monitoring-rules.md#cm-c-wlm-query-monitoring-metrics-serverless).

_Required_: No

_Type_: Array of [ConfigParameter](aws-properties-redshiftserverless-workgroup-configparameter.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnhancedVpcRouting`

The value that specifies whether to enable enhanced virtual
private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxCapacity`

The maximum data-warehouse capacity Amazon Redshift Serverless uses to serve queries. The max capacity is specified in RPUs.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NamespaceName`

The namespace the workgroup is associated with.

_Required_: No

_Type_: String

_Pattern_: `^(?=^[a-z0-9-]+$).{3,64}$`

_Minimum_: `3`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Port`

The custom port to use when connecting to a workgroup. Valid port ranges are 5431-5455 and 8191-8215.
The default is 5439.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PricePerformanceTarget`

An object that represents the price performance target settings for the workgroup.

_Required_: No

_Type_: [PerformanceTarget](aws-properties-redshiftserverless-workgroup-performancetarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PubliclyAccessible`

A value that specifies whether the workgroup
can be accessible from a public network.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecoveryPointId`

The recovery point id to restore from.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIds`

A list of security group IDs to associate with the workgroup.

_Required_: No

_Type_: Array of String

_Minimum_: `0 | 1`

_Maximum_: `255 | 32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotArn`

The Amazon Resource Name (ARN) of the snapshot to restore from.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotName`

The snapshot name to restore from.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotOwnerAccount`

The Amazon Web Services account that owns the snapshot.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

A list of subnet IDs the workgroup is associated with.

_Required_: No

_Type_: Array of String

_Minimum_: `0 | 1`

_Maximum_: `255 | 32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The map of the key-value pairs used to tag the workgroup.

_Required_: No

_Type_: Array of [Tag](aws-properties-redshiftserverless-workgroup-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrackName`

An optional parameter for the name of the track for the workgroup. If you don't provide
a track name, the workgroup is assigned to the current track.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Workgroup`

The collection of computing resources from which an endpoint is created.

_Required_: No

_Type_: [Workgroup](aws-properties-redshiftserverless-workgroup-workgroup.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkgroupName`

The name of the workgroup.

_Required_: Yes

_Type_: String

_Pattern_: `^(?=^[a-z0-9-]+$).{3,64}$`

_Minimum_: `3`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function, Ref
returns the WorkgroupName, such as `sample-workgroup.` For more
information about using the Ref function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

GetAtt returns a value for a specified attribute of this type. For more information, see
[Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md). The following are the available attributes and sample return
values.

`Workgroup.BaseCapacity`

The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).

`Workgroup.ConfigParameters`

Property description not available.

`Workgroup.CreationDate`

The creation date of the workgroup.

`Workgroup.Endpoint.Address`

The DNS address of the VPC endpoint.

`Workgroup.Endpoint.Port`

The custom port to use when connecting to a workgroup. Valid port ranges are 5431-5455 and 8191-8215.
The default is 5439.

`Workgroup.EnhancedVpcRouting`

The value that specifies whether to enable enhanced virtual
private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.

`Workgroup.MaxCapacity`

The maximum data-warehouse capacity Amazon Redshift Serverless uses to serve queries. The max capacity is specified in RPUs.

`Workgroup.NamespaceName`

The namespace the workgroup is associated with.

`Workgroup.PubliclyAccessible`

A value that specifies whether the workgroup can be accessible from a public network.

`Workgroup.SecurityGroupIds`

An array of security group IDs to associate with the workgroup.

`Workgroup.Status`

The status of the workgroup.

`Workgroup.SubnetIds`

An array of subnet IDs the workgroup is associated with.

`Workgroup.TrackName`

An optional parameter for the name of the track for the workgroup. If you don't provide
a track name, the workgroup is assigned to the current track.

`Workgroup.WorkgroupArn`

The Amazon Resource Name (ARN) that links to the workgroup.

`Workgroup.WorkgroupId`

The unique identifier of the workgroup.

`Workgroup.WorkgroupName`

The name of the workgroup.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ConfigParameter

All content copied from https://docs.aws.amazon.com/.
