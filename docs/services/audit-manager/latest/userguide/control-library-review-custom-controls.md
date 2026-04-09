AWS Audit Manager will no longer be open to new customers starting
April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers
can continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](audit-manager-availability-change.md).

# Reviewing a custom control

You can review the details of a custom control by using the Audit Manager console, the Audit Manager API,
or the AWS Command Line Interface (AWS CLI).

## Prerequisites

Make sure your IAM identity has appropriate permissions to view controls in
AWS Audit Manager. Two suggested policies that grant these permissions are [AWSAuditManagerAdministratorAccess](../../../aws-managed-policy/latest/reference/awsauditmanageradministratoraccess.md) and [Allow users management access to AWS Audit Manager](security-iam-id-based-policy-examples.md#management-access).

## Procedure

You can review the details of a custom control by using the Audit Manager console, the Audit Manager
API, or the AWS Command Line Interface (AWS CLI).

Audit Manager console

###### To view custom control details on the Audit Manager console

1. Open the AWS Audit Manager console at [https://console.aws.amazon.com/auditmanager/home](https://console.aws.amazon.com/auditmanager/home).

2. In the navigation pane, choose **Control library**.

3. Choose **Custom** to see the custom controls that you
    created.

4. Choose any custom control name to view the details for that control.

5. Review the custom control details using the following information as
    reference.

**Overview section**

This section describes the custom control and lists the [data\
source types](concepts.md#control-data-source) that it uses to collect evidence. It also provides
information about when the control was created and last updated.

**Evidence sources tab**

This tab shows where the custom control collects evidence from. It
includes the following information:

NameDescription

**Common controls**

These are the common controls that collect evidence to support the custom control.

Common controls collect evidence using underlying data sources
that AWS manages for you. For every common control that’s listed,
Audit Manager collects the relevant evidence for all of the supporting core
controls. Choose a common control to see the related core
controls.

**Core controls**

These are the core controls that collect evidence to support the
custom control.

Core controls collect evidence by using a predefined group
of data sources that AWS manages for you. Choose a core
control to see the underlying data sources.

**Data sources**

These are the data sources that collect evidence to support the
custom control.

###### Note

These data sources aren't managed for you by AWS. You're
responsible for maintaining them.

- **Name** – The name of the data
source.

- **Type** – The type of data source
that the evidence comes from.

- If Audit Manager collects the evidence, the type can be _AWS Security Hub CSPM_, _AWS Config_, _AWS CloudTrail_, or _AWS API_
_calls_.

- If you upload your own evidence, the type is _Manual_. A description indicates
if the required manual evidence is a _File upload_ or a _Text_
_response_.

- **Mapping** – The specific keyword
that's used to collect evidence.

- If the type is _AWS Config_, the mapping is an AWS Config rule
(such as `SNS_ENCRYPTED_KMS`).

- If the type is _AWS Security Hub CSPM_, the mapping is a Security Hub CSPM control
(such as `EC2.1`).

- If the type is _AWS API_
_calls_, the mapping is an API call (such as
`kms_ListKeys`).

- If the type is _AWS CloudTrail_, the mapping is a CloudTrail event (such as
`CreateAccessKey`).

- **Frequency** – How often Audit Manager
collects evidence for an AWS API call data source.

**Details tab**

This tab includes the following information:

NameDescription

**Instructions**

The directions that describe how to test and remediate the
control.**Testing information**

The recommended testing procedures.

**Action plan**

The recommended actions to take if you need to remediate the
control.

**Tags**

The tags that are associated with the control.

**Key**

The tag key (for example, a compliance standard, regulation, or
category).

**Value**

The tag value.

AWS CLI

###### To view custom control details in the AWS CLI

1. Follow the steps to [find a\
    control](access-available-controls.md). Make sure to set the `--control-type` as
    `Custom`, and apply any optional filters as needed.

```

aws auditmanager list-controls --control-type Custom
```

2. In the response, identify the control that you want to review and take note of
    the control ID and Amazon Resource Name (ARN).

3. Run the [get-control](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/auditmanager/get-control.html) command and specify the `--control-id`. In the
    following example, replace the `placeholder text` with
    your own information.

```nohighlight

aws auditmanager get-control --control-id a1b2c3d4-5678-90ab-cdef-EXAMPLE11111
```

###### Tip

The control details are returned in JSON format. To help you understand this
data, see [get-control Output](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/auditmanager/get-control.html) in the _AWS CLI Command_
_Reference_.

4. To see the tags for a control, use the [list-tags-for-resource](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/auditmanager/list-tags-for-resource.html) command and specify the
    `--resource-arn`. In the following example, replace the
    `placeholder text` with your own information:

```nohighlight

aws auditmanager list-tags-for-resource --resource-arn arn:aws:auditmanager:us-east-1:111122223333:control/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111
```

Audit Manager API

###### To view custom control details using the API

1. Follow the steps to [find a\
    control](access-available-controls.md). Make sure to set the [controlType](../../../../reference/audit-manager/latest/apireference/api-listcontrols.md#auditmanager-ListControls-request-controlType) as `Custom`, and apply any optional filters as
    needed.

2. In the response, identify the control that you want to review and take note of
    the control ID and its Amazon Resource Name (ARN).

3. Use the [GetControl](../../../../reference/audit-manager/latest/apireference/api-getcontrol.md)
    operation and specify the [controlId](../../../../reference/audit-manager/latest/apireference/api-getcontrol.md#auditmanager-GetControl-request-controlId) that you noted in step 2.

###### Tip

The control details are returned in JSON format. To help you understand this
data, see [GetControl Response Elements](../../../../reference/audit-manager/latest/apireference/api-getcontrol.md#API_GetControl_ResponseElements) in the _AWS Audit Manager_
_API Reference_.

4. To see tags for the control, use the [ListTagsForResource](../../../../reference/audit-manager/latest/apireference/api-listtagsforresource.md) operation and specify the control [resourceArn](../../../../reference/audit-manager/latest/apireference/api-listtagsforresource.md#auditmanager-ListTagsForResource-request-resourceArn) that you noted in step 2.

For more information about these API operations, choose any of the links in this
procedure to read more in the _AWS Audit Manager API_
_Reference_. This includes information about how to use these operations
and parameters in one of the language-specific AWS SDKs.

## Next steps

You can add a custom control to any of your custom frameworks. For instructions, see
[Creating a custom framework in AWS Audit Manager](custom-frameworks.md).

You can also [edit a custom control](edit-controls.md),
[make an\
editable copy of a custom control](customize-control-from-existing.md), or [delete a custom\
control](delete-controls.md) that you no longer need.

## Additional resources

- [Reviewing a common control](control-library-review-common-controls.md)

- [Reviewing a core control](control-library-review-core-controls.md)

- [Reviewing a standard control](control-library-review-standard-controls.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Standard controls

Creating a custom control

All content copied from https://docs.aws.amazon.com/.
