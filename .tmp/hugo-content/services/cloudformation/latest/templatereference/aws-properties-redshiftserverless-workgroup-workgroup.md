This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RedshiftServerless::Workgroup Workgroup

The collection of computing resources from which an endpoint is created.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BaseCapacity" : Integer,
  "ConfigParameters" : [ ConfigParameter, ... ],
  "CreationDate" : String,
  "Endpoint" : Endpoint,
  "EnhancedVpcRouting" : Boolean,
  "MaxCapacity" : Integer,
  "NamespaceName" : String,
  "PricePerformanceTarget" : PerformanceTarget,
  "PubliclyAccessible" : Boolean,
  "SecurityGroupIds" : [ String, ... ],
  "Status" : String,
  "SubnetIds" : [ String, ... ],
  "TrackName" : String,
  "WorkgroupArn" : String,
  "WorkgroupId" : String,
  "WorkgroupName" : String
}

```

### YAML

```yaml

  BaseCapacity: Integer
  ConfigParameters:
    - ConfigParameter
  CreationDate: String
  Endpoint:
    Endpoint
  EnhancedVpcRouting: Boolean
  MaxCapacity: Integer
  NamespaceName: String
  PricePerformanceTarget:
    PerformanceTarget
  PubliclyAccessible: Boolean
  SecurityGroupIds:
    - String
  Status: String
  SubnetIds:
    - String
  TrackName: String
  WorkgroupArn: String
  WorkgroupId: String
  WorkgroupName: String

```

## Properties

`BaseCapacity`

The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfigParameters`

An array of parameters to set for advanced control over a database. The
options are `auto_mv`, `datestyle`, `enable_case_sensitive_identifier`, `enable_user_activity_logging`,
`query_group`, `search_path`, `require_ssl`, `use_fips_ssl`, and either `wlm_json_configuration` or query monitoring metrics that
let you define performance boundaries. You can either specify individual query monitoring metrics (such as `max_scan_row_count`, `max_query_execution_time`) or use
`wlm_json_configuration` to define query queues with rules, but not both. If you're using `wlm_json_configuration`, the
maximum size of `parameterValue` is 8000 characters.
For more information about query monitoring rules and available metrics, see
[Query monitoring metrics for Amazon Redshift Serverless](../../../redshift/latest/dg/cm-c-wlm-query-monitoring-rules.md#cm-c-wlm-query-monitoring-metrics-serverless).

_Required_: No

_Type_: Array of [ConfigParameter](aws-properties-redshiftserverless-workgroup-configparameter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreationDate`

The creation date of the workgroup.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Endpoint`

The endpoint that is created from the workgroup.

_Required_: No

_Type_: [Endpoint](aws-properties-redshiftserverless-workgroup-endpoint.md)

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

_Pattern_: `^[a-z0-9-]+$`

_Minimum_: `3`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PricePerformanceTarget`

An object that represents the price performance target settings for the workgroup.

_Required_: No

_Type_: [PerformanceTarget](aws-properties-redshiftserverless-workgroup-performancetarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PubliclyAccessible`

A value that specifies whether the workgroup can be accessible from a public network.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIds`

An array of security group IDs to associate with the workgroup.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the workgroup.

_Required_: No

_Type_: String

_Allowed values_: `CREATING | AVAILABLE | MODIFYING | DELETING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

An array of subnet IDs the workgroup is associated with.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrackName`

The name of the track for the workgroup.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkgroupArn`

The Amazon Resource Name (ARN) that links to the workgroup.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkgroupId`

The unique identifier of the workgroup.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkgroupName`

The name of the workgroup.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9-]*$`

_Minimum_: `3`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcEndpoint

Next

All content copied from https://docs.aws.amazon.com/.
