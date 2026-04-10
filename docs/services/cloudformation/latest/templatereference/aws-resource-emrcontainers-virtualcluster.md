This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::VirtualCluster

The `AWS::EMRContainers::VirtualCluster` resource specifies a virtual cluster. A virtual cluster is a managed entity on Amazon EMR on EKS. You can create, describe, list, and delete virtual clusters. They do not consume any additional resources in your system. A single virtual cluster maps to a single Kubernetes namespace. Given this relationship, you can model virtual clusters the same way you model Kubernetes namespaces to meet your requirements.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EMRContainers::VirtualCluster",
  "Properties" : {
      "ContainerProvider" : ContainerProvider,
      "Name" : String,
      "SecurityConfigurationId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EMRContainers::VirtualCluster
Properties:
  ContainerProvider:
    ContainerProvider
  Name: String
  SecurityConfigurationId: String
  Tags:
    - Tag

```

## Properties

`ContainerProvider`

The container provider of the virtual cluster.

_Required_: Yes

_Type_: [ContainerProvider](aws-properties-emrcontainers-virtualcluster-containerprovider.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the virtual cluster.

_Required_: Yes

_Type_: String

_Pattern_: `[\.\-_/#A-Za-z0-9]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityConfigurationId`

The ID of the security configuration.

_Required_: No

_Type_: String

_Pattern_: `[0-9a-z]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-emrcontainers-virtualcluster-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the virtual cluster.

For more information about using the `Ref` function, see [`Ref`](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](../userguide/intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the project, such as `arn:aws:emr-containers:us-east-1:123456789012:/virtualclusters/ab4rp1abcs8xz47n3x0example`.

`Id`

The ID of the virtual cluster, such as `ab4rp1abcs8xz47n3x0example`.

## Examples

### Specifies a virtual cluster

#### JSON

```json

{
   "Resources": {
      "TestVirtualCluster": {
         "Type": "AWS::EMRContainers::VirtualCluster",
         "Properties": {
            "Name": "VirtualClusterName",
            "ContainerProvider": {
               "Type": "EKS",
               "Id": "EKSClusterName",
               "Info": {
                  "EksInfo": {
                     "Namespace": "EKSNamespace"
                  }
               }
            },
            "Tags": [
               {
                  "Key": "Key1",
                  "Value": "Value1"
               }
            ]
         }
      }
   },
   "Outputs": {
      "PrimaryId": {
         "Value": null
      }
   }
}
```

#### YAML

```yaml

Resources:
  TestVirtualCluster:
    Type: 'AWS::EMRContainers::VirtualCluster'
    Properties:
      Name: 'VirtualClusterName'
      ContainerProvider:
        Type: 'EKS'
        Id: 'EKSClusterName'
        Info:
          EksInfo:
            Namespace: 'EKSNamespace'
      Tags:
        - Key: Key1
          Value: Value1
Outputs:
  PrimaryId:
    Value: !Ref TestVirtualCluster
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TLSCertificateConfiguration

ContainerInfo

All content copied from https://docs.aws.amazon.com/.
