Menu

- [Aws](namespace-aws.md)

## Credentials

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.credentials.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.credentials.html\#toc-interfaces)

[CredentialsInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsInterface.html)Provides access to the AWS credentials used for accessing AWS services: AWS
access key ID, secret access key, and security token. These credentials are
used to securely sign requests to AWS services.

#### Classes  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.credentials.html\#toc-classes)

[AssumeRoleCredentialProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleCredentialProvider.html)Credential provider that provides credentials via assuming a role
More Information, see: http://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sts-2011-06-15.html#assumerole[AssumeRoleWithWebIdentityCredentialProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.AssumeRoleWithWebIdentityCredentialProvider.html)Credential provider that provides credentials via assuming a role with a web identity
More Information, see: https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sts-2011-06-15.html#assumerolewithwebidentity[CredentialProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialProvider.html)Credential providers are functions that accept no arguments and return a
promise that is fulfilled with an {@see \\Aws\\Credentials\\CredentialsInterface}
or rejected with an {@see \\Aws\\Exception\\CredentialsException}.[Credentials](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.Credentials.html)Basic implementation of the AWS Credentials interface that allows callers to
pass in the AWS Access Key and AWS Secret Access Key in the constructor.[CredentialsUtils](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.CredentialsUtils.html)[EcsCredentialProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.EcsCredentialProvider.html)Credential provider that fetches container credentials with GET request.[InstanceProfileProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.InstanceProfileProvider.html)Credential provider that provides credentials from the EC2 metadata service.[LoginCredentialProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Credentials.LoginCredentialProvider.html)Credential provider for login using console credentials

```

```

×

**On this page**

- Table Of Contents
  - [Interfaces](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.credentials.html#toc-interfaces)
  - [Classes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.credentials.html#toc-classes)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.credentials.html#top)
