This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::VpcEndpoint

Creates an OpenSearch Serverless-managed interface VPC endpoint. For more information, see [Access\
Amazon OpenSearch Serverless using an interface endpoint](../../../opensearch-service/latest/developerguide/serverless-vpc.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::OpenSearchServerless::VpcEndpoint",
  "Properties" : {
      "Name" : String,
      "SecurityGroupIds" : [ String, ... ],
      "SubnetIds" : [ String, ... ],
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::OpenSearchServerless::VpcEndpoint
Properties:
  Name: String
  SecurityGroupIds:
    - String
  SubnetIds:
    - String
  VpcId: String

```

## Properties

`Name`

The name of the endpoint.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z][a-z0-9-]{2,31}$`

_Minimum_: `3`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupIds`

The unique identifiers of the security groups that define the ports, protocols, and
sources for inbound traffic that you are authorizing into your endpoint.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `128 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

The ID of the subnets from which you access OpenSearch Serverless.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `32 | 6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the VPC from which you access OpenSearch Serverless.

_Required_: Yes

_Type_: String

_Pattern_: `^vpc-[0-9a-z]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the endpoint ID. For more information about using the
`Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

`GetAtt` returns a value for a specified attribute of this type. For more
information, see [Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md). The following are the available attributes and sample return
values.

`Id`

The unique identifier of the endpoint. For example,
`vpce-050f79086ee71ac05`.

## Examples

### Create a VPC endpoint

The following example specifies an OpenSearch Serverless-managed interface VPC
endpoint named `test-vpcendpoint`. The endpoint has one subnet and one
security group.

#### JSON

```json

{
   "Description":"OpenSearch Serverless VPC endpoint template",
   "Resources":{
      "TestAOSSVpcEndpoint":{
         "Type":"AWS::OpenSearchServerless::VpcEndpoint",
         "Properties":{
            "Name":"test-vpcendpoint",
            "VpcId":"vpc-0d728b8430292b3f4",
            "SubnetIds":[
               "subnet-0e855f5722a9598ee"
            ],
            "SecurityGroupIds":[
               "sg-03843b03f369eb245"
            ]
         }
      }
   }
}
```

#### YAML

```yaml

AAWSTemplateFormatVersion: '2010-09-09'
Description: OpenSearch Serverless VPC endpoint template
Resources:
  TestAOSSVpcEndpoint:
    Type: 'AWS::OpenSearchServerless::VpcEndpoint'
    Properties:
      Name: test-vpcendpoint
      VpcId: vpc-0d728b8430292b3f4
      SubnetIds:
        - subnet-0e855f5722a9598ee
      SecurityGroupIds:
        - sg-03843b03f369eb245
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::OpenSearchServerless::SecurityPolicy

Next

All content copied from https://docs.aws.amazon.com/.
