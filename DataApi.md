# DataApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createFileGet**](DataApi.md#createFileGet) | **GET** /create-file | Created .csv file with data according to your parameters you want to know
[**downloadGet**](DataApi.md#downloadGet) | **GET** /download | Download .csv file with link what hardcode in app.
[**searchGet**](DataApi.md#searchGet) | **GET** /search | Get transaction items


<a name="createFileGet"></a>
# **createFileGet**
> createFileGet()

Created .csv file with data according to your parameters you want to know



### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

<a name="downloadGet"></a>
# **downloadGet**
> downloadGet()

Download .csv file with link what hardcode in app.


### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

<a name="searchGet"></a>
# **searchGet**
> ModelsDataParse searchGet(transaction, terminal, status, payment, dateFrom, dateTo, narrative)

Get transaction items


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transaction** | **Integer**| Search transaction by transaction ID. You can use several IDs to see several transaction (like transaction&#x3D;1,2,10 ). This parameter is incompatible with other. | [optional]
 **terminal** | **Integer**| Search transaction by terminal ID. You can use several IDs to see several transaction (like transaction&#x3D;3507,3510 ). This parameter is incompatible with other. | [optional]
 **status** | **String**| Search transaction by status. You can use several statuses to see several transaction (like transaction&#x3D;accepted,declined ). This parameter is incompatible with other. | [optional]
 **payment** | **String**| Search transaction by payment type. You can use several types to see several transaction (like transaction&#x3D;cash,card ). This parameter is incompatible with other. | [optional]
 **dateFrom** | **String**| Search transaction in the selected date interval. Must be used together with &#x60;dateTo&#x60; and is incompatible with other. | [optional]
 **dateTo** | **String**| Search transaction in the selected date interval. Must be used together with &#x60;dateFrom&#x60; and is incompatible with other. | [optional]
 **narrative** | **String**| Search transaction by partially specified payment narrative. This parameter is incompatible with other. | [optional]

### Return type

[**ModelsDataParse**](ModelsDataParse.md)


### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

