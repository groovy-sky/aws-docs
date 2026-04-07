This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPCDHCPOptionsAssociation

Associates a set of DHCP options with a VPC, or associates no DHCP options with the
VPC.

After you associate the options with the VPC, any existing instances and all new
instances that you launch in that VPC use the options. You don't need to restart or
relaunch the instances. They automatically pick up the changes within a few hours,
depending on how frequently the instance renews its DHCP lease. You can explicitly renew
the lease using the operating system on the instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VPCDHCPOptionsAssociation",
  "Properties" : {
      "DhcpOptionsId" : String,
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VPCDHCPOptionsAssociation
Properties:
  DhcpOptionsId: String
  VpcId: String

```

## Properties

`DhcpOptionsId`

The ID of the DHCP options set, or `default` to associate
no DHCP options with the VPC.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcId`

The ID of the VPC.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the DHCP options association.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### VPC DHCP options association

The following example uses the `Ref` intrinsic function to associate
the myDHCPOptions DHCP options with the myVPC VPC. The VPC and DHCP options can be
declared in the same template or added as input parameters. For more information
about the VPC or the DHCP options resources, see AWS::EC2::VPC or
AWS::EC2::DHCPOptions.

#### JSON

```json

"myVPCDHCPOptionsAssociation" : {
     "Type" : "AWS::EC2::VPCDHCPOptionsAssociation",
      "Properties" : {
         "VpcId" : {"Ref" : "myVPC"},
         "DhcpOptionsId" : {"Ref" : "myDHCPOptions"}
         }
}
```

#### YAML

```yaml

myVPCDHCPOptionsAssociation:
  Type: AWS::EC2::VPCDHCPOptionsAssociation
  Properties:
    VpcId:
       Ref: myVPC
    DhcpOptionsId:
       Ref: myDHCPOptions
```

## See also

- [AssociateDhcpOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AssociateDhcpOptions.html) in the _Amazon EC2 API_
_Reference_

- [DHCP options sets](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_DHCP_Options.html) in the _Amazon VPC User_
_Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::VPCCidrBlock

AWS::EC2::VPCEncryptionControl
