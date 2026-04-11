This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::PullThroughCacheRule

The `AWS::ECR::PullThroughCacheRule` resource creates or updates a pull
through cache rule. A pull through cache rule provides a way to cache images from an
upstream registry in your Amazon ECR private registry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ECR::PullThroughCacheRule",
  "Properties" : {
      "CredentialArn" : String,
      "CustomRoleArn" : String,
      "EcrRepositoryPrefix" : String,
      "UpstreamRegistry" : String,
      "UpstreamRegistryUrl" : String,
      "UpstreamRepositoryPrefix" : String
    }
}

```

### YAML

```yaml

Type: AWS::ECR::PullThroughCacheRule
Properties:
  CredentialArn: String
  CustomRoleArn: String
  EcrRepositoryPrefix: String
  UpstreamRegistry: String
  UpstreamRegistryUrl: String
  UpstreamRepositoryPrefix: String

```

## Properties

`CredentialArn`

The ARN of the Secrets Manager secret associated with the pull through cache rule.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws:secretsmanager:[a-zA-Z0-9-:]+:secret:ecr\-pullthroughcache\/[a-zA-Z0-9\/_+=.@-]+$`

_Minimum_: `50`

_Maximum_: `612`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomRoleArn`

The ARN of the IAM role associated with the pull through cache rule.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EcrRepositoryPrefix`

The Amazon ECR repository prefix associated with the pull through cache rule.

_Required_: No

_Type_: String

_Pattern_: `^([a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*\/?|ROOT)$`

_Minimum_: `2`

_Maximum_: `30`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UpstreamRegistry`

The name of the upstream source registry associated with the pull through cache
rule.

_Required_: No

_Type_: String

_Allowed values_: `ecr | ecr-public | quay | k8s | docker-hub | github-container-registry | azure-container-registry | gitlab-container-registry`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UpstreamRegistryUrl`

The upstream registry URL associated with the pull through cache rule.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UpstreamRepositoryPrefix`

The upstream repository prefix associated with the pull through cache rule.

_Required_: No

_Type_: String

_Pattern_: `^([a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*\/?|ROOT)$`

_Minimum_: `2`

_Maximum_: `30`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

The following resource examples show how to create a pull through cache rule for a
private registry.

- [Create a pull through cache rule for an upstream registry that requires authentication](#aws-resource-ecr-pullthroughcacherule--examples--Create_a_pull_through_cache_rule_for_an_upstream_registry_that_requires_authentication)

- [Create a pull through cache rule for an upstream registry that does not require authentication](#aws-resource-ecr-pullthroughcacherule--examples--Create_a_pull_through_cache_rule_for_an_upstream_registry_that_does_not_require_authentication)

### Create a pull through cache rule for an upstream registry that requires authentication

The following example creates a pull through cache rule for the upstream
registry Docker Hub, which requires authentication. The authentication
credentials for the upstream registry must be stored in a Secrets Manager secret
with a secret name with a `ecr-pullthroughcache/` prefix. You specify
the full Amazon Resource Name (ARN) of the secret. When the pull through cache
rule is used to pull images from the upstream registry, Amazon ECR will create
repositories in your private registry on your behalf with the
`docker-hub` prefix.

#### JSON

```json

{
    "Resources": {
        "MyECRPullThroughCacheRule": {
            "Type": "AWS::ECR::PullThroughCacheRule",
            "Properties": {
                "EcrRepositoryPrefix": "docker-hub",
                "UpstreamRegistryUrl": "registry-1.docker.io",
                "CredentialArn": "arn:aws:secretsmanager:us-east-2:111122223333:secret:ecr-pullthroughcache/example1234"
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  MyECRPullThroughCacheRule:
    Type: 'AWS::ECR::PullThroughCacheRule'
    Properties:
      EcrRepositoryPrefix: 'docker-hub'
      UpstreamRegistryUrl: 'registry-1.docker.io'
      CredentialArn: 'arn:aws:secretsmanager:us-east-2:111122223333:secret:ecr-pullthroughcache/example1234'
      UpstreamRegistry: 'docker-hub'
```

### Create a pull through cache rule for an upstream registry that does not require authentication

The following example creates a pull through cache rule that caches
repositories with the name prefix `ecr-public` from the Amazon ECR
Public registry into your private registry.

#### JSON

```json

{
    "Resources": {
        "MyECRPullThroughCacheRule": {
            "Type": "AWS::ECR::PullThroughCacheRule",
            "Properties": {
                "EcrRepositoryPrefix": "ecr-public",
                "UpstreamRegistryUrl": "public.ecr.aws"
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  MyECRPullThroughCacheRule:
    Type: 'AWS::ECR::PullThroughCacheRule'
    Properties:
      EcrRepositoryPrefix: 'ecr-public'
      UpstreamRegistryUrl: 'public.ecr.aws'
      UpstreamRegistry: 'ecr-public'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::ECR::PullTimeUpdateExclusion

All content copied from https://docs.aws.amazon.com/.
