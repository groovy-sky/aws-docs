This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::CodeRepository GitConfig

Specifies configuration details for a Git repository in your AWS
account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Branch" : String,
  "RepositoryUrl" : String,
  "SecretArn" : String
}

```

### YAML

```yaml

  Branch: String
  RepositoryUrl: String
  SecretArn: String

```

## Properties

`Branch`

The default branch for the Git repository.

_Required_: No

_Type_: String

_Pattern_: `[^ ~^:?*\[]+`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RepositoryUrl`

The URL where the Git repository is located.

_Required_: Yes

_Type_: String

_Pattern_: `https://([^/]+)/?.{3,1016}`

_Minimum_: `11`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecretArn`

The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that
contains the credentials used to access the git repository. The secret must have a
staging label of `AWSCURRENT` and must be in the following format:

`{"username": UserName, "password":
                    Password}`

_Required_: No

_Type_: String

_Pattern_: `arn:aws[a-z\-]*:secretsmanager:[a-z0-9\-]*:[0-9]{12}:secret:.*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SageMaker::CodeRepository

Tag
