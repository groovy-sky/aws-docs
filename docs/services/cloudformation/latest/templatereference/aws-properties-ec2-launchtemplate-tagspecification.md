This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate TagSpecification

Specifies the tags to apply to resources that are created during instance launch.

`TagSpecification` is a property type of [`TagSpecifications`](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata.md#cfn-ec2-launchtemplate-launchtemplatedata-tagspecifications).
[`TagSpecifications`](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata.md#cfn-ec2-launchtemplate-launchtemplatedata-tagspecifications) is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata.md).

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

The type of resource to tag. You can specify tags for the following resource types
only: `instance` \| `volume` \| `network-interface` \| `spot-instances-request`.
If the instance does not include the resource type that you specify, the instance
launch fails. For example, not all instance types include a volume.

To tag a resource after it has been created, see [CreateTags](../../../../reference/awsec2/latest/apireference/api-createtags.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to apply to the resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-launchtemplate-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following example adds the tag `Stack=Production` to the instances
created by the launch template.

#### YAML

```yaml

TagSpecifications:
  - ResourceType: "instance"
    Tags:
    - Key: "Stack"
      Value: "Production"
```

#### JSON

```json

"TagSpecifications": [
    {
        "ResourceType": "instance",
        "Tags": [
            {
                "Key": "Stack",
                "Value": "Production"
            }
        ]
    }
]
```

## See also

- [LaunchTemplateTagSpecificationRequest](../../../../reference/awsec2/latest/apireference/api-launchtemplatetagspecificationrequest.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TotalLocalStorageGB

All content copied from https://docs.aws.amazon.com/.
