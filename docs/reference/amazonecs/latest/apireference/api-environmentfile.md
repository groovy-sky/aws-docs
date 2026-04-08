# EnvironmentFile

A list of files containing the environment variables to pass to a container. You can
specify up to ten environment files. The file must have a `.env` file
extension. Each line in an environment file should contain an environment variable in
`VARIABLE=VALUE` format. Lines beginning with `#` are treated
as comments and are ignored.

If there are environment variables specified using the `environment`
parameter in a container definition, they take precedence over the variables contained
within an environment file. If multiple environment files are specified that contain the
same variable, they're processed from the top down. We recommend that you use unique
variable names. For more information, see [Use a file to pass\
environment variables to a container](../../../../services/amazonecs/latest/developerguide/use-environment-file.md) in the _Amazon Elastic_
_Container Service Developer Guide_.

Environment variable files are objects in Amazon S3 and all Amazon S3 security
considerations apply.

You must use the following platforms for the Fargate launch type:

- Linux platform version `1.4.0` or later.

- Windows platform version `1.0.0` or later.

Consider the following when using the Fargate launch type:

- The file is handled like a native Docker env-file.

- There is no support for shell escape handling.

- The container entry point interperts the `VARIABLE` values.

## Contents

**type**

The file type to use. Environment files are objects in Amazon S3. The only supported
value is `s3`.

Type: String

Valid Values: `s3`

Required: Yes

**value**

The Amazon Resource Name (ARN) of the Amazon S3 object containing the environment
variable file.

Type: String

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/environmentfile.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/environmentfile.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/environmentfile.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EFSVolumeConfiguration

EphemeralStorage
