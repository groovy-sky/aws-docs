# Getting access to your customer's Amazon Q index

There are two non-mutually exclusive ways to get access to a customer's Amazon Q
index:

## Customer grants ISVs (data accessors) access to their Amazon Q index

Amazon Q Business customers add the ISV as a _data accessor_ to their
Amazon Q Business application environment and underlying Amazon Q index. This includes Amazon Q
customers selecting which data sources and end users can retrieve data from and
grant ISVs _cross-account_ access to their Amazon Q index based on
those permissions.

###### Note

Only ISVs who successfully complete registration with Amazon Q Business can be listed
as data accessors.

Only Amazon Q Business application environments that are created using
AWS Identity and Access Management (IAM) identity center (IDC) are supported.

This process gives Amazon Q customers more flexibility, oversight, and control of
their Amazon Q index. Amazon Q customers can create their Q index once, and then grant
multiple ISVs access to the same minimizing replication of enterprise data and
permissions across multiple ISVs. For more information, see [Data\
accessors](data-accessors.md) and [Accessing a\
customer's Amazon Q index as a data accessor using cross-account\
access](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/isv-accessing-cross-account.html).

## The ISV sets up the Amazon Q index on behalf of the customer within their own AWS account

ISVs can manage the entire process themselves, creating an Amazon Q Business
application environment and the underlying Amazon Q index on behalf of each customer in their
AWS account. Customers must provide ISV credentials to connect to their data
sources.

At the end of both workflows, ISVs can retrieve the data that they have access to
using the Amazon Q index.

Comparison of onboarding optionsScenarioOnboarding with Amazon Q BusinessOnboarding with the ISVWhere is the Amazon Q Business application environment created?Amazon Q Business customer's accountISV's accountWho sets up Amazon Q Business application environment and underlying Amazon Q
index?Amazon Q Business customerISV on behalf of the customer

We recommend that ISVs use both onboarding workflows to best support all
customers.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ISV information

Creating an Amazon Q index
