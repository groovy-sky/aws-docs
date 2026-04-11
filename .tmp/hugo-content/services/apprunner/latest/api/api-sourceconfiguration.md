# SourceConfiguration

Describes the source deployed to an AWS App Runner service. It can be a code or an image repository.

## Contents

**AuthenticationConfiguration**

Describes the resources that are needed to authenticate access to some source repositories.

Type: [AuthenticationConfiguration](api-authenticationconfiguration.md) object

Required: No

**AutoDeploymentsEnabled**

If `true`, continuous integration from the source repository is enabled for the App Runner service. Each repository change (including any source
code commit or new image version) starts a deployment.

Default: App Runner sets to `false` for a source image that uses an ECR Public repository or an ECR repository that's in an AWS account other than the one that the service is in. App Runner sets to `true` in all other cases (which currently include a source code
repository or a source image using a same-account ECR repository).

Type: Boolean

Required: No

**CodeRepository**

The description of a source code
repository.

You must provide either this member or `ImageRepository` (but not both).

Type: [CodeRepository](api-coderepository.md) object

Required: No

**ImageRepository**

The description of a source image
repository.

You must provide either this member or `CodeRepository` (but not both).

Type: [ImageRepository](api-imagerepository.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/sourceconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/sourceconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/sourceconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceCodeVersion

Tag

All content copied from https://docs.aws.amazon.com/.
