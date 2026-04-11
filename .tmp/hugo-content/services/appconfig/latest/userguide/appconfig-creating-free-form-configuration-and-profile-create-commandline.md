# Creating an AWS AppConfig freeform configuration profile (command line)

The following procedure describes how to use the AWS CLI (on Linux or Windows) or
AWS Tools for PowerShell to create an AWS AppConfig freeform configuration profile. If you prefer, you can
use AWS CloudShell to run the commands listed below. For more information, see [What is\
AWS CloudShell?](../../../cloudshell/latest/userguide/welcome.md) in the _AWS CloudShell User Guide_.

###### Note

For freeform configurations hosted in the AWS AppConfig hosted configuration store, you
specify `hosted` for the location URI.

###### To create a configuration profile by using the AWS CLI

1. Open the AWS CLI.

2. Run the following command to create a freeform configuration profile.
Linux

```nohighlight

aws appconfig create-configuration-profile \
     --application-id APPLICATION_ID \
     --name NAME \
     --description CONFIGURATION_PROFILE_DESCRIPTION \
     --location-uri CONFIGURATION_URI or hosted \
     --retrieval-role-arn IAM_ROLE_ARN \
     --tags TAGS \
     --validators "Content=SCHEMA_CONTENT or LAMBDA_FUNCTION_ARN,Type=JSON_SCHEMA or LAMBDA"
```

Windows

```nohighlight

aws appconfig create-configuration-profile ^
     --application-id APPLICATION_ID ^
     --name NAME ^
     --description CONFIGURATION_PROFILE_DESCRIPTION ^
     --location-uri CONFIGURATION_URI or hosted  ^
     --retrieval-role-arn IAM_ROLE_ARN ^
     --tags TAGS ^
     --validators "Content=SCHEMA_CONTENT or LAMBDA_FUNCTION_ARN,Type=JSON_SCHEMA or LAMBDA"
```

PowerShell

```powershell

New-APPCConfigurationProfile `
     -Name NAME `
     -ApplicationId APPLICATION_ID `
     -Description CONFIGURATION_PROFILE_DESCRIPTION `
     -LocationUri CONFIGURATION_URI or hosted `
     -RetrievalRoleArn IAM_ROLE_ARN `
     -Tag TAGS `
     -Validators "Content=SCHEMA_CONTENT or LAMBDA_FUNCTION_ARN,Type=JSON_SCHEMA or LAMBDA"
```

###### Important

Note the following important information.

- If you created a configuration profile for AWS CodePipeline, then you must create a
pipeline in CodePipeline that specifies AWS AppConfig as the _deploy provider_.
You don't need to perform [Deploying feature flags and configuration data in AWS AppConfig](deploying-feature-flags.md). However, you must configure a client to
receive application configuration updates as described in [Retrieving configuration data without AWS AppConfig Agent](about-data-plane.md). For information
about creating a pipeline that specifies AWS AppConfig as the deploy provider, see [Tutorial: Create a Pipeline that Uses\
AWS AppConfig as a Deployment Provider](../../../codepipeline/latest/userguide/tutorials-appconfig.md) in the
_AWS CodePipeline User Guide_.

- If you created a configuration in the AWS AppConfig hosted configuration store, you can
create new versions of the configuration by using the [CreateHostedConfigurationVersion](../../../../reference/appconfig/2019-10-09/apireference/api-createhostedconfigurationversion.md) API operations. To view AWS CLI details
and sample commands for this API operation, see [create-hosted-configuration-version](../../../cli/latest/reference/appconfig/create-hosted-configuration-version.md) in the
_AWS CLI Command Reference_.

Proceed to [Deploying feature flags and configuration data in AWS AppConfig](deploying-feature-flags.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Creating an AWS AppConfig freeform configuration profile (console)

Creating a configuration profile for non-native data sources

All content copied from https://docs.aws.amazon.com/.
