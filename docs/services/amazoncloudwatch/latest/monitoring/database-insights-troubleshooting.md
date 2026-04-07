# Troubleshooting for CloudWatch Database Insights

Use the following information to troubleshoot issues for CloudWatch Database Insights.

## Applying tags to Amazon RDS resources

To apply tags to your databases, use the Amazon RDS API, AWS CLI, or Amazon RDS console. For more information, see the following topics.

- [AddTagsToResource](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_AddTagsToResource.html) in the _Amazon RDS API Reference_

- [add-tags-to-resource](https://docs.aws.amazon.com/cli/latest/reference/rds/add-tags-to-resource.html) in the _Amazon RDS Command Line Reference_

- [Tagging Amazon Aurora and Amazon RDS resources](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Tagging.html) in the _Amazon Aurora User Guide_

## Maximum DB instances for fleets

You can't monitor more than 500 DB instances in a database fleet. You can use filters to create a fleet health view with less than 500 DB instances.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Monitoring Aurora Limitless databases with Database Insights

Use Contributor Insights to analyze high-cardinality data
