This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeConnections::Connection

A resource that is used to connect third-party source providers with services like CodePipeline.

Note: A connection created through CloudFormation, the CLI, or the SDK is in \`PENDING\` status by default. You can make its status \`AVAILABLE\` by updating the
connection in the console.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeConnections::Connection",
  "Properties" : {
      "ConnectionName" : String,
      "HostArn" : String,
      "ProviderType" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CodeConnections::Connection
Properties:
  ConnectionName: String
  HostArn: String
  ProviderType: String
  Tags:
    - Tag

```

## Properties

`ConnectionName`

The name of the connection. Connection names must be unique in an AWS account.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HostArn`

The Amazon Resource Name (ARN) of the host associated with the connection.

_Required_: No

_Type_: String

_Pattern_: `arn:aws(-[\w]+)*:.+:.+:[0-9]{12}:.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProviderType`

The name of the external provider where your third-party code repository is
configured.

_Required_: No

_Type_: String

_Allowed values_: `Bitbucket | GitHub | GitHubEnterpriseServer | GitLab | GitLabSelfManaged`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codeconnections-connection-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`ConnectionArn`

The Amazon Resource Name (ARN) of the connection. The ARN is used as the connection
reference when the connection is shared between AWS services.

###### Note

The ARN is never reused if the connection is deleted.

`ConnectionStatus`

The current status of the connection.

`OwnerAccountId`

The identifier of the external provider where your third-party code repository is configured.
For Bitbucket, this is the account ID of the owner of the Bitbucket repository.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS CodeConnections

Tag
