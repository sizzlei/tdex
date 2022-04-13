# TDEX
Table Definition Exporter(이하 tdex) 는 간단한 테이블 명세서 추출 툴로서 필요시 원하는 데이터 베이스 및 테이블을 선택하여 테이블 명세서를 추출할 수 있는 툴입니다. 

## Usge Packet
```
go get github.com/360EntSecGroup-Skylar/excelize
go get github.com/go-sql-driver/mysql
```

## Usage
---
```
$ ./tdex
Host: localhost
Port(3306):     
User: root
Pass: *********
Database: nft_uat
Table name Like Condition: 
FilePath: ./
Table Summary Export Success.
```
>컴파일된 툴은 Max 지원입니다. 
- Host : 접근 DB Host를 입력합니다. 
- Port : 기본 포트 3306 으로 지정됩니다.
- User : DB 접근 계정
- Pass : DB 접근 패스워드 ( 평문으로 입력되니 유의!)
- Database : 추출 대상 스키마
- Table name Like Condition : 테이블 선택적으로 추출하기 위한 항목, 공백인 경우 전체 테이블에 대해 추출 합니다. 
    + tablename% : Wildcard(%)가 포함된 경우 Like 검색으로 테이블 추출합니다. (1개만 가능)
    ```
    Table name Like Condition: user_%
    ```
    + tablename,tablename : 복수 테이블을 작성하는 경우 쉼표로 구분합니다. 
    ```
    Table name Like Condition: users,assets
    ```
    + tablename : 단일 테이블 추출시 해당 테이블 명을 작성합니다. 
    ```
    Table name Like Condition: logs
    ```
- FilePath: 추출된 파일이 저장될 경로를 작성합니다.


## Export
---
![추출](/img/tdex_export_excel.png)