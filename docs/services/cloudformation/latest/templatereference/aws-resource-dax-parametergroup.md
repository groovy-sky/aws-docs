This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DAX::ParameterGroup

A named set of parameters that are applied to all of the nodes in a DAX cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DAX::ParameterGroup",
  "Properties" : {
      "Description" : String,
      "ParameterGroupName" : String,
      "ParameterNameValues" : Json
    }
}

```

### YAML

```yaml

Type: AWS::DAX::ParameterGroup
Properties:
  Description: String
  ParameterGroupName: String
  ParameterNameValues: Json

```

## Properties

`Description`

A description of the parameter group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterGroupName`

The name of the parameter group.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ParameterNameValues`

An array of name-value pairs for the parameters in the group. Each element in the
array represents a single parameter.

###### Note

`record-ttl-millis` and `query-ttl-millis` are the only
supported parameter names. For more details, see [Configuring TTL Settings](../../../dynamodb/latest/developerguide/dax-cluster-management.md#DAX.cluster-management.custom-settings.ttl).

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the created parameter group. For example:

```nohighlight

{ "Ref": "MyDAXParameterGroup" }
```

Returns a value similar to the
following:

```

my-dax-parameter-group
```

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Create Parameter Group

The following example creates a DAX parameter group.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "DAX parameter group",
  "Resources": {
    "daxParamGroup": {
      "Type": "AWS::DAX::ParameterGroup",
      "Properties": {
        "ParameterGroupName": "MyDAXParameterGroup",
        "Description": "Description for my DAX parameter group",
        "ParameterNameValues": {
          "query-ttl-millis": "75000",
          "record-ttl-millis": "88000"
        }
      }
    }
  },
  "Outputs": {
    "ParameterGroup": {
      "Value": {
        "Ref": "daxParamGroup"
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"
Description: "DAX parameter group"
Resources:
  daxParamGroup:
    Type: AWS::DAX::ParameterGroup
    Properties:
      ParameterGroupName: "MyDAXParameterGroup"
      Description: "Description for my DAX parameter group"
      ParameterNameValues:
         "query-ttl-millis" : "75000"
         "record-ttl-millis" : "88000"
Outputs:
  ParameterGroup:
    Value: !Ref daxParamGroup
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SSESpecification

AWS::DAX::SubnetGroup

All content copied from https://docs.aws.amazon.com/.
