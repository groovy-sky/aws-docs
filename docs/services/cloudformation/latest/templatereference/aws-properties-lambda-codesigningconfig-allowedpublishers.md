This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::CodeSigningConfig AllowedPublishers

List of signing profiles that can sign a code package.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SigningProfileVersionArns" : [ String, ... ]
}

```

### YAML

```yaml

  SigningProfileVersionArns:
    - String

```

## Properties

`SigningProfileVersionArns`

The Amazon Resource Name (ARN) for each of the signing profiles. A signing profile defines a trusted user
who can sign a code package.

_Required_: Yes

_Type_: Array of String

_Minimum_: `12 | 1`

_Maximum_: `1024 | 20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Lambda::CodeSigningConfig

CodeSigningPolicies
