# Removing a source identifier from an Amazon RDS event notification subscription

You can remove a source identifier (the Amazon RDS source generating the event) from a subscription if you no longer
want to be notified of events for that source.

You can easily add or remove source identifiers using the Amazon RDS console by selecting or deselecting
them when modifying a subscription. For more information, see [Modifying an Amazon RDS event notification subscription](user-events-modifying.md).

To remove a source identifier from an Amazon RDS event notification subscription, use the AWS CLI [`remove-source-identifier-from-subscription`](https://docs.aws.amazon.com/cli/latest/reference/rds/remove-source-identifier-from-subscription.html) command. Include the following
required parameters:

- `--subscription-name`

- `--source-identifier`

###### Example

The following example removes the source identifier `mysqldb` from the
`myrdseventsubscription` subscription.

For Linux, macOS, or Unix:

```nohighlight

aws rds remove-source-identifier-from-subscription \
    --subscription-name myrdseventsubscription \
    --source-identifier mysqldb
```

For Windows:

```nohighlight

aws rds remove-source-identifier-from-subscription ^
    --subscription-name myrdseventsubscription ^
    --source-identifier mysqldb
```

To remove a source identifier from an Amazon RDS event notification subscription, use the Amazon RDS API [`RemoveSourceIdentifierFromSubscription`](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_RemoveSourceIdentifierFromSubscription.html) command. Include the following
required parameters:

- `SubscriptionName`

- `SourceIdentifier`

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Adding a source identifier to an Amazon RDS event notification subscription

Listing the Amazon RDS event notification categories
