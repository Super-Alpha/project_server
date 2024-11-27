
##天语计费中心类接口

**生产环境地址**：https://callcenter.58corp.com

**测试环境地址**：http://t-callcenter.58v5.cn

**本地接口地址**：http://callcenter.local

### 一、账单查询

#### **1、账单汇总**

1.1 根据城市、业务线、通话起止时间、供应商、话务类型、号码类型等查询账单汇总数据

* （1）请求格式

     >|URL| /api/bill/summary|
     >|:----:|:----:|
     >|支持格式|json|
     >|请求方式|get|

* （2）参数说明  （Query）

    >|名称|类型|是否必选|描述|
    >|:-----:|:----:|:----:|:----:|
    >|city_id|int|否|城市|
    >|bg|string|否|BG|
    >|bu|string|否|BU| 
    >|call_start_time|string|否|通话开始时间|
    >|call_end_time|string|否|通话结束时间|
    >|operator_code|int|否|供应商|
    >|call_type|string|否|话务类型|
    >|number_type|string|否|号码类型|
    >|page|int|否|分页数|
    >|limit|int|否|条目| 
  
* （3）返回结果说明

   >|名称|类型|说明|
   >|:-----:|:----:|:-----:|
   >|code|int|状态码|
   >|message|string|状态信息|
   >|data|map|返回数据|
   >|data.items|array|账单数据|
   >|data.items.*.city_id|int|城市|
   >|data.items.*.business_line|string|业务线|
   >|data.items.*.call_time|datetime|通话时间|
   >|data.items.*.workplace_id|int|职场|
   >|data.items.*.operator_code|int|供应商|
   >|data.items.*.call_type|string|话务类型|
   >|data.items.*.number_type|string|号码类型|
   >|data.items.*.call_amount|float|话务量汇总|
   >|data.items.*.call_duration|float|通话时长汇总|
   >|data.items.*.weight_duration|float|加权时长汇总|
   >|data.items.*.fee_total|float|费用汇总|
   >|data.total|int|账单总数量|
   >|user|string|当前用户|

   例如：
    ```json
         {
             "code": 0,
             "message": "success",
             "data": {
                "items":[
                    {
                    "id":1,
                    "city_id": 79,
                    "business_line": "黄页",
                    "call_time": "2022-03-01 15:00:00",
                    "workplace_id": 10,
                    "operator_code": 53,
                    "call_type": "本地固话",
                    "number_type": "联通手机",
                    "call_amount": 2310.10,
                    "call_duration": 2310.10,
                    "weight_duration": 2310.10,
                    "cost_total": 1528.59
                    }
                 ],
                "total":100
             },
             "user": "Jack"
         }
    ```
    ```json
        {
            "code": -1,
            "message": "error",
            "user": "Jack"
        }
    ```

* （4）状态码说明

   >|code|释义|备注|
   >|:----:|:----:|:----:|
   >|0|查询成功|success|
   >|-1|查询失败|参数有误|
  

#### **2、账单上报**

2.1 填写账单信息，上报账单数据

* （1）请求格式

  >|URL| /api/bill/operator/create|
  >|:----:|:----:|
  >|支持格式|json|
  >|请求方式|post|
  
* （2）参数说明（Body）
  
  >|名称|类型|是否必选|描述|
  >|:----:|:----:|:----:|:----:|
  >|bill_date|string|是|账单日期|
  >|did|string|是|DID 号段|
  >|city_id|int|是|城市|
  >|workplace_id|int|是|职场|
  >|operator_code|int|是|供应商|
  >|bg|string|是|BG|
  >|bu|string|是|BU|
  >|fee_strategy|string|是|资费标准|
  >|fee_state|bool|是|资费状态：true/一致、false/不一致|
  >|call_duration|int|是|通话总时长（单位：秒）|
  >|tel_fee_total|float|是|话费总额（单位：元）|
  >|remark|string|否|备注说明|

