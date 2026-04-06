# Deploying feature flags and configuration data in AWS AppConfig

After you [create\
required artifacts](https://docs.aws.amazon.com/appconfig/latest/userguide/creating-feature-flags-and-configuration-data.html) for working with feature flags and freeform configuration data, you
can create a new deployment. When you create a new deployment, you specify the following
information:

- An application ID

- A configuration profile ID

- A configuration version

- An environment ID where you want to deploy the configuration data

- A deployment strategy ID that defines how fast you want the changes to take
effect

- An AWS Key Management Service (AWS KMS) key ID to encrypt the data using a customer managed key.

When you call the [StartDeployment](https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_StartDeployment.html) API
action, AWS AppConfig performs the following tasks:

1. Retrieves the configuration data from the underlying data store by using the location
    URI in the configuration profile.

2. Verifies the configuration data is syntactically and semantically correct by using the
    validators you specified when you created your configuration profile.

3. Caches a copy of the data so it is ready to be retrieved by your application. This
    cached copy is called the _deployed data_.

You can mitigate situations where deploying configuration data causes errors in your
application by using a combination of AWS AppConfig deployment strategies and automatic rollbacks based
on Amazon CloudWatch alarms. A deployment strategy enables you to slowly release changes to all targets or specific segments over minutes or hours—either session-based or along your own target dimension by leveraging entity-based deployments. After you configure CloudWatch, if one or more alarms go into the
alarm state during a deployment, AWS AppConfig automatically rolls back your configuration data to the
previous version. For more information about deployment strategies, see [Working with deployment strategies](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-deployment-strategy.html). For more information about automatic
rollbacks, see [Monitoring deployments for automatic rollback](https://docs.aws.amazon.com/appconfig/latest/userguide/monitoring-deployments.html).

###### Note

AWS AppConfig Agent (version 2.0.136060 or later) supports deploying feature flag or free-form configuration data to specific segments or individual users during a gradual rollout. Entity-based gradual deployments ensure that once a user or segment receives a configuration version, they continue to receive that same version throughout the deployment period, regardless of which compute resource serves their requests. For more information, see [Using AWS AppConfig Agent for user- or entity-based gradual deployments](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-agent-how-to-use.html#appconfig-entity-based-gradual-deployments).

###### Topics

- [Working with deployment strategies](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-deployment-strategy.html)

- [Deploying a configuration](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-deploying.html)

- [Deploying AWS AppConfig configurations using CodePipeline](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-integration-codepipeline.html)

- [Reverting a configuration](https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-deploying-reverting.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Creating a configuration profile for non-native data sources

Working with deployment strategies
