# Delivery Service API 
### 1.Report Daily
 รายงานสรุปการจัดส่งสินค้าประจำวัน 
 
 > Report Daily for delivery system for check list of DO Document By date filter
 > time process around 10-15 sec 

##### REQUEST 
###### URL : /delivery/report
Content type : application/json
Request payload

```sh
         { "date" : "2018-09-03"}
```

Response 
```json 
        {
        "data": {
                "data": [
                            {
                                "id": 211749,
                                "do_no": "s02-do6109-0038 ",
                                "so_no": "s02-shv6108-1020 ",
                                "confirm_date": "2018-09-03T00:00:00Z",
                                "do_date": "2018-09-03T00:00:00Z",
                                "diff_date": 0,
                                "description": "ตรงเวลา",
                                "ar_name": "คุณธนยงค์ แก้วก่ำ ",
                                "item_amount": 95933.5,
                                "item_group": " ",
                                "remark": " ",
                                "invoice": "s02-ihv6109-0103 ",
                                "car_license": "ผอ.4291 ",
                                "sale_code": "55285 ",
                                "saleman": "พรรณิดา ภูมี "
                                },
                            {
                                "id": 211525,
                                "do_no": "s02-do6109-0054 ",...}]}}
```


#### Report By ProfitCenter 
###### URL : /delivery/report/team
Content type : application/json
Request :

```json
 { 
    "date":"2018-09-01",
    "profit":"S01"
 }
```

Response :
```json 
        {
        "data": {
                "data": [
                            {
                                "id": 211749,
                                "do_no": "s02-do6109-0038 ",
                                "so_no": "s02-shv6108-1020 ",
                                "confirm_date": "2018-09-03T00:00:00Z",
                                "do_date": "2018-09-03T00:00:00Z",
                                "diff_date": 0,
                                "description": "ตรงเวลา",
                                "ar_name": "คุณธนยงค์ แก้วก่ำ ",
                                "item_amount": 95933.5,
                                "item_group": " ",
                                "remark": " ",
                                "invoice": "s02-ihv6109-0103 ",
                                "car_license": "ผอ.4291 ",
                                "sale_code": "55285 ",
                                "saleman": "พรรณิดา ภูมี ",
                                "profit":"S01"
                                },
                            {
                                "id": 211525,
                                "do_no": "s02-do6109-0054 ",...}]}}
```

### List of sale person
###### URL : /delivery/report/sales
Content type : application/json

Method : POST

Request 

```json
{
  "profit_center":"S01"
}
```

Response 

```json
{
    "data": [
      {
        "sale_code": "56187 ",
        "sale_name": "นิธิพิชญ์ พรหมอนัต์ "
        },
          {
        "sale_code": "57217 ",
        "sale_name": "อังคณา คุ้มวานิช "
        },
          {
        "sale_code": "61102 ",
        "sale_name": "วัชรนนท์ ธนโพธิ์สน "
      }
    ]
}
```


### List of ProfitCenter
###### URL : /delivery/report/team
Content type : application/json

Method : POST

Request 

```json

```

Response 

```json
{
    "data": [
          {
           "profit": "H00 "
        },
          {
          "profit": "S01 "
        },
          {
          "profit": "S02 "
        },
          {
        "profit": "S03 "
        },
          {
        "profit": "W01 "
        },
          {
        "profit": "W02 "
        }
    ]
}
```
