This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeStarConnections::Connection

The AWS::CodeStarConnections::Connection resource can be used to connect external source
providers with services like AWS CodePipeline.

**Note:** A connection created through CloudFormation is
in `PENDING` status by default. You can make its status `AVAILABLE`
by updating the connection in the console.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeStarConnections::Connection",
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

Type: AWS::CodeStarConnections::Connection
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

Specifies the tags applied to the resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-codestarconnections-connection-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the connection. The ARN is
used as the connection reference when the connection is shared between AWS
services. For example:

`arn:aws:codestar-connections:us-west-2:123456789012:connection/39e4c34d-e13a-4e94-a886-ea67651bf042`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ConnectionArn`

The Amazon Resource Name (ARN) of the connection. The ARN is used as the connection
reference when the connection is shared between AWS services. For example:
`arn:aws:codestar-connections:us-west-2:123456789012:connection/39e4c34d-e13a-4e94-a886-ea67651bf042`.

`ConnectionStatus`

The current status of the connection. For example: `PENDING`,
`AVAILABLE`, or `ERROR`.

`OwnerAccountId`

The AWS account ID of the owner of the connection. For Bitbucket, this
is the account ID of the owner of the Bitbucket repository. For example:
`123456789012`.

## Examples

- [Bitbucket Connection Configuration](#aws-resource-codestarconnections-connection--examples--Bitbucket_Connection_Configuration)

- [GitHub Enterprise Server Connection Configuration](#aws-resource-codestarconnections-connection--examples--GitHub_Enterprise_Server_Connection_Configuration)

### Bitbucket Connection Configuration

The following example creates a connection with Bitbucket.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "SampleConnection": {
            "Type": "AWS::CodeStarConnections::Connection",
            "Properties": {
                "ConnectionName": "MyConnection",
                "ProviderType": "Bitbucket",
                "Tags": [
                    {
                        "Key": "Project",
                        "Value": "ProjectB"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  SampleConnection:
    Type: AWS::CodeStarConnections::Connection
    Properties:
      ConnectionName: MyConnection
      ProviderType: Bitbucket
      Tags:
      - Key: Project
        Value: ProjectB
```

### GitHub Enterprise Server Connection Configuration

The following example creates a connection with GitHub Enterprise Server.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "SampleConnection": {
            "Type": "AWS::CodeStarConnections::Connection",
            "Properties": {
                "ConnectionName": "MyConnection",
                "ProviderType": "GitHubEnterpriseServer",
                "HostArn": "arn:aws:codestar-connections:us-west-2:123456789123:host/abc123-example",
                "Tags": [
                    {
                        "Key": "Project",
                        "Value": "ProjectB"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  SampleConnection:
    Type: AWS::CodeStarConnections::Connection
    Properties:
      ConnectionName: MyConnection
      ProviderType: GitHubEnterpriseServer
      HostArn: arn:aws:codestar-connections:us-west-2:123456789123:host/abc123-example
      Tags:
      - Key: Project
        Value: ProjectB
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS CodeStar Connections

Tag

All content copied from https://docs.aws.amazon.com/.
