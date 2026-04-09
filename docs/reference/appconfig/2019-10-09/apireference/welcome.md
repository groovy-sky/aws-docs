# Welcome

## AWS AppConfig

AWS AppConfig feature flags and dynamic configurations help software builders
quickly and securely adjust application behavior in production environments without full
code deployments. AWS AppConfig speeds up software release frequency, improves
application resiliency, and helps you address emergent issues more quickly. With feature
flags, you can gradually release new capabilities to users and measure the impact of those
changes before fully deploying the new capabilities to all users. With operational flags
and dynamic configurations, you can update block lists, allow lists, throttling limits,
logging verbosity, and perform other operational tuning to quickly respond to issues in
production environments.

###### Note

AWS AppConfig is a tool in AWS Systems Manager.

Despite the fact that application configuration content can vary greatly from
application to application, AWS AppConfig supports the following use cases, which
cover a broad spectrum of customer needs:

- **Feature flags and toggles** \- Safely release new
capabilities to your customers in a controlled environment. Instantly roll back
changes if you experience a problem.

- **Application tuning** \- Carefully introduce
application changes while testing the impact of those changes with users in
production environments.

- **Allow list or block list** \- Control access to
premium features or instantly block specific users without deploying new code.

- **Centralized configuration storage** \- Keep your
configuration data organized and consistent across all of your workloads. You can use
AWS AppConfig to deploy configuration data stored in the AWS AppConfig
hosted configuration store, AWS Secrets Manager, Systems Manager, Parameter
Store, or Amazon S3.

**How AWS AppConfig works**

This section provides a high-level description of how AWS AppConfig works and how
you get started.

1\. Identify configuration values in code you want to manage in the cloud

Before you start creating AWS AppConfig artifacts, we recommend you
identify configuration data in your code that you want to dynamically manage using
AWS AppConfig. Good examples include feature flags or toggles, allow and
block lists, logging verbosity, service limits, and throttling rules, to name a
few.

If your configuration data already exists in the cloud, you can take advantage
of AWS AppConfig validation, deployment, and extension features to further
streamline configuration data management.

2\. Create an application namespace

To create a namespace, you create an AWS AppConfig artifact called an
application. An application is simply an organizational construct like a
folder.

3\. Create environments

For each AWS AppConfig application, you define one or more environments.
An environment is a logical grouping of targets, such as applications in a
`Beta` or `Production` environment, Lambda functions,
or containers. You can also define environments for application subcomponents,
such as the `Web`, `Mobile`, and
`Back-end`.

You can configure Amazon CloudWatch alarms for each environment. The system monitors
alarms during a configuration deployment. If an alarm is triggered, the system
rolls back the configuration.

4\. Create a configuration profile

A configuration profile includes, among other things, a URI that enables
AWS AppConfig to locate your configuration data in its stored location
and a profile type. AWS AppConfig supports two configuration profile types:
feature flags and freeform configurations. Feature flag configuration profiles
store their data in the AWS AppConfig hosted configuration store and the URI
is simply `hosted`. For freeform configuration profiles, you can store
your data in the AWS AppConfig hosted configuration store or any AWS
service that integrates with AWS AppConfig, as described in [Creating\
a free form configuration profile](../../../../services/appconfig/latest/userguide/appconfig-free-form-configurations-creating.md) in the the _AWS AppConfig User Guide_.

A configuration profile can also include optional validators to ensure your
configuration data is syntactically and semantically correct. AWS AppConfig
performs a check using the validators when you start a deployment. If any errors
are detected, the deployment rolls back to the previous configuration data.

5\. Deploy configuration data

When you create a new deployment, you specify the following:

- An application ID

- A configuration profile ID

- A configuration version

- An environment ID where you want to deploy the configuration data

- A deployment strategy ID that defines how fast you want the changes to
take effect

When you call the [StartDeployment](api-startdeployment.md) API action, AWS AppConfig performs the following
tasks:

1. Retrieves the configuration data from the underlying data store by using
    the location URI in the configuration profile.

2. Verifies the configuration data is syntactically and semantically correct
    by using the validators you specified when you created your configuration
    profile.

3. Caches a copy of the data so it is ready to be retrieved by your
    application. This cached copy is called the _deployed_
_data_.

6\. Retrieve the configuration

You can configure AWS AppConfig Agent as a local host and have the agent
poll AWS AppConfig for configuration updates. The agent calls the [StartConfigurationSession](api-appconfigdata-startconfigurationsession.md) and [GetLatestConfiguration](api-appconfigdata-getlatestconfiguration.md) API actions and caches your configuration data
locally. To retrieve the data, your application makes an HTTP call to the
localhost server. AWS AppConfig Agent supports several use cases, as
described in [Simplified\
retrieval methods](../../../../services/appconfig/latest/userguide/appconfig-retrieving-simplified-methods.md) in the the _AWS AppConfig User_
_Guide_.

If AWS AppConfig Agent isn't supported for your use case, you can
configure your application to poll AWS AppConfig for configuration updates
by directly calling the [StartConfigurationSession](api-appconfigdata-startconfigurationsession.md) and [GetLatestConfiguration](api-appconfigdata-getlatestconfiguration.md) API actions.

This reference is intended to be used with the [AWS AppConfig User\
Guide](../../../../services/appconfig/latest/userguide/what-is-appconfig.md).

## AWS AppConfig Data

AWS AppConfig Data provides the data plane APIs your application uses to retrieve
configuration data. Here's how it works:

Your application retrieves configuration data by first establishing a configuration
session using the AWS AppConfig Data
[StartConfigurationSession](api-appconfigdata-startconfigurationsession.md) API action.
Your session's client then makes periodic calls to [GetLatestConfiguration](api-appconfigdata-getlatestconfiguration.md)
to check for and retrieve the latest data available.

When calling `StartConfigurationSession`, your code sends the following
information:

- Identifiers (ID or name) of an AWS AppConfig application, environment, and
configuration profile that the session tracks.

- (Optional) The minimum amount of time the session's client must wait between calls
to `GetLatestConfiguration`.

In response, AWS AppConfig provides an `InitialConfigurationToken` to be given to
the session's client and used the first time it calls `GetLatestConfiguration`
for that session.

###### Important

This token should only be used once in your first call to
`GetLatestConfiguration`. You _must_ use the new token
in the `GetLatestConfiguration` response
( `NextPollConfigurationToken`) in each subsequent call to
`GetLatestConfiguration`.

When calling `GetLatestConfiguration`, your client code sends the most recent
`ConfigurationToken` value it has and receives in response:

- `NextPollConfigurationToken`: the `ConfigurationToken` value
to use on the next call to `GetLatestConfiguration`.

- `NextPollIntervalInSeconds`: the duration the client should wait before
making its next call to `GetLatestConfiguration`. This duration may vary
over the course of the session, so it should be used instead of the value sent on the
`StartConfigurationSession` call.

- The configuration: the latest data intended for the session. This may be empty if
the client already has the latest version of the configuration.

###### Important

The `InitialConfigurationToken` and
`NextPollConfigurationToken` should only be used once. To support long
poll use cases, the tokens are valid for up to 24 hours. If a
`GetLatestConfiguration` call uses an expired token, the system returns
`BadRequestException`.

For more information and to view example AWS CLI commands that show how to retrieve a
configuration using the AWS AppConfig Data
`StartConfigurationSession` and
`GetLatestConfiguration` API actions, see [Retrieving feature flags and\
configuration data in AWS AppConfig](../../../../services/appconfig/latest/userguide/retrieving-feature-flags.md) in the
_AWS AppConfig User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Actions

All content copied from https://docs.aws.amazon.com/.
