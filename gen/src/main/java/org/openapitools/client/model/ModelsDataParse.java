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


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.math.BigDecimal;

/**
 * ModelsDataParse
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-10-30T20:18:54.300124+02:00[Europe/Kiev]")
public class ModelsDataParse {
  public static final String SERIALIZED_NAME_TRANSACTION_ID = "transaction_id";
  @SerializedName(SERIALIZED_NAME_TRANSACTION_ID)
  private Integer transactionId;

  public static final String SERIALIZED_NAME_REQUEST_ID = "request_id";
  @SerializedName(SERIALIZED_NAME_REQUEST_ID)
  private Integer requestId;

  public static final String SERIALIZED_NAME_TERMINAL_ID = "terminal_id";
  @SerializedName(SERIALIZED_NAME_TERMINAL_ID)
  private Integer terminalId;

  public static final String SERIALIZED_NAME_PARTNER_OBJECT_ID = "partner_object_id";
  @SerializedName(SERIALIZED_NAME_PARTNER_OBJECT_ID)
  private Integer partnerObjectId;

  public static final String SERIALIZED_NAME_AMOUNT_TOTAL = "amount_total";
  @SerializedName(SERIALIZED_NAME_AMOUNT_TOTAL)
  private Integer amountTotal;

  public static final String SERIALIZED_NAME_AMOUNT_ORIGINAL = "amount_original";
  @SerializedName(SERIALIZED_NAME_AMOUNT_ORIGINAL)
  private Integer amountOriginal;

  public static final String SERIALIZED_NAME_COMMISSION_PS = "commission_ps";
  @SerializedName(SERIALIZED_NAME_COMMISSION_PS)
  private BigDecimal commissionPs;

  public static final String SERIALIZED_NAME_COMMISSION_CLIENT = "commission_client";
  @SerializedName(SERIALIZED_NAME_COMMISSION_CLIENT)
  private Integer commissionClient;

  public static final String SERIALIZED_NAME_COMMISSION_PROVIDER = "commission_provider";
  @SerializedName(SERIALIZED_NAME_COMMISSION_PROVIDER)
  private BigDecimal commissionProvider;

  public static final String SERIALIZED_NAME_DATE_INPUT = "date_input";
  @SerializedName(SERIALIZED_NAME_DATE_INPUT)
  private String dateInput;

  public static final String SERIALIZED_NAME_DATE_POST = "date_post";
  @SerializedName(SERIALIZED_NAME_DATE_POST)
  private String datePost;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;

  public static final String SERIALIZED_NAME_PAYMENT_TYPE = "payment_type";
  @SerializedName(SERIALIZED_NAME_PAYMENT_TYPE)
  private String paymentType;

  public static final String SERIALIZED_NAME_PAYMENT_NUMBER = "payment_number";
  @SerializedName(SERIALIZED_NAME_PAYMENT_NUMBER)
  private String paymentNumber;

  public static final String SERIALIZED_NAME_SERVICE_ID = "service_id";
  @SerializedName(SERIALIZED_NAME_SERVICE_ID)
  private Integer serviceId;

  public static final String SERIALIZED_NAME_SERVICE = "service";
  @SerializedName(SERIALIZED_NAME_SERVICE)
  private String service;

  public static final String SERIALIZED_NAME_PAYEE_ID = "payee_id";
  @SerializedName(SERIALIZED_NAME_PAYEE_ID)
  private Integer payeeId;

  public static final String SERIALIZED_NAME_PAYEE_NAME = "payee_name";
  @SerializedName(SERIALIZED_NAME_PAYEE_NAME)
  private String payeeName;

  public static final String SERIALIZED_NAME_PAYEE_BANK_MFO = "payee_bank_mfo";
  @SerializedName(SERIALIZED_NAME_PAYEE_BANK_MFO)
  private Integer payeeBankMfo;

  public static final String SERIALIZED_NAME_PAYEE_BANK_ACCOUNT = "payee_bank_account";
  @SerializedName(SERIALIZED_NAME_PAYEE_BANK_ACCOUNT)
  private String payeeBankAccount;

  public static final String SERIALIZED_NAME_PAYMENT_NARRATIVE = "payment_narrative";
  @SerializedName(SERIALIZED_NAME_PAYMENT_NARRATIVE)
  private String paymentNarrative;


  public ModelsDataParse transactionId(Integer transactionId) {
    
    this.transactionId = transactionId;
    return this;
  }

   /**
   * Get transactionId
   * @return transactionId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getTransactionId() {
    return transactionId;
  }


  public void setTransactionId(Integer transactionId) {
    this.transactionId = transactionId;
  }


  public ModelsDataParse requestId(Integer requestId) {
    
    this.requestId = requestId;
    return this;
  }

   /**
   * Get requestId
   * @return requestId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getRequestId() {
    return requestId;
  }


  public void setRequestId(Integer requestId) {
    this.requestId = requestId;
  }


  public ModelsDataParse terminalId(Integer terminalId) {
    
    this.terminalId = terminalId;
    return this;
  }

   /**
   * Get terminalId
   * @return terminalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getTerminalId() {
    return terminalId;
  }


  public void setTerminalId(Integer terminalId) {
    this.terminalId = terminalId;
  }


  public ModelsDataParse partnerObjectId(Integer partnerObjectId) {
    
    this.partnerObjectId = partnerObjectId;
    return this;
  }

   /**
   * Get partnerObjectId
   * @return partnerObjectId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getPartnerObjectId() {
    return partnerObjectId;
  }


  public void setPartnerObjectId(Integer partnerObjectId) {
    this.partnerObjectId = partnerObjectId;
  }


  public ModelsDataParse amountTotal(Integer amountTotal) {
    
    this.amountTotal = amountTotal;
    return this;
  }

   /**
   * Get amountTotal
   * @return amountTotal
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getAmountTotal() {
    return amountTotal;
  }


  public void setAmountTotal(Integer amountTotal) {
    this.amountTotal = amountTotal;
  }


  public ModelsDataParse amountOriginal(Integer amountOriginal) {
    
    this.amountOriginal = amountOriginal;
    return this;
  }

   /**
   * Get amountOriginal
   * @return amountOriginal
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getAmountOriginal() {
    return amountOriginal;
  }


  public void setAmountOriginal(Integer amountOriginal) {
    this.amountOriginal = amountOriginal;
  }


  public ModelsDataParse commissionPs(BigDecimal commissionPs) {
    
    this.commissionPs = commissionPs;
    return this;
  }

   /**
   * Get commissionPs
   * @return commissionPs
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCommissionPs() {
    return commissionPs;
  }


  public void setCommissionPs(BigDecimal commissionPs) {
    this.commissionPs = commissionPs;
  }


  public ModelsDataParse commissionClient(Integer commissionClient) {
    
    this.commissionClient = commissionClient;
    return this;
  }

   /**
   * Get commissionClient
   * @return commissionClient
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getCommissionClient() {
    return commissionClient;
  }


  public void setCommissionClient(Integer commissionClient) {
    this.commissionClient = commissionClient;
  }


  public ModelsDataParse commissionProvider(BigDecimal commissionProvider) {
    
    this.commissionProvider = commissionProvider;
    return this;
  }

   /**
   * Get commissionProvider
   * @return commissionProvider
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCommissionProvider() {
    return commissionProvider;
  }


  public void setCommissionProvider(BigDecimal commissionProvider) {
    this.commissionProvider = commissionProvider;
  }


  public ModelsDataParse dateInput(String dateInput) {
    
    this.dateInput = dateInput;
    return this;
  }

   /**
   * Get dateInput
   * @return dateInput
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDateInput() {
    return dateInput;
  }


  public void setDateInput(String dateInput) {
    this.dateInput = dateInput;
  }


  public ModelsDataParse datePost(String datePost) {
    
    this.datePost = datePost;
    return this;
  }

   /**
   * Get datePost
   * @return datePost
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDatePost() {
    return datePost;
  }


  public void setDatePost(String datePost) {
    this.datePost = datePost;
  }


  public ModelsDataParse status(String status) {
    
    this.status = status;
    return this;
  }

   /**
   * Get status
   * @return status
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStatus() {
    return status;
  }


  public void setStatus(String status) {
    this.status = status;
  }


  public ModelsDataParse paymentType(String paymentType) {
    
    this.paymentType = paymentType;
    return this;
  }

   /**
   * Get paymentType
   * @return paymentType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPaymentType() {
    return paymentType;
  }


  public void setPaymentType(String paymentType) {
    this.paymentType = paymentType;
  }


  public ModelsDataParse paymentNumber(String paymentNumber) {
    
    this.paymentNumber = paymentNumber;
    return this;
  }

   /**
   * Get paymentNumber
   * @return paymentNumber
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPaymentNumber() {
    return paymentNumber;
  }


  public void setPaymentNumber(String paymentNumber) {
    this.paymentNumber = paymentNumber;
  }


  public ModelsDataParse serviceId(Integer serviceId) {
    
    this.serviceId = serviceId;
    return this;
  }

   /**
   * Get serviceId
   * @return serviceId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getServiceId() {
    return serviceId;
  }


  public void setServiceId(Integer serviceId) {
    this.serviceId = serviceId;
  }


  public ModelsDataParse service(String service) {
    
    this.service = service;
    return this;
  }

   /**
   * Get service
   * @return service
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getService() {
    return service;
  }


  public void setService(String service) {
    this.service = service;
  }


  public ModelsDataParse payeeId(Integer payeeId) {
    
    this.payeeId = payeeId;
    return this;
  }

   /**
   * Get payeeId
   * @return payeeId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getPayeeId() {
    return payeeId;
  }


  public void setPayeeId(Integer payeeId) {
    this.payeeId = payeeId;
  }


  public ModelsDataParse payeeName(String payeeName) {
    
    this.payeeName = payeeName;
    return this;
  }

   /**
   * Get payeeName
   * @return payeeName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPayeeName() {
    return payeeName;
  }


  public void setPayeeName(String payeeName) {
    this.payeeName = payeeName;
  }


  public ModelsDataParse payeeBankMfo(Integer payeeBankMfo) {
    
    this.payeeBankMfo = payeeBankMfo;
    return this;
  }

   /**
   * Get payeeBankMfo
   * @return payeeBankMfo
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getPayeeBankMfo() {
    return payeeBankMfo;
  }


  public void setPayeeBankMfo(Integer payeeBankMfo) {
    this.payeeBankMfo = payeeBankMfo;
  }


  public ModelsDataParse payeeBankAccount(String payeeBankAccount) {
    
    this.payeeBankAccount = payeeBankAccount;
    return this;
  }

   /**
   * Get payeeBankAccount
   * @return payeeBankAccount
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPayeeBankAccount() {
    return payeeBankAccount;
  }


  public void setPayeeBankAccount(String payeeBankAccount) {
    this.payeeBankAccount = payeeBankAccount;
  }


  public ModelsDataParse paymentNarrative(String paymentNarrative) {
    
    this.paymentNarrative = paymentNarrative;
    return this;
  }

   /**
   * Get paymentNarrative
   * @return paymentNarrative
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPaymentNarrative() {
    return paymentNarrative;
  }


  public void setPaymentNarrative(String paymentNarrative) {
    this.paymentNarrative = paymentNarrative;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ModelsDataParse modelsDataParse = (ModelsDataParse) o;
    return Objects.equals(this.transactionId, modelsDataParse.transactionId) &&
        Objects.equals(this.requestId, modelsDataParse.requestId) &&
        Objects.equals(this.terminalId, modelsDataParse.terminalId) &&
        Objects.equals(this.partnerObjectId, modelsDataParse.partnerObjectId) &&
        Objects.equals(this.amountTotal, modelsDataParse.amountTotal) &&
        Objects.equals(this.amountOriginal, modelsDataParse.amountOriginal) &&
        Objects.equals(this.commissionPs, modelsDataParse.commissionPs) &&
        Objects.equals(this.commissionClient, modelsDataParse.commissionClient) &&
        Objects.equals(this.commissionProvider, modelsDataParse.commissionProvider) &&
        Objects.equals(this.dateInput, modelsDataParse.dateInput) &&
        Objects.equals(this.datePost, modelsDataParse.datePost) &&
        Objects.equals(this.status, modelsDataParse.status) &&
        Objects.equals(this.paymentType, modelsDataParse.paymentType) &&
        Objects.equals(this.paymentNumber, modelsDataParse.paymentNumber) &&
        Objects.equals(this.serviceId, modelsDataParse.serviceId) &&
        Objects.equals(this.service, modelsDataParse.service) &&
        Objects.equals(this.payeeId, modelsDataParse.payeeId) &&
        Objects.equals(this.payeeName, modelsDataParse.payeeName) &&
        Objects.equals(this.payeeBankMfo, modelsDataParse.payeeBankMfo) &&
        Objects.equals(this.payeeBankAccount, modelsDataParse.payeeBankAccount) &&
        Objects.equals(this.paymentNarrative, modelsDataParse.paymentNarrative);
  }

  @Override
  public int hashCode() {
    return Objects.hash(transactionId, requestId, terminalId, partnerObjectId, amountTotal, amountOriginal, commissionPs, commissionClient, commissionProvider, dateInput, datePost, status, paymentType, paymentNumber, serviceId, service, payeeId, payeeName, payeeBankMfo, payeeBankAccount, paymentNarrative);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ModelsDataParse {\n");
    sb.append("    transactionId: ").append(toIndentedString(transactionId)).append("\n");
    sb.append("    requestId: ").append(toIndentedString(requestId)).append("\n");
    sb.append("    terminalId: ").append(toIndentedString(terminalId)).append("\n");
    sb.append("    partnerObjectId: ").append(toIndentedString(partnerObjectId)).append("\n");
    sb.append("    amountTotal: ").append(toIndentedString(amountTotal)).append("\n");
    sb.append("    amountOriginal: ").append(toIndentedString(amountOriginal)).append("\n");
    sb.append("    commissionPs: ").append(toIndentedString(commissionPs)).append("\n");
    sb.append("    commissionClient: ").append(toIndentedString(commissionClient)).append("\n");
    sb.append("    commissionProvider: ").append(toIndentedString(commissionProvider)).append("\n");
    sb.append("    dateInput: ").append(toIndentedString(dateInput)).append("\n");
    sb.append("    datePost: ").append(toIndentedString(datePost)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    paymentType: ").append(toIndentedString(paymentType)).append("\n");
    sb.append("    paymentNumber: ").append(toIndentedString(paymentNumber)).append("\n");
    sb.append("    serviceId: ").append(toIndentedString(serviceId)).append("\n");
    sb.append("    service: ").append(toIndentedString(service)).append("\n");
    sb.append("    payeeId: ").append(toIndentedString(payeeId)).append("\n");
    sb.append("    payeeName: ").append(toIndentedString(payeeName)).append("\n");
    sb.append("    payeeBankMfo: ").append(toIndentedString(payeeBankMfo)).append("\n");
    sb.append("    payeeBankAccount: ").append(toIndentedString(payeeBankAccount)).append("\n");
    sb.append("    paymentNarrative: ").append(toIndentedString(paymentNarrative)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}
