AWS App Runner will no longer be open to new customers starting April 30, 2026. If you would like to use
App Runner, sign up prior to that date. Existing customers can continue to use the service as normal. For more information, see
[AWS App Runner availability\
change](../dg/apprunner-availability-change.md).

# Release: App Runner adds support for referencing secrets and configurations from AWS Secrets Manager and AWS Systems Manager Parameter Store on January 5, 2023

AWS App Runner now supports referencing secrets and configuration data stored in AWS Secrets Manager and AWS SSM Parameter Store in App Runner service.

**Release date:** January 5, 2023

## Changes

AWS App Runner now
supports securely referencing secrets and configuration data that are stored in Secrets Manager and
SSM  Parameter Store by adding them as environment variables in your App Runner service.

Secrets Manager and SSM  Parameter Store are AWS services that provide secure storage and reliable management of sensitive data such as configuration data
and secrets. Examples of data you can store include passwords, database strings or connection parameters, license codes, API keys, and application
versions. App Runner leverages the capabilities of Secrets Manager and SSM  Parameter Store to integrate a more secure experience for your applications using App Runner
service.

App Runner only stores reference to the Amazon Resource Name (ARN) of the secret or parameter stored in Secrets Manager or SSM  Parameter Store. This ensures that
your sensitive data isn't visible to others in App Runner service configurations and application logs. As such, secrets and parameters are managed in a way
that's completely isolated from your App Runner application code and your App Runner service configuration.

For more information, see [Referencing environment variables](../dg/env-variable.md) in the
_AWS App Runner Developer Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HTTP1.0 support 2023-02-01

2022 release notes

All content copied from https://docs.aws.amazon.com/.
