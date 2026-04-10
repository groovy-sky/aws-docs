This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::CloudFormationProduct CodeStarParameters

The subtype containing details about the Codestar connection `Type`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArtifactPath" : String,
  "Branch" : String,
  "ConnectionArn" : String,
  "Repository" : String
}

```

### YAML

```yaml

  ArtifactPath: String
  Branch: String
  ConnectionArn: String
  Repository: String

```

## Properties

`ArtifactPath`

The absolute path wehre the artifact resides within the repo and branch, formatted as
"folder/file.json."

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Branch`

The specific branch where the artifact resides.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `250`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionArn`

The CodeStar ARN, which is the connection between AWS Service Catalog and the external repository.

_Required_: Yes

_Type_: String

_Pattern_: `arn:[a-z0-9][-.a-z0-9]{0,62}:(codestar-connections|codeconnections):([a-z0-9][-.a-z0-9]{0,62})?:([a-z0-9][-.a-z0-9]{0,62})?:[^/].{0,1023}$`

_Minimum_: `1`

_Maximum_: `1224`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Repository`

The specific repository where the product’s artifact-to-be-synced resides, formatted as
"Account/Repo."

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ServiceCatalog::CloudFormationProduct

ConnectionParameters

All content copied from https://docs.aws.amazon.com/.
