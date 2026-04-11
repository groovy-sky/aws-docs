This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::SourceCredential

Information about the credentials for a GitHub, GitHub Enterprise, or Bitbucket repository. We strongly recommend that you
use AWS Secrets Manager to store your credentials. If you use
Secrets Manager, you must have secrets in your secrets manager. For more
information, see [Using Dynamic References to Specify Template Values](../userguide/dynamic-references.md#dynamic-references-secretsmanager).

###### Important

For security purposes, do not use plain text in your CloudFormation template to store your credentials.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeBuild::SourceCredential",
  "Properties" : {
      "AuthType" : String,
      "ServerType" : String,
      "Token" : String,
      "Username" : String
    }
}

```

### YAML

```yaml

Type: AWS::CodeBuild::SourceCredential
Properties:
  AuthType: String
  ServerType: String
  Token: String
  Username: String

```

## Properties

`AuthType`

The type of authentication used by the credentials. Valid options are OAUTH,
BASIC\_AUTH, PERSONAL\_ACCESS\_TOKEN, CODECONNECTIONS, or SECRETS\_MANAGER.

_Required_: Yes

_Type_: String

_Allowed values_: `OAUTH | BASIC_AUTH | PERSONAL_ACCESS_TOKEN | CODECONNECTIONS | SECRETS_MANAGER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerType`

The type of source provider. The valid options are GITHUB, GITHUB\_ENTERPRISE, GITLAB, GITLAB\_SELF\_MANAGED, or
BITBUCKET.

_Required_: Yes

_Type_: String

_Allowed values_: `GITHUB | BITBUCKET | GITHUB_ENTERPRISE | GITLAB | GITLAB_SELF_MANAGED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Token`

For GitHub or GitHub Enterprise, this is the personal access token. For Bitbucket,
this is either the access token or the app password. For the `authType` CODECONNECTIONS,
this is the `connectionArn`. For the `authType` SECRETS\_MANAGER, this is the `secretArn`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

The Bitbucket username when the `authType` is BASIC\_AUTH. This parameter
is not valid for other types of source providers or connections.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Create Bitbucket source credentials using AWS Secrets Manager](#aws-resource-codebuild-sourcecredential--examples--Create_Bitbucket_source_credentials_using)

- [Create GitHub Enterprise source credentials using AWS Secrets Manager](#aws-resource-codebuild-sourcecredential--examples--Create_GitHub_Enterprise_source_credentials_using)

- [Create GitHub source credentials using AWS Secrets Manager](#aws-resource-codebuild-sourcecredential--examples--Create_GitHub_source_credentials_using)

- [Import source credentials for Bitbucket](#aws-resource-codebuild-sourcecredential--examples--Import_source_credentials_for_Bitbucket)

- [Import source credentials for Github](#aws-resource-codebuild-sourcecredential--examples--Import_source_credentials_for_Github)

### Create Bitbucket source credentials using AWS Secrets Manager

#### YAML

```yaml

CodeBuildSourceCredential:
  Type: 'AWS::CodeBuild::SourceCredential'
  Properties:
    Token: '{{resolve:secretsmanager:bitbucket:SecretString:token}}'
    ServerType: BITBUCKET
    Username: '{{resolve:secretsmanager:bitbucket:SecretString:username}}'
    AuthType: BASIC_AUTH
```

#### JSON

```json

{
   "CodeBuildSourceCredential": {
      "Type": "AWS::CodeBuild::SourceCredential",
      "Properties": {
         "Token": "{{resolve:secretsmanager:bitbucket:SecretString:token}}",
         "ServerType": "BITBUCKET",
         "Username": "{{resolve:secretsmanager:bitbucket:SecretString:username}}",
         "AuthType": "BASIC_AUTH"
      }
   }
}
```

### Create GitHub Enterprise source credentials using AWS Secrets Manager

#### YAML

```yaml

Resources:
  CodeBuildSourceCredential:
    Type: 'AWS::CodeBuild::SourceCredential'
    Properties:
      Token: '{{resolve:secretsmanager:github_enterprise:SecretString:token}}'
      ServerType: GITHUB_ENTERPRISE
      AuthType: PERSONAL_ACCESS_TOKEN
```

#### JSON

```json

{
   "Resources": {
      "CodeBuildSourceCredential": {
         "Type": "AWS::CodeBuild::SourceCredential",
         "Properties": {
            "Token": "{{resolve:secretsmanager:github_enterprise:SecretString:token}}",
            "ServerType": "GITHUB_ENTERPRISE",
            "AuthType": "PERSONAL_ACCESS_TOKEN"
         }
      }
   }
}
```

### Create GitHub source credentials using AWS Secrets Manager

#### YAML

```yaml

