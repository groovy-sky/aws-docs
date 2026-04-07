This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# Alexa::ASK::Skill Overrides

The `Overrides` property type provides overrides to the skill package to
apply when creating or updating the skill. Values provided here do not modify the
contents of the original skill package. Currently, only overriding values inside of the
skill manifest component of the package is supported.

`Overrides` is a property of the `Alexa::ASK::Skill SkillPackage`
property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Manifest" : Json
}

```

### YAML

```yaml

  Manifest: Json

```

## Properties

`Manifest`

Overrides to apply to the skill manifest inside of the skill package. The skill
manifest contains metadata about the skill. For more information, see [Skill Manifest Schemas](https://developer.amazon.com/docs/smapi/skill-manifest.html).

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Alexa Skill Overrides Resource Configuration

#### JSON

```json

"Manifest": {
  "publishingInformation": {
    "locales": {
      "en-US": {
        "summary": "Sample Short Description",
      }
    },
    "category": "EDUCATION_AND_REFERENCE"
  }
}
```

#### YAML

```yaml

Manifest:
  publishingInformation:
    locales:
      en-US:
        summary: "Sample Short Description"
    category: "EDUCATION_AND_REFERENCE"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AuthenticationConfiguration

SkillPackage
