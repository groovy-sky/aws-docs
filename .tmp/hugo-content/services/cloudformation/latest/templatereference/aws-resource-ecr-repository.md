This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::Repository

The `AWS::ECR::Repository` resource specifies an Amazon Elastic Container
Registry (Amazon ECR) repository, where users can push and pull Docker images, Open
Container Initiative (OCI) images, and OCI compatible artifacts. For more information,
see [Amazon ECR private repositories](../../../amazonecr/latest/userguide/repositories.md) in the _Amazon ECR User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ECR::Repository",
  "Properties" : {
      "EmptyOnDelete" : Boolean,
      "EncryptionConfiguration" : EncryptionConfiguration,
      "ImageScanningConfiguration" : ImageScanningConfiguration,
      "ImageTagMutability" : String,
      "ImageTagMutabilityExclusionFilters" : [ ImageTagMutabilityExclusionFilter, ... ],
      "LifecyclePolicy" : LifecyclePolicy,
      "RepositoryName" : String,
      "RepositoryPolicyText" : Json,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ECR::Repository
Properties:
  EmptyOnDelete: Boolean
  EncryptionConfiguration:
    EncryptionConfiguration
  ImageScanningConfiguration:
    ImageScanningConfiguration
  ImageTagMutability: String
  ImageTagMutabilityExclusionFilters:
    - ImageTagMutabilityExclusionFilter
  LifecyclePolicy:
    LifecyclePolicy
  RepositoryName: String
  RepositoryPolicyText: Json
  Tags:
    - Tag

```

## Properties

`EmptyOnDelete`

If true, deleting the repository force deletes the contents of the repository. If
false, the repository must be empty before attempting to delete it.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfiguration`

The encryption configuration for the repository. This determines how the contents of
your repository are encrypted at rest.

_Required_: No

_Type_: [EncryptionConfiguration](aws-properties-ecr-repository-encryptionconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ImageScanningConfiguration`

###### Important

The `imageScanningConfiguration` parameter is being deprecated, in
favor of specifying the image scanning configuration at the registry level. For more
information, see `PutRegistryScanningConfiguration`.

The image scanning configuration for the repository. This determines whether images
are scanned for known vulnerabilities after being pushed to the repository.

_Required_: No

_Type_: [ImageScanningConfiguration](aws-properties-ecr-repository-imagescanningconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageTagMutability`

The tag mutability setting for the repository. If this parameter is omitted, the
default setting of `MUTABLE` will be used which will allow image tags to be
overwritten. If `IMMUTABLE` is specified, all image tags within the
repository will be immutable which will prevent them from being overwritten.

_Required_: No

_Type_: String

_Allowed values_: `MUTABLE | IMMUTABLE | MUTABLE_WITH_EXCLUSION | IMMUTABLE_WITH_EXCLUSION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageTagMutabilityExclusionFilters`

A list of filters that specify which image tags are excluded from the repository's
image tag mutability setting.

_Required_: No

_Type_: Array of [ImageTagMutabilityExclusionFilter](aws-properties-ecr-repository-imagetagmutabilityexclusionfilter.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LifecyclePolicy`

Creates or updates a lifecycle policy. For information about lifecycle policy syntax,
see [Lifecycle policy template](../../../amazonecr/latest/userguide/lifecyclepolicies.md).

_Required_: No

_Type_: [LifecyclePolicy](aws-properties-ecr-repository-lifecyclepolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepositoryName`

The name to use for the repository. The repository name may be specified on its own
(such as `nginx-web-app`) or it can be prepended with a namespace to group
the repository into a category (such as `project-a/nginx-web-app`). If you
don't specify a name, AWS CloudFormation generates a unique physical ID and uses
that ID for the repository name. For more information, see [Name\
type](../userguide/aws-properties-name.md).

The repository name must start with a letter and can only contain lowercase letters,
numbers, hyphens, underscores, and forward slashes.

###### Note

If you specify a name, you cannot perform updates that require replacement of this
resource. You can perform updates that require no or some interruption. If you must
replace the resource, specify a new name.

_Required_: No

_Type_: String

_Pattern_: `^(?=.{2,256}$)([a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*)$`

_Minimum_: `2`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RepositoryPolicyText`

The JSON repository policy text to apply to the repository. For more information, see
[Amazon ECR repository\
policies](../../../amazonecr/latest/userguide/repository-policy-examples.md) in the _Amazon Elastic Container Registry User Guide_.

_Required_: No

_Type_: Json

_Minimum_: `0`

_Maximum_: `10240`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-ecr-repository-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name, such as
`test-repository`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) for the specified
`AWS::ECR::Repository` resource. For example,
`arn:aws:ecr:eu-west-1:123456789012:repository/test-repository`.

`RepositoryUri`

Returns the URI for the specified `AWS::ECR::Repository` resource. For
example,
`123456789012.dkr.ecr.us-west-2.amazonaws.com/repository`.

## Examples

- [Specify a repository](#aws-resource-ecr-repository--examples--Specify_a_repository)

- [Specify a repository with an image scanning configuration](#aws-resource-ecr-repository--examples--Specify_a_repository_with_an_image_scanning_configuration)

- [Specify a repository with a lifecycle policy](#aws-resource-ecr-repository--examples--Specify_a_repository_with_a_lifecycle_policy)

### Specify a repository

The following example specifies a repository named
`test-repository`. Its policy permits the users `Bob`
and `Alice` to push and pull images. Note that the IAM
users actually need to exist, or stack creation will fail.

#### JSON

```json

"MyRepository": {
  "Type": "AWS::ECR::Repository",
  "Properties": {
    "RepositoryName" : "test-repository",
    "RepositoryPolicyText" : {
      "Version": "2012-10-17",
      "Statement": [
        {
          "Sid": "AllowPushPull",
          "Effect": "Allow",
          "Principal": {
            "AWS": [
              "arn:aws:iam::123456789012:user/Bob",
              "arn:aws:iam::123456789012:user/Alice"
            ]
          },
          "Action": [
            "ecr:GetDownloadUrlForLayer",
            "ecr:BatchGetImage",
            "ecr:BatchCheckLayerAvailability",
            "ecr:PutImage",
            "ecr:InitiateLayerUpload",
            "ecr:UploadLayerPart",
            "ecr:CompleteLayerUpload"
          ]
        }
      ]
    }
  }
}
```

#### YAML

```yaml

MyRepository:
  Type: AWS::ECR::Repository
  Properties:
    RepositoryName: "test-repository"
    RepositoryPolicyText:
      Version: "2012-10-17"
      Statement:
        -
          Sid: AllowPushPull
          Effect: Allow
          Principal:
            AWS:
              - "arn:aws:iam::123456789012:user/Bob"
              - "arn:aws:iam::123456789012:user/Alice"
          Action:
            - "ecr:GetDownloadUrlForLayer"
            - "ecr:BatchGetImage"
            - "ecr:BatchCheckLayerAvailability"
            - "ecr:PutImage"
            - "ecr:InitiateLayerUpload"
            - "ecr:UploadLayerPart"
            - "ecr:CompleteLayerUpload"
```

### Specify a repository with an image scanning configuration

The following example creates a repository named `test-repository`
with image scanning enabled. For more information on image scanning, see [Image scanning](../../../amazonecr/latest/userguide/image-scanning.md) in the _Amazon ECR User_
_Guide_.

#### JSON

```json

"MyRepository": {
  "Type": "AWS::ECR::Repository",
  "Properties": {
    "RepositoryName" : "test-repository",
    "ImageScanningConfiguration" : {
      "ScanOnPush": true
    }
  }
}
```

#### YAML

```yaml

MyRepository:
  Type: AWS::ECR::Repository
  Properties:
    RepositoryName: "test-repository"
    ImageScanningConfiguration:
      ScanOnPush: true
```

### Specify a repository with a lifecycle policy

The following example creates a repository with a lifecycle policy.

#### JSON

```json

{
  "Parameters": {
    "lifecyclePolicyText": {
      "Type": "String"
    },
    "repositoryName": {
      "Type": "String"
    },
    "registryId": {
      "Type": "String"
    }
  },
  "Resources": {
    "MyRepository": {
      "Type": "AWS::ECR::Repository",
      "Properties": {
        "LifecyclePolicy": {
          "LifecyclePolicyText": {
            "Ref": "lifecyclePolicyText"
          },
          "RegistryId": {
            "Ref": "registryId"
          }
        },
        "RepositoryName": {
          "Ref": "repositoryName"
        }
      }
    }
  },
  "Outputs": {
    "Arn": {
      "Value": {
        "Fn::GetAtt": [
          "MyRepository",
          "Arn"
        ]
      }
    }
  }
}
```

#### YAML

```yaml

Parameters:
  lifecyclePolicyText:
    Type: String
  repositoryName:
    Type: String
  registryId:
    Type: String
Resources:
  MyRepository:
    Type: AWS::ECR::Repository
    Properties:
      LifecyclePolicy:
        LifecyclePolicyText: !Ref lifecyclePolicyText
        RegistryId: !Ref registryId
      RepositoryName: !Ref repositoryName
Outputs:
  Arn:
    Value: !GetAtt MyRepository.Arn
```

## See also

- [Creating a lifecycle\
policy](../../../amazonecr/latest/userguide/lp-creation.md) in the _Amazon ECR User Guide_

- [PutLifecyclePolicy](../../../../reference/amazonecr/latest/apireference/api-putlifecyclepolicy.md) in the _Amazon ECR API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RepositoryFilter

EncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
