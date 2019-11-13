--订单主表
CREATE TABLE [dbo].[SEOrder](
    [FInterID] [INT] identity(1,1) PRIMARY KEY,   --内码
    [FBillNo] [nvarchar](255) NOT NULL, --单据号
    [FCurrencyID] [int] NULL,--币别
    [FCustID] [int] NULL,  --客户
    [FDate] [datetime] NULL,--日期
    [FBillerID] [int] NULL,--制单人
    [FCheckerID] [int] NULL,--审核人
    [FNote] [varchar](30) NULL,--备注
    [FStatus] [smallint] NOT NULL--是否审核
    )


CREATE TABLE [dbo].[SEOrderEntry](
    [FInterID] [INT] NOT NULL,   --内码
    [FEntryID] [int] NOT NULL,  --行号
    [FItemID] [int] NULL, --物料
    [FQty] [decimal](28, 10) NOT NULL, --数量
    [FPrice] [decimal](28, 10) NOT NULL,--单价
    [FAmount] [decimal](20, 2) NOT NULL, --金额
    )