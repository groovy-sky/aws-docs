This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL CaptchaConfig

Specifies how AWS WAF should handle `CAPTCHA` evaluations for rules that don't have their own `CaptchaConfig` settings. If you don't specify this, AWS WAF uses its default settings for `CaptchaConfig`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ImmunityTimeProperty" : ImmunityTimeProperty
}

```

### YAML

```yaml

  ImmunityTimeProperty:
    ImmunityTimeProperty

```

## Properties

`ImmunityTimeProperty`

Determines how long a `CAPTCHA` timestamp in the token remains valid after the client
successfully solves a `CAPTCHA` puzzle.

_Required_: No

_Type_: [ImmunityTimeProperty](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-immunitytimeproperty.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CaptchaAction

ChallengeAction
