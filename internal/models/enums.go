package models

// Status enums
type StatusUmum string

const (
	StatusAktif      StatusUmum = "Aktif"
	StatusTidakAktif StatusUmum = "Tidak Aktif"
)

type ProductUsageStatus string

const (
	ProductUsageJual        ProductUsageStatus = "Jual"
	ProductUsagePakaiSendiri ProductUsageStatus = "Pakai Sendiri"
	ProductUsageRusak       ProductUsageStatus = "Rusak"
)

type SNStatus string

const (
	SNStatusTersedia SNStatus = "Tersedia"
	SNStatusTerpakai SNStatus = "Terpakai"
	SNStatusRusak    SNStatus = "Rusak"
)

type ServiceStatusEnum string

const (
	ServiceStatusAntri     ServiceStatusEnum = "Antri"
	ServiceStatusDikerjakan ServiceStatusEnum = "Dikerjakan"
	ServiceStatusSelesai   ServiceStatusEnum = "Selesai"
	ServiceStatusDiambil   ServiceStatusEnum = "Diambil"
	ServiceStatusKomplain  ServiceStatusEnum = "Komplain"
)

type TransactionStatus string

const (
	TransactionStatusPending TransactionStatus = "pending"
	TransactionStatusSukses  TransactionStatus = "sukses"
	TransactionStatusGagal   TransactionStatus = "gagal"
)

type PurchaseStatus string

const (
	PurchaseStatusSelesai PurchaseStatus = "Selesai"
	PurchaseStatusPending PurchaseStatus = "Pending"
)

type PaymentTypeEnum string

const (
	PaymentTypeTunai    PaymentTypeEnum = "tunai"
	PaymentTypeTransfer PaymentTypeEnum = "transfer"
	PaymentTypeCicilan  PaymentTypeEnum = "cicilan"
)

type APARStatus string

const (
	APARStatusBelumLunas APARStatus = "Belum Lunas"
	APARStatusLunas      APARStatus = "Lunas"
)

type CashFlowType string

const (
	CashFlowTypePemasukan   CashFlowType = "Pemasukan"
	CashFlowTypePengeluaran CashFlowType = "Pengeluaran"
)

type ReportTypeEnum string

const (
	ReportTypePenjualan ReportTypeEnum = "Penjualan"
	ReportTypeKeuangan  ReportTypeEnum = "Keuangan"
	ReportTypeInventory ReportTypeEnum = "Inventory"
)

type ReportStatus string

const (
	ReportStatusPending ReportStatus = "Pending"
	ReportStatusSelesai ReportStatus = "Selesai"
	ReportStatusGagal   ReportStatus = "Gagal"
)

type PromotionType string

const (
	PromotionTypePercentage PromotionType = "percentage"
	PromotionTypeFixed      PromotionType = "fixed"
)