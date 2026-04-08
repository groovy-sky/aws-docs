Menu

- [Aws](namespace-aws.md)

## Credentials

### Table of Contents  [header link](namespace-aws-credentials-toc.md)

#### Interfaces  [header link](namespace-aws-credentials-toc-interfaces.md)

[CredentialsInterface](class-aws-credentials-credentialsinterface.md)Provides access to the AWS credentials used for accessing AWS services: AWS
access key ID, secret access key, and security token. These credentials are
used to securely sign requests to AWS services.

#### Classes  [header link](namespace-aws-credentials-toc-classes.md)

[AssumeRoleCredentialProvider](class-aws-credentials-assumerolecredentialprovider.md)Credential provider that provides credentials via assuming a role
More Information, see: http://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sts-2011-06-15.html#assumerole[AssumeRoleWithWebIdentityCredentialProvider](class-aws-credentials-assumerolewithwebidentitycredentialprovider.md)Credential provider that provides credentials via assuming a role with a web identity
More Information, see: https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sts-2011-06-15.html#assumerolewithwebidentity[CredentialProvider](class-aws-credentials-credentialprovider.md)Credential providers are functions that accept no arguments and return a
promise that is fulfilled with an {@see \\Aws\\Credentials\\CredentialsInterface}
or rejected with an {@see \\Aws\\Exception\\CredentialsException}.[Credentials](class-aws-credentials-credentials.md)Basic implementation of the AWS Credentials interface that allows callers to
pass in the AWS Access Key and AWS Secret Access Key in the constructor.[CredentialsUtils](class-aws-credentials-credentialsutils.md)[EcsCredentialProvider](class-aws-credentials-ecscredentialprovider.md)Credential provider that fetches container credentials with GET request.[InstanceProfileProvider](class-aws-credentials-instanceprofileprovider.md)Credential provider that provides credentials from the EC2 metadata service.[LoginCredentialProvider](class-aws-credentials-logincredentialprovider.md)Credential provider for login using console credentials

```

```

×

**On this page**

- Table Of Contents
  - [Interfaces](namespace-aws-credentials-toc-interfaces.md)
  - [Classes](namespace-aws-credentials-toc-classes.md)

[Back To Top](namespace-aws-credentials-top.md)
