This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate LaunchTemplateTagSpecification

Specifies the tags to apply to the launch template during creation.

To specify the tags for the resources that are created during instance launch, use
[AWS::EC2::LaunchTemplate TagSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-tagspecification.html).

`LaunchTemplateTagSpecification` is a property of [AWS::EC2::LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html).

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

The type of resource. To tag a launch template, `ResourceType` must be
`launch-template`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-launchtemplate-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following example adds the tag `Stack=Production` to the
launch template.

#### YAML

```yaml

TagSpecifications:
  - ResourceType: "launch-template"
    Tags:
    - Key: "Stack"
      Value: "Production"
```

#### JSON

```json

"TagSpecifications": [
    {
        "ResourceType": "launch-template",
        "Tags": [
            {
                "Key": "Stack",
                "Value": "Production"
            }
        ]
    }
]
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LaunchTemplateData

LicenseSpecification
