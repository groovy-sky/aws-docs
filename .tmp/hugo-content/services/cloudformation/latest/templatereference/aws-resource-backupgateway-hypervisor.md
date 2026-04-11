This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BackupGateway::Hypervisor

Represents the hypervisor's permissions to which the gateway will connect.

A hypervisor is hardware, software, or firmware that creates and manages virtual machines,
and allocates resources to them.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BackupGateway::Hypervisor",
  "Properties" : {
      "Host" : String,
      "KmsKeyArn" : String,
      "LogGroupArn" : String,
      "Name" : String,
      "Password" : String,
      "Tags" : [ Tag, ... ],
      "Username" : String
    }
}

```

### YAML

```yaml

Type: AWS::BackupGateway::Hypervisor
Properties:
  Host: String
  KmsKeyArn: String
  LogGroupArn: String
  Name: String
  Password: String
  Tags:
    - Tag
  Username: String

```

## Properties

`Host`

The server host of the hypervisor. This can be either an IP address or a fully-qualified
domain name (FQDN).

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Minimum_: `3`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The Amazon Resource Name (ARN) of the AWS Key Management Service used to encrypt the
hypervisor.

_Required_: No

_Type_: String

_Pattern_: `^(^arn:(aws|aws-cn|aws-us-gov):kms:([a-zA-Z0-9-]+):([0-9]+):(key|alias)/(\S+)$)|(^alias/(\S+)$)$`

_Minimum_: `50`

_Maximum_: `500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogGroupArn`

The Amazon Resource Name (ARN) of the group of gateways within
the requested log.

_Required_: No

_Type_: String

_Pattern_: `^$|^arn:(aws|aws-cn|aws-us-gov):logs:([a-zA-Z0-9-]+):([0-9]+):log-group:[a-zA-Z0-9_\-\/\.]+:\*$`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the hypervisor.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Password`

The password for the hypervisor.

_Required_: No

_Type_: String

_Pattern_: `^[ -~]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags of the hypervisor configuration to import.

_Required_: No

_Type_: Array of [Tag](aws-properties-backupgateway-hypervisor-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Username`

The username for the hypervisor.

_Required_: No

_Type_: String

_Pattern_: `^[ -\.0-\[\]-~]*[!-\.0-\[\]-~][ -\.0-\[\]-~]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `HypervisorArn`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`HypervisorArn`

Returns `HypervisorArn`, an Amazon Resource Name (ARN) that uniquely identifies a
Hypervisor. For example:
`arn:aws:backup-gateway:us-east-1:123456789012:hypervisor/hype-1234D67D`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Backup gateway

Tag

All content copied from https://docs.aws.amazon.com/.
