This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::CapacityReservationFleet TagSpecification

The tags to apply to a resource when the resource is being created. When you specify a tag, you must
specify the resource type to tag, otherwise the request will fail.

###### Note

The `Valid Values` lists all the resource types that can be tagged.
However, the action you're using might not support tagging all of these resource types.
If you try to tag a resource type that is unsupported for the action you're using,
you'll get an error.

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

The type of resource to tag on creation. Specify
`capacity-reservation-fleet`.

To tag a resource after it has been created, see [CreateTags](../../../../reference/awsec2/latest/apireference/api-createtags.md).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to apply to the resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-capacityreservationfleet-tag.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::EC2::CarrierGateway
