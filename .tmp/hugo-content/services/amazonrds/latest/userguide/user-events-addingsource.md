# Adding a source identifier to an Amazon RDS event notification subscription

You can add a source identifier (the Amazon RDS source generating the event) to an existing subscription.

You can easily add or remove source identifiers using the Amazon RDS console by selecting or deselecting
them when modifying a subscription. For more information, see [Modifying an Amazon RDS event notification subscription](user-events-modifying.md).

To add a source identifier to an Amazon RDS event notification subscription, use the AWS CLI [`add-source-identifier-to-subscription`](../../../../general/index.md) command. Include
the following required parameters:

- `--subscription-name`

- `--source-identifier`

###### Example

The following example adds the source identifier `mysqldb` to the
`myrdseventsubscription` subscription.

For Linux, macOS, or Unix:

```nohighlight

aws rds add-source-identifier-to-subscription \
    --subscription-name myrdseventsubscription \
    --source-identifier mysqldb
```

For Windows:

```nohighlight

aws rds add-source-identifier-to-subscription ^
    --subscription-name myrdseventsubscription ^
    --source-identifier mysqldb
```

To add a source identifier to an Amazon RDS event notification subscription, call the Amazon RDS API [`AddSourceIdentifierToSubscription`](../../../../reference/amazonrds/latest/apireference/api-addsourceidentifiertosubscription.md). Include the following required
parameters:

- `SubscriptionName`

- `SourceIdentifier`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Modifying an Amazon RDS event notification subscription

Removing a source identifier from an Amazon RDS event notification subscription

All content copied from https://docs.aws.amazon.com/.
