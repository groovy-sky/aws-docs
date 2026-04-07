This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Project Environment

`Environment` is a property of the [AWS::CodeBuild::Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies
the environment for an AWS CodeBuild project.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Certificate" : String,
  "ComputeType" : String,
  "EnvironmentVariables" : [ EnvironmentVariable, ... ],
  "Fleet:" : ProjectFleet,
  "Image" : String,
  "ImagePullCredentialsType" : String,
  "PrivilegedMode" : Boolean,
  "RegistryCredential" : RegistryCredential,
  "Type" : String
}

```

### YAML

```yaml

  Certificate: String
  ComputeType: String
  EnvironmentVariables:
    - EnvironmentVariable
  Fleet::
    ProjectFleet
  Image: String
  ImagePullCredentialsType: String
  PrivilegedMode: Boolean
  RegistryCredential:
    RegistryCredential
  Type: String

```

## Properties

`Certificate`

The ARN of the Amazon S3 bucket, path prefix, and object key that contains the PEM-encoded
certificate for the build project. For more information, see [certificate](https://docs.aws.amazon.com/codebuild/latest/userguide/create-project-cli.html#cli.environment.certificate) in the
_AWS CodeBuild User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComputeType`

The type of compute environment. This determines the number of CPU cores and memory the build environment uses. Available values
include:

- `ATTRIBUTE_BASED_COMPUTE`: Specify the amount of vCPUs, memory, disk space, and the type of machine.

###### Note

If you use `ATTRIBUTE_BASED_COMPUTE`, you must define your attributes by using `computeConfiguration`. AWS CodeBuild
will select the cheapest instance that satisfies your specified attributes. For more information, see [Reserved capacity environment \
types](https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment-reserved-capacity.types) in the _AWS CodeBuild User Guide_.

- `BUILD_GENERAL1_SMALL`: Use up to 4 GiB memory and 2 vCPUs for
builds.

- `BUILD_GENERAL1_MEDIUM`: Use up to 8 GiB memory and 4 vCPUs for
builds.

- `BUILD_GENERAL1_LARGE`: Use up to 16 GiB memory and 8 vCPUs for
builds, depending on your environment type.

- `BUILD_GENERAL1_XLARGE`: Use up to 72 GiB memory and 36 vCPUs for
builds, depending on your environment type.

- `BUILD_GENERAL1_2XLARGE`: Use up to 144 GiB memory, 72 vCPUs, and
824 GB of SSD storage for builds. This compute type supports Docker images up to
100 GB uncompressed.

- `BUILD_LAMBDA_1GB`: Use up to 1 GiB memory for
builds. Only available for environment type `LINUX_LAMBDA_CONTAINER` and `ARM_LAMBDA_CONTAINER`.

- `BUILD_LAMBDA_2GB`: Use up to 2 GiB memory for
builds. Only available for environment type `LINUX_LAMBDA_CONTAINER` and `ARM_LAMBDA_CONTAINER`.

- `BUILD_LAMBDA_4GB`: Use up to 4 GiB memory for
builds. Only available for environment type `LINUX_LAMBDA_CONTAINER` and `ARM_LAMBDA_CONTAINER`.

- `BUILD_LAMBDA_8GB`: Use up to 8 GiB memory for
builds. Only available for environment type `LINUX_LAMBDA_CONTAINER` and `ARM_LAMBDA_CONTAINER`.

- `BUILD_LAMBDA_10GB`: Use up to 10 GiB memory for
builds. Only available for environment type `LINUX_LAMBDA_CONTAINER` and `ARM_LAMBDA_CONTAINER`.

If you use `BUILD_GENERAL1_SMALL`:

- For environment type `LINUX_CONTAINER`, you can use up to 4 GiB
memory and 2 vCPUs for builds.

- For environment type `LINUX_GPU_CONTAINER`, you can use up to 16
GiB memory, 4 vCPUs, and 1 NVIDIA A10G Tensor Core GPU for builds.

- For environment type `ARM_CONTAINER`, you can use up to 4 GiB
memory and 2 vCPUs on ARM-based processors for builds.

If you use `BUILD_GENERAL1_LARGE`:

