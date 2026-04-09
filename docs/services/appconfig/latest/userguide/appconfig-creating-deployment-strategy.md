# Working with deployment strategies

A deployment strategy enables you to slowly release changes to all targets or specific segments over minutes or hours—either session based or along your own target dimension by leveraging entity-based deployments.

###### Note

AWS AppConfig Agent (version 2.0.136060 or later) supports deploying feature flag or free-form configuration data to specific segments or individual users during a gradual rollout. Entity-based gradual deployments ensure that once a user or segment receives a configuration version, they continue to receive that same version throughout the deployment period, regardless of which compute resource serves their requests. For more information, see [Using AWS AppConfig Agent for user- or entity-based gradual deployments](appconfig-agent-how-to-use.md#appconfig-entity-based-gradual-deployments).

An AWS AppConfig deployment strategy defines the following important aspects
of a configuration deployment.

SettingDescription

Deployment type

Deployment type defines how the configuration deploys or _rolls_
_out_. AWS AppConfig supports **Linear** and
**Exponential** deployment types.

- **Linear**: For this type, AWS AppConfig processes the
deployment by increments of the growth factor evenly distributed over the
deployment. Here's an example timeline for a 10 hour deployment that uses 20%
linear growth:

Elapsed timeDeployment progress

0 hour

0%

2 hour

20%

4 hour

40%

6 hour

60%

8 hour

80%

10 hour

100%

- **Exponential**: For this type, AWS AppConfig
processes the deployment exponentially using the following formula:
`G*(2^N)`. In this formula, `G` is the step percentage
specified by the user and `N` is the number of steps until the
configuration is deployed to all targets. For example, if you specify a growth
factor of 2, then the system rolls out the configuration as follows:

```

2*(2^0)
2*(2^1)
2*(2^2)
```

Expressed numerically, the deployment rolls out as follows: 2% of the
targets, 4% of the targets, 8% of the targets, and continues until the
configuration has been deployed to all targets.

Step percentage (growth factor)

This setting specifies the percentage of callers to target during each step of
the deployment.

###### Note

In the SDK and the [AWS AppConfig API\
Reference](../../../../reference/appconfig/2019-10-09/apireference/api-createdeploymentstrategy.md), `step percentage` is called `growth
                    factor`.

Deployment time

This setting specifies an amount of time during which AWS AppConfig deploys to hosts.
This is not a timeout value. It is a window of time during which the deployment is
processed in intervals.

Bake time

This setting specifies the amount of time AWS AppConfig monitors for Amazon CloudWatch alarms
after the configuration has been deployed to 100% of its targets, before considering
the deployment to be complete. If an alarm is triggered during this time, AWS AppConfig
rolls back the deployment. You must configure permissions for AWS AppConfig to roll back
based on CloudWatch alarms. For more information, see [Configure permissions for automatic rollback](setting-up-appconfig.md#getting-started-with-appconfig-cloudwatch-alarms-permissions).

You can choose a predefined strategy included with AWS AppConfig or create your own.

###### Note

AWS AppConfig Agent (version 2.0.136060 or later) supports deploying feature flag or free-form configuration data to specific segments or individual users during a gradual rollout. Entity-based gradual deployments ensure that once a user or segment receives a configuration version, they continue to receive that same version throughout the deployment period, regardless of which compute resource serves their requests. For more information, see [Using AWS AppConfig Agent for user- or entity-based gradual deployments](appconfig-agent-how-to-use.md#appconfig-entity-based-gradual-deployments).

###### Topics

- [Using predefined deployment strategies](appconfig-creating-deployment-strategy-predefined.md)

- [Create a deployment strategy](appconfig-creating-deployment-strategy-create.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Deploying

Using predefined deployment strategies

All content copied from https://docs.aws.amazon.com/.
