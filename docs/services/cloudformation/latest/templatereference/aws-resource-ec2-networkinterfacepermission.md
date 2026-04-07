This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkInterfacePermission

Specifies a permission for the network interface, For example, you can grant an
AWS-authorized account permission to attach the network interface
to an instance in their account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::NetworkInterfacePermission",
  "Properties" : {
      "AwsAccountId" : String,
      "NetworkInterfaceId" : String,
      "Permission" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::NetworkInterfacePermission
Properties:
  AwsAccountId: String
  NetworkInterfaceId: String
  Permission: String

```

## Properties

`AwsAccountId`

The AWS account ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkInterfaceId`

The ID of the network interface.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Permission`

The type of permission to grant: `INSTANCE-ATTACH` or
`EIP-ASSOCIATE`.

_Required_: Yes

_Type_: String

_Allowed values_: `INSTANCE-ATTACH | EIP-ASSOCIATE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:
`eni-perm-055663b682ea24b48`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Grant INSTANCE-ATTACH permission

The following example creates a permission ( `INSTANCE-ATTACH`) for a
specified network interface and AWS account.

#### JSON

```json

"MyNetworkInterfacePermission": {
   "Type": "AWS::EC2::NetworkInterfacePermission",
   "Properties": {
      "NetworkInterfaceId": "eni-030e3xxx",
      "AwsAccountId": "11111111111",
      "Permission": "INSTANCE-ATTACH"
   }
}

```

#### YAML

```yaml

   MyNetworkInterfacePermission:
      Type: AWS::EC2::NetworkInterfacePermission
      Properties:
         NetworkInterfaceId: eni-030e3xxx
         AwsAccountId: '11111111111'
         Permission: INSTANCE-ATTACH
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EnaSrdUdpSpecification

AWS::EC2::NetworkPerformanceMetricSubscription
