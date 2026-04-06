# Query Application Load Balancer logs

An Application Load Balancer is a load balancing option for Elastic Load Balancing that enables traffic distribution in a
microservices deployment using containers. Querying Application Load Balancer logs allows you to see the source
of traffic, latency, and bytes transferred to and from Elastic Load Balancing instances and backend
applications. For more information, see [Access logs\
for your Application Load Balancer](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-access-logs.html) and [Connection logs for your Application Load Balancer](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-connection-logs.html) in the
_User Guide for Application Load Balancers_.

## Prerequisites

- Enable [access logging](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-access-logs.html) or [connection logging](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-connection-logs.html) so that Application Load Balancer logs can be saved to your Amazon S3
bucket.

- A database to hold the table that you will create for Athena. To create a
database, you can use the Athena or AWS Glue console. For more information, see
[Create databases in Athena](https://docs.aws.amazon.com/athena/latest/ug/creating-databases.html) in
this guide or [Working with databases on the\
AWS glue console](https://docs.aws.amazon.com/glue/latest/dg/console-databases.html) in the _AWS Glue Developer Guide_.

###### Topics

- [Create the table for ALB access logs](https://docs.aws.amazon.com/athena/latest/ug/create-alb-access-logs-table.html)

- [Create the table for ALB access logs in Athena using partition projection](https://docs.aws.amazon.com/athena/latest/ug/create-alb-access-logs-table-partition-projection.html)

- [Example queries for ALB access logs](https://docs.aws.amazon.com/athena/latest/ug/query-alb-access-logs-examples.html)

- [Create the table for ALB connection logs](https://docs.aws.amazon.com/athena/latest/ug/create-alb-connection-logs-table.html)

- [Create the table for ALB connection logs in Athena using partition projection](https://docs.aws.amazon.com/athena/latest/ug/create-alb-connection-logs-table-partition-projection.html)

- [Example queries for ALB connection logs](https://docs.aws.amazon.com/athena/latest/ug/query-alb-connection-logs-examples.html)

- [Additional resources](https://docs.aws.amazon.com/athena/latest/ug/application-load-balancer-logs-additional-resources.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Query AWS service logs

ALB access logs
