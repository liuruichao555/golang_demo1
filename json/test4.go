package main

import (
	"fmt"
	"encoding/json"
)

type FoodInformation struct{
	FoodName						string
	Manufacturer					string
}

//the struct of the product
type ProduceInformation struct {
	Address         string
	ProductionBatch string
	ProductionTime  int64
	ValidDate       int64
}

type DrugInformation struct {
	DrugName            string
	ApprovalNumber      string
	Size                string
	Form                string
	Manufacturer        string
	NDCNumber           string
	NDCNumberRemark     string
	MedicineInstruction string
}

type Goods struct{
	Type 							string
	GoodsType						string
	ParentCode						string
	ProduceInfoId					string
	ProductId						string
	ProductCode						string
	UniqueCode						string
	CommodityCode					string
	ProductionBatch					string
	DrugInformation				    []DrugInformation
	FoodInformation				    []FoodInformation
	ProduceInformation			    []ProduceInformation
}

func main() {
	jsonStr := `{"Id":"Id","Type":"Type","GoodsType":"GoodsType","ParentCode":"ParentCode","ProduceInfoId":"ProduceInfoId","ProductCode":"ProductCode","UniqueCode":"UniqueCode","CommodityCode":"CommodityCode","ProductionBatch":"ProductionBatch3","DrugInformation":[{"DrugName":"heheda","ApprovalNumber":"ApprovalNumber","Size":"Size","Form":"Form","Manufacturer":"Manufacturer","NDCNumber":"NDCNumber","NDCNumberRemark":"NDCNumberRemark","MedicineInstruction":"MedicineInstruction"}],"FoodInformation":[{"Id":0,"FoodName":"FoodName","Manufacturer":"Manufacturer"}],"ProduceInformation":[{"Id":0,"Address":"Address","ProductionBatch":"ProductionBatch","ProductionTime":12345678,"ValidDate":12345678}]}`
	//jsonStr := "{'Id':'Id','Type':'Type','GoodsType':'GoodsType','ParentCode':'ParentCode','ProduceInfoId':'ProduceInfoId','ProductCode':'ProductCode','UniqueCode':'UniqueCode','CommodityCode':'CommodityCode','ProductionBatch':'ProductionBatch3','DrugInformation':[{'DrugName':'DrugName','ApprovalNumber':'ApprovalNumber','Size':'Size','Form':'Form','Manufacturer':'Manufacturer','NDCNumber':'NDCNumber','NDCNumberRemark':'NDCNumberRemark','MedicineInstruction':'MedicineInstruction'}],'FoodInformation':[{'Id':0,'FoodName':'FoodName','Manufacturer':'Manufacturer'}],'ProduceInformation':[{'Id':0,'Address':'Address','ProductionBatch':'ProductionBatch','ProductionTime':12345678,'ValidDate':12345678}]}"
	goods := &Goods{}
	err := json.Unmarshal([]byte(jsonStr), &goods)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", goods)
}
