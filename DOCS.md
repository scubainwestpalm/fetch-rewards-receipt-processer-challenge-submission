# Documentation

Documentation for packages and methods in source of written solution.

### Project Directory
```
./
    main.go
    go.mod
    go.sum
    .gitignore
    README.md
    DOCS.md
    ./models
        item.go
        receipt.go
    ./router
        receipts.go
        router.go
    ./utils
        points_helpers.go
```

### _*Packages*_

Project go packages

* [Models](#models)
    * [Process](#processMethod)
    * [parseTimeDataFromDateAndTimeStrings](#parsetimedatafromdateandtimestringsMethod)
    * [GetReceiptById](#getreceiptbyidMethod)
* [Router](#router)
    * [GET](#get-receiptsidpoints)
    * [POST](#post-receiptsprocess)
* [Utils](#utils)
    * [CountAlphanumericCharsInString](#countalphanumericcharsinstringMethod)
    * [Float64IsWholeNumber](#float64iswholenumberMethod)
    * [Float64IsMultipleOf](#float64ismultipleofMethod)
    * [IntIsMultipleOf](#intismultipleofMethod)

### [Models](/models/)

Contains structs and their methods for receipts and items within a receipt

* ##### [Receipts](/models/receipt.go)

    > Struct

    https://github.com/scubainwestpalm/fetch-rewards-receipt-processer-challenge-submission/blob/79d04f22707fd94d962b74be44a9d70438831107/models/receipt.go#L13-L20

    ###### Marshaled JSON:

    ```json
    {
        "retailer": "Target",
        "purchaseDate": "2022-01-02",
        "purchaseTime": "13:13",
        "total": "1.25",
        "items": [
            {"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
        ]
    }
    ```

    > Struct Methods

    <span id="processMethod"></span>
    * #### _**Process**_ <sup> _uuid, err_</sup>
    
        Determines points awarded to a receipt and returns a uuid after successfully aggregating points

        https://github.com/scubainwestpalm/fetch-rewards-receipt-processer-challenge-submission/blob/79d04f22707fd94d962b74be44a9d70438831107/models/receipt.go#L24-L72

        ###### _Example_

        ```golang
        ...
        uuid, err := Receipt.Process()
        if err != nil {
            panic(err)
        }
        fmt.Println(uuid)
        ...
        ```

        ###### **Outputs:**

        `a4566f0b-c7de-4b58-948e-7467993abd64`

    <span id="parsetimedatafromdateandtimestringsMethod"></span>
    * #### _**parseTimeDataFromDateAndTimeStrings**_ <sup> _purchaseDay, hours, mins, err_ </sup>

        Parses time data from purchase date and purchase time strings within receipt, and returns clock info as integers for point calculations

        https://github.com/scubainwestpalm/fetch-rewards-receipt-processer-challenge-submission/blob/79d04f22707fd94d962b74be44a9d70438831107/models/receipt.go#L74-L85

        ###### _Example_

        ```golang
        ...
        // E.G Purchase Date is 2020-01-02 and Purchase Time is 13:13
        purchaseDay, hours, mins, err := Receipt.parseTimeDataFromDateAndTimeStrings()
        if err != nil {
            panic(err)
        }
        fmt.Println(purchaseDay, hours, mins)
        ...
        ```

        ###### **Outputs:**

        `02 13 13`


    > Methods

    <span id="getreceiptbyidMethod"></span>
    * #### _**GetReceiptById**_ <sub> _id **string**_</sub> <sup> _Receipt, err_ </sup>

        Fetches receipt by uuid from map stored in application memory

        https://github.com/scubainwestpalm/fetch-rewards-receipt-processer-challenge-submission/blob/79d04f22707fd94d962b74be44a9d70438831107/models/receipt.go#L87-L99

         ###### _Example_

        ```golang
        ...
        Receipt, err := models.GetReceiptById("a4566f0b-c7de-4b58-948e-7467993abd64")
        if err != nil {
            panic(err)
        }
        fmt.Println(Receipt.Points)
        ...
        ```

        ###### **Outputs:**

        `31`

* ##### [Items](/models/item.go)

    > Struct

    https://github.com/scubainwestpalm/fetch-rewards-receipt-processer-challenge-submission/blob/79d04f22707fd94d962b74be44a9d70438831107/models/item.go#L3-L6

    ###### Marshaled JSON:

    ```json
    { 
        "shortDescription": "Pepsi - 12-oz", 
        "price": "1.25"
    } 
    ```

### [Router](/router/)

Adds routes for processing a receipt and getting receipt points by id

https://github.com/scubainwestpalm/fetch-rewards-receipt-processer-challenge-submission/blob/79d04f22707fd94d962b74be44a9d70438831107/router/router.go#L7-L10

###### _**POST**_ `/receipts/process`

https://github.com/scubainwestpalm/fetch-rewards-receipt-processer-challenge-submission/blob/79d04f22707fd94d962b74be44a9d70438831107/router/receipts.go#L11-L27

###### _**GET**_ `/receipts/:id/points`

https://github.com/scubainwestpalm/fetch-rewards-receipt-processer-challenge-submission/blob/79d04f22707fd94d962b74be44a9d70438831107/router/receipts.go#L29-L43

### [Utils](/utils/)

Contains methods to help with the process of calculating points, and to improve readability of code in receipt [process](#process--uuid-err) method.

> Methods

<span id="countalphanumericcharsinstringMethod"></span>
* #### _**CountAlphanumericCharsInString**_ <sub> _str **string**_</sub> <sup> _int_ </sup>

    Counts number of alphanumeric characters in string, and returns count as an integer.

    https://github.com/scubainwestpalm/fetch-rewards-receipt-processer-challenge-submission/blob/79d04f22707fd94d962b74be44a9d70438831107/utils/point_helpers.go#L9-L16

    ###### _Example_

    ```golang
    ...
    alphaNumericCount := utils.CountAlphanumericCharsInString("AbCd3!!!")
    fmt.Println(alphaNumericCount)
    ...
    ```

    ###### **Outputs:**

    `5`

<span id="float64iswholenumberMethod"></span>
* #### _**Float64IsWholeNumber**_ <sub> _number **float64**_</sub> <sup> _bool_ </sup>

    Checks if float64 is a whole number, returns true if float64 is a whole number.

    https://github.com/scubainwestpalm/fetch-rewards-receipt-processer-challenge-submission/blob/79d04f22707fd94d962b74be44a9d70438831107/utils/point_helpers.go#L18-L20

    ###### _Example_

    ```golang
    ...
    float64IsWholeNumber := utils.Float64IsWholeNumber(20.00)
    fmt.Println(float64IsWholeNumber)
    ...
    ```

    ###### **Outputs:**

    `true`

<span id="float64ismultipleofMethod"></span>
* #### _**Float64IsMultipleOf**_ <sub> _number **float64**, divisor **float64**_</sub> <sup> _bool_ </sup>

    Checks if float64 is a multiple of a given float64 divisor, returns true if _number_ is **multiple** of _divisor_.

    https://github.com/scubainwestpalm/fetch-rewards-receipt-processer-challenge-submission/blob/79d04f22707fd94d962b74be44a9d70438831107/utils/point_helpers.go#L22-L24

    ###### _Example_

    ```golang
    ...
    isMultipleOf := utils.Float64IsMultipleOf(1.00, 0.25)
    fmt.Println(isMultipleOf)
    ...
    ```

    ###### **Outputs:**

    `true`

<span id="intismultipleofMethod"></span>
* #### _**IntIsMultipleOf**_ <sub> _number **int**, divisor **int**_</sub> <sup> _bool_ </sup>

    Checks if integer is a multiple of a given integer divisor, returns true if _number_ is **multiple** of _divisor_.

    https://github.com/scubainwestpalm/fetch-rewards-receipt-processer-challenge-submission/blob/79d04f22707fd94d962b74be44a9d70438831107/utils/point_helpers.go#L26-L28

    ###### _Example_

    ```golang
    ...
    isMultipleOf := utils.IntIsMultipleOf(6, 3)
    fmt.Println(isMultipleOf)
    ...
    ```

    ###### **Outputs:**

    `true`