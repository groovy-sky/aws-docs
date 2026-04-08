# Listing the Amazon RDS event notification categories

All events for a resource type are grouped into categories. To view the list of categories available, use the
following procedures.

When you create or modify an event notification subscription, the event categories are displayed in the
Amazon RDS console. For more information, see [Modifying an Amazon RDS event notification subscription](user-events-modifying.md).

![List DB event notification categories](https://docs.aws.amazon.com/images/AmazonRDS/latest/UserGuide/images/EventNotification-Categories.png)

To list the Amazon RDS event notification categories, use the AWS CLI [`describe-event-categories`](../../../cli/latest/reference/rds/describe-event-categories.md)
command. This command has no required parameters.

###### Example

```nohighlight

aws rds describe-event-categories
```

To list the Amazon RDS event notification categories, use the Amazon RDS API [`DescribeEventCategories`](../../../../reference/amazonrds/latest/apireference/api-describeeventcategories.md)
command. This command has no required parameters.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Removing a source identifier from an Amazon RDS event notification subscription

Deleting an Amazon RDS event notification subscription

All content copied from https://docs.aws.amazon.com/.
