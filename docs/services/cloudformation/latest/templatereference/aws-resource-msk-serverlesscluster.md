This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::ServerlessCluster

Specifies the properties required for creating a serverless cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MSK::ServerlessCluster",
  "Properties" : {
      "ClientAuthentication" : ClientAuthentication,
      "ClusterName" : String,
      "Tags" : {Key: Value, ...},
      "VpcConfigs" : [ VpcConfig, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MSK::ServerlessCluster
Properties:
  ClientAuthentication:
    ClientAuthentication
  ClusterName: String
  Tags:
    Key: Value
  VpcConfigs:
    - VpcConfig

```

## Properties

`ClientAuthentication`

Includes all client authentication related information.

_Required_: Yes

_Type_: [ClientAuthentication](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-serverlesscluster-clientauthentication.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClusterName`

The name of the cluster.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An arbitrary set of tags (key-value pairs) for the cluster.

_Required_: No

_Type_: Object of String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcConfigs`

VPC configuration information for the serverless cluster.

_Required_: Yes

_Type_: Array of [VpcConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-serverlesscluster-vpcconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you provide the logical ID of this resource to the `Ref` intrinsic function, it returns the ARN of the created MSK cluster.

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

`Arn`

The Amazon Resource Name (ARN) of the MSK cluster.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TopicReplication

ClientAuthentication
