This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Redshift::ClusterParameterGroup

Describes a parameter group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Redshift::ClusterParameterGroup",
  "Properties" : {
      "Description" : String,
      "ParameterGroupFamily" : String,
      "ParameterGroupName" : String,
      "Parameters" : [ Parameter, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Redshift::ClusterParameterGroup
Properties:
  Description: String
  ParameterGroupFamily: String
  ParameterGroupName: String
  Parameters:
    - Parameter
  Tags:
    - Tag

```

## Properties

`Description`

The description of the parameter group.

_Required_: Yes

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ParameterGroupFamily`

The name of the cluster parameter group family that this cluster parameter group is
compatible with. You can create a custom parameter group and then associate your cluster with it. For more information, see
[Amazon Redshift parameter groups](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html).

_Required_: Yes

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ParameterGroupName`

The name of the cluster parameter group.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Parameters`

An array of parameters to be modified. A maximum of 20 parameters can be modified
in a single request.

For each parameter to be modified, you must supply at least the parameter name and
parameter value; other name-value pairs of the parameter are optional.

For the workload management (WLM) configuration, you must supply all the name-value
pairs in the wlm\_json\_configuration parameter.

_Required_: No

_Type_: Array of [Parameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-redshift-clusterparametergroup-parameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The list of tags for the cluster parameter group.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-redshift-clusterparametergroup-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myClusterParameterGroup" }`

For the Amazon Redshift cluster parameter group
`myClusterParameterGroup`, `Ref` returns the name of the cluster
parameter group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

- [Single Parameter](#aws-resource-redshift-clusterparametergroup--examples--Single_Parameter)

- [Workload Management Configuration](#aws-resource-redshift-clusterparametergroup--examples--Workload_Management_Configuration)

### Single Parameter

The following example describes a parameter group with one parameter that's
specified:

#### JSON

```json

"myClusterParameterGroup" : {
  "Type" : "AWS::Redshift::ClusterParameterGroup",
  "Properties" : {
    "Description" : "My parameter group",
    "ParameterGroupFamily" : "redshift-1.0",
    "Parameters" : [ {
      "ParameterName" : "enable_user_activity_logging",
      "ParameterValue" : "true"
    }]
  }
}
```

#### YAML

```yaml

myClusterParameterGroup:
  Type: "AWS::Redshift::ClusterParameterGroup"
  Properties:
    Description: "My parameter group"
    ParameterGroupFamily: "redshift-1.0"
    Parameters:
      -
        ParameterName: "enable_user_activity_logging"
        ParameterValue: "true"
```

### Workload Management Configuration

The following example modifies the workload management configuration using the
`wlm_json_configuration` parameter. The parameter value is a JSON
object that must be passed as a string enclosed in quotation marks (").

#### JSON

```json

"RedshiftClusterParameterGroup": {
  "Type": "AWS::Redshift::ClusterParameterGroup",
  "Properties": {
    "Description": "Cluster parameter group",
    "ParameterGroupFamily": "redshift-1.0",
    "Parameters": [{
      "ParameterName": "wlm_json_configuration",
      "ParameterValue": "[{\"user_group\":[\"example_user_group1\"],\"query_group\":[\"example_query_group1\"],\"query_concurrency\":7},{\"query_concurrency\":5}]"
    }],
    "Tags": [
      {
        "Key": "foo",
        "Value": "bar"
      }
    ]
  }
}
```

#### YAML

```yaml

RedshiftClusterParameterGroup:
  Type: "AWS::Redshift::ClusterParameterGroup"
  Properties:
    Description: "Cluster parameter group"
    ParameterGroupFamily: "redshift-1.0"
    Parameters:
      -
        ParameterName: "wlm_json_configuration"
        ParameterValue: "[{\"user_group\":[\"example_user_group1\"],\"query_group\":[\"example_query_group1\"],\"query_concurrency\":7},{\"query_concurrency\":5}]"
    Tags:
      - Key: foo
        Value: bar
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Parameter
