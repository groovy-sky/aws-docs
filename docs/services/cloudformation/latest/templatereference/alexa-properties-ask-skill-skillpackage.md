This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# Alexa::ASK::Skill SkillPackage

The `SkillPackage` property type contains configuration details for the
skill package that contains the components of the Alexa skill. Skill packages are
retrieved from an Amazon S3 bucket and key and used to create and update the skill. More
details about the skill package format are located in the [Skill Package API Reference](https://developer.amazon.com/docs/smapi/skill-package-api-reference.html).

`SkillPackage` is a property of the `Alexa::ASK::Skill`
resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Overrides" : Overrides,
  "S3Bucket" : String,
  "S3BucketRole" : String,
  "S3Key" : String,
  "S3ObjectVersion" : String
}

```

### YAML

```yaml

  Overrides:
    Overrides
  S3Bucket: String
  S3BucketRole: String
  S3Key: String
  S3ObjectVersion: String

```

## Properties

`Overrides`

Overrides to the skill package to apply when creating or updating the skill. Values
provided here do not modify the contents of the original skill package. Currently, only
overriding values inside of the skill manifest component of the package is
supported.

_Required_: No

_Type_: [Overrides](alexa-properties-ask-skill-overrides.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Bucket`

The name of the Amazon S3 bucket where the .zip file that contains the skill
package is stored.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketRole`

ARN of the IAM role that grants the Alexa service ( `alexa-appkit.amazon.com`) permission to access the bucket and
retrieve the skill package. This property is optional. If you do not provide it, the bucket
must be publicly accessible or configured with a policy that allows this access.
Otherwise, CloudFormation cannot create the skill.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Key`

The location and name of the skill package .zip file.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3ObjectVersion`

If you have S3 versioning enabled, the version ID of the skill package.zip
file.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Overrides

Next

All content copied from https://docs.aws.amazon.com/.
