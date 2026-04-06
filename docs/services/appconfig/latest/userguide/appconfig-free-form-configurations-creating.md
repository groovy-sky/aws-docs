# Creating a free form configuration profile in AWS AppConfig

_Configuration data_ is a collection of settings that influence the
behavior of your application. A _configuration profile_ includes, among
other things, a URI that enables AWS AppConfig to locate your configuration data in its stored
location and a configure type. With freeform configuration profiles, you can store your data
in the AWS AppConfig hosted configuration store or any of the following AWS services and Systems Manager
tools:

LocationSupported file types

AWS AppConfig hosted configuration store

YAML, JSON, and text if added using the AWS Management Console. Any file type if added
using the AWS AppConfig [CreateHostedConfigurationVersion](https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_CreateHostedConfigurationVersion.html) API action.

[Amazon Simple Storage Service (Amazon S3)](../../../s3/latest/userguide/welcome.md)

Any

[AWS CodePipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/welcome.html)

Pipeline (as defined by the service)

[AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html)

Secret (as defined by the service)

[AWS Systems Manager Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html)

Standard and secure string parameters (as defined by Parameter Store)

[AWS Systems Manager document store\
(SSM documents)](https://docs.aws.amazon.com/systems-manager/latest/userguide/documents.html)

YAML, JSON, text

A configuration profile can also include optional validators to ensure your
configuration data is syntactically and semantically correct. AWS AppConfig performs a check using
the validators when you start a deployment. If any errors are detected, the deployment stops
before making any changes to the targets of the configuration.

###### Note

If possible, we recommend hosting your configuration data in the AWS AppConfig hosted
configuration store as it offers the most features and enhancements.

For freeform configurations stored in the AWS AppConfig hosted configuration store or SSM
documents, you can create the freeform configuration by using the Systems Manager console at the time
you create a configuration profile. The process is described later in this topic.

For freeform configurations stored in Parameter Store, Secrets Manager, or Amazon S3, you must create the
parameter, secret, or object first and store it in the relevant configuration store. After
you store the configuration data, use the procedure in this topic to create the
configuration profile.

###### Topics

- [Understanding validators](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-configuration-and-profile-validators.html)

- [Understanding configuration store quotas and limitations](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-configuration-and-profile-quotas.html)

- [Understanding the AWS AppConfig hosted configuration store](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-configuration-and-profile-about-hosted-store.html)

- [Understanding configurations stored in Amazon S3](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-configuration-and-profile-S3-source.html)

- [Creating an AWS AppConfig freeform configuration profile (console)](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-free-form-configuration-and-profile-create-console.html)

- [Creating an AWS AppConfig freeform configuration profile (command line)](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-free-form-configuration-and-profile-create-commandline.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Saving a previous feature flag version to a new version

Understanding validators
