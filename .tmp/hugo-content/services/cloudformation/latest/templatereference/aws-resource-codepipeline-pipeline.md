This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline

The `AWS::CodePipeline::Pipeline` resource creates a CodePipeline pipeline
that describes how software changes go through a release process. For more information, see
[What Is \
CodePipeline?](../../../codepipeline/latest/userguide/welcome.md) in the _CodePipeline User Guide_.

For an example in YAML and JSON that contains the parameters in this reference, see
[Examples](../userguide/aws-resource-codepipeline-pipeline.md#aws-resource-codepipeline-pipeline--examples).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodePipeline::Pipeline",
  "Properties" : {
      "ArtifactStore" : ArtifactStore,
      "ArtifactStores" : [ ArtifactStoreMap, ... ],
      "DisableInboundStageTransitions" : [ StageTransition, ... ],
      "ExecutionMode" : String,
      "Name" : String,
      "PipelineType" : String,
      "RestartExecutionOnUpdate" : Boolean,
      "RoleArn" : String,
      "Stages" : [ StageDeclaration, ... ],
      "Tags" : [ Tag, ... ],
      "Triggers" : [ PipelineTriggerDeclaration, ... ],
      "Variables" : [ VariableDeclaration, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CodePipeline::Pipeline
Properties:
  ArtifactStore:
    ArtifactStore
  ArtifactStores:
    - ArtifactStoreMap
  DisableInboundStageTransitions:
    - StageTransition
  ExecutionMode: String
  Name: String
  PipelineType: String
  RestartExecutionOnUpdate: Boolean
  RoleArn: String
  Stages:
    - StageDeclaration
  Tags:
    - Tag
  Triggers:
    - PipelineTriggerDeclaration
  Variables:
    - VariableDeclaration

```

## Properties

`ArtifactStore`

The S3 bucket where artifacts for the pipeline are stored.

###### Note

You must include either `artifactStore` or
`artifactStores` in your pipeline, but you cannot use both. If you
create a cross-region action in your pipeline, you must use
`artifactStores`.

_Required_: Conditional

_Type_: [ArtifactStore](aws-properties-codepipeline-pipeline-artifactstore.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ArtifactStores`

A mapping of `artifactStore` objects and their corresponding AWS Regions. There must be an artifact store for the pipeline Region and for
each cross-region action in the pipeline.

###### Note

You must include either `artifactStore` or
`artifactStores` in your pipeline, but you cannot use both. If you
create a cross-region action in your pipeline, you must use
`artifactStores`.

_Required_: Conditional

_Type_: Array of [ArtifactStoreMap](aws-properties-codepipeline-pipeline-artifactstoremap.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisableInboundStageTransitions`

Represents the input of a `DisableStageTransition` action.

_Required_: No

_Type_: Array of [StageTransition](aws-properties-codepipeline-pipeline-stagetransition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionMode`

The method that the pipeline will use to handle multiple executions. The default
mode is SUPERSEDED.

_Required_: No

_Type_: String

_Allowed values_: `QUEUED | SUPERSEDED | PARALLEL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the pipeline.

_Required_: No

_Type_: String

_Pattern_: `[A-Za-z0-9.@\-_]+`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PipelineType`

CodePipeline provides the following pipeline types, which differ in
characteristics and price, so that you can tailor your pipeline features and cost to the
needs of your applications.

- V1 type pipelines have a JSON structure that contains standard pipeline,
stage, and action-level parameters.

- V2 type pipelines have the same structure as a V1 type, along with additional
parameters for release safety and trigger configuration.

###### Important

Including V2 parameters, such as triggers on Git tags, in the pipeline JSON when
creating or updating a pipeline will result in the pipeline having the V2 type of
pipeline and the associated costs.

For information about pricing for CodePipeline, see [Pricing](https://aws.amazon.com/codepipeline/pricing).

For information about which type of pipeline to choose, see [What type of\
pipeline is right for me?](../../../codepipeline/latest/userguide/pipeline-types-planning.md).

_Required_: No

_Type_: String

_Allowed values_: `V1 | V2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestartExecutionOnUpdate`

Indicates whether to rerun the CodePipeline pipeline after you update it.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) for CodePipeline to use to either perform
actions with no `actionRoleArn`, or to use to assume roles for actions with
an `actionRoleArn`.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws(-[\w]+)*:iam::[0-9]{12}:role/.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Stages`

Represents information about a stage and its definition.

_Required_: Yes

_Type_: Array of [StageDeclaration](aws-properties-codepipeline-pipeline-stagedeclaration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Specifies the tags applied to the pipeline.

_Required_: No

_Type_: Array of [Tag](aws-properties-codepipeline-pipeline-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Triggers`

The trigger configuration specifying a type of event, such as Git tags, that starts
the pipeline.

###### Note

When a trigger configuration is specified, default change detection for
repository and branch commits is disabled.

_Required_: No

_Type_: Array of [PipelineTriggerDeclaration](aws-properties-codepipeline-pipeline-pipelinetriggerdeclaration.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Variables`

A list that defines the pipeline variables for a pipeline resource. Variable names can
have alphanumeric and underscore characters, and the values must match
`[A-Za-z0-9@\-_]+`.

_Required_: No

_Type_: Array of [VariableDeclaration](aws-properties-codepipeline-pipeline-variabledeclaration.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the pipeline name, such as
mysta-MyPipeline-A1BCDEFGHIJ2.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Version`

The version of the pipeline.

###### Note

A new pipeline is always assigned a version number of 1. This number increments when
a pipeline is updated.

## Examples

- [Pipeline with two stages and trigger configuration](#aws-resource-codepipeline-pipeline--examples--Pipeline_with_two_stages_and_trigger_configuration)

- [Pipeline with source stage and a stage configured for automatic rollback on failure](#aws-resource-codepipeline-pipeline--examples--Pipeline_with_source_stage_and_a_stage_configured_for_automatic_rollback_on_failure)

### Pipeline with two stages and trigger configuration

The following example creates a pipeline with a source stage and a beta stage. For the
source stage, CodePipeline uses a connection to a GitHub repository. The beta stage builds
those changes by using CodeBuild. The pipeline is configured to start using trigger
filtering. The pipeline will start when push events meet the Git tags, branches, and file
path filter criteria specified. Also, the pipeline will start when pull requests meet the
filter criteria for the branch names, file paths, and pull request events specified. For
example, when a pull request for branches and file paths containing `release-*`
is closed, the pipeline will start.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "CodePipeline sample",
    "Parameters": {
        "CodePipelineServiceRole": {
            "Type": "String"
        },
        "ArtifactStoreS3Location": {
            "Type": "String"
        },
        "SourceActionName": {
            "Type": "String"
        }
    },
    "Resources": {
        "CodePipelineForIntegTest": {
            "Type": "AWS::CodePipeline::Pipeline",
            "Properties": {
                "RoleArn": {
                    "Ref": "CodePipelineServiceRole"
                },
                "Stages": [
                    {
                        "Name": "Source",
                        "Actions": [
                            {
                                "Name": {
                                    "Ref": "SourceActionName"
                                },
                                "InputArtifacts": [],
                                "ActionTypeId": {
                                    "Category": "Source",
                                    "Owner": "AWS",
                                    "Version": "1",
                                    "Provider": "CodeStarSourceConnection"
                                },
                                "OutputArtifacts": [
                                    {
                                        "Name": "SourceOutput"
                                    }
                                ],
                                "Configuration": {
                                    "BranchName": "main",
                                    "ConnectionArn": "arn:aws:codestar-connections:us-east-1:123456789123:connection/id",
                                    "FullRepositoryId": "repo-owner/sample-project"
                                },
                                "RunOrder": 1
                            }
                        ]
                    },
                    {
                        "Name": "Beta",
                        "Actions": [
                            {
                                "Name": "BetaAction",
                                "InputArtifacts": [
                                    {
                                        "Name": "SourceOutput"
                                    }
                                ],
                                "ActionTypeId": {
                                    "Category": "Build",
                                    "Owner": "AWS",
                                    "Provider": "CodeBuild",
                                    "Version": "1"
                                },
                                "Configuration": {
                                    "ProjectName": "Sample"
                                },
                                "RunOrder": 1
                            }
                        ]
                    }
                ],
                "Triggers": [
                    {
                        "ProviderType": "CodeStarSourceConnection",
                        "GitConfiguration": {
                            "Push": [
                                {
                                    "Tags": {
                                        "Excludes": [
                                            "beta-*"
                                        ],
                                        "Includes": [
                                            "release-*"
                                        ]
                                    }
                                },
                                {
                                    "Branches": {
                                        "Excludes": [
                                            "beta-*"
                                        ],
                                        "Includes": [
                                            "release-*"
                                        ]
                                    },
                                    "FilePaths": {
                                        "Includes": [
                                            "projectA/**",
                                            "common/**/*.js"
                                        ],
                                        "Excludes": [
                                            "**/README.md",
                                            "**/LICENSE",
                                            "**/CONTRIBUTING.md"
                                        ]
                                    }
                                }
                            ],
                            "PullRequest": [
                                {
                                    "Branches": {
                                        "Excludes": [
                                            "stable-v1-*"
                                        ],
                                        "Includes": [
                                            "stable-*",
                                            "release-*"
                                        ]
                                    },
                                    "FilePaths": {
                                        "Includes": [
                                            "projectA/**",
                                            "common/**/*.js"
                                        ],
                                        "Excludes": [
                                            "**/README.md",
                                            "**/LICENSE",
                                            "**/CONTRIBUTING.md"
                                        ]
                                    },
                                    "Events": [
                                        "CLOSED"
                                    ]
                                }
                            ],
                            "SourceActionName": {
                                "Ref": "SourceActionName"
                            }
                        }
                    }
                ],
                "PipelineType": "V2",
                "ExecutionMode": "PARALLEL",
                "ArtifactStore": {
                    "Type": "S3",
                    "Location": {
                        "Ref": "ArtifactStoreS3Location"
                    }
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: CodePipeline sample
Parameters:
  CodePipelineServiceRole:
    Type: String
  ArtifactStoreS3Location:
    Type: String
  SourceActionName:
    Type: String
Resources:
  CodePipelineForIntegTest:
    Type: 'AWS::CodePipeline::Pipeline'
    Properties:
      RoleArn: !Ref CodePipelineServiceRole
      Stages:
        - Name: Source
          Actions:
            - Name: !Ref SourceActionName
              InputArtifacts: []
              ActionTypeId:
                Category: Source
                Owner: AWS
                Version: '1'
                Provider: CodeStarSourceConnection
              OutputArtifacts:
                - Name: SourceOutput
              Configuration:
                BranchName: main
                ConnectionArn: >-
                  arn:aws:codestar-connections:us-east-1:123456789123:connection/id
                FullRepositoryId: repo-owner/sample-project
              RunOrder: 1
        - Name: Beta
          Actions:
            - Name: BetaAction
              InputArtifacts:
                - Name: SourceOutput
              ActionTypeId:
                Category: Build
                Owner: AWS
                Provider: CodeBuild
                Version: '1'
              Configuration:
                ProjectName: Sample
              RunOrder: 1
      Triggers:
        - ProviderType: CodeStarSourceConnection
          GitConfiguration:
            Push:
              - Tags:
                  Excludes:
                    - beta-*
                  Includes:
                    - release-*
              - Branches:
                  Excludes:
                    - beta-*
                  Includes:
                    - release-*
                FilePaths:
                  Includes:
                    - projectA/**
                    - common/**/*.js
                  Excludes:
                    - '**/README.md'
                    - '**/LICENSE'
                    - '**/CONTRIBUTING.md'
            PullRequest:
              - Branches:
                  Excludes:
                    - stable-v1-*
                  Includes:
                    - stable-*
                    - release-*
                FilePaths:
                  Includes:
                    - projectA/**
                    - common/**/*.js
                  Excludes:
                    - '**/README.md'
                    - '**/LICENSE'
                    - '**/CONTRIBUTING.md'
                Events:
                  - CLOSED
            SourceActionName: !Ref SourceActionName
      PipelineType: V2
      ExecutionMode: PARALLEL
      ArtifactStore:
        Type: S3
        Location: !Ref ArtifactStoreS3Location

```

### Pipeline with source stage and a stage configured for automatic rollback on failure

The following example creates a pipeline with a source stage and a release
stage.

#### JSON

```json

{
    "AppPipeline": {
        "Type": "AWS::CodePipeline::Pipeline",
        "Properties": {
            "RoleArn": {
                "Ref": "CodePipelineServiceRole"
            },
            "Stages": [
                {
                    "Name": "Source",
                    "Actions": [
                        {
                            "Name": "SourceAction",
                            "ActionTypeId": {
                                "Category": "Source",
                                "Owner": "AWS",
                                "Version": 1,
                                "Provider": "S3"
                            },
                            "OutputArtifacts": [
                                {
                                    "Name": "SourceOutput"
                                }
                            ],
                            "Configuration": {
                                "S3Bucket": {
                                    "Ref": "SourceS3Bucket"
                                },
                                "S3ObjectKey": {
                                    "Ref": "SourceS3ObjectKey"
                                }
                            },
                            "RunOrder": 1
                        }
                    ]
                },
                {
                    "Name": "Release",
                    "Actions": [
                        {
                            "Name": "ReleaseAction",
                            "InputArtifacts": [
                                {
                                    "Name": "SourceOutput"
                                }
                            ],
                            "ActionTypeId": {
                                "Category": "Deploy",
                                "Owner": "AWS",
                                "Version": 1,
                                "Provider": "CodeDeploy"
                            },
                            "Configuration": {
                                "ApplicationName": {
                                    "Ref": "ApplicationName"
                                },
                                "DeploymentGroupName": {
                                    "Ref": "DeploymentGroupName"
                                }
                            },
                            "RunOrder": 1
                        }
                    ],
                    "OnFailure": {
                        "Result": "ROLLBACK"
                    }
                }
            ],
            "ArtifactStore": {
                "Type": "S3",
                "Location": {
                    "Ref": "ArtifactStoreS3Location"
                },
                "EncryptionKey": {
                    "Id": "arn:aws:kms:useast-1:ACCOUNT-ID:key/KEY-ID",
                    "Type": "KMS"
                }
            },
            "DisableInboundStageTransitions": [
                {
                    "StageName": "Release",
                    "Reason": "Disabling the transition until integration tests are completed"
                }
            ],
            "Tags": [
                {
                    "Key": "Project",
                    "Value": "ProjectA"
                },
                {
                    "Key": "IsContainerBased",
                    "Value": "true"
                }
            ]
        }
    }
}
```

#### YAML

```yaml

AppPipeline:
  Type: AWS::CodePipeline::Pipeline
  Properties:
    RoleArn:
      Ref: CodePipelineServiceRole
    Stages:
      -
        Name: Source
        Actions:
          -
            Name: SourceAction
            ActionTypeId:
              Category: Source
              Owner: AWS
              Version: 1
              Provider: S3
            OutputArtifacts:
              -
                Name: SourceOutput
            Configuration:
              S3Bucket:
                Ref: SourceS3Bucket
              S3ObjectKey:
                Ref: SourceS3ObjectKey
            RunOrder: 1
      -
        Name: Release
        Actions:
          -
            Name: ReleaseAction
            InputArtifacts:
              -
                Name: SourceOutput
            ActionTypeId:
              Category: Deploy
              Owner: AWS
              Version: 1
              Provider: CodeDeploy
            Configuration:
              ApplicationName:
                Ref: ApplicationName
              DeploymentGroupName:
                Ref: DeploymentGroupName
            RunOrder: 1
        OnFailure:
            Result: ROLLBACK
    ArtifactStore:
      Type: S3
      Location:
        Ref: ArtifactStoreS3Location
      EncryptionKey:
        Id: arn:aws:kms:useast-1:ACCOUNT-ID:key/KEY-ID
        Type: KMS
    DisableInboundStageTransitions:
      -
        StageName: Release
        Reason: "Disabling the transition until integration tests are completed"
    Tags:
      - Key: Project
        Value: ProjectA
      - Key: IsContainerBased
        Value: 'true'

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ActionDeclaration

All content copied from https://docs.aws.amazon.com/.
