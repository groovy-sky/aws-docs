This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Amplify::Domain SubDomainSetting

The SubDomainSetting property type enables you to connect a subdomain (for example,
example.exampledomain.com) to a specific branch.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BranchName" : String,
  "Prefix" : String
}

```

### YAML

```yaml

  BranchName: String
  Prefix: String

```

## Properties

`BranchName`

The branch name setting for the subdomain.

_Length Constraints:_ Minimum length of 1. Maximum length of
255.

_Pattern:_ (?s).+

_Required_: Yes

_Type_: String

_Pattern_: `(?s).+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The prefix setting for the subdomain.

_Required_: Yes

_Type_: String

_Pattern_: `(?s).*`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CertificateSettings

Next