* （3）返回结果说明

  >|名称|类型|说明|
  >|:-----:|:----:|:-----:|
  >|code|int|状态码|
  >|message|string|状态信息|

  例如：

    ```json
         {
             "code": 0,
             "message": "success",
             "user": "Jack01"
         }
    ```
    ```json
        {
            "code": -1,
            "message": "error",
            "user": "Jack"
        }
    ```

* （4）状态码说明

  >|code|释义|备注|
  >|:----:|:----:|:----:|
  >|0|上报成功|success|
  >|-1|上报失败|参数有误|


#### **3、供应商账单查询**

3.1 根据城市、职场、账期起止时间、供应商、上报状态、是否满足抵消等查询供应商账单数据

* （1）请求格式
  
  >|URL| /api/bill/operator/query|
  >|:----:|:----:|
  >|支持格式|json|
  >|请求方式|get|

* （2）参数说明 （Query）

  >|名称|类型|是否必选|描述|
  >|:-----:|:----:|:----:|:----:|
  >|city_id|int|否|城市|
  >|workplace_id|int|否|职场|
  >|bill_start_date|string|否|账期开始时间|
  >|bill_end_date|string|否|账期结束时间|
  >|operator_code|int|否|供应商|
  >|report_state|bool|否|上报状态|
  >|satisfy_minimum_charge|bool|否|是否满足抵消|

* （3）返回结果说明

  >|名称|类型|说明|
  >|:-----:|:----:|:-----:|
  >|code|int|状态码|
  >|message|string|状态信息|
  >|data|map|返回数据|
  >|items|array|供应商账单数据|
  >|data.items.*.city_id|int|城市|
  >|data.items.*.workplace_id|int|职场|
  >|data.items.*.operator_code|int|供应商|
  >|data.items.*.bill_date|datetime|账单日期|
  >|data.items.*.call_duration|float|通话时长汇总|
  >|data.items.*.weight_duration|float|加权时长汇总|
  >|data.items.*.call_amount|float|话务量汇总|
  >|data.items.*.tel_fee|float|实际话费汇总|
  >|data.items.*.call_fee|float|通话费用结算|
  >|data.items.*.minimum_charge_fee|float|低消费用|
  >|data.items.*.fixed_fee|float|固定费用|
  >|data.items.*.total_tel_fee|float|话费总费用|
  >|data.items.*.minimum_charge_compare|bool|抵消比较|
  >|data.items.*.report_state|bool|上报状态|
  >|total|int|供应商账单数量|
  >|user|string|当前用户|
  
  例如：
  
    ```json
         {
             "code": 0,
             "message": "success",
             "data": {
                "items":[
                    {
                      "id":1,
                      "city_id": 71,
                      "workplace_id": 67,
                      "operator_code": 43,
                      "bill_date": "2022-03-01",
                      "call_duration": 2310.10,
                      "weight_duration": 2310.00,
                      "call_amount": 2340.87,
                      "tel_fee": 1528.59,
                      "call_fee": 2310.32,
                      "minimum_charge_fee": 2345.76,
                      "fixed_fee": 4532.22,
                      "total_tel_fee": 9876.90,
                      "minimum_charge_compare": false,
                      "report_state": false
                    }
                 ],
                "total":68
             },
             "user": "Jack"
         }
    ```
    ```json
        {
            "code": -1,
            "message": "error",
            "user": "Jack"
        }
    ```
  
* （4）状态码说明

  >|code|释义|备注|
  >|:----:|:----:|:----:|
  >|0|查询成功|success|
  >|-1|查询失败|参数有误|


#### **4、业务线账单查询**

4.1 根据城市、业务线、账期起止时间等查询业务线账单数据

* （1）请求格式

  >|URL| /api/bill/businessline/query|
  >|:----:|:----:|
  >|支持格式|json|
  >|请求方式|get|
  
* （2）参数说明 （Query）

  >|参数|类型|是否必选|描述|
  >|:-----:|:----:|:----:|:----:|
  >|city_id|int|否|城市|
  >|bg|string|否|BG|
  >|bu|string|否|BU|
  >|bill_start_date|string|否|账期开始时间|
  >|bill_end_date|string|否|账期结束时间|

