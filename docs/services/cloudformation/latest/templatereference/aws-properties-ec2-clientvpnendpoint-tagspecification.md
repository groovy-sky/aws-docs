This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::ClientVpnEndpoint TagSpecification

Specifies the tags to apply to the Client VPN endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourceType" : String,
  "Tags" : [ Tag, ... ]
}

```

### YAML

```yaml

  ResourceType: String
  Tags:
    - Tag

```

## Properties

`ResourceType`

The type of resource to tag. To tag a Client VPN endpoint, `ResourceType`
must be `client-vpn-endpoint`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to apply to the resource.

_Required_: Yes

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-clientvpnendpoint-tag.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::EC2::ClientVpnRoute
