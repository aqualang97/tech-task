/*
 * Go Restful API with Swagger
 * Simple swagger implementation in Go HTTP
 *
 * The version of the OpenAPI document: 0.1.0
 * Contact: aqualang9797@gmail.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.ModelsDataParse;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for DataApi
 */
@Ignore
public class DataApiTest {

    private final DataApi api = new DataApi();

    
    /**
     * Download .csv file with link what hardcode in app.
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void downloadGetTest() throws ApiException {
        api.downloadGet();

        // TODO: test validations
    }
    
    /**
     * Get transaction items
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void searchGetTest() throws ApiException {
        Integer transaction = null;
        Integer terminal = null;
        String status = null;
        String payment = null;
        String dateFrom = null;
        String dateTo = null;
        String narrative = null;
        ModelsDataParse response = api.searchGet(transaction, terminal, status, payment, dateFrom, dateTo, narrative);

        // TODO: test validations
    }
    
}