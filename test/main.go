package main

import (
	"fmt"
	"github.com/civet148/html2pdf"
)

type Signature struct {
	Title     string //签名标题（例如：安检签字、生产领导签字等等...）
	ImagePath string //签字图片路径
}

type ProductionCap struct {
	ProductName       string //产品名称
	ProductNo         string //产品编码
	ProductSpecsValue string //产品规格值
	OutNum            string //领货件数
	OutWeight         string //领货重量
	InNum             string //交货件数
	InWeight          string //交货重量
	QualifiedNum      string //交货合格数
	QualifiedWeight   string //合格重量
	QualifiedRate     string //交货合格率
	WasteNum          string //交货废品数
	WasteRate         string //交货废品率
	NumDiff           string //数量差
	WeightDiff        string //重量差
	BackNum           string //返车间数量
	BackWeight        string //返车间重量
}

type WorkshopBill struct {
	BillDate         string           //账单日期
	CreateName       string           //创建人真实姓名（制单人）
	CreateTime       string           //创建时间
	FlowNo           string           //账单流水号
	UserName         string           //结账人姓名
	DepartName       string           //车间名称
	CheckoutMethod   string           //结账方式
	GoldWastage      float64          //金计损流量
	StandardWastage  float64          //金标准损耗
	ActualWastage    float64          //金实际损耗
	OverWastage      float64          //金超损
	KGoldWastage     float64          //K金计损流量
	KStandardWastage float64          //K金标准损耗
	KOverWastage     float64          //K金超损
	InNum            string           //总收货数量
	InWeight         string           //总收货重量
	OutNum           string           //总发货数量
	OutWeight        string           //总发货重量
	QualifiedNum     string           //收货合格数
	QualifiedRate    string           //收货合格率
	WasteNum         string           //收货废品数
	WasteRate        string           //收货废品率
	NumDiff          string           //数量差
	WeightDiff       string           //重量差
	DepositNum       string           //寄存数量
	DepositWeight    string           //寄存重量
	Signatures       []*Signature     //签名数据（标题和图片路径）
	ProductionCaps   []*ProductionCap //产能明细数据
}

func main() {
	var err error
	var templatePath = "workshop_bill.html"

	var bill = &WorkshopBill{
		BillDate:         "2025-03-07",
		CreateName:       "张三",
		CreateTime:       "2025-03-07 16:58:03",
		FlowNo:           "2025030700001",
		UserName:         "李四",
		DepartName:       "油压组",
		CheckoutMethod:   "车间",
		GoldWastage:      0,
		StandardWastage:  0,
		ActualWastage:    0,
		OverWastage:      0,
		KGoldWastage:     0,
		KStandardWastage: 0,
		KOverWastage:     0,
		InNum:            "2699.13",
		InWeight:         "2",
		OutNum:           "3",
		OutWeight:        "4",
		QualifiedNum:     "5",
		QualifiedRate:    "6",
		WasteNum:         "7",
		WasteRate:        "8",
		NumDiff:          "9",
		WeightDiff:       "10",
		DepositNum:       "11",
		DepositWeight:    "12",
		Signatures: []*Signature{
			{
				Title:     "安检签字",
				ImagePath: "https://p26.toutiaoimg.com/origin/pgc-image/69042025ec3140b1b59a4f2f998d9ef0",
			},
		},
		ProductionCaps: []*ProductionCap{
			{
				ProductName:       "产品名称",
				ProductNo:         "产品编码",
				ProductSpecsValue: "产品规格值",
				OutNum:            "11",
				OutWeight:         "22",
				InNum:             "33",
				InWeight:          "44",
				QualifiedNum:      "55",
				QualifiedWeight:   "66",
				QualifiedRate:     "99.9%",
				WasteNum:          "1",
				WasteRate:         "0.1%",
				NumDiff:           "3",
				WeightDiff:        "4",
				BackNum:           "5",
				BackWeight:        "6",
			},
			{
				ProductName:       "产品名称",
				ProductNo:         "产品编码",
				ProductSpecsValue: "产品规格值",
				OutNum:            "11",
				OutWeight:         "22",
				InNum:             "33",
				InWeight:          "44",
				QualifiedNum:      "55",
				QualifiedWeight:   "66",
				QualifiedRate:     "99.9%",
				WasteNum:          "1",
				WasteRate:         "0.1%",
				NumDiff:           "3",
				WeightDiff:        "4",
				BackNum:           "5",
				BackWeight:        "6",
			},
		},
	}

	generator := html2pdf.NewHtml2PDF()
	err = generator.GenerateWithFile(templatePath, "test.pdf", bill)
	if err != nil {
		fmt.Printf("generate PDF error：%s\n", err)
		return
	}
}