* （3）返回结果说明

  >|名称|类型|说明|
  >|:-----:|:----:|:-----:|
  >|code|int|状态码|
  >|message|string|状态信息|
  >|data|map|返回数据|
  >|items|array|业务线账单数据|
  >|data.items.*.city_id|int|城市|
  >|data.items.*.bg|string|BG|
  >|data.items.*.bu|string|BU|
  >|data.items.*.bill_date|string|账单日期|
  >|data.items.*.call_duration|float|通话时长汇总|
  >|data.items.*.weight_duration|float|加权时长汇总|
  >|data.items.*.fee|float|费用汇总|
  >|data.items.*.fee_after_discount|float|优惠后费用|
  >|data.items.*.call_amount|float|话务量汇总|
  >|total|int|业务线账单数量|
  >|user|string|当前用户|

  例如：

    ```json
         {
             "code": 0,
             "message": "success",
             "data": {
                "items":[
                 {
                  "id":1,
                  "city_id":71,
                  "business_line": "黄页",
                  "bill_date": "2021-01-22",
                  "call_duration": 2310.10,
                  "weight_duration": 2345.98,
                  "fee": 1528.59,
                  "fee_after_discount": 2310.0,
                  "tel_traffic": 2354.78
                  }
                ],
                "total":688
             },
             "user": "Jack"
         }
    ```
    ```json
        {
            "code": -1,
            "message": "error",
            "user": "Jack"
        }
    ```
  
* （4）状态码说明

  >|code|释义|备注|
  >|:----:|:----:|:----:|
  >|0|查询成功|success|
  >|-1|查询失败|参数有误|
  

#### **5、账单对账**

5.1 根据城市、职场、账期起止时间、供应商等查询账单对账数据

* （1）请求格式

  >|URL| /api/bill/verify|
  >|:----:|:----:|
  >|支持格式|json|
  >|请求方式|get|

* （2）参数说明 （Query）

  >|参数|类型|是否必选|描述|
  >|:-----:|:----:|:----:|:----:|
  >|city_id|int|否|城市|
  >|workplace_id|int|否|职场|
  >|bill_start_date|string|否|账期开始时间|
  >|bill_end_date|string|否|账期结束时间|
  >|operator_code|int|否|供应商|

* （3）返回结果说明

  >|名称|类型|说明|
  >|:-----:|:----:|:-----:|
  >|code|int|状态码|
  >|message|string|状态信息|
  >|data|map|返回数据|
  >|items|array|账单对账数据|
  >|data.items.*.city_id|int|城市|
  >|data.items.*.workplace_id|int|职场|
  >|data.items.*.operator_code|int|供应商|
  >|data.items.*.bill_date|string|账单日期|
  >|data.items.*.did|string|DID/虚拟号码|
  >|data.items.*.line_prefix|string|线路编号|
  >|data.items.*.operator_call_duration|float|供应商通话总时长|
  >|data.items.*.billing_call_duration|float|计费通话总时长|
  >|data.items.*.time_diff|string|时长差异|
  >|data.items.*.operator_total_fee|float|供应商话费总金额|
  >|data.items.*.billing_total_fee|float|计费话费总金额|
  >|data.items.*.fee_diff|string|金额差异|
  >|total|int|账单对账数量|
  >|user|string|当前用户|

  例如：

    ```json
         {
             "code": 0,
             "data": {
                "items":[
                    {
                      "id":1,
                      "city_id":71,
                      "workplace_id": 23,
                      "operator_code": 21,
                      "bill_date": "2022-02-02",
                      "did": "041188254440",
                      "line_prefix": "JN-02-DX-01",
                      "operator_call_duration": 673552,
                      "billing_duration": 34625.87,
                      "time_diff": "0.07%",
                      "operator_total_fee": 48000,
                      "billing_total_fee": 5342.97,
                      "fee_diff": "-0.017%"
                    }
                 ],
                "total":688
             },
             "message": "success",
             "user": "Jack01"
         }
    ```
    ```json
        {
            "code": -1,
            "message": "error",
            "user": "Jack"
        }
   ```

