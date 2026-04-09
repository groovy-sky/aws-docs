AWS Audit Manager will no longer be open to new customers starting
April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers
can continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](audit-manager-availability-change.md).

# Tagging AWS Audit Manager resources

A _tag_ is a metadata label that you assign or that AWS
assigns to an AWS resource. Each tag consists of a _key_ and
a _value_. For tags that you assign, you define the key and
value. For example, you might define the key as `stage` and the value for one
resource as `test`.

Tags help you do the following:

- Easily locate your Audit Manager resources. You can use tags as search criteria when browsing the
framework library and the control library.

- Associate your resource with a compliance type. You can tag multiple resources with a
compliance-specific tag to associate those resources with a specific framework.

- Identify and organize your AWS resources. Many AWS services support tagging, so you
can assign the same tag to resources from different services to indicate that the resources
are related.

- Track your AWS costs. You activate these tags on the AWS Billing and Cost Management dashboard. AWS uses the
tags to categorize your costs and deliver a monthly cost allocation report to you. For more
information, see [Use cost allocation\
tags](../../../awsaccountbilling/latest/aboutv2/cost-alloc-tags.md) in the _AWS Billing and Cost Management User Guide_.

The following sections provide more information about tags for AWS Audit Manager.

###### Contents

- [Supported resources in Audit Manager](tagging.md#supported-resources)

- [Tag restrictions](tagging.md#tag-restrictions)

- [Additional resources](tagging.md#managing-tags)

## Supported resources in Audit Manager

The following Audit Manager resources support tagging:

- Assessments

- Controls

- Frameworks

## Tag restrictions

The following basic restrictions apply to tags on Audit Manager resources:

- Maximum number of tags that you can assign to a resource — 50

- Maximum key length — 128 Unicode characters

- Maximum value length — 256 Unicode characters

- Valid characters for key and value — a-z, A-Z, 0-9, space, and the following
characters: \_ . : / = + - and @

- Keys and values are case sensitive

- Don't use `aws:` as a prefix for keys; it's reserved for AWS use

## Additional resources

You can set tags as properties when you create an assessment, framework, or control. You
can add, edit, and delete tags through the Audit Manager console, the AWS Command Line Interface (AWS CLI), and the Audit Manager
API. For more information, see the following links.

- For tagging assessments:

- [Creating an assessment in AWS Audit Manager](create-assessments.md) and [Editing an assessment in AWS Audit Manager](edit-assessment.md) in the
_Assessments_ section of this guide

- [Tags tab](review-assessments.md#review-assessment-tags) in
the _Review an assessment_ page of this guide

- [CreateAssessment](../../../../reference/audit-manager/latest/apireference/api-createassessment.md) and [UpdateAssessment](../../../../reference/audit-manager/latest/apireference/api-updateassessment.md) in the _AWS Audit Manager API_
_Reference_

- [TagResource](../../../../reference/audit-manager/latest/apireference/api-tagresource.md) and [UntagResource](../../../../reference/audit-manager/latest/apireference/api-untagresource.md) in the _AWS Audit Manager API_
_Reference_

- For tagging frameworks:

- [Creating a custom framework in AWS Audit Manager](custom-frameworks.md) and [Editing a custom framework in AWS Audit Manager](edit-custom-frameworks.md) in the
_Framework library_ section of this guide

- The [Tags tab](review-frameworks.md#framework-tags-tab) on the
_View framework details_ page of this guide

- [CreateAssessmentFramework](../../../../reference/audit-manager/latest/apireference/api-createassessmentframework.md) and [UpdateAssessmentFramework](../../../../reference/audit-manager/latest/apireference/api-updateassessmentframework.md) in the _AWS Audit Manager API_
_Reference_

- [TagResource](../../../../reference/audit-manager/latest/apireference/api-tagresource.md) and [UntagResource](../../../../reference/audit-manager/latest/apireference/api-untagresource.md) in the _AWS Audit Manager API_
_Reference_

- For tagging controls:

- [Creating a custom control in AWS Audit Manager](create-controls.md) and [Editing a custom control in AWS Audit Manager](edit-controls.md) in the _Control_
_library_ section of this guide

- The [Tags](control-library-review-custom-controls.md#custom-control-tags-tab) section on the _Reviewing a_
_custom control_ page of this guide

- The [Tags](control-library-review-standard-controls.md#standard-control-tags-tab) section on the _Reviewing a_
_standard control_ page of this guide

- [CreateControl](../../../../reference/audit-manager/latest/apireference/api-createcontrol.md) and [UpdateControl](../../../../reference/audit-manager/latest/apireference/api-updatecontrol.md) in the _AWS Audit Manager API_
_Reference_

- [TagResource](../../../../reference/audit-manager/latest/apireference/api-tagresource.md) and [UntagResource](../../../../reference/audit-manager/latest/apireference/api-untagresource.md) in the _AWS Audit Manager API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Troubleshooting permissions and access

Quotas

All content copied from https://docs.aws.amazon.com/.
