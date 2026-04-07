This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::Listener

Creates a listener for a service. Before you start using your Amazon VPC Lattice service, you must
add one or more listeners. A listener is a process that checks for connection requests to your
services. For more information, see [Listeners](https://docs.aws.amazon.com/vpc-lattice/latest/ug/listeners.html) in the
_Amazon VPC Lattice User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VpcLattice::Listener",
  "Properties" : {
      "DefaultAction" : DefaultAction,
      "Name" : String,
      "Port" : Integer,
      "Protocol" : String,
      "ServiceIdentifier" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::VpcLattice::Listener
Properties:
  DefaultAction:
    DefaultAction
  Name: String
  Port: Integer
  Protocol: String
  ServiceIdentifier: String
  Tags:
    - Tag

```

## Properties

`DefaultAction`

The action for the default rule. Each listener has a default rule. The default rule is used
if no other rules match.

_Required_: Yes

_Type_: [DefaultAction](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-vpclattice-listener-defaultaction.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the listener. A listener name must be unique within a service. The valid
characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last
character, or immediately after another hyphen.

If you don't specify a name, CloudFormation generates one. However, if
you specify a name, and later want to replace the resource, you must specify a new
name.

_Required_: No

_Type_: String

_Pattern_: `^(?!listener-)(?![-])(?!.*[-]$)(?!.*[-]{2})[a-z0-9-]+$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Port`

The listener port. You can specify a value from 1 to 65535. For HTTP, the default is 80. For
HTTPS, the default is 443.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Protocol`

The listener protocol.

_Required_: Yes

_Type_: String

_Allowed values_: `HTTP | HTTPS | TLS_PASSTHROUGH`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServiceIdentifier`

The ID or ARN of the service.

_Required_: No

_Type_: String

_Pattern_: `^((svc-[0-9a-z]{17})|(arn:[a-z0-9\-]+:vpc-lattice:[a-zA-Z0-9\-]+:\d{12}:service/svc-[0-9a-z]{17}))$`

_Minimum_: `21`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the listener.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-vpclattice-listener-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the listener.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the listener.

`Id`

The ID of the listener.

`ServiceArn`

The Amazon Resource Name (ARN) of the service.

`ServiceId`

The ID of the service.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TxtMethodConfig

DefaultAction