- For environment type `LINUX_CONTAINER`, you can use up to 16 GiB
memory and 8 vCPUs for builds.

- For environment type `LINUX_GPU_CONTAINER`, you can use up to 255
GiB memory, 32 vCPUs, and 4 NVIDIA Tesla V100 GPUs for builds.

- For environment type `ARM_CONTAINER`, you can use up to 16 GiB
memory and 8 vCPUs on ARM-based processors for builds.

For more information, see [On-demand environment types](https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment.types)
in the _AWS CodeBuild User Guide._

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentVariables`

A set of environment variables to make available to builds for this build
project.

_Required_: No

_Type_: Array of [EnvironmentVariable](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codebuild-project-environmentvariable.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Fleet:`

Property description not available.

_Required_: No

_Type_: [ProjectFleet](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codebuild-project-projectfleet.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Image`

The image tag or image digest that identifies the Docker image to use for this build
project. Use the following formats:

- For an image tag: `<registry>/<repository>:<tag>`. For
example, in the Docker repository that CodeBuild uses to manage its Docker
images, this would be `aws/codebuild/standard:4.0`.

- For an image digest: `<registry>/<repository>@<digest>`.
For example, to specify an image with the digest
"sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0382cfbdbf," use
`<registry>/<repository>@sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0382cfbdbf`.

For more information, see [Docker images provided by CodeBuild](https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html) in the _AWS CodeBuild user_
_guide_.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImagePullCredentialsType`

The type of credentials AWS CodeBuild uses to pull images in your build. There are two valid
values:

- `CODEBUILD` specifies that AWS CodeBuild uses its own credentials.
This requires that you modify your ECR repository policy to trust AWS CodeBuild service principal.

- `SERVICE_ROLE` specifies that AWS CodeBuild uses your build project's service
role.

When you use a cross-account or private registry image, you must use SERVICE\_ROLE
credentials. When you use an AWS CodeBuild curated image, you must use CODEBUILD credentials.

_Required_: No

_Type_: String

_Allowed values_: `CODEBUILD | SERVICE_ROLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivilegedMode`

Enables running the Docker daemon inside a Docker container. Set to true only if the
build project is used to build Docker images. Otherwise, a build that attempts
to interact with the Docker daemon fails. The default setting is `false`.

You can initialize the Docker daemon during the install phase of your build by adding one of the following sets of commands
to the install phase of your buildspec file:

If the operating system's base image is Ubuntu Linux:

`- nohup /usr/local/bin/dockerd --host=unix:///var/run/docker.sock --host=tcp://0.0.0.0:2375 --storage-driver=overlay&`

`- timeout 15 sh -c "until docker info; do echo .; sleep 1; done"`

If the operating system's base image is Alpine Linux and the previous command does not work, add the `-t` argument
to `timeout`:

`- nohup /usr/local/bin/dockerd --host=unix:///var/run/docker.sock --host=tcp://0.0.0.0:2375 --storage-driver=overlay&`

`- timeout -t 15 sh -c "until docker info; do echo .; sleep 1; done"`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegistryCredential`

`RegistryCredential` is a property of the [AWS::CodeBuild::Project Environment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-environment) property that specifies information about credentials that provide access to a private Docker registry. When this is set:

- `imagePullCredentialsType` must be set to `SERVICE_ROLE`.

- images cannot be curated or an Amazon ECR image.

_Required_: No

_Type_: [RegistryCredential](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codebuild-project-registrycredential.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of build environment to use for related builds.

###### Note

If you're using compute fleets during project creation, `type` will be ignored.

For more information, see [Build environment compute types](https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html) in the _AWS CodeBuild_
_user guide_.

_Required_: Yes

_Type_: String

_Allowed values_: `WINDOWS_CONTAINER | LINUX_CONTAINER | LINUX_GPU_CONTAINER | ARM_CONTAINER | WINDOWS_SERVER_2019_CONTAINER | WINDOWS_SERVER_2022_CONTAINER | LINUX_LAMBDA_CONTAINER | ARM_LAMBDA_CONTAINER | LINUX_EC2 | ARM_EC2 | WINDOWS_EC2 | MAC_ARM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CloudWatchLogsConfig

EnvironmentVariable