* （4）状态码说明

  >|code|释义|备注|
  >|:----:|:----:|:----:|
  >|0|查询成功|success|
  >|-1|查询失败|参数有误|


#### **6、账单分摊**

6.1 根据城市、职场、账期起止时间、供应商、上报状态、是否满足抵消、成本中心编码等查询账单对账数据

* （1）请求格式

  >|URL| /api/bill/share|
  >|:----:|:----:|
  >|支持格式|json|
  >|请求方式|get|

* （2）参数说明 （Query）

  >|名称|类型|是否必选|描述|
  >|:-----:|:----:|:----:|:----:|
  >|city_id|int|否|城市|
  >|workplace_id|int|否|职场|
  >|bill_start_date|string|否|账期开始日期|
  >|bill_end_date|string|否|账期结束日期|
  >|operator_code|int|否|供应商|
  >|report_state|bool|否|上报状态|
  >|satisfy_minimum_charge|bool|否|是否满足低消|
  >|cost_center_code|string|否|成本中心编码|

* （3）返回结果说明

  >|名称|类型|说明|
  >|:-----:|:----:|:-----:|
  >|code|int|状态码|
  >|message|string|状态信息|
  >|data|map|返回数据|
  >|items|array|账单分摊数据|
  >|data.items.*.city_id|int|城市|
  >|data.items.*.workplace_id|int|职场|
  >|data.items.*.bill_date|string|账单时间|
  >|data.items.*.operator_code|int|供应商|
  >|data.items.*.bg|string|BG|
  >|data.items.*.bu|string|BU|
  >|data.items.*.cost_center|string|成本中心|
  >|data.items.*.cost_center_code|string|成本中心编码|
  >|data.items.*.operator_bill_share|float|供应商账单分摊|
  >|data.items.*.system_bill_share|float|系统账单分摊|
  >|data.items.*.report_state|bool|上报状态|
  >|data.items.*.satisfy_minimum_charge|bool|是否满足低消|
  >|total|int|账单分摊数量|
  >|user|string|当前用户|

  例如：

    ```json
         {
             "code": 0,
             "message": "success",
             "data": {
                "items":[
                    {
                      "id":1,
                      "city_id":71,
                      "workplace_id": 12,
                      "bill_date": "2022-03-08",
                      "operator_code": 56,
                      "bg": "",
                      "bu": "",
                      "cost_center": "房产人力资源部",
                      "cost_center_code": "67.10.00.34.09",
                      "operator_bill_share": 106.92,
                      "system_bill_share": 39.46,
                      "report_state": true,
                      "satisfy_minimum_charge": false
                    }
                 ],
                "total":688
             },
             "user": "Jack01"
        }
    ```
    ```json
        {
            "code": -1,
            "message": "error",
            "user": "Jack"
        }
    ```

* （4）状态码说明

  >|code|释义|备注|
  >|:----:|:----:|:----:|
  >|0|查询成功|success|
  >|-1|查询失败|参数有误|


### 二、资费管理

#### **1、资费配置**

1.1 填写资费配置信息，创建资费配置记录

* （1）请求格式

  >|URL| /api/fee/config/:id|
  >|:----:|:----:|
  >|id|线路 ID|
  >|支持格式|json|
  >|请求方式|post|
  
* （2）参数说明 （Body）
  >|名称|类型|是否必选|描述|
  >|:----:|:----:|:----:|:----:|
  >|city_id|int|否|城市|
  >|workplace_id|int|否|职场|
  >|operator_code|int|否|供应商|
  >|line_batch|string|否|线路批次|
  >|line_quantity|int|否|线路数量| 
  >|number_detail|string|否|号码明细|
  >|did|string|否|DID号段|
  >|minimum_charge_unit_price|float|否|低消单价|
  >|minimum_charge_total_price|float|否|低消总价|
  >|month_rent_fee|float|否|月租费用|
  >|number_use_fee|float|否|号码使用费用|
  >|caller_identify_fee|float|否|来电显示费用|
  >|ring_fee|float|否|炫铃费|
  >|call_type|string|否|话务类型|
  >|billing_mode|string|否|计费模式|
  >|first_section_price|float|否|分段一单价|
  >|section_time|int|否|分段时间|
  >|first_section_discount|float|否|分段一折扣|
  >|second_section_price|float|否|分段二单价|
  >|second_section_discount|float|否|分段二折扣|
  >|billing_unit|string|否|计费单位|
  >|discount_name|string|否|优惠名称|
  >|tel_fee_threshold|float|否|话费阈值|
  >|discount_type|string|否|优惠类型|
  >|max_discount|string|否|最大折扣|
  >|fee_remark|string|否|资费说明|

