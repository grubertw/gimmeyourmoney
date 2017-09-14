FORMAT: 1A
HOST: https://gimmeyourmoney.local.orderservice.io/

# Gimmeyourmoney Order Service

This service represents the main entry point for processing order requests of various types.  The most common type of order is initiated by a lender, for purposes of obtaining a report.  The contents of the report is both controlled and described by a product.  Using the product definition, along with a series of other information specified in the order, the system generates the report, which usually consists of consumer account transactions (i.e banking transactions), formatted accoring to an output template (i.e. PDF, html, or json document).

# Group Orders

This section provides details on order submission (i.e. **POST**) and order status (i.e. **GET**)

## Order Submission [/]

### Submit Order [POST]

Orders are submitted in the form of a JSON document.  The structure of this document can be highly variable.  Required fields are as follows: 

+ lenderId - Value chosen by the lender to uniquely represent itself.  This avoids the need to share the internally generated lenderId, whilst also allowing the lender a more user-friendly way to track itself.
+ orderId -  Value chosen by lender to uniquely identify the order.
+ itmes - Selected products in the order (there must be at least one of these).
    + productSku - Product idendifier for the order item.

Optional fields are as follows:
+ userId - User responsible for the order (usually an employee of the lender).

+ Attributes (OrderRequest)

+ Request (application/json)

+ Response 202 (application/json)

    + Headers

        Location: /orders/ord-556

    + Attributes (OrderResponse)

+ Response 400 (text/plain)

        JSON document was malformed.

+ Response 401 (application/json)

    + Attributes
        + errorCode: `0x0001` (string, required)
        + errorDescription: Lender was not authorized (string, required)

+ Response 401 (application/json)

    + Attributes
        + errorCode: `0x0002` (string, required)
        + errorDescription: Representitive responsible for order was not authorized (string, required)

+ Response 403 (application/json)

    + Attributes
        + errorCode: `0x0101` (string, required)
        + errorDescription: Order is not refreshable (string, required)

+ Response 403 (application/json)

    + Attributes
        + errorCode: `0x0102` (string, required)
        + errorDescription: Order refresh duration has been exceeded. (string, required)

+ Response 403 (application/json)

    + Attributes
        + errorCode: `0x0103` (string, required)
        + errorDescription: Order is not upgradable. (string, required)

+ Response 403 (application/json)

    + Attributes
        + errorCode: `0x0104` (string, required)
        + errorDescription: Order cannot be upgraded becuase consumer refereces cannot be found. (string, required)

## Order [/order/{id}] 

### Query a single Order [GET]
Query a single order by it's agreed on reference ID.

+ Parameters
    + id: `ord-556` (string) - Agreed upon reference ID for the order.

+ Response 200 (application/json)
    + Attributes (OrderQuery)

## Lender Orders [/orders/{lenderId}/{?status}]

### Get Orders for a Lender [GET]

+ Parameters
    + lenderId: MorganStandly (string) - Agreed upon ID for the lender.
    + status (string, optional)

        Query an Order by it's status: ACTIVE, WAITING_FOR_CAM, WAITING_FOR_CAM_PARTIAL, AVAILABLE, PARTIALLY_AVAILABLE, EXPIRED, PARTIALLY_EXPIRED

        + Default: `ACTIVE`

+ Response 200 (application/json)
    + Attributes (array[OrderQuery])



# Data Structures

## OrderRequest (object)
+ lenderId: MorgnStandly (string, required)
+ orderId: `ord-556` (string, required)
+ type: INITIAL (string, required) - Type of order.  See POST docs for enums.
+ items (array[OrderItem], required) - Selected products in the order.
+ userId: hendrickj (string) - User responsible for the order (usually an employee of the lender).

## OrderItem (object)
+ productSku: `report-951` (string, required)
+ consumerItems (array[ConsumerItem], required)
+ parameterCardId: `card-5` (string)
+ customRuleParams (array[RuleParameter])
+ extras: `statement-w2:PDF`, `PRK-5:HTML` (array[string])

## ConsumerItem (object)
+ consumerId: `c-269501692` (string)
+ consumerIdent (ConsumerIdent)
+ requestedAccounts (array[RequestedAccount])
+ consumerIsEngaged: yes (string)
+ lenderFlow: CLOSED (string)

## ConsumerIdent (object)
+ firstName: John (string)
+ middleName: Wilks (string)
+ lastName: Doe (string)
+ addresses (array[Address])
+ ssn: `123-45-6789` (string)
+ emails: `johnwdoe@hotmail.com`, `doej@yahoo.com` (array[string])
+ phones: `465-397-2266`, `(818) 332-2184` (array[string])

## Address (object)
+ country: USA (string)
+ state: CA (string)
+ city: Los Angeles (string)
+ street: `31 Green` (string)
+ zipCode: 92650 (number)

## RequestedAccount (object)
+ accountNum: 9378298758 (number, required)
+ required: yes (string)

## RuleParameter (object)
+ rule: Sum transactions in catigory x for the last y weeks (string)
+ arguments: RECREATION, 3 (array[string])

## OrderResponse (object)
+ orderId: `ord-556` (string, required)
+ created: 415203908 (number) - Time stamp
+ status: ACTIVE (string, required)
+ refreshable: yes (string, required)
+ refreshDuration: `60 days` (string, required)
+ consumersFoundLocally: no (string, required)

## OrderQuery (object)
+ orderId: `ord-556` (string)
+ status: ACTIVE (string)
+ type: INITAL (string)
+ fufilledReports: `report-951`, `report-952` (array[string])
