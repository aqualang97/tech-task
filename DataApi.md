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

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DataApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8080");

    DataApi apiInstance = new DataApi(defaultClient);
    try {
      apiInstance.createFileGet();
    } catch (ApiException e) {
      System.err.println("Exception when calling DataApi#createFileGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

<a name="downloadGet"></a>
# **downloadGet**
> downloadGet()

Download .csv file with link what hardcode in app.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DataApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8080");

    DataApi apiInstance = new DataApi(defaultClient);
    try {
      apiInstance.downloadGet();
    } catch (ApiException e) {
      System.err.println("Exception when calling DataApi#downloadGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

No authorization required

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

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DataApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8080");

    DataApi apiInstance = new DataApi(defaultClient);
    Integer transaction = 56; // Integer | Search transaction by transaction ID. You can use several IDs to see several transaction (like transaction=1,2,10 ). This parameter is incompatible with other.
    Integer terminal = 56; // Integer | Search transaction by terminal ID. You can use several IDs to see several transaction (like transaction=3507,3510 ). This parameter is incompatible with other.
    String status = "status_example"; // String | Search transaction by status. You can use several statuses to see several transaction (like transaction=accepted,declined ). This parameter is incompatible with other.
    String payment = "payment_example"; // String | Search transaction by payment type. You can use several types to see several transaction (like transaction=cash,card ). This parameter is incompatible with other.
    String dateFrom = "dateFrom_example"; // String | Search transaction in the selected date interval. Must be used together with `dateTo` and is incompatible with other.
    String dateTo = "dateTo_example"; // String | Search transaction in the selected date interval. Must be used together with `dateFrom` and is incompatible with other.
    String narrative = "narrative_example"; // String | Search transaction by partially specified payment narrative. This parameter is incompatible with other.
    try {
      ModelsDataParse result = apiInstance.searchGet(transaction, terminal, status, payment, dateFrom, dateTo, narrative);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DataApi#searchGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

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

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