* (3) 请求体举例

   ```json
         {
              "city_id": 1,
              "workplace_id": 2,
              "operator_code": 3,
              "line_batch": "E1-联通-北京-03",
              "line_quantity": 3,
              "did_detail": "",
              "did": "3436437",
              "minimum_charge_unit_price": 465.8,
              "minimum_charge_total_price": 3435.98,
              "month_rent_fee": 48578.4,
              "number_use_fee": 386.38,
              "caller_identify_fee": 387,
              "ring_fee": 3232,
              "call_fee_config": {
                  "0":["本地手机","单一",0.16,5,0.98,1.0,0.97],
                  "1":["本地固化","单一",0.16,5,0.98,1.0,0.97]
              },
              "discount_fee_config": {
                  "0":["超阈值优惠",50,""],
                  "1":["超阈值优惠",50,""]
              },
              "fee_remark": ""
         }
    ```

* （4）返回结果说明

  >|名称|类型|说明|
  >|:-----:|:----:|:-----:|
  >|code|int|状态码|
  >|message|string|状态信息|

  例如：

    ```json
         {
             "code": 0,
             "message": "success",
             "user": "Jack"
         }
    ```
    ```json
        {
            "code": -1,
            "message": "error",
            "user": "Jack"
        }
    ```

* （5）状态码说明

  >|code|释义|备注|
  >|:----:|:----:|:----:|
  >|0|上报成功|success|
  >|-1|上报失败|参数有误|

#### **2、资费管理**

2.1 根据城市、职场、供应商等查询资费管理数据

* （1）请求格式

  >|URL| /api/bill/fee/manage|
  >|:----:|:----:|
  >|支持格式|json|
  >|请求方式|get|
  
* （2）参数说明 （Query）

  >|名称|类型|是否必选|描述|
  >|:----:|:----:|:----:|:----:|
  >|city_id|int|否|城市|
  >|workplace_id|int|否|职场|
  >|operator_code|int|否|供应商|


* （3）返回结果说明

  >|名称|类型|说明|
  >|:-----:|:----:|:-----:|
  >|code|int|状态码|
  >|message|string|状态信息|
  >|data|map|返回数据|
  >|items|array|资费管理数据|
  >|data.items.*.city_id|int|城市|
  >|data.items.*.workplace_id|int|职场|
  >|data.items.*.operator_code|int|供应商|
  >|data.items.*.line_batch_num|string|线路批次号|
  >|data.items.*.line_quantity|int|线路数量|
  >|data.items.*.did|string|DID号码|
  >|data.items.*.minimum_charge_unit_price|float|低消单价|
  >|data.items.*.minimum_charge_total_price|float|低消总价|
  >|data.items.*.month_rent_total_price|float|月租总价|
  >|data.items.*.number_use_fee|float|号码使用费用|
  >|data.items.*.caller_identify_fee|float|来电显示费用|
  >|data.items.*.ring_fee|float|炫铃使用费用|
  >|data.items.*.fee_remark|string|资费说明|
  >|data.items.*.take_effect_time|datetime|生效时间|
  >|data.items.*.invalid_time|datetime|失效时间|
  >|data.items.*.state|bool|状态|
  >|total|int|资费管理数量|
  >|user|string|当前用户|
  
  例如：

    ```json
        {
        "code": 0,
        "message": "success",
        "data": {
                "items":[
                    {
                      "id":1,
                      "city_id":1,
                      "workplace_id": 54,
                      "operator_code": 78,
                      "line_batch_code": "JD-NAME-KING",
                      "line_amount": 8,
                      "did": "8343",
                      "minimum_charge_unit_price": 8000.0,
                      "minimum_charge_total_price": 40000.0,
                      "month_rent_total_price": 24000.0,
                      "number_use_fee": 0.0,
                      "caller_identification_fee": 0.0,
                      "ring_fee": 0.0,
                      "charge_remark": "",
                      "take_effect_time": "2018-09-03",
                      "invalid_time": "2022-09-29"
                    }
                 ],
                "total":88
             },
             "user": "Jack"
        }
    ```
    ```json
        {
            "code": -1,
            "message": "error",
            "user": "Jack"
        }
    ```
  
