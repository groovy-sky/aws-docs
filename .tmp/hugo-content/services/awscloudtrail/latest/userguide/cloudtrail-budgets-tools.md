# Using AWS Budgets to manage costs

AWS Budgets a feature of AWS Billing and Cost Management, allows you to set custom budgets that alert you when
your costs or usage exceed (or are forecasted to exceed) your budgeted amount.

Creating a budget for CloudTrail by using AWS Budgets is a
recommended best practice, and can help you track your CloudTrail spending. Cost-based budgets
help promote awareness of how much you might be billed for your CloudTrail use. [Budget alerts](../../../cost-management/latest/userguide/budgets-best-practices.md#budgets-best-practices-alerts) notify you when your bill reaches a
threshold that you define. When you receive a budget alert, you can make changes before
the end of the billing cycle to manage your costs.

###### Note

Though you can apply tags to CloudTrail trails, AWS Billing cannot currently use tags
applied to trails for cost allocation. Cost Explorer can show costs for CloudTrail Lake event data
stores and for the CloudTrail service as a whole.

To get started with AWS Budgets, open [AWS Billing and Cost Management](https://console.aws.amazon.com/billing), and then choose
**Budgets** in the left navigation bar. We recommend configuring
budget alerts as you create a budget to track CloudTrail spending. For more information about
how to use AWS Budgets, see [Managing your costs with AWS Budgets](../../../cost-management/latest/userguide/budgets-managing-costs.md) and [Best practices for AWS Budgets](../../../cost-management/latest/userguide/budgets-best-practices.md).

## Creating user-defined cost allocation tags for CloudTrail Lake event data stores

You can create [user-defined cost\
allocation tags](../../../awsaccountbilling/latest/aboutv2/custom-tags.md) to track the query and ingestion costs for your CloudTrail
Lake event data stores. A _user-defined cost allocation tag_ is a
key-value pair that you can associate with an event data store. After you activate
cost allocation tags, AWS uses the tags to organize your resource costs on your
cost allocation report.

- To create tags in the console, see step 9 of the [To create an event data store for CloudTrail events](query-event-data-store-cloudtrail.md#query-event-data-store-cloudtrail-procedure)
procedure.

- To create tags using the CloudTrail API, see [CreateEventDataStore](../../../../reference/awscloudtrail/latest/apireference/api-createeventdatastore.md) and [AddTags](../../../../reference/awscloudtrail/latest/apireference/api-addtags.md)
in the _AWS CloudTrail API Reference_.

- To create tags using the AWS CLI, see [create-event-data-store](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/cloudtrail/create-event-data-store.html) and [add-tags](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/cloudtrail/add-tags.html) in the _AWS CLI Command Reference_.

For more information about activating tags, see [Activating\
user-defined cost allocation tags](../../../awsaccountbilling/latest/aboutv2/activating-tags.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Viewing CloudTrail cost and usage

Managing CloudTrail trail costs

All content copied from https://docs.aws.amazon.com/.