Resources:
  CodeBuildSourceCredential:
    Type: 'AWS::CodeBuild::SourceCredential'
    Properties:
      Token: '{{resolve:secretsmanager:github:SecretString:token}}'
      ServerType: GITHUB
      AuthType: PERSONAL_ACCESS_TOKEN
```

#### JSON

```json

{
   "Resources": {
      "CodeBuildSourceCredential": {
         "Type": "AWS::CodeBuild::SourceCredential",
         "Properties": {
            "Token": "{{resolve:secretsmanager:github:SecretString:token}}",
            "ServerType": "GITHUB",
            "AuthType": "PERSONAL_ACCESS_TOKEN"
         }
      }
   }
}
```

### Import source credentials for Bitbucket

#### YAML

```yaml

Resources:
  MySourceCreds:
    Type: 'AWS::CodeBuild::SourceCredential'
    Properties:
      Token: '{{resolve:secretsmanager:bitbucket:SecretString:token}}'
      ServerType: BITBUCKET
      Username: '{{resolve:secretsmanager:bitbucket:SecretString:username}}'
      AuthType: BASIC_AUTH
  MyProject:
    Type: 'AWS::CodeBuild::Project'
    Properties:
      Name: myProjectName
      Description: A description about my project
      ServiceRole: testServiceRoleArn
      Artifacts:
        Type: NO_ARTIFACTS
      Environment:
        Type: LINUX_CONTAINER
        ComputeType: BUILD_GENERAL1_SMALL
        Image: 'aws/codebuild/standard:5.0'
      Source:
        Type: BITBUCKET
        Location: 'your-bitbucket-repo-url'
    DependsOn: MySourceCreds
```

#### JSON

```json

{
    "Resources": {
        "MySourceCreds": {
            "Type": "AWS::CodeBuild::SourceCredential",
            "Properties": {
                "Token": "{{resolve:secretsmanager:bitbucket:SecretString:token}}",
                "ServerType": "BITBUCKET",
                "Username": "{{resolve:secretsmanager:bitbucket:SecretString:username}}",
                "AuthType": "BASIC_AUTH"
            }
        },
        "MyProject": {
            "Type": "AWS::CodeBuild::Project",
            "Properties": {
                "Name": "myProjectName",
                "Description": "A description about my project",
                "ServiceRole": "testServiceRoleAr",
                "Artifacts": {
                    "Type": "NO_ARTIFACTS"
                },
                "Environment": {
                    "Type": "LINUX_CONTAINER",
                    "ComputeType": "BUILD_GENERAL1_SMALL",
                    "Image": "aws/codebuild/standard:5.0"
                },
                "Source": {
                    "Type": "BITBUCKET",
                    "Location": "your-bitbucket-repo-url"
                }
            },
            "DependsOn": "MySourceCreds"
        }
    }
}
```

### Import source credentials for Github

#### YAML

```yaml

Resources:
  MySourceCreds:
    Type: 'AWS::CodeBuild::SourceCredential'
    Properties:
      Token: '{{resolve:secretsmanager:github:SecretString:token}}'
      ServerType: GITHUB
      AuthType: PERSONAL_ACCESS_TOKEN
  MyProject:
    Type: 'AWS::CodeBuild::Project'
    Properties:
      Name: myProjectName
      Description: A description about my project
      ServiceRole: testServiceRoleArn
      Artifacts:
        Type: NO_ARTIFACTS
      Environment:
        Type: LINUX_CONTAINER
        ComputeType: BUILD_GENERAL1_SMALL
        Image: 'aws/codebuild/standard:5.0'
      Source:
        Type: GITHUB
        Location: 'your-github-repo-url'
    DependsOn: MySourceCreds
```

#### JSON

```json

{
    "Resources": {
        "MySourceCreds": {
            "Type": "AWS::CodeBuild::SourceCredential",
            "Properties": {
                "Token": "{{resolve:secretsmanager:github:SecretString:token}}",
                "ServerType": "GITHUB",
                "AuthType": "PERSONAL_ACCESS_TOKEN"
            }
        },
        "MyProject": {
            "Type": "AWS::CodeBuild::Project",
            "Properties": {
                "Name": "myProjectName",
                "Description": "A description about my project",
                "ServiceRole": "testServiceRoleArn",
                "Artifacts": {
                    "Type": "NO_ARTIFACTS"
                },
                "Environment": {
                    "Type": "LINUX_CONTAINER",
                    "ComputeType": "BUILD_GENERAL1_SMALL",
                    "Image": "aws/codebuild/standard:5.0"
                },
                "Source": {
                    "Type": "GITHUB",
                    "Location": "your-github-repo-url"
                }
            },
            "DependsOn": "MySourceCreds"
        }
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Next

All content copied from https://docs.aws.amazon.com/.
