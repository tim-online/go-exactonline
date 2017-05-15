package logistics

import "github.com/tim-online/go-exactonline/edm"

type Items []Item

type Item struct {
	ID                      edm.GUID     `json:"ID"`
	Barcode                 edm.String   `json:"Barcode"`
	Class01                 edm.String   `json:"Class_01"`
	Class02                 edm.String   `json:"Class_02"`
	Class03                 edm.String   `json:"Class_03"`
	Class04                 edm.String   `json:"Class_04"`
	Class05                 edm.String   `json:"Class_05"`
	Class06                 edm.String   `json:"Class_06"`
	Class07                 edm.String   `json:"Class_07"`
	Class08                 edm.String   `json:"Class_08"`
	Class09                 edm.String   `json:"Class_09"`
	Class10                 edm.String   `json:"Class_10"`
	Code                    edm.String   `json:"Code"`
	CopyRemarks             edm.Byte     `json:"CopyRemarks"`
	CostPriceCurrency       edm.String   `json:"CostPriceCurrency"`
	CostPriceNew            edm.Double   `json:"CostPriceNew"`
	CostPriceStandard       edm.Double   `json:"CostPriceStandard"`
	Created                 edm.DateTime `json:"Created"`
	Creator                 edm.GUID     `json:"Creator"`
	CreatorFullName         edm.String   `json:"CreatorFullName"`
	Description             edm.String   `json:"Description"`
	Division                edm.Int32    `json:"Division"`
	EndDate                 edm.DateTime `json:"EndDate"`
	ExtraDescription        edm.String   `json:"ExtraDescription"`
	FreeBoolField01         edm.Boolean  `json:"FreeBoolField_01"`
	FreeBoolField02         edm.Boolean  `json:"FreeBoolField_02"`
	FreeBoolField03         edm.Boolean  `json:"FreeBoolField_03"`
	FreeBoolField04         edm.Boolean  `json:"FreeBoolField_04"`
	FreeBoolField05         edm.Boolean  `json:"FreeBoolField_05"`
	FreeDateField01         edm.DateTime `json:"FreeDateField_01"`
	FreeDateField02         edm.DateTime `json:"FreeDateField_02"`
	FreeDateField03         edm.DateTime `json:"FreeDateField_03"`
	FreeDateField04         edm.DateTime `json:"FreeDateField_04"`
	FreeDateField05         edm.DateTime `json:"FreeDateField_05"`
	FreeNumberField01       edm.Double   `json:"FreeNumberField_01"`
	FreeNumberField02       edm.Double   `json:"FreeNumberField_02"`
	FreeNumberField03       edm.Double   `json:"FreeNumberField_03"`
	FreeNumberField04       edm.Double   `json:"FreeNumberField_04"`
	FreeNumberField05       edm.Double   `json:"FreeNumberField_05"`
	FreeNumberField06       edm.Double   `json:"FreeNumberField_06"`
	FreeNumberField07       edm.Double   `json:"FreeNumberField_07"`
	FreeNumberField08       edm.Double   `json:"FreeNumberField_08"`
	FreeTextField01         edm.String   `json:"FreeTextField_01"`
	FreeTextField02         edm.String   `json:"FreeTextField_02"`
	FreeTextField03         edm.String   `json:"FreeTextField_03"`
	FreeTextField04         edm.String   `json:"FreeTextField_04"`
	FreeTextField05         edm.String   `json:"FreeTextField_05"`
	FreeTextField06         edm.String   `json:"FreeTextField_06"`
	FreeTextField07         edm.String   `json:"FreeTextField_07"`
	FreeTextField08         edm.String   `json:"FreeTextField_08"`
	FreeTextField09         edm.String   `json:"FreeTextField_09"`
	FreeTextField10         edm.String   `json:"FreeTextField_10"`
	GLCosts                 edm.GUID     `json:"GLCosts"`
	GLCostsCode             edm.String   `json:"GLCostsCode"`
	GLCostsDescription      edm.String   `json:"GLCostsDescription"`
	GLRevenue               edm.GUID     `json:"GLRevenue"`
	GLRevenueCode           edm.String   `json:"GLRevenueCode"`
	GLRevenueDescription    edm.String   `json:"GLRevenueDescription"`
	GLStock                 edm.GUID     `json:"GLStock"`
	GLStockCode             edm.String   `json:"GLStockCode"`
	GLStockDescription      edm.String   `json:"GLStockDescription"`
	GrossWeight             edm.Double   `json:"GrossWeight"`
	IsBatchItem             edm.Byte     `json:"IsBatchItem"`
	IsBatchNumberItem       edm.Byte     `json:"IsBatchNumberItem"`
	IsFractionAllowedItem   edm.Boolean  `json:"IsFractionAllowedItem"`
	IsMakeItem              edm.Byte     `json:"IsMakeItem"`
	IsNewContract           edm.Byte     `json:"IsNewContract"`
	IsOnDemandItem          edm.Byte     `json:"IsOnDemandItem"`
	IsPackageItem           edm.Boolean  `json:"IsPackageItem"`
	IsPurchaseItem          edm.Boolean  `json:"IsPurchaseItem"`
	IsRegistrationCodeItem  edm.Byte     `json:"IsRegistrationCodeItem"`
	IsSalesItem             edm.Boolean  `json:"IsSalesItem"`
	IsSerialItem            edm.Boolean  `json:"IsSerialItem"`
	IsSerialNumberItem      edm.Boolean  `json:"IsSerialNumberItem"`
	IsStockItem             edm.Boolean  `json:"IsStockItem"`
	IsSubcontractedItem     edm.Boolean  `json:"IsSubcontractedItem"`
	IsTaxableItem           edm.Byte     `json:"IsTaxableItem"`
	IsTime                  edm.Byte     `json:"IsTime"`
	IsWebshopItem           edm.Byte     `json:"IsWebshopItem"`
	ItemGroup               edm.GUID     `json:"ItemGroup"`
	ItemGroupCode           edm.String   `json:"ItemGroupCode"`
	ItemGroupDescription    edm.String   `json:"ItemGroupDescription"`
	Modified                edm.DateTime `json:"Modified"`
	Modifier                edm.GUID     `json:"Modifier"`
	ModifierFullName        edm.String   `json:"ModifierFullName"`
	NetWeight               edm.Double   `json:"NetWeight"`
	NetWeightUnit           edm.String   `json:"NetWeightUnit"`
	Notes                   edm.String   `json:"Notes"`
	PictureName             edm.String   `json:"PictureName"`
	PictureThumbnailUrl     edm.String   `json:"PictureThumbnailUrl"`
	PictureUrl              edm.String   `json:"PictureUrl"`
	SalesVatCode            edm.String   `json:"SalesVatCode"`
	SalesVatCodeDescription edm.String   `json:"SalesVatCodeDescription"`
	SearchCode              edm.String   `json:"SearchCode"`
	SecurityLevel           edm.Int32    `json:"SecurityLevel"`
	StartDate               edm.DateTime `json:"StartDate"`
	Stock                   edm.Double   `json:"Stock"`
	Unit                    edm.String   `json:"Unit"`
	UnitDescription         edm.String   `json:"UnitDescription"`
	UnitType                UnitType     `json:"UnitType"`
}

const (
	UnitTypeArea   UnitType = "A" // Area
	UnitTypeLength UnitType = "L" // Length
	UnitTypeOther  UnitType = "O" // Other
	UnitTypeTime   UnitType = "T" // Time
	UnitTypeVolume UnitType = "V" // Volume
	UnitTypeWeight UnitType = "W" // Weight
)

type UnitType edm.String

func (u *UnitType) String() string {
	switch string(*u) {
	case "A":
		return "Area"
	case "L":
		return "Length"
	case "O":
		return "Other"
	case "T":
		return "Time"
	case "V":
		return "Volume"
	case "W":
		return "Weight"
	}
	return ""
}
