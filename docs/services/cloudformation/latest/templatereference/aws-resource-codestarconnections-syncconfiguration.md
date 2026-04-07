This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeStarConnections::SyncConfiguration

Information, such as repository, branch, provider, and resource names for a specific sync configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeStarConnections::SyncConfiguration",
  "Properties" : {
      "Branch" : String,
      "ConfigFile" : String,
      "PublishDeploymentStatus" : String,
      "RepositoryLinkId" : String,
      "ResourceName" : String,
      "RoleArn" : String,
      "SyncType" : String,
      "TriggerResourceUpdateOn" : String
    }
}

```

### YAML

```yaml

Type: AWS::CodeStarConnections::SyncConfiguration
Properties:
  Branch: String
  ConfigFile: String
  PublishDeploymentStatus: String
  RepositoryLinkId: String
  ResourceName: String
  RoleArn: String
  SyncType: String
  TriggerResourceUpdateOn: String

```

## Properties

`Branch`

The branch associated with a specific sync configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfigFile`

The file path to the configuration file associated with a specific sync configuration. The path should point to an actual file in the sync configurations linked repository.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublishDeploymentStatus`

Whether to enable or disable publishing of deployment status to source providers.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepositoryLinkId`

The ID of the repository link associated with a specific sync configuration.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceName`

The name of the connection resource associated with a specific sync configuration.

_Required_: Yes

_Type_: String

_Pattern_: `[a-za-z0-9_\.-]+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role associated with a specific sync configuration.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws(-[\w]+)*:iam::\d{12}:role/[a-zA-Z_0-9+=,.@\-_/]+`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SyncType`

The type of sync for a specific sync configuration.

_Required_: Yes

_Type_: String

_Allowed values_: `CFN_STACK_SYNC`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TriggerResourceUpdateOn`

When to trigger Git sync to begin the stack update.

_Required_: No

_Type_: String

_Allowed values_: `ANY_CHANGE | FILE_CHANGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the sync configuration.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`OwnerId`

The owner ID for the repository associated with a specific sync configuration, such as
the owner ID in GitHub.

`ProviderType`

Property description not available.

`RepositoryName`

Property description not available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Next
