package models

import (
	"time"

	"gorm.io/gorm"
)

// ServiceJobs table
type ServiceJob struct {
	ServiceJobID               uint              `gorm:"primaryKey;autoIncrement" json:"service_job_id"`
	ServiceCode                string            `gorm:"size:50;unique;not null" json:"service_code"`
	QueueNumber                int               `gorm:"not null" json:"queue_number"`
	CustomerID                 uint              `gorm:"not null;index" json:"customer_id"`
	VehicleID                  uint              `gorm:"not null;index" json:"vehicle_id"`
	TechnicianID               *uint             `gorm:"index" json:"technician_id"`
	ReceivedByUserID           uint              `gorm:"not null;index" json:"received_by_user_id"`
	OutletID                   uint              `gorm:"not null;index" json:"outlet_id"`
	ProblemDescription         string            `gorm:"type:text;not null" json:"problem_description"`
	TechnicianNotes            *string           `gorm:"type:text" json:"technician_notes"`
	Status                     ServiceStatusEnum `gorm:"not null" json:"status"`
	ServiceInDate              time.Time         `gorm:"not null" json:"service_in_date"`
	PickedUpDate               *time.Time        `json:"picked_up_date"`
	ComplainDate               *time.Time        `json:"complain_date"`
	WarrantyExpiresAt          *time.Time        `gorm:"type:date" json:"warranty_expires_at"`
	NextServiceReminderDate    *time.Time        `gorm:"type:date" json:"next_service_reminder_date"`
	DownPayment                float64           `gorm:"type:decimal(15,2);default:0" json:"down_payment"`
	GrandTotal                 float64           `gorm:"type:decimal(15,2);default:0" json:"grand_total"`
	TechnicianCommission       float64           `gorm:"type:decimal(15,2);default:0" json:"technician_commission"`
	ShopProfit                 float64           `gorm:"type:decimal(15,2);default:0" json:"shop_profit"`
	CreatedAt                  time.Time         `json:"created_at"`
	UpdatedAt                  time.Time         `json:"updated_at"`
	DeletedAt                  gorm.DeletedAt    `gorm:"index" json:"deleted_at"`
	CreatedBy                  *uint             `json:"created_by"`

	// Relationships
	Customer       *Customer       `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	Vehicle        *CustomerVehicle `gorm:"foreignKey:VehicleID" json:"vehicle,omitempty"`
	Technician     *User           `gorm:"foreignKey:TechnicianID" json:"technician,omitempty"`
	ReceivedByUser *User           `gorm:"foreignKey:ReceivedByUserID" json:"received_by_user,omitempty"`
	Outlet         *Outlet         `gorm:"foreignKey:OutletID" json:"outlet,omitempty"`
	ServiceDetails []ServiceDetail `gorm:"foreignKey:ServiceJobID" json:"service_details,omitempty"`
	Histories      []ServiceJobHistory `gorm:"foreignKey:ServiceJobID" json:"histories,omitempty"`
}

// ServiceDetails table
type ServiceDetail struct {
	DetailID        uint    `gorm:"primaryKey;autoIncrement" json:"detail_id"`
	ServiceJobID    uint    `gorm:"not null;index" json:"service_job_id"`
	ItemID          uint    `gorm:"not null" json:"item_id"`
	ItemType        string  `gorm:"not null" json:"item_type"`
	Description     string  `gorm:"size:255;not null" json:"description"`
	SerialNumberUsed *string `gorm:"size:255" json:"serial_number_used"`
	Quantity        int     `gorm:"not null" json:"quantity"`
	PricePerItem    float64 `gorm:"type:decimal(15,2);not null" json:"price_per_item"`
	CostPerItem     float64 `gorm:"type:decimal(15,2);not null" json:"cost_per_item"`

	// Relationships
	ServiceJob *ServiceJob `gorm:"foreignKey:ServiceJobID" json:"service_job,omitempty"`
}

// ServiceJobHistories table
type ServiceJobHistory struct {
	HistoryID    uint       `gorm:"primaryKey;autoIncrement" json:"history_id"`
	ServiceJobID uint       `gorm:"not null;index" json:"service_job_id"`
	UserID       uint       `gorm:"not null;index" json:"user_id"`
	Notes        *string    `gorm:"type:text" json:"notes"`
	ChangedAt    time.Time  `json:"changed_at"`

	// Relationships
	ServiceJob *ServiceJob `gorm:"foreignKey:ServiceJobID" json:"service_job,omitempty"`
	User       *User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
}