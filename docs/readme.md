# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [Health.proto](#Health.proto)
    - [Health](#health.Health)
  
    - [HealthService](#health.HealthService)
  
- [Store.proto](#Store.proto)
    - [ChangeStatusRequest](#store.ChangeStatusRequest)
    - [CreateRequest](#store.CreateRequest)
    - [FindRequest](#store.FindRequest)
    - [ListRequest](#store.ListRequest)
    - [Store](#store.Store)
    - [Stores](#store.Stores)
  
    - [Status](#store.Status)
  
    - [StoreService](#store.StoreService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="Health.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## Health.proto



<a name="health.Health"></a>

### Health



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| alive | [bool](#bool) |  |  |





 

 

 


<a name="health.HealthService"></a>

### HealthService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Check | [.google.protobuf.Empty](#google.protobuf.Empty) | [Health](#health.Health) |  |

 



<a name="Store.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## Store.proto



<a name="store.ChangeStatusRequest"></a>

### ChangeStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [Status](#store.Status) |  |  |






<a name="store.CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| uri | [string](#string) |  |  |






<a name="store.FindRequest"></a>

### FindRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |






<a name="store.ListRequest"></a>

### ListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| status | [Status](#store.Status) |  |  |






<a name="store.Store"></a>

### Store



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| uri | [string](#string) |  |  |
| status | [Status](#store.Status) |  |  |
| created | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="store.Stores"></a>

### Stores



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stores | [Store](#store.Store) | repeated |  |





 


<a name="store.Status"></a>

### Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| ACTIVE | 0 |  |
| INACTIVE | 1 |  |


 

 


<a name="store.StoreService"></a>

### StoreService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Find | [FindRequest](#store.FindRequest) | [Store](#store.Store) |  |
| List | [ListRequest](#store.ListRequest) | [Stores](#store.Stores) |  |
| Create | [CreateRequest](#store.CreateRequest) | [Store](#store.Store) |  |
| ChangeStatus | [ChangeStatusRequest](#store.ChangeStatusRequest) | [Store](#store.Store) |  |

 



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

