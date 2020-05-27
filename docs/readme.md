# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [Health.proto](#Health.proto)
    - [Health](#health.v1.Health)
  
    - [HealthService](#health.v1.HealthService)
  
- [Store.proto](#Store.proto)
    - [ChangeStatusRequest](#store.v1.ChangeStatusRequest)
    - [CreateRequest](#store.v1.CreateRequest)
    - [FindRequest](#store.v1.FindRequest)
    - [ListRequest](#store.v1.ListRequest)
    - [Store](#store.v1.Store)
    - [Stores](#store.v1.Stores)
  
    - [Status](#store.v1.Status)
  
    - [StoreService](#store.v1.StoreService)
  
- [Swagger.proto](#Swagger.proto)
- [Scalar Value Types](#scalar-value-types)



<a name="Health.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## Health.proto



<a name="health.v1.Health"></a>

### Health



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alive | [bool](#bool) |  |  |





 

 

 


<a name="health.v1.HealthService"></a>

### HealthService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Check | [.google.protobuf.Empty](#google.protobuf.Empty) | [Health](#health.v1.Health) |  |

 



<a name="Store.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## Store.proto



<a name="store.v1.ChangeStatusRequest"></a>

### ChangeStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id of the store |
| status | [Status](#store.v1.Status) |  | status to be changed |






<a name="store.v1.CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Request name |
| uri | [string](#string) |  | Request URI |






<a name="store.v1.FindRequest"></a>

### FindRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="store.v1.ListRequest"></a>

### ListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| status | [Status](#store.v1.Status) |  |  |






<a name="store.v1.Store"></a>

### Store



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| uri | [string](#string) |  |  |
| status | [Status](#store.v1.Status) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="store.v1.Stores"></a>

### Stores



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stores | [Store](#store.v1.Store) | repeated |  |





 


<a name="store.v1.Status"></a>

### Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| ACTIVE | 0 |  |
| INACTIVE | 1 |  |


 

 


<a name="store.v1.StoreService"></a>

### StoreService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Find | [FindRequest](#store.v1.FindRequest) | [Store](#store.v1.Store) |  |
| List | [ListRequest](#store.v1.ListRequest) | [Stores](#store.v1.Stores) |  |
| Create | [CreateRequest](#store.v1.CreateRequest) | [Store](#store.v1.Store) |  |
| ChangeStatus | [ChangeStatusRequest](#store.v1.ChangeStatusRequest) | [Store](#store.v1.Store) |  |

 



<a name="Swagger.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## Swagger.proto


 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

