# Creating flexible CloudWatch dashboards with dashboard variables

Use _dashboard variables_ to create flexible dashboards that can
quickly display different content in multiple widgets, depending on the value of an
input field within the dashboard. For example, you can create a dashboard that can
quickly switch between different Lambda functions or Amazon EC2 instance IDs, or one that can
switch to different AWS Regions.

After you create a dashboard that uses a variable, you can copy the same variable
pattern to other existing dashboards.

Using dashboard variables improves the operational workflow for people who use your
dashboards. It can also reduce your costs because you're using dashboard variables in
one dashboard instead of creating multiple similar dashboards.

###### Note

If you share a dashboard that contains dashboard variables, the people that you share it with
won't be able to change between the variable values.

## Types of dashboard variables

The dashboard variable can be a _property variable_ or a
_pattern variable_.

- _Property variables_ change all instances of a property in all widgets in
the dashboard. This property can be any JSON property in the JSON source of a
dashboard, such as `region`. Or it can be a dimension name for a
metric, such as `InstanceID` or `FunctionName`.

For a tutorial that uses a property variable,
see [Tutorial: Creating a CloudWatch Lambda dashboard with function name as the variable](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_dashboard_variables_property.html).

For more information about the JSON source of dashboards,
see [Dashboard Body Structure and Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Dashboard-Body-Structure.html).
In the CloudWatch console, you can see the JSON source for any custom dashboard by choosing **Actions**,
**View/edit source**.

- _Pattern variables_ use a regular expression pattern to change all of a
JSON property or only a certain part of it.

For a tutorial that uses a pattern variable,
see [Tutorial: Creating a dashboard that uses a regular expression pattern to switch between AWS Regions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_dashboard_variables_pattern.html).

Property variables apply to most use cases and are less complex to set up.

###### Topics

- [Copying a variable to another CloudWatch dashboard](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_dashboard_variables_copy.html)

- [Tutorial: Creating a dashboard that uses a regular expression pattern to switch between AWS Regions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_dashboard_variables_pattern.html)

- [Tutorial: Creating a CloudWatch Lambda dashboard with function name as the variable](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_dashboard_variables_property.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Adding an alarm from a different account to a cross-account dashboard

Copying a variable to another dashboard
