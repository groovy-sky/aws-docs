This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::StaticIp

The `AWS::Lightsail::StaticIp` resource specifies a static IP that can be
attached to an Amazon Lightsail instance that is in the same AWS Region and Availability Zone.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lightsail::StaticIp",
  "Properties" : {
      "AttachedTo" : String,
      "StaticIpName" : String
    }
}

```

### YAML

```yaml

Type: AWS::Lightsail::StaticIp
Properties:
  AttachedTo: String
  StaticIpName: String

```

## Properties

`AttachedTo`

The instance that the static IP is attached to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StaticIpName`

The name of the static IP.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`IpAddress`

The IP address of the static IP.

`IsAttached`

A Boolean value indicating whether the static IP is attached to an instance.

`StaticIpArn`

The Amazon Resource Name (ARN) of the static IP (for example,
`arn:aws:lightsail:us-east-2:123456789101:StaticIp/244ad76f-8aad-4741-809f-12345EXAMPLE`).

## Remarks

_An instance must be in a running state to attach a static_
_IP_

To attach a static IP to an instance, the instance must be in a `running`
state. If the instance does not come to `running` state within 15 minutes
after performing an attach static IP request, the attach static IP request
times-out.

_You can attach only one static IP to an instance_

You can attach one static IP to a single instance. You cannot attach multiple static
IPs to one instance. If multiple static IPs have the same instance in the
`AttachedTo` parameter, the behavior is unpredictable and any of the
static IPs (but only one) could be attached to the instance. This will cause the stack
to drift because only one of the static IPs will be attached to the instance but the
template will show multiple.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Lightsail::LoadBalancerTlsCertificate

Next
