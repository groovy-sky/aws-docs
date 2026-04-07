This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::EIPAssociation

Associates an Elastic IP address with an instance or a network interface. Before you can
use an Elastic IP address, you must allocate it to your account. For more information about
working with Elastic IP addresses, see [Elastic IP address concepts and rules](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#vpc-eip-overview).

You must specify `AllocationId` and either `InstanceId`,
`NetworkInterfaceId`, or `PrivateIpAddress`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::EIPAssociation",
  "Properties" : {
      "AllocationId" : String,
      "InstanceId" : String,
      "NetworkInterfaceId" : String,
      "PrivateIpAddress" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::EIPAssociation
Properties:
  AllocationId: String
  InstanceId: String
  NetworkInterfaceId: String
  PrivateIpAddress: String

```

## Properties

`AllocationId`

The allocation ID. This is required.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceId`

The ID of the instance. The instance must have exactly one attached network interface.
You can specify either the instance ID or the network interface ID, but not both.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkInterfaceId`

The ID of the network interface. If the instance has more than one network interface, you must specify a network interface ID.

You can specify either the instance ID or the network interface ID, but not both.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrivateIpAddress`

The primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the association.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the association.

## Examples

### Associate an Elastic IP address to an instance

The following example creates an Elastic IP address and a network interface,
and associates the Elastic IP address with the network interface. The example
uses the ID of an existing subnet and an example IP address from the subnet
CIDR range.

#### JSON

```json

  "Resources" : {
    "myEIP" : {
        "Type" : "AWS::EC2::EIP",
        "Properties" : {
            "Domain" : "vpc"
        }
    },
    "myENI" : {
        "Type" : "AWS::EC2::NetworkInterface",
        "Properties" : {
            "SubnetId" : "subnet-0231a7e21ca967d2c",
            "PrivateIpAddress": "10.0.0.16"
        }
    },
    "eniAssociation" : {
        "Type" : "AWS::EC2::EIPAssociation",
        "Properties" : {
            "AllocationId" : {
                "Fn::GetAtt" : [ "myEIP", "AllocationId" ]
            },
            "NetworkInterfaceId" : {
                "Ref" : "myENI"
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  myEIP:
    Type: AWS::EC2::EIP
    Properties:
      Domain: vpc
  myENI:
    Type: AWS::EC2::NetworkInterface
    Properties:
      SubnetId: subnet-0231a7e21ca967d2c
      PrivateIpAddress: 10.0.0.16
  eniAssociation:
    Type: AWS::EC2::EIPAssociation
    Properties:
      AllocationId: !GetAtt myEIP.AllocationId
      NetworkInterfaceId: !Ref myENI
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::EC2::EnclaveCertificateIamRoleAssociation