* （4）状态码说明

  >|code|释义|备注|
  >|:----:|:----:|:----:|
  >|0|查询成功|success|
  >|-1|查询失败|参数有误|


2.2 编辑资费管理数据

* （1）请求格式

  >|URL| /api/bill/fee/edit/:id|
  >|:----:|:----:|
  >|id|供应商资费数据 ID|
  >|支持格式|json|
  >|请求方式|put|

* （2）参数说明 （body）

  >|名称|类型|是否必选|描述|
  >|:----:|:----:|:----:|:----:|
  >|city_id|int|否|城市|
  >|workplace_id|int|否|职场|
  >|operator_code|int|否|供应商|
  >|line_batch|string|否|线路批次|
  >|line_quantity|int|否|线路数量| 
  >|did|string|否|DID号段|
  >|minimum_charge_unit_price|float|否|低消单价|
  >|minimum_charge_total_price|float|否|低消总价|
  >|month_rent_total_price|float|否|月租总价|
  >|number_use_fee|float|否|号码使用费用|
  >|caller_identify_fee|float|否|来电显示费用|
  >|ring_fee|float|否|炫铃使用费用|
  >|fee_remark|string|否|资费说明|
  >|take_effect_time|datetime|否|生效时间|
  >|invalid_time|datetime|否|失效时间|

* （3）返回结果说明

    ```json
        {
            "code": 0,
            "message": "success",
            "data": {
              "city_id": 1,
              "workplace_id": 2,
              ".....": ""
            },
            "user": "Jack"
        }
    ```

* （4）状态码说明

  >|code|释义|备注|
  >|:----:|:----:|:----:|
  >|0|编辑成功|success|
  >|-1|编辑失败|参数有误|

2.3 资费管理数据状态变更

* （1）请求格式

  >|URL| /api/bill/fee/manage/state/:id|
    >|:----:|:----:|
  >|id|资费数据 ID|
  >|支持格式|json|
  >|请求方式|post|

* （2）参数说明 （body）

  >|名称|类型|是否必选|描述|
  >|:----:|:----:|:----:|:----:|
  >|state|bool|是|状态（true:启用，false:停用）|
  
* （3）返回参数说明

  ```json
        {
            "code": 0,
            "message": "success",
            "data": "change state success",
            "user": "Jack"
        }
    ```

* （4）状态码说明

  >|code|释义|备注|
  >|:----:|:----:|:----:|
  >|0|变更成功|success|
  >|-1|变更失败|参数有误|


### 三、其他

#### **1、话务类型**

1.1 获得话务类型数据

* （1）请求格式（Query）

  >|URL|/api/bill/calltype|
  >|:----:|:----:|
  >|支持格式|json|
  >|请求方式|get|

* （2）返回结果说明

  >|名称|类型|说明|
  >|:-----:|:----:|:-----:|
  >|code|int|状态码|
  >|message|string|状态信息|
  >|data|map|返回数据|

#### **2、号码类型**

1.1 获得号码类型数据

* （1）请求格式（Query）

  >|URL|/api/bill/numtype|
  >|:----:|:----:|
  >|支持格式|json|
  >|请求方式|get|

* （2）返回结果说明

  >|名称|类型|说明|
  >|:-----:|:----:|:-----:|
  >|code|int|状态码|
  >|message|string|状态信息|
  >|data|map|返回数据|