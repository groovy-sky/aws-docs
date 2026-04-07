This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::SigningConfiguration Rule

A signing rule that specifies a signing profile and optional
repository filters. When an image is pushed to a matching repository, a
signing job is created using the specified profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RepositoryFilters" : [ RepositoryFilter, ... ],
  "SigningProfileArn" : String
}

```

### YAML

```yaml

  RepositoryFilters:
    - RepositoryFilter
  SigningProfileArn: String

```

## Properties

`RepositoryFilters`

A list of repository filters that determine which repositories
have their images signed on push. If no filters are specified, all
images pushed to the registry are signed using the rule's signing
profile. Maximum of 100 filters per rule.

_Required_: No

_Type_: Array of [RepositoryFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecr-signingconfiguration-repositoryfilter.html)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SigningProfileArn`

The ARN of the AWS Signer signing profile to use for signing images that match this
rule. For more information about signing profiles, see [Signing profiles](https://docs.aws.amazon.com/signer/latest/developerguide/signing-profiles.html) in
the _AWS Signer Developer Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[a-z]+)*:signer:[a-z0-9-]+:[0-9]{12}:\/signing-profiles\/[a-zA-Z0-9_]{2,}$`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RepositoryFilter

Next
