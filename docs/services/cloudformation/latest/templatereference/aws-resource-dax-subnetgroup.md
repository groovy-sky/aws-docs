This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DAX::SubnetGroup

Creates a new subnet group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DAX::SubnetGroup",
  "Properties" : {
      "Description" : String,
      "SubnetGroupName" : String,
      "SubnetIds" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DAX::SubnetGroup
Properties:
  Description: String
  SubnetGroupName: String
  SubnetIds:
    - String

```

## Properties

`Description`

The description of the subnet group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetGroupName`

The name of the subnet group.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

A list of VPC subnet IDs for the subnet group.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the created subnet group. For example

```nohighlight

{ "Ref": "MyDAXSubnetGroup" }
```

Returns
a value similar to the
following:

```

my-dax-subnet-group
```

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Create Subnet group

The following example creates a DAX subnet group.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "Create a DAX subnet group",
  "Resources": {
    "MyDAXSubnetGroup": {
      "Type": "AWS::DAX::SubnetGroup",
      "Properties": {
        "SubnetGroupName": "my-dax-subnet-group",
        "Description": "Description of my DAX subnet group",
        "SubnetIds": [
          "subnet1",
          "subnet2"
        ]
      }
    },
    "subnet1": {
      "Type": "AWS::EC2::Subnet",
      "Properties": {
        "VpcId": "daxVPC",
        "CidrBlock": "172.13.17.0/24",
        "AvailabilityZone": {
          "Fn::Select": [
            0,
            {
              "Fn::GetAZs": ""
            }
          ]
        }
      }
    },
    "subnet2": {
      "Type": "AWS::EC2::Subnet",
      "Properties": {
        "VpcId": "daxVPC",
        "CidrBlock": "172.13.18.0/24",
        "AvailabilityZone": {
          "Fn::Select": [
            1,
            {
              "Fn::GetAZs": ""
            }
          ]
        }
      }
    },
    "daxVpc": {
      "Type": "AWS::EC2::VPC",
      "Properties": {
        "CidrBlock": "172.13.0.0/16"
      }
    }
  },
  "Outputs": {
    "ParameterGroup": {
      "Value": "MyDAXSubnetGroup"
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"
Description: "DAX subnet group"
Resources:
  MyDAXSubnetGroup:
    Type: AWS::DAX::SubnetGroup
    Properties:
      SubnetGroupName: "my-dax-subnet-group"
      Description: "Description of my DAX subnet group"
      SubnetIds:
        - !Ref subnet1
        - !Ref subnet2
  subnet1:
    Type: AWS::EC2::Subnet
    Properties:
      VpcId:
        !Ref daxVpc
      CidrBlock: 172.13.17.0/24
      AvailabilityZone:
        Fn::Select:
          - 0
          - Fn::GetAZs: ""
  subnet2:
    Type: AWS::EC2::Subnet
    Properties:
      VpcId:
        !Ref daxVpc
      CidrBlock: 172.13.18.0/24
      AvailabilityZone:
        Fn::Select:
          - 1
          - Fn::GetAZs: ""
  daxVpc:
     Type: AWS::EC2::VPC
     Properties:
       CidrBlock: 172.13.0.0/16
Outputs:
  ParameterGroup:
Value: !Ref MyDAXSubnetGroup
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DAX::ParameterGroup

Next
