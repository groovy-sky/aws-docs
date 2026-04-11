# Quotas and restrictions

You can use the following tables to find the current quotas, restrictions, and naming
constraints within the AWS Billing and Cost Management console.

###### Notes

- To learn more about quotas and restrictions for AWS Cost Management, see [Quotas and\
restrictions](../../../cost-management/latest/userguide/management-limits.md) in the _AWS Cost Management User Guide_.

- For more information about other AWS service quotas, see [AWS\
service quotas](../../../../general/latest/gr/aws-service-limits.md) in the _AWS General Reference_.

###### Topics

- [Cost categories](#limits-categories)

- [Purchase orders](#limits-po)

- [Advance Pay](#limits-ap)

- [Cost allocation tags](#limits-cat)

- [AWS Price List](#price-list-api-quotas)

- [Bulk policy migrator](#limits-bulk)

- [Payment methods](#limits-payments)

- [AWS invoice configuration](#limits-invoicing)

## Cost categories

See the following quotas and restrictions for cost categories.

DescriptionQuotas and restrictionsThe total number of cost categories for a management account.`50`The number of cost category rules for a cost category
(API).`500`The number of cost category rules for a cost category
(UI).`100`

Cost category names.

- Names must be unique

- Case sensitive

Cost category value names.Names don't have to be unique

The type and number of characters allowed in a cost category name
and value name.

- Numbers: `0-9`

- Unicode letters

- `Space`, if it's not used at the beginning or
end of the name

- The following symbols: underscore ( `_`) or en
dash (-)

The number of split charge rules for a cost category.`10`

## Purchase orders

See the following quotas and restrictions for purchase orders.

DescriptionQuotas and restrictionsThe type of characters that you can use in a purchase order
ID.

- `A-Z` and `a-z`

- `Space`

- The following symbols: `_.:/=+-%@`

The number of characters allowed in a purchase order ID.`100`The number of contacts allowed for a purchase order.`20`The number of tags allowed for a purchase order.`50`The number of line items allowed for a purchase order.`100`

## Advance Pay

See the following quotas and restrictions for Advance Pay.

DescriptionQuotas and restrictionsUser entity AWS Inc. or AWS EuropeCurrencyUSDFund usage after funds are added to your Advance Pay.

- Funds can only be used to pay for eligible AWS charges.
Non-eligible charges (for example, AWS Marketplace invoices) are
charged using the default payment method at the time of
Advance Pay registration.

- Advance Pay funds added in AWS Europe can only be used to pay AWS Europe invoices.

- Funds can't be withdrawn, refunded, or transferred.

- Funds can't be converted to other currencies.

If there are unused funds in your Advance Pay.

- You can't change your seller on record.

- You can't change your preferred currency.

- You can't change your default payment method.

## Cost allocation tags

You can adjust the maximum number of active cost allocation tag
keys from Service Quotas. For more information, see [Requesting a quota increase](../../../servicequotas/latest/userguide/request-quota-increase.md) in the _Service Quotas User_
_Guide_.

###### Note

Tags that are automatically activated don’t count towards your cost allocation tag
quota, such as the `awsApplication` tag.

See the following quotas and restrictions for cost allocation tags.

DescriptionQuotas and restrictionsThe maximum number of active cost allocation tag keys for each payer
account.500The number of cost allocation tags that can be activated or deactivated for
one request, by using the API or the console.20

## AWS Price List

Some Price List Query API and Price List Bulk API operations are throttled by using a token bucket scheme to
maintain service availability. These quotas are per AWS account on a per Region basis.
The following table shows the quotas for each API operation.

Price List Query APIAPI operationToken bucket sizeRefill rate per second`DescribeServices`105`GetAttributeValues`105`GetProducts`105

Price List Bulk APIAPI operationToken bucket sizeRefill rate per second`DescribeServices`105`GetPriceListFileUrl`105`ListPriceLists`105

## Bulk policy migrator

See the following quotas and restrictions for bulk policy migrator.

DescriptionQuotas and restrictionsThe maximum number of affected accounts in an organization that you can migrate.200The maximum number of affected policies in an organization that you can migrate.1,000

## Payment methods

See the following quotas and restrictions for payments.

DescriptionQuotas and restrictions

Tagging payment instruments.

This feature supports the following payment methods:

- Credit cards

- Bank accounts (ACH)

This feature doesn't support the following payment methods:

- Advance pay

- Net Banking

- China bank redirect

- PIX

- United Payments Interface (UPI)

- Pay by invoice

## AWS invoice configuration

See the following quotas and restrictions for Invoice configuration.

DescriptionQuotas and restrictions

The number of invoice units for a payer account.

`500`

The type of characters allowed in an invoice unit name.

- The name must be between 1-50 characters.

- Letters: `A-Z` and `a-z`

- Numbers: `0-9`

- `Space`

- The following symbols: hyphen ( `-`), underscore ( `_`)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS PrivateLink

Document history

All content copied from https://docs.aws.amazon.com/.
