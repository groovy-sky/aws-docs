This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::RepositoryCreationTemplate

The details of the repository creation template associated with the request.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ECR::RepositoryCreationTemplate",
  "Properties" : {
      "AppliedFor" : [ String, ... ],
      "CustomRoleArn" : String,
      "Description" : String,
      "EncryptionConfiguration" : EncryptionConfiguration,
      "ImageTagMutability" : String,
      "ImageTagMutabilityExclusionFilters" : [ ImageTagMutabilityExclusionFilter, ... ],
      "LifecyclePolicy" : String,
      "Prefix" : String,
      "RepositoryPolicy" : String,
      "ResourceTags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ECR::RepositoryCreationTemplate
Properties:
  AppliedFor:
    - String
  CustomRoleArn: String
  Description: String
  EncryptionConfiguration:
    EncryptionConfiguration
  ImageTagMutability: String
  ImageTagMutabilityExclusionFilters:
    - ImageTagMutabilityExclusionFilter
  LifecyclePolicy: String
  Prefix: String
  RepositoryPolicy: String
  ResourceTags:
    - Tag

```

## Properties

`AppliedFor`

A list of enumerable Strings representing the repository creation scenarios that this
template will apply towards. The supported scenarios are PULL\_THROUGH\_CACHE, REPLICATION, and
CREATE\_ON\_PUSH

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomRoleArn`

The ARN of the role to be assumed by Amazon ECR. Amazon ECR will assume your supplied role
when the customRoleArn is specified. When this field isn't specified, Amazon ECR will use the
service-linked role for the repository creation template.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:iam::[0-9]{12}:role/[A-Za-z0-9+=,-.@_]*$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description associated with the repository creation template.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfiguration`

The encryption configuration associated with the repository creation template.

_Required_: No

_Type_: [EncryptionConfiguration](aws-properties-ecr-repositorycreationtemplate-encryptionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageTagMutability`

The tag mutability setting for the repository. If this parameter is omitted, the
default setting of `MUTABLE` will be used which will allow image tags to be
overwritten. If `IMMUTABLE` is specified, all image tags within the
repository will be immutable which will prevent them from being overwritten.

_Required_: No

_Type_: String

_Allowed values_: `MUTABLE | IMMUTABLE | IMMUTABLE_WITH_EXCLUSION | MUTABLE_WITH_EXCLUSION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageTagMutabilityExclusionFilters`

A list of filters that specify which image tags are excluded from the repository
creation template's image tag mutability setting.

_Required_: No

_Type_: Array of [ImageTagMutabilityExclusionFilter](aws-properties-ecr-repositorycreationtemplate-imagetagmutabilityexclusionfilter.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LifecyclePolicy`

The lifecycle policy to use for repositories created using the template.

_Required_: No

_Type_: String

_Minimum_: `100`

_Maximum_: `30720`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The repository namespace prefix associated with the repository creation
template.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*\/?|ROOT)$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RepositoryPolicy`

The repository policy to apply to repositories created using the template. A
repository policy is a permissions policy associated with a repository to control access
permissions.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `10240`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTags`

The metadata to apply to the repository to help you categorize and organize. Each tag
consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have
a maximum length of 256 characters.

_Required_: No

_Type_: Array of [Tag](aws-properties-ecr-repositorycreationtemplate-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreatedAt`

The date and time, in JavaScript date format, when the repository creation template
was created.

`UpdatedAt`

The date and time, in JavaScript date format, when the repository creation template
was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

EncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
